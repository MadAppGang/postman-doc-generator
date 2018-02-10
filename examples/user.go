package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model for storage user
type User struct {
	ID        *bson.ObjectId `json:"id" bson:"_id,omitempty" export:"string" description:"Unique identifier of the user"`
	Name      *string        `json:"name" bson:"name,omitempty" export:"string" description:"User name"`
	Balance   *int           `json:"balance" bson:"balance,omitempty" export:"number" description:"User balance"`
	CreatedAt *time.Time     `json:"created_at" bson:"created_at,omitempty" export:"string" description:"The time of user creation"`
	UpdatedAt *time.Time     `json:"updated_at" bson:"updated_at,omitempty" export:"string" description:"Last user update time"`
}
