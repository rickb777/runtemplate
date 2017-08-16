// A simple type derived from map[string]big.Int.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=big.Int
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> Mutable:always

package simple


import (

	"math/big"

)

// SP1StringIntMap is the primary type that represents a map
type SP1StringIntMap map[*string]*big.Int

// SP1StringIntTuple represents a key/value pair.
type SP1StringIntTuple struct {
	Key *string
	Val *big.Int
}

// SP1StringIntTuples can be used as a builder for unmodifiable maps.
type SP1StringIntTuples []SP1StringIntTuple

func (ts SP1StringIntTuples) Append1(k *string, v *big.Int) SP1StringIntTuples {
	return append(ts, SP1StringIntTuple{k, v})
}

func (ts SP1StringIntTuples) Append2(k1 *string, v1 *big.Int, k2 *string, v2 *big.Int) SP1StringIntTuples {
	return append(ts, SP1StringIntTuple{k1, v1}, SP1StringIntTuple{k2, v2})
}

func (ts SP1StringIntTuples) Append3(k1 *string, v1 *big.Int, k2 *string, v2 *big.Int, k3 *string, v3 *big.Int) SP1StringIntTuples {
	return append(ts, SP1StringIntTuple{k1, v1}, SP1StringIntTuple{k2, v2}, SP1StringIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1StringIntMap() SP1StringIntMap {
	return SP1StringIntMap(make(map[*string]*big.Int))
}

// NewSP1StringIntMap creates and returns a reference to a map containing one item.
func NewSP1StringIntMap1(k *string, v *big.Int) SP1StringIntMap {
	mm := newSP1StringIntMap()
	mm[k] = v
	return mm
}

// NewSP1StringIntMap creates and returns a reference to a map, optionally containing some items.
func NewSP1StringIntMap(kv ...SP1StringIntTuple) SP1StringIntMap {
	mm := newSP1StringIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1StringIntMap) Keys() []*string {
	var s []*string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1StringIntMap) Values() []*big.Int {
	var s []*big.Int
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1StringIntMap) ToSlice() []SP1StringIntTuple {
	var s []SP1StringIntTuple
	for k, v := range mm {
		s = append(s, SP1StringIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1StringIntMap) Get(k *string) (*big.Int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1StringIntMap) Put(k *string, v *big.Int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1StringIntMap) ContainsKey(k *string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1StringIntMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1StringIntMap) Clear() {
	*mm = make(map[*string]*big.Int)
}

// Remove allows the removal of a single item from the map.
func (mm SP1StringIntMap) Remove(k *string) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1StringIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1StringIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1StringIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1StringIntMap) Forall(fn func(*string, *big.Int) bool) bool {
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
func (mm SP1StringIntMap) Exists(fn func(*string, *big.Int) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SP1StringIntMap) Filter(fn func(*string, *big.Int) bool) SP1StringIntMap {
	result := NewSP1StringIntMap()
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
func (mm SP1StringIntMap) Partition(fn func(*string, *big.Int) bool) (matching SP1StringIntMap, others SP1StringIntMap) {
	matching = NewSP1StringIntMap()
	others = NewSP1StringIntMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SP1StringIntMap) Clone() SP1StringIntMap {
	result := NewSP1StringIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


