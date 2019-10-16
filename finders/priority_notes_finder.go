package finders

import (
	"importantnotes/enums"
	"importantnotes/models"
)

// FindPriorityNotes finds those notes that are very important or important.
// They should be acted on as a priority.
func FindPriorityNotes(actionList *models.ActionList) *models.ActionList {
	var priorityNotes []models.Note

	for _, note := range actionList.Notes {
		// TODO: notes types can be passed as parameter
		if note.Importance == enums.Important ||
			note.Importance == enums.VeryImportant {
			priorityNotes = append(priorityNotes, note)
		}
	}

	return &models.ActionList{Notes: priorityNotes}
}
