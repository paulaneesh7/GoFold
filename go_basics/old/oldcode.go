package old

import "fmt"

func OldCode() {
	age := 24
	name := "Alex Ryan"
	height := 1888.5
	fmt.Println("Name: ", name, "Age: ", age, "height: ", height) // println creates a new line for each print

	/*
		The Printf function (print formatted) is used for formatted printing. It allows you to
		control the output format by using format specifiers similar to those in the C
		programming language
	*/
	fmt.Printf("Name is %s, Age is %d, Height is %f", name, age, height)
	fmt.Printf("\nData-type of 'name' is %T\n", name)
}