// Generated from fast/collection.tpl with Type=*Pear
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:always
// by runtemplate v3.3.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

// P1PearSizer defines an interface for sizing methods on *Pear collections.
type P1PearSizer interface {
	// IsEmpty tests whether P1PearCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether P1PearCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// P1PearCollection defines an interface for common collection methods on *Pear.
type P1PearCollection interface {
	P1PearSizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []*Pear

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of P1PearCollection return true for the predicate p.
	Exists(p func(*Pear) bool) bool

	// Forall verifies that all elements of P1PearCollection return true for the predicate p.
	Forall(p func(*Pear) bool) bool

	// Foreach iterates over P1PearCollection and executes the function f against each element.
	Foreach(f func(*Pear))

	// Find returns the first *Pear that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(*Pear) bool) (*Pear, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *Pear

	// CountBy gives the number elements of P1PearCollection that return true for the predicate p.
	CountBy(p func(*Pear) bool) int

	// Clear the entire collection.
	Clear()

	// Add adds items to the current collection.
	Add(more ...*Pear)

	// MinBy returns an element of P1PearCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(*Pear, *Pear) bool) *Pear

	// MaxBy returns an element of P1PearCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(*Pear, *Pear) bool) *Pear
}
