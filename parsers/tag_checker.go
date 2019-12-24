package parsers

import (
	"fmt"
	"regexp"
)

type TagChecker struct {
	Tag Tag
}

func (checker TagChecker) Check(text string) (isTagFound bool, outErr error) {
	match, err := isMatch(text, checker.Tag.Name)
	if err != nil {
		return false, err
	}

	return match, nil
}

func (checker TagChecker) GetTag() Tag {
	return checker.Tag
}

func isMatch(text string, tag string) (isMatch bool, outErr error) {
	if len(tag) == 0 {
		return false, nil
	}

	match, err := regexp.MatchString(fmt.Sprintf("^%s", tag), text)
	if err != nil {
		return false, err
	}

	return match, nil
}
