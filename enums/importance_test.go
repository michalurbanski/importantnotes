package enums

import (
	"testing"
)

func TestDetermineImportanceBasedOnLine(t *testing.T) {
	line := "First line"
	importance := DetermineNoteImportance(line)

	if importance != Regular {
		t.Error("Importance is not correct")
	}
}
