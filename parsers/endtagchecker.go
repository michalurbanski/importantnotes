package parsers

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
