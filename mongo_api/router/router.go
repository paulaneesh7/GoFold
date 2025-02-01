package router

import (
	"github.com/gorilla/mux"
	controllers "github.com/paulaneesh7/mongo_api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/movie", controllers.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/movies", controllers.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllers.UpdateMovieController).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovieController).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMoviesController).Methods("DELETE")

	return router
}