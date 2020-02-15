package configuration

import (
	"github.com/tkanos/gonfig"
)

// Configuration provides configuration values based.
type Configuration struct {
	FileReader struct {
		Start_Tag   string
		End_Tag     string
		File_Name   string
		Output_Path string
	}
}

// GetConfig reads configuration from yaml file.
func GetConfig(fileName string) Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
