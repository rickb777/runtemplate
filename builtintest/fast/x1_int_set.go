// An encapsulated map[int]struct{} used as a set.
// Thread-safe.
//
// Generated from fast/set.tpl with Type=int
// options: Comparable:always Numeric:true Ordered:true Stringer:true

package fast

import (

	"bytes"
	"fmt"
)

// X1IntSet is the primary type that represents a set
type X1IntSet struct {
	m map[int]struct{}
}

// NewX1IntSet creates and returns a reference to an empty set.
func NewX1IntSet(values ...int) X1IntSet {
	set := X1IntSet{
		m: make(map[int]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertX1IntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertX1IntSet(values ...interface{}) (X1IntSet, bool) {
	set := NewX1IntSet()
	good := true

	for _, i := range values {
		v, ok := i.(int)
		if !ok {
			good = false
		} else {
			set.m[v] = struct{}{}
		}
	}
	return set, good
}

// BuildX1IntSetFromChan constructs a new X1IntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1IntSetFromChan(source <-chan int) X1IntSet {
	set := NewX1IntSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set X1IntSet) ToSlice() []int {

	var s []int
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set X1IntSet) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set X1IntSet) Clone() X1IntSet {
	clonedSet := NewX1IntSet()


	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X1IntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X1IntSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set X1IntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set X1IntSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X1IntSet) Size() int {

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X1IntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set X1IntSet) Add(more ...int) {

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set X1IntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set X1IntSet) Contains(i int) bool {

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set X1IntSet) ContainsAll(i ...int) bool {

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set X1IntSet) IsSubset(other X1IntSet) bool {

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set X1IntSet) IsSuperset(other X1IntSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set X1IntSet) Union(other X1IntSet) X1IntSet {
	unionedSet := set.Clone()


	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X1IntSet) Intersect(other X1IntSet) X1IntSet {
	intersection := NewX1IntSet()


	// loop over smaller set
	if set.Size() < other.Size() {
		for v, _ := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v, _ := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set X1IntSet) Difference(other X1IntSet) X1IntSet {
	differencedSet := NewX1IntSet()


	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X1IntSet) SymmetricDifference(other X1IntSet) X1IntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *X1IntSet) Clear() {

	set.m = make(map[int]struct{})
}

// Remove removes a single item from the set.
func (set X1IntSet) Remove(i int) {

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X1IntSet) Send() <-chan int {
	ch := make(chan int)
	go func() {

		for v, _ := range set.m {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set X1IntSet) Forall(fn func(int) bool) bool {

	for v, _ := range set.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set X1IntSet) Exists(fn func(int) bool) bool {

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over intSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set X1IntSet) Foreach(fn func(int)) {

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (set X1IntSet) Find(fn func(int) bool) (int, bool) {

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}


	var empty int
	return empty, false

}

// Filter returns a new X1IntSet whose elements return true for func.
func (set X1IntSet) Filter(fn func(int) bool) X1IntSet {
	result := NewX1IntSet()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set X1IntSet) Partition(p func(int) bool) (X1IntSet, X1IntSet) {
	matching := NewX1IntSet()
	others := NewX1IntSet()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of X1IntSet that return true for the passed predicate.
func (set X1IntSet) CountBy(predicate func(int) bool) (result int) {

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set X1IntSet) Min() int {

	var m int
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set X1IntSet) Max() (result int) {

	var m int
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the set.
func (set X1IntSet) Sum() int {

	sum := int(0)
	for v, _ := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set X1IntSet) Equals(other X1IntSet) bool {

	if set.Size() != other.Size() {
		return false
	}
	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}


//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set X1IntSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (set X1IntSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// implements json.Marshaler interface {
func (set X1IntSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set X1IntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set X1IntSet) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set X1IntSet) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for v, _ := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set X1IntSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

