package db

import (
	"DGraph-Example/model"
	"errors"

	"github.com/dgraph-io/dgo"
)

type DB interface {
	GetUsers() ([]*model.User, error)
}

type DGraph struct {
	client *dgo.Dgraph
}

func NewDgraph(client *dgo.Dgraph) DB {
	return DGraph{client: client}
}

func (d DGraph) GetUsers() ([]*model.User, error) {
	return nil, errors.New("Not implemented")
}
