// A simple type derived from map[int]int.
// Not thread-safe.
//
// Generated from map.tpl with Key=int Type=int
// options: Comparable=true Stringer=true Mutable=always

package simple


import (
	"bytes"
	"fmt"
)

// SXIntIntMap is the primary type that represents a map
type SXIntIntMap map[int]int

// SXIntIntTuple represents a key/value pair.
type SXIntIntTuple struct {
	Key int
	Val int
}

// SXIntIntTuples can be used as a builder for unmodifiable maps.
type SXIntIntTuples []SXIntIntTuple

func (ts SXIntIntTuples) Append1(k int, v int) SXIntIntTuples {
	return append(ts, SXIntIntTuple{k, v})
}

func (ts SXIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) SXIntIntTuples {
	return append(ts, SXIntIntTuple{k1, v1}, SXIntIntTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSXIntIntMap() SXIntIntMap {
	return SXIntIntMap(make(map[int]int))
}

// NewSXIntIntMap creates and returns a reference to a map containing one item.
func NewSXIntIntMap1(k int, v int) SXIntIntMap {
	mm := newSXIntIntMap()
	mm[k] = v
	return mm
}

// NewSXIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewSXIntIntMap(kv ...SXIntIntTuple) SXIntIntMap {
	mm := newSXIntIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SXIntIntMap) Keys() []int {
	var s []int
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SXIntIntMap) ToSlice() []SXIntIntTuple {
	var s []SXIntIntTuple
	for k, v := range mm {
		s = append(s, SXIntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SXIntIntMap) Get(k int) (int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SXIntIntMap) Put(k int, v int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SXIntIntMap) ContainsKey(k int) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SXIntIntMap) ContainsAllKeys(kk ...int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Remove allows the removal of a single item from the map.
func (mm SXIntIntMap) Remove(k int) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SXIntIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SXIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SXIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SXIntIntMap) Forall(fn func(int, int) bool) bool {
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
func (mm SXIntIntMap) Exists(fn func(int, int) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SXIntIntMap) Filter(fn func(int, int) bool) SXIntIntMap {
	result := NewSXIntIntMap()
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
func (mm SXIntIntMap) Partition(fn func(int, int) bool) (matching SXIntIntMap, others SXIntIntMap) {
	matching = NewSXIntIntMap()
	others = NewSXIntIntMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm SXIntIntMap) Equals(other SXIntIntMap) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SXIntIntMap) Clone() SXIntIntMap {
	result := NewSXIntIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm SXIntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm SXIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (mm SXIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (mm SXIntIntMap) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm SXIntIntMap) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
