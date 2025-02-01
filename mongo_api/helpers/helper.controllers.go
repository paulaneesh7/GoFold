package helpers

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



// GET All Records
func GetAllMovies() [] model.Netflix {

	// we created a seperated variable for context here in this case because we will be using it frequently
	ctx := context.Background()
	cursor, err := db.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}


	var movies[] model.Netflix


	// Iterate over the cursor and print each record
	for cursor.Next(ctx) {
		var movie model.Netflix
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err.Error())
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(ctx)

	fmt.Println("All movies: ", movies)
	return movies
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


// DELETE 1 Record
func DeleteOneMovie(id string) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	filter := bson.M{"_id": Id}
	res, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Deleted Count: ", res.DeletedCount)
}


// DELETE Many Records
func DeleteManyMovies(movie model.Netflix ) int64 {
	
	// M -> Unordered , D -> Ordered and by passing bson.M{{}} we mean that we want to delete everything present in the collection
	filter := bson.D{{}}
	res, err := db.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Number of movies deleted : ", res.DeletedCount)
	return res.DeletedCount
}