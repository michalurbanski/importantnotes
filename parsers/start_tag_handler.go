package parsers

//TODO: clean comments in this file

import (
	"fmt"
	"importantnotes/models"
	"regexp"
)

type LineHandler interface {
	Handle(lineNumber int, text string) (*models.InputLine, error)
}

type StartTagHandler struct { // TODO: add constructor for both handlers to set IsEnabled to true
	Next        LineHandler
	IsEnabled   bool
	SearchedTag Tag
	Matcher     Matcher
}

func NewStartTagHandler(lineHandler LineHandler, tag Tag) *StartTagHandler {
	return &StartTagHandler{
		IsEnabled:   true,
		SearchedTag: tag,
		Next:        lineHandler,
	}
}

func (handler *StartTagHandler) Handle(lineNumber int, text string) (*models.InputLine, error) {
	if handler.IsEnabled {
		isMatch, err := handler.Matcher.IsMatch(text, handler.SearchedTag.Name)
		if err != nil {
			return nil, err
		}

		if isMatch {
			handler.IsEnabled = false
		}

		return nil, nil // line should not be read

		// Check if line has start tag (use SearchedTag fieldvalue)
		// if not then continue to the next line
	} else {
		// Start tag checker is disbaled that means start tag was found,
		// and if second handler is OK with this line, then it can be read.
		if handler.Next != nil {
			return handler.Next.Handle(lineNumber, text)
		} else {
			// If there's no second handler, then all notes after StartTag was found should be read
			return &models.InputLine{Number: lineNumber, Text: text}, nil
		}
	}
}

type EndTagHandler struct {
	IsEnabled   bool
	SearchedTag Tag
	Matcher     Matcher
}

func NewEndTagHandler(tag Tag) *EndTagHandler {
	return &EndTagHandler{
		IsEnabled:   true,
		SearchedTag: tag,
	}
}

func (handler *EndTagHandler) Handle(lineNumber int, text string) (*models.InputLine, error) {
	if handler.IsEnabled {
		// If endtag was found then disable it
		// If it wasn't found in this line then return line

		isMatch, err := handler.Matcher.IsMatch(text, handler.SearchedTag.Name)
		if err != nil {
			return nil, err
		}

		if isMatch {
			handler.IsEnabled = false
			return nil, nil // When end tag was found then the current line should not be read, as it doesn't have note
		}

		return &models.InputLine{Number: lineNumber, Text: text}, nil
	} else {
		// tag was found so it doesn't make sense to search further
		return nil, nil
	}
}

type Matcher struct{}

func (matcher Matcher) IsMatch(text string, tag string) (isMatch bool, outErr error) {
	if len(tag) == 0 {
		return false, nil
	}

	match, err := regexp.MatchString(fmt.Sprintf("^%s", tag), text)
	if err != nil {
		return false, err
	}

	return match, nil
}
