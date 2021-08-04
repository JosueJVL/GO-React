package bd

import (
	"context"
	"log"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(model models.User) (string, bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("users")

	model.Password, _ = Encrypt(model.Password)
	model.Id = primitive.NewObjectID()

	result, err := col.InsertOne(context, model)

	if err != nil {
		log.Println("Error")
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	log.Println("INSERT CORRECTO" + ObjID.String())
	return ObjID.String(), true, nil
}
