package bank_package

import (
	"fmt"
)

const accountBalanceFile = "balance.txt"


func Bank() {

	var accoountBalance, err = getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Failed to read balance from file", err)
		fmt.Println("------------------")
	}
	
	for {
		
		presentOptions()
		
		var choice int
		fmt.Print("Enter you choice: ")
		fmt.Scan(&choice)
		

		if choice == 1{
			fmt.Println("Your account balance is: ", accoountBalance)
		}else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount 😢, Amount must be greater than 0")
				return
			}

			accoountBalance += depositAmount
			fmt.Println("Balance updated 🫡! New account balance: ", accoountBalance)
			writeFloatToFile(accoountBalance, accountBalanceFile)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accoountBalance {
				fmt.Println("Insufficient balance 😢")
				return
			}else{
				accoountBalance -= withdrawAmount
				fmt.Println("Balance updated 🫡! New account balance: ", accoountBalance)
				writeFloatToFile(accoountBalance, accountBalanceFile)
			}
		} else if choice == 4 {
			fmt.Println("Thank you for using Go Bank")
			break
		}else{
			fmt.Println("Invalid choice")
			break
		}
	}

	fmt.Println(("Thank you for using Go Bank"))
}



