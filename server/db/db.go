package db

import (
	"DGraph-Example/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

type DB interface {
	GetUsers() ([]*model.Person, error)
}

type DGraph struct {
	client *dgo.Dgraph
}

func NewDgraph(client *dgo.Dgraph, init bool) DB {
	d := DGraph{client: client}
	if init {
		err := d.initDB()
		if err != nil {
			log.Fatal(err)
		}
	}
	return d
}
func (d DGraph) initDB() error {

	err := d.client.Alter(context.Background(), &api.Operation{
		Schema: `
	name: string @index(exact) .
	age: int .
	married: bool .
	loc: geo .
	dob: datetime .

	type Person {
		name
		age
		dob
		married
		
		
		
	}

	type Loc {
		
	}

	type Institution {
		name
	}
		`,
	})
	if err != nil {
		return err
	}
	ctx := context.Background()
	txn := d.client.NewTxn()
	defer txn.Discard(ctx)

	dob := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)
	// While setting an object if a struct has a Uid then its properties in the graph are updated
	// else a new node is created.
	// In the example below new nodes for Alice, Bob and Charlie and school are created (since they
	// dont have a Uid).
	alice := model.Person{
		UID:     "_:alice",
		Name:    "Alice",
		DType:   []string{"Person"},
		Age:     26,
		Married: true,
		Location: model.Loc{
			Type:   "Point",
			Coords: []float64{1.1, 2},
		},
		Dob: &dob,
		Raw: []byte("raw_bytes"),
		Friends: []model.Person{{
			Name: "Bob",
			Age:  24,
		}, {
			Name: "Charlie",
			Age:  29,
		}},
		School: []model.School{{
			Name: "Crown Public School",
		}},
	}
	pb, err := json.Marshal(alice)
	if err != nil {
		return err
	}
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}

	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": assigned.Uids["alice"]}
	q := `query Me($id: string){
	me(func: uid($id)) {
		name
		dob
		age
		loc
		raw_bytes
		married
		dgraph.type
		friend @filter(eq(name, "Bob")){
			name
			age
			dgraph.type
		}
		school {
			name
			dgraph.type
			}
		}
	}`

	resp, err := d.client.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return err
	}

	type Root struct {
		Me []model.Person `json:"me"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return err
	}
	// fmt.Printf("Me: %+v\n", r.Me)
	// R.Me would be same as the person that we set above.
	fmt.Println("Initalization complete")
	fmt.Println(string(resp.Json))
	return nil
}
func (d DGraph) GetUsers() ([]*model.Person, error) {
	return nil, errors.New("Not implemented")
}
