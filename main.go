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
	priorityNotes := finders.FindPriorityNotes(actionList) // TODO: this action can be in ActionList
	processors.SortByPriorityAscending(*priorityNotes)     // TODO: this could be object chaining or sth similar https://medium.com/@yuseferi/method-chaining-in-golang-fa29a8c40c97

	fmt.Println("Following priority tasks were found:")
	printers.ColorPrinter{}.Print(*priorityNotes)

	fmt.Println("Number of read lines is: ", fileReader.TotalReadLines())
	fmt.Println(parser.Stats())

	// Actions stats
	summary := stats.NewSummary(priorityNotes)
	summary = summary.Calculate()
	fmt.Println(summary)

	saver := stats.NewSaver(summary, config.Read.OutputPath)
	if err := saver.SaveToFile(); err != nil {
		log.Printf("Error while saving results to output file. %v\n", err)
	}

	fmt.Println("Program finished.")
}
