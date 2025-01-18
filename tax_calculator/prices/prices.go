package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) LoadData() {
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

	// Print the content of the file from the slice where we appended it
	for _, line := range lines {
		fmt.Println(line)
	}

	defer content.Close();
}



func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)	
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
