package parsers

import "fmt"

// ParserStats contains information about parser execution.
type ParserStats struct {
	IsStartTagHandlerFired bool
	IsEndTagHandlerFired   bool
}

func (p ParserStats) String() string {
	start := fmt.Sprintf("Start tag handler was fired? %t", p.IsStartTagHandlerFired)
	end := fmt.Sprintf("End tag handler was fired? %t", p.IsEndTagHandlerFired)

	return fmt.Sprintf("%s\n%s", start, end)
}
