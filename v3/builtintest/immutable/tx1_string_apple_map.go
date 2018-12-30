// An encapsulated immutable map[string]Apple.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.1.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TX1StringAppleMap is the primary type that represents a thread-safe map
type TX1StringAppleMap struct {
	m map[string]Apple
}

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

func newTX1StringAppleMap() *TX1StringAppleMap {
	return &TX1StringAppleMap{
		m: make(map[string]Apple),
	}
}

// NewTX1StringAppleMap1 creates and returns a reference to a map containing one item.
func NewTX1StringAppleMap1(k string, v Apple) *TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	mm.m[k] = v
	return mm
}

// NewTX1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTX1StringAppleMap(kv ...TX1StringAppleTuple) *TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TX1StringAppleMap) Keys() []string {
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TX1StringAppleMap) Values() []Apple {
	if mm == nil {
		return nil
	}

	s := make([]Apple, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TX1StringAppleMap) slice() []TX1StringAppleTuple {
	if mm == nil {
		return nil
	}

	s := make([]TX1StringAppleTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TX1StringAppleTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TX1StringAppleMap) ToSlice() []TX1StringAppleTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TX1StringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm *TX1StringAppleMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TX1StringAppleMap) ContainsAllKeys(kk ...string) bool {
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

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TX1StringAppleMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TX1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TX1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TX1StringAppleMap) Foreach(f func(string, Apple)) {
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
func (mm *TX1StringAppleMap) Forall(f func(string, Apple) bool) bool {
	if mm == nil {
		return true
	}

	for k, v := range mm.m {
		if !f(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *TX1StringAppleMap) Exists(p func(string, Apple) bool) bool {
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

// Find returns the first Apple that returns true for the predicate p.
// False is returned if none match.
func (mm *TX1StringAppleMap) Find(p func(string, Apple) bool) (TX1StringAppleTuple, bool) {
	if mm == nil {
		return TX1StringAppleTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return TX1StringAppleTuple{k, v}, true
		}
	}

	return TX1StringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *TX1StringAppleMap) Filter(p func(string, Apple) bool) *TX1StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringAppleMap()

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
func (mm *TX1StringAppleMap) Partition(p func(string, Apple) bool) (matching *TX1StringAppleMap, others *TX1StringAppleMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTX1StringAppleMap()
	others = NewTX1StringAppleMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TX1AppleMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1StringAppleMap) Map(f func(string, Apple) (string, Apple)) *TX1StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringAppleMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1AppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1StringAppleMap) FlatMap(f func(string, Apple) []TX1StringAppleTuple) *TX1StringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringAppleMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm *TX1StringAppleMap) Clone() *TX1StringAppleMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *TX1StringAppleMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *TX1StringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *TX1StringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
// The map entries are sorted by their keys.
func (mm *TX1StringAppleMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *TX1StringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *TX1StringAppleMap) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *TX1StringAppleMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(mm.m)
}
