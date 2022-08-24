package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User defines the user mongo object.
type ToDoList struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Content     string             `json:"content,omitempty"`
	Type        string             `json:"type,omitempty"`
	Remark      string             `json:"remark,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	UpdateAt    time.Time          `json:"update_at,omitempty"`
	CompletedAt time.Time          `json:"completed_at,omitempty"`
	StartdateAt time.Time          `json:"startdate_at,omitempty"`
	EnddateAt   time.Time          `json:"enddate_at,omitempty"`
	Status      bool               `json:"status,omitempty"`
}
