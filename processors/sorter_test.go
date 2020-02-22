package processors

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestProcessors_SortByPriority_MostImportantFirst(t *testing.T) {
	asserter := test.Asserter{T: t}

	inputLines := []models.InputLine{
		models.InputLine{
			Number: 1,
			Text:   "! Important note",
		},
		models.InputLine{
			Number: 2,
			Text:   "!!! Very important note",
		},
	}

	actionList := models.NewActionList(inputLines)
	SortByPriorityAscending(*actionList)

	asserter.Equal(actionList.Notes[0].LineNumber, 2)
}

// Proves that Stable function was used
func TestProcessors_SortByPriority_MostImportantFirstInReadOrder(t *testing.T) {
	asserter := test.Asserter{T: t}

	inputLines := []models.InputLine{
		models.InputLine{
			Number: 1,
			Text:   "! Important note",
		},
		models.InputLine{
			Number: 2,
			Text:   "!!! Very important note",
		},
		models.InputLine{
			Number: 3,
			Text:   "!!! Another very important note in sequence",
		},
	}

	actionList := models.NewActionList(inputLines)
	SortByPriorityAscending(*actionList)

	asserter.Equal(actionList.Notes[0].LineNumber, 2)
	asserter.Equal(actionList.Notes[1].LineNumber, 3)
}
