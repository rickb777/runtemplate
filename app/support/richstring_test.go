package support

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestRichStringNoDots(t *testing.T) {
	a := RichString(".foo.Bar.Baz.").NoDots()
	expect.String(a).ToBe(t, RichString("fooBarBaz"))
}

func TestRichStringDivideOr0A(t *testing.T) {
	a, b := RichString("foo,bar").DivideLastOr0(',')
	expect.String(a).ToBe(t, "foo")
	expect.String(b).ToBe(t, "bar")
}

func TestRichStringDivideOr0B(t *testing.T) {
	a, b := RichString("foo bar").DivideLastOr0(',')
	expect.String(a).ToBe(t, "foo bar")
	expect.String(b).ToBe(t, "")
}

func TestRichStringDivideOr1A(t *testing.T) {
	a, b := RichString("foo,bar").DivideLastOr1(',')
	expect.String(a).ToBe(t, "foo")
	expect.String(b).ToBe(t, "bar")
}

func TestRichStringDivideOr1B(t *testing.T) {
	a, b := RichString("foo bar").DivideLastOr1(',')
	expect.String(a).ToBe(t, "")
	expect.String(b).ToBe(t, "foo bar")
}

func TestRichStringRemoveBeforeA(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBeforeLast('/')
	expect.String(a).ToBe(t, "baz")
}

func TestRichStringRemoveBeforeB(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBeforeLast(',')
	expect.String(a).ToBe(t, "foo/bar/baz")
}
