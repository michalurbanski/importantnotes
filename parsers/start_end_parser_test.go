package parsers

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

// Here, for test purposes, it doesn't matter what value is set.
// In application configuration file these values will be more unique,
// comparing to the real content that might appear.
var startTag = "aaa"
var endTag = "bbb"

func TestLineWithStartTagIsNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{Name: startTag}, EndTag: &Tag{}}

	results := []models.InputLine{}
	results, err := parser.Parse(1, startTag, results)
	if err != nil {
		t.Error(err)
	}

	asserter.Equal(0, len(results))
}

func TestEmptyStartTagLineIsRead(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{Name: ""}, EndTag: &Tag{}} // No end tag specified, all lines should be read

	results := []models.InputLine{}
	results, err := parser.Parse(1, startTag, results)
	if err != nil {
		t.Error(err)
	}

	asserter.Equal(1, len(results))
}

func TestLineWithEndTagIsNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{}, EndTag: &Tag{Name: endTag}} // TODO: make constructor to ensure that properties are filled

	results := []models.InputLine{}
	results, err := parser.Parse(1, endTag, results)
	if err != nil {
		t.Error(err)
	}

	asserter.Equal(0, len(results))
}

func TestLinesBeforeStartTagAreNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{Name: startTag}, EndTag: &Tag{}}

	results := []models.InputLine{}
	lines := []string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		startTag,
		"This line should be read",
	}

	var err error
	for index, line := range lines {
		results, err = parser.Parse(index, line, results)
		if err != nil {
			t.Error(err)
		}
	}

	asserter.Equal(1, len(results))
}

func TestLinesAfterEndTagShouldNotBeRead(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{}, EndTag: &Tag{Name: endTag}}

	results := []models.InputLine{}
	lines := []string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		endTag,
		"This line should be read",
	}

	var err error
	for index, line := range lines {
		results, err = parser.Parse(index, line, results)
		if err != nil {
			t.Error(err)
		}
	}

	asserter.Equal(2, len(results))
}

func TestOnlyLinesBetweenTagsShouldBeRead(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := TempParser{StartTag: &Tag{Name: startTag}, EndTag: &Tag{Name: endTag}}

	lineThatShouldBeFound := "This line should be found"

	results := []models.InputLine{}
	lines := []string{
		"This is first line, and it shouldn't be read",
		"Second line shouldn't be read",
		startTag,
		lineThatShouldBeFound, // Only this line should be in the results
		endTag,
		"Another insignificant line",
	}

	var err error
	for index, line := range lines {
		results, err = parser.Parse(index, line, results)
		if err != nil {
			t.Error(err)
		}
	}

	asserter.Equal(1, len(results))
	asserter.Equal(3, results[0].Number)
	asserter.Equal(lineThatShouldBeFound, results[0].Text)
}
