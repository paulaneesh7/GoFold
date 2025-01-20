package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)


type FileManager struct {
	InputFilePath string
	OutputFilePath string
}




func (fm FileManager) ReadLines() ([]string, error) {
	// Read the content of the file
	content, err := os.Open(fm.InputFilePath)
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



func (fm FileManager) WriteJSON(data interface{}) error {
	content, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create the file")
	}


	// Create a new JSON encoder
	encoder := json.NewEncoder(content)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to encode the data")
	}

	defer content.Close()
	return nil
}



// Constructor function to create a new FileManager object
func NewFileManager(inputFilePath string, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath: inputFilePath,
		OutputFilePath: outputFilePath,
	}
}