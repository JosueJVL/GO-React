package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPerfil(ID string) (models.User, error) {

	context, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	dataBase := MongoConection.Database("twitter")
	collection := dataBase.Collection("users")

	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(context, condition).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
