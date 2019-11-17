package parsers

import (
	"importantnotes/configuration"
	"importantnotes/models"
)

// InputLineParser is used to parse input line one by line.
// This parsers can be switched to parse line in a different way, using this interface.
type InputLineParser interface {
	ParseLine(lineNumber int, text string, results []models.InputLine) ([]models.InputLine, error)
}

// SelectInputLinesParser based on configuration determines which parser
// should be used.
// NOTE: actually this is overengineering for only one checker (or two, counting the default one).
func SelectInputLinesParser(config configuration.Configuration) InputLineParser {
	inputLinesParsers := map[configuration.ConfigurationChecker]InputLineParser{
		configuration.StartEndChecker{}: StartEndLineParser{
			StartTag: config.FileReader.Start_Tag,
			EndTag:   config.FileReader.End_Tag,
		},
	}

	for checker, parser := range inputLinesParsers {
		if checker.Check(config) {
			return parser
		}
	}

	return StandardLineParser{}
}
