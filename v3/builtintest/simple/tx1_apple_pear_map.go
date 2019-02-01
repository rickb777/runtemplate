// A simple type derived from map[Apple]Pear.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=Pear
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.2.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"fmt"
)

// TX1ApplePearMap is the primary type that represents a map
type TX1ApplePearMap map[Apple]Pear

// TX1ApplePearTuple represents a key/value pair.
type TX1ApplePearTuple struct {
	Key Apple
	Val Pear
}

// TX1ApplePearTuples can be used as a builder for unmodifiable maps.
type TX1ApplePearTuples []TX1ApplePearTuple

// Append1 adds one item.
func (ts TX1ApplePearTuples) Append1(k Apple, v Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k, v})
}

// Append2 adds two items.
func (ts TX1ApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1ApplePearTuples) Append3(k1 Apple, v1 Pear, k2 Apple, v2 Pear, k3 Apple, v3 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2}, TX1ApplePearTuple{k3, v3})
}

// TX1ApplePearZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1ApplePearMap
// constructor function.
func TX1ApplePearZip(keys ...Apple) TX1ApplePearTuples {
	ts := make(TX1ApplePearTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1ApplePearZip.
func (ts TX1ApplePearTuples) Values(values ...Pear) TX1ApplePearTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX1ApplePearMap() TX1ApplePearMap {
	return TX1ApplePearMap(make(map[Apple]Pear))
}

// NewTX1ApplePearMap1 creates and returns a reference to a map containing one item.
func NewTX1ApplePearMap1(k Apple, v Pear) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	mm[k] = v
	return mm
}

// NewTX1ApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTX1ApplePearMap(kv ...TX1ApplePearTuple) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1ApplePearMap) Keys() []Apple {
	s := make([]Apple, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1ApplePearMap) Values() []Pear {
	s := make([]Pear, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TX1ApplePearMap) slice() []TX1ApplePearTuple {
	s := make([]TX1ApplePearTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, TX1ApplePearTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1ApplePearMap) ToSlice() []TX1ApplePearTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm TX1ApplePearMap) Get(k Apple) (Pear, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1ApplePearMap) Put(k Apple, v Pear) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1ApplePearMap) ContainsKey(k Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1ApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1ApplePearMap) Clear() {
	*mm = make(map[Apple]Pear)
}

// Remove a single item from the map.
func (mm TX1ApplePearMap) Remove(k Apple) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TX1ApplePearMap) Pop(k Apple) (Pear, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1ApplePearMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX1ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm TX1ApplePearMap) DropWhere(fn func(Apple, Pear) bool) TX1ApplePearTuples {
	removed := make(TX1ApplePearTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, TX1ApplePearTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1ApplePearMap) Foreach(f func(Apple, Pear)) {
	for k, v := range mm {
		f(k, v)
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1ApplePearMap) Forall(p func(Apple, Pear) bool) bool {
	for k, v := range mm {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TX1ApplePearMap) Exists(p func(Apple, Pear) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first Pear that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TX1ApplePearMap) Find(p func(Apple, Pear) bool) (TX1ApplePearTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return TX1ApplePearTuple{(k), v}, true
		}
	}

	return TX1ApplePearTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TX1ApplePearMap) Filter(p func(Apple, Pear) bool) TX1ApplePearMap {
	result := NewTX1ApplePearMap()
	for k, v := range mm {
		if p(k, v) {
			result[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm TX1ApplePearMap) Partition(p func(Apple, Pear) bool) (matching TX1ApplePearMap, others TX1ApplePearMap) {
	matching = NewTX1ApplePearMap()
	others = NewTX1ApplePearMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX1PearMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1ApplePearMap) Map(f func(Apple, Pear) (Apple, Pear)) TX1ApplePearMap {
	result := NewTX1ApplePearMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1PearMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1ApplePearMap) FlatMap(f func(Apple, Pear) []TX1ApplePearTuple) TX1ApplePearMap {
	result := NewTX1ApplePearMap()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1ApplePearMap) Clone() TX1ApplePearMap {
	result := NewTX1ApplePearMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm TX1ApplePearMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1ApplePearMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1ApplePearMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TX1ApplePearMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1ApplePearMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}
