// An encapsulated map[Apple]*string.
// Note that the api uses *Apple but the map uses Apple keys.
// Not thread-safe.
//
// Generated from fast/map.tpl with Key=*Apple Type=*string
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.3.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

import (
	"fmt"
)

// TP1AppleStringMap is the primary type that represents a thread-safe map
type TP1AppleStringMap struct {
	m map[Apple]*string
}

// TP1AppleStringTuple represents a key/value pair.
type TP1AppleStringTuple struct {
	Key *Apple
	Val *string
}

// TP1AppleStringTuples can be used as a builder for unmodifiable maps.
type TP1AppleStringTuples []TP1AppleStringTuple

// Append1 adds one item.
func (ts TP1AppleStringTuples) Append1(k *Apple, v *string) TP1AppleStringTuples {
	return append(ts, TP1AppleStringTuple{k, v})
}

// Append2 adds two items.
func (ts TP1AppleStringTuples) Append2(k1 *Apple, v1 *string, k2 *Apple, v2 *string) TP1AppleStringTuples {
	return append(ts, TP1AppleStringTuple{k1, v1}, TP1AppleStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1AppleStringTuples) Append3(k1 *Apple, v1 *string, k2 *Apple, v2 *string, k3 *Apple, v3 *string) TP1AppleStringTuples {
	return append(ts, TP1AppleStringTuple{k1, v1}, TP1AppleStringTuple{k2, v2}, TP1AppleStringTuple{k3, v3})
}

// TP1AppleStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1AppleStringMap
// constructor function.
func TP1AppleStringZip(keys ...*Apple) TP1AppleStringTuples {
	ts := make(TP1AppleStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1AppleStringZip.
func (ts TP1AppleStringTuples) Values(values ...*string) TP1AppleStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTP1AppleStringMap() *TP1AppleStringMap {
	return &TP1AppleStringMap{
		m: make(map[Apple]*string),
	}
}

// NewTP1AppleStringMap1 creates and returns a reference to a map containing one item.
func NewTP1AppleStringMap1(k *Apple, v *string) *TP1AppleStringMap {
	mm := newTP1AppleStringMap()
	mm.m[*k] = v
	return mm
}

// NewTP1AppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewTP1AppleStringMap(kv ...TP1AppleStringTuple) *TP1AppleStringMap {
	mm := newTP1AppleStringMap()
	for _, t := range kv {
		mm.m[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TP1AppleStringMap) Keys() []*Apple {
	if mm == nil {
		return nil
	}

	s := make([]*Apple, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, &k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TP1AppleStringMap) Values() []*string {
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
func (mm *TP1AppleStringMap) slice() []TP1AppleStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]TP1AppleStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TP1AppleStringTuple{(&k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TP1AppleStringMap) ToSlice() []TP1AppleStringTuple {
	if mm == nil {
		return nil
	}

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TP1AppleStringMap) Get(k *Apple) (*string, bool) {

	v, found := mm.m[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TP1AppleStringMap) Put(k *Apple, v *string) bool {

	_, found := mm.m[*k]
	mm.m[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TP1AppleStringMap) ContainsKey(k *Apple) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TP1AppleStringMap) ContainsAllKeys(kk ...*Apple) bool {
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
func (mm *TP1AppleStringMap) Clear() {
	if mm != nil {

		mm.m = make(map[Apple]*string)
	}
}

// Remove a single item from the map.
func (mm *TP1AppleStringMap) Remove(k *Apple) {
	if mm != nil {

		delete(mm.m, *k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TP1AppleStringMap) Pop(k *Apple) (*string, bool) {
	if mm == nil {
		return nil, false
	}

	v, found := mm.m[*k]
	delete(mm.m, *k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TP1AppleStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TP1AppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TP1AppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TP1AppleStringMap) DropWhere(fn func(*Apple, *string) bool) TP1AppleStringTuples {

	removed := make(TP1AppleStringTuples, 0)
	for k, v := range mm.m {
		if fn(&k, v) {
			removed = append(removed, TP1AppleStringTuple{(&k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TP1AppleStringMap) Foreach(f func(*Apple, *string)) {
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
func (mm *TP1AppleStringMap) Forall(p func(*Apple, *string) bool) bool {
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
func (mm *TP1AppleStringMap) Exists(p func(*Apple, *string) bool) bool {
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
func (mm *TP1AppleStringMap) Find(p func(*Apple, *string) bool) (TP1AppleStringTuple, bool) {
	if mm == nil {
		return TP1AppleStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(&k, v) {
			return TP1AppleStringTuple{(&k), v}, true
		}
	}

	return TP1AppleStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TP1AppleStringMap) Filter(p func(*Apple, *string) bool) *TP1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1AppleStringMap()

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
func (mm *TP1AppleStringMap) Partition(p func(*Apple, *string) bool) (matching *TP1AppleStringMap, others *TP1AppleStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTP1AppleStringMap()
	others = NewTP1AppleStringMap()

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
func (mm *TP1AppleStringMap) Map(f func(*Apple, *string) (*Apple, *string)) *TP1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1AppleStringMap()

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
func (mm *TP1AppleStringMap) FlatMap(f func(*Apple, *string) []TP1AppleStringTuple) *TP1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1AppleStringMap()

	for k1, v1 := range mm.m {
		ts := f(&k1, v1)
		for _, t := range ts {
			result.m[*t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TP1AppleStringMap) Clone() *TP1AppleStringMap {
	if mm == nil {
		return nil
	}

	result := NewTP1AppleStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
