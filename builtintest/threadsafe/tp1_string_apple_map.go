// An encapsulated map[string]Apple.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always

package threadsafe

import (

	"sync"
)

// TP1StringAppleMap is the primary type that represents a thread-safe map
type TP1StringAppleMap struct {
	s *sync.RWMutex
	m map[*string]*Apple
}

// TP1StringAppleTuple represents a key/value pair.
type TP1StringAppleTuple struct {
	Key *string
	Val *Apple
}

// TP1StringAppleTuples can be used as a builder for unmodifiable maps.
type TP1StringAppleTuples []TP1StringAppleTuple

func (ts TP1StringAppleTuples) Append1(k *string, v *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k, v})
}

func (ts TP1StringAppleTuples) Append2(k1 *string, v1 *Apple, k2 *string, v2 *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k1, v1}, TP1StringAppleTuple{k2, v2})
}

func (ts TP1StringAppleTuples) Append3(k1 *string, v1 *Apple, k2 *string, v2 *Apple, k3 *string, v3 *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k1, v1}, TP1StringAppleTuple{k2, v2}, TP1StringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTP1StringAppleMap() TP1StringAppleMap {
	return TP1StringAppleMap{
		s: &sync.RWMutex{},
		m: make(map[*string]*Apple),
	}
}

// NewTP1StringAppleMap creates and returns a reference to a map containing one item.
func NewTP1StringAppleMap1(k *string, v *Apple) TP1StringAppleMap {
	mm := newTP1StringAppleMap()
	mm.m[k] = v
	return mm
}

// NewTP1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTP1StringAppleMap(kv ...TP1StringAppleTuple) TP1StringAppleMap {
	mm := newTP1StringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TP1StringAppleMap) Keys() []*string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []*string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TP1StringAppleMap) Values() []*Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []*Apple
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TP1StringAppleMap) ToSlice() []TP1StringAppleTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TP1StringAppleTuple
	for k, v := range mm.m {
		s = append(s, TP1StringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TP1StringAppleMap) Get(k *string) (*Apple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TP1StringAppleMap) Put(k *string, v *Apple) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TP1StringAppleMap) ContainsKey(k *string) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TP1StringAppleMap) ContainsAllKeys(kk ...*string) bool {
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
func (mm *TP1StringAppleMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[*string]*Apple)
}

// Remove a single item from the map.
func (mm TP1StringAppleMap) Remove(k *string) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm TP1StringAppleMap) Pop(k *string) (*Apple, bool) {
	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TP1StringAppleMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TP1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TP1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TP1StringAppleMap) DropWhere(fn func(*string, *Apple) bool) TP1StringAppleTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TP1StringAppleTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TP1StringAppleTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TP1StringAppleMap) Foreach(fn func(*string, *Apple)) {
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
func (mm TP1StringAppleMap) Forall(fn func(*string, *Apple) bool) bool {
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
func (mm TP1StringAppleMap) Exists(fn func(*string, *Apple) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (mm TP1StringAppleMap) Find(fn func(*string, *Apple) bool) (TP1StringAppleTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return TP1StringAppleTuple{k, v}, true
		}
	}

	return TP1StringAppleTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TP1StringAppleMap) Filter(fn func(*string, *Apple) bool) TP1StringAppleMap {
	result := NewTP1StringAppleMap()
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
func (mm TP1StringAppleMap) Partition(fn func(*string, *Apple) bool) (matching TP1StringAppleMap, others TP1StringAppleMap) {
	matching = NewTP1StringAppleMap()
	others = NewTP1StringAppleMap()
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

// Map returns a new TP1AppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringAppleMap) Map(fn func(*string, *Apple) (*string, *Apple)) TP1StringAppleMap {
	result := NewTP1StringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TP1AppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringAppleMap) FlatMap(fn func(*string, *Apple) []TP1StringAppleTuple) TP1StringAppleMap {
	result := NewTP1StringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TP1StringAppleMap) Clone() TP1StringAppleMap {
	result := NewTP1StringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


