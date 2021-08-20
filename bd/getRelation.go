package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/JosueJVL/GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(model models.Relation) (bool, error) {
	_context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("relation")

	//condicion
	filter := bson.M{
		"userId":         model.UserId,
		"userRelationId": model.UserRelationId,
	}

	var result models.Relation

	fmt.Println(result)

	err := col.FindOne(_context, filter).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
