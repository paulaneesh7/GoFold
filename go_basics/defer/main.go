package main

import "fmt"

func main() {
	fmt.Println("Starting of the program")
	defer fmt.Println("Middle of the program") // this part will run just before the main function ends
	fmt.Println("End of the program")


	data1 := add(10, 20)
	data2 := add(10, 2)
	// when 2 defer keywords are used, the last one will be executed first (LIFO)
	defer fmt.Println("Sum of 10 and 20 is: ", data1)
	defer fmt.Println("Sum of 10 and 20 is: ", data2)
}

func add(a int, b int) int {
	return a + b
}