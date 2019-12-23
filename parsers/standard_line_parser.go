package parsers

import "importantnotes/models"

// StandardLineParser has no specific way of input parsing.
// It always adds each line to the results.
type StandardLineParser struct{}

// ParseLine adds InputLine object to the results, as it is.
func (s StandardLineParser) ParseLine(lineNumber int, text string, results []models.InputLine) ([]models.InputLine, error) {
	line := models.InputLine{Number: lineNumber, Text: text}
	results = append(results, line)
	return results, nil
}
