// An encapsulated map[int64]struct{} used as a set.
// Thread-safe.
//
// Generated from set.tpl with Type=int64
// options: Numeric=true Ordered=true Stringer=true Mutable=false

package threadsafe


import (
	"bytes"
	"fmt"
	"sync"
)

// XInt64Set is the primary type that represents a set
type XInt64Set struct {
	s *sync.RWMutex
	m map[int64]struct{}
}

// NewXInt64Set creates and returns a reference to an empty set.
func NewXInt64Set(a ...int64) *XInt64Set {
	set := &XInt64Set{
		s: &sync.RWMutex{},
		m: make(map[int64]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set *XInt64Set) ToSlice() []int64 {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []int64
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set *XInt64Set) Clone() *XInt64Set {
	clonedSet := NewXInt64Set()

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *XInt64Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *XInt64Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set *XInt64Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set *XInt64Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *XInt64Set) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *XInt64Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

func (set *XInt64Set) doAdd(i int64) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set *XInt64Set) Contains(i int64) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set *XInt64Set) ContainsAll(i ...int64) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set *XInt64Set) IsSubset(other *XInt64Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set *XInt64Set) IsSuperset(other *XInt64Set) bool {
	return other.IsSubset(set)
}

// Append returns a new set with all original items and all in `more`.
func (set *XInt64Set) Append(more ...int64) *XInt64Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set *XInt64Set) Union(other *XInt64Set) *XInt64Set {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *XInt64Set) Intersect(other *XInt64Set) *XInt64Set {
	intersection := NewXInt64Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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
func (set *XInt64Set) Difference(other *XInt64Set) *XInt64Set {
	differencedSet := NewXInt64Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *XInt64Set) SymmetricDifference(other *XInt64Set) *XInt64Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *XInt64Set) Send() <-chan int64 {
	ch := make(chan int64)
	go func() {
		set.s.RLock()
		defer set.s.RUnlock()

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
func (set *XInt64Set) Forall(fn func(int64) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *XInt64Set) Exists(fn func(int64) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over int64Set and executes the passed func against each element.
func (set *XInt64Set) Foreach(fn func(int64)) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new XInt64Set whose elements return true for func.
func (set *XInt64Set) Filter(fn func(int64) bool) *XInt64Set {
	result := NewXInt64Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new int64Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *XInt64Set) Partition(p func(int64) bool) (*XInt64Set, *XInt64Set) {
	matching := NewXInt64Set()
	others := NewXInt64Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of XInt64Set that return true for the passed predicate.
func (set *XInt64Set) CountBy(predicate func(int64) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XInt64Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *XInt64Set) MinBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m int64
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of XInt64Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *XInt64Set) MaxBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m int64
	first := true
	for v, _ := range set.m {
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
// These methods are included when int64 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *XInt64Set) Min() int64 {
	return set.MinBy(func(a int64, b int64) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set *XInt64Set) Max() (result int64) {
	return set.MaxBy(func(a int64, b int64) bool {
		return a < b
	})
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is numeric.

// Sum returns the sum of all the elements in the set.
func (set *XInt64Set) Sum() int64 {
	set.s.RLock()
	defer set.s.RUnlock()

	sum := int64(0)
	for v, _ := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *XInt64Set) Equals(other *XInt64Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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

func (set *XInt64Set) StringList() []string {
	strings := make([]string, 0)
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set *XInt64Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set *XInt64Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *XInt64Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *XInt64Set) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set *XInt64Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

