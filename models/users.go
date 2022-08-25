package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	UserId            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Email             string             `json:"email,omitempty"`
	Email_verified_at time.Time          `json:"email_verified_at,omitempty"`
	Password          string             `json:"password,omitempty"`
	RememberToken     time.Time          `json:"rememberToken,omitempty"`
	Timestamps        time.Time          `json:"timestamps,omitempty"`
}

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
