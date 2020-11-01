package configuration

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/tkanos/gonfig"
)

// Reader allows to read configuration from file
type Reader struct {
	inputFilePath string
}

// MakeReader creates a new Reader
func MakeReader(inputFilePath string) Reader {
	return Reader{inputFilePath: inputFilePath}
}

// GetConfig returns configuration read from file.
func (reader Reader) GetConfig() (Configuration, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Configuration{}, err
	}

	config, err := reader.readConfig(configFilePath)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
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

// readConfig reads configuration from yaml file.
func (reader Reader) readConfig(configFileName string) (Configuration, error) {
	configuration := makeConfiguration(reader.inputFilePath)
	err := gonfig.GetConf(configFileName, &configuration)
	if err != nil {
		return configuration, fmt.Errorf("No config file found. %v", err)
	}

	if err := reader.checkInputFilePresence(configuration, configFileName); err != nil {
		return configuration, err
	}

	return configuration, nil
}

// checkInputFilePresence checks if path to input file is provide either in config
// or as a command line argument.
func (reader Reader) checkInputFilePresence(config Configuration, configFileName string) error {
	// If value is provide as cmd line argument than it overwrites config value.
	if len(reader.inputFilePath) > 0 {
		return nil
	}

	configValue := config.FileReader.FileName()
	if len(configValue) > 0 {
		return nil
	}

	message := fmt.Sprintf("Input file path has to be provided in %s or using 'file' argument.\n", configFileName)
	message += "Consider also running the application using 'run.zsh' script."

	return errors.New(message)
}
