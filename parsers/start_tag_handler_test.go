package parsers

import (
	"importantnotes/helpers/test"
	"testing"
)

func TestParsers_StartTagHandlerWasFired(t *testing.T) {
	asserter := &test.Asserter{T: t}
	startTag := Tag{Name: "@<"}

	startTagHandler := NewStartTagHandler(nil, startTag)
	if _, err := startTagHandler.Handle(1, "@<"); err == nil {
		asserter.Equal(true, startTagHandler.WasFired())
	}
}

func TestParsers_EndTagHandlerWasFired(t *testing.T) {
	asserter := &test.Asserter{T: t}
	endTag := Tag{Name: "@>"}

	endTagHandler := NewEndTagHandler(endTag)
	if _, err := endTagHandler.Handle(1, "@>"); err == nil {
		asserter.Equal(true, endTagHandler.WasFired())
	}
}
