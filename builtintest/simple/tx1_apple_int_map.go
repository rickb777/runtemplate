// A simple type derived from map[Apple]big.Int.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=big.Int
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always

package simple


import (

	"math/big"

)

// TX1AppleIntMap is the primary type that represents a map
type TX1AppleIntMap map[Apple]big.Int

// TX1AppleIntTuple represents a key/value pair.
type TX1AppleIntTuple struct {
	Key Apple
	Val big.Int
}

// TX1AppleIntTuples can be used as a builder for unmodifiable maps.
type TX1AppleIntTuples []TX1AppleIntTuple

func (ts TX1AppleIntTuples) Append1(k Apple, v big.Int) TX1AppleIntTuples {
	return append(ts, TX1AppleIntTuple{k, v})
}

func (ts TX1AppleIntTuples) Append2(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int) TX1AppleIntTuples {
	return append(ts, TX1AppleIntTuple{k1, v1}, TX1AppleIntTuple{k2, v2})
}

func (ts TX1AppleIntTuples) Append3(k1 Apple, v1 big.Int, k2 Apple, v2 big.Int, k3 Apple, v3 big.Int) TX1AppleIntTuples {
	return append(ts, TX1AppleIntTuple{k1, v1}, TX1AppleIntTuple{k2, v2}, TX1AppleIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1AppleIntMap() TX1AppleIntMap {
	return TX1AppleIntMap(make(map[Apple]big.Int))
}

// NewTX1AppleIntMap creates and returns a reference to a map containing one item.
func NewTX1AppleIntMap1(k Apple, v big.Int) TX1AppleIntMap {
	mm := newTX1AppleIntMap()
	mm[k] = v
	return mm
}

// NewTX1AppleIntMap creates and returns a reference to a map, optionally containing some items.
func NewTX1AppleIntMap(kv ...TX1AppleIntTuple) TX1AppleIntMap {
	mm := newTX1AppleIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1AppleIntMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1AppleIntMap) Values() []big.Int {
	var s []big.Int
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1AppleIntMap) ToSlice() []TX1AppleIntTuple {
	var s []TX1AppleIntTuple
	for k, v := range mm {
		s = append(s, TX1AppleIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1AppleIntMap) Get(k Apple) (big.Int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1AppleIntMap) Put(k Apple, v big.Int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1AppleIntMap) ContainsKey(k Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1AppleIntMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1AppleIntMap) Clear() {
	*mm = make(map[Apple]big.Int)
}

// Remove allows the removal of a single item from the map.
func (mm TX1AppleIntMap) Remove(k Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1AppleIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX1AppleIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1AppleIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1AppleIntMap) Forall(fn func(Apple, big.Int) bool) bool {
	for k, v := range mm {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TX1AppleIntMap) Exists(fn func(Apple, big.Int) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm TX1AppleIntMap) Filter(fn func(Apple, big.Int) bool) TX1AppleIntMap {
	result := NewTX1AppleIntMap()
	for k, v := range mm {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified
func (mm TX1AppleIntMap) Partition(fn func(Apple, big.Int) bool) (matching TX1AppleIntMap, others TX1AppleIntMap) {
	matching = NewTX1AppleIntMap()
	others = NewTX1AppleIntMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX1IntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1AppleIntMap) Map(fn func(Apple, big.Int) (Apple, big.Int)) TX1AppleIntMap {
	result := NewTX1AppleIntMap()

	for k1, v1 := range mm {
	    k2, v2 := fn(k1, v1)
	    result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1IntMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1AppleIntMap) FlatMap(fn func(Apple, big.Int) []TX1AppleIntTuple) TX1AppleIntMap {
	result := NewTX1AppleIntMap()

	for k1, v1 := range mm {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1AppleIntMap) Clone() TX1AppleIntMap {
	result := NewTX1AppleIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


