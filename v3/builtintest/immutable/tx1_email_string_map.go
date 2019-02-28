// An encapsulated immutable map[Email]string.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=Email Type=string
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"fmt"
)

// TX1EmailStringMap is the primary type that represents a thread-safe map
type TX1EmailStringMap struct {
	m map[Email]string
}

// TX1EmailStringTuple represents a key/value pair.
type TX1EmailStringTuple struct {
	Key Email
	Val string
}

// TX1EmailStringTuples can be used as a builder for unmodifiable maps.
type TX1EmailStringTuples []TX1EmailStringTuple

// Append1 adds one item.
func (ts TX1EmailStringTuples) Append1(k Email, v string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k, v})
}

// Append2 adds two items.
func (ts TX1EmailStringTuples) Append2(k1 Email, v1 string, k2 Email, v2 string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k1, v1}, TX1EmailStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1EmailStringTuples) Append3(k1 Email, v1 string, k2 Email, v2 string, k3 Email, v3 string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k1, v1}, TX1EmailStringTuple{k2, v2}, TX1EmailStringTuple{k3, v3})
}

// TX1EmailStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1EmailStringMap
// constructor function.
func TX1EmailStringZip(keys ...Email) TX1EmailStringTuples {
	ts := make(TX1EmailStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1EmailStringZip.
func (ts TX1EmailStringTuples) Values(values ...string) TX1EmailStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX1EmailStringMap() *TX1EmailStringMap {
	return &TX1EmailStringMap{
		m: make(map[Email]string),
	}
}

// NewTX1EmailStringMap1 creates and returns a reference to a map containing one item.
func NewTX1EmailStringMap1(k Email, v string) *TX1EmailStringMap {
	mm := newTX1EmailStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1EmailStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1EmailStringMap(kv ...TX1EmailStringTuple) *TX1EmailStringMap {
	mm := newTX1EmailStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TX1EmailStringMap) Keys() []Email {
	if mm == nil {
		return nil
	}

	s := make([]Email, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TX1EmailStringMap) Values() []string {
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TX1EmailStringMap) slice() []TX1EmailStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]TX1EmailStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TX1EmailStringTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TX1EmailStringMap) ToSlice() []TX1EmailStringTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TX1EmailStringMap) Get(k Email) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *TX1EmailStringMap) Put(k Email, v string) *TX1EmailStringMap {
	if mm == nil {
		return NewTX1EmailStringMap1(k, v)
	}

	result := NewTX1EmailStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *TX1EmailStringMap) ContainsKey(k Email) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TX1EmailStringMap) ContainsAllKeys(kk ...Email) bool {
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

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TX1EmailStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TX1EmailStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TX1EmailStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TX1EmailStringMap) Foreach(f func(Email, string)) {
	if mm != nil {
		for k, v := range mm.m {
			f(k, v)
		}
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *TX1EmailStringMap) Forall(f func(Email, string) bool) bool {
	if mm == nil {
		return true
	}

	for k, v := range mm.m {
		if !f(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *TX1EmailStringMap) Exists(p func(Email, string) bool) bool {
	if mm == nil {
		return false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
func (mm *TX1EmailStringMap) Find(p func(Email, string) bool) (TX1EmailStringTuple, bool) {
	if mm == nil {
		return TX1EmailStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return TX1EmailStringTuple{k, v}, true
		}
	}

	return TX1EmailStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *TX1EmailStringMap) Filter(p func(Email, string) bool) *TX1EmailStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1EmailStringMap()

	for k, v := range mm.m {
		if p(k, v) {
			result.m[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm *TX1EmailStringMap) Partition(p func(Email, string) bool) (matching *TX1EmailStringMap, others *TX1EmailStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTX1EmailStringMap()
	others = NewTX1EmailStringMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TX1StringMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1EmailStringMap) Map(f func(Email, string) (Email, string)) *TX1EmailStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1EmailStringMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1EmailStringMap) FlatMap(f func(Email, string) []TX1EmailStringTuple) *TX1EmailStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1EmailStringMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns the same map, which is immutable.
func (mm *TX1EmailStringMap) Clone() *TX1EmailStringMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *TX1EmailStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *TX1EmailStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *TX1EmailStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
// The map entries are sorted by their keys.
func (mm *TX1EmailStringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *TX1EmailStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
