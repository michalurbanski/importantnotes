package configuration

// Configuration provides configuration values based.
type Configuration struct {
	FileReader FileReader
}

func makeConfiguration(inputFilePath string) Configuration {
	return Configuration{FileReader: makeFileReader(inputFilePath)}
}
