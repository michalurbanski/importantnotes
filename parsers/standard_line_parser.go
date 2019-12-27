package parsers

import "importantnotes/models"

// StandardLineParser has no specific way of input parsing.
// It always adds each line to the results.
type StandardLineParser struct{}

// ParseLine in case of this parser returns read line.
func (s StandardLineParser) ParseLine(lineNumber int, text string) (*models.InputLine, error) {
	line := &models.InputLine{Number: lineNumber, Text: text}
	return line, nil
}
