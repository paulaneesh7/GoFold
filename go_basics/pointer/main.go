package main

import "fmt"

func main() {
	var num int
	num = 2

	var ptr *int // ptr int because this is a pointer to an integer variable
	ptr = &num   // storing the address of num in ptr

	fmt.Println("Value of num: ", num)
	fmt.Println("Address of num: ", ptr)


	// Short-form to do the above steps
	str := "Hello"
	ptr2 := &str
	fmt.Println("Value of str: ", str)
	fmt.Println("Address of str: ", ptr2)
	fmt.Println("Data contain through Pointer: ", *ptr2)

	value := 10
	modifyValueByReference(&value)

}

func modifyValueByReference(num *int) {
	*num = 20
	fmt.Println("Modified value: ", *num)
}