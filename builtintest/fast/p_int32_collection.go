// Generated from collection.tpl with Type=*int32
// options: Comparable=true Numeric=true Ordered=true Stringer=true

package fast

// PInt32Sizer defines an interface for sizing methods on *int32 collections.
type PInt32Sizer interface {
	// IsEmpty tests whether PInt32Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether PInt32Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}


// PInt32MkStringer defines an interface for stringer methods on *int32 collections.
type PInt32MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

    // implements json.Marshaler interface {
    MarshalJSON() ([]byte, error)

    // StringList gets a list of strings that depicts all the elements.
    StringList() []string
}

// PInt32Collection defines an interface for common collection methods on *int32.
type PInt32Collection interface {
	PInt32Sizer

	PInt32MkStringer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []*int32

	// Exists verifies that one or more elements of PInt32Collection return true for the passed func.
	Exists(fn func(*int32) bool) bool

	// Forall verifies that all elements of PInt32Collection return true for the passed func.
	Forall(fn func(*int32) bool) bool

	// Foreach iterates over PInt32Collection and executes the passed func against each element.
	Foreach(fn func(*int32))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *int32

    // CountBy gives the number elements of PInt32Collection that return true for the passed predicate.
	CountBy(predicate func(*int32) bool) int


    // Contains determines if a given item is already in the collection.
    Contains(v int32) bool

    // ContainsAll determines if the given items are all in the collection.
    ContainsAll(v ...int32) bool


	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() int32

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() int32


	// Sum returns the sum of all the elements in the collection.
	Sum() int32
}
