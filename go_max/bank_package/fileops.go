package bank_package

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeFloatToFile(value float64, fileName string) {

	// converting value to string
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644) // 0644 is the permission to read and write the file
}

func getFloatFromFile(fileName string) (float64, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	valueText := string(content)

	// now we will convert the balanceText to float64
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}