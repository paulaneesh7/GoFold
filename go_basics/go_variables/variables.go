package go_variables

import "fmt"

func LearnVariables() {
	// if we don't assign data-type to a variable, then golang gets to know about them on its own
	// basically infers the variable on its own
	var name string = "Aneesh"
	var age int = 21;

	var isAdult bool = true
	if age > 18 {
		fmt.Println(name, "is a SDE and he's an Adult", isAdult)
	} else if age < 18 {
		isAdult = false
		fmt.Println(name, "is a SDE and he's an Adult", isAdult)
	}


	// short-hand syntax (Used to declare a variable and assign it's value at the same time)
	friendName := "Suvankar"
	owes := 57.4
	fmt.Println(name,"friend's name is", friendName, "and he still owes him", owes)

}