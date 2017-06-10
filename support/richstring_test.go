package support

import (
	"testing"
)

func TestRichStringNoDots(t *testing.T) {
	a := RichString(".foo.Bar.Baz.").NoDots()
	if a != RichString("fooBarBaz") {
		t.Errorf("Want 'fooBarBaz', got %q", a)
	}
}

func TestRichStringDivideOr0A(t *testing.T) {
	a, b := RichString("foo,bar").DivideLastOr0(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got %q", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got %q", b)
	}
}

func TestRichStringDivideOr0B(t *testing.T) {
	a, b := RichString("foo bar").DivideLastOr0(',')
	if a != "foo bar" {
		t.Errorf("Want 'foo bar', got %q", a)
	}
	if b != "" {
		t.Errorf("Want '', got %q", b)
	}
}

func TestRichStringDivideOr1A(t *testing.T) {
	a, b := RichString("foo,bar").DivideLastOr1(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got %q", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got %q", b)
	}
}

func TestRichStringDivideOr1B(t *testing.T) {
	a, b := RichString("foo bar").DivideLastOr1(',')
	if a != "" {
		t.Errorf("Want 'foo bar', got %q", a)
	}
	if b != "foo bar" {
		t.Errorf("Want '', got %q", b)
	}
}

func TestRichStringRemoveBeforeA(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBeforeLast('/')
	if a != "baz" {
		t.Errorf("Want 'baz', got %q", a)
	}
}

func TestRichStringRemoveBeforeB(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBeforeLast(',')
	if a != "foo/bar/baz" {
		t.Errorf("Want 'foo/bar/baz', got %q", a)
	}
}
