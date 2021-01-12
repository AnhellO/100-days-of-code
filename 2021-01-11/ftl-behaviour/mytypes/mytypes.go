package mytypes

import "strings"

// MyInt is a custom version of the `int` type.
type MyInt int

// MyString is a custom version of the `int` type
type MyString string

// MultiString is a representation of a []string
type MultiString []string

// MyBuilder wraps a strings.Builder
type MyBuilder struct {
	strings.Builder
}

// StringUppercaser wraps a strings.Builder
type StringUppercaser struct {
	strings.Builder
}

// IntBuilder builds an int from a []ints
type IntBuilder struct {
	Integers []int
}

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// Double takes a pointer receiver, and doubles
// the value it points to
func (i *MyInt) Double() {
	*i *= 2
}

// MultiplyBy takes an MyInt parameter
// and multiplies its receiver by that value
func (i *MyInt) MultiplyBy(mi MyInt) {
	*i *= mi
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

// Hello returns the string "Hello, Gophers!".
func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

// ToUpper returns the wrapped string.Builder object
// converted to upper case.
func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.String())
}

// Sum returns the sum of the ints within the []int
func (ib IntBuilder) Sum() int {
	sum := 0
	for _, v := range ib.Integers {
		sum += v
	}

	return sum
}

// Square multiplies the original int value by itself
func Square(p *int) {
	*p *= *p
}

// SwapInts takes two pointers to int, and swaps the
// values of the variables pointed to
func SwapInts(a, b *int) {
	*a, *b = *b, *a
}
