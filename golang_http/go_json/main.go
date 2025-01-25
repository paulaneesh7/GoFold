package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course_name"`
	Price    int64
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Creating JSON In Golang")

	EncodeJSON()
	DecoderJSON()
}

// Converting JSON to a suitable format
func EncodeJSON() {
	kzCourses := []course{
		{"React", 299, "Udemy", "abc", []string{"web-dev", "js"}},
		{"Angular", 199, "Udemy", "abc", []string{"web-dev", "js"}},
		{"Golang", 199, "LCO", "abc", nil},
	}

	// package this data as JSON data
	finalJSON, err := json.Marshal(kzCourses)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJSON))
}


func DecoderJSON() {

	// Any data that comes from the web comes in byte[] format
	jsonDataFromWeb := []byte(
		`{"course_name":"React", "price":299, "platform":"Udemy", "password":"abc", "tags":["web-dev", "js"]}`,
	)

	var kzCourse course

	// Check whether the data is valid JSON (it doesn't throw an error it just gives a bool)
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid",checkValid)
		json.Unmarshal(jsonDataFromWeb, &kzCourse)
		fmt.Println(kzCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases where you might want to add data to key-value pair
	// ( here we took interface for value because we don't know the type of value )
	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	for key, value := range myOnlineData {
		fmt.Printf("Key is %s and value is %v\n", key, value)
	}
}