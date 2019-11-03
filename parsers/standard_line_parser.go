package parsers

import "importantnotes/models"

type StandardLineParser struct{}

func (s StandardLineParser) ParseLine(lineNumber int, text string, results []models.InputLine) []models.InputLine {
	line := models.InputLine{Number: lineNumber, Text: text}
	results = append(results, line)
	return results
}
