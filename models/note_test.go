package models

import (
	"importantnotes/helpers/test"
	"importantnotes/importance"
	"testing"
)

func TestModels_Note_CreateFromLine(t *testing.T) {
	a := &test.Asserter{T: t}

	cases := [...]struct {
		input      string
		importance importance.Importance
	}{
		{"This is a regular line", importance.Regular},
		{"! This is important line", importance.Important},
		{"!!! This is very important line", importance.VeryImportant},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			inputLine := &InputLine{Number: 1, Text: c.input}
			note := NewNote(inputLine)
			a.Equal(1, note.LineNumber)
			a.Equal(c.input, note.Text)
			a.Equal(c.importance, note.Importance)
		})
	}
}
