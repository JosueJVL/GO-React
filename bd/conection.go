package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://adminMongo:UWtmxo8vgIVxDEdJ@cluster.sl8k1.mongodb.net/twitter")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la Base de Datos")

	return client
}

func CheckConnection() int {
	err := MongoConection.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	return 1
}
