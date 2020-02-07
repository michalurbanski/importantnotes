package stats

import (
	"fmt"
	"importantnotes/importance"
	"importantnotes/models"
	"strings"
	"time"
)

// Summary holds stats from program execution.
type Summary struct {
	ImportantCount     int
	VeryImportantCount int
	TotalCount         int
	timeStamp          time.Time
	actions            *models.ActionList
}

// NewSummary creates a new object with summary.
func NewSummary(actions *models.ActionList) *Summary {
	return &Summary{
		timeStamp: time.Now(),
		actions:   actions,
	}
}

// Calculate counts stats from execution.
func (summary *Summary) Calculate() *Summary {
	summary.VeryImportantCount = len(summary.actions.Filter(func(note models.Note) bool {
		return note.Importance == importance.VeryImportant
	}))
	summary.ImportantCount = len(summary.actions.Filter(func(note models.Note) bool {
		return note.Importance == importance.Important
	}))

	// Calculated, instead of just summing important and very important
	// to make sure that there are no errors in logic.
	summary.TotalCount = summary.actions.Len()

	return summary
}

func (summary Summary) String() string {
	var builder strings.Builder

	const timeLayout = "2006-01-02T15:04:05"

	builder.WriteString(fmt.Sprintf("\n%s\n", "Summary"))
	builder.WriteString(fmt.Sprintf("Time: %s\n", summary.timeStamp.Format(timeLayout)))
	builder.WriteString(fmt.Sprintf("Very important tasks: %d\n", summary.VeryImportantCount))
	builder.WriteString(fmt.Sprintf("Important tasks: %d\n", summary.ImportantCount))
	builder.WriteString(fmt.Sprintf("Total tasks: %d\n", summary.TotalCount))

	return builder.String()
}
