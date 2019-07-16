// A simple type derived from map[int]int.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"fmt"
)

// TX1IntIntMap is the primary type that represents a map
type TX1IntIntMap map[int]int

// TX1IntIntTuple represents a key/value pair.
type TX1IntIntTuple struct {
	Key int
	Val int
}

// TX1IntIntTuples can be used as a builder for unmodifiable maps.
type TX1IntIntTuples []TX1IntIntTuple

// Append1 adds one item.
func (ts TX1IntIntTuples) Append1(k int, v int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k, v})
}

// Append2 adds two items.
func (ts TX1IntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k1, v1}, TX1IntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1IntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) TX1IntIntTuples {
	return append(ts, TX1IntIntTuple{k1, v1}, TX1IntIntTuple{k2, v2}, TX1IntIntTuple{k3, v3})
}

// TX1IntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1IntIntMap
// constructor function.
func TX1IntIntZip(keys ...int) TX1IntIntTuples {
	ts := make(TX1IntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1IntIntZip.
func (ts TX1IntIntTuples) Values(values ...int) TX1IntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX1IntIntMap() TX1IntIntMap {
	return TX1IntIntMap(make(map[int]int))
}

// NewTX1IntIntMap1 creates and returns a reference to a map containing one item.
func NewTX1IntIntMap1(k int, v int) TX1IntIntMap {
	mm := newTX1IntIntMap()
	mm[k] = v
	return mm
}

// NewTX1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX1IntIntMap(kv ...TX1IntIntTuple) TX1IntIntMap {
	mm := newTX1IntIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1IntIntMap) Keys() []int {
	s := make([]int, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1IntIntMap) Values() []int {
	s := make([]int, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TX1IntIntMap) slice() []TX1IntIntTuple {
	s := make([]TX1IntIntTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, TX1IntIntTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1IntIntMap) ToSlice() []TX1IntIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm TX1IntIntMap) Get(k int) (int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1IntIntMap) Put(k int, v int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1IntIntMap) ContainsKey(k int) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1IntIntMap) ContainsAllKeys(kk ...int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1IntIntMap) Clear() {
	*mm = make(map[int]int)
}

// Remove a single item from the map.
func (mm TX1IntIntMap) Remove(k int) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TX1IntIntMap) Pop(k int) (int, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1IntIntMap) Size() int {
	return len(mm)
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
// This is similar to Filter except that the map is modified.
func (mm TX1IntIntMap) DropWhere(fn func(int, int) bool) TX1IntIntTuples {
	removed := make(TX1IntIntTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, TX1IntIntTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1IntIntMap) Foreach(f func(int, int)) {
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
func (mm TX1IntIntMap) Forall(p func(int, int) bool) bool {
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
func (mm TX1IntIntMap) Exists(p func(int, int) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TX1IntIntMap) Find(p func(int, int) bool) (TX1IntIntTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return TX1IntIntTuple{(k), v}, true
		}
	}

	return TX1IntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TX1IntIntMap) Filter(p func(int, int) bool) TX1IntIntMap {
	result := NewTX1IntIntMap()
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
func (mm TX1IntIntMap) Partition(p func(int, int) bool) (matching TX1IntIntMap, others TX1IntIntMap) {
	matching = NewTX1IntIntMap()
	others = NewTX1IntIntMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX1IntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1IntIntMap) Map(f func(int, int) (int, int)) TX1IntIntMap {
	result := NewTX1IntIntMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1IntIntMap) FlatMap(f func(int, int) []TX1IntIntTuple) TX1IntIntMap {
	result := NewTX1IntIntMap()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TX1IntIntMap) Equals(other TX1IntIntMap) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1IntIntMap) Clone() TX1IntIntMap {
	result := NewTX1IntIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm TX1IntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
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

	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}
