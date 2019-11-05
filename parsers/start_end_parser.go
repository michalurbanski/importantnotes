package parsers

import "importantnotes/models"

type StartEndLineParser struct {
	StartTag        string
	EndTag          string
	isStartTagFound bool
	isEndTagFound   bool
	// TODO: another method has to be exposed - like in case of scanner - to not proceed
	// if end tag was found
}

func (s StartEndLineParser) ParseLine(lineNumber int, text string, results []models.InputLine) []models.InputLine {
	// TODO: implementation to be changed - add lines only when start tag was encountered;
	// store this fact in local variable
	// test this
	line := models.InputLine{Number: lineNumber, Text: text}
	results = append(results, line)
	return results
}
