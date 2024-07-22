package userInput

import (
	"bufio"
	"fmt"
	"os"
)

func UserProvidedInput1() {
	fmt.Println("What's your friend's name?")
	var friendName string

	// Providing th user-input, fmt.Scan reads until the first whitespace character
	fmt.Scan(&friendName)
	fmt.Println("My Friend's name is ", friendName)
	
}

func UserProvidedInput2(){
	/*
		so if you want to
		read a whole line, you might want to use bufio package's NewReader or
		ReadString functions for more complex scenarios
	*/
	fmt.Println("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("My name is ",name)
}