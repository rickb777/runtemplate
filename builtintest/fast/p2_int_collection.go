// Generated from fast/collection.tpl with Type=*big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:always

package fast


import (
	"math/big"
)

// P2IntSizer defines an interface for sizing methods on *big.Int collections.
type P2IntSizer interface {
	// IsEmpty tests whether P2IntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether P2IntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// P2IntCollection defines an interface for common collection methods on *big.Int.
type P2IntCollection interface {
	P2IntSizer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []*big.Int

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of P2IntCollection return true for the passed func.
	Exists(fn func(*big.Int) bool) bool

	// Forall verifies that all elements of P2IntCollection return true for the passed func.
	Forall(fn func(*big.Int) bool) bool

	// Foreach iterates over P2IntCollection and executes the passed func against each element.
	Foreach(fn func(*big.Int))

	// Find returns the first big.Int that returns true for some function.
	// False is returned if none match.
	Find(fn func(*big.Int) bool) (*big.Int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *big.Int

	// CountBy gives the number elements of P2IntCollection that return true for the passed predicate.
	CountBy(predicate func(*big.Int) bool) int

// Add adds items to the current collection.
	Add(more ...big.Int)

// MinBy returns an element of P2IntCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(*big.Int, *big.Int) bool) *big.Int

	// MaxBy returns an element of P2IntCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(*big.Int, *big.Int) bool) *big.Int

}
