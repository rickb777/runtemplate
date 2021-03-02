// A simple type derived from map[string]Apple.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
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

// Append1 adds one item.
func (ts SimpleStringAppleTuples) Append1(k string, v Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k, v})
}

// Append2 adds two items.
func (ts SimpleStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k1, v1}, SimpleStringAppleTuple{k2, v2})
}

// Append3 adds three items.
func (ts SimpleStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) SimpleStringAppleTuples {
	return append(ts, SimpleStringAppleTuple{k1, v1}, SimpleStringAppleTuple{k2, v2}, SimpleStringAppleTuple{k3, v3})
}

// SimpleStringAppleZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSimpleStringAppleMap
// constructor function.
func SimpleStringAppleZip(keys ...string) SimpleStringAppleTuples {
	ts := make(SimpleStringAppleTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SimpleStringAppleZip.
func (ts SimpleStringAppleTuples) Values(values ...Apple) SimpleStringAppleTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts SimpleStringAppleTuples) ToMap() SimpleStringAppleMap {
	return NewSimpleStringAppleMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newSimpleStringAppleMap() SimpleStringAppleMap {
	return SimpleStringAppleMap(make(map[string]Apple))
}

// NewSimpleStringAppleMap1 creates and returns a reference to a map containing one item.
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
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SimpleStringAppleMap) Values() []Apple {
	if mm == nil {
		return nil
	}

	s := make([]Apple, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm SimpleStringAppleMap) slice() SimpleStringAppleTuples {
	s := make(SimpleStringAppleTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, SimpleStringAppleTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm SimpleStringAppleMap) ToSlice() SimpleStringAppleTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm SimpleStringAppleMap) OrderedSlice(keys []string) SimpleStringAppleTuples {
	s := make(SimpleStringAppleTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, SimpleStringAppleTuple{k, v})
		}
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

// Pop removes a single item from the map, returning the value present prior to removal.
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
// This is similar to Filter except that the map is modified.
func (mm SimpleStringAppleMap) DropWhere(fn func(string, Apple) bool) SimpleStringAppleTuples {
	removed := make(SimpleStringAppleTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, SimpleStringAppleTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm SimpleStringAppleMap) Foreach(f func(string, Apple)) {
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
func (mm SimpleStringAppleMap) Forall(p func(string, Apple) bool) bool {
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
func (mm SimpleStringAppleMap) Exists(p func(string, Apple) bool) bool {
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
func (mm SimpleStringAppleMap) Find(p func(string, Apple) bool) (SimpleStringAppleTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return SimpleStringAppleTuple{(k), v}, true
		}
	}

	return SimpleStringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm SimpleStringAppleMap) Filter(p func(string, Apple) bool) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()
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
func (mm SimpleStringAppleMap) Partition(p func(string, Apple) bool) (matching SimpleStringAppleMap, others SimpleStringAppleMap) {
	matching = NewSimpleStringAppleMap()
	others = NewSimpleStringAppleMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new SimpleAppleMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SimpleStringAppleMap) Map(f func(string, Apple) (string, Apple)) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new SimpleAppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SimpleStringAppleMap) FlatMap(f func(string, Apple) []SimpleStringAppleTuple) SimpleStringAppleMap {
	result := NewSimpleStringAppleMap()

	for k1, v1 := range mm {
		ts := f(k1, v1)
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

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm SimpleStringAppleMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm SimpleStringAppleMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm SimpleStringAppleMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm SimpleStringAppleMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

func (ts SimpleStringAppleTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts SimpleStringAppleTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts SimpleStringAppleTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts SimpleStringAppleTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t SimpleStringAppleTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t SimpleStringAppleTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
