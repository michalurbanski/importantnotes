package models

import "importantnotes/enums"

// Note represents a note - one line in input file
type Note struct {
	Line       int
	Text       string
	Importance enums.Importance
}

// NewNote creates a new note based on input line
func NewNote(line int, text string, importance enums.Importance) *Note {
	return &Note{line, text, importance}
}
