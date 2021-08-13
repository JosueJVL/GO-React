package bd

import (
	"context"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(user models.User, UserId string) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	dataBase := MongoConection.Database("twitter")
	collection := dataBase.Collection("users")

	register := make(map[string]interface{})
	//Se crea una interfaz para actualizar con Mongo
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		register["lastName"] = user.LastName
	}

	register["BirthDay"] = user.BirthDay

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.WebSite) > 0 {
		register["webSite"] = user.WebSite
	}

	updateString := bson.M{
		"$set": register,
	}

	// Se hace el conver de String a ObjectId
	objID, _ := primitive.ObjectIDFromHex(UserId)

	//Se crea el filtro por el Id
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, error := collection.UpdateOne(context, filter, updateString)
	if error != nil {
		return false, error
	}

	return true, nil
}
