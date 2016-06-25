package main

import (
	"strings"
	"testing"
)

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
