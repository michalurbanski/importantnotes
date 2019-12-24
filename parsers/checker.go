package parsers

type Checker interface {
	Check(text string) (checkNext bool, outErr error)
	IsTagFound() bool
	GetTag() Tag
}
