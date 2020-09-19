package db

import (
	"context"
	"dgraph-example/model"
	"encoding/json"

	"github.com/dgraph-io/dgo/v200"
)

type DB interface {
	SearchActors(phrase string) ([]model.Actor, error)
	SearchMovies(phrase string) ([]model.Film, error)
}

type DGraph struct {
	client *dgo.Dgraph
}

func NewDgraph(client *dgo.Dgraph) DB {
	d := DGraph{client: client}
	return d
}

//SearchActors returns a list of actors whose name matches the provided phrase
//Together with the other actors that acted each respective movie
func (d DGraph) SearchActors(phrase string) ([]model.Actor, error) {
	txn := d.client.NewTxn()

	var result struct {
		Actors []model.Person
	}
	var q = `
		{
			actors(func: regexp(name@en, /.*` + phrase + `.*/i), first: 100)

			##remove results that are not actors

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

						## Traverse the director node in reverse

						~director.film{
							name@en
							uid
						}
						
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

//SearchMovies returns a list of movies whose name matches the provided phrase
func (d DGraph) SearchMovies(phrase string) ([]model.Film, error) {

	txn := d.client.NewTxn()

	var result struct {
		Movies []model.Movie
	}
	var q = `
		{
			movies(func: regexp(name@en, /.*` + phrase + `.*/i), first: 100)
			## remove results that are not movies
			@filter(has(genre)) @cascade{
				uid
				name@en
				genre{
					name@en
				}
				initial_release_date
				starring{
					performance.actor{
							name@en
							uid
						}
				}
				~director.film{
					name@en
					uid
				}
			}
		}
	`
	resp, err := txn.Query(context.Background(), q)

	if err != nil {
		return []model.Film{}, err
	}
	if err := json.Unmarshal(resp.GetJson(), &result); err != nil {
		return []model.Film{}, err
	}
	films := func() []model.Film {
		films := []model.Film{}
		for _, movie := range result.Movies {
			films = append(films, *movie.ConvertFilm())
		}
		return films
	}()

	return films, nil
}
