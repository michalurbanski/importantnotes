package filereader

import (
	"bytes"
	"fmt"
	"importantnotes/helpers/test"
	"importantnotes/parsers"
	"strings"
	"testing"
)

func Test_reader_reads_all_lines_number_of_read_lines_is_equal_to_input_lines(t *testing.T) {
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
		t.Errorf("Unexpected error")
	}

	a.Equal(2, linesReader.TotalReadLines())
}

func Test_reader_reads_only_lines_between_tags(t *testing.T) {
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
		t.Errorf("Unexpected error")
	}

	// Should stop reading after finding end tag
	a.Equal(4, linesReader.TotalReadLines())
}
