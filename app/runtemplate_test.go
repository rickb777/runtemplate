package app_test

import (
	"github.com/rickb777/expect"
	"github.com/rickb777/runtemplate/v4/app"
	"github.com/rickb777/runtemplate/v4/app/support"
	"os"
	"testing"
)

func TestHappy(t *testing.T) {
	var failMessageArgs []any
	support.Fail = func(args ...any) {
		failMessageArgs = args
	}

	app.Generate("test.tpl", "test-out.txt", true, nil,
		support.Tuples{{
			Key:  "Type",
			Type: support.NewType("Foo"),
		}}, nil, "test")

	expect.Slice(failMessageArgs).ToBeNil(t)
	bs, err := os.ReadFile("test-out.txt")
	expect.Error(err).ToBeNil(t)
	expect.String(bs).ToEqual(t,
		`// Foo is a specialised kind of string.
type Foo string

// Ptr returns the address of a Foo.
func (foo Foo) Ptr() *Foo {
	return &foo
}

// String converts to a string and implements fmt.Stringer.
func (foo Foo) String() string {
	return string(foo)
}

// FooSlice attaches the methods of sort.Interface to []Foo, sorting in increasing order.
type FooSlice []Foo
`)

	os.Remove("test-out.txt")
}

func TestNotFound(t *testing.T) {
	var failMessageArgs []any
	support.Fail = func(args ...any) {
		failMessageArgs = args
	}

	app.Generate("missing.tpl", "test-out.txt", true, nil,
		support.Tuples{{
			Key:  "Type",
			Type: support.NewType("Foo"),
		}}, nil, "test")

	expect.Slice(failMessageArgs).ToContainAll(t, "not found:", "missing.tpl")
}
