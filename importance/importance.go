package importance

import (
	"regexp"
)

// Note: Implementation based on http://golang-basic.blogspot.com/2014/07/step-by-step-guide-to-declaring-enums.html

// Importance states how note is important.
//
// There are three categories:
// !!! Very important
// ! Important
// Regular note
type Importance int

const (
	// VeryImportant means it's a top priority.
	VeryImportant Importance = iota

	// Important means it should be done relatively fast.
	Important

	// Regular means it can be done at any time. Low priority, nice to have.
	Regular
)

func (i Importance) String() string {
	names := [...]string{
		"VeryImportant",
		"Important",
		"Regular",
	}

	return names[i]
}

// DetermineNoteImportance determine how note is important, and returns correct Importance enum value.
func DetermineNoteImportance(line string) (Importance, error) {
	// At first matching Very important lines has to be implemented
	// as they will fit also into second condition of important notes.
	// So order does matter.
	match, err := regexp.MatchString("^!!!", line)
	if err != nil {
		return Regular, err
	}

	if match {
		return VeryImportant, nil
	}

	// Check for important notes
	match, err = regexp.MatchString("^!", line)
	if err != nil {
		return Regular, err
	}

	if match {
		return Important, nil
	}

	return Regular, nil
}
