package support

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestFindTemplateArg1(t *testing.T) {
	tpl, args := FindTemplateArg("", nil)
	expect.Any(tpl).ToBe(t, "")
	expect.Slice(args).ToHaveLength(t, 0)
}

func TestFindTemplateArg2(t *testing.T) {
	tpl, args := FindTemplateArg("", []string{"x=1", "y=2", "f.tpl", "z=3"})
	expect.Any(tpl).ToBe(t, "f.tpl")
	expect.Slice(args).ToBe(t, "x=1", "y=2", "z=3")
}

func TestSplitKeyValArgs1(t *testing.T) {
	types, others, left := SplitKeyValArgs(nil)
	expect.Slice(types).ToHaveLength(t, 0)
	expect.Slice(others).ToHaveLength(t, 0)
	expect.Slice(left).ToHaveLength(t, 0)
}

func TestSplitKeyValArgs2(t *testing.T) {
	types, others, left := SplitKeyValArgs([]string{"w=t1", "x=t2", "foo", "y=xyz/abc/123", "a:v1", "z=t3", "b:v2", `c:a\n\tb`, "q:", ""})

	expect.Any(types).ToBe(t, []Tuple{NewTuple("w=t1"), NewTuple("x=t2"), NewTuple("y=xyz/abc/123"), NewTuple("z=t3")})
	expect.Any(others).ToBe(t, []Pair{NewPair("a:v1"), NewPair("b:v2"), NewPair("c:a\n\tb"), NewPair("q:")})
	expect.Slice(left).ToBe(t, "foo", "")
}

func TestSplitKeyValArgs3(t *testing.T) {
	types, others, left := SplitKeyValArgs([]string{"Type=date.PeriodOfDays/PeriodOfDays/0", "Comparable:true", "Ordered:true", "Numeric:true", `Import:"github.com/rickb777/date"`})

	expect.Any(types).ToBe(t, []Tuple{{Key: "Type", Type: Type{s: "date.PeriodOfDays", ident: "PeriodOfDays", zero: "0"}}})
	expect.Any(others).ToBe(t, []Pair{NewPair("Comparable:true"), NewPair("Ordered:true"), NewPair("Numeric:true"), NewPair(`Import:"github.com/rickb777/date"`)})
	expect.Any(left).ToBeNil(t)
}
