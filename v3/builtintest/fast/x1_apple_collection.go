// Generated from fast/collection.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:false Mutable:always
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package fast

// X1AppleSizer defines an interface for sizing methods on Apple collections.
type X1AppleSizer interface {
	// IsEmpty tests whether X1AppleCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X1AppleCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// X1AppleCollection defines an interface for common collection methods on Apple.
type X1AppleCollection interface {
	X1AppleSizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() *X1AppleSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []Apple

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of X1AppleCollection return true for the predicate p.
	Exists(p func(Apple) bool) bool

	// Forall verifies that all elements of X1AppleCollection return true for the predicate p.
	Forall(p func(Apple) bool) bool

	// Foreach iterates over X1AppleCollection and executes the function f against each element.
	Foreach(f func(Apple))

	// Find returns the first Apple that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(Apple) bool) (Apple, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Apple

	// CountBy gives the number elements of X1AppleCollection that return true for the predicate p.
	CountBy(p func(Apple) bool) int

	// Clear the entire collection.
	Clear()

	// Add adds items to the current collection.
	Add(more ...Apple)

	// MinBy returns an element of X1AppleCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Apple, Apple) bool) Apple

	// MaxBy returns an element of X1AppleCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Apple, Apple) bool) Apple
}
