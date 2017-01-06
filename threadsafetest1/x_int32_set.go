// Generated from set.tpl with Type=int32
// options: Numeric=true Ordered=true Stringer=true Mutable=always-true

package threadsafetest1


import (
	"bytes"
	"fmt"
	"sync"
)

// Int32Set is the primary type that represents a set
type Int32Set struct {
	s *sync.RWMutex
	m map[int32]struct{}
}

// NewInt32Set creates and returns a reference to an empty set.
func NewInt32Set(a ...int32) Int32Set {
	set := Int32Set{
		s: &sync.RWMutex{},
		m: make(map[int32]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set Int32Set) ToSlice() []int32 {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []int32
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set Int32Set) Clone() Int32Set {
	clonedSet := NewInt32Set()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set Int32Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set Int32Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set Int32Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set Int32Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set Int32Set) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set Int32Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds an item to the current set if it doesn't already exist in the set.
func (set Int32Set) Add(i int32) bool {
	set.s.Lock()
	defer set.s.Unlock()

	_, found := set.m[i]
	set.m[i] = struct{}{}
	return !found //False if it existed already
}

func (set Int32Set) doAdd(i int32) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set Int32Set) Contains(i int32) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set Int32Set) ContainsAll(i ...int32) bool {
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
func (set Int32Set) IsSubset(other Int32Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set Int32Set) IsSuperset(other Int32Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set Int32Set) Append(more ...int32) Int32Set {
	set.s.Lock()
	defer set.s.Unlock()

	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set Int32Set) Union(other Int32Set) Int32Set {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.m[v] = struct{}{}
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set Int32Set) Intersect(other Int32Set) Int32Set {
	intersection := NewInt32Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(v) {
				intersection.Add(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set Int32Set) Difference(other Int32Set) Int32Set {
	differencedSet := NewInt32Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.Add(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set Int32Set) SymmetricDifference(other Int32Set) Int32Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *Int32Set) Clear() {
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[int32]struct{})
}

// Remove allows the removal of a single item from the set.
func (set Int32Set) Remove(i int32) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set Int32Set) Send() <-chan int32 {
	ch := make(chan int32)
	go func() {
		set.s.RLock()
		defer set.s.RUnlock()

		for v := range set.m {
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
func (set Int32Set) Forall(fn func(int32) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set Int32Set) Exists(fn func(int32) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over int32Set and executes the passed func against each element.
func (set Int32Set) Foreach(fn func(int32)) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new Int32Set whose elements return true for func.
func (set Int32Set) Filter(fn func(int32) bool) Int32Set {
	result := NewInt32Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if fn(v) {
			result.m[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new int32Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set Int32Set) Partition(p func(int32) bool) (Int32Set, Int32Set) {
	matching := NewInt32Set()
	others := NewInt32Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			matching.m[v] = struct{}{}
		} else {
			others.m[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of Int32Set that return true for the passed predicate.
func (set Int32Set) CountBy(predicate func(int32) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Int32Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set Int32Set) MinBy(less func(int32, int32) bool) int32 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m int32
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of Int32Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set Int32Set) MaxBy(less func(int32, int32) bool) int32 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m int32
	first := true
	for v := range set.m {
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
// These methods are included when int32 is numeric.

// Sum returns the sum of all the elements in the set.
func (set Int32Set) Sum() int32 {
	set.s.RLock()
	defer set.s.RUnlock()

	sum := int32(0)
	for v, _ := range set.m {
		sum = sum + v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set Int32Set) Equals(other Int32Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	if set.Size() != other.Size() {
		return false
	}
	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int32Set) Min() int32 {
	return list.MinBy(func(a int32, b int32) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int32Set) Max() (result int32) {
	return list.MaxBy(func(a int32, b int32) bool {
		return a < b
	})
}



//-------------------------------------------------------------------------------------------------

func (set Int32Set) StringList() []string {
	strings := make([]string, 0)
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set Int32Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set Int32Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set Int32Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set Int32Set) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set Int32Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

