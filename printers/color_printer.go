package printers

import (
	"importantnotes/importance"
	"importantnotes/models"

	"github.com/gookit/color"
)

// ColorPrinter prints ActionList in colors, based on note priority.
type ColorPrinter struct{}

var colors = map[importance.Importance]color.Color{
	importance.VeryImportant: color.Red,
	importance.Important:     color.Yellow,
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
