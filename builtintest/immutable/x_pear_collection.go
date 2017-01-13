// Generated from collection.tpl with Type=Pear
// options: Comparable=<no value> Numeric=<no value> Ordered=<no value> Stringer=<no value> Mutable=disabled

package immutable

// XPearSizer defines an interface for sizing methods on Pear collections.
type XPearSizer interface {
	// IsEmpty tests whether XPearCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether XPearCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// XPearCollection defines an interface for common collection methods on Pear.
type XPearCollection interface {
	XPearSizer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []Pear

	// Exists verifies that one or more elements of XPearCollection return true for the passed func.
	Exists(fn func(Pear) bool) bool

	// Forall verifies that all elements of XPearCollection return true for the passed func.
	Forall(fn func(Pear) bool) bool

	// Foreach iterates over XPearCollection and executes the passed func against each element.
	Foreach(fn func(Pear))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Pear

	// CountBy gives the number elements of XPearCollection that return true for the passed predicate.
	CountBy(predicate func(Pear) bool) int

// MinBy returns an element of XPearCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Pear, Pear) bool) Pear

	// MaxBy returns an element of XPearCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Pear, Pear) bool) Pear

}
