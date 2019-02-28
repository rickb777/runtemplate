// A simple type derived from map[string]*Apple.
// Note that the api uses *string but the map uses string keys.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=*string Type=*Apple
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"fmt"
)

// TP1StringAppleMap is the primary type that represents a map
type TP1StringAppleMap map[string]*Apple

// TP1StringAppleTuple represents a key/value pair.
type TP1StringAppleTuple struct {
	Key *string
	Val *Apple
}

// TP1StringAppleTuples can be used as a builder for unmodifiable maps.
type TP1StringAppleTuples []TP1StringAppleTuple

// Append1 adds one item.
func (ts TP1StringAppleTuples) Append1(k *string, v *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k, v})
}

// Append2 adds two items.
func (ts TP1StringAppleTuples) Append2(k1 *string, v1 *Apple, k2 *string, v2 *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k1, v1}, TP1StringAppleTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1StringAppleTuples) Append3(k1 *string, v1 *Apple, k2 *string, v2 *Apple, k3 *string, v3 *Apple) TP1StringAppleTuples {
	return append(ts, TP1StringAppleTuple{k1, v1}, TP1StringAppleTuple{k2, v2}, TP1StringAppleTuple{k3, v3})
}

// TP1StringAppleZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1StringAppleMap
// constructor function.
func TP1StringAppleZip(keys ...*string) TP1StringAppleTuples {
	ts := make(TP1StringAppleTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1StringAppleZip.
func (ts TP1StringAppleTuples) Values(values ...*Apple) TP1StringAppleTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTP1StringAppleMap() TP1StringAppleMap {
	return TP1StringAppleMap(make(map[string]*Apple))
}

// NewTP1StringAppleMap1 creates and returns a reference to a map containing one item.
func NewTP1StringAppleMap1(k *string, v *Apple) TP1StringAppleMap {
	mm := newTP1StringAppleMap()
	mm[*k] = v
	return mm
}

// NewTP1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTP1StringAppleMap(kv ...TP1StringAppleTuple) TP1StringAppleMap {
	mm := newTP1StringAppleMap()
	for _, t := range kv {
		mm[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TP1StringAppleMap) Keys() []*string {
	s := make([]*string, 0, len(mm))
	for k := range mm {
		s = append(s, &k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TP1StringAppleMap) Values() []*Apple {
	s := make([]*Apple, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TP1StringAppleMap) slice() []TP1StringAppleTuple {
	s := make([]TP1StringAppleTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, TP1StringAppleTuple{(&k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TP1StringAppleMap) ToSlice() []TP1StringAppleTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm TP1StringAppleMap) Get(k *string) (*Apple, bool) {
	v, found := mm[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TP1StringAppleMap) Put(k *string, v *Apple) bool {
	_, found := mm[*k]
	mm[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TP1StringAppleMap) ContainsKey(k *string) bool {
	_, found := mm[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TP1StringAppleMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TP1StringAppleMap) Clear() {
	*mm = make(map[string]*Apple)
}

// Remove a single item from the map.
func (mm TP1StringAppleMap) Remove(k *string) {
	delete(mm, *k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TP1StringAppleMap) Pop(k *string) (*Apple, bool) {
	v, found := mm[*k]
	delete(mm, *k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TP1StringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TP1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TP1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm TP1StringAppleMap) DropWhere(fn func(*string, *Apple) bool) TP1StringAppleTuples {
	removed := make(TP1StringAppleTuples, 0)
	for k, v := range mm {
		if fn(&k, v) {
			removed = append(removed, TP1StringAppleTuple{(&k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TP1StringAppleMap) Foreach(f func(*string, *Apple)) {
	for k, v := range mm {
		f(&k, v)
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TP1StringAppleMap) Forall(p func(*string, *Apple) bool) bool {
	for k, v := range mm {
		if !p(&k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TP1StringAppleMap) Exists(p func(*string, *Apple) bool) bool {
	for k, v := range mm {
		if p(&k, v) {
			return true
		}
	}

	return false
}

// Find returns the first *Apple that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TP1StringAppleMap) Find(p func(*string, *Apple) bool) (TP1StringAppleTuple, bool) {
	for k, v := range mm {
		if p(&k, v) {
			return TP1StringAppleTuple{(&k), v}, true
		}
	}

	return TP1StringAppleTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TP1StringAppleMap) Filter(p func(*string, *Apple) bool) TP1StringAppleMap {
	result := NewTP1StringAppleMap()
	for k, v := range mm {
		if p(&k, v) {
			result[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm TP1StringAppleMap) Partition(p func(*string, *Apple) bool) (matching TP1StringAppleMap, others TP1StringAppleMap) {
	matching = NewTP1StringAppleMap()
	others = NewTP1StringAppleMap()
	for k, v := range mm {
		if p(&k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TP1AppleMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringAppleMap) Map(f func(*string, *Apple) (*string, *Apple)) TP1StringAppleMap {
	result := NewTP1StringAppleMap()

	for k1, v1 := range mm {
		k2, v2 := f(&k1, v1)
		result[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1AppleMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1StringAppleMap) FlatMap(f func(*string, *Apple) []TP1StringAppleTuple) TP1StringAppleMap {
	result := NewTP1StringAppleMap()

	for k1, v1 := range mm {
		ts := f(&k1, v1)
		for _, t := range ts {
			result[*t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TP1StringAppleMap) Clone() TP1StringAppleMap {
	result := NewTP1StringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}
