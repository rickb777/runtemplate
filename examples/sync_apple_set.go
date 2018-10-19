// An encapsulated map[Apple]struct{} used as a set.
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=Apple
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:false

package examples

import (

	"sync"
)

// SyncAppleSet is the primary type that represents a set
type SyncAppleSet struct {
	s *sync.RWMutex
	m map[Apple]struct{}
}

// NewSyncAppleSet creates and returns a reference to an empty set.
func NewSyncAppleSet(values ...Apple) SyncAppleSet {
	set := SyncAppleSet{
		s: &sync.RWMutex{},
		m: make(map[Apple]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertSyncAppleSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertSyncAppleSet(values ...interface{}) (SyncAppleSet, bool) {
	set := NewSyncAppleSet()

	for _, i := range values {
		v, ok := i.(Apple)
		if ok {
			set.m[v] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildSyncAppleSetFromChan constructs a new SyncAppleSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildSyncAppleSetFromChan(source <-chan Apple) SyncAppleSet {
	set := NewSyncAppleSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set SyncAppleSet) ToSlice() []Apple {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []Apple
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set SyncAppleSet) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []interface{}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set SyncAppleSet) Clone() SyncAppleSet {
	clonedSet := NewSyncAppleSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set SyncAppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set SyncAppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set SyncAppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set SyncAppleSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set SyncAppleSet) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set SyncAppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set SyncAppleSet) Add(more ...Apple) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set SyncAppleSet) doAdd(i Apple) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set SyncAppleSet) Contains(i Apple) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set SyncAppleSet) ContainsAll(i ...Apple) bool {
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
func (set SyncAppleSet) IsSubset(other SyncAppleSet) bool {
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
func (set SyncAppleSet) IsSuperset(other SyncAppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set SyncAppleSet) Union(other SyncAppleSet) SyncAppleSet {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set SyncAppleSet) Intersect(other SyncAppleSet) SyncAppleSet {
	intersection := NewSyncAppleSet()

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
func (set SyncAppleSet) Difference(other SyncAppleSet) SyncAppleSet {
	differencedSet := NewSyncAppleSet()

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
func (set SyncAppleSet) SymmetricDifference(other SyncAppleSet) SyncAppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *SyncAppleSet) Clear() {
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[Apple]struct{})
}

// Remove removes a single item from the set.
func (set SyncAppleSet) Remove(i Apple) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set SyncAppleSet) Send() <-chan Apple {
	ch := make(chan Apple)
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
func (set SyncAppleSet) Forall(fn func(Apple) bool) bool {
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
func (set SyncAppleSet) Exists(fn func(Apple) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set SyncAppleSet) Foreach(fn func(Apple)) {
	set.s.Lock()
	defer set.s.Unlock()

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (set SyncAppleSet) Find(fn func(Apple) bool) (Apple, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}


	var empty Apple
	return empty, false

}

// Filter returns a new SyncAppleSet whose elements return true for func.
//
// The original set is not modified
func (set SyncAppleSet) Filter(fn func(Apple) bool) SyncAppleSet {
	result := NewSyncAppleSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set SyncAppleSet) Partition(p func(Apple) bool) (SyncAppleSet, SyncAppleSet) {
	matching := NewSyncAppleSet()
	others := NewSyncAppleSet()
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

// Map returns a new SyncAppleSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SyncAppleSet) Map(fn func(Apple) Apple) SyncAppleSet {
	result := NewSyncAppleSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
        result.m[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new SyncAppleSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SyncAppleSet) FlatMap(fn func(Apple) []Apple) SyncAppleSet {
	result := NewSyncAppleSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
	    for _, x := range fn(v) {
            result.m[x] = struct{}{}
	    }
	}

	return result
}

// CountBy gives the number elements of SyncAppleSet that return true for the passed predicate.
func (set SyncAppleSet) CountBy(predicate func(Apple) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of SyncAppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set SyncAppleSet) MinBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m Apple
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

// MaxBy returns an element of SyncAppleSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set SyncAppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m Apple
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

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set SyncAppleSet) Equals(other SyncAppleSet) bool {
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


