package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News model for storage users news
type News struct {
	ID        *bson.ObjectId `json:"id" bson:"_id,omitempty" export:"string" description:"Unique identifier of the news"`
	UserID    *bson.ObjectId `json:"user_id" bson:"user_id,omitempty" export:"string" description:"Identifier of the user"`
	Title     *string        `json:"title" bson:"title,omitempty" export:"string" description:"Title of the news"`
	Content   *string        `json:"content" bson:"content,omitempty" export:"string" description:"Content of the news"`
	CreatedAt *time.Time     `json:"created_at" bson:"created_at,omitempty" export:"string" description:"Time of news creation"`
	UpdatedAt *time.Time     `json:"updated_at" bson:"updated_at,omitempty" export:"string" description:"Last news update time"`
}
