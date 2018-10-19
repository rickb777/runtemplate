// An encapsulated map[string]Apple.
// Thread-safe.
//
// Generated from immutable/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled

package examples


import (

	"bytes"
	"fmt"
)

// ImmutableStringAppleMap is the primary type that represents a thread-safe map
type ImmutableStringAppleMap struct {
	m map[string]Apple
}

// ImmutableStringAppleTuple represents a key/value pair.
type ImmutableStringAppleTuple struct {
	Key string
	Val Apple
}

// ImmutableStringAppleTuples can be used as a builder for unmodifiable maps.
type ImmutableStringAppleTuples []ImmutableStringAppleTuple

func (ts ImmutableStringAppleTuples) Append1(k string, v Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k, v})
}

func (ts ImmutableStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k1, v1}, ImmutableStringAppleTuple{k2, v2})
}

func (ts ImmutableStringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) ImmutableStringAppleTuples {
	return append(ts, ImmutableStringAppleTuple{k1, v1}, ImmutableStringAppleTuple{k2, v2}, ImmutableStringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newImmutableStringAppleMap() ImmutableStringAppleMap {
	return ImmutableStringAppleMap{
		m: make(map[string]Apple),
	}
}

// NewImmutableStringAppleMap creates and returns a reference to a map containing one item.
func NewImmutableStringAppleMap1(k string, v Apple) ImmutableStringAppleMap {
	mm := newImmutableStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewImmutableStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewImmutableStringAppleMap(kv ...ImmutableStringAppleTuple) ImmutableStringAppleMap {
	mm := newImmutableStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm ImmutableStringAppleMap) Keys() []string {
	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm ImmutableStringAppleMap) Values() []Apple {

	var s []Apple
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm ImmutableStringAppleMap) ToSlice() []ImmutableStringAppleTuple {
	var s []ImmutableStringAppleTuple
	for k, v := range mm.m {
		s = append(s, ImmutableStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm ImmutableStringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm ImmutableStringAppleMap) ContainsKey(k string) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm ImmutableStringAppleMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm ImmutableStringAppleMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm ImmutableStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm ImmutableStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm ImmutableStringAppleMap) Forall(fn func(string, Apple) bool) bool {
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
func (mm ImmutableStringAppleMap) Exists(fn func(string, Apple) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (mm ImmutableStringAppleMap) Find(fn func(string, Apple) bool) (ImmutableStringAppleTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return ImmutableStringAppleTuple{k, v}, true
		}
	}

	return ImmutableStringAppleTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm ImmutableStringAppleMap) Filter(fn func(string, Apple) bool) ImmutableStringAppleMap {
	result := NewImmutableStringAppleMap()

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
func (mm ImmutableStringAppleMap) Partition(fn func(string, Apple) bool) (matching ImmutableStringAppleMap, others ImmutableStringAppleMap) {
	matching = NewImmutableStringAppleMap()
	others = NewImmutableStringAppleMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new ImmutableAppleMap by transforming every element with a function fn.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm ImmutableStringAppleMap) Map(fn func(string, Apple) (string, Apple)) ImmutableStringAppleMap {
	result := NewImmutableStringAppleMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new ImmutableAppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm ImmutableStringAppleMap) FlatMap(fn func(string, Apple) []ImmutableStringAppleTuple) ImmutableStringAppleMap {
	result := NewImmutableStringAppleMap()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm ImmutableStringAppleMap) Clone() ImmutableStringAppleMap {
	return mm
}


//-------------------------------------------------------------------------------------------------

func (mm ImmutableStringAppleMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm ImmutableStringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm ImmutableStringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm ImmutableStringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm ImmutableStringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}
