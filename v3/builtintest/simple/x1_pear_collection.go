// Generated from simple/collection.tpl with Type=Pear
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

// X1PearSizer defines an interface for sizing methods on Pear collections.
type X1PearSizer interface {
	// IsEmpty tests whether X1PearCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X1PearCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// X1PearCollection defines an interface for common collection methods on Pear.
type X1PearCollection interface {
	X1PearSizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []Pear

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of X1PearCollection return true for the predicate p.
	Exists(p func(Pear) bool) bool

	// Forall verifies that all elements of X1PearCollection return true for the predicate p.
	Forall(p func(Pear) bool) bool

	// Foreach iterates over X1PearCollection and executes the function f against each element.
	Foreach(f func(Pear))

	// Find returns the first Pear that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(Pear) bool) (Pear, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Pear

	// CountBy gives the number elements of X1PearCollection that return true for the predicate p.
	CountBy(p func(Pear) bool) int

	// Add adds items to the current collection.
	//Add(more ...Pear)

	// MinBy returns an element of X1PearCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Pear, Pear) bool) Pear

	// MaxBy returns an element of X1PearCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Pear, Pear) bool) Pear
}
