// An encapsulated immutable map[Apple]string.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"fmt"
)

// TX1AppleStringMap is the primary type that represents a thread-safe map
type TX1AppleStringMap struct {
	m map[Apple]string
}

// TX1AppleStringTuple represents a key/value pair.
type TX1AppleStringTuple struct {
	Key Apple
	Val string
}

// TX1AppleStringTuples can be used as a builder for unmodifiable maps.
type TX1AppleStringTuples []TX1AppleStringTuple

// Append1 adds one item.
func (ts TX1AppleStringTuples) Append1(k Apple, v string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k, v})
}

// Append2 adds two items.
func (ts TX1AppleStringTuples) Append2(k1 Apple, v1 string, k2 Apple, v2 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1AppleStringTuples) Append3(k1 Apple, v1 string, k2 Apple, v2 string, k3 Apple, v3 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2}, TX1AppleStringTuple{k3, v3})
}

// TX1AppleStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1AppleStringMap
// constructor function.
func TX1AppleStringZip(keys ...Apple) TX1AppleStringTuples {
	ts := make(TX1AppleStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1AppleStringZip.
func (ts TX1AppleStringTuples) Values(values ...string) TX1AppleStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts TX1AppleStringTuples) ToMap() *TX1AppleStringMap {
	return NewTX1AppleStringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newTX1AppleStringMap() *TX1AppleStringMap {
	return &TX1AppleStringMap{
		m: make(map[Apple]string),
	}
}

// NewTX1AppleStringMap1 creates and returns a reference to a map containing one item.
func NewTX1AppleStringMap1(k Apple, v string) *TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1AppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1AppleStringMap(kv ...TX1AppleStringTuple) *TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TX1AppleStringMap) Keys() []Apple {
	if mm == nil {
		return nil
	}

	s := make([]Apple, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TX1AppleStringMap) Values() []string {
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TX1AppleStringMap) slice() TX1AppleStringTuples {
	if mm == nil {
		return nil
	}

	s := make(TX1AppleStringTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TX1AppleStringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TX1AppleStringMap) ToSlice() TX1AppleStringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *TX1AppleStringMap) OrderedSlice(keys []Apple) TX1AppleStringTuples {
	if mm == nil {
		return nil
	}

	s := make(TX1AppleStringTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, TX1AppleStringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *TX1AppleStringMap) Get(k Apple) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *TX1AppleStringMap) Put(k Apple, v string) *TX1AppleStringMap {
	if mm == nil {
		return NewTX1AppleStringMap1(k, v)
	}

	result := NewTX1AppleStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *TX1AppleStringMap) ContainsKey(k Apple) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TX1AppleStringMap) ContainsAllKeys(kk ...Apple) bool {
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
func (mm *TX1AppleStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TX1AppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TX1AppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TX1AppleStringMap) Foreach(f func(Apple, string)) {
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
func (mm *TX1AppleStringMap) Forall(f func(Apple, string) bool) bool {
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
func (mm *TX1AppleStringMap) Exists(p func(Apple, string) bool) bool {
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

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
func (mm *TX1AppleStringMap) Find(p func(Apple, string) bool) (TX1AppleStringTuple, bool) {
	if mm == nil {
		return TX1AppleStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return TX1AppleStringTuple{k, v}, true
		}
	}

	return TX1AppleStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *TX1AppleStringMap) Filter(p func(Apple, string) bool) *TX1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1AppleStringMap()

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
func (mm *TX1AppleStringMap) Partition(p func(Apple, string) bool) (matching *TX1AppleStringMap, others *TX1AppleStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTX1AppleStringMap()
	others = NewTX1AppleStringMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TX1StringMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1AppleStringMap) Map(f func(Apple, string) (Apple, string)) *TX1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1AppleStringMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1AppleStringMap) FlatMap(f func(Apple, string) []TX1AppleStringTuple) *TX1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1AppleStringMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm *TX1AppleStringMap) Clone() *TX1AppleStringMap {
	return mm
}
