// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from map.tpl with Key=int Type=int
// options: Comparable=true Stringer=true Mutable=true

package fast

import (

	"bytes"
	"fmt"
)

// TPIntIntMap is the primary type that represents a thread-safe map
type TPIntIntMap struct {
	m map[*int]*int
}

// TPIntIntTuple represents a key/value pair.
type TPIntIntTuple struct {
	Key *int
	Val *int
}

// TPIntIntTuples can be used as a builder for unmodifiable maps.
type TPIntIntTuples []TPIntIntTuple

func (ts TPIntIntTuples) Append1(k *int, v *int) TPIntIntTuples {
	return append(ts, TPIntIntTuple{k, v})
}

func (ts TPIntIntTuples) Append2(k1 *int, v1 *int, k2 *int, v2 *int) TPIntIntTuples {
	return append(ts, TPIntIntTuple{k1, v1}, TPIntIntTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newTPIntIntMap() TPIntIntMap {
	return TPIntIntMap{
		m: make(map[*int]*int),
	}
}

// NewTPIntIntMap creates and returns a reference to a map containing one item.
func NewTPIntIntMap1(k *int, v *int) TPIntIntMap {
	mm := newTPIntIntMap()
	mm.m[k] = v
	return mm
}

// NewTPIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTPIntIntMap(kv ...TPIntIntTuple) TPIntIntMap {
	mm := newTPIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TPIntIntMap) Keys() []*int {

	var s []*int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TPIntIntMap) ToSlice() []TPIntIntTuple {

	var s []TPIntIntTuple
	for k, v := range mm.m {
		s = append(s, TPIntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TPIntIntMap) Get(k *int) (*int, bool) {

	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm TPIntIntMap) Put(k *int, v *int) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TPIntIntMap) ContainsKey(k *int) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TPIntIntMap) ContainsAllKeys(kk ...*int) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *TPIntIntMap) Clear() {

	mm.m = make(map[*int]*int)
}

// Remove allows the removal of a single item from the map.
func (mm TPIntIntMap) Remove(k *int) {

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TPIntIntMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TPIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TPIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TPIntIntMap) Forall(fn func(*int, *int) bool) bool {

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
func (mm TPIntIntMap) Exists(fn func(*int, *int) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TPIntIntMap) Filter(fn func(*int, *int) bool) TPIntIntMap {
	result := NewTPIntIntMap()

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
func (mm TPIntIntMap) Partition(fn func(*int, *int) bool) (matching TPIntIntMap, others TPIntIntMap) {
	matching = NewTPIntIntMap()
	others = NewTPIntIntMap()

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
func (mm TPIntIntMap) Equals(other TPIntIntMap) bool {

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || *v1 != *v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TPIntIntMap) Clone() TPIntIntMap {
	result := NewTPIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TPIntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TPIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (mm TPIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (mm TPIntIntMap) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm TPIntIntMap) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
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

