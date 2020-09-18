package db

import (
	"context"
	"dgraph-example/model"
	"encoding/json"

	"github.com/dgraph-io/dgo/v200"
)

type DB interface {
	SearchActors(phrase string) ([]model.Actor, error)
	SearchMovies(phrase string) ([]*model.Movie, error)
}

type DGraph struct {
	client *dgo.Dgraph
}

func NewDgraph(client *dgo.Dgraph) DB {
	d := DGraph{client: client}
	return d
}

//SearchActors returns a list of whose name matches the provided phrase
//Together with the other actors that acted each respective movie
func (d DGraph) SearchActors(phrase string) ([]model.Actor, error) {
	txn := d.client.NewTxn()

	var result struct {
		Actors []model.Person
	}
	var q = `
		{
			actors(func: regexp(name@en, /.*` + phrase + `*/i) (first: 3)) 
			@filter(has(actor.film)) @cascade{
				uid
				name@en
				actor.film {
					performance.film{
						uid
						name@en
						genre{
							name@en
						}
						starring{
							performance.actor{
								name@en
								uid
							}
						}
						initial_release_date
					}
				}
			}
		}
	`
	resp, err := txn.Query(context.Background(), q)

	if err != nil {
		return []model.Actor{}, err
	}

	if err := json.Unmarshal(resp.GetJson(), &result); err != nil {
		return []model.Actor{}, err
	}
	actors := func() []model.Actor {
		actors := []model.Actor{}
		for _, person := range result.Actors {
			actors = append(actors, *person.Act())
		}
		return actors
	}()

	return actors, nil
}
func (d DGraph) SearchMovies(phrase string) ([]*model.Movie, error) {

	// tx := dgman.NewReadOnlyTxn(d.client)

	// movies := []*model.Movie{}

	// regex := "regexp(name@en, /.*" + phrase + ".*/i)"
	// // get node with node type `user` that matches filter
	// err := tx.Get(&movies).
	// 	// Filter("has(genre)"). // dgraph filter
	// 	Filter(regex). // dgraph filter
	// 	All(6).        // returns all predicates, expand on 1 level of edge predicates
	// 	First(100).
	// 	Nodes() // get single node from query
	// if err != nil {
	// 	if err == dgman.ErrNodeNotFound {
	// 		return []*model.Movie{}, err
	// 	}
	// }
	// return movies, nil
	panic("NOt implemented")
}
