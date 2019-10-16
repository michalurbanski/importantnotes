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

// SortByPriority orders action list from highest priority to smallest.
func SortByPriority(actionList models.ActionList) {
	// Sorting can also be done using this function
	// sort.SliceStable(actionList.Notes, func(i, j int) bool {
	// 	return actionList.Notes[i].Importance < actionList.Notes[j].Importance
	// })
	// return actionList

	sort.Stable(ByMostImportantNote(actionList.Notes))

	// how to reverse sorting (if ever needed)
	// sort.Sort(sort.Reverse(SortByMostImportant(actionList.Notes)))
}
