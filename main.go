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
var configFileName = "config.yaml"
var outputPath = "./realdata/output.txt"

func init() {
	flag.StringVar(&inputFilePath, "file", "", "Path to file with notes")
	flag.Parse()
}

func main() {
	fmt.Println("Starting program...")

	config := configuration.GetConfig(configFileName)
	inputFilePath, err := GetInputFileName(config)
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	parser := parsers.SelectInputLinesParser(config)

	fileReader := filereader.NewFileReader(inputFilePath, parser)
	lines, err := fileReader.ReadLines()
	if err != nil {
		log.Fatal(err)
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

// GetInputFileName gets file name from config or command line argument
func GetInputFileName(config configuration.Configuration) (string, error) {
	// If value is provide as cmd line argument than it overwrites config value.
	if len(inputFilePath) > 0 {
		return inputFilePath, nil
	}

	configValue := config.FileReader.File_Name
	if len(configValue) > 0 {
		return configValue, nil
	}

	return "", fmt.Errorf("Input file path has to be provided in %s or using '-file' argument", configFileName)
}
