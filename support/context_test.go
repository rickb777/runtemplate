package support

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func expectPresent(t *testing.T, ctx map[string]interface{}, key string) {
	if _, ok := ctx[key]; !ok {
		t.Fatalf("Missing %s; Got len %d %+v", key, len(ctx), ctx)
	}
	delete(ctx, key)
}

func TestCreateContext(t *testing.T) {
	m := FileMeta{"/a/b/c", "foo", time.Time{}}
	types := Pairs([]Pair{{"Type", "big.Int"}, {"B", "*FooBar"}, {"C", "vv3"}})
	others := Pairs([]Pair{{"I1", "X1"}, {"I1", "X2"}, {"I1", "X3"}})
	ctx := CreateContext(m, "output.txt", types, others)

	if len(ctx) != 35 {
		t.Fatalf("Got len %d %+v", len(ctx), ctx)
	}

	expectPresent(t, ctx, "PWD")
	expectPresent(t, ctx, "GOOS")
	expectPresent(t, ctx, "GOROOT")
	expectPresent(t, ctx, "GOARCH")
	expectPresent(t, ctx, "GOPATH")

	if !reflect.DeepEqual(ctx, map[string]interface{}{
		"OutFile":      "output.txt",
		"Type":         "big.Int",
		"UType":        "BigInt",
		"LType":        "bigInt",
		"PType":        "big.Int",
		"Prefix":       "",
		"UPrefix":      "",
		"LPrefix":      "",
		"B":            "FooBar",
		"UB":           "FooBar",
		"LB":           "fooBar",
		"PB":           "*FooBar",
		"C":            "vv3",
		"UC":           "Vv3",
		"LC":           "vv3",
		"PC":           "vv3",
		"HasType":      true,
		"HasB":         true,
		"HasC":         true,
		"TypeAmp":      "",
		"BAmp":         "&",
		"CAmp":         "",
		"TypeStar":     "",
		"BStar":        "*",
		"CStar":        "",
		"Package":      "support",
		"TemplatePath": "/a/b/c",
		"TemplateFile": "foo",
		"I1":           []string{"X1", "X2", "X3"},
		"HasI1":        true,
	}) {
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
