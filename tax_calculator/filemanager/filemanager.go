package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Read the content of the file
	content, err := os.Open(path)
	if err != nil {
		return nil, errors.New("could not open the file")
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
		return nil, errors.New("reading the file content failed")
	}

	

	defer content.Close()
	return lines, nil
}