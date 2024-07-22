package main

import (
	"fmt"
	"go_basics/myutil"
	"go_basics/old"
	"go_basics/userInput"
	"go_basics/variables"
)

func main() {
	fmt.Println("Hello World")
	myutil.PrintMessage("Aneesh")
	variables.LearnVariables()
	fmt.Println("---------------------------------------")

	old.OldCode()
	// userInput.UserProvidedInput1()
	userInput.UserProvidedInput2()
}