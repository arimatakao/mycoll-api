package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Links       []string           `json:"links" bson:"links,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}
