package printers

import (
	"importantnotes/enums"
	"importantnotes/models"

	"github.com/gookit/color"
)

// ColorPrinter prints ActionList in colors, based on note priority.
type ColorPrinter struct{}

var colors = map[enums.Importance]color.Color{
	enums.VeryImportant: color.Red,
	enums.Important:     color.Yellow,
}

// Print prints note in color based on priority.
//
// Red: very important note
// Yellow: important note
func (ColorPrinter) Print(actionList models.ActionList) {
	for _, note := range actionList.Notes {
		c := colors[note.Importance]
		c.Println(note)
	}
}
