package models

import (
	"fmt"
	"strings"
)

// ActionList represents notes for further processing.
// Each note in this collection has already determined importance.
type ActionList struct {
	Notes []Note
}

// NewActionList creates ActionList with notes based on input lines.
func NewActionList(inputLines []InputLine) *ActionList {
	actionList := ActionList{}

	for _, line := range inputLines {
		actionList.Notes = append(actionList.Notes, *NewNote(&line))
	}

	return &actionList
}

// Len returns number of notes.
func (a *ActionList) Len() int {
	return len(a.Notes)
}

func (a *ActionList) String() string {
	var builder strings.Builder

	for _, note := range a.Notes {
		builder.WriteString(fmt.Sprintf("%s\n", note.String()))
	}

	return strings.TrimRight(builder.String(), "\n")
}
