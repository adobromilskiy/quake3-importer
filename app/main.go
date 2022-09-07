package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var opts struct {
	DbConn string `long:"dbconn" description:"Database connection." required:"true"`
	DbName string `long:"dbname" description:"Database name." required:"true"`
	Path   string `long:"path" description:"Path to the directory where the stats files are located." required:"true"`
}

func main() {
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("[ERROR] failed to parse flags: %v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(opts.DbConn))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	log.Println(opts.Path)
}
