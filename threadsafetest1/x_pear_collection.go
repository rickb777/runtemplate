// Generated from collection.tpl with Type=Pear
// options: Comparable=<no value> Numeric=<no value> Ordered=<no value> Stringer=<no value>

package threadsafetest1

// PearCollection defines an interface for common collection methods on Pear.
type PearCollection interface {
	// IsEmpty tests whether PearCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether PearCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int

	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// Exists verifies that one or more elements of PearCollection return true for the passed func.
	Exists(fn func(Pear) bool) bool

	// Forall verifies that all elements of PearCollection return true for the passed func.
	Forall(fn func(Pear) bool) bool

	// Foreach iterates over PearCollection and executes the passed func against each element.
	Foreach(fn func(Pear))

	// Send returns a channel that will send all the elements in order.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan Pear

	// CountBy gives the number elements of PearCollection that return true for the passed predicate.
	CountBy(predicate func(Pear) bool) int

	// MinBy returns an element of PearCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(Pear, Pear) bool) Pear

	// MaxBy returns an element of PearCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(Pear, Pear) bool) Pear

	
	
    
	
}
