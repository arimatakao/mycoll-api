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

func (c *Connection) CreateLinks(idOwner, title string, links []string, description string) (int, error) {
	return 0, nil
}

func (c *Connection) FindLinksById(id string) (int, error) {
	return 0, nil
}

func (c *Connection) FindLinksByIdOwner(idOwner string) (int, error) {
	return 0, nil
}

func (c *Connection) UpdateLinksById(id int) (int, error) {
	return 0, nil
}

func (c *Connection) DeleteLinksById(id int) (int, error) {
	return 0, nil
}

func (c *Connection) CreateUser(name, password string) (interface{}, error) {
	result, err := c.users.InsertOne(context.TODO(), bson.D{{"name", name}, {"password", password}})
	if err != nil {
		return result.InsertedID, err
	}
	return 0, nil
}

func (c *Connection) IsUserExist(name string) bool {
	res := c.links.FindOne(context.TODO(), bson.D{{"name", name}})
	return res.Err() != nil
}

func (c *Connection) GetUserNamePassword(name, password string) (string, string) {
	var user User
	err := c.users.FindOne(context.TODO(), bson.D{{"name", name}, {"password", password}}).Decode(user)
	log.Println("User: ", user.Name, "Password: ", user.Password)
	if err != nil {
		return "", ""
	}
	return user.Name, user.Password
}

func (c *Connection) DeleteUser(name, password string) (int, error) {
	return 0, nil
}
