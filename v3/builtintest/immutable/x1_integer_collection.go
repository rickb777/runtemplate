// Generated from immutable/collection.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:true Mutable:disabled
// by runtemplate v3.5.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"math/big"
)

// X1IntegerSizer defines an interface for sizing methods on big.Int collections.
type X1IntegerSizer interface {
	// IsEmpty tests whether X1IntegerCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X1IntegerCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// X1IntegerMkStringer defines an interface for stringer methods on big.Int collections.
type X1IntegerMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// X1IntegerCollection defines an interface for common collection methods on big.Int.
type X1IntegerCollection interface {
	X1IntegerSizer
	X1IntegerMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToList returns a shallow copy as a list.
	ToList() *X1IntegerList

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []big.Int

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of X1IntegerCollection return true for the predicate p.
	Exists(p func(big.Int) bool) bool

	// Forall verifies that all elements of X1IntegerCollection return true for the predicate p.
	Forall(p func(big.Int) bool) bool

	// Foreach iterates over X1IntegerCollection and executes the function f against each element.
	Foreach(f func(big.Int))

	// Find returns the first big.Int that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(big.Int) bool) (big.Int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan big.Int

	// CountBy gives the number elements of X1IntegerCollection that return true for the predicate p.
	CountBy(p func(big.Int) bool) int

	// MinBy returns an element of X1IntegerCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(big.Int, big.Int) bool) big.Int

	// MaxBy returns an element of X1IntegerCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(big.Int, big.Int) bool) big.Int
}
