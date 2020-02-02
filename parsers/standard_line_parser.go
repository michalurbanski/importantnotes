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

// IsEnabled checks whether parsers is active.
// As for StandardLineParser it's always enabled.
func (s StandardLineParser) IsEnabled() bool {
	return true
}

// Stats for this parser always returns information that handlers were not fired as they're not defined.
func (s StandardLineParser) Stats() ParserStats {
	return ParserStats{IsEndTagHandlerFired: false, IsStartTagHandlerFired: false}
}
