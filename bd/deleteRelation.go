package bd

import (
	"context"
	"time"

	"github.com/JosueJVL/GO-React/models"
)

func DeleteRelation(model models.Relation) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConection.Database("twitter")
	col := db.Collection("relation")

	_, err := col.DeleteOne(context, model)
	if err != nil {
		return false, err
	}

	return true, err
}
