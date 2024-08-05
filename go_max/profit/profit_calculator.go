package profit

import (
	"errors"
	"fmt"
	"math"
	"os"
)


func ProfitCalculator(){
	fmt.Println("Profit Calculator Project")

	revenue, err1 := getUserInput("Enter revenue: ")
	expense, err2 := getUserInput("Enter expense: ")
	taxRate, err3 := getUserInput("Enter tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}


	ebt, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Earnings After Tax: %.2f\n", profit)
	fmt.Printf("Earnings Before Tax to Earnings After Tax Ratio: %s\n", ratio)
}


func getUserInput(infoText string) (float64, error){

	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if (userInput < 0){
		return 0, errors.New("revenue cannot be negative")
	}


	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, string){
	ebt := revenue - expenses
	eat := ebt * math.Abs(1-taxRate/100)
	ratio := fmt.Sprintf("%.2f", ebt/eat) // did this to fix the final result's decimal places upto 2 fixed digits

	storeResults(ebt, eat, ebt/eat)

	return ebt, eat, ratio
	
}

/*Storing all the results in a file in a formatted way*/
func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %0.1f\nProfit: %0.1f\nRatio: %0.1f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

