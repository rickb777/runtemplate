// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> Mutable:always

package threadsafe

import (

	"bytes"
	"fmt"
	"sync"
)

// TX1IntIntMap is the primary type that represents a thread-safe map
type TX1IntIntMap struct {
	s *sync.RWMutex
	m map[int]int
}

// TX1IntIntTuple represents a key/value pair.
type TX1IntIntTuple struct {
	Key int
	Val int
}

// TX1IntIntTuples can be used as a builder for unmodifiable maps.
type TX1IntIntTuples []TX1IntIntTuple

func (ts TX1IntIntTuples) Append1(k int, v int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k, v})
}

func (ts TX1IntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k1, v1}, TX1IntIntTuple{k2, v2})
}

func (ts TX1IntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k1, v1}, TX1IntIntTuple{k2, v2}, TX1IntIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1IntIntMap() TX1IntIntMap {
	return TX1IntIntMap{
		s: &sync.RWMutex{},
		m: make(map[int]int),
	}
}

// NewTX1IntIntMap creates and returns a reference to a map containing one item.
func NewTX1IntIntMap1(k int, v int) TX1IntIntMap {
	mm := newTX1IntIntMap()
	mm.m[k] = v
	return mm
}

// NewTX1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX1IntIntMap(kv ...TX1IntIntTuple) TX1IntIntMap {
	mm := newTX1IntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1IntIntMap) Keys() []int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1IntIntMap) Values() []int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1IntIntMap) ToSlice() []TX1IntIntTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TX1IntIntTuple
	for k, v := range mm.m {
		s = append(s, TX1IntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1IntIntMap) Get(k int) (int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1IntIntMap) Put(k int, v int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1IntIntMap) ContainsKey(k int) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1IntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *TX1IntIntMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[int]int)
}

// Remove allows the removal of a single item from the map.
func (mm TX1IntIntMap) Remove(k int) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1IntIntMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TX1IntIntMap) DropWhere(fn func(int, int) bool) TX1IntIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TX1IntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TX1IntIntTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1IntIntMap) Foreach(fn func(int, int)) {
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
func (mm TX1IntIntMap) Forall(fn func(int, int) bool) bool {
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
func (mm TX1IntIntMap) Exists(fn func(int, int) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (mm TX1IntIntMap) Find(fn func(int, int) bool) (TX1IntIntTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return TX1IntIntTuple{k, v}, true
		}
	}

	return TX1IntIntTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1IntIntMap) Filter(fn func(int, int) bool) TX1IntIntMap {
	result := NewTX1IntIntMap()
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
func (mm TX1IntIntMap) Partition(fn func(int, int) bool) (matching TX1IntIntMap, others TX1IntIntMap) {
	matching = NewTX1IntIntMap()
	others = NewTX1IntIntMap()
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

// Transform returns a new TX1IntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1IntIntMap) Transform(fn func(int, int) (int, int)) TX1IntIntMap {
	result := NewTX1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TX1IntIntMap) Equals(other TX1IntIntMap) bool {
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
func (mm TX1IntIntMap) Clone() TX1IntIntMap {
	result := NewTX1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TX1IntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TX1IntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

