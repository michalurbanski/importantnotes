package parsers

// Tag represents tag read from config file (for now it's start or end tag)
type Tag struct {
	// Name is a content of a tag.
	Name string
}

// IsEmpty checks if tag is empty.
func (t Tag) IsEmpty() bool {
	return len(t.Name) == 0
}
