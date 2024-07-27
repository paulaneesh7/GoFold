package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, mango, banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts) // will output [apple  orange  mango  banana]

	str := "one two three four five six seven eight nine ten"
	count := strings.Count(str, "two")
	fmt.Println(count) // will output 1

	str = "    Hello World!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed: ",trimmed) // will output Hello World!

	str1 := "Aneesh"
	str2 := "Paul"
	result := strings.Join([]string{str1, str2}, " ") // we have to provide an array of strings everytime
	fmt.Println("Resultant string: ",result) // will output Aneesh Paul
}