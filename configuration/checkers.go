package configuration

// ConfigurationChecker checks if configuration has specific element.
type ConfigurationChecker interface {
	Check(Configuration) bool
}

// StartEndChecker checks if content is delimited by 'start' and/or 'end' tag (based on config).
type StartEndChecker struct {
}

// Check verifies if StartEndChecker can be used.
func (checker StartEndChecker) Check(configuration Configuration) bool {
	return len(configuration.FileReader.Start_Tag) > 0 ||
		len(configuration.FileReader.End_Tag) > 0
}
