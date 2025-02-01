package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulaneesh7/mongo_api/router"
)


func main() {
	
	router := router.Router()
	fmt.Println("Server is starting...")


	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Listening on port 8080")

}