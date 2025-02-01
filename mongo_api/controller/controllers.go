package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paulaneesh7/mongo_api/helpers"
	"github.com/paulaneesh7/mongo_api/model"
)

// Actual controller functions

// GET
func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	movies := helpers.GetAllMovies()
	json.NewEncoder(w).Encode(movies)

}

// POST
func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	helpers.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// PUT
func UpdateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")

	params := mux.Vars(r)
	id := params["id"]

	helpers.UpdateOneMovie(id)
	json.NewEncoder(w).Encode("Movie marked as watched")
}

// DELETE - ONE
func DeleteMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	params := mux.Vars(r)
	id := params["id"]

	helpers.DeleteOneMovie(id)
	json.NewEncoder(w).Encode("Movie with id " + id + " deleted successfully")
}

// DELETE - MANY
func DeleteAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count := helpers.DeleteManyMovies()

	// converted the count into int from string
	json.NewEncoder(w).Encode("Deleted " + strconv.Itoa(int(count)) + " movies")

}
