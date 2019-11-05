package configuration

type ConfigurationChecker interface {
	Check(Configuration) bool
}

type StartEndChecker struct {
}

// Check verifies if both tags are filled.
func (checker StartEndChecker) Check(configuration Configuration) bool {
	return len(configuration.FileReader.Start_Tag) > 0 &&
		len(configuration.FileReader.End_Tag) > 0
}
