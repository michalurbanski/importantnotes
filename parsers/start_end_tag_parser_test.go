package parsers

import (
	"fmt"
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

// Here, for test purposes, it doesn't matter what value is set.
// In application configuration file these values will be more unique,
// comparing to the real content that might appear.
const startTagName = "aaa"
const endTagName = "bbb"

var startTag = Tag{Name: startTagName}
var endTag = Tag{Name: endTagName}

func Test_StartEndTagParser_line_with_start_tag_is_not_included_in_results(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	line, err := parser.ParseLine(1, startTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNil(line)
}

func Test_StartEndTagParser_when_start_tag_is_empty_then_line_is_read(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(Tag{Name: ""}, endTag)

	line, err := parser.ParseLine(1, startTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNotNil(line)
}

func Test_StartEndTagParser_line_with_end_tag_is_not_included_in_results(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	line, err := parser.ParseLine(1, endTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNil(line)
}

func Test_StartEndTagParser_lines_before_start_tag_are_not_included_in_results(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	results := []*models.InputLine{}
	lines := [...]string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		startTagName,
		"This line should be read",
	}

	for index, line := range lines {
		result, err := parser.ParseLine(index, line)
		if err != nil {
			t.Error(err)
		}

		if result != nil {
			results = append(results, result)
		}
	}

	asserter.Equal(1, len(results))
}

func Test_StartEndTagParser_lines_after_end_tag_should_not_be_read(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(Tag{Name: ""}, endTag) // start tag is empty so lines before end tag should be read

	results := []*models.InputLine{}
	lines := [...]string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		endTagName,
		"This line should be read",
	}

	for index, line := range lines {
		result, err := parser.ParseLine(index, line)
		if err != nil {
			t.Error(err)
		}

		if result != nil {
			results = append(results, result)
		}
	}

	asserter.Equal(2, len(results))
}

func Test_StartEndTagParser_only_lines_between_tags_should_be_read(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	lineThatShouldBeFound := "This line should be found"

	results := []*models.InputLine{}
	lines := [...]string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		startTagName,
		lineThatShouldBeFound, // Only this line should be in the results
		endTagName,
		"Another insignificant line",
	}

	for index, line := range lines {
		result, err := parser.ParseLine(index, line)
		if err != nil {
			t.Error(err)
		}

		if result != nil {
			results = append(results, result)
		}
	}

	asserter.Equal(1, len(results))
	asserter.Equal(3, results[0].Number)
	asserter.Equal(lineThatShouldBeFound, results[0].Text)
}

func Test_StartEndTagParser_ParseLine_starttag_noendtag_lines_before_starttag_should_not_be_read(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, Tag{Name: ""})

	lineThatShouldBeFound := "This line should be found"

	results := []*models.InputLine{}
	lines := [...]string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		startTagName,
		lineThatShouldBeFound, // all lines starting from here should be read
		endTagName,
		"Another insignificant line",
	}

	for index, line := range lines {
		result, err := parser.ParseLine(index, line)
		if err != nil {
			t.Error(err)
		}

		if result != nil {
			results = append(results, result)
		}
	}

	fmt.Printf("Results: %v", results)
	asserter.Equal(3, len(results))
	asserter.Equal(3, results[0].Number)
	asserter.Equal(lineThatShouldBeFound, results[0].Text)
}
