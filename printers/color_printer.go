package printers

import (
	"importantnotes/enums"
	"importantnotes/models"

	"github.com/gookit/color"
)

type Printer struct{}

var colors = map[enums.Importance]color.Color{
	enums.VeryImportant: color.Red,
	enums.Important:     color.Yellow,
}

func (Printer) Print(actionList models.ActionList) {
	for _, note := range actionList.Notes {
		c := colors[note.Importance]
		c.Println(note)
	}
}
