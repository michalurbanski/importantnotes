package test

import "testing"

// Asserter wraps *testing.T object.
type Asserter struct {
	T *testing.T
}

// Equal checks if two values are equal.
// a - expected
// b - actual
func (m *Asserter) Equal(a, b interface{}) {
	// https://www.tobstarr.com/2017/06/16/better-test-helpers-in-go/
	m.T.Helper()

	if a != b {
		m.T.Errorf("Expected %#v (%T), was %#v (%T)", a, a, b, b)
	}
}
