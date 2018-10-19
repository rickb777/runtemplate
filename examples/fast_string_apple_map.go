// An encapsulated map[string]Apple.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package examples

import (

	"bytes"
	"fmt"
)

// FastStringAppleMap is the primary type that represents a thread-safe map
type FastStringAppleMap struct {
	m map[string]Apple
}

// FastStringAppleTuple represents a key/value pair.
type FastStringAppleTuple struct {
	Key string
	Val Apple
}

// FastStringAppleTuples can be used as a builder for unmodifiable maps.
type FastStringAppleTuples []FastStringAppleTuple

func (ts FastStringAppleTuples) Append1(k string, v Apple) FastStringAppleTuples {
	return append(ts, FastStringAppleTuple{k, v})
}

func (ts FastStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) FastStringAppleTuples {
	return append(ts, FastStringAppleTuple{k1, v1}, FastStringAppleTuple{k2, v2})
}

func (ts FastStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) FastStringAppleTuples {
	return append(ts, FastStringAppleTuple{k1, v1}, FastStringAppleTuple{k2, v2}, FastStringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newFastStringAppleMap() FastStringAppleMap {
	return FastStringAppleMap{
		m: make(map[string]Apple),
	}
}

// NewFastStringAppleMap creates and returns a reference to a map containing one item.
func NewFastStringAppleMap1(k string, v Apple) FastStringAppleMap {
	mm := newFastStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewFastStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewFastStringAppleMap(kv ...FastStringAppleTuple) FastStringAppleMap {
	mm := newFastStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm FastStringAppleMap) Keys() []string {

	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm FastStringAppleMap) Values() []Apple {

	var s []Apple
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm FastStringAppleMap) ToSlice() []FastStringAppleTuple {

	var s []FastStringAppleTuple
	for k, v := range mm.m {
		s = append(s, FastStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm FastStringAppleMap) Get(k string) (Apple, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm FastStringAppleMap) Put(k string, v Apple) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm FastStringAppleMap) ContainsKey(k string) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm FastStringAppleMap) ContainsAllKeys(kk ...string) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *FastStringAppleMap) Clear() {

	mm.m = make(map[string]Apple)
}

// Remove a single item from the map.
func (mm FastStringAppleMap) Remove(k string) {

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm FastStringAppleMap) Pop(k string) (Apple, bool) {

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm FastStringAppleMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm FastStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm FastStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm FastStringAppleMap) DropWhere(fn func(string, Apple) bool) FastStringAppleTuples {

	removed := make(FastStringAppleTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, FastStringAppleTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm FastStringAppleMap) Foreach(fn func(string, Apple)) {

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
func (mm FastStringAppleMap) Forall(fn func(string, Apple) bool) bool {

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
func (mm FastStringAppleMap) Exists(fn func(string, Apple) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (mm FastStringAppleMap) Find(fn func(string, Apple) bool) (FastStringAppleTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return FastStringAppleTuple{k, v}, true
		}
	}

	return FastStringAppleTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm FastStringAppleMap) Filter(fn func(string, Apple) bool) FastStringAppleMap {
	result := NewFastStringAppleMap()

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
func (mm FastStringAppleMap) Partition(fn func(string, Apple) bool) (matching FastStringAppleMap, others FastStringAppleMap) {
	matching = NewFastStringAppleMap()
	others = NewFastStringAppleMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new FastAppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm FastStringAppleMap) Map(fn func(string, Apple) (string, Apple)) FastStringAppleMap {
	result := NewFastStringAppleMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new FastAppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm FastStringAppleMap) FlatMap(fn func(string, Apple) []FastStringAppleTuple) FastStringAppleMap {
	result := NewFastStringAppleMap()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm FastStringAppleMap) Clone() FastStringAppleMap {
	result := NewFastStringAppleMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm FastStringAppleMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm FastStringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm FastStringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm FastStringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm FastStringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
