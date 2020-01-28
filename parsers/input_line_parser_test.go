package parsers

import (
	"importantnotes/configuration"
	"testing"
)

func TestParsers_InputLineParser_WhenNoTagsInConfigurationCreatesStandardLineParser(t *testing.T) {
	config := configuration.Configuration{}

	lineParser := SelectInputLinesParser(config)
	switch lp := lineParser.(type) {
	case StandardLineParser:
		break
	default:
		t.Errorf("Incorrect type passed %T", lp)
	}
}

func TestParser_InputLineParser_WhenTagsInConfigurationCreatesStartEndTagParser(t *testing.T) {
	config := configuration.Configuration{}
	config.FileReader = struct {
		Start_Tag string
		End_Tag   string
	}{"start", "end"}

	lineParser := SelectInputLinesParser(config)
	switch lp := lineParser.(type) {
	case *StartEndTagParser:
		break
	default:
		t.Errorf("Incorrect type passed %T", lp)
	}
}
