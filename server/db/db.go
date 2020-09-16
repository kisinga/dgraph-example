package db

import (
	"dgraph-example/model"
	"errors"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dolan-in/dgman"
)

type DB interface {
	SearchActors(phrase string) ([]*model.Person, error)
	SearchMovies(phrase string) ([]*model.Movie, error)
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
func (d DGraph) SearchMovies(phrase string) ([]*model.Movie, error) {

	tx := dgman.NewReadOnlyTxn(d.client)

	movies := []*model.Movie{}

	regex := "regexp(name@en, /.*" + phrase + "*/i)"
	// get node with node type `user` that matches filter
	err := tx.Get(&movies).
		Filter(regex). // dgraph filter
		All(1).        // returns all predicates, expand on 1 level of edge predicates
		Nodes()        // get single node from query
	if err != nil {
		if err == dgman.ErrNodeNotFound {
			return []*model.Movie{}, nil
		}
	}
	return movies, nil
}
