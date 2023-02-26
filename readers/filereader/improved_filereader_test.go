package filereader

import (
	"bytes"
	"fmt"
	"importantnotes/helpers/test"
	"importantnotes/parsers"
	"strings"
	"testing"
)

func TestFilereader_LinesReaderReadLines_AllLinesReadLinesEqualToInputLines(t *testing.T) {
	a := &test.Asserter{T: t}

	var b strings.Builder
	fmt.Fprintln(&b, "First line")
	fmt.Fprintln(&b, "Second line")
	slice := []byte(b.String())

	reader := bytes.NewReader(slice)
	parser := parsers.StandardLineParser{}
	linesReader := NewLinesReader(parser, reader)
	_, err := linesReader.ReadLines()
	if err != nil {
		t.Errorf("unexpected error")
	}

	a.Equal(2, linesReader.TotalReadLines())
}

func TestFilereader_LinesReaderReadLines_SomeLinesReadOnlyBetweenTags(t *testing.T) {
	a := &test.Asserter{T: t}

	var b strings.Builder

	fmt.Fprintln(&b, "@")
	fmt.Fprintln(&b, "First line")
	fmt.Fprintln(&b, "Second line")
	fmt.Fprintln(&b, ">")
	fmt.Fprintln(&b, "Third line that shouldn't be read")
	slice := []byte(b.String())

	reader := bytes.NewReader(slice)
	parser := parsers.NewStartEndTagParser(parsers.Tag{Name: "@"}, parsers.Tag{Name: ">"})
	linesReader := NewLinesReader(parser, reader)
	_, err := linesReader.ReadLines()
	if err != nil {
		t.Errorf("unexpected error")
	}

	// Should stop reading after finding end tag
	a.Equal(4, linesReader.TotalReadLines())
}
