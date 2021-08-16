package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseTweets struct {
	Id      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
