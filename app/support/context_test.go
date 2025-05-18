package support

import (
	"github.com/rickb777/expect"
	"strings"
	"testing"
	"time"
)

func expectPresent(t *testing.T, ctx map[string]any, key string) {
	expect.Map(ctx).ToContain(t, key)
	delete(ctx, key)
}

func TestCreateContextCore(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, ""}
	types := Tuples([]Tuple{})
	others := Pairs([]Pair{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")
	expectPresent(t, ctx, "AppVersion")

	exp := map[string]any{
		"OutFile":      "output.txt",
		"Package":      "support",
		"TemplatePath": "/a/b/c",
		"TemplateFile": "foo",
	}
	expect.Any(ctx).ToBe(t, exp)
}

func TestCreateContext(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, ""}
	b := NewTuple("B=*FooBar")
	c := NewTuple("C=vv3")
	types := Tuples([]Tuple{b, c})
	others := Pairs([]Pair{{"I1", "X1"}, {"I1", "X2"}, {"I1", "X3"}})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")
	expectPresent(t, ctx, "OutFile")
	expectPresent(t, ctx, "AppVersion")
	expectPresent(t, ctx, "TemplatePath")
	expectPresent(t, ctx, "TemplateFile")
	expectPresent(t, ctx, "Package")
	expectPresent(t, ctx, "I1")

	exp := map[string]any{
		"B": b.Type,
		//"UB":     "FooBar",
		//"LB":     "fooBar",
		//"PB":     "*FooBar",
		"C": c.Type,
		//"UC":     "Vv3",
		//"LC":     "vv3",
		//"PC":     "vv3",
		"HasB":  true,
		"HasC":  true,
		"HasI1": true,
		//"BAmp":   "&",
		//"CAmp":   "",
		//"BStar":  "*",
		//"CStar":  "",
		//"BIsPtr": true,
		//"CIsPtr": false,
		//"BZero":  "nil",
		//"CZero":  "*(new(vv3))",
	}
	expect.Any(ctx).ToBe(t, exp)
}

func TestCreateContextWithDottedType(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, ""}
	bigInt := NewTuple("Type=*big.Int")
	types := Tuples([]Tuple{bigInt})
	others := Pairs([]Pair{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")
	expectPresent(t, ctx, "OutFile")
	expectPresent(t, ctx, "AppVersion")
	expectPresent(t, ctx, "TemplatePath")
	expectPresent(t, ctx, "TemplateFile")
	expectPresent(t, ctx, "Package")

	exp := map[string]any{
		"Type": bigInt.Type,
		//"UType":     "BigInt",
		//"LType":     "bigInt",
		//"PType":     "*big.Int",
		"Prefix": NewTuple("Prefix=").Type,
		//"UPrefix":   "",
		//"LPrefix":   "",
		"HasType": true,
		//"TypeIsPtr": true,
		//"TypeAmp":   "&",
		//"TypeStar":  "*",
		//"TypeZero":  "nil",
	}
	expect.Any(ctx).ToBe(t, exp)
}

func TestCreateContextWithPrefix(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, ""}
	types := Tuples([]Tuple{NewTuple("OneType=Apple"), NewTuple("TwoType=Pear/Pear/nil"), NewTuple("OnePrefix=Foo")})
	others := Pairs([]Pair{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")
	expectPresent(t, ctx, "OutFile")
	expectPresent(t, ctx, "AppVersion")
	expectPresent(t, ctx, "TemplatePath")
	expectPresent(t, ctx, "TemplateFile")
	expectPresent(t, ctx, "Package")

	exp := map[string]any{
		"OneType": NewTuple("OneType=Apple").Type,
		"TwoType": NewTuple("TwoType=Pear/Pear/nil").Type,
		//"UOneType":     "Apple",
		//"UTwoType":     "Pear",
		//"LOneType":     "apple",
		//"LTwoType":     "pear",
		//"POneType":     "Apple",
		//"PTwoType":     "Pear",
		"OnePrefix": NewTuple("OnePrefix=Foo").Type,
		"TwoPrefix": NewTuple("TwoPrefix=").Type,
		//"UOnePrefix":   "Foo",
		//"UTwoPrefix":   "",
		//"LOnePrefix":   "foo",
		//"LTwoPrefix":   "",
		"HasOneType":   true,
		"HasTwoType":   true,
		"HasOnePrefix": true,
		//"OneTypeIsPtr": false,
		//"TwoTypeIsPtr": false,
		//"OneTypeAmp":   "",
		//"TwoTypeAmp":   "",
		//"OneTypeStar":  "",
		//"TwoTypeStar":  "",
		//"OneTypeZero":  "*(new(Apple))",
		//"TwoTypeZero":  "nil",
	}
	expect.Any(ctx).ToBe(t, exp)
}

func TestChoosePackage(t *testing.T) {
	wd, pkg := choosePackage("foo.go")
	expect.Bool(strings.HasSuffix(wd, pkg)).ToBeTrue(t)

	wd, pkg = choosePackage("aaa/foo.go")
	expect.Any(pkg).ToBe(t, "aaa")
	expect.Bool(strings.HasSuffix(wd, pkg)).ToBeFalse(t)

	wd, pkg = choosePackage("bbb/aaa/foo.go")
	expect.Any(pkg).ToBe(t, "aaa")
	expect.Bool(strings.HasSuffix(wd, pkg)).ToBeFalse(t)

	wd, pkg = choosePackage("./foo.go")
	expect.Bool(strings.HasSuffix(wd, pkg)).ToBeTrue(t)
}

func diffMaps(t *testing.T, a, b map[string]any) {
	t.Helper()
	if len(a) > len(b) {
		diffMaps(t, b, a)
	} else {
		for k, vb := range b {
			va, ok := a[k]
			if !ok {
				t.Logf("Missing: %s\n", k)
			} else if va != vb {
				t.Logf("Differ: %s: %v and %v\n", k, va, vb)
			}
		}
	}
}
