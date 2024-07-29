package investment

import (
	"fmt"
	"math"
)

func InvestmentCalculator(){
	fmt.Println("Investment Calculator Project")
	fmt.Println("------------------------------------------")

	const inflationRate = 2.5
	var investmentAmount int
	var years int
	expectedReturnRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// everything should be in float here as expectedReturnRate is float
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) 
	futureRealValue := futureValue / math.Pow(1 + inflationRate/ 100, float64(years))
	
	fmt.Println("Future value of investment is: ", futureValue)
	fmt.Println("Future real value of investment is: ", futureRealValue)


}