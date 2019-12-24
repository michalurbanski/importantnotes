package parsers

type EndTagChecker struct {
	TagChecker TagChecker
	IsFound    bool
}

func (checker *EndTagChecker) IsTagFound() bool {
	return checker.IsFound
}

func (checker *EndTagChecker) GetTag() Tag {
	return checker.TagChecker.Tag
}

func (checker *EndTagChecker) Check(text string) (checkNext bool, outErr error) {
	match, err := checker.TagChecker.Check(text)
	if err != nil {
		return false, nil
	}

	if match {
		checker.IsFound = true
	}

	// No tag in the current line - allow to read until it is found.
	return true, nil

	// match, err := isMatch(text, checker.Tag.Name)
	// if err != nil {
	// 	return false, err
	// }

	// if match {
	// 	checker.Tag.IsFound = true
	// 	return false, nil
	// }

	// // No tag in the current line - allow to read until it is found.
	// return true, nil
}
