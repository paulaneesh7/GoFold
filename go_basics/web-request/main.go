package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web requests in Go")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()
	

	// Read the response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response: ", err)
		return
	}
	fmt.Println("Response: ", string(data)) // without string() it will print the byte array
}