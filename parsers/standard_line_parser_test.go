package parsers

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestAllLinesAreUsed(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := StandardLineParser{}

	input := map[int]string{
		1: "First line",
		2: "Second line",
	}

	results := []models.InputLine{}
	for lineNumber, text := range input {
		results = parser.ParseLine(lineNumber, text, results)
	}

	asserter.Equal(len(results), len(input))
}
