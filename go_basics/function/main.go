package main

import "fmt"

func main() {
	fmt.Println("Learning functions in go...")
	simpleFunction()

	val1:=add(3,4)
	fmt.Println("Value after addition", val1)

	val2 := multiply(2,3,4)
	fmt.Println("Value after multiplication", val2)

}

func simpleFunction(){
	fmt.Println("Simple function")
}

// Function with return type mentioned
func add(a int, b int) int{
	return a+b
}

// Function with named returned variable, we are returning a variable(int) which stores the computed operation
func multiply(a, b, c int) (result int){
	return a*b*c
}
