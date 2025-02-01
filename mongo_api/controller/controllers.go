package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/paulaneesh7/mongo_api/helpers"
)

// Actual controller functions


func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-wwww-form-urlencoded")
	movies := helpers.GetAllMovies()
	json.NewEncoder(w).Encode(movies)

}