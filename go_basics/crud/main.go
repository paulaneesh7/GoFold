package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest() {
	fmt.Println("Performing GET request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("Error: ", res.Status)
		return
	}


	// Read the response body and print it - TYPE 1
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response: ", err)
		return
	}
	fmt.Println("Response: ", string(data))

}


func performPostRequest(){
	fmt.Println("Performing POST request")

	todo := Todo{
		UserId: 23,
		Title: "Learn Go",
		Completed: false,
	}

	// Convert todo into JSON Encoding (Marshalling) as we want to send JSON data in POST request
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while marshalling: ", err)
		return
	}

	// Convert Json data to string
	jsonString := string(jsonData)

	// Covert string data to reader
	jsonReader := strings.NewReader(jsonString)

	// Send POST request
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer res.Body.Close()


	// Converting the response to a readable form
	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error while reading response: ", err)
		return
	}
	fmt.Println("Response: ", string(data))
}

func performPutRequest(){
	fmt.Println("Performing PUT request")

	todo := Todo{
		UserId: 54,
		Title: "Learn Golang's CRUD",
		Completed: false,
	}

	// Convert todo into JSON Encoding (Marshalling) as we want to send JSON data in PUT request
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error while marshalling data: ", err)
		return
	}

	// Convert Json data to string
	jsonString := string(jsonData)


	// Covert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"


	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Errow while creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send PUT request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending request: ", err)
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response: ", err)
		return
	}

	fmt.Println("Response: ", string(data))

}


func performDeleteRequest(){
	fmt.Println("Performing DELETE request")

	const myURL = "https://jsonplaceholder.typicode.com/todos/2"

	// create DELETE request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil) // as we are just deleting so content is nil
	if err != nil {
		fmt.Println("Error while creating request: ", err)
		return
	}

	// Send DELETE request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error while sending request: ", err)
		return
	}

	defer res.Body.Close()


	// Converting the response to a readable form
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response: ", err)
		return
	}

	// as data will originally won't be deleted from the db but we will get a response status that it is deleted
	fmt.Println("Response: ", string(data))
	fmt.Println("Response Status: ", res.Status)

}

func main(){
	fmt.Println("We are learning CRUD in Go")

	performGetRequest()
	performPostRequest()
	performPutRequest()
	performDeleteRequest()
}