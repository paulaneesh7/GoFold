package profit

import (
	"fmt"
	"math"
)


func ProfitCalculator(){
	fmt.Println("Profit Calculator Project")

	var revenue float64
	var expense float64
	var taxRate float64


	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)
	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	fmt.Println("Earning Before Tax (EBT): ", ebt)

	eat := ebt * math.Abs(1-taxRate/100)
	fmt.Println("Earning After Tax (EAT) i.e Profit: ", eat)

	ratio := fmt.Sprintf("%.2f", ebt/eat) // did this to fix the final result's decimal places upto 2 fixed digits
	fmt.Println("Ratio: ", ratio)
}