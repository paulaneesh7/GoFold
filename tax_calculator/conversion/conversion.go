package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	var floats []float64

	// Print the content of the file from the slice where we appended it
	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}