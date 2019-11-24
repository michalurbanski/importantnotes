package parsers

type Tag struct {
	Name    string
	IsFound bool
}

func (t *Tag) IsEmpty() bool {
	return len(t.Name) == 0
}
