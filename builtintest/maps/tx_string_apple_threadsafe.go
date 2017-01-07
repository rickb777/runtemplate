// Generated from threadsafe.tpl with Key=string Type=Apple
// options: Comparable=<no value> Numeric=<no value> Stringer=<no value> Mutable=true

package maps

import (

	"sync"
)

// TXStringAppleMap is the primary type that represents a thread-safe map
type TXStringAppleMap struct {
	s *sync.RWMutex
	m map[string]Apple
}

// TXStringAppleTuple represents a key/value pair.
type TXStringAppleTuple struct {
	Key string
	Val Apple
}

// TXStringAppleTuples can be used as a builder for unmodifiable maps.
type TXStringAppleTuples []TXStringAppleTuple

func (ts TXStringAppleTuples) Append1(k string, v Apple) TXStringAppleTuples {
    return append(ts, TXStringAppleTuple{k, v})
}

func (ts TXStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) TXStringAppleTuples {
    return append(ts, TXStringAppleTuple{k1, v1}, TXStringAppleTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

// NewTXStringAppleMap creates and returns a reference to a map containing one item.
func NewTXStringAppleMap1(k string, v Apple) TXStringAppleMap {
	mm := TXStringAppleMap{
		m: make(map[string]Apple),
	}
    mm.m[k] = v
	return mm
}

// NewTXStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTXStringAppleMap(kv ...TXStringAppleTuple) TXStringAppleMap {
	mm := TXStringAppleMap{
		m: make(map[string]Apple),
	}
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TXStringAppleMap) Keys() []string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TXStringAppleMap) ToSlice() []TXStringAppleTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TXStringAppleTuple
	for k, v := range mm.m {
		s = append(s, TXStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *TXStringAppleMap) Get(k string) (Apple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *TXStringAppleMap) Put(k string, v Apple) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *TXStringAppleMap) ContainsKey(k string) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TXStringAppleMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *TXStringAppleMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[string]Apple)
}

// Remove allows the removal of a single item from the map.
func (mm *TXStringAppleMap) Remove(k string) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TXStringAppleMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TXStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TXStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TXStringAppleMap) Forall(fn func(string, Apple) bool) bool {
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
func (mm *TXStringAppleMap) Exists(fn func(string, Apple) bool) bool {
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
func (mm *TXStringAppleMap) Filter(fn func(string, Apple) bool) TXStringAppleMap {
	result := NewTXStringAppleMap()
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
func (mm *TXStringAppleMap) Partition(fn func(string, Apple) bool) (matching TXStringAppleMap, others TXStringAppleMap) {
	matching = NewTXStringAppleMap()
	others = NewTXStringAppleMap()
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


// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TXStringAppleMap) Clone() TXStringAppleMap {
	result := NewTXStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


