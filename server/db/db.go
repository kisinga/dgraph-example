package db

import (
	"DGraph-Example/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
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

//@deprecated
//I decided to leave this here since it demostrates how to
//perform some CRUD operation, so it might be useful
//Otherwise, this code is not useful since initialization
//with data is no longer handled within the app
func (d DGraph) initDB() error {
	//The code in the official documentation is throwing an error so I
	//in the meantime I edited out the erroneous lines until I figure out why

	// s := `
	// name: string @index(exact) .
	// age: int .
	// married: bool .
	// loc: geo .
	// dob: datetime .

	// type Person {
	// 	name
	// 	age
	// 	dob
	// 	married
	// 	raw
	// 	friends
	// 	loc
	// 	school
	// }

	// type Loc {
	// 	type
	// 	coords
	// }

	// type Institution {
	// 	name
	// }
	// 	`
	s := `
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
	}`
	err := d.client.Alter(context.Background(), &api.Operation{
		Schema: s,
	})
	if err != nil {
		return err
	}
	ctx := context.Background()
	txn := d.client.NewTxn()
	defer txn.Discard(ctx)
	people, err := parsePeople()
	if err != nil {
		return err
	}
	pb, err := json.Marshal(people)
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
func (d DGraph) SearchActors(phrase string) ([]*model.Person, error) {
	return nil, errors.New("SearchActors Not implemented")
}
func (d DGraph) SearchMovies(phrase string) ([]*model.Person, error) {
	return nil, errors.New("SearchMovies Not implemented")
}

func parsePeople() ([]model.Person, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("../sample/people.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	// we initialize our Users array
	var people []model.Person

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	//This is probably unecessary but helps catch any
	//parsing errors within the json before values are sent over
	//to the database, as well as allow initialization of missing
	//fields with their nil values
	err = json.Unmarshal(byteValue, &people)

	if err != nil {
		return nil, err
	}
	return people, nil
}
