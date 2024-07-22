package main

import (
	"fmt"
	"go_basics/array"
	"go_basics/error_handling"
	"go_basics/myutil"
	"go_basics/old"
	"go_basics/variables"
)

func main() {
	fmt.Println("Hello World")
	myutil.PrintMessage("Aneesh")
	variables.LearnVariables()
	fmt.Println("---------------------------------------")

	old.OldCode()
	// userInput.UserProvidedInput1()
	// userInput.UserProvidedInput2()

	fmt.Println("---------------------------------------")
	error_handling.Error()

	fmt.Println("---------------------------------------")
	array.Array()
	array.UserDefinedArray2()

	
}