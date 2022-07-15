package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Links       []string           `json:"links" bson:"links,omitempty"`
	Title       string             `json:"title" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}
