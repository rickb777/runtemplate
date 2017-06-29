// An encapsulated map[Apple]Pear.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=Apple Type=Pear
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package threadsafe

import (

	"sync"
)

// TX1ApplePearMap is the primary type that represents a thread-safe map
type TX1ApplePearMap struct {
	s *sync.RWMutex
	m map[Apple]Pear
}

// TX1ApplePearTuple represents a key/value pair.
type TX1ApplePearTuple struct {
	Key Apple
	Val Pear
}

// TX1ApplePearTuples can be used as a builder for unmodifiable maps.
type TX1ApplePearTuples []TX1ApplePearTuple

func (ts TX1ApplePearTuples) Append1(k Apple, v Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k, v})
}

func (ts TX1ApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2})
}

func (ts TX1ApplePearTuples) Append3(k1 Apple, v1 Pear, k2 Apple, v2 Pear, k3 Apple, v3 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2}, TX1ApplePearTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1ApplePearMap() TX1ApplePearMap {
	return TX1ApplePearMap{
		s: &sync.RWMutex{},
		m: make(map[Apple]Pear),
	}
}

// NewTX1ApplePearMap creates and returns a reference to a map containing one item.
func NewTX1ApplePearMap1(k Apple, v Pear) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	mm.m[k] = v
	return mm
}

// NewTX1ApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTX1ApplePearMap(kv ...TX1ApplePearTuple) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1ApplePearMap) Keys() []Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1ApplePearMap) Values() []Pear {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Pear
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1ApplePearMap) ToSlice() []TX1ApplePearTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TX1ApplePearTuple
	for k, v := range mm.m {
		s = append(s, TX1ApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1ApplePearMap) Get(k Apple) (Pear, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1ApplePearMap) Put(k Apple, v Pear) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1ApplePearMap) ContainsKey(k Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1ApplePearMap) ContainsAllKeys(kk ...Apple) bool {
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
func (mm *TX1ApplePearMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[Apple]Pear)
}

// Remove allows the removal of a single item from the map.
func (mm TX1ApplePearMap) Remove(k Apple) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1ApplePearMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TX1ApplePearMap) DropWhere(fn func(Apple, Pear) bool) TX1ApplePearTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TX1ApplePearTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
		    removed = append(removed, TX1ApplePearTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1ApplePearMap) Foreach(fn func(Apple, Pear)) {
	mm.s.Lock()
	defer mm.s.Unlock()

	for k, v := range mm.m {
		fn(k, v)
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1ApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm TX1ApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first Pear that returns true for some function.
// False is returned if none match.
func (mm TX1ApplePearMap) Find(fn func(Apple, Pear) bool) (TX1ApplePearTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return TX1ApplePearTuple{k, v}, true
		}
	}

	return TX1ApplePearTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1ApplePearMap) Filter(fn func(Apple, Pear) bool) TX1ApplePearMap {
	result := NewTX1ApplePearMap()
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
func (mm TX1ApplePearMap) Partition(fn func(Apple, Pear) bool) (matching TX1ApplePearMap, others TX1ApplePearMap) {
	matching = NewTX1ApplePearMap()
	others = NewTX1ApplePearMap()
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
func (mm TX1ApplePearMap) Clone() TX1ApplePearMap {
	result := NewTX1ApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


