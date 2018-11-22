package support

import (
	"github.com/benmoss/matchers"
	. "github.com/onsi/gomega"
	"strings"
	"testing"
	"time"
)

func expectPresent(g *GomegaWithT, ctx map[string]interface{}, key string) {
	g.Expect(ctx).To(HaveKey(key))
	delete(ctx, key)
}

func TestCreateContextCore(t *testing.T) {
	g := NewGomegaWithT(t)

	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Types([]Type{})
	others := Pairs([]Pair{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 10 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

	expectPresent(g, ctx, "PWD")
	expectPresent(g, ctx, "GOOS")
	expectPresent(g, ctx, "GOROOT")
	expectPresent(g, ctx, "GOARCH")
	expectPresent(g, ctx, "GOPATH")
	expectPresent(g, ctx, "AppVersion")

	exp := map[string]interface{}{
		"OutFile":      "output.txt",
		"Package":      "support",
		"TemplatePath": "/a/b/c",
		"TemplateFile": "foo",
	}
	g.Expect(ctx).To(matchers.DeepEqual(exp))
}

func TestCreateContext(t *testing.T) {
	g := NewGomegaWithT(t)

	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Types([]Type{NewType("B=*FooBar"), NewType("C=vv3")})
	others := Pairs([]Pair{{"I1", "X1"}, {"I1", "X2"}, {"I1", "X3"}})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 30 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

	expectPresent(g, ctx, "PWD")
	expectPresent(g, ctx, "GOOS")
	expectPresent(g, ctx, "GOROOT")
	expectPresent(g, ctx, "GOARCH")
	expectPresent(g, ctx, "GOPATH")
	expectPresent(g, ctx, "OutFile")
	expectPresent(g, ctx, "AppVersion")
	expectPresent(g, ctx, "TemplatePath")
	expectPresent(g, ctx, "TemplateFile")
	expectPresent(g, ctx, "Package")
	expectPresent(g, ctx, "I1")

	exp := map[string]interface{}{
		"B":      "FooBar",
		"UB":     "FooBar",
		"LB":     "fooBar",
		"PB":     "*FooBar",
		"C":      "vv3",
		"UC":     "Vv3",
		"LC":     "vv3",
		"PC":     "vv3",
		"HasB":   true,
		"HasC":   true,
		"HasI1":  true,
		"BAmp":   "&",
		"CAmp":   "",
		"BStar":  "*",
		"CStar":  "",
		"BIsPtr": true,
		"CIsPtr": false,
		"BZero":  "nil",
		"CZero":  "*(new(vv3))",
	}
	g.Expect(ctx).To(matchers.DeepEqual(exp))
}

func TestCreateContextWithDottedType(t *testing.T) {
	g := NewGomegaWithT(t)

	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Types([]Type{NewType("Type=*big.Int")})
	others := Pairs([]Pair{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 22 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

	expectPresent(g, ctx, "PWD")
	expectPresent(g, ctx, "GOOS")
	expectPresent(g, ctx, "GOROOT")
	expectPresent(g, ctx, "GOARCH")
	expectPresent(g, ctx, "GOPATH")
	expectPresent(g, ctx, "OutFile")
	expectPresent(g, ctx, "AppVersion")
	expectPresent(g, ctx, "TemplatePath")
	expectPresent(g, ctx, "TemplateFile")
	expectPresent(g, ctx, "Package")

	exp := map[string]interface{}{
		"Type":      "big.Int",
		"UType":     "BigInt",
		"LType":     "bigInt",
		"PType":     "*big.Int",
		"Prefix":    "",
		"UPrefix":   "",
		"LPrefix":   "",
		"HasType":   true,
		"TypeIsPtr": true,
		"TypeAmp":   "&",
		"TypeStar":  "*",
		"TypeZero":  "nil",
	}
	g.Expect(ctx).To(matchers.DeepEqual(exp))
}

func TestChoosePackage(t *testing.T) {
	g := NewGomegaWithT(t)

	wd, pkg := choosePackage("foo.go")
	g.Expect(strings.HasSuffix(wd, pkg)).To(BeTrue())

	wd, pkg = choosePackage("aaa/foo.go")
	g.Expect(pkg).To(Equal("aaa"))
	g.Expect(strings.HasSuffix(wd, pkg)).To(BeFalse())

	wd, pkg = choosePackage("bbb/aaa/foo.go")
	g.Expect(pkg).To(Equal("aaa"))
	g.Expect(strings.HasSuffix(wd, pkg)).To(BeFalse())

	wd, pkg = choosePackage("./foo.go")
	g.Expect(strings.HasSuffix(wd, pkg)).To(BeTrue())
}

func diffMaps(t *testing.T, a, b map[string]interface{}) {
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
