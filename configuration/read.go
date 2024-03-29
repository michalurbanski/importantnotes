package configuration

// Read contains properties that are expected to be found in a configuration file.
// Configuration file follows the naming convention 'config.{env}.yaml'.
//
// Fields names in the struct must match those list in the config file,
// but they are case-insensitive.
type Read struct {
	// StartTag, put at the beginning of a line, marks the place in a file starting from which lines are read.
	StartTag string
	// EndTag, put at the beginning of a line, marks the place in a file until which lines are read.
	EndTag string
	// TODO: This is write property, not read
	// OutputPath is a path to a file with the results.
	OutputPath string
}
