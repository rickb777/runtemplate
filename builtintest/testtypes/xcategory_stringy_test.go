package testtypes

import (
	"testing"
	"sort"
	"reflect"
)

func TestString(t *testing.T) {
	a := Category("foo")
	s := a.String()

	if s != "foo" {
		t.Errorf("Expected foo but got %#v", s)
	}
}

func TestPtr(t *testing.T) {
	a := Category("foo")
	s := a.Ptr()

	if *s != "foo" {
		t.Errorf("Expected foo but got %#v", s)
	}
}

func TestTrimSpace(t *testing.T) {
	a := Category(" foo ")
	s := a.TrimSpace()

	if s != "foo" {
		t.Errorf("Expected foo but got %#v", s)
	}
}

func TestToLower(t *testing.T) {
	a := Category("FOO")
	s := a.ToLower()

	if s != "foo" {
		t.Errorf("Expected foo but got %#v", s)
	}
}

func TestToUpper(t *testing.T) {
	a := Category("foo")
	s := a.ToUpper()

	if s != "FOO" {
		t.Errorf("Expected foo but got %#v", s)
	}
}

func TestScan1(t *testing.T) {
	a := new(Category)
	e := a.Scan("foo")
	if e != nil {
		t.Errorf("Expected nil but got %v", e)
	}

	if a.String() != "foo" {
		t.Errorf("Expected foo but got %v", *a)
	}
}

func TestScan2(t *testing.T) {
	a := new(Category)
	e := a.Scan([]byte("foo"))
	if e != nil {
		t.Errorf("Expected nil but got %v", e)
	}

	if a.String() != "foo" {
		t.Errorf("Expected foo but got %v", *a)
	}
}

func TestScan3(t *testing.T) {
	a := new(Category)
	e := a.Scan(123)
	if e.Error() != "123" {
		t.Errorf("Expected error but got %#v", e)
	}
}

func TestValue(t *testing.T) {
	a := Category("foo")
	v, e := a.Value()
	if e != nil {
		t.Errorf("Expected nil but got %v", e)
	}

	if v.(string) != "foo" {
		t.Errorf("Expected foo but got %v", v)
	}
}

func TestMarshall(t *testing.T) {
	a := Category("foo")
	v, e := a.MarshalText()
	if e != nil {
		t.Errorf("Expected nil but got %v", e)
	}

	if string(v) != "foo" {
		t.Errorf("Expected foo but got %v", v)
	}
}

func TestUnmarshall(t *testing.T) {
	a := new(Category)
	e := a.UnmarshalText([]byte("foo"))
	if e != nil {
		t.Errorf("Expected nil but got %v", e)
	}

	if a.String() != "foo" {
		t.Errorf("Expected foo but got %v", *a)
	}
}

func TestSort(t *testing.T) {
	a := Category("a")
	b := Category("b")
	c := Category("c")
	s := []Category{b, c, a}
	sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})

	if !reflect.DeepEqual(s, []Category{a, b, c}) {
		t.Errorf("Expected [a b c] but got %v", s)
	}
}

