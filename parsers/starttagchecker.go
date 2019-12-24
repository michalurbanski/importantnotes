package parsers

type StartTagChecker struct {
	//Tag        *Tag
	TagChecker TagChecker
	IsFound    bool
}

func (checker *StartTagChecker) IsTagFound() bool {
	return checker.IsFound
}

func (checker *StartTagChecker) GetTag() Tag {
	return checker.TagChecker.Tag
}

func (checker *StartTagChecker) Check(text string) (checkNext bool, outErr error) {
	match, err := checker.TagChecker.Check(text)
	if err != nil {
		return false, nil
	}

	if match {
		checker.IsFound = true
	}

	// Checker exists only because tag was not found.
	// So if it's not found in currently processed line, it means that we're before
	// start tag, and lines should not be read.

	return false, nil

	// match, err := isMatch(text, checker.Tag.Name)
	// if err != nil {
	// 	return false, err
	// }

	// if match {
	// 	checker.Tag.IsFound = true
	// }

	// // Checker exists only because tag was not found.
	// // So if it's not found in currently processed line, it means that we're before
	// // start tag, and lines should not be read.

	// // If tag was found in current line, this line also shouldn't be processed.
	// return false, nil
}
