package main

import (
	"fmt"
	"importantnotes/finders"
	"importantnotes/models"
	"importantnotes/parsers"
	"importantnotes/printers"
	"importantnotes/processors"
	"importantnotes/readers/filereader"
	"log"
)

var configFileName = "config.yaml"

func main() {
	fmt.Println("Starting program...")

	// TODO: use line parsers based on config values
	//config := configuration.GetConfig(configFileName)

	path := "./data/input.txt"
	lines, err := filereader.ReadLines(path, parsers.StandardLineParser{})
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
