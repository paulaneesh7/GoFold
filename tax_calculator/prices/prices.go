package prices

import (
	"bufio"
	"fmt"
	"os"
	"github.com/paulaneesh7/tax_calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	content, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open the file: ", err)
		return
	}
	
	// to store the scanned content from the file
	var lines []string

	// Scan the content of the file and append it in the slice we created above
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors during the scan and using the err variable we already defined above to handle the error
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed: ", err)
		return
	}

	// Here we will store all the prices from the file
	prices := make([]float64, len(lines))

	prices, err = conversion.StringsToFloats(lines) // Convert the string to float and store it in the prices slice
	if err != nil {
		fmt.Println("Failed to convert the string to float: ", err)
		defer content.Close();
		return
	}


	job.InputPrices = prices // Assign the prices to the struct field

	defer content.Close();
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
