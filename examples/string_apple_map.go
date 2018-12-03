// An encapsulated map[string]Apple.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v2.4.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

// StringAppleMap is the primary type that represents a thread-safe map
type StringAppleMap struct {
	s *sync.RWMutex
	m map[string]Apple
}

// StringAppleTuple represents a key/value pair.
type StringAppleTuple struct {
	Key string
	Val Apple
}

// StringAppleTuples can be used as a builder for unmodifiable maps.
type StringAppleTuples []StringAppleTuple

func (ts StringAppleTuples) Append1(k string, v Apple) StringAppleTuples {
	return append(ts, StringAppleTuple{k, v})
}

func (ts StringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) StringAppleTuples {
	return append(ts, StringAppleTuple{k1, v1}, StringAppleTuple{k2, v2})
}

func (ts StringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) StringAppleTuples {
	return append(ts, StringAppleTuple{k1, v1}, StringAppleTuple{k2, v2}, StringAppleTuple{k3, v3})
}

// StringAppleZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewStringAppleMap
// constructor function.
func StringAppleZip(keys ...string) StringAppleTuples {
	ts := make(StringAppleTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with StringAppleZip.
func (ts StringAppleTuples) Values(values ...Apple) StringAppleTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newStringAppleMap() *StringAppleMap {
	return &StringAppleMap{
		s: &sync.RWMutex{},
		m: make(map[string]Apple),
	}
}

// NewStringAppleMap creates and returns a reference to a map containing one item.
func NewStringAppleMap1(k string, v Apple) *StringAppleMap {
	mm := newStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewStringAppleMap(kv ...StringAppleTuple) *StringAppleMap {
	mm := newStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *StringAppleMap) Keys() []string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *StringAppleMap) Values() []Apple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *StringAppleMap) slice() []StringAppleTuple {
	if mm == nil {
		return nil
	}

	var s []StringAppleTuple
	for k, v := range mm.m {
		s = append(s, StringAppleTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *StringAppleMap) ToSlice() []StringAppleTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *StringAppleMap) Get(k string) (Apple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *StringAppleMap) Put(k string, v Apple) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *StringAppleMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *StringAppleMap) ContainsAllKeys(kk ...string) bool {
	if mm == nil {
		return len(kk) == 0
	}

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
func (mm *StringAppleMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[string]Apple)
	}
}

// Remove a single item from the map.
func (mm *StringAppleMap) Remove(k string) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *StringAppleMap) Pop(k string) (Apple, bool) {
	if mm == nil {
		return *(new(Apple)), false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *StringAppleMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *StringAppleMap) DropWhere(fn func(string, Apple) bool) StringAppleTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(StringAppleTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, StringAppleTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *StringAppleMap) Foreach(f func(string, Apple)) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		for k, v := range mm.m {
			f(k, v)
		}
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *StringAppleMap) Forall(p func(string, Apple) bool) bool {
	if mm == nil {
		return true
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *StringAppleMap) Exists(p func(string, Apple) bool) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first Apple that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *StringAppleMap) Find(p func(string, Apple) bool) (StringAppleTuple, bool) {
	if mm == nil {
		return StringAppleTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return StringAppleTuple{k, v}, true
		}
	}

	return StringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *StringAppleMap) Filter(p func(string, Apple) bool) *StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			result.m[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm *StringAppleMap) Partition(p func(string, Apple) bool) (matching *StringAppleMap, others *StringAppleMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewStringAppleMap()
	others = NewStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new AppleMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringAppleMap) Map(f func(string, Apple) (string, Apple)) *StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new AppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringAppleMap) FlatMap(f func(string, Apple) []StringAppleTuple) *StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *StringAppleMap) Clone() *StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewStringAppleMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register Apple with the 'gob' package before this method is used.
func (mm *StringAppleMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobDecode implements 'gob' encoding for this map type.
// You must register Apple with the 'gob' package before this method is used.
func (mm *StringAppleMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
