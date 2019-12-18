package configuration

// Checker checks if configuration has specific element.
type Checker interface {
	Check(Configuration) bool
}

// StartEndChecker checks if content is delimited by 'start' and/or 'end' tag (based on config).
type StartEndChecker struct {
}

// Check verifies if StartEndChecker can be used.
// It can be used when at least one of the tags (start_tag, end_tag) is defined.
func (checker StartEndChecker) Check(configuration Configuration) bool {
	return len(configuration.FileReader.Start_Tag) > 0 ||
		len(configuration.FileReader.End_Tag) > 0
}
