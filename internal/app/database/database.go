package database

import (
	"context"
	"log"

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

	return &Connection{
		links: client.Database("Mycoll").Collection("Links"),
		users: client.Database("Mycoll").Collection("Users"),
	}
}

func (c *Connection) CreateLinks(l *Links) bool {
	_, err := c.links.InsertOne(context.TODO(), l)
	return err == nil
}

func (c *Connection) FindLinks(l *Links) *Links {
	var result *Links
	err := c.links.FindOne(context.TODO(), l).Decode(result)
	if err != nil {
		return nil
	}
	return result
}

func (c *Connection) UpdateLinks(filter bson.D, l *Links) bool {
	_, err := c.links.UpdateOne(context.TODO(), filter, l)
	return err == nil
}

func (c *Connection) DeleteLinks(l *Links) bool {
	_, err := c.links.DeleteOne(context.TODO(), l)
	return err == nil
}

func (c *Connection) CreateUser(u *User) bool {
	_, err := c.links.InsertOne(context.TODO(), u)
	return err == nil
}

func (c *Connection) FindUser(u *User) *User {
	var result *User
	err := c.users.FindOne(context.TODO(), u).Decode(result)
	if err != nil {
		return nil
	}
	return result
}

func (c *Connection) DeleteUser(u *User) bool {
	_, err := c.users.DeleteOne(context.TODO(), u)
	return err == nil
}
