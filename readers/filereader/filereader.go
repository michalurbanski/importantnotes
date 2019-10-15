package filereader

import (
	"bufio"
	"importantnotes/models"
	"os"
)

// ReadLines reads all lines from a specified file.
func ReadLines(path string) ([]models.InputLine, error) {
	results := []models.InputLine{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// We want to capture also lines to be able to capture them later in ediotr.
	// Due to this lines' calculation starts from 1.
	lineCounter := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := models.InputLine{Number: lineCounter, Text: scanner.Text()}
		results = append(results, line)
		lineCounter++
	}

	// Check for any errors while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return results, nil
}
