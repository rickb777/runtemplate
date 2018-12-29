// An encapsulated immutable map[string]Apple.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// ImmutableStringAppleMap is the primary type that represents a thread-safe map
type ImmutableStringAppleMap struct {
	m map[string]Apple
}

// ImmutableStringAppleTuple represents a key/value pair.
type ImmutableStringAppleTuple struct {
	Key string
	Val Apple
}

// ImmutableStringAppleTuples can be used as a builder for unmodifiable maps.
type ImmutableStringAppleTuples []ImmutableStringAppleTuple

// Append1 adds one item.
func (ts ImmutableStringAppleTuples) Append1(k string, v Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k, v})
}

// Append2 adds two items.
func (ts ImmutableStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k1, v1}, ImmutableStringAppleTuple{k2, v2})
}

// Append3 adds three items.
func (ts ImmutableStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k1, v1}, ImmutableStringAppleTuple{k2, v2}, ImmutableStringAppleTuple{k3, v3})
}

// ImmutableStringAppleZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewImmutableStringAppleMap
// constructor function.
func ImmutableStringAppleZip(keys ...string) ImmutableStringAppleTuples {
	ts := make(ImmutableStringAppleTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with ImmutableStringAppleZip.
func (ts ImmutableStringAppleTuples) Values(values ...Apple) ImmutableStringAppleTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newImmutableStringAppleMap() *ImmutableStringAppleMap {
	return &ImmutableStringAppleMap{
		m: make(map[string]Apple),
	}
}

// NewImmutableStringAppleMap1 creates and returns a reference to a map containing one item.
func NewImmutableStringAppleMap1(k string, v Apple) *ImmutableStringAppleMap {
	mm := newImmutableStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewImmutableStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewImmutableStringAppleMap(kv ...ImmutableStringAppleTuple) *ImmutableStringAppleMap {
	mm := newImmutableStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *ImmutableStringAppleMap) Keys() []string {
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
func (mm *ImmutableStringAppleMap) Values() []Apple {
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
func (mm *ImmutableStringAppleMap) slice() []ImmutableStringAppleTuple {
	if mm == nil {
		return nil
	}

	s := make([]ImmutableStringAppleTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, ImmutableStringAppleTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *ImmutableStringAppleMap) ToSlice() []ImmutableStringAppleTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *ImmutableStringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm *ImmutableStringAppleMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *ImmutableStringAppleMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *ImmutableStringAppleMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *ImmutableStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *ImmutableStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *ImmutableStringAppleMap) Foreach(f func(string, Apple)) {
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
func (mm *ImmutableStringAppleMap) Forall(f func(string, Apple) bool) bool {
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
func (mm *ImmutableStringAppleMap) Exists(p func(string, Apple) bool) bool {
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
func (mm *ImmutableStringAppleMap) Find(p func(string, Apple) bool) (ImmutableStringAppleTuple, bool) {
	if mm == nil {
		return ImmutableStringAppleTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return ImmutableStringAppleTuple{k, v}, true
		}
	}

	return ImmutableStringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *ImmutableStringAppleMap) Filter(p func(string, Apple) bool) *ImmutableStringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableStringAppleMap()

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
func (mm *ImmutableStringAppleMap) Partition(p func(string, Apple) bool) (matching *ImmutableStringAppleMap, others *ImmutableStringAppleMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewImmutableStringAppleMap()
	others = NewImmutableStringAppleMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new ImmutableAppleMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *ImmutableStringAppleMap) Map(f func(string, Apple) (string, Apple)) *ImmutableStringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableStringAppleMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new ImmutableAppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *ImmutableStringAppleMap) FlatMap(f func(string, Apple) []ImmutableStringAppleTuple) *ImmutableStringAppleMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableStringAppleMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm *ImmutableStringAppleMap) Clone() *ImmutableStringAppleMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register Apple with the 'gob' package before this method is used.
func (mm *ImmutableStringAppleMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (mm *ImmutableStringAppleMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
