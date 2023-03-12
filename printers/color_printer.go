package printers

import (
	"importantnotes/importance"
	"importantnotes/models"

	"github.com/gookit/color"
)

type Printer interface {
	Print(actionList *models.ActionList)
}

// ColorPrinter prints ActionList in colors, based on note priority.
type colorPrinter struct{}

func MakeColorPrinter() Printer {
	return colorPrinter{}
}

var colors = map[importance.Importance]color.Color{
	importance.VeryImportant: color.Red,
	importance.Important:     color.Yellow,
}

// Print prints notes in color based on a note priority.
//
// Red: very important note
// Yellow: important note
func (colorPrinter) Print(actionList *models.ActionList) {
	for _, note := range actionList.Notes {
		c := colors[note.Importance]
		c.Println(note)
	}
}
