package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message represents the structure of a message
type Message struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID  string             `json:"userId" bson:"userId"`
	Content string             `json:"content" bson:"content"`
	Created int64              `json:"created" bson:"created"`
}
