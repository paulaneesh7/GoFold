package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL in Go")

	myURL := "https://example.com:8080/path/to/resource?name=abc#fragment"

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while parsing URL: ", err)
		return
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("Query: ", parsedURL.RawQuery)


	// Modifying the URL
	parsedURL.Scheme = "http"
	parsedURL.Host = "example.org"
	parsedURL.Path = "/new-path"
	parsedURL.RawQuery = "name=xyz"

	// Constructing a URL string from a URL object
	newURL := parsedURL.String()
	fmt.Println("New URL: ", newURL)
	
}