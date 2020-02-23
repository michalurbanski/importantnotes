package models

import (
	"fmt"
	"importantnotes/importance"
)

// Note is a line in an input file.
// Each line has its own importance.
// Line number is important to easily find this note later in the input file.
type Note struct {
	LineNumber int
	Text       string
	Importance importance.Importance
}

// NewNote creates a note based on input line.
func NewNote(inputLine *InputLine) *Note {
	importance, err := importance.DetermineNoteImportance(inputLine.Text)
	if err != nil {
		panic("Incorrect regular expression: " + err.Error())
	}

	return &Note{
		LineNumber: inputLine.Number,
		Text:       inputLine.Text,
		Importance: importance}
}

func (n Note) String() string {
	return fmt.Sprintf("%d: %s", n.LineNumber, n.Text)
}
