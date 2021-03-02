// Generated from immutable/collection.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Integer:<no value> Ordered:<no value> Stringer:false Mutable:disabled
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

// ImmutableAppleSizer defines an interface for sizing methods on Apple collections.
type ImmutableAppleSizer interface {
	// IsEmpty tests whether ImmutableAppleCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether ImmutableAppleCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// ImmutableAppleCollection defines an interface for common collection methods on Apple.
type ImmutableAppleCollection interface {
	ImmutableAppleSizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []Apple

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of ImmutableAppleCollection return true for the predicate p.
	Exists(p func(Apple) bool) bool

	// Forall verifies that all elements of ImmutableAppleCollection return true for the predicate p.
	Forall(p func(Apple) bool) bool

	// Foreach iterates over ImmutableAppleCollection and executes the function f against each element.
	Foreach(f func(Apple))

	// Find returns the first Apple that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(Apple) bool) (Apple, bool)

	// MapToString returns a new []string by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToString(f func(Apple) string) []string

	// FlatMapString returns a new []string by transforming every element with function f
	// that returns zero or more items in a slice. The resulting slice may have a different size to the
	// collection. The collection is not modified.
	FlatMapToString(f func(Apple) []string) []string

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Apple

	// CountBy gives the number elements of ImmutableAppleCollection that return true for the predicate p.
	CountBy(p func(Apple) bool) int

	// MinBy returns an element of ImmutableAppleCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Apple, Apple) bool) Apple

	// MaxBy returns an element of ImmutableAppleCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Apple, Apple) bool) Apple

	// Fold aggregates all the values in the collection using a supplied function, starting from some initial value.
	Fold(initial Apple, fn func(Apple, Apple) Apple) Apple
}

// ImmutableAppleSequence defines an interface for sequence methods on Apple.
type ImmutableAppleSequence interface {
	ImmutableAppleCollection

	// Head gets the first element in the sequence. Head plus Tail include the whole sequence. Head is the opposite of Last.
	Head() Apple

	// HeadOption gets the first element in the sequence, if possible.
	HeadOption() (Apple, bool)

	// Last gets the last element in the sequence. Init plus Last include the whole sequence. Last is the opposite of Head.
	Last() Apple

	// LastOption gets the last element in the sequence, if possible.
	LastOption() (Apple, bool)
}
