package investment

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func InvestmentCalculator(){
	fmt.Println("Investment Calculator Project")
	fmt.Println("------------------------------------------")

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
	futureValue , futureRealValue := calculateFutureValues(float64(investmentAmount), years, expectedReturnRate)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/ 100, float64(years))
	
	fmt.Printf("Future value of investment is: %v\n", futureValue)
	fmt.Println("Future real value of investment is: ", futureRealValue)


}

func outputText(text string) string{
	return fmt.Sprintf("Output: %v", text)
}

func calculateFutureValues(investmentAmount float64, years int, expectedReturnRate float64) (float64, float64){
	fv := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) 
	rfv := fv / math.Pow(1 + inflationRate/ 100, float64(years))

	return fv, rfv
}