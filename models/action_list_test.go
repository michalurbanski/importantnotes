package models

import (
	"importantnotes/enums"
	"importantnotes/helpers/test"
	"testing"
)

func TestCreateValidActionList(t *testing.T) {
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
	asserter.Equal(actionList.Notes[1].Importance, enums.VeryImportant)
}
