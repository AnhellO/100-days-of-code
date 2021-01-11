package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestMyIntTwice(t *testing.T) {
	i := mytypes.MyInt(10)
	want := mytypes.MyInt(20)
	got := i.Twice()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	s := mytypes.MyString("Hello, Gophers!")
	want := 15
	got := s.Len()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMultiStringJoin(t *testing.T) {
	m := mytypes.MultiString{"a", "b", "c"}
	want := "a plus b plus c"
	got := m.Join(" plus ")

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestStringUppercaser(t *testing.T) {
	m := mytypes.StringUppercaser{}
	m.WriteString("Testing Upper Case")
	want := "TESTING UPPER CASE"
	got := m.ToUpper()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestIntBuilderSum(t *testing.T) {
	m := mytypes.IntBuilder{Integers: []int{1, 2, 3, 4, 5, 6}}
	want := 21
	got := m.Sum()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
