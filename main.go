package main

import (
	"flag"
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

var inputFilePath string

func init() {
	flag.StringVar(&inputFilePath, "file", "", "Path to file with notes")
	flag.Parse()
}

func main() {
	fmt.Println("Starting program...")

	configurationReader := configuration.MakeReader(inputFilePath)
	config, err := configurationReader.GetConfig()
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	parser := parsers.SelectInputLinesParser(config)

	fileReader := filereader.NewFileReader(config.FileReader.FileName(), parser)
	lines, err := fileReader.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	// Find very important and important notes
	actionList := models.NewActionList(lines)
	priorityNotes := finders.FindPriorityNotes(actionList) // TODO: this action can be in ActionList
	processors.SortByPriorityAscending(*priorityNotes)

	fmt.Println("Following priority tasks were found:")
	printers.ColorPrinter{}.Print(*priorityNotes)

	fmt.Println("Number of read lines is: ", fileReader.TotalReadLines())
	fmt.Println(parser.Stats())

	// Actions stats
	summary := stats.NewSummary(priorityNotes)
	summary = summary.Calculate()
	fmt.Println(summary)

	saver := stats.NewSaver(summary, config.FileReader.Output_Path)
	if err := saver.SaveToFile(); err != nil {
		log.Printf("Error while saving results to output file. %v\n", err)
	}

	fmt.Println("Program finished.")
}
