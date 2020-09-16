package main

import (
	"dgraph-example/config"
	"dgraph-example/db"
	"dgraph-example/web"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

var prod bool = false

func main() {
	cfg, err := config.ReadYaml("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("prod") == "true" {
		prod = true
	}
	client, conn := newClient(cfg.DBConfig.URI)
	defer conn.Close()
	db := db.NewDgraph(client)
	// CORS is enabled only in prod profile
	app := web.NewApp(db, prod)
	err = app.Serve()
	log.Println("Error", err)
}
func newClient(url string) (*dgo.Dgraph, *grpc.ClientConn) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	d, err := grpc.Dial(url, dialOpts...)
	if err != nil {
		log.Fatal(err)
	}
	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	), d
}
