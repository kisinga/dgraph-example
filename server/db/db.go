package db

import (
	"dgraph-example/model"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dolan-in/dgman"
)

type DB interface {
	SearchActors(phrase string) ([]*model.Person, error)
	SearchMovies(phrase string) ([]*model.Person, error)
}

type DGraph struct {
	client *dgo.Dgraph
}

func NewDgraph(client *dgo.Dgraph) DB {
	d := DGraph{client: client}
	return d
}

func (d DGraph) SearchActors(phrase string) ([]*model.Person, error) {
	return nil, errors.New("SearchActors Not implemented")
}
func (d DGraph) SearchMovies(phrase string) ([]*model.Person, error) {

	tx := dgman.NewReadOnlyTxn(d.client)

	movie := []model.Movie{}
	// get node with node type `user` that matches filter
	err := tx.Get(&movie).
		Filter("allofterms(name@en, $1)", phrase). // dgraph filter
		All(1).                                    // returns all predicates, expand on 1 level of edge predicates
		Nodes()                                    // get single node from query
	if err != nil {
		if err == dgman.ErrNodeNotFound {
			// node using the specified filter not found
		}
	}
	// struct will be populated if found
	fmt.Println(movie)
	return nil, errors.New("SearchMovies Not implemented")
}
