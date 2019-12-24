package parsers

import (
	"fmt"
	"importantnotes/models"
)

// StartEndParser is used to parse lines when start tag and/or end tag are provided in the configuration.
// It contains information on how Start and End tags are set.
type StartEndParser struct {
	StartTag Tag
	EndTag   Tag
}

// SetTagFound sets tag as found when it's encountered in a line.
func (parser *StartEndParser) SetTagFound(tag Tag) {
	switch tag.Name {
	case parser.StartTag.Name:
		parser.StartTag.IsFound = true
	case parser.EndTag.Name:
		parser.EndTag.IsFound = true
	default:
		panic(fmt.Sprintf("Incorrect tag name, found %s", tag.Name))
	}
}

// ParseLine adds line to the results only when it's between start and end tags.
func (parser *StartEndParser) ParseLine(lineNumber int, text string, results []models.InputLine) ([]models.InputLine, error) {
	read, err := parser.shouldRead(text)
	if err != nil {
		return results, err
	}

	if read {
		line := models.InputLine{Number: lineNumber, Text: text}
		results = append(results, line)
	}

	return results, nil
}

// shouldRead determines whether line should be read or skipped.
func (parser *StartEndParser) shouldRead(text string) (bool, error) {
	shouldRead, checkers := parser.validate()
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
		if shouldRead { // necessary as it can be either true of false at this stage
			if shouldRead, err = checker.Check(text); err != nil {
				break
			}

			if checker.IsTagFound() {
				parser.SetTagFound(checker.GetTag())
			}
		}

		// isTagFound, err := checker.Check(text)
		// if err != nil {
		// 	break
		// }

		// If tag was found in the line, then this line should not be read
		// as it does not contain notes.
		// if isTagFound {
		// 	parser.SetTagFound(checker.GetTag())
		// 	return false, nil
		// }
		// if shouldRead {
		// 	if shouldRead, err = checker.Check(text); err != nil {
		// 		break
		// 	}
		// }
	}

	return shouldRead, err
}

// Validate checks whether before reading next line it can be already be determined if it should be read or not.
// If it cannot be determined, then returns collection of checkers that should be used to check whether line should be read.
func (parser *StartEndParser) validate() (shouldRead bool, checkers []Checker) {
	cks := []Checker{}

	// If end tag was already found, then no more lines should be read.
	if parser.EndTag.IsFound {
		return false, cks
	}

	// If there are no tags provided, then each line should be read.
	if parser.StartTag.IsEmpty() && parser.EndTag.IsEmpty() {
		return true, cks
	}

	if parser.StartTag.IsEmpty() == false && parser.StartTag.IsFound == false {
		//cks = append(cks, &StartTagChecker{Tag: parser.StartTag})
		cks = append(cks, &StartTagChecker{TagChecker: TagChecker{Tag: parser.StartTag}})
	}

	if parser.EndTag.IsEmpty() == false && parser.EndTag.IsFound == false {
		//cks = append(cks, &EndTagChecker{Tag: parser.EndTag})
		//cks = append(cks, &TagChecker{Tag: parser.EndTag})
		cks = append(cks, &EndTagChecker{TagChecker: TagChecker{Tag: parser.EndTag}})
	}

	return true, cks
}

// func isMatch(text string, tag string) (isMatch bool, outErr error) {
// 	if len(tag) == 0 {
// 		return false, nil
// 	}

// 	match, err := regexp.MatchString(fmt.Sprintf("^%s", tag), text)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false, err
// 	}

// 	return match, nil
// }
