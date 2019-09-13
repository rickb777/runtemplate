// Generated from threadsafe/collection.tpl with Type=*string
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:true Mutable:always
// by runtemplate v3.5.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

// P1StringSizer defines an interface for sizing methods on *string collections.
type P1StringSizer interface {
	// IsEmpty tests whether P1StringCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether P1StringCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// P1StringMkStringer defines an interface for stringer methods on *string collections.
type P1StringMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// implements json.Unmarshaler interface {
	UnmarshalJSON(b []byte) error

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// P1StringCollection defines an interface for common collection methods on *string.
type P1StringCollection interface {
	P1StringSizer
	P1StringMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToList returns a shallow copy as a list.
	ToList() *P1StringList

	// ToSet returns a shallow copy as a set.
	ToSet() *P1StringSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []*string

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of P1StringCollection return true for the predicate p.
	Exists(p func(*string) bool) bool

	// Forall verifies that all elements of P1StringCollection return true for the predicate p.
	Forall(p func(*string) bool) bool

	// Foreach iterates over P1StringCollection and executes the function f against each element.
	Foreach(f func(*string))

	// Find returns the first *string that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(*string) bool) (*string, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *string

	// CountBy gives the number elements of P1StringCollection that return true for the predicate p.
	CountBy(p func(*string) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v *string) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...*string) bool

	// Clear the entire collection.
	Clear()

	// Add adds items to the current collection.
	Add(more ...*string)

	// MinBy returns an element of P1StringCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(*string, *string) bool) *string

	// MaxBy returns an element of P1StringCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(*string, *string) bool) *string
}
