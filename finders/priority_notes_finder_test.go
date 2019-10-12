package finders

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestFindPriorityNotes(t *testing.T) {
	asserter := test.Asserter{T: t}

	notes := []string{
		"!important notes",
		"!!! very important notes",
		"regular note"}

	actionList := models.NewActionList(notes)

	// Find priority notes
	priorityNotes := FindPriorityNotes(actionList)
	asserter.Equal(priorityNotes.Len(), 2)
}
