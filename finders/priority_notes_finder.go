package finders

import (
	"importantnotes/importance"
	"importantnotes/models"
)

// FindPriorityNotes finds those notes that are very important or important.
// They should be acted on as a priority.
func FindPriorityNotes(actionList *models.ActionList) *models.ActionList {
	seek := []importance.Importance{importance.Important, importance.VeryImportant}

	return findNotesWithImportance(actionList, seek)
}

func findNotesWithImportance(actionList *models.ActionList, importance []importance.Importance) *models.ActionList {
	var priorityNotes []models.Note

	for _, note := range actionList.Notes {
		if contains(importance, note.Importance) {
			priorityNotes = append(priorityNotes, note)
		}
	}

	return &models.ActionList{Notes: priorityNotes}
}

func contains(statuses []importance.Importance, current importance.Importance) bool {
	for _, a := range statuses {
		if a == current {
			return true
		}
	}

	return false
}
