package filereader

import (
	"bufio"
	"importantnotes/models"
	"importantnotes/parsers"
	"os"
)

// ReadLines reads all lines from a specified file.
func ReadLines(path string, inputLineParser parsers.InputLineParser) ([]models.InputLine, error) {
	results := []models.InputLine{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lineCounter := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results, err = inputLineParser.ParseLine(lineCounter, scanner.Text(), results)
		if err != nil {
			return results, err
		}
		lineCounter++
	}

	// Check for any errors while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return results, nil
}
