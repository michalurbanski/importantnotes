package models

import (
	"importantnotes/helpers/test"
	"importantnotes/importance"
	"testing"
)

func TestModels_ActionList_CreateValidActionList(t *testing.T) {
	asserter := test.Asserter{T: t}
	inputLines := []InputLine{
		InputLine{
			1, "First line",
		},
		InputLine{
			2, "!!! Second line",
		},
	}

	actionList := NewActionList(inputLines)

	asserter.Equal(actionList.Len(), len(inputLines))
	asserter.Equal(actionList.Notes[1].Importance, importance.VeryImportant)
}
