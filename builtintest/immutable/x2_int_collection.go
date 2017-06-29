// Generated from immutable/collection.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:disabled

package immutable


import (
    "math/big"
)

// X2IntSizer defines an interface for sizing methods on big.Int collections.
type X2IntSizer interface {
	// IsEmpty tests whether X2IntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X2IntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// X2IntCollection defines an interface for common collection methods on big.Int.
type X2IntCollection interface {
	X2IntSizer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []big.Int

	// Exists verifies that one or more elements of X2IntCollection return true for the passed func.
	Exists(fn func(big.Int) bool) bool

	// Forall verifies that all elements of X2IntCollection return true for the passed func.
	Forall(fn func(big.Int) bool) bool

	// Foreach iterates over X2IntCollection and executes the passed func against each element.
	Foreach(fn func(big.Int))

    // Find returns the first big.Int that returns true for some function.
    // False is returned if none match.
    Find(fn func(big.Int) bool) (big.Int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan big.Int

	// CountBy gives the number elements of X2IntCollection that return true for the passed predicate.
	CountBy(predicate func(big.Int) bool) int

// MinBy returns an element of X2IntCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(big.Int, big.Int) bool) big.Int

	// MaxBy returns an element of X2IntCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(big.Int, big.Int) bool) big.Int

}
