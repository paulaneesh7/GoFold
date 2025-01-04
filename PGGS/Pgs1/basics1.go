package pgs1

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// A switch case function to tell on which OS Go is running
func switch_case() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// Switch Case function to tell which day it is Today
func switch_day() {
	fmt.Print("Today is: ")
	day := time.Now().Weekday()
	switch day {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown Day")
	}
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

	// Switch Case
	switch_case()
	switch_day()

}
