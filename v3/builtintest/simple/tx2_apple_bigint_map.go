// A simple type derived from map[Apple]big.Int.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=big.Int
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"fmt"
	"math/big"
)

// TX2AppleBigIntMap is the primary type that represents a map
type TX2AppleBigIntMap map[Apple]big.Int

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

func newTX2AppleBigIntMap() TX2AppleBigIntMap {
	return TX2AppleBigIntMap(make(map[Apple]big.Int))
}

// NewTX2AppleBigIntMap1 creates and returns a reference to a map containing one item.
func NewTX2AppleBigIntMap1(k Apple, v big.Int) TX2AppleBigIntMap {
	mm := newTX2AppleBigIntMap()
	mm[k] = v
	return mm
}

// NewTX2AppleBigIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX2AppleBigIntMap(kv ...TX2AppleBigIntTuple) TX2AppleBigIntMap {
	mm := newTX2AppleBigIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX2AppleBigIntMap) Keys() []Apple {
	s := make([]Apple, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX2AppleBigIntMap) Values() []big.Int {
	s := make([]big.Int, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TX2AppleBigIntMap) slice() []TX2AppleBigIntTuple {
	s := make([]TX2AppleBigIntTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, TX2AppleBigIntTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX2AppleBigIntMap) ToSlice() []TX2AppleBigIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm TX2AppleBigIntMap) Get(k Apple) (big.Int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX2AppleBigIntMap) Put(k Apple, v big.Int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX2AppleBigIntMap) ContainsKey(k Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX2AppleBigIntMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX2AppleBigIntMap) Clear() {
	*mm = make(map[Apple]big.Int)
}

// Remove a single item from the map.
func (mm TX2AppleBigIntMap) Remove(k Apple) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TX2AppleBigIntMap) Pop(k Apple) (big.Int, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX2AppleBigIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX2AppleBigIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX2AppleBigIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm TX2AppleBigIntMap) DropWhere(fn func(Apple, big.Int) bool) TX2AppleBigIntTuples {
	removed := make(TX2AppleBigIntTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, TX2AppleBigIntTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX2AppleBigIntMap) Foreach(f func(Apple, big.Int)) {
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
func (mm TX2AppleBigIntMap) Forall(p func(Apple, big.Int) bool) bool {
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
func (mm TX2AppleBigIntMap) Exists(p func(Apple, big.Int) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first big.Int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TX2AppleBigIntMap) Find(p func(Apple, big.Int) bool) (TX2AppleBigIntTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return TX2AppleBigIntTuple{(k), v}, true
		}
	}

	return TX2AppleBigIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TX2AppleBigIntMap) Filter(p func(Apple, big.Int) bool) TX2AppleBigIntMap {
	result := NewTX2AppleBigIntMap()
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
func (mm TX2AppleBigIntMap) Partition(p func(Apple, big.Int) bool) (matching TX2AppleBigIntMap, others TX2AppleBigIntMap) {
	matching = NewTX2AppleBigIntMap()
	others = NewTX2AppleBigIntMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX2BigIntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX2AppleBigIntMap) Map(f func(Apple, big.Int) (Apple, big.Int)) TX2AppleBigIntMap {
	result := NewTX2AppleBigIntMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX2BigIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX2AppleBigIntMap) FlatMap(f func(Apple, big.Int) []TX2AppleBigIntTuple) TX2AppleBigIntMap {
	result := NewTX2AppleBigIntMap()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX2AppleBigIntMap) Clone() TX2AppleBigIntMap {
	result := NewTX2AppleBigIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}
