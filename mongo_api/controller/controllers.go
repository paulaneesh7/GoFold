package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/paulaneesh7/mongo_api/db"
	"github.com/paulaneesh7/mongo_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MONGO helper functions

// INSERT 1 Record

func InsertOneMovie(movie model.Netflix) {
	// Now you can use db.Collection here
	inserted, err := db.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Inserted 1 movie in the collection with Id: ", inserted.InsertedID)

}



// UPDATE 1 Record
func UpdateOneMovie(id string) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	filter := bson.M{"_id": Id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Modified Count: ", res.ModifiedCount)
}