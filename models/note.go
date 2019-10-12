package models

import "importantnotes/enums"

// Note represents a note - one line in an input file
// Each line has its own importance
// TODO: avoid uninitialized structures by making type private?
type Note struct {
	LineNumber int
	Text       string
	Importance enums.Importance
}

func createNote(lineNumber int, text string, importance enums.Importance) *Note {
	return &Note{lineNumber, text, importance}
}

// NewNote creates Note object based on input line and its number.
func NewNote(lineNumber int, line string) *Note {
	importance, err := enums.DetermineNoteImportance(line)
	if err != nil {
		panic("Incorrect line parsing")
	}

	return &Note{LineNumber: lineNumber, Text: line, Importance: importance}
}
