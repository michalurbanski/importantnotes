package parsers

import (
	"fmt"
	"importantnotes/models"
	"regexp"
)

type StartEndLineParser struct {
	StartTag        string
	EndTag          string
	isStartTagFound bool
	isEndTagFound   bool
	// TODO: another method has to be exposed - like in case of scanner - to not proceed
	// if end tag was found
}

// ParseLine parses input lines according to the following logic:
// - Skip all lines before (and including) StartTag
// - Take lines
// - Skip EndTag line and all subsequent lines
func (s StartEndLineParser) ParseLine(lineNumber int, text string, results []models.InputLine) []models.InputLine {
	// TODO: implementation to be changed - add lines only when start tag was encountered;
	// store this fact in local variable
	// test this

	match, err := regexp.MatchString(fmt.Sprintf("^%s", s.StartTag), text)
	if err != nil {
		return results // Maybe it could be better handled?
	}

	if match {
		s.isStartTagFound = true // TODO: won't work in this way because it's passed by value
		return results
	}

	line := models.InputLine{Number: lineNumber, Text: text}
	results = append(results, line)
	return results
}
