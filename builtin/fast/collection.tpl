// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable={{.Comparable}} Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}}

package {{.Package}}

// {{.UPrefix}}{{.UType}}Sizer defines an interface for sizing methods on {{.PType}} collections.
type {{.UPrefix}}{{.UType}}Sizer interface {
	// IsEmpty tests whether {{.UPrefix}}{{.UType}}Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether {{.UPrefix}}{{.UType}}Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

{{if .Stringer}}
// {{.UPrefix}}{{.UType}}MkStringer defines an interface for stringer methods on {{.PType}} collections.
type {{.UPrefix}}{{.UType}}MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

}

{{end -}}
// {{.UPrefix}}{{.UType}}Collection defines an interface for common collection methods on {{.PType}}.
type {{.UPrefix}}{{.UType}}Collection interface {
	{{.UPrefix}}{{.UType}}Sizer
{{if .Stringer}}
	{{.UPrefix}}{{.UType}}MkStringer
{{end}}

	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []{{.PType}}

	// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}Collection return true for the passed func.
	Exists(fn func({{.PType}}) bool) bool

	// Forall verifies that all elements of {{.UPrefix}}{{.UType}}Collection return true for the passed func.
	Forall(fn func({{.PType}}) bool) bool

	// Foreach iterates over {{.UPrefix}}{{.UType}}Collection and executes the passed func against each element.
	Foreach(fn func({{.PType}}))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan {{.PType}}

	// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Collection that return true for the passed predicate.
	CountBy(predicate func({{.PType}}) bool) int

	// MinBy returns an element of {{.UPrefix}}{{.UType}}Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}

	// MaxBy returns an element of {{.UPrefix}}{{.UType}}Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}

	{{if .Ordered}}
	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() {{.Type}}

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() {{.Type}}

	{{end -}}
	{{if .Numeric}}
	// Sum returns the sum of all the elements in the collection.
	Sum() {{.Type}}

	{{end -}}
	{{if .Comparable}}
	// ContainsAll determines if two collections have the same size and contain the same items.
	// The order of items does not matter.
	//TODO ContainsAll(other {{.UType}}Collection) bool

{{end -}}
}
