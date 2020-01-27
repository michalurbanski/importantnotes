package parsers

import (
	"importantnotes/configuration"
	"testing"
)

func Test_InputLineParser_when_no_tags_in_configuration_creates_StandardLineParser(t *testing.T) {
	config := configuration.Configuration{}

	lineParser := SelectInputLinesParser(config)
	switch lp := lineParser.(type) {
	case StandardLineParser:
		break
	default:
		t.Errorf("Incorrect type passed %T", lp)
	}
}

func Test_InputLineParser_when_tags_are_in_configuration_creates_StandardEndTagParser(t *testing.T) {
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
