package bd

import (
	"context"
	"time"

	"github.com/JosueJVL/GO-React/models"
)

func AddRelation(model models.Relation) (bool, error) {
	// Se crea el contexto y cancel para el TimeOut
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Tipo SaveChange
	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("relation")

	_, err := col.InsertOne(context, model)

	if err != nil {
		return false, err
	}

	return true, err
}
