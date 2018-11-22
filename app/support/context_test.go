package support

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func expectPresent(t *testing.T, ctx map[string]interface{}, key string) {
	t.Helper()
	if _, ok := ctx[key]; !ok {
		t.Fatalf("Missing %s; Got len %d %+v", key, len(ctx), ctx)
	}
	delete(ctx, key)
}

func TestCreateContextCore(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Triples([]Triple{})
	others := Triples([]Triple{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 10 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")
	expectPresent(t, ctx, "AppVersion")

	exp := map[string]interface{}{
		"OutFile":      "output.txt",
		"Package":      "support",
		"TemplatePath": "/a/b/c",
		"TemplateFile": "foo",
	}
	if !reflect.DeepEqual(ctx, exp) {
		diffMaps(t, ctx, exp)
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}
}

func TestCreateContext(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Triples([]Triple{{"B", "*FooBar", ""}, {"C", "vv3", ""}})
	others := Triples([]Triple{{"I1", "X1", ""}, {"I1", "X2", ""}, {"I1", "X3", ""}})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 30 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

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
	if !reflect.DeepEqual(ctx, exp) {
		diffMaps(t, ctx, exp)
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}
}

func TestCreateContextWithDottedType(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}, false}
	types := Triples([]Triple{{"Type", "*big.Int", ""}})
	others := Triples([]Triple{})
	ctx := CreateContext(m, "output.txt", types, others, "(app version)")

	if len(ctx) != 22 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

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
	if !reflect.DeepEqual(ctx, exp) {
		diffMaps(t, ctx, exp)
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}
}

func TestChoosePackage(t *testing.T) {
	wd, pkg := choosePackage("foo.go")
	if !strings.HasSuffix(wd, pkg) {
		t.Errorf("wd=%s, pkg=%s", wd, pkg)
	}

	wd, pkg = choosePackage("aaa/foo.go")
	if pkg != "aaa" {
		t.Errorf("Want aaa, got %s", pkg)
	}
	if strings.HasSuffix(wd, pkg) {
		t.Errorf("wd=%s, pkg=%s", wd, pkg)
	}

	wd, pkg = choosePackage("bbb/aaa/foo.go")
	if pkg != "aaa" {
		t.Errorf("Want aaa, got %s", pkg)
	}
	if strings.HasSuffix(wd, pkg) {
		t.Errorf("wd=%s, pkg=%s", wd, pkg)
	}

	wd, pkg = choosePackage("./foo.go")
	if !strings.HasSuffix(wd, pkg) {
		t.Errorf("wd=%s, pkg=%s", wd, pkg)
	}
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
