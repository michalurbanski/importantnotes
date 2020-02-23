package configuration

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

// Configuration provides configuration values based.
type Configuration struct {
	FileReader FileReader
}

// FileReader contains properties that are expected to be found in configuration file.
type FileReader struct {
	// Start_Tag put at the line beginning marks the place in file from which lines are read.
	Start_Tag string
	// End_Tag put at the line beginning marks the place in file until which lines are read.
	End_Tag string
	// File_Name is a path to file from which data is read.
	File_Name string
	// Output_path is a path to file with results.
	Output_Path string
}

// GetConfig reads configuration from yaml file.
func GetConfig(fileName string) (Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		return configuration, fmt.Errorf("No config file found. %v", err)
	}

	return configuration, nil
}
