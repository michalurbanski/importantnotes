package main

import (
	"fmt"
	"importantnotes/finders"
	"importantnotes/models"
	"importantnotes/printers"
	"importantnotes/processors"
	"importantnotes/readers/filereader"
	"log"
)

func main() {
	fmt.Println("Starting program...")

	path := "./data/input.txt"
	lines, err := filereader.ReadLines(path)
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	// Find very important and important notes
	actionList := models.NewActionList(lines)
	priorityNotes := finders.FindPriorityNotes(actionList)
	processors.SortByPriority(*priorityNotes)

	fmt.Println("Following priority tasks were found:")
	//fmt.Println(priorityNotes)

	colorPrinter := printers.Printer{}
	colorPrinter.Print(*priorityNotes)

	fmt.Println("Program finished.")
}
