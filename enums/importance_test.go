package enums

import (
	"testing"
)

func TestDetermineImportanceBasedOnLine(t *testing.T) {
	cases := []struct {
		input  string
		result Importance
	}{
		{"First line", Regular},
		{"! Second line", Important},
		{"!!! Third line", VeryImportant},
	}

	for _, c := range cases {
		result := DetermineNoteImportance(c.input)

		if result != c.result {
			t.Errorf("Error: %s is '%s', but should be '%s'\n", c.input, result, c.result)
		}
	}
}
