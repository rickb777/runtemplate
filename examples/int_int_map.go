// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v2.3.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"fmt"
	"sync"
)

// IntIntMap is the primary type that represents a thread-safe map
type IntIntMap struct {
	s *sync.RWMutex
	m map[int]int
}

// IntIntTuple represents a key/value pair.
type IntIntTuple struct {
	Key int
	Val int
}

// IntIntTuples can be used as a builder for unmodifiable maps.
type IntIntTuples []IntIntTuple

func (ts IntIntTuples) Append1(k int, v int) IntIntTuples {
	return append(ts, IntIntTuple{k, v})
}

func (ts IntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2})
}

func (ts IntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2}, IntIntTuple{k3, v3})
}

// IntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIntIntMap
// constructor function.
func IntIntZip(keys ...int) IntIntTuples {
	ts := make(IntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IntIntZip.
func (ts IntIntTuples) Values(values ...int) IntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIntIntMap() *IntIntMap {
	return &IntIntMap{
		s: &sync.RWMutex{},
		m: make(map[int]int),
	}
}

// NewIntIntMap creates and returns a reference to a map containing one item.
func NewIntIntMap1(k int, v int) *IntIntMap {
	mm := newIntIntMap()
	mm.m[k] = v
	return mm
}

// NewIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewIntIntMap(kv ...IntIntTuple) *IntIntMap {
	mm := newIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *IntIntMap) Keys() []int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for k, _ := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *IntIntMap) Values() []int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []int
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *IntIntMap) slice() []IntIntTuple {
	if mm == nil {
		return nil
	}

	var s []IntIntTuple
	for k, v := range mm.m {
		s = append(s, IntIntTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *IntIntMap) ToSlice() []IntIntTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *IntIntMap) Get(k int) (int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *IntIntMap) Put(k int, v int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *IntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *IntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *IntIntMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int]int)
	}
}

// Remove a single item from the map.
func (mm *IntIntMap) Remove(k int) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present until removal.
// The boolean result is true only if the key had been present.
func (mm *IntIntMap) Pop(k int) (int, bool) {
	if mm == nil {
		return 0, false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *IntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *IntIntMap) DropWhere(fn func(int, int) bool) IntIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(IntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, IntIntTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *IntIntMap) Foreach(fn func(int, int)) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		for k, v := range mm.m {
			fn(k, v)
		}
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *IntIntMap) Forall(fn func(int, int) bool) bool {
	if mm == nil {
		return true
	}

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
func (mm *IntIntMap) Exists(fn func(int, int) bool) bool {
	if mm == nil {
		return false
	}

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
// The original map is not modified.
func (mm *IntIntMap) Find(fn func(int, int) bool) (IntIntTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return IntIntTuple{k, v}, true
		}
	}

	return IntIntTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *IntIntMap) Filter(fn func(int, int) bool) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()
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
// The original map is not modified.
func (mm *IntIntMap) Partition(fn func(int, int) bool) (matching *IntIntMap, others *IntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewIntIntMap()
	others = NewIntIntMap()
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

// Map returns a new IntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IntIntMap) Map(fn func(int, int) (int, int)) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := fn(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new IntMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IntIntMap) FlatMap(fn func(int, int) []IntIntTuple) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()
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
func (mm *IntIntMap) Equals(other *IntIntMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

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
func (mm *IntIntMap) Clone() *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *IntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *IntIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
