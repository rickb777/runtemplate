// Generated from fast/collection.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

// X1IntSizer defines an interface for sizing methods on int collections.
type X1IntSizer interface {
	// IsEmpty tests whether X1IntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X1IntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the collection - an alias of Len().
	Size() int
}

// X1IntMkStringer defines an interface for stringer methods on int collections.
type X1IntMkStringer interface {
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

// X1IntCollection defines an interface for common collection methods on int.
type X1IntCollection interface {
	X1IntSizer
	X1IntMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToList returns a shallow copy as a list.
	ToList() *X1IntList

	// ToSet returns a shallow copy as a set.
	ToSet() *X1IntSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []int

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of X1IntCollection return true for the predicate p.
	Exists(p func(int) bool) bool

	// Forall verifies that all elements of X1IntCollection return true for the predicate p.
	Forall(p func(int) bool) bool

	// Foreach iterates over X1IntCollection and executes the function f against each element.
	Foreach(f func(int))

	// Find returns the first int that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(int) bool) (int, bool)

	// MapToString returns a new []string by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToString(f func(int) string) []string

	// MapToInt64 returns a new []int64 by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToInt64(f func(int) int64) []int64

	// FlatMapString returns a new []string by transforming every element with function f
	// that returns zero or more items in a slice. The resulting slice may have a different size to the
	// collection. The collection is not modified.
	FlatMapToString(f func(int) []string) []string

	// FlatMapInt64 returns a new []int64 by transforming every element with function f
	// that returns zero or more items in a slice. The resulting slice may have a different size to the
	// collection. The collection is not modified.
	FlatMapToInt64(f func(int) []int64) []int64

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan int

	// CountBy gives the number elements of X1IntCollection that return true for the predicate p.
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

	// MinBy returns an element of X1IntCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(int, int) bool) int

	// MaxBy returns an element of X1IntCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(int, int) bool) int

	// Fold aggregates all the values in the collection using a supplied function, starting from some initial value.
	Fold(initial int, fn func(int, int) int) int

	// Sum returns the sum of all the elements in the collection.
	Sum() int
}

// X1IntSequence defines an interface for sequence methods on int.
type X1IntSequence interface {
	X1IntCollection

	// Head gets the first element in the sequence. Head plus Tail include the whole sequence. Head is the opposite of Last.
	Head() int

	// HeadOption gets the first element in the sequence, if possible.
	HeadOption() (int, bool)

	// Last gets the last element in the sequence. Init plus Last include the whole sequence. Last is the opposite of Head.
	Last() int

	// LastOption gets the last element in the sequence, if possible.
	LastOption() (int, bool)
}
