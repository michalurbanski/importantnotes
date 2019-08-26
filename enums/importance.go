package enums

// TODO: http://golang-basic.blogspot.com/2014/07/step-by-step-guide-to-declaring-enums.html

// Importance states how note is important
type Importance int

// Status denotes whether note is not started, done, or aborted
type Status int

// TODO: how to wrap it
const (
	// Regular means it can be done at any time
	Regular Importance = iota

	// Important means it should be done relatively fast
	Important

	// VeryImportant means it's a priority
	VeryImportant
)

func DetermineNoteImportance(line string) Importance {
	// TODO: based on line prefix
	return Regular
}

const (
	NotStarted Status = iota
	Done
	Aborted
)
