// Generated from fast/collection.tpl with Type=int
// options: Comparable:true Numeric:<no value> Integer:true Ordered:true Stringer:true Mutable:always
// by runtemplate v3.10.2
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

// FastIntSizer defines an interface for sizing methods on int collections.
type FastIntSizer interface {
	// IsEmpty tests whether FastIntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether FastIntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the collection - an alias of Len().
	Size() int
}

// FastIntMkStringer defines an interface for stringer methods on int collections.
type FastIntMkStringer interface {
	// String implements the Stringer interface to render the collection as a comma-separated string enclosed
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

	// StringList gets a collection of strings that depicts all the elements.
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

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v int) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...int) bool

	// Clear the entire collection.
	Clear()

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

	// Fold aggregates all the values in the collection using a supplied function, starting from some initial value.
	Fold(initial int, fn func(int, int) int) int

	// Sum returns the sum of all the elements in the collection.
	Sum() int
}

// FastIntSequence defines an interface for sequence methods on int.
type FastIntSequence interface {
	FastIntCollection

	// Head gets the first element in the sequence. Head plus Tail include the whole sequence. Head is the opposite of Last.
	Head() int

	// HeadOption gets the first element in the sequence, if possible.
	HeadOption() (int, bool)

	// Last gets the last element in the sequence. Init plus Last include the whole sequence. Last is the opposite of Head.
	Last() int

	// LastOption gets the last element in the sequence, if possible.
	LastOption() (int, bool)
}
