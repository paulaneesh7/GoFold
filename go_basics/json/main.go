package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string  `json: "name"`
	Age int	`json: "age"`
	IsAdult bool `json: "is_adult"`
}


func main() {
	fmt.Println("We are learning JSON in Go")
	person := Person{Name: "John", Age: 30, IsAdult: true}
	fmt.Println("Person: ", person)

	// Convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error while marshalling: ", err)
		return
	}
	fmt.Println("Person data is: ", string(jsonData))


	// Convert JSON Encoding into person object (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err)
		return
	}
	fmt.Println("Person data after unmarshalling: ", personData)
}