// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable={{.Comparable}} Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}}

package {{.Package}}

// {{.UType}}Collection defines an interface for common collection methods on {{.PType}}.
type {{.UType}}Collection interface {
    // IsEmpty tests whether {{.UType}}Collection is empty.
    IsEmpty() bool

    // NonEmpty tests whether {{.UType}}Collection is empty.
    NonEmpty() bool

    // Size returns the number of items in the list - an alias of Len().
    Size() int

    // IsSequence returns true for lists.
    IsSequence() bool

    // IsSet returns false for lists.
    IsSet() bool

    // Exists verifies that one or more elements of {{.UType}}Collection return true for the passed func.
    Exists(fn func({{.PType}}) bool) bool

    // Forall verifies that all elements of {{.UType}}Collection return true for the passed func.
    Forall(fn func({{.PType}}) bool) bool

    // Foreach iterates over {{.UType}}Collection and executes the passed func against each element.
    Foreach(fn func({{.PType}}))

    // Send returns a channel that will send all the elements in order.
    Send() <-chan {{.PType}}

    // CountBy gives the number elements of {{.UType}}Collection that return true for the passed predicate.
    CountBy(predicate func({{.PType}}) bool) int

    // MinBy returns an element of {{.UType}}Collection containing the minimum value, when compared to other elements
    // using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
    // element is returned. Panics if there are no elements.
    MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}

    // MaxBy returns an element of {{.UType}}Collection containing the maximum value, when compared to other elements
    // using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
    // element is returned. Panics if there are no elements.
    MaxBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}
    {{if .Ordered}}

    // Min returns the minimum value of all the items in the collection. Panics if there are no elements.
    Min() {{.PType}}

    // Max returns the minimum value of all the items in the collection. Panics if there are no elements.
    Max() {{.PType}}
    {{end}}
    {{if .Numeric}}

    // Sum returns the sum of all the elements in the collection.
    Sum() {{.PType}}
    {{end}}
    {{if .Stringer}}

    // String implements the Stringer interface to render the list as a comma-separated string enclosed
    // in square brackets.
    String() string

    // MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
    MkString(sep string) string

    // MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
    MkString3(pfx, mid, sfx string) string
    {{end}}
}
