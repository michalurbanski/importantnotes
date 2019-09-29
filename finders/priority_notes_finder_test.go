package finders

import (
	"importantnotes/enums"
	"importantnotes/models"
	"testing"
)

func TestFindPriorityNotes(t *testing.T) {
	notes := [...]string{
		"!important notes",
		"!!! very important notes",
		"regular note"}

	var actionList = models.ActionList{}

	// Convert to action list
	for i, text := range notes {
		importance, _ := enums.DetermineNoteImportance(text)
		actionList.Notes = append(actionList.Notes,
			*models.NewNote(i, text, importance))
	}

	// Find priority notes
	priorityNotes := FindPriorityNotes(&actionList)
	if len(priorityNotes.Notes) != 2 {
		t.Errorf("Priority notes count should be %d, but %d determined",
			2, len(priorityNotes.Notes))
	}
}
