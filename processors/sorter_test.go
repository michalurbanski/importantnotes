package processors

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

func TestProcessors_SortByPriority_LeastImportantFirst(t *testing.T) {
	asserter := test.Asserter{T: t}

	inputLines := []models.InputLine{
		models.InputLine{
			Number: 1,
			Text:   "!!! Very important note",
		},
		models.InputLine{
			Number: 2,
			Text:   "! Important note",
		},
	}

	actionList := models.NewActionList(inputLines)
	SortByPriorityAscending(*actionList)

	asserter.Equal(actionList.Notes[0].LineNumber, 2)
}

// Proves that Stable function was used
func TestProcessors_SortByPriority_LeastImportantFirstInReadOrder(t *testing.T) {
	asserter := test.Asserter{T: t}

	inputLines := []models.InputLine{
		models.InputLine{
			Number: 1,
			Text:   "!!! Very important note",
		},
		models.InputLine{
			Number: 2,
			Text:   "!!! Another very important note in sequence",
		},
		models.InputLine{
			Number: 3,
			Text:   "! Important note",
		},
	}

	actionList := models.NewActionList(inputLines)
	SortByPriorityAscending(*actionList)

	asserter.Equal(actionList.Notes[1].LineNumber, 1)
	asserter.Equal(actionList.Notes[2].LineNumber, 2)
}
