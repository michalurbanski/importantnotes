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
		{"Line not starting from ! is not an important line", Regular},
	}

	for _, c := range cases {
		result, err := DetermineNoteImportance(c.input)
		if err != nil {
			t.Errorf("Error parsing line %s, %s", c.input, err)
		}

		if result != c.result {
			t.Errorf("Error: %s is '%s', but should be '%s'\n",
				c.input, result, c.result)
		}
	}
}
