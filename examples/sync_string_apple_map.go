// An encapsulated map[string]Apple.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package examples

import (

	"bytes"
	"fmt"
	"sync"
)

// SyncStringAppleMap is the primary type that represents a thread-safe map
type SyncStringAppleMap struct {
	s *sync.RWMutex
	m map[string]Apple
}

// SyncStringAppleTuple represents a key/value pair.
type SyncStringAppleTuple struct {
	Key string
	Val Apple
}

// SyncStringAppleTuples can be used as a builder for unmodifiable maps.
type SyncStringAppleTuples []SyncStringAppleTuple

func (ts SyncStringAppleTuples) Append1(k string, v Apple) SyncStringAppleTuples {
	return append(ts, SyncStringAppleTuple{k, v})
}

func (ts SyncStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) SyncStringAppleTuples {
	return append(ts, SyncStringAppleTuple{k1, v1}, SyncStringAppleTuple{k2, v2})
}

func (ts SyncStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) SyncStringAppleTuples {
	return append(ts, SyncStringAppleTuple{k1, v1}, SyncStringAppleTuple{k2, v2}, SyncStringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSyncStringAppleMap() SyncStringAppleMap {
	return SyncStringAppleMap{
		s: &sync.RWMutex{},
		m: make(map[string]Apple),
	}
}

// NewSyncStringAppleMap creates and returns a reference to a map containing one item.
func NewSyncStringAppleMap1(k string, v Apple) SyncStringAppleMap {
	mm := newSyncStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewSyncStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewSyncStringAppleMap(kv ...SyncStringAppleTuple) SyncStringAppleMap {
	mm := newSyncStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SyncStringAppleMap) Keys() []string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SyncStringAppleMap) Values() []Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SyncStringAppleMap) ToSlice() []SyncStringAppleTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []SyncStringAppleTuple
	for k, v := range mm.m {
		s = append(s, SyncStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SyncStringAppleMap) Get(k string) (Apple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SyncStringAppleMap) Put(k string, v Apple) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SyncStringAppleMap) ContainsKey(k string) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SyncStringAppleMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *SyncStringAppleMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[string]Apple)
}

// Remove a single item from the map.
func (mm SyncStringAppleMap) Remove(k string) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm SyncStringAppleMap) Pop(k string) (Apple, bool) {
	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SyncStringAppleMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm SyncStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SyncStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm SyncStringAppleMap) DropWhere(fn func(string, Apple) bool) SyncStringAppleTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SyncStringAppleTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SyncStringAppleTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm SyncStringAppleMap) Foreach(fn func(string, Apple)) {
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
func (mm SyncStringAppleMap) Forall(fn func(string, Apple) bool) bool {
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
func (mm SyncStringAppleMap) Exists(fn func(string, Apple) bool) bool {
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
func (mm SyncStringAppleMap) Find(fn func(string, Apple) bool) (SyncStringAppleTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return SyncStringAppleTuple{k, v}, true
		}
	}

	return SyncStringAppleTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SyncStringAppleMap) Filter(fn func(string, Apple) bool) SyncStringAppleMap {
	result := NewSyncStringAppleMap()
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
func (mm SyncStringAppleMap) Partition(fn func(string, Apple) bool) (matching SyncStringAppleMap, others SyncStringAppleMap) {
	matching = NewSyncStringAppleMap()
	others = NewSyncStringAppleMap()
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

// Map returns a new SyncAppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SyncStringAppleMap) Map(fn func(string, Apple) (string, Apple)) SyncStringAppleMap {
	result := NewSyncStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SyncAppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SyncStringAppleMap) FlatMap(fn func(string, Apple) []SyncStringAppleTuple) SyncStringAppleMap {
	result := NewSyncStringAppleMap()
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
func (mm SyncStringAppleMap) Clone() SyncStringAppleMap {
	result := NewSyncStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm SyncStringAppleMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm SyncStringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm SyncStringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm SyncStringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm SyncStringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

