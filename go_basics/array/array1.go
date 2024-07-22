package array

import "fmt"

func UserDefinedArray2() {
	var n int
	fmt.Println("Enter the range for the array:")
	fmt.Scan(&n)

	// use make function to create a dynamic array instead of fixed size
	nums := make([]int, n)
	
	fmt.Println("Enter the elements into the array: ")
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	fmt.Println("The elements in the array are: ", nums)
}