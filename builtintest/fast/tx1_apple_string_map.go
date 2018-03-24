// An encapsulated map[Apple]string.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always

package fast

import (

)

// TX1AppleStringMap is the primary type that represents a thread-safe map
type TX1AppleStringMap struct {
	m map[Apple]string
}

// TX1AppleStringTuple represents a key/value pair.
type TX1AppleStringTuple struct {
	Key Apple
	Val string
}

// TX1AppleStringTuples can be used as a builder for unmodifiable maps.
type TX1AppleStringTuples []TX1AppleStringTuple

func (ts TX1AppleStringTuples) Append1(k Apple, v string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k, v})
}

func (ts TX1AppleStringTuples) Append2(k1 Apple, v1 string, k2 Apple, v2 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2})
}

func (ts TX1AppleStringTuples) Append3(k1 Apple, v1 string, k2 Apple, v2 string, k3 Apple, v3 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2}, TX1AppleStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1AppleStringMap() TX1AppleStringMap {
	return TX1AppleStringMap{
		m: make(map[Apple]string),
	}
}

// NewTX1AppleStringMap creates and returns a reference to a map containing one item.
func NewTX1AppleStringMap1(k Apple, v string) TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1AppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1AppleStringMap(kv ...TX1AppleStringTuple) TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1AppleStringMap) Keys() []Apple {

	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1AppleStringMap) Values() []string {

	var s []string
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1AppleStringMap) ToSlice() []TX1AppleStringTuple {

	var s []TX1AppleStringTuple
	for k, v := range mm.m {
		s = append(s, TX1AppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1AppleStringMap) Get(k Apple) (string, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1AppleStringMap) Put(k Apple, v string) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1AppleStringMap) ContainsKey(k Apple) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1AppleStringMap) ContainsAllKeys(kk ...Apple) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1AppleStringMap) Clear() {

	mm.m = make(map[Apple]string)
}

// Remove a single item from the map.
func (mm TX1AppleStringMap) Remove(k Apple) {

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm TX1AppleStringMap) Pop(k Apple) (string, bool) {

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1AppleStringMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1AppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1AppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TX1AppleStringMap) DropWhere(fn func(Apple, string) bool) TX1AppleStringTuples {

	removed := make(TX1AppleStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TX1AppleStringTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1AppleStringMap) Foreach(fn func(Apple, string)) {

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
func (mm TX1AppleStringMap) Forall(fn func(Apple, string) bool) bool {

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
func (mm TX1AppleStringMap) Exists(fn func(Apple, string) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (mm TX1AppleStringMap) Find(fn func(Apple, string) bool) (TX1AppleStringTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return TX1AppleStringTuple{k, v}, true
		}
	}

	return TX1AppleStringTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1AppleStringMap) Filter(fn func(Apple, string) bool) TX1AppleStringMap {
	result := NewTX1AppleStringMap()

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
func (mm TX1AppleStringMap) Partition(fn func(Apple, string) bool) (matching TX1AppleStringMap, others TX1AppleStringMap) {
	matching = NewTX1AppleStringMap()
	others = NewTX1AppleStringMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TX1StringMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1AppleStringMap) Map(fn func(Apple, string) (Apple, string)) TX1AppleStringMap {
	result := NewTX1AppleStringMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1StringMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1AppleStringMap) FlatMap(fn func(Apple, string) []TX1AppleStringTuple) TX1AppleStringMap {
	result := NewTX1AppleStringMap()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1AppleStringMap) Clone() TX1AppleStringMap {
	result := NewTX1AppleStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


