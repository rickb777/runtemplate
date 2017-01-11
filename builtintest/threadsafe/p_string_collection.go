// Generated from collection.tpl with Type=*string
// options: Comparable=true Numeric=<no value> Ordered=<no value> Stringer=true

package threadsafe

// StringCollection defines an interface for common collection methods on *string.
type PStringCollection interface {
	// IsEmpty tests whether PStringCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether PStringCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int

	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// Exists verifies that one or more elements of PStringCollection return true for the passed func.
	Exists(fn func(*string) bool) bool

	// Forall verifies that all elements of PStringCollection return true for the passed func.
	Forall(fn func(*string) bool) bool

	// Foreach iterates over PStringCollection and executes the passed func against each element.
	Foreach(fn func(*string))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *string

	// CountBy gives the number elements of PStringCollection that return true for the passed predicate.
	CountBy(predicate func(*string) bool) int

	// MinBy returns an element of PStringCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(*string, *string) bool) *string

	// MaxBy returns an element of PStringCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(*string, *string) bool) *string

	
	// ContainsAll determines if two collections have the same size and contain the same items.
	// The order of items does not matter.
	//TODO ContainsAll(other StringCollection) bool

	
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

}
