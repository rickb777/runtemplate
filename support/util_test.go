package support

import (
	"reflect"
	"testing"
)

func TestRichStringNoDots(t *testing.T) {
	a := RichString("foo.Bar.Baz").NoDots()
	if a != RichString("fooBarBaz") {
		t.Errorf("Want 'fooBarBaz', got %q", a)
	}
}

func TestRichStringDivideOr0A(t *testing.T) {
	a, b := RichString("foo,bar").DivideOr0(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got %q", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got %q", b)
	}
}

func TestRichStringDivideOr0B(t *testing.T) {
	a, b := RichString("foo bar").DivideOr0(',')
	if a != "foo bar" {
		t.Errorf("Want 'foo bar', got %q", a)
	}
	if b != "" {
		t.Errorf("Want '', got %q", b)
	}
}

func TestRichStringDivideOr1A(t *testing.T) {
	a, b := RichString("foo,bar").DivideOr1(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got %q", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got %q", b)
	}
}

func TestRichStringDivideOr1B(t *testing.T) {
	a, b := RichString("foo bar").DivideOr1(',')
	if a != "" {
		t.Errorf("Want 'foo bar', got %q", a)
	}
	if b != "foo bar" {
		t.Errorf("Want '', got %q", b)
	}
}

func TestRichStringRemoveBeforeA(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBefore('/')
	if a != "baz" {
		t.Errorf("Want 'baz', got %q", a)
	}
}

func TestRichStringRemoveBeforeB(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBefore(',')
	if a != "foo/bar/baz" {
		t.Errorf("Want 'foo/bar/baz', got %q", a)
	}
}

func TestFindTemplateArg1(t *testing.T) {
	tpl, args := FindTemplateArg("", nil)
	if tpl != "" {
		t.Errorf("Want '', got %q", tpl)
	}
	if len(args) != 0 {
		t.Errorf("Want [], got %+v", args)
	}
}

func TestFindTemplateArg2(t *testing.T) {
	tpl, args := FindTemplateArg("", []string{"x=1", "y=2", "f.tpl", "z=3"})
	if tpl != "f.tpl" {
		t.Errorf("Want 'f.tpl', got %q", tpl)
	}
	if !reflect.DeepEqual(args, []string{"x=1", "y=2", "z=3"}) {
		t.Fatalf("Got %+v", args)
	}
}

func TestSplitKeyValArgs1(t *testing.T) {
	types, others, left := SplitKeyValArgs(nil)
	if len(types) != 0 {
		t.Errorf("Want [], got types '%+v'", types)
	}
	if len(others) != 0 {
		t.Errorf("Want [], got others '%+v'", others)
	}
	if len(left) != 0 {
		t.Errorf("Want [], got leftovers %+v", left)
	}
}

func TestSplitKeyValArgs2(t *testing.T) {
	types, others, left := SplitKeyValArgs([]string{"x=t1", "y=t2", "foo", "a:v1", "z=t3", "b:v2", ""})
	if !reflect.DeepEqual(types, Pairs([]Pair{{"x", "t1"}, {"y", "t2"}, {"z", "t3"}})) {
		t.Fatalf("Got types %+v", types)
	}
	if !reflect.DeepEqual(others, Pairs([]Pair{{"a", "v1"}, {"b", "v2"}})) {
		t.Fatalf("Got others %+v", others)
	}
	if !reflect.DeepEqual(left, []string{"foo", ""}) {
		t.Fatalf("Got leftovers %+v", left)
	}
}
