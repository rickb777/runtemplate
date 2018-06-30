// An encapsulated map[string]string.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always

package fast

import (

)

// TP1StringStringMap is the primary type that represents a thread-safe map
type TP1StringStringMap struct {
	m map[*string]*string
}

// TP1StringStringTuple represents a key/value pair.
type TP1StringStringTuple struct {
	Key *string
	Val *string
}

// TP1StringStringTuples can be used as a builder for unmodifiable maps.
type TP1StringStringTuples []TP1StringStringTuple

func (ts TP1StringStringTuples) Append1(k *string, v *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k, v})
}

func (ts TP1StringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k1, v1}, TP1StringStringTuple{k2, v2})
}

func (ts TP1StringStringTuples) Append3(k1 *string, v1 *string, k2 *string, v2 *string, k3 *string, v3 *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k1, v1}, TP1StringStringTuple{k2, v2}, TP1StringStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTP1StringStringMap() TP1StringStringMap {
	return TP1StringStringMap{
		m: make(map[*string]*string),
	}
}

// NewTP1StringStringMap creates and returns a reference to a map containing one item.
func NewTP1StringStringMap1(k *string, v *string) TP1StringStringMap {
	mm := newTP1StringStringMap()
	mm.m[k] = v
	return mm
}

// NewTP1StringStringMap creates and returns a reference to a map, optionally containing some items.
func NewTP1StringStringMap(kv ...TP1StringStringTuple) TP1StringStringMap {
	mm := newTP1StringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TP1StringStringMap) Keys() []*string {

	var s []*string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TP1StringStringMap) Values() []*string {

	var s []*string
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TP1StringStringMap) ToSlice() []TP1StringStringTuple {

	var s []TP1StringStringTuple
	for k, v := range mm.m {
		s = append(s, TP1StringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TP1StringStringMap) Get(k *string) (*string, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TP1StringStringMap) Put(k *string, v *string) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TP1StringStringMap) ContainsKey(k *string) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TP1StringStringMap) ContainsAllKeys(kk ...*string) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TP1StringStringMap) Clear() {

	mm.m = make(map[*string]*string)
}

// Remove a single item from the map.
func (mm TP1StringStringMap) Remove(k *string) {

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm TP1StringStringMap) Pop(k *string) (*string, bool) {

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TP1StringStringMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TP1StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TP1StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TP1StringStringMap) DropWhere(fn func(*string, *string) bool) TP1StringStringTuples {

	removed := make(TP1StringStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TP1StringStringTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TP1StringStringMap) Foreach(fn func(*string, *string)) {

	for k, v := range mm.m {
		fn(k, v)
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TP1StringStringMap) Forall(fn func(*string, *string) bool) bool {

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
func (mm TP1StringStringMap) Exists(fn func(*string, *string) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (mm TP1StringStringMap) Find(fn func(*string, *string) bool) (TP1StringStringTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return TP1StringStringTuple{k, v}, true
		}
	}

	return TP1StringStringTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TP1StringStringMap) Filter(fn func(*string, *string) bool) TP1StringStringMap {
	result := NewTP1StringStringMap()

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
func (mm TP1StringStringMap) Partition(fn func(*string, *string) bool) (matching TP1StringStringMap, others TP1StringStringMap) {
	matching = NewTP1StringStringMap()
	others = NewTP1StringStringMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TP1StringMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringStringMap) Map(fn func(*string, *string) (*string, *string)) TP1StringStringMap {
	result := NewTP1StringStringMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TP1StringMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringStringMap) FlatMap(fn func(*string, *string) []TP1StringStringTuple) TP1StringStringMap {
	result := NewTP1StringStringMap()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TP1StringStringMap) Equals(other TP1StringStringMap) bool {

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
func (mm TP1StringStringMap) Clone() TP1StringStringMap {
	result := NewTP1StringStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


