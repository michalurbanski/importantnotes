package parsers

import (
	"errors"
	"importantnotes/models"
)

// StartEndTagParser is used to read lines between start and end tag.
type StartEndTagParser struct {
	handler LineHandler
}

// NewStartEndTagParser creates a new StartEndTagParser based on tags.
// Can created only when one of the tags is present.
func NewStartEndTagParser(startTag Tag, endTag Tag) *StartEndTagParser {
	parser := new(StartEndTagParser)

	// Chain handlers in the correct order.
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

// ParseLine processes lines between start and end tag.
func (parser StartEndTagParser) ParseLine(lineNumber int, text string) (*models.InputLine, error) {
	if parser.handler == nil {
		return nil, errors.New("No handlers defined for this parser. Do you intend to use a different parser?")
	}

	return parser.handler.Handle(lineNumber, text)
}

// IsEnabled checks if any of the defined handlers are still enabled.
func (parser StartEndTagParser) IsEnabled() bool {
	return parser.handler.IsEnabled()
}

// Stats returns parser stats (whether line handlers were fired)
func (parser StartEndTagParser) Stats() ParserStats {
	return ParserStats{
		IsStartTagHandlerFired: parser.handler.WasFired(),
		IsEndTagHandlerFired:   parser.handler.Next().WasFired(),
	}
}
