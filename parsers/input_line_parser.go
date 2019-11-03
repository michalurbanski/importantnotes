package parsers

import "importantnotes/models"

// InputLineParser is used to parse input line one by line.
// This parsers can be switched to parse line in a different way, using this interface.
type InputLineParser interface {
	ParseLine(lineNumber int, text string, results []models.InputLine) []models.InputLine
}
