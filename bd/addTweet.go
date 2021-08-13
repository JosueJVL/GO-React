package bd

import (
	"context"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTweet(model models.Tweet) (string, bool, error) {
	// Se crea el contexto y cancel para el TimeOut
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Tipo SaveChange
	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("tweet")

	regitro := bson.M{
		"userId":  model.UserId,
		"message": model.Message,
		"date":    model.Date,
	}

	result, error := col.InsertOne(context, regitro)

	if error != nil {
		return "", false, error
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
