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

func TestParsers_StartEndTagParserParseLine_LineWithStartTagIsNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	line, err := parser.ParseLine(1, startTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNil(line)
}

func TestParsers_StartEndTagParserParseLine_WhenStartTagIsEmptyLineIsRead(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(Tag{Name: ""}, endTag)

	line, err := parser.ParseLine(1, startTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNotNil(line)
}

func TestParsers_StartEndTagParserParseLine_LineWithEndTagIsNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	line, err := parser.ParseLine(1, endTagName)
	if err != nil {
		t.Error(err)
	}

	asserter.IsNil(line)
}

func TestParsers_StartEndTagParserParseLine_LinesBeforeStartTagNotIncludedInResults(t *testing.T) {
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

func TestParsers_StartEndTagParserParseLine_LinesAfterEndTagNotRead(t *testing.T) {
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

func TestParsers_StartEndTagParserParseLine_LinesBetweenTagsAreRead(t *testing.T) {
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

func TestParsers_StartEndTagParserParseLine_StartTagEndTagLinesBeforeStartTagAreNotRead(t *testing.T) {
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

func TestParsers_StartEndTagParserParseLine_NoTags(t *testing.T) {
	parser := NewStartEndTagParser(Tag{}, Tag{})
	if _, err := parser.ParseLine(1, "sample text"); err == nil {
		t.Error("Expected error but none found")
	}
}

func TestParsers_StartEndTagParserParseLine_TagsDefinedButNoneInContent(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := NewStartEndTagParser(startTag, endTag)

	results := []*models.InputLine{}
	lines := [...]string{
		"First line",
		"Second line",
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

	asserter.Equal(0, len(results))
}
