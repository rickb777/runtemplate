// Generated from collection.tpl with Type=string
// options: Comparable=true Numeric=<no value> Ordered=<no value> Stringer=true Mutable=disabled

package immutable

// XStringSizer defines an interface for sizing methods on string collections.
type XStringSizer interface {
	// IsEmpty tests whether XStringCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether XStringCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}


// XStringMkStringer defines an interface for stringer methods on string collections.
type XStringMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// XStringCollection defines an interface for common collection methods on string.
type XStringCollection interface {
	XStringSizer

	XStringMkStringer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []string

	// Exists verifies that one or more elements of XStringCollection return true for the passed func.
	Exists(fn func(string) bool) bool

	// Forall verifies that all elements of XStringCollection return true for the passed func.
	Forall(fn func(string) bool) bool

	// Foreach iterates over XStringCollection and executes the passed func against each element.
	Foreach(fn func(string))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan string

	// CountBy gives the number elements of XStringCollection that return true for the passed predicate.
	CountBy(predicate func(string) bool) int


	// Contains determines if a given item is already in the collection.
	Contains(v string) bool

	// ContainsAll determines if the given items are all in the collection.
	ContainsAll(v ...string) bool

// MinBy returns an element of XStringCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(string, string) bool) string

	// MaxBy returns an element of XStringCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(string, string) bool) string

}
