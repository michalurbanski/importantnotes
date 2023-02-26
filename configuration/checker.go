package configuration

// Checker checks if configuration has a specific element(s).
type Checker interface {
	Check(Configuration) bool
}

// StartEndChecker checks if content is delimited by 'start' and/or 'end' tag (based on config.{environment}.yaml file).
type StartEndChecker struct {
}

// Check verifies if StartEndChecker can be used.
// It can be used when at least one of the tags (startTag, endTag) is defined in the config.{environment}.yaml file.
func (checker StartEndChecker) Check(configuration Configuration) bool {
	return len(configuration.Read.StartTag) > 0 ||
		len(configuration.Read.EndTag) > 0
}
