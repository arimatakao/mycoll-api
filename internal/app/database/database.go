package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c *Connection) Shutdown() error {
	err := c.links.Database().Client().Disconnect(context.TODO())
	if err != nil {
		return err
	}
	err = c.users.Database().Client().Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) CreateGroupLinks(idOwner, title string, tags []string, description string, links []string) (interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(idOwner)
	if err != nil {
		return nil, nil
	}
	result, err := c.links.InsertOne(context.TODO(),
		bson.D{{Key: "id_owner", Value: objectId},
			{Key: "title", Value: title},
			{Key: "tags", Value: tags},
			{Key: "description", Value: description},
			{Key: "links", Value: links}})
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}

func (c *Connection) FindAllGroupLinksByIdOwner(idOwner string) []GroupLinks {
	objectId, err := primitive.ObjectIDFromHex(idOwner)
	if err != nil {
		return nil
	}
	var groupLinks []GroupLinks
	cursor, err := c.links.Find(context.TODO(), bson.D{{Key: "id_owner", Value: objectId}})
	if err != nil {
		return groupLinks
	}
	if err = cursor.All(context.TODO(), &groupLinks); err != nil {
		return groupLinks
	}
	return groupLinks
}

func (c *Connection) UpdateGroupLinksById(id, title string, tags []string, description string, links []string) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0
	}
	result, err := c.links.UpdateByID(context.TODO(),
		objectId,
		bson.D{
			{Key: "$set",
				Value: bson.D{{Key: "title", Value: title},
					{Key: "tags", Value: tags},
					{Key: "description", Value: description},
					{Key: "links", Value: links}}},
		})
	if err != nil {
		return 0
	}
	return result.ModifiedCount
}

func (c *Connection) DeleteGroupLinksById(id string) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0
	}
	result, err := c.links.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return 0
	}
	return result.DeletedCount
}

func (c *Connection) DeleteAllGroupLinksByIdOwner(idOwner string) int64 {
	objectId, err := primitive.ObjectIDFromHex(idOwner)
	if err != nil {
		return 0
	}
	result, err := c.links.DeleteMany(context.TODO(), bson.D{{Key: "id_owner", Value: objectId}})
	if err != nil {
		return 0
	}
	return result.DeletedCount
}

func (c *Connection) CountGroupLinks() int64 {
	result, err := c.links.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return 0
	}
	return result
}

func (c *Connection) CreateUser(name, passwordHash string) (interface{}, error) {
	result, err := c.users.InsertOne(context.TODO(), bson.D{{Key: "name", Value: name},
		{Key: "password", Value: passwordHash}})
	if err != nil {
		return result.InsertedID, err
	}
	return 0, nil
}

func (c *Connection) IsUserExist(name string) bool {
	res := c.users.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}})
	return res.Err() == nil
}

func (c *Connection) GetUserNamePassword(name string) (string, string) {
	var user User
	err := c.users.FindOne(context.TODO(),
		bson.D{{Key: "name", Value: name}}).
		Decode(&user)
	if err != nil {
		return "", ""
	}
	return user.Name, user.Password
}

func (c *Connection) GetUserId(name string) string {
	var user User
	err := c.users.FindOne(context.TODO(),
		bson.D{{Key: "name", Value: name}}).
		Decode(&user)
	if err != nil {
		return ""
	}
	return user.Id.Hex()
}

func (c *Connection) DeleteUser(id string) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0
	}
	result, err := c.users.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return 0
	}
	return result.DeletedCount
}

func (c *Connection) CountUsers() int64 {
	result, err := c.users.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return 0
	}
	return result
}
