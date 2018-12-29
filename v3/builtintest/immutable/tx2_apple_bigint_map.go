// An encapsulated immutable map[Apple]big.Int.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=Apple Type=big.Int
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"fmt"
	"math/big"

)

// TX2AppleBigIntMap is the primary type that represents a thread-safe map
type TX2AppleBigIntMap struct {
	m map[Apple]big.Int
}

// TX2AppleBigIntTuple represents a key/value pair.
type TX2AppleBigIntTuple struct {
	Key Apple
	Val big.Int
}

// TX2AppleBigIntTuples can be used as a builder for unmodifiable maps.
type TX2AppleBigIntTuples []TX2AppleBigIntTuple

// Append1 adds one item.
func (ts TX2AppleBigIntTuples) Append1(k Apple, v big.Int) TX2AppleBigIntTuples {
	return append(ts, TX2AppleBigIntTuple{k, v})
}

// Append2 adds two items.
func (ts TX2AppleBigIntTuples) Append2(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int) TX2AppleBigIntTuples {
	return append(ts, TX2AppleBigIntTuple{k1, v1}, TX2AppleBigIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX2AppleBigIntTuples) Append3(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int, k3 Apple, v3 big.Int) TX2AppleBigIntTuples {
	return append(ts, TX2AppleBigIntTuple{k1, v1}, TX2AppleBigIntTuple{k2, v2}, TX2AppleBigIntTuple{k3, v3})
}

// TX2AppleBigIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX2AppleBigIntMap
// constructor function.
func TX2AppleBigIntZip(keys ...Apple) TX2AppleBigIntTuples {
	ts := make(TX2AppleBigIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX2AppleBigIntZip.
func (ts TX2AppleBigIntTuples) Values(values ...big.Int) TX2AppleBigIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX2AppleBigIntMap() *TX2AppleBigIntMap {
	return &TX2AppleBigIntMap{
		m: make(map[Apple]big.Int),
	}
}

// NewTX2AppleBigIntMap1 creates and returns a reference to a map containing one item.
func NewTX2AppleBigIntMap1(k Apple, v big.Int) *TX2AppleBigIntMap {
	mm := newTX2AppleBigIntMap()
	mm.m[k] = v
	return mm
}

// NewTX2AppleBigIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX2AppleBigIntMap(kv ...TX2AppleBigIntTuple) *TX2AppleBigIntMap {
	mm := newTX2AppleBigIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TX2AppleBigIntMap) Keys() []Apple {
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
func (mm *TX2AppleBigIntMap) Values() []big.Int {
	if mm == nil {
		return nil
	}

	s := make([]big.Int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TX2AppleBigIntMap) slice() []TX2AppleBigIntTuple {
	if mm == nil {
		return nil
	}

	s := make([]TX2AppleBigIntTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TX2AppleBigIntTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TX2AppleBigIntMap) ToSlice() []TX2AppleBigIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TX2AppleBigIntMap) Get(k Apple) (big.Int, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm *TX2AppleBigIntMap) ContainsKey(k Apple) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TX2AppleBigIntMap) ContainsAllKeys(kk ...Apple) bool {
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
func (mm *TX2AppleBigIntMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TX2AppleBigIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TX2AppleBigIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TX2AppleBigIntMap) Foreach(f func(Apple, big.Int)) {
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
func (mm *TX2AppleBigIntMap) Forall(f func(Apple, big.Int) bool) bool {
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
func (mm *TX2AppleBigIntMap) Exists(p func(Apple, big.Int) bool) bool {
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

// Find returns the first big.Int that returns true for the predicate p.
// False is returned if none match.
func (mm *TX2AppleBigIntMap) Find(p func(Apple, big.Int) bool) (TX2AppleBigIntTuple, bool) {
	if mm == nil {
		return TX2AppleBigIntTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return TX2AppleBigIntTuple{k, v}, true
		}
	}

	return TX2AppleBigIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *TX2AppleBigIntMap) Filter(p func(Apple, big.Int) bool) *TX2AppleBigIntMap {
	if mm == nil {
		return nil
	}

	result := NewTX2AppleBigIntMap()

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
func (mm *TX2AppleBigIntMap) Partition(p func(Apple, big.Int) bool) (matching *TX2AppleBigIntMap, others *TX2AppleBigIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTX2AppleBigIntMap()
	others = NewTX2AppleBigIntMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TX2BigIntMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX2AppleBigIntMap) Map(f func(Apple, big.Int) (Apple, big.Int)) *TX2AppleBigIntMap {
	if mm == nil {
		return nil
	}

	result := NewTX2AppleBigIntMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX2BigIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX2AppleBigIntMap) FlatMap(f func(Apple, big.Int) []TX2AppleBigIntTuple) *TX2AppleBigIntMap {
	if mm == nil {
		return nil
	}

	result := NewTX2AppleBigIntMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm *TX2AppleBigIntMap) Clone() *TX2AppleBigIntMap {
	return mm
}
