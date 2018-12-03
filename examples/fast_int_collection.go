// Generated from fast/collection.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always
// by runtemplate v2.4.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

// FastIntSizer defines an interface for sizing methods on int collections.
type FastIntSizer interface {
	// IsEmpty tests whether FastIntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether FastIntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// FastIntMkStringer defines an interface for stringer methods on int collections.
type FastIntMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// implements json.Unmarshaler interface {
	UnmarshalJSON(b []byte) error

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// FastIntCollection defines an interface for common collection methods on int.
type FastIntCollection interface {
	FastIntSizer
	FastIntMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []int

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of FastIntCollection return true for the predicate p.
	Exists(p func(int) bool) bool

	// Forall verifies that all elements of FastIntCollection return true for the predicate p.
	Forall(p func(int) bool) bool

	// Foreach iterates over FastIntCollection and executes the function f against each element.
	Foreach(f func(int))

	// Find returns the first int that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(int) bool) (int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan int

	// CountBy gives the number elements of FastIntCollection that return true for the predicate p.
	CountBy(p func(int) bool) int

	// Contains determines if a given item is already in the collection.
	Contains(v int) bool

	// ContainsAll determines if the given items are all in the collection.
	ContainsAll(v ...int) bool

	// Add adds items to the current collection.
	Add(more ...int)

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() int

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() int

	// MinBy returns an element of FastIntCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(int, int) bool) int

	// MaxBy returns an element of FastIntCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(int, int) bool) int

	// Sum returns the sum of all the elements in the collection.
	Sum() int
}
