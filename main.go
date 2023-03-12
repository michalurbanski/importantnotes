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
	"os"
)

var inputFilePath string

func init() {
	missingFileParameter := "Path to the file with notes to process.\n"
	missingFileParameter += "Optionally use run.zsh to run the app."

	flag.StringVar(&inputFilePath, "file", "", missingFileParameter)
	flag.Parse()

	if len(inputFilePath) == 0 {
		fmt.Println("ERROR: -file required parameter was not provided.")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting program...")

	configurationReader := configuration.MakeReader(inputFilePath)
	config, err := configurationReader.GetConfig()
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	parser := parsers.SelectInputLinesParser(config)

	fileReader := filereader.NewFileReader(config.InputFilePath, parser)
	lines, err := fileReader.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	// Find very important and important notes
	actionList := models.NewActionList(lines)
	priorityNotes := finders.FindPriorityNotes(actionList)
	processors.SortByPriorityAscending(*priorityNotes)

	// Output to screen found priority notes
	fmt.Println("Following priority tasks were found:")
	printers.MakeColorPrinter().Print(priorityNotes)

	// Stats about the notes
	fmt.Println("Number of read lines is: ", fileReader.TotalReadLines())
	fmt.Println(parser.Stats())

	summary := stats.NewSummary(priorityNotes)
	summary = summary.Calculate()
	fmt.Println(summary)

	// Save stats to a file
	saver := stats.NewSaver(summary, config.Read.OutputPath)
	if err := saver.SaveToFile(); err != nil {
		log.Printf("Error while saving results to output file. %v\n", err)
	}

	fmt.Println("Program finished.")
}
