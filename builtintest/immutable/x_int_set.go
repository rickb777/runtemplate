// An encapsulated map[int]struct{} used as a set.
// Thread-safe.
//
// Generated from immutable/set.tpl with Type=int
// options: Comparable=always Numeric=true Ordered=true Stringer=true Mutable=disabled

package immutable


import (
	"bytes"
	"fmt"
)

// XIntSet is the primary type that represents a set
type XIntSet struct {
	m map[int]struct{}
}

// NewXIntSet creates and returns a reference to an empty set.
func NewXIntSet(a ...int) XIntSet {
	set := XIntSet{
		m: make(map[int]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// BuildXIntSetFromChan constructs a new XIntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildXIntSetFromChan(source <-chan int) XIntSet {
	result := NewXIntSet()
	for v := range source {
		result.m[v] = struct{}{}
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (set XIntSet) ToSlice() []int {
	var s []int
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set XIntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set XIntSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set XIntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set XIntSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set XIntSet) Size() int {
	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set XIntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set XIntSet) Add(more ...int) XIntSet {
	newSet := NewXIntSet()

	for v, _ := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set XIntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set XIntSet) Contains(i int) bool {
	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set XIntSet) ContainsAll(i ...int) bool {

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set XIntSet) IsSubset(other XIntSet) bool {

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set XIntSet) IsSuperset(other XIntSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set XIntSet) Union(other XIntSet) XIntSet {
	unionedSet := NewXIntSet()

	for v, _ := range set.m {
		unionedSet.doAdd(v)
	}

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set XIntSet) Intersect(other XIntSet) XIntSet {
	intersection := NewXIntSet()

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
func (set XIntSet) Difference(other XIntSet) XIntSet {
	differencedSet := NewXIntSet()

	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set XIntSet) SymmetricDifference(other XIntSet) XIntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set XIntSet) Remove(i int) XIntSet {
	clonedSet := NewXIntSet()

	for v, _ := range set.m {
		if i != v {
			clonedSet.doAdd(v)
		}
	}

	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set XIntSet) Send() <-chan int {
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
func (set XIntSet) Forall(fn func(int) bool) bool {

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
func (set XIntSet) Exists(fn func(int) bool) bool {

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over intSet and executes the passed func against each element.
func (set XIntSet) Foreach(fn func(int)) {

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new XIntSet whose elements return true for func.
func (set XIntSet) Filter(fn func(int) bool) XIntSet {
	result := NewXIntSet()

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
func (set XIntSet) Partition(p func(int) bool) (XIntSet, XIntSet) {
	matching := NewXIntSet()
	others := NewXIntSet()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of XIntSet that return true for the passed predicate.
func (set XIntSet) CountBy(predicate func(int) bool) (result int) {

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
func (set XIntSet) Min() int {

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
func (set XIntSet) Max() (result int) {

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
func (set XIntSet) Sum() int {

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
func (set XIntSet) Equals(other XIntSet) bool {

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
func (set XIntSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (set XIntSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// implements json.Marshaler interface {
func (set XIntSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set XIntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set XIntSet) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set XIntSet) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
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
func (set XIntSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

