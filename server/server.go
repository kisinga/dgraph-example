package main

import (
	"DGraph-Example/db"
	"DGraph-Example/web"
	"log"
	"os"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

var prod bool = false

func main() {

	if os.Getenv("prod") == "true" {
		prod = true
	}
	client := newClient("localhost:9080")

	db := db.NewDgraph(client)
	// CORS is enabled only in prod profile
	cors := os.Getenv("prod") != "true"
	app := web.NewApp(db, cors)
	err := app.Serve()
	log.Println("Error", err)
}
func newClient(url string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	d, err := grpc.Dial(url, dialOpts...)
	defer d.Close()
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
