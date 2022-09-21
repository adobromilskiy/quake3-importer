package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adobromilskiy/quake3-importer/app/importer"
	"github.com/jessevdk/go-flags"
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

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	if err := importer.Go(ctx, opts.DbConn, opts.DbName, opts.Path); err != nil {
		log.Printf("[ERROR] failed to import: %v", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}
