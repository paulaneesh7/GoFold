package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Printf("Type of num is %T and num is %d\n", num, num)

	// Converting from int to float
	data := float64(num)
	fmt.Printf("Type data is %T and data is %f\n", data,data)

	// Converting from int to string
	val := 34
	str := strconv.Itoa(val)
	fmt.Printf("Type of str is %T and str is %s\n", str, str)

	// Converting from string to int
	num_str := "23342"
	num_int, _ := strconv.Atoi(num_str)
	fmt.Printf("Type of num is %T and num is %d\n", num_int, num_int)

	// Converting from string to float
	float_str := "23342.23"
	num_float, _ := strconv.ParseFloat(float_str, 64)
	fmt.Printf("Type of num is %T and num is %f\n", num_float, num_float)
}