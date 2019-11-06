package parsers

import (
	"importantnotes/helpers/test"
	"importantnotes/models"
	"testing"
)

// Here, for test purposes, it doesn't matter what value is set.
// In application configuration file these values will be more unique, comparing
// to the real content that might appear.
var startTag = "aaa"
var endTag = "bbb"

func TestLineWithStartTagIsNotIncludedInResults(t *testing.T) {
	asserter := test.Asserter{T: t}
	parser := StartEndLineParser{StartTag: startTag}

	results := []models.InputLine{}
	results = parser.ParseLine(1, startTag, results)

	asserter.Equal(0, len(results))
}

// TODO: add test for end tag
// TODO: add tests for skipping lines before start tag, and after end tag.
