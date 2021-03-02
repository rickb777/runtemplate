// An encapsulated map[int]int.
//
// Not thread-safe.
//
// Generated from fast/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
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

// Append1 adds one item.
func (ts FastIntIntTuples) Append1(k int, v int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k, v})
}

// Append2 adds two items.
func (ts FastIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k1, v1}, FastIntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts FastIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) FastIntIntTuples {
	return append(ts, FastIntIntTuple{k1, v1}, FastIntIntTuple{k2, v2}, FastIntIntTuple{k3, v3})
}

// FastIntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewFastIntIntMap
// constructor function.
func FastIntIntZip(keys ...int) FastIntIntTuples {
	ts := make(FastIntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with FastIntIntZip.
func (ts FastIntIntTuples) Values(values ...int) FastIntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts FastIntIntTuples) ToMap() *FastIntIntMap {
	return NewFastIntIntMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newFastIntIntMap() *FastIntIntMap {
	return &FastIntIntMap{
		m: make(map[int]int),
	}
}

// NewFastIntIntMap1 creates and returns a reference to a map containing one item.
func NewFastIntIntMap1(k int, v int) *FastIntIntMap {
	mm := newFastIntIntMap()
	mm.m[k] = v
	return mm
}

// NewFastIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewFastIntIntMap(kv ...FastIntIntTuple) *FastIntIntMap {
	mm := newFastIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *FastIntIntMap) Keys() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *FastIntIntMap) Values() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *FastIntIntMap) slice() FastIntIntTuples {
	if mm == nil {
		return nil
	}

	s := make(FastIntIntTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, FastIntIntTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *FastIntIntMap) ToSlice() FastIntIntTuples {
	if mm == nil {
		return nil
	}

	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *FastIntIntMap) OrderedSlice(keys []int) FastIntIntTuples {
	if mm == nil {
		return nil
	}

	s := make(FastIntIntTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, FastIntIntTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *FastIntIntMap) Get(k int) (int, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *FastIntIntMap) Put(k int, v int) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *FastIntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *FastIntIntMap) ContainsAllKeys(kk ...int) bool {
	if mm == nil {
		return len(kk) == 0
	}

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *FastIntIntMap) Clear() {
	if mm != nil {

		mm.m = make(map[int]int)
	}
}

// Remove a single item from the map.
func (mm *FastIntIntMap) Remove(k int) {
	if mm != nil {

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *FastIntIntMap) Pop(k int) (int, bool) {
	if mm == nil {
		return 0, false
	}

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *FastIntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *FastIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *FastIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *FastIntIntMap) DropWhere(fn func(int, int) bool) FastIntIntTuples {

	removed := make(FastIntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, FastIntIntTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *FastIntIntMap) Foreach(f func(int, int)) {
	if mm != nil {

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
func (mm *FastIntIntMap) Forall(p func(int, int) bool) bool {
	if mm == nil {
		return true
	}

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
func (mm *FastIntIntMap) Exists(p func(int, int) bool) bool {
	if mm == nil {
		return false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *FastIntIntMap) Find(p func(int, int) bool) (FastIntIntTuple, bool) {
	if mm == nil {
		return FastIntIntTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return FastIntIntTuple{(k), v}, true
		}
	}

	return FastIntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *FastIntIntMap) Filter(p func(int, int) bool) *FastIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewFastIntIntMap()

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
func (mm *FastIntIntMap) Partition(p func(int, int) bool) (matching *FastIntIntMap, others *FastIntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewFastIntIntMap()
	others = NewFastIntIntMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new FastIntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *FastIntIntMap) Map(f func(int, int) (int, int)) *FastIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewFastIntIntMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new FastIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *FastIntIntMap) FlatMap(f func(int, int) []FastIntIntTuple) *FastIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewFastIntIntMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *FastIntIntMap) Equals(other *FastIntIntMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

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
func (mm *FastIntIntMap) Clone() *FastIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewFastIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *FastIntIntMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *FastIntIntMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *FastIntIntMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *FastIntIntMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

func (ts FastIntIntTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts FastIntIntTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts FastIntIntTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts FastIntIntTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	sep := before
	for _, t := range ts {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", t.Key, equals, t.Val)
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t FastIntIntTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t FastIntIntTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
