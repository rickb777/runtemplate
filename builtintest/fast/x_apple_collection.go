// Generated from collection.tpl with Type=Apple
// options: Comparable=<no value> Numeric=<no value> Ordered=<no value> Stringer=false

package fast

// XAppleSizer defines an interface for sizing methods on Apple collections.
type XAppleSizer interface {
	// IsEmpty tests whether XAppleCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether XAppleCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// XAppleCollection defines an interface for common collection methods on Apple.
type XAppleCollection interface {
	XAppleSizer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []Apple

	// Exists verifies that one or more elements of XAppleCollection return true for the passed func.
	Exists(fn func(Apple) bool) bool

	// Forall verifies that all elements of XAppleCollection return true for the passed func.
	Forall(fn func(Apple) bool) bool

	// Foreach iterates over XAppleCollection and executes the passed func against each element.
	Foreach(fn func(Apple))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Apple

	// CountBy gives the number elements of XAppleCollection that return true for the passed predicate.
	CountBy(predicate func(Apple) bool) int

	// MinBy returns an element of XAppleCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Apple, Apple) bool) Apple

	// MaxBy returns an element of XAppleCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Apple, Apple) bool) Apple

	}
