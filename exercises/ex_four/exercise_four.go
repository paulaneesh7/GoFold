package exfour

import "fmt"

func FourthExercise(nums []int) {
	largest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	smallest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < smallest {
			smallest = nums[i]
		}
	}

	fmt.Println("Largest number is: ", largest, "Smallest number is: ", smallest)
}