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

	var err error
	results := []models.InputLine{}
	for lineNumber, text := range input {
		results, err = parser.ParseLine(lineNumber, text, results)
		if err != nil {
			t.Error(err)
		}
	}

	asserter.Equal(len(results), len(input))
}
