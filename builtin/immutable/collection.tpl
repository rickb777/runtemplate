// Generated from {{.TemplateFile}} with Type={{.Type.Name}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} Mutable:disabled
// by runtemplate {{.AppVersion}}
// See https://github.com/johanbrandhorst/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

{{if .HasImport}}
import (
	{{.Import}}
)

{{end -}}
// {{.Prefix.U}}{{.Type.U}}Sizer defines an interface for sizing methods on {{.Type.Name}} collections.
type {{.Prefix.U}}{{.Type.U}}Sizer interface {
	// IsEmpty tests whether {{.Prefix.U}}{{.Type.U}}Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether {{.Prefix.U}}{{.Type.U}}Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}
{{- if .Stringer}}

// {{.Prefix.U}}{{.Type.U}}MkStringer defines an interface for stringer methods on {{.Type.Name}} collections.
type {{.Prefix.U}}{{.Type.U}}MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}
{{- end}}

// {{.Prefix.U}}{{.Type.U}}Collection defines an interface for common collection methods on {{.Type.Name}}.
type {{.Prefix.U}}{{.Type.U}}Collection interface {
	{{.Prefix.U}}{{.Type.U}}Sizer
{{- if .Stringer}}
	{{.Prefix.U}}{{.Type.U}}MkStringer
{{- end}}

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool
{{- if .ToList}}

    // ToList returns a shallow copy as a list.
    ToList() *{{.Prefix.U}}{{.Type.U}}List
{{- end}}
{{- if .ToSet}}

    // ToSet returns a shallow copy as a set.
    ToSet() *{{.Prefix.U}}{{.Type.U}}Set
{{- end}}

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []{{.Type.Name}}

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of {{.Prefix.U}}{{.Type.U}}Collection return true for the predicate p.
	Exists(p func({{.Type.Name}}) bool) bool

	// Forall verifies that all elements of {{.Prefix.U}}{{.Type.U}}Collection return true for the predicate p.
	Forall(p func({{.Type.Name}}) bool) bool

	// Foreach iterates over {{.Prefix.U}}{{.Type.U}}Collection and executes the function f against each element.
	Foreach(f func({{.Type.Name}}))

	// Find returns the first {{.Type.Name}} that returns true for the predicate p.
	// False is returned if none match.
	Find(p func({{.Type.Name}}) bool) ({{.Type.Name}}, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan {{.Type.Name}}

	// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}Collection that return true for the predicate p.
	CountBy(p func({{.Type.Name}}) bool) int
{{- if .Comparable}}

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v {{.Type.Name}}) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...{{.Type.Name}}) bool
{{- end}}
{{- if .Ordered}}

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() {{.Type.Name}}

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() {{.Type.Name}}
{{- end}}

	// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func({{.Type.Name}}, {{.Type.Name}}) bool) {{.Type.Name}}

	// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func({{.Type.Name}}, {{.Type.Name}}) bool) {{.Type.Name}}
{{- if .Numeric}}

	// Sum returns the sum of all the elements in the collection.
	Sum() {{.Type.Name}}
{{- end}}
}
