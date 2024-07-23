package array

import "fmt"

func Array(){
	fmt.Println("Arrays un Go")
	var arr[5] int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3 
	fmt.Println("First Format of Writing Array: ", arr)

	arr1 := [3] int {1,2,3}
	fmt.Println("Second Format of Writing Array:", arr1)

	fmt.Println("Length of array is: ", len(arr))
	fmt.Println("Length of array1 is: ", len(arr1))

	// user-defined array
	userDefinedArray()

}

func userDefinedArray(){
	var nums[5] int
	fmt.Println("Enter the elements into the array: ")
	for i:=0; i<5; i++{
		fmt.Scan(&nums[i])
	}
	fmt.Println("The elements in the array are: ", nums)
}

func twoDimensionalArray(){
	var arr[2][2] int
	fmt.Println("Enter the elements into the 2D array: ")
	for i:=0; i<2; i++{
		for j:=0; j<2; j++{
			fmt.Scan(&arr[i][j])
		}
	}
	fmt.Println("The elements in the 2D array are: ", arr)
}
