package pgs2

import (
	"fmt"
)


func Maps_Go() {
	fmt.Println("Maps in Go")


	websites := map[string]string{
		"Google": "https://www.google.com",
		"AWS": "https://aws.amazon.com",
	}

	fmt.Println(websites)
	fmt.Println("Google website: ", websites["Google"])
	websites["LinkedIn"] = "https://www.linkedin.com"

	delete(websites, "AWS")
	fmt.Println("After deleting AWS: ", websites)
}