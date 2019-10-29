package configuration

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	FileReader struct {
		Start_Tag string
		End_Tag   string
	}
}

func GetConfig(fileName string) Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
