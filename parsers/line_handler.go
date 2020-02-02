package parsers

import "importantnotes/models"

// LineHandler defines interface for line handlers.
type LineHandler interface {
	Handle(lineNumber int, text string) (*models.InputLine, error)
	IsEnabled() bool
	WasFired() bool
	Next() LineHandler
}
