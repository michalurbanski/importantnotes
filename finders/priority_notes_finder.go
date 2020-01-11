package finders

import (
	"importantnotes/enums"
	"importantnotes/models"
)

// FindPriorityNotes finds those notes that are very important or important.
// They should be acted on as a priority.
func FindPriorityNotes(actionList *models.ActionList) *models.ActionList {
	seek := []enums.Importance{enums.Important, enums.VeryImportant}

	return findNotesWithImportance(actionList, seek)
}

func findNotesWithImportance(actionList *models.ActionList, importance []enums.Importance) *models.ActionList {
	var priorityNotes []models.Note

	for _, note := range actionList.Notes {
		if contains(importance, note.Importance) {
			priorityNotes = append(priorityNotes, note)
		}
	}

	return &models.ActionList{Notes: priorityNotes}
}

func contains(statuses []enums.Importance, current enums.Importance) bool {
	for _, a := range statuses {
		if a == current {
			return true
		}
	}

	return false
}
