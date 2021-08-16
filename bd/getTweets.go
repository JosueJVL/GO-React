package bd

import (
	"context"
	"log"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(userId string, page int64) ([]*models.ResponseTweets, bool) {
	_context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("tweet")

	var resultado []*models.ResponseTweets

	//se crea el filtro
	filter := bson.M{
		"userId": userId,
	}

	options := options.Find()

	options.SetLimit(20)
	// Key, el valor por el cual se va a ordenar
	// Value, -1 es desend ente
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	inf, err := col.Find(_context, filter, options)
	if err != nil {
		log.Fatal(err)
		return resultado, false
	}

	for inf.Next(context.TODO()) {
		var files models.ResponseTweets
		err := inf.Decode(&files)
		if err != nil {
			return resultado, false
		}

		resultado = append(resultado, &files)
	}

	return resultado, true
}
