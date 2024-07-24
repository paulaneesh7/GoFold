package main

import "fmt"

func main() {
	x := 10
	if x == 0 {
		fmt.Println("x is zero")
	} else{
		fmt.Println("x is not zero")
	}

	z  := 5
	if z == 0{
		fmt.Println("z is zero")
	} else if z >= 5 {
		fmt.Println("z is greater than or equal to 5")
	} else {
		fmt.Println("z is less than 5")
	}

	y := 10
	if x >= 10 && y <= 10 {
		fmt.Println("x is greater than or equal to 10 and y is less than or equal to 10")
	} else {
		fmt.Println("x is less than 10 and y is greater than 10")
	}

	SwitchStatement()
}

func SwitchStatement(){
	a := 10
	switch a {
		case 5:
			fmt.Println("a is 5")
		case 10:
			fmt.Println("a is 10")
		case 15:
			fmt.Println("a is 15")
		case 20:
			fmt.Println("a is 20")
		default:
			fmt.Println("a is not 5, 10, 15, 20")
	}

	month := "November"

	switch month {
		case "January":
			fmt.Println("It's January")
		case "February":
			fmt.Println("It's February")
		case "March":
			fmt.Println("It's March")
		case "April":
			fmt.Println("It's April")
		case "May":
			fmt.Println("It's May")
		case "June":
			fmt.Println("It's June")
		case "July":
			fmt.Println("It's July")
		case "August":
			fmt.Println("It's August")
		case "September":
			fmt.Println("It's September")
		case "October":
			fmt.Println("It's October")
		case "November":
			fmt.Println("It's November")
		case "December":
			fmt.Println("It's December")
		default:
			fmt.Println("Invalid month")
	}
}