// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package examples

import (

	"bytes"
	"fmt"
	"sync"
)

// SyncIntIntMap is the primary type that represents a thread-safe map
type SyncIntIntMap struct {
	s *sync.RWMutex
	m map[int]int
}

// SyncIntIntTuple represents a key/value pair.
type SyncIntIntTuple struct {
	Key int
	Val int
}

// SyncIntIntTuples can be used as a builder for unmodifiable maps.
type SyncIntIntTuples []SyncIntIntTuple

func (ts SyncIntIntTuples) Append1(k int, v int) SyncIntIntTuples {
	return append(ts, SyncIntIntTuple{k, v})
}

func (ts SyncIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) SyncIntIntTuples {
	return append(ts, SyncIntIntTuple{k1, v1}, SyncIntIntTuple{k2, v2})
}

func (ts SyncIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) SyncIntIntTuples {
	return append(ts, SyncIntIntTuple{k1, v1}, SyncIntIntTuple{k2, v2}, SyncIntIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSyncIntIntMap() SyncIntIntMap {
	return SyncIntIntMap{
		s: &sync.RWMutex{},
		m: make(map[int]int),
	}
}

// NewSyncIntIntMap creates and returns a reference to a map containing one item.
func NewSyncIntIntMap1(k int, v int) SyncIntIntMap {
	mm := newSyncIntIntMap()
	mm.m[k] = v
	return mm
}

// NewSyncIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewSyncIntIntMap(kv ...SyncIntIntTuple) SyncIntIntMap {
	mm := newSyncIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SyncIntIntMap) Keys() []int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SyncIntIntMap) Values() []int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SyncIntIntMap) ToSlice() []SyncIntIntTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []SyncIntIntTuple
	for k, v := range mm.m {
		s = append(s, SyncIntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SyncIntIntMap) Get(k int) (int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SyncIntIntMap) Put(k int, v int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SyncIntIntMap) ContainsKey(k int) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SyncIntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *SyncIntIntMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[int]int)
}

// Remove a single item from the map.
func (mm SyncIntIntMap) Remove(k int) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm SyncIntIntMap) Pop(k int) (int, bool) {
	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SyncIntIntMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm SyncIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SyncIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm SyncIntIntMap) DropWhere(fn func(int, int) bool) SyncIntIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SyncIntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SyncIntIntTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm SyncIntIntMap) Foreach(fn func(int, int)) {
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
func (mm SyncIntIntMap) Forall(fn func(int, int) bool) bool {
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
func (mm SyncIntIntMap) Exists(fn func(int, int) bool) bool {
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
func (mm SyncIntIntMap) Find(fn func(int, int) bool) (SyncIntIntTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return SyncIntIntTuple{k, v}, true
		}
	}

	return SyncIntIntTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SyncIntIntMap) Filter(fn func(int, int) bool) SyncIntIntMap {
	result := NewSyncIntIntMap()
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
func (mm SyncIntIntMap) Partition(fn func(int, int) bool) (matching SyncIntIntMap, others SyncIntIntMap) {
	matching = NewSyncIntIntMap()
	others = NewSyncIntIntMap()
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

// Map returns a new SyncIntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SyncIntIntMap) Map(fn func(int, int) (int, int)) SyncIntIntMap {
	result := NewSyncIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SyncIntMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SyncIntIntMap) FlatMap(fn func(int, int) []SyncIntIntTuple) SyncIntIntMap {
	result := NewSyncIntIntMap()
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


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm SyncIntIntMap) Equals(other SyncIntIntMap) bool {
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
func (mm SyncIntIntMap) Clone() SyncIntIntMap {
	result := NewSyncIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm SyncIntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm SyncIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm SyncIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm SyncIntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm SyncIntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

