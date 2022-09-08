package importer

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Go(ctx context.Context, dbconn, dbname, dir string) error {
	files, err := getFiles(dir)
	if err != nil {
		return err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbconn))
	if err != nil {
		return err
	}
	col := client.Database(dbname).Collection("match")

	for _, file := range files {
		match, err := parseFile(file)
		if err != nil {
			log.Printf("[ERROR] failed while parse file: %s", err)
			continue
		}

		if err := col.FindOne(ctx, bson.M{"map": match.Map, "type": match.Type, "duration": match.Duration, "datetime": match.Datetime}).Err(); err == nil {
			log.Println("[INFO] match already exists.")
			continue
		}

		result, err := col.InsertOne(ctx, match)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("[INFO] inserted match: %s", result.InsertedID)
	}

	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
