// An encapsulated map[int]struct{} used as a set.
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=int
// options: Comparable:always Numeric:true Ordered:true Stringer:true

package examples

import (

	"bytes"
	"fmt"
	"sync"
)

// SyncIntSet is the primary type that represents a set
type SyncIntSet struct {
	s *sync.RWMutex
	m map[int]struct{}
}

// NewSyncIntSet creates and returns a reference to an empty set.
func NewSyncIntSet(values ...int) SyncIntSet {
	set := SyncIntSet{
		s: &sync.RWMutex{},
		m: make(map[int]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertSyncIntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertSyncIntSet(values ...interface{}) (SyncIntSet, bool) {
	set := NewSyncIntSet()

	for _, i := range values {
		switch i.(type) {
		case int:
			set.m[int(i.(int))] = struct{}{}
		case int8:
			set.m[int(i.(int8))] = struct{}{}
		case int16:
			set.m[int(i.(int16))] = struct{}{}
		case int32:
			set.m[int(i.(int32))] = struct{}{}
		case int64:
			set.m[int(i.(int64))] = struct{}{}
		case uint:
			set.m[int(i.(uint))] = struct{}{}
		case uint8:
			set.m[int(i.(uint8))] = struct{}{}
		case uint16:
			set.m[int(i.(uint16))] = struct{}{}
		case uint32:
			set.m[int(i.(uint32))] = struct{}{}
		case uint64:
			set.m[int(i.(uint64))] = struct{}{}
		case float32:
			set.m[int(i.(float32))] = struct{}{}
		case float64:
			set.m[int(i.(float64))] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildSyncIntSetFromChan constructs a new SyncIntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildSyncIntSetFromChan(source <-chan int) SyncIntSet {
	set := NewSyncIntSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set SyncIntSet) ToSlice() []int {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []int
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set SyncIntSet) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []interface{}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set SyncIntSet) Clone() SyncIntSet {
	clonedSet := NewSyncIntSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set SyncIntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set SyncIntSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set SyncIntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set SyncIntSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set SyncIntSet) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set SyncIntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set SyncIntSet) Add(more ...int) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set SyncIntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set SyncIntSet) Contains(i int) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set SyncIntSet) ContainsAll(i ...int) bool {
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
func (set SyncIntSet) IsSubset(other SyncIntSet) bool {
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
func (set SyncIntSet) IsSuperset(other SyncIntSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set SyncIntSet) Union(other SyncIntSet) SyncIntSet {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set SyncIntSet) Intersect(other SyncIntSet) SyncIntSet {
	intersection := NewSyncIntSet()

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
func (set SyncIntSet) Difference(other SyncIntSet) SyncIntSet {
	differencedSet := NewSyncIntSet()

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
func (set SyncIntSet) SymmetricDifference(other SyncIntSet) SyncIntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *SyncIntSet) Clear() {
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[int]struct{})
}

// Remove removes a single item from the set.
func (set SyncIntSet) Remove(i int) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set SyncIntSet) Send() <-chan int {
	ch := make(chan int)
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
func (set SyncIntSet) Forall(fn func(int) bool) bool {
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
func (set SyncIntSet) Exists(fn func(int) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over intSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set SyncIntSet) Foreach(fn func(int)) {
	set.s.Lock()
	defer set.s.Unlock()

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (set SyncIntSet) Find(fn func(int) bool) (int, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}


	var empty int
	return empty, false

}

// Filter returns a new SyncIntSet whose elements return true for func.
//
// The original set is not modified
func (set SyncIntSet) Filter(fn func(int) bool) SyncIntSet {
	result := NewSyncIntSet()
	set.s.RLock()
	defer set.s.RUnlock()

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
//
// The original set is not modified
func (set SyncIntSet) Partition(p func(int) bool) (SyncIntSet, SyncIntSet) {
	matching := NewSyncIntSet()
	others := NewSyncIntSet()
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

// Map returns a new SyncIntSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SyncIntSet) Map(fn func(int) int) SyncIntSet {
	result := NewSyncIntSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
        result.m[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new SyncIntSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SyncIntSet) FlatMap(fn func(int) []int) SyncIntSet {
	result := NewSyncIntSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
	    for _, x := range fn(v) {
            result.m[x] = struct{}{}
	    }
	}

	return result
}

// CountBy gives the number elements of SyncIntSet that return true for the passed predicate.
func (set SyncIntSet) CountBy(predicate func(int) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set SyncIntSet) Min() int {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set SyncIntSet) Max() (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set SyncIntSet) Sum() int {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set SyncIntSet) Equals(other SyncIntSet) bool {
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

// StringList gets a list of strings that depicts all the elements.
func (set SyncIntSet) StringList() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (set SyncIntSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// implements json.Marshaler interface {
func (set SyncIntSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set SyncIntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set SyncIntSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set SyncIntSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set SyncIntSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
