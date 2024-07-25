package exsix

import (
	"fmt"
	"strings"
)

func ExerciseSix(str string) {
	str = strings.ToLower(str) // using ToLower from strings package of Go
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			fmt.Printf("%s is not a palindrome\n", str)
			return
		}
		start++
		end--
	}
	fmt.Printf("%s is Palindromic\n", str)
}