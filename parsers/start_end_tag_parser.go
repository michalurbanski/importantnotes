package parsers

import (
	"importantnotes/models"
)

// New version of start_end_parser.go - TODO: remove the old one after rewrite is completed.

type StartEndTagParser struct {
	handler LineHandler
}

// Can be created only when one of the tags is present.
func NewStartEndTagParser(startTag Tag, endTag Tag) *StartEndTagParser {
	parser := new(StartEndTagParser)

	// Chain handlers in a correct order.
	var endTagHandler LineHandler
	var firstTagHandler LineHandler

	if endTag.IsEmpty() == false {
		endTagHandler = NewEndTagHandler(endTag)
	}

	if startTag.IsEmpty() == false {
		firstTagHandler = NewStartTagHandler(endTagHandler, startTag)
	} else {
		firstTagHandler = endTagHandler
	}

	parser.handler = firstTagHandler
	return parser
}

func (parser *StartEndTagParser) ParseLine(lineNumber int, text string) (*models.InputLine, error) {
	return parser.handler.Handle(lineNumber, text)
}
