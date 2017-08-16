// A simple type derived from map[int]int.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> Mutable:always

package simple


import (

	"bytes"
	"fmt"
)

// SP1IntIntMap is the primary type that represents a map
type SP1IntIntMap map[*int]*int

// SP1IntIntTuple represents a key/value pair.
type SP1IntIntTuple struct {
	Key *int
	Val *int
}

// SP1IntIntTuples can be used as a builder for unmodifiable maps.
type SP1IntIntTuples []SP1IntIntTuple

func (ts SP1IntIntTuples) Append1(k *int, v *int) SP1IntIntTuples {
	return append(ts, SP1IntIntTuple{k, v})
}

func (ts SP1IntIntTuples) Append2(k1 *int, v1 *int, k2 *int, v2 *int) SP1IntIntTuples {
	return append(ts, SP1IntIntTuple{k1, v1}, SP1IntIntTuple{k2, v2})
}

func (ts SP1IntIntTuples) Append3(k1 *int, v1 *int, k2 *int, v2 *int, k3 *int, v3 *int) SP1IntIntTuples {
	return append(ts, SP1IntIntTuple{k1, v1}, SP1IntIntTuple{k2, v2}, SP1IntIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1IntIntMap() SP1IntIntMap {
	return SP1IntIntMap(make(map[*int]*int))
}

// NewSP1IntIntMap creates and returns a reference to a map containing one item.
func NewSP1IntIntMap1(k *int, v *int) SP1IntIntMap {
	mm := newSP1IntIntMap()
	mm[k] = v
	return mm
}

// NewSP1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewSP1IntIntMap(kv ...SP1IntIntTuple) SP1IntIntMap {
	mm := newSP1IntIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1IntIntMap) Keys() []*int {
	var s []*int
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1IntIntMap) Values() []*int {
	var s []*int
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1IntIntMap) ToSlice() []SP1IntIntTuple {
	var s []SP1IntIntTuple
	for k, v := range mm {
		s = append(s, SP1IntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1IntIntMap) Get(k *int) (*int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1IntIntMap) Put(k *int, v *int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1IntIntMap) ContainsKey(k *int) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1IntIntMap) ContainsAllKeys(kk ...*int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1IntIntMap) Clear() {
	*mm = make(map[*int]*int)
}

// Remove allows the removal of a single item from the map.
func (mm SP1IntIntMap) Remove(k *int) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1IntIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1IntIntMap) Forall(fn func(*int, *int) bool) bool {
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
func (mm SP1IntIntMap) Exists(fn func(*int, *int) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SP1IntIntMap) Filter(fn func(*int, *int) bool) SP1IntIntMap {
	result := NewSP1IntIntMap()
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
func (mm SP1IntIntMap) Partition(fn func(*int, *int) bool) (matching SP1IntIntMap, others SP1IntIntMap) {
	matching = NewSP1IntIntMap()
	others = NewSP1IntIntMap()
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
func (mm SP1IntIntMap) Equals(other SP1IntIntMap) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || *v1 != *v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SP1IntIntMap) Clone() SP1IntIntMap {
	result := NewSP1IntIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm SP1IntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm SP1IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm SP1IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm SP1IntIntMap) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm SP1IntIntMap) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
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

