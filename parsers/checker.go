package parsers

type Checker interface {
	Check(text string) (bool, error)
}
