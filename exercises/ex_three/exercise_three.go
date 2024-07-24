package exthree

import "fmt"

func ThirdExercise(num int){
	for i := 1; i<= 10; i++{
		fmt.Println(num, "x", i, "=", num*i)
	}
}