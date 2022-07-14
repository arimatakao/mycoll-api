package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	links *mongo.Collection
	users *mongo.Collection
}

func NewConnection(ctx context.Context, uri string) *Connection {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	db := &Connection{
		links: client.Database("Mycoll").Collection("Links"),
		users: client.Database("Mycoll").Collection("Users"),
	}
	return db
}

func (c *Connection) CreateLinks(l Links) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	result, err := c.links.InsertOne(ctx, l)
	if err != nil {
		log.Println("Cannot create links")
	}
	log.Println(result.InsertedID)
}

func (c *Connection) ReadAllLinks() string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cursor, err := c.links.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Links not found")
	}

	var links []bson.M
	if err = cursor.All(ctx, &links); err != nil {
		log.Println("Links not Found")
	}
	log.Println(links)
	return " "
}

func (c *Connection) UpdateAllLinks() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	result, err := c.links.UpdateMany(ctx, "123", "123")
	if err != nil {
		log.Println("Links not updated")
	}
	log.Printf("Updated %v  Documents", result.ModifiedCount)
}

func (c *Connection) DeleteAllLinks() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	result, err := c.links.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("Links not be deleted")
	}
	log.Printf("Delete %v Documents", result.DeletedCount)
}
