package importer

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Go(ctx context.Context, dbconn, dbname, dir string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbconn))
	if err != nil {
		return err
	}

	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
