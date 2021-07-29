package bd

import (
	"context"
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

	result, err := col.InsertOne(context, model)
	if result != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil
}
