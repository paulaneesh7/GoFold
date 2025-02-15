package bank

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Bank() {

	var accoountBalance, err = readBalanceFromFile()
	if err != nil {
		fmt.Println("Failed to read balance from file", err)
		fmt.Println("------------------")
	}
	
	for {
		
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

			if depositAmount <= 0 {
				fmt.Println("Invalid amount 😢, Amount must be greater than 0")
				return
			}

			accoountBalance += depositAmount
			fmt.Println("Balance updated 🫡! New account balance: ", accoountBalance)
			writeBalanceToFile(accoountBalance)
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
				writeBalanceToFile(accoountBalance)
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

func writeBalanceToFile(balance float64){

	// converting balance to string
	balanceText := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceText), 0644) // 0644 is the permission to read and write the file
}

func readBalanceFromFile() (float64, error) {
	content, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(content)

	// now we wikll convert the balanceText to float64
	accountBalance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value.")
	}

	return accountBalance, nil
}

