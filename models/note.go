package models

import "importantnotes/enums"

// Note represents a note - one line in input file
type Note struct {
	LineNumber int
	Text       string
	Importance enums.Importance
}

// NewNote creates a new note based on input line
func NewNote(lineNumber int, text string, importance enums.Importance) *Note {
	return &Note{lineNumber, text, importance}
}
