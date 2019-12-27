package parsers

import (
	"fmt"
	"importantnotes/models"
	"regexp"
)

// LineHandler defines interface for line handlers.
type LineHandler interface {
	Handle(lineNumber int, text string) (*models.InputLine, error)
}

// StartTagHandler decides how to handle line when start tag is present.
type StartTagHandler struct {
	next        LineHandler
	isEnabled   bool
	searchedTag Tag
	matcher     Matcher
}

// NewStartTagHandler creates a new StartTagHandler.
func NewStartTagHandler(lineHandler LineHandler, tag Tag) *StartTagHandler {
	return &StartTagHandler{
		isEnabled:   true,
		searchedTag: tag,
		next:        lineHandler,
	}
}

// Handle for StartTagHandler returns line if start tag was found.
// It does not return lines before start tag, or when line starts with start tag.
func (handler *StartTagHandler) Handle(lineNumber int, text string) (*models.InputLine, error) {
	if handler.isEnabled {
		isMatch, err := handler.matcher.IsMatch(text, handler.searchedTag.Name)
		if err != nil {
			return nil, err
		}

		if isMatch {
			handler.isEnabled = false
		}

		return nil, nil // line should not be read

		// Check if line has start tag (use SearchedTag fieldvalue)
		// if not then continue to the next line
	} else {
		// Start tag checker is disbaled that means start tag was found,
		// and if second handler is OK with this line, then it can be read.
		if handler.next != nil {
			return handler.next.Handle(lineNumber, text)
		} else {
			// If there's no second handler, then all notes after StartTag was found should be read
			return &models.InputLine{Number: lineNumber, Text: text}, nil
		}
	}
}

// EndTagHandler decides how to handle line when end tag is present.
type EndTagHandler struct {
	IsEnabled   bool
	SearchedTag Tag
	Matcher     Matcher
}

// NewEndTagHandler creates a new EndTagHandler.
func NewEndTagHandler(tag Tag) *EndTagHandler {
	return &EndTagHandler{
		IsEnabled:   true,
		SearchedTag: tag,
	}
}

// Handle for EndTagHandler returns line until end tag is found.
// It does not return lines after end tag, or when line starts with end tag.
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
		// Tag was found so it doesn't make sense to search further
		return nil, nil
	}
}

// Matcher contains helper methods used to find tag in line.
type Matcher struct{}

// IsMatch checks if line starts from tag.
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
