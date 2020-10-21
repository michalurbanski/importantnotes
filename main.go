package main

import (
	"errors"
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
	"path"
)

var inputFilePath string

func init() {
	flag.StringVar(&inputFilePath, "file", "", "Path to file with notes")
	flag.Parse()
}

func main() {
	fmt.Println("Starting program...")

	configFilePath, err := getConfigFilePath()
	if err != nil {
		log.Fatal(err) // calls os.Exit(1) automatically
	}

	config, err := configuration.GetConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: this method can be moved to configuration.Configuration
	inputFilePath, err := getInputFileName(config, configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	parser := parsers.SelectInputLinesParser(config)

	fileReader := filereader.NewFileReader(inputFilePath, parser)
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

// getInputFileName gets file name from config or command line argument
func getInputFileName(config configuration.Configuration, configFileName string) (string, error) {
	// If value is provide as cmd line argument than it overwrites config value.
	if len(inputFilePath) > 0 {
		return inputFilePath, nil
	}

	configValue := config.FileReader.File_Name
	if len(configValue) > 0 {
		return configValue, nil
	}

	message := fmt.Sprintf("Input file path has to be provided in %s or using 'file' argument.\n", configFileName)
	message += "Consider also running the application using 'run.zsh' script."

	return "", errors.New(message)
}

// getConfigFilePath reads configuration values from config.{env}.yaml file.
// {env} can be set using environment variable.
// If not set, then by default 'development' value is used.
func getConfigFilePath() (string, error) {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}

	configFileName := fmt.Sprintf("config.%s.yaml", env)

	// Search for config file in the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(currentDir, configFileName), nil
}
