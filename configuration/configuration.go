package configuration

// Configuration provides configuration values based.
type Configuration struct {
	// File path provided as the application argument.
	InputFilePath string
	// Read is configuration read from config file.
	Read Read // This variable name has to be the same as root key in config.yaml file.
}

func makeConfiguration(inputFilePath string) Configuration {
	return Configuration{
		InputFilePath: inputFilePath,
		Read:          Read{},
	}
}
