package models

import (
	"doit/common"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User defines the user mongo object.
type ToDoList struct {
	DoitId      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Content     string             `json:"content,omitempty"`
	Type        string             `json:"type,omitempty"`
	Remark      string             `json:"remark,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	UpdateAt    time.Time          `json:"updateAt,omitempty"`
	CompletedAt time.Time          `json:"completedAt,omitempty"`
	StartdateAt common.JsonTime    `json:"startdateAt,omitempty"`
	EnddateAt   common.JsonTime    `json:"enddateAt,omitempty"`
	Timeleft    string             `json:"timeleft,omitempty"`
	Complete    bool               `json:"complete,omitempty"`
	Status      string             `json:"status,omitempty"`
	Useremail   string             `json:"userEmail,omitempty"`
}
