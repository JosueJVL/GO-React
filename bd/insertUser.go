package bd

import (
	"context"
	"log"
	"time"

	"github.com/JosueJVL/GO-React/models"
)

func InsertUser(model models.User) (string, bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("users")

	model.Password, _ = Encrypt(model.Password)

	result, err := col.InsertOne(context, model)

	if result != nil {
		log.Println("ERROR")
		return "", false, err
	}

	log.Println("INSERT CORRECTO")
	return "", true, nil
}
