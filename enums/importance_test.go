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
		t.Run(c.input, func(t *testing.T) {
			result, err := DetermineNoteImportance(c.input)
			if err != nil {
				t.Errorf("Error parsing line %s, %s", c.input, err)
			}

			if result != c.result {
				t.Errorf("Error: %q is %q, but should be %q\n",
					c.input, result, c.result)
			}
		})
	}
}
