package models

import (
	"fmt"
	"importantnotes/enums"
)

// Note represents a note - one line in an input file
// Each line has its own importance
// TODO: avoid uninitialized structures by making type private?
type Note struct {
	LineNumber int
	Text       string
	Importance enums.Importance
}

// NewNote creates a note based on input line.
func NewNote(inputLine *InputLine) *Note {
	importance, err := enums.DetermineNoteImportance(inputLine.Text)
	if err != nil {
		panic("Incorrect line parsing")
	}

	return &Note{
		LineNumber: inputLine.Number,
		Text:       inputLine.Text,
		Importance: importance}
}

func (n Note) String() string {
	return fmt.Sprintf("%d: %s", n.LineNumber, n.Text)
}
