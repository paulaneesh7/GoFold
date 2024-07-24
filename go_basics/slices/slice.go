package slices

import "fmt"

func SliceBase() {
	numbers := []int{1, 2, 3}

	fmt.Println("Slice now: ", numbers, "it's length is: ", len(numbers))

	// appending elements to the slice
	numbers = append(numbers, 4,5,6)
	fmt.Println("After appending elements: ", numbers, "it's length is: ", len(numbers))
	fmt.Println("After appending elements: ", numbers, "it's capacity is: ", cap(numbers)) // cap and len are same

	name := []string{"Aneesh", "Swastik", "Suvankar"}
	fmt.Println("Name slice: ", name)

	// creating slice with make function
	nums := make([]int, 3, 5)
	for i:=0; i<3; i++{
		nums[i] = i
	}

	fmt.Println("Slice with make function: ", nums, "it's length is: ", len(nums), "it's capacity is: ", cap(nums))


}