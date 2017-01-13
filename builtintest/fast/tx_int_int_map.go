// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=int Type=int
// options: Comparable=true Stringer=true Mutable=always

package fast

import (

	"bytes"
	"fmt"
)

// TXIntIntMap is the primary type that represents a thread-safe map
type TXIntIntMap struct {
	m map[int]int
}

// TXIntIntTuple represents a key/value pair.
type TXIntIntTuple struct {
	Key int
	Val int
}

// TXIntIntTuples can be used as a builder for unmodifiable maps.
type TXIntIntTuples []TXIntIntTuple

func (ts TXIntIntTuples) Append1(k int, v int) TXIntIntTuples {
	return append(ts, TXIntIntTuple{k, v})
}

func (ts TXIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) TXIntIntTuples {
	return append(ts, TXIntIntTuple{k1, v1}, TXIntIntTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newTXIntIntMap() TXIntIntMap {
	return TXIntIntMap{
		m: make(map[int]int),
	}
}

// NewTXIntIntMap creates and returns a reference to a map containing one item.
func NewTXIntIntMap1(k int, v int) TXIntIntMap {
	mm := newTXIntIntMap()
	mm.m[k] = v
	return mm
}

// NewTXIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTXIntIntMap(kv ...TXIntIntTuple) TXIntIntMap {
	mm := newTXIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TXIntIntMap) Keys() []int {

	var s []int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TXIntIntMap) ToSlice() []TXIntIntTuple {

	var s []TXIntIntTuple
	for k, v := range mm.m {
		s = append(s, TXIntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TXIntIntMap) Get(k int) (int, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TXIntIntMap) Put(k int, v int) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TXIntIntMap) ContainsKey(k int) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TXIntIntMap) ContainsAllKeys(kk ...int) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TXIntIntMap) Clear() {

	mm.m = make(map[int]int)
}

// Remove allows the removal of a single item from the map.
func (mm TXIntIntMap) Remove(k int) {

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TXIntIntMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TXIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TXIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TXIntIntMap) Forall(fn func(int, int) bool) bool {

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
func (mm TXIntIntMap) Exists(fn func(int, int) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TXIntIntMap) Filter(fn func(int, int) bool) TXIntIntMap {
	result := NewTXIntIntMap()

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
func (mm TXIntIntMap) Partition(fn func(int, int) bool) (matching TXIntIntMap, others TXIntIntMap) {
	matching = NewTXIntIntMap()
	others = NewTXIntIntMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TXIntIntMap) Equals(other TXIntIntMap) bool {

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
func (mm TXIntIntMap) Clone() TXIntIntMap {
	result := NewTXIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TXIntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TXIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (mm TXIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (mm TXIntIntMap) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm TXIntIntMap) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

