package parsers

import (
	"fmt"
	"importantnotes/models"
	"regexp"
)

type Tag struct {
	Name    string
	IsFound bool
}

func (t *Tag) IsEmpty() bool {
	return len(t.Name) == 0
}

type TempParser struct {
	StartTag *Tag
	EndTag   *Tag
}

type Checker interface {
	Check(text string) (bool, error)
}

type StartTagChecker struct {
	Tag *Tag
}

func (checker *StartTagChecker) Check(text string) (checkNext bool, outErr error) {
	match, err := isMatch(text, checker.Tag.Name)
	if err != nil {
		return false, err
	}

	if match {
		checker.Tag.IsFound = true
	}

	// Checker exists only because tag was not found.
	// So if it's not found in currently processed line, it means that we're before
	// start tag, and lines should not be read.

	// If tag was found in current line, this line also shouldn't be processed.
	return false, nil
}

type EndTagChecker struct {
	Tag *Tag
}

func (checker *EndTagChecker) Check(text string) (checkNext bool, outErr error) {
	match, err := isMatch(text, checker.Tag.Name)
	if err != nil {
		return false, err
	}

	if match {
		checker.Tag.IsFound = true
		return false, nil
	}

	// No tag in the current line
	return true, nil
}

func (parser TempParser) ParseLine(lineNumber int, text string, results []models.InputLine) ([]models.InputLine, error) {
	read, err := parser.shouldR(text)
	if err != nil {
		fmt.Println(err)
		return results, err
	}

	if read {
		line := models.InputLine{Number: lineNumber, Text: text}
		results = append(results, line)
	}

	return results, nil
}

// Validate checks whether before reading next line it can be already be determined if it should be read or not.
// If it cannot be determined, then returns collection of checkers that should be used to check whether line should be read.
func (parser TempParser) Validate() (shouldRead bool, checkers []Checker) {
	cks := []Checker{}

	// If there are not tags provided, then each line should be read.
	if parser.StartTag.IsEmpty() && parser.EndTag.IsEmpty() {
		return true, cks
	}

	// If end tag was already found, then no more lines should be read.
	if parser.EndTag.IsFound {
		return false, cks
	}

	if parser.StartTag.IsEmpty() == false && parser.StartTag.IsFound == false {
		cks = append(cks, &StartTagChecker{Tag: parser.StartTag})
	}

	if parser.EndTag.IsEmpty() == false && parser.EndTag.IsFound == false {
		cks = append(cks, &EndTagChecker{Tag: parser.EndTag})
	}

	return true, cks
}

func (parser TempParser) shouldR(text string) (bool, error) {
	shouldRead, checkers := parser.Validate()
	if shouldRead == false {
		return false, nil
	}

	if shouldRead && len(checkers) == 0 {
		return true, nil
	}

	// Line should be further checked and checkers exist - use them to check if line should be read.
	// NOTE: if first checker finds tag, than the second one does not hve to be checked
	// as both tags cannot exist in the same line.
	var err error
	for _, checker := range checkers {
		if shouldRead {
			if shouldRead, err = checker.Check(text); err != nil {
				break
			}
		}
	}

	return shouldRead, nil
}

func isMatch(text string, tag string) (isMatch bool, outErr error) {
	if len(tag) == 0 {
		return false, nil
	}

	match, err := regexp.MatchString(fmt.Sprintf("^%s", tag), text)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return match, nil
}
