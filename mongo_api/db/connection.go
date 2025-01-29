package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aneesh16117:aGW3OndfzLb0kFmT@cluster0.jtbbc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "Netflix"
const colName = "Netflix" // here colName is the collection name

// MOST IMPORTANT
var collection *mongo.Collection


// Connect with MongoDB
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// This collection instance is what we will work with throughout the project
	collection = client.Database(dbName).Collection(colName)


	// Check whether collection instance is ready
	fmt.Println("Collection instance is Ready✅")
}