package main

import (
	"fmt"
	"importantnotes/configuration"
	"importantnotes/finders"
	"importantnotes/models"
	"importantnotes/parsers"
	"importantnotes/printers"
	"importantnotes/processors"
	"importantnotes/readers/filereader"
	"importantnotes/stats"
	"log"
)

var configFileName = "config.yaml"
var path = "./realdata/realdata.dat"
var outputPath = "./realdata/output.txt"

func main() {
	fmt.Println("Starting program...")

	config := configuration.GetConfig(configFileName)
	parser := parsers.SelectInputLinesParser(config)

	fileReader := filereader.NewFileReader(path, parser)
	lines, err := fileReader.ReadLines()
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	// Find very important and important notes
	actionList := models.NewActionList(lines)
	priorityNotes := finders.FindPriorityNotes(actionList)
	processors.SortByPriority(*priorityNotes)

	fmt.Println("Following priority tasks were found:")

	printers.ColorPrinter{}.Print(*priorityNotes)

	fmt.Println("Number of read lines is: ", fileReader.TotalReadLines())
	fmt.Println(parser.Stats())

	// Actions stats
	summary := stats.NewSummary(priorityNotes)
	summary = summary.Calculate()
	fmt.Println(summary)

	saver := stats.NewSaver(summary, outputPath)
	saver.SaveToFile()

	fmt.Println("Program finished.")
}
