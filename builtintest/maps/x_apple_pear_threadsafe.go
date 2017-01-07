// Generated from threadsafe.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Numeric=<no value> Ordered=<no value> Stringer=<no value>

package maps

import (
	"sync"
)

// ApplePearMap is the primary type that represents a thread-safe map
type ApplePearMap struct {
	s *sync.RWMutex
	m map[Apple]Pear
}

// ApplePearTuple represents a key/value pair.
type ApplePearTuple struct {
	Key Apple
	Val Pear
}

// NewApplePearMap creates and returns a reference to an empty map.
func NewApplePearMap(kv ...ApplePearTuple) ApplePearMap {
	mm := ApplePearMap{
		s: &sync.RWMutex{},
		m: make(map[Apple]Pear),
	}
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *ApplePearMap) Keys() []Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *ApplePearMap) ToSlice() []ApplePearTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []ApplePearTuple
	for k, v := range mm.m {
		s = append(s, ApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *ApplePearMap) Get(k Apple) (Pear, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *ApplePearMap) Put(k Apple, v Pear) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *ApplePearMap) ContainsKey(k Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *ApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *ApplePearMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[Apple]Pear)
}

// Remove allows the removal of a single item from the map.
func (mm *ApplePearMap) Remove(k Apple) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *ApplePearMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *ApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *ApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *ApplePearMap) Filter(fn func(Apple, Pear) bool) ApplePearMap {
	result := NewApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			result.m[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm *ApplePearMap) Partition(fn func(Apple, Pear) bool) (matching ApplePearMap, others ApplePearMap) {
	matching = NewApplePearMap()
	others = NewApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (mm *ApplePearMap) Equals(other ApplePearMap) bool {
	mm.s.RLock()
	other.s.RLock()
	defer mm.s.RUnlock()
	defer other.s.RUnlock()

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *ApplePearMap) Clone() ApplePearMap {
	result := NewApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
