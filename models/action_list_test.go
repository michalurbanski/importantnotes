package models

import (
	"importantnotes/enums"
	"importantnotes/helpers/test"
	"testing"
)

func TestCreateValidActionList(t *testing.T) {
	asserter := test.Asserter{T: t}
	lines := []string{
		"First line",
		"!!! Second line",
	}
	actionList := NewActionList(lines)

	asserter.Equal(actionList.Len(), len(lines))
	asserter.Equal(actionList.Notes[1].Importance, enums.VeryImportant)
}
