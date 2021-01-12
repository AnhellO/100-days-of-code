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

func TestSquare(t *testing.T) {
	want := 64
	got := 8
	mytypes.Square(&got)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSwapInts(t *testing.T) {
	x := 666
	y := 999
	mytypes.SwapInts(&x, &y)

	if x == 666 || y == 999 {
		t.Errorf("want x=999 and y=666, got x=%d and y=%d", x, y)
	}
}

func TestMyIntDouble(t *testing.T) {
	got := mytypes.MyInt(10)
	want := mytypes.MyInt(20)
	got.Double()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMyIntMultiplyBy(t *testing.T) {
	got := mytypes.MyInt(10)
	want := mytypes.MyInt(100)
	got.MultiplyBy(mytypes.MyInt(10))

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
