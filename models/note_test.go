package models

import (
	"importantnotes/enums"
	"importantnotes/helpers/test"
	"testing"
)

// TODO: make these tests parameterized
func TestCreatesNewNoteFromLine(t *testing.T) {
	a := &test.Asserter{T: t}

	lineNumber := 1
	line := "This is a regular line"
	inputLine := &InputLine{Number: lineNumber, Text: line}
	note := NewNote(inputLine)

	a.Equal(note.LineNumber, lineNumber)
	a.Equal(note.Text, line)
	a.Equal(note.Importance, enums.Regular)
}

func TestCreatesNewNoteForImportantLine(t *testing.T) {
	a := &test.Asserter{T: t}

	lineNumber := 2
	line := "!!! This is very important line"
	inputLine := &InputLine{Number: lineNumber, Text: line}
	note := NewNote(inputLine)

	a.Equal(note.LineNumber, lineNumber)
	a.Equal(note.Text, line)
	a.Equal(note.Importance, enums.VeryImportant)
}

//TODO: skip beginning of the line indicating importance when parsing line
