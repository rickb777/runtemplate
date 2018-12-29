// A simple type derived from map[string]Apple.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"fmt"
)

// TX1StringAppleMap is the primary type that represents a map
type TX1StringAppleMap map[string]Apple

// TX1StringAppleTuple represents a key/value pair.
type TX1StringAppleTuple struct {
	Key string
	Val Apple
}

// TX1StringAppleTuples can be used as a builder for unmodifiable maps.
type TX1StringAppleTuples []TX1StringAppleTuple

// Append1 adds one item.
func (ts TX1StringAppleTuples) Append1(k string, v Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k, v})
}

// Append2 adds two items.
func (ts TX1StringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k1, v1}, TX1StringAppleTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1StringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k1, v1}, TX1StringAppleTuple{k2, v2}, TX1StringAppleTuple{k3, v3})
}

// TX1StringAppleZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1StringAppleMap
// constructor function.
func TX1StringAppleZip(keys ...string) TX1StringAppleTuples {
	ts := make(TX1StringAppleTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1StringAppleZip.
func (ts TX1StringAppleTuples) Values(values ...Apple) TX1StringAppleTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX1StringAppleMap() TX1StringAppleMap {
	return TX1StringAppleMap(make(map[string]Apple))
}

// NewTX1StringAppleMap1 creates and returns a reference to a map containing one item.
func NewTX1StringAppleMap1(k string, v Apple) TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	mm[k] = v
	return mm
}

// NewTX1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTX1StringAppleMap(kv ...TX1StringAppleTuple) TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1StringAppleMap) Keys() []string {
	s := make([]string, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1StringAppleMap) Values() []Apple {
	s := make([]Apple, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TX1StringAppleMap) slice() []TX1StringAppleTuple {
	s := make([]TX1StringAppleTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, TX1StringAppleTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1StringAppleMap) ToSlice() []TX1StringAppleTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm TX1StringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1StringAppleMap) Put(k string, v Apple) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1StringAppleMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1StringAppleMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1StringAppleMap) Clear() {
	*mm = make(map[string]Apple)
}

// Remove a single item from the map.
func (mm TX1StringAppleMap) Remove(k string) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TX1StringAppleMap) Pop(k string) (Apple, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1StringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm TX1StringAppleMap) DropWhere(fn func(string, Apple) bool) TX1StringAppleTuples {
	removed := make(TX1StringAppleTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, TX1StringAppleTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1StringAppleMap) Foreach(f func(string, Apple)) {
	for k, v := range mm {
		f(k, v)
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1StringAppleMap) Forall(p func(string, Apple) bool) bool {
	for k, v := range mm {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TX1StringAppleMap) Exists(p func(string, Apple) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first Apple that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TX1StringAppleMap) Find(p func(string, Apple) bool) (TX1StringAppleTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return TX1StringAppleTuple{(k), v}, true
		}
	}

	return TX1StringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TX1StringAppleMap) Filter(p func(string, Apple) bool) TX1StringAppleMap {
	result := NewTX1StringAppleMap()
	for k, v := range mm {
		if p(k, v) {
			result[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm TX1StringAppleMap) Partition(p func(string, Apple) bool) (matching TX1StringAppleMap, others TX1StringAppleMap) {
	matching = NewTX1StringAppleMap()
	others = NewTX1StringAppleMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX1AppleMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1StringAppleMap) Map(f func(string, Apple) (string, Apple)) TX1StringAppleMap {
	result := NewTX1StringAppleMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1AppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1StringAppleMap) FlatMap(f func(string, Apple) []TX1StringAppleTuple) TX1StringAppleMap {
	result := NewTX1StringAppleMap()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1StringAppleMap) Clone() TX1StringAppleMap {
	result := NewTX1StringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm TX1StringAppleMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1StringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1StringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
// The map entries are sorted by their keys.
func (mm TX1StringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1StringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
