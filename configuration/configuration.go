package configuration

// Configuration provides configuration values based.
type Configuration struct {
	Read Read // This variable name has to be the same as root key in config.yaml file.
}

func makeConfiguration(inputFilePath string) Configuration {
	return Configuration{Read: makeRead(inputFilePath)}
}
