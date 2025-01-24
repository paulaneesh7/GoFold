package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=123"

func main() {
	fmt.Println("URL Handling In Golang")

	// parsing the url
	result, _ := url.Parse(myUrl)
	

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())


	// getting the query parameters
	queryParams := result.Query()
	fmt.Println("First Query Params is: ", queryParams["coursename"])
	fmt.Println("Second Query Params is: ", queryParams["paymentid"])


	// iterating over the query params
	for key, value := range queryParams {
		fmt.Println("Query is: ", key, " Value is: ", value)
	}
}
