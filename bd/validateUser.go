package bd

import (
	"context"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUser(email string) (models.User, bool, string) {

	// Se crea el contexto y cancel para el TimeOut
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Tipo SaveChange
	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("users")

	// M -- regresa un formato json mapeado
	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(context, condition).Decode(&result)
	Id := result.Id.Hex()

	if err != nil {
		return result, false, Id
	}

	return result, true, Id
}
