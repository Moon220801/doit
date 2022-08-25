package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User defines the user mongo object.
type Daily struct {
	DoilyId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}
