// A simple type derived from map[string]Apple.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package examples

import (
	"bytes"
	"fmt"
)

// SimpleStringAppleMap is the primary type that represents a map
type SimpleStringAppleMap map[string]Apple

// SimpleStringAppleTuple represents a key/value pair.
type SimpleStringAppleTuple struct {
	Key string
	Val Apple
}

// SimpleStringAppleTuples can be used as a builder for unmodifiable maps.
type SimpleStringAppleTuples []SimpleStringAppleTuple

func (ts SimpleStringAppleTuples) Append1(k string, v Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k, v})
}

func (ts SimpleStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k1, v1}, SimpleStringAppleTuple{k2, v2})
}

func (ts SimpleStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k1, v1}, SimpleStringAppleTuple{k2, v2}, SimpleStringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSimpleStringAppleMap() SimpleStringAppleMap {
	return SimpleStringAppleMap(make(map[string]Apple))
}

// NewSimpleStringAppleMap creates and returns a reference to a map containing one item.
func NewSimpleStringAppleMap1(k string, v Apple) SimpleStringAppleMap {
	mm := newSimpleStringAppleMap()
	mm[k] = v
	return mm
}

// NewSimpleStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewSimpleStringAppleMap(kv ...SimpleStringAppleTuple) SimpleStringAppleMap {
	mm := newSimpleStringAppleMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SimpleStringAppleMap) Keys() []string {
	var s []string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SimpleStringAppleMap) Values() []Apple {
	var s []Apple
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SimpleStringAppleMap) ToSlice() []SimpleStringAppleTuple {
	var s []SimpleStringAppleTuple
	for k, v := range mm {
		s = append(s, SimpleStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SimpleStringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SimpleStringAppleMap) Put(k string, v Apple) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SimpleStringAppleMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SimpleStringAppleMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SimpleStringAppleMap) Clear() {
	*mm = make(map[string]Apple)
}

// Remove a single item from the map.
func (mm SimpleStringAppleMap) Remove(k string) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm SimpleStringAppleMap) Pop(k string) (Apple, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SimpleStringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SimpleStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SimpleStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm SimpleStringAppleMap) DropWhere(fn func(string, Apple) bool) SimpleStringAppleTuples {
	removed := make(SimpleStringAppleTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, SimpleStringAppleTuple{k, v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm SimpleStringAppleMap) Foreach(fn func(string, Apple)) {
	for k, v := range mm {
		fn(k, v)
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SimpleStringAppleMap) Forall(fn func(string, Apple) bool) bool {
	for k, v := range mm {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm SimpleStringAppleMap) Exists(fn func(string, Apple) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
// The original map is not modified
func (mm SimpleStringAppleMap) Find(fn func(string, Apple) bool) (SimpleStringAppleTuple, bool) {
	for k, v := range mm {
		if fn(k, v) {
			return SimpleStringAppleTuple{k, v}, true
		}
	}

	return SimpleStringAppleTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm SimpleStringAppleMap) Filter(fn func(string, Apple) bool) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified
func (mm SimpleStringAppleMap) Partition(fn func(string, Apple) bool) (matching SimpleStringAppleMap, others SimpleStringAppleMap) {
	matching = NewSimpleStringAppleMap()
	others = NewSimpleStringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new SimpleAppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SimpleStringAppleMap) Map(fn func(string, Apple) (string, Apple)) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()

	for k1, v1 := range mm {
		k2, v2 := fn(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new SimpleAppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SimpleStringAppleMap) FlatMap(fn func(string, Apple) []SimpleStringAppleTuple) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()

	for k1, v1 := range mm {
		ts := fn(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SimpleStringAppleMap) Clone() SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm SimpleStringAppleMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm SimpleStringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm SimpleStringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm SimpleStringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm SimpleStringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}
