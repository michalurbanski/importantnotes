package models

// ActionList represents notes for further processing.
// TODO: based on priority_notes_finder_test this might be converted to []*Note
type ActionList struct {
	Notes []Note
}

// NewActionList creates ActionList with notes based on plain lines.
func NewActionList(lines []string) *ActionList {
	actionList := ActionList{}

	// TODO: index here is not a line number, but just subsequent item - to be fixed
	for index, line := range lines {
		actionList.Notes = append(actionList.Notes, *NewNote(index, line))
	}

	return &actionList
}

// Len returns number of notes.
func (a *ActionList) Len() int {
	return len(a.Notes)
}
