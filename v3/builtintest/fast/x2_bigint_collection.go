// Generated from fast/collection.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast


import (
	"math/big"
)

// X2BigIntSizer defines an interface for sizing methods on big.Int collections.
type X2BigIntSizer interface {
	// IsEmpty tests whether X2BigIntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X2BigIntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// X2BigIntCollection defines an interface for common collection methods on big.Int.
type X2BigIntCollection interface {
	X2BigIntSizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []big.Int

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of X2BigIntCollection return true for the predicate p.
	Exists(p func(big.Int) bool) bool

	// Forall verifies that all elements of X2BigIntCollection return true for the predicate p.
	Forall(p func(big.Int) bool) bool

	// Foreach iterates over X2BigIntCollection and executes the function f against each element.
	Foreach(f func(big.Int))

	// Find returns the first big.Int that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(big.Int) bool) (big.Int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan big.Int

	// CountBy gives the number elements of X2BigIntCollection that return true for the predicate p.
	CountBy(p func(big.Int) bool) int

    // Clear the entire collection.
    Clear()

	// Add adds items to the current collection.
	Add(more ...big.Int)

	// MinBy returns an element of X2BigIntCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(big.Int, big.Int) bool) big.Int

	// MaxBy returns an element of X2BigIntCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(big.Int, big.Int) bool) big.Int
}
