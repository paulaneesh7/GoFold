package prices

import (
	"fmt"

	"github.com/paulaneesh7/tax_calculator/conversion"
	"github.com/paulaneesh7/tax_calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt") // Read the prices from the file
	if err != nil {
		fmt.Println("Failed to read the prices from the file: ", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines) // Convert the string to float and store it in the prices slice
	if err != nil {
		fmt.Println("Failed to convert the string to float: ", err)
		return
	}

	job.InputPrices = prices // Assign the prices to the struct field

}

func (job TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		// the above line to store the price and the tax included price in the map as string
		// also why we changed the return type of the function to map[string]string instead of float64
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
