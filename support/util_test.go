package support

import (
	"reflect"
	"testing"
)

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
