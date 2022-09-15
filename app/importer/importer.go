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
	col := client.Database(dbname).Collection("matches")

	for _, file := range files {
		match, err := parseFile(file)
		if err != nil {
			log.Printf("[WARN] failed while parse file '%s': %s", file, err)
			continue
		}

		if err := col.FindOne(ctx, bson.M{"map": match.Map, "type": match.Type, "duration": match.Duration, "datetime": match.Datetime}).Err(); err == nil {
			log.Println("[WARN] match already exists.")
			continue
		}

		_, err = col.InsertOne(ctx, transform(match))
		if err != nil {
			log.Printf("[WARN] failed to insert a document: %s", err)
			continue
		}
	}

	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func transform(match XMLMatch) Match {
	m := Match{
		Map:      match.Map,
		Type:     match.Type,
		Duration: match.Duration,
		Datetime: match.Datetime,
	}

	for _, player := range match.Players {
		var p Player
		p.Name = player.Name

		for _, stat := range player.Stats {
			switch stat.Name {
			case "Score":
				p.Score = stat.Value
			case "Kills":
				p.Kills = uint(stat.Value)
			case "Deaths":
				p.Deaths = uint(stat.Value)
			case "Suicides":
				p.Suicides = uint(stat.Value)
			case "DamageGiven":
				p.DamageGiven = uint(stat.Value)
			case "DamageTaken":
				p.DamageTaken = uint(stat.Value)
			case "HealthTotal":
				p.HealthTotal = uint(stat.Value)
			case "ArmorTotal":
				p.ArmorTotal = uint(stat.Value)
			}
		}

		p.Weapons = player.Weapons
		p.Items = player.Items
		p.Powerups = player.Powerups

		m.Players = append(m.Players, p)
	}

	return m
}
