package test

import (
	"reflect"
	"testing"
)

// Asserter wraps *testing.T object.
type Asserter struct {
	T *testing.T
}

// Equal checks if two values are equal.
//
// a - expected
//
// b - actual
func (m *Asserter) Equal(a, b interface{}) {
	// https://www.tobstarr.com/2017/06/16/better-test-helpers-in-go/
	m.T.Helper()

	if a != b {
		m.T.Errorf("expected %#v (%T), was %#v (%T)", a, a, b, b)
	}
}

// IsNil checks if value (or pointer) is nil.
//
// a - actual value that should be nil
func (m *Asserter) IsNil(a interface{}) {
	m.T.Helper()

	var isNil = a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil())
	if !isNil {
		m.T.Errorf("%#v expected to be nil, but found not nil", a)
	}
}

// IsNotNil checks if value (or pointer) is not nil.
//
// a - actual value that should not be nil
func (m *Asserter) IsNotNil(a interface{}) {
	m.T.Helper()

	var isNil = a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil())
	if isNil {
		m.T.Errorf("%#v expected not to be nil, but found nil", a)
	}
}
