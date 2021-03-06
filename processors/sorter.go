package processors

import (
	"importantnotes/models"
	"sort"
)

// Note: Sorting custom type https://yourbasic.org/golang/how-to-sort-in-go/

// ByMostImportantNote sorts from the most important notes to the least ones.
type ByMostImportantNote []models.Note

func (a ByMostImportantNote) Len() int {
	return len(a)
}

func (a ByMostImportantNote) Less(i, j int) bool {
	return a[i].Importance < a[j].Importance
}

func (a ByMostImportantNote) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// SortByPriorityAscending orders action list from highest priority to smallest.
func SortByPriorityAscending(actionList models.ActionList) {
	// Sorting can also be done using this function
	// sort.SliceStable(actionList.Notes, func(i, j int) bool {
	// 	return actionList.Notes[i].Importance < actionList.Notes[j].Importance
	// })
	// return actionList

	// how to sort using normal sort order (if even needed)
	//sort.Stable(ByMostImportantNote(actionList.Notes))

	// sort descending
	sort.Stable(sort.Reverse(ByMostImportantNote(actionList.Notes)))
}
