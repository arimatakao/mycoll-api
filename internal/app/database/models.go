package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	_id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	links       []string           `json:"links" bson:"links,omitempty"`
	name        string             `json:"name" bson:"name,omitempty"`
	description string             `json:"description" bson:"description,omitempty"`
}
