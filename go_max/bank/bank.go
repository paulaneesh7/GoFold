package bank

import "fmt"

func Bank() {

	var accoountBalance float64 = 1000.0

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter you choice: ")
	fmt.Scan(&choice)

	if choice == 1{
		fmt.Println("Your account balance is: ", accoountBalance)
	}else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter the amount you want to deposit: ")
		fmt.Scan(&depositAmount)
		accoountBalance += depositAmount
		fmt.Println("Balance updated 🫡! New account balance: ", accoountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter the amount you want to withdraw: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accoountBalance {
			fmt.Println("Insufficient balance 😢")
		}else{
			accoountBalance -= withdrawAmount
			fmt.Println("Balance updated 🫡! New account balance: ", accoountBalance)
		}
	} else if choice == 4 {
		fmt.Println("Thank you for using Go Bank")
	}else{
		fmt.Println("Invalid choice")
	}
	

}