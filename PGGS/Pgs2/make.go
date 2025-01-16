package pgs2

import (
	"fmt"
)

func Make_Array() {
	fmt.Println("Make function in Go for creating Arrays and Slices")

	userName := make([]string, 3, 5)

	userName[0] = "John"
	userName[1] = "Doe"
	userName[2] = "Smith"

	userName = append(userName, "Jane")
	userName = append(userName, "Doe")
	userName = append(userName, "Aneesh")
	userName = append(userName, "Paul")

	fmt.Println("Username slice: ", userName)

	// Loops usage with Arrays
	loops_array_map(userName)
}


type intMap map[string]int

// func (m intMap) output () {
// 	fmt.Println(m)
// }

func Make_Maps() {
	fmt.Println("Make function for Maps")

	ages := make(intMap)

	ages["John"] = 30
	ages["Doe"] = 25
	ages["Smith"] = 40

	fmt.Println("Ages map:", ages)

	johnAge := ages["John"]
	fmt.Println("John's age:", johnAge)

	if age, exists := ages["Doe"]; exists {
		fmt.Println("Doe's age:", age)
	} else {
		fmt.Println("Doe not found in the map.")
	}

	delete(ages, "Smith")
	fmt.Println("Ages map after deleting Smith:", ages)

	fmt.Println("Iterating over the map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}


}


func loops_array_map(nums []string) {
	for i, val := range nums {
		fmt.Println("Index: ", i, "Value: ", val)
	}
}