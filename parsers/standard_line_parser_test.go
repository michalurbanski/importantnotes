package parsers

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestParsers_StandardLineParserParseLine_AllInputLinesAreRead(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := StandardLineParser{}
	input := [...]string{"First line", "Second line"}

	results := []*models.InputLine{}
	for lineNumber, text := range input {
		line, err := parser.ParseLine(lineNumber, text)
		if err != nil {
			t.Error(err)
		}

		results = append(results, line)
	}

	asserter.Equal(len(input), len(results))
}
