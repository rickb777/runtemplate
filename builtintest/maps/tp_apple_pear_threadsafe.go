// Generated from threadsafe.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Numeric=<no value> Stringer=<no value> Mutable=<no value>

package maps

import (

	"sync"
)


// TPApplePearMap is the primary type that represents a thread-safe map
type TPApplePearMap struct {
	s *sync.RWMutex
	m map[*Apple]*Pear
}

// TPApplePearTuple represents a key/value pair.
type TPApplePearTuple struct {
	Key *Apple
	Val *Pear
}

// TPApplePearTuples can be used as a builder for unmodifiable maps.
type TPApplePearTuples []TPApplePearTuple

func (ts TPApplePearTuples) Append1(k *Apple, v *Pear) TPApplePearTuples {
	return append(ts, TPApplePearTuple{k, v})
}

func (ts TPApplePearTuples) Append2(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear) TPApplePearTuples {
	return append(ts, TPApplePearTuple{k1, v1}, TPApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

// NewTPApplePearMap creates and returns a reference to a map containing one item.
func NewTPApplePearMap1(k *Apple, v *Pear) TPApplePearMap {
	mm := TPApplePearMap{
	    s: &sync.RWMutex{},
		m: make(map[*Apple]*Pear),
	}
	mm.m[k] = v
	return mm
}

// NewTPApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTPApplePearMap(kv ...TPApplePearTuple) TPApplePearMap {
	mm := TPApplePearMap{
	    s: &sync.RWMutex{},
		m: make(map[*Apple]*Pear),
	}
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TPApplePearMap) Keys() []*Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []*Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TPApplePearMap) ToSlice() []TPApplePearTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TPApplePearTuple
	for k, v := range mm.m {
		s = append(s, TPApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *TPApplePearMap) Get(k *Apple) (*Pear, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}


// ContainsKey determines if a given item is already in the map.
func (mm *TPApplePearMap) ContainsKey(k *Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TPApplePearMap) ContainsAllKeys(kk ...*Apple) bool {
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
func (mm *TPApplePearMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TPApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TPApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TPApplePearMap) Forall(fn func(*Apple, *Pear) bool) bool {
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
func (mm *TPApplePearMap) Exists(fn func(*Apple, *Pear) bool) bool {
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
func (mm *TPApplePearMap) Filter(fn func(*Apple, *Pear) bool) TPApplePearMap {
	result := NewTPApplePearMap()
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
func (mm *TPApplePearMap) Partition(fn func(*Apple, *Pear) bool) (matching TPApplePearMap, others TPApplePearMap) {
	matching = NewTPApplePearMap()
	others = NewTPApplePearMap()
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
func (mm *TPApplePearMap) Clone() TPApplePearMap {
	result := NewTPApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


