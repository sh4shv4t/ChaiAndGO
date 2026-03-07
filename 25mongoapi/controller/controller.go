package controller

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://23ucs745_db_user:KR7pt55lpWgVuYI5@cluster0.sqd8dj6.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {

	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo conn success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Reached collection instance")
}
