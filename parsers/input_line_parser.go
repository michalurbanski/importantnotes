package parsers

import (
	"importantnotes/configuration"
	"importantnotes/models"
)

// InputLineParser is used to parse input read line by line.
// Parsers can be switched to parse line in a different way, using this interface.
type InputLineParser interface {
	ParseLine(lineNumber int, text string) (*models.InputLine, error)
	IsEnabled() bool
	Stats() ParserStats
}

// SelectInputLinesParser determines which parser should be used based on configuration.
func SelectInputLinesParser(config configuration.Configuration) InputLineParser {
	inputLinesParsers := map[configuration.Checker]InputLineParser{
		configuration.StartEndChecker{}: NewStartEndTagParser(
			Tag{Name: config.Read.StartTag},
			Tag{Name: config.Read.EndTag},
		),
	}

	// NOTE: actually this is overengineering for only one checker that is currently defined.
	for checker, parser := range inputLinesParsers {
		if checker.Check(config) {
			return parser
		}
	}

	return StandardLineParser{}
}
