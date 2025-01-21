package main

import (
	"fmt"

	"github.com/paulaneesh7/tax_calculator/filemanager"
	"github.com/paulaneesh7/tax_calculator/prices"
)



func main() {
	fmt.Println("Price Tax Calculator")

	taxRate := []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJobs := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJobs.Process()
		priceJobs.LoadData()
	}

}
