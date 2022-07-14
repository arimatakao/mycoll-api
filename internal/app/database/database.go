package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	db *mongo.Database
}

func NewClient(ctx context.Context, uri string) Database {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return Database{
		db: client.Database("Mycoll"),
	}
}

func (d *Database) CreateLinks(l Links) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	links := d.db.Collection("Links")

	result, err := links.InsertOne(ctx, l)
	if err != nil {
		log.Println("Cannot create links")
	}
	log.Println(result.InsertedID)
}
