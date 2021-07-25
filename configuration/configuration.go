package configuration

// Configuration provides configuration values based.
type Configuration struct {
	// TODO: change variable name
	FileReader Read
}

func makeConfiguration(inputFilePath string) Configuration {
	return Configuration{FileReader: makeRead(inputFilePath)}
}
