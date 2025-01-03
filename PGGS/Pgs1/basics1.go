package pgs1

import (
	"fmt"
	"math"
	"math/rand"
)


func swap(x, y string) (string, string) {
	return y, x
}

func split (sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Basics1() {
	fmt.Println("Welcome to Go Language Basics 1")

	// Math Module
	fmt.Println("Randome Number: ", rand.Int())
	fmt.Println("Random Number between 1 and 10: ", rand.Intn(10)+1)
	fmt.Println("PI Value: ", math.Pi)

	// Swap Function
	a, b := swap("Hello", "World")
	fmt.Println("Swapped: ", a, b)

	// Split Function
	x, y := split(17)
	fmt.Println("Split: ", x, y)

}