// A simple type derived from map[int32]struct{}
// Not thread-safe.
//
// Generated from set.tpl with Type=int32
// options: Numeric=true Stringer=true Mutable=always

package simple


import (
	"bytes"
	"fmt"
)

// XInt32Set is the primary type that represents a set
type XInt32Set map[int32]struct{}

// NewXInt32Set creates and returns a reference to an empty set.
func NewXInt32Set(a ...int32) XInt32Set {
	set := make(XInt32Set)
	for _, i := range a {
		set[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set XInt32Set) ToSlice() []int32 {
	var s []int32
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set XInt32Set) Clone() XInt32Set {
	clonedSet := NewXInt32Set()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set XInt32Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set XInt32Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set XInt32Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set XInt32Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set XInt32Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set XInt32Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set XInt32Set) Add(i ...int32) XInt32Set {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set XInt32Set) doAdd(i int32) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set XInt32Set) Contains(i int32) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set XInt32Set) ContainsAll(i ...int32) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set XInt32Set) IsSubset(other XInt32Set) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set XInt32Set) IsSuperset(other XInt32Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set XInt32Set) Append(more ...int32) XInt32Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set XInt32Set) Union(other XInt32Set) XInt32Set {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set XInt32Set) Intersect(other XInt32Set) XInt32Set {
	intersection := NewXInt32Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set XInt32Set) Difference(other XInt32Set) XInt32Set {
	differencedSet := NewXInt32Set()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set XInt32Set) SymmetricDifference(other XInt32Set) XInt32Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}


// Clear clears the entire set to be the empty set.
func (set *XInt32Set) Clear() {
	*set = NewXInt32Set()
}

// Remove allows the removal of a single item from the set.
func (set XInt32Set) Remove(i int32) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set XInt32Set) Send() <-chan int32 {
	ch := make(chan int32)
	go func() {
		for v := range set {
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
func (set XInt32Set) Forall(fn func(int32) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set XInt32Set) Exists(fn func(int32) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over int32Set and executes the passed func against each element.
func (set XInt32Set) Foreach(fn func(int32)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new XInt32Set whose elements return true for func.
func (set XInt32Set) Filter(fn func(int32) bool) XInt32Set {
	result := NewXInt32Set()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new int32Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set XInt32Set) Partition(p func(int32) bool) (XInt32Set, XInt32Set) {
	matching := NewXInt32Set()
	others := NewXInt32Set()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of XInt32Set that return true for the passed predicate.
func (set XInt32Set) CountBy(predicate func(int32) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XInt32Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set XInt32Set) MinBy(less func(int32, int32) bool) int32 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m int32
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of XInt32Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set XInt32Set) MaxBy(less func(int32, int32) bool) int32 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m int32
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list XInt32Set) Min() int32 {
	return list.MinBy(func(a int32, b int32) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list XInt32Set) Max() (result int32) {
	return list.MaxBy(func(a int32, b int32) bool {
		return a < b
	})
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is numeric.

// Sum returns the sum of all the elements in the set.
func (set XInt32Set) Sum() int32 {
	sum := int32(0)
	for v, _ := range set {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set XInt32Set) Equals(other XInt32Set) bool {
	if set.Size() != other.Size() {
		return false
	}
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}


//-------------------------------------------------------------------------------------------------

func (set XInt32Set) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set XInt32Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set XInt32Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set XInt32Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set XInt32Set) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set XInt32Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
