package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	IdOwner     primitive.ObjectID `json:"id_owner" bson:"_id"`
	Title       string             `json:"title,omitempty" bson:"name,omitempty"`
	Tags        string             `json:"tags,omitempty" bson:"tags,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Links       []string           `json:"links,omitempty" bson:"links,omitempty"`
}

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
}
