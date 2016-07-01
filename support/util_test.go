package support

import (
	"testing"
	"reflect"
)

func TestRichStringDivideOr0A(t *testing.T) {
	a, b := RichString("foo,bar").DivideOr0(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got '%s'", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got '%s'", b)
	}
}

func TestRichStringDivideOr0B(t *testing.T) {
	a, b := RichString("foo bar").DivideOr0(',')
	if a != "foo bar" {
		t.Errorf("Want 'foo bar', got '%s'", a)
	}
	if b != "" {
		t.Errorf("Want '', got '%s'", b)
	}
}

func TestRichStringDivideOr1A(t *testing.T) {
	a, b := RichString("foo,bar").DivideOr1(',')
	if a != "foo" {
		t.Errorf("Want 'foo', got '%s'", a)
	}
	if b != "bar" {
		t.Errorf("Want 'bar', got '%s'", b)
	}
}

func TestRichStringDivideOr1B(t *testing.T) {
	a, b := RichString("foo bar").DivideOr1(',')
	if a != "" {
		t.Errorf("Want 'foo bar', got '%s'", a)
	}
	if b != "foo bar" {
		t.Errorf("Want '', got '%s'", b)
	}
}

func TestRichStringRemoveBeforeA(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBefore('/')
	if a != "baz" {
		t.Errorf("Want 'baz', got '%s'", a)
	}
}

func TestRichStringRemoveBeforeB(t *testing.T) {
	a := RichString("foo/bar/baz").RemoveBefore(',')
	if a != "foo/bar/baz" {
		t.Errorf("Want 'foo/bar/baz', got '%s'", a)
	}
}

func TestFindTemplateArg1(t *testing.T) {
	tpl, args := FindTemplateArg("", nil)
	if tpl != "" {
		t.Errorf("Want '', got '%s'", tpl)
	}
	if len(args) != 0 {
		t.Errorf("Want [], got %+v", args)
	}
}

func TestFindTemplateArg2(t *testing.T) {
	tpl, args := FindTemplateArg("", []string{"x=1", "y=2", "f.tpl", "z=3"})
	if tpl != "f.tpl" {
		t.Errorf("Want 'f.tpl', got '%s'", tpl)
	}
	if !reflect.DeepEqual(args, []string{"x=1", "y=2", "z=3"}) {
		t.Fatalf("Got %+v", args)
	}
}

func TestSplitKeyValArgs1(t *testing.T) {
	pairs, left := SplitKeyValArgs(nil)
	if len(pairs) != 0 {
		t.Errorf("Want [], got '%+v'", pairs)
	}
	if len(left) != 0 {
		t.Errorf("Want [], got %+v", left)
	}
}

func TestSplitKeyValArgs2(t *testing.T) {
	pairs, left := SplitKeyValArgs([]string{"x=1", "y=2", "foo", "z=3", ""})
	if !reflect.DeepEqual(pairs, Pairs([]Pair{{"x", "1"}, {"y", "2"}, {"z", "3"}})) {
		t.Fatalf("Got %+v", pairs)
	}
	if !reflect.DeepEqual(left, []string{"foo", ""}) {
		t.Fatalf("Got %+v", left)
	}
}

