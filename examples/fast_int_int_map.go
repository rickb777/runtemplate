// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package examples

import (

	"bytes"
	"fmt"
)

// FastIntIntMap is the primary type that represents a thread-safe map
type FastIntIntMap struct {
	m map[int]int
}

// FastIntIntTuple represents a key/value pair.
type FastIntIntTuple struct {
	Key int
	Val int
}

// FastIntIntTuples can be used as a builder for unmodifiable maps.
type FastIntIntTuples []FastIntIntTuple

func (ts FastIntIntTuples) Append1(k int, v int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k, v})
}

func (ts FastIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k1, v1}, FastIntIntTuple{k2, v2})
}

func (ts FastIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k1, v1}, FastIntIntTuple{k2, v2}, FastIntIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newFastIntIntMap() FastIntIntMap {
	return FastIntIntMap{
		m: make(map[int]int),
	}
}

// NewFastIntIntMap creates and returns a reference to a map containing one item.
func NewFastIntIntMap1(k int, v int) FastIntIntMap {
	mm := newFastIntIntMap()
	mm.m[k] = v
	return mm
}

// NewFastIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewFastIntIntMap(kv ...FastIntIntTuple) FastIntIntMap {
	mm := newFastIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm FastIntIntMap) Keys() []int {

	var s []int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm FastIntIntMap) Values() []int {

	var s []int
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm FastIntIntMap) ToSlice() []FastIntIntTuple {

	var s []FastIntIntTuple
	for k, v := range mm.m {
		s = append(s, FastIntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm FastIntIntMap) Get(k int) (int, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm FastIntIntMap) Put(k int, v int) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm FastIntIntMap) ContainsKey(k int) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm FastIntIntMap) ContainsAllKeys(kk ...int) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *FastIntIntMap) Clear() {

	mm.m = make(map[int]int)
}

// Remove a single item from the map.
func (mm FastIntIntMap) Remove(k int) {

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm FastIntIntMap) Pop(k int) (int, bool) {

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm FastIntIntMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm FastIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm FastIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm FastIntIntMap) DropWhere(fn func(int, int) bool) FastIntIntTuples {

	removed := make(FastIntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, FastIntIntTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm FastIntIntMap) Foreach(fn func(int, int)) {

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
func (mm FastIntIntMap) Forall(fn func(int, int) bool) bool {

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
func (mm FastIntIntMap) Exists(fn func(int, int) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (mm FastIntIntMap) Find(fn func(int, int) bool) (FastIntIntTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return FastIntIntTuple{k, v}, true
		}
	}

	return FastIntIntTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm FastIntIntMap) Filter(fn func(int, int) bool) FastIntIntMap {
	result := NewFastIntIntMap()

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
func (mm FastIntIntMap) Partition(fn func(int, int) bool) (matching FastIntIntMap, others FastIntIntMap) {
	matching = NewFastIntIntMap()
	others = NewFastIntIntMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new FastIntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm FastIntIntMap) Map(fn func(int, int) (int, int)) FastIntIntMap {
	result := NewFastIntIntMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new FastIntMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm FastIntIntMap) FlatMap(fn func(int, int) []FastIntIntTuple) FastIntIntMap {
	result := NewFastIntIntMap()

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
func (mm FastIntIntMap) Equals(other FastIntIntMap) bool {

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
func (mm FastIntIntMap) Clone() FastIntIntMap {
	result := NewFastIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm FastIntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm FastIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm FastIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm FastIntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm FastIntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

