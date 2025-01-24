package main

import (
	"fmt"
	"io"
	"net/http"
)


const url = "https://paulaneesh7.in"

func main() {
	fmt.Println("Http Request Handling In Golang")

	response , err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type: %T\n", response)

	defer response.Body.Close() // caller's responsobility to close the connection



	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// converting the data into string format before printing it
	content := string(data)
	fmt.Println(content)
}