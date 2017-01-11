// Generated from collection.tpl with Type=int32
// options: Comparable=true Numeric=true Ordered=true Stringer=true

package fast

// XInt32Sizer defines an interface for sizing methods on int32 collections.
type XInt32Sizer interface {
	// IsEmpty tests whether XInt32Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether XInt32Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}


// XInt32MkStringer defines an interface for stringer methods on int32 collections.
type XInt32MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

}

// XInt32Collection defines an interface for common collection methods on int32.
type XInt32Collection interface {
	XInt32Sizer

	XInt32MkStringer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []int32

	// Exists verifies that one or more elements of XInt32Collection return true for the passed func.
	Exists(fn func(int32) bool) bool

	// Forall verifies that all elements of XInt32Collection return true for the passed func.
	Forall(fn func(int32) bool) bool

	// Foreach iterates over XInt32Collection and executes the passed func against each element.
	Foreach(fn func(int32))

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan int32

	// CountBy gives the number elements of XInt32Collection that return true for the passed predicate.
	CountBy(predicate func(int32) bool) int

	// MinBy returns an element of XInt32Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(int32, int32) bool) int32

	// MaxBy returns an element of XInt32Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(int32, int32) bool) int32

	
	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() int32

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() int32

	
	// Sum returns the sum of all the elements in the collection.
	Sum() int32

	
	// ContainsAll determines if two collections have the same size and contain the same items.
	// The order of items does not matter.
	//TODO ContainsAll(other Int32Collection) bool

}
