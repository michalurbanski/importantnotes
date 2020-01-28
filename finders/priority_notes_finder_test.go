package finders

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestFinders_FindPriorityNotes(t *testing.T) {
	asserter := test.Asserter{T: t}
	inputLines := []models.InputLine{
		models.InputLine{
			Number: 1, Text: "!important notes",
		},
		models.InputLine{
			Number: 2, Text: "!!! very important notes",
		},
		models.InputLine{
			Number: 3, Text: "regular note",
		},
	}

	actionList := models.NewActionList(inputLines)

	// Find priority notes
	priorityNotes := FindPriorityNotes(actionList)
	asserter.Equal(priorityNotes.Len(), 2)
}
