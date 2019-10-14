package models

// ActionList represents notes for further processing.
// TODO: based on priority_notes_finder_test this might be converted to []*Note
type ActionList struct {
	Notes []Note
}

// NewActionList creates ActionList with notes based on input lines.
func NewActionList(inputLines []InputLine) *ActionList {
	actionList := ActionList{}

	for _, line := range inputLines {
		actionList.Notes =
			append(actionList.Notes, *NewNote(&line))
	}

	return &actionList
}

// Len returns number of notes.
func (a *ActionList) Len() int {
	return len(a.Notes)
}
