// An encapsulated map[Apple]big.Int.
// Thread-safe.
//
// Generated from immutable/map.tpl with Key=Apple Type=big.Int
// options: Comparable:<no value> Stringer:<no value> Mutable:disabled

package immutable


import (

    "math/big"

)

// TX2AppleIntMap is the primary type that represents a thread-safe map
type TX2AppleIntMap struct {
	m map[Apple]big.Int
}

// TX2AppleIntTuple represents a key/value pair.
type TX2AppleIntTuple struct {
	Key Apple
	Val big.Int
}

// TX2AppleIntTuples can be used as a builder for unmodifiable maps.
type TX2AppleIntTuples []TX2AppleIntTuple

func (ts TX2AppleIntTuples) Append1(k Apple, v big.Int) TX2AppleIntTuples {
	return append(ts, TX2AppleIntTuple{k, v})
}

func (ts TX2AppleIntTuples) Append2(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int) TX2AppleIntTuples {
	return append(ts, TX2AppleIntTuple{k1, v1}, TX2AppleIntTuple{k2, v2})
}

func (ts TX2AppleIntTuples) Append3(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int, k3 Apple, v3 big.Int) TX2AppleIntTuples {
	return append(ts, TX2AppleIntTuple{k1, v1}, TX2AppleIntTuple{k2, v2}, TX2AppleIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX2AppleIntMap() TX2AppleIntMap {
	return TX2AppleIntMap{
		m: make(map[Apple]big.Int),
	}
}

// NewTX2AppleIntMap creates and returns a reference to a map containing one item.
func NewTX2AppleIntMap1(k Apple, v big.Int) TX2AppleIntMap {
	mm := newTX2AppleIntMap()
	mm.m[k] = v
	return mm
}

// NewTX2AppleIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX2AppleIntMap(kv ...TX2AppleIntTuple) TX2AppleIntMap {
	mm := newTX2AppleIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX2AppleIntMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX2AppleIntMap) ToSlice() []TX2AppleIntTuple {
	var s []TX2AppleIntTuple
	for k, v := range mm.m {
		s = append(s, TX2AppleIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX2AppleIntMap) Get(k Apple) (big.Int, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm TX2AppleIntMap) ContainsKey(k Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX2AppleIntMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX2AppleIntMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX2AppleIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX2AppleIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX2AppleIntMap) Forall(fn func(Apple, big.Int) bool) bool {
	for k, v := range mm.m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TX2AppleIntMap) Exists(fn func(Apple, big.Int) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX2AppleIntMap) Filter(fn func(Apple, big.Int) bool) TX2AppleIntMap {
	result := NewTX2AppleIntMap()

	for k, v := range mm.m {
		if fn(k, v) {
			result.m[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm TX2AppleIntMap) Partition(fn func(Apple, big.Int) bool) (matching TX2AppleIntMap, others TX2AppleIntMap) {
	matching = NewTX2AppleIntMap()
	others = NewTX2AppleIntMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Clone returns the same map, which is immutable.
func (mm TX2AppleIntMap) Clone() TX2AppleIntMap {
	return mm
}


