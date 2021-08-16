package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(IdTweet string, UserId string) error {
	_context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(IdTweet)

	filter := bson.M{
		"_id":    objID,
		"userId": UserId,
	}

	_, err := col.DeleteOne(_context, filter)

	return err
}
