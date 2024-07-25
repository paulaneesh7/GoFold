package main

import (
	exfour "exercies/ex_four"
	exseven "exercies/ex_seven"
	exsix "exercies/ex_six"
	exthree "exercies/ex_three"
	extwo "exercies/ex_two"
	"fmt"
)

func main(){
	fmt.Println("Hello World")
	extwo.SecondExercise("Aneesh", 25)
	exthree.ThirdExercise(5)
	exfour.FourthExercise([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	exsix.ExerciseSix("aneesh")

	result := exseven.ExerciseSeven(5)
	fmt.Println(result)
}
