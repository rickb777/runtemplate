// Generated from threadsafe.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Numeric=<no value> Stringer=<no value> Mutable=<no value>

package maps

import (

	"sync"
)

// TXApplePearMap is the primary type that represents a thread-safe map
type TXApplePearMap struct {
	s *sync.RWMutex
	m map[Apple]Pear
}

// TXApplePearTuple represents a key/value pair.
type TXApplePearTuple struct {
	Key Apple
	Val Pear
}

// TXApplePearTuples can be used as a builder for unmodifiable maps.
type TXApplePearTuples []TXApplePearTuple

func (ts TXApplePearTuples) Append1(k Apple, v Pear) TXApplePearTuples {
    return append(ts, TXApplePearTuple{k, v})
}

func (ts TXApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) TXApplePearTuples {
    return append(ts, TXApplePearTuple{k1, v1}, TXApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

// NewTXApplePearMap creates and returns a reference to a map containing one item.
func NewTXApplePearMap1(k Apple, v Pear) TXApplePearMap {
	mm := TXApplePearMap{
		m: make(map[Apple]Pear),
	}
    mm.m[k] = v
	return mm
}

// NewTXApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTXApplePearMap(kv ...TXApplePearTuple) TXApplePearMap {
	mm := TXApplePearMap{
		m: make(map[Apple]Pear),
	}
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TXApplePearMap) Keys() []Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TXApplePearMap) ToSlice() []TXApplePearTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TXApplePearTuple
	for k, v := range mm.m {
		s = append(s, TXApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *TXApplePearMap) Get(k Apple) (Pear, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}


// ContainsKey determines if a given item is already in the map.
func (mm *TXApplePearMap) ContainsKey(k Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TXApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TXApplePearMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TXApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TXApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TXApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm *TXApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
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
func (mm *TXApplePearMap) Filter(fn func(Apple, Pear) bool) TXApplePearMap {
	result := NewTXApplePearMap()
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
func (mm *TXApplePearMap) Partition(fn func(Apple, Pear) bool) (matching TXApplePearMap, others TXApplePearMap) {
	matching = NewTXApplePearMap()
	others = NewTXApplePearMap()
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
func (mm *TXApplePearMap) Clone() TXApplePearMap {
	result := NewTXApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


