package main

import (
	"github.com/paulaneesh7/go_advance/waitgroups"
)

func main() {
	// go routines.Greeter("Hello")
	// routines.Greeter("World")

	sites := []string{
		"https://google.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://netflix.com",
		"https://paulaneesh7.in",
	}

	for _, site := range sites {
		go waitgroups.GetStatusCode(site)

		// wait group to keep track of the number of go routines, in this case we have 1 in the above line
		waitgroups.Wg.Add(1)
	}

	// wait group to wait for all the go routines to finish
	waitgroups.Wg.Wait()
}
