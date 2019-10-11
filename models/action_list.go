package models

// ActionList represents notes for further processing.
// TODO: based on priority_notes_finder_test this might be converted to []*Note
type ActionList struct {
	Notes []Note
}

// NewActionList creates ActionList with notes based on plain lines.
func NewActionList(lines []string) {
	// TODO: Create action list based on []string, because this is how it's used
	// in main method
	panic("not implemented")
}
