package mytypes

import "strings"

// MyInt is a custom version of the `int` type.
type MyInt int

// MyString is a custom version of the `int` type.
type MyString string

// MultiString is a representation of a []string
type MultiString []string

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// Len returns the length of the string
func (s MyString) Len() int {
	return len(s)
}

// Join returns the MultiString joined by using
// the provided string as the "glue"
func (s MultiString) Join(glue string) string {
	return strings.Join(s, glue)
}
