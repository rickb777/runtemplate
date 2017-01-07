// An encapsulated map[string]string
// Thread-safe.
//
// Generated from threadsafe.tpl with Key=string Type=string
// options: Comparable=true Stringer=<no value> Mutable=true

package maps

import (

	"sync"
)

// TPStringStringMap is the primary type that represents a thread-safe map
type TPStringStringMap struct {
	s *sync.RWMutex
	m map[*string]*string
}

// TPStringStringTuple represents a key/value pair.
type TPStringStringTuple struct {
	Key *string
	Val *string
}

// TPStringStringTuples can be used as a builder for unmodifiable maps.
type TPStringStringTuples []TPStringStringTuple

func (ts TPStringStringTuples) Append1(k *string, v *string) TPStringStringTuples {
	return append(ts, TPStringStringTuple{k, v})
}

func (ts TPStringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) TPStringStringTuples {
	return append(ts, TPStringStringTuple{k1, v1}, TPStringStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newTPStringStringMap() TPStringStringMap {
	return TPStringStringMap{
	    s: &sync.RWMutex{},
		m: make(map[*string]*string),
	}
}

// NewTPStringStringMap creates and returns a reference to a map containing one item.
func NewTPStringStringMap1(k *string, v *string) TPStringStringMap {
	mm := newTPStringStringMap()
	mm.m[k] = v
	return mm
}

// NewTPStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewTPStringStringMap(kv ...TPStringStringTuple) TPStringStringMap {
	mm := newTPStringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TPStringStringMap) Keys() []*string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []*string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TPStringStringMap) ToSlice() []TPStringStringTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TPStringStringTuple
	for k, v := range mm.m {
		s = append(s, TPStringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *TPStringStringMap) Get(k *string) (*string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *TPStringStringMap) Put(k *string, v *string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *TPStringStringMap) ContainsKey(k *string) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TPStringStringMap) ContainsAllKeys(kk ...*string) bool {
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
func (mm *TPStringStringMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[*string]*string)
}

// Remove allows the removal of a single item from the map.
func (mm *TPStringStringMap) Remove(k *string) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TPStringStringMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TPStringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TPStringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TPStringStringMap) Forall(fn func(*string, *string) bool) bool {
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
func (mm *TPStringStringMap) Exists(fn func(*string, *string) bool) bool {
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
func (mm *TPStringStringMap) Filter(fn func(*string, *string) bool) TPStringStringMap {
	result := NewTPStringStringMap()
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
func (mm *TPStringStringMap) Partition(fn func(*string, *string) bool) (matching TPStringStringMap, others TPStringStringMap) {
	matching = NewTPStringStringMap()
	others = NewTPStringStringMap()
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


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *TPStringStringMap) Equals(other TPStringStringMap) bool {
	mm.s.RLock()
	other.s.RLock()
	defer mm.s.RUnlock()
	defer other.s.RUnlock()

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || *v1 != *v2 {
			return false
		}
	}
	return true
}


// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TPStringStringMap) Clone() TPStringStringMap {
	result := NewTPStringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


