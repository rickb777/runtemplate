// An encapsulated map[string]*string.
// Note that the api uses *string but the map uses string keys.
// Not thread-safe.
//
// Generated from fast/map.tpl with Key=*string Type=*string
// options: Comparable:true Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.2.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

import (
	"fmt"
)

// TP1StringStringMap is the primary type that represents a thread-safe map
type TP1StringStringMap struct {
	m map[string]*string
}

// TP1StringStringTuple represents a key/value pair.
type TP1StringStringTuple struct {
	Key *string
	Val *string
}

// TP1StringStringTuples can be used as a builder for unmodifiable maps.
type TP1StringStringTuples []TP1StringStringTuple

// Append1 adds one item.
func (ts TP1StringStringTuples) Append1(k *string, v *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k, v})
}

// Append2 adds two items.
func (ts TP1StringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k1, v1}, TP1StringStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1StringStringTuples) Append3(k1 *string, v1 *string, k2 *string, v2 *string, k3 *string, v3 *string) TP1StringStringTuples {
	return append(ts, TP1StringStringTuple{k1, v1}, TP1StringStringTuple{k2, v2}, TP1StringStringTuple{k3, v3})
}

// TP1StringStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1StringStringMap
// constructor function.
func TP1StringStringZip(keys ...*string) TP1StringStringTuples {
	ts := make(TP1StringStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1StringStringZip.
func (ts TP1StringStringTuples) Values(values ...*string) TP1StringStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTP1StringStringMap() *TP1StringStringMap {
	return &TP1StringStringMap{
		m: make(map[string]*string),
	}
}

// NewTP1StringStringMap1 creates and returns a reference to a map containing one item.
func NewTP1StringStringMap1(k *string, v *string) *TP1StringStringMap {
	mm := newTP1StringStringMap()
	mm.m[*k] = v
	return mm
}

// NewTP1StringStringMap creates and returns a reference to a map, optionally containing some items.
func NewTP1StringStringMap(kv ...TP1StringStringTuple) *TP1StringStringMap {
	mm := newTP1StringStringMap()
	for _, t := range kv {
		mm.m[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TP1StringStringMap) Keys() []*string {
	if mm == nil {
		return nil
	}

	s := make([]*string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, &k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TP1StringStringMap) Values() []*string {
	if mm == nil {
		return nil
	}

	s := make([]*string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TP1StringStringMap) slice() []TP1StringStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]TP1StringStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TP1StringStringTuple{(&k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TP1StringStringMap) ToSlice() []TP1StringStringTuple {
	if mm == nil {
		return nil
	}

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TP1StringStringMap) Get(k *string) (*string, bool) {

	v, found := mm.m[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TP1StringStringMap) Put(k *string, v *string) bool {

	_, found := mm.m[*k]
	mm.m[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TP1StringStringMap) ContainsKey(k *string) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TP1StringStringMap) ContainsAllKeys(kk ...*string) bool {
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

// Clear clears the entire map.
func (mm *TP1StringStringMap) Clear() {
	if mm != nil {

		mm.m = make(map[string]*string)
	}
}

// Remove a single item from the map.
func (mm *TP1StringStringMap) Remove(k *string) {
	if mm != nil {

		delete(mm.m, *k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TP1StringStringMap) Pop(k *string) (*string, bool) {
	if mm == nil {
		return nil, false
	}

	v, found := mm.m[*k]
	delete(mm.m, *k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TP1StringStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TP1StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TP1StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TP1StringStringMap) DropWhere(fn func(*string, *string) bool) TP1StringStringTuples {

	removed := make(TP1StringStringTuples, 0)
	for k, v := range mm.m {
		if fn(&k, v) {
			removed = append(removed, TP1StringStringTuple{(&k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TP1StringStringMap) Foreach(f func(*string, *string)) {
	if mm != nil {

		for k, v := range mm.m {
			f(&k, v)
		}
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TP1StringStringMap) Forall(p func(*string, *string) bool) bool {
	if mm == nil {
		return true
	}

	for k, v := range mm.m {
		if !p(&k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *TP1StringStringMap) Exists(p func(*string, *string) bool) bool {
	if mm == nil {
		return false
	}

	for k, v := range mm.m {
		if p(&k, v) {
			return true
		}
	}

	return false
}

// Find returns the first *string that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *TP1StringStringMap) Find(p func(*string, *string) bool) (TP1StringStringTuple, bool) {
	if mm == nil {
		return TP1StringStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(&k, v) {
			return TP1StringStringTuple{(&k), v}, true
		}
	}

	return TP1StringStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TP1StringStringMap) Filter(p func(*string, *string) bool) *TP1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringStringMap()

	for k, v := range mm.m {
		if p(&k, v) {
			result.m[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm *TP1StringStringMap) Partition(p func(*string, *string) bool) (matching *TP1StringStringMap, others *TP1StringStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTP1StringStringMap()
	others = NewTP1StringStringMap()

	for k, v := range mm.m {
		if p(&k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TP1StringMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1StringStringMap) Map(f func(*string, *string) (*string, *string)) *TP1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringStringMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(&k1, v1)
		result.m[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1StringStringMap) FlatMap(f func(*string, *string) []TP1StringStringTuple) *TP1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringStringMap()

	for k1, v1 := range mm.m {
		ts := f(&k1, v1)
		for _, t := range ts {
			result.m[*t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *TP1StringStringMap) Equals(other *TP1StringStringMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

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
func (mm *TP1StringStringMap) Clone() *TP1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
