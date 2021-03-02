// An encapsulated immutable map[int]int.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ImmutableIntIntMap is the primary type that represents a thread-safe map
type ImmutableIntIntMap struct {
	m map[int]int
}

// ImmutableIntIntTuple represents a key/value pair.
type ImmutableIntIntTuple struct {
	Key int
	Val int
}

// ImmutableIntIntTuples can be used as a builder for unmodifiable maps.
type ImmutableIntIntTuples []ImmutableIntIntTuple

// Append1 adds one item.
func (ts ImmutableIntIntTuples) Append1(k int, v int) ImmutableIntIntTuples {
	return append(ts, ImmutableIntIntTuple{k, v})
}

// Append2 adds two items.
func (ts ImmutableIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) ImmutableIntIntTuples {
	return append(ts, ImmutableIntIntTuple{k1, v1}, ImmutableIntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts ImmutableIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) ImmutableIntIntTuples {
	return append(ts, ImmutableIntIntTuple{k1, v1}, ImmutableIntIntTuple{k2, v2}, ImmutableIntIntTuple{k3, v3})
}

// ImmutableIntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewImmutableIntIntMap
// constructor function.
func ImmutableIntIntZip(keys ...int) ImmutableIntIntTuples {
	ts := make(ImmutableIntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with ImmutableIntIntZip.
func (ts ImmutableIntIntTuples) Values(values ...int) ImmutableIntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts ImmutableIntIntTuples) ToMap() *ImmutableIntIntMap {
	return NewImmutableIntIntMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newImmutableIntIntMap() *ImmutableIntIntMap {
	return &ImmutableIntIntMap{
		m: make(map[int]int),
	}
}

// NewImmutableIntIntMap1 creates and returns a reference to a map containing one item.
func NewImmutableIntIntMap1(k int, v int) *ImmutableIntIntMap {
	mm := newImmutableIntIntMap()
	mm.m[k] = v
	return mm
}

// NewImmutableIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewImmutableIntIntMap(kv ...ImmutableIntIntTuple) *ImmutableIntIntMap {
	mm := newImmutableIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *ImmutableIntIntMap) Keys() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *ImmutableIntIntMap) Values() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *ImmutableIntIntMap) slice() ImmutableIntIntTuples {
	if mm == nil {
		return nil
	}

	s := make(ImmutableIntIntTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, ImmutableIntIntTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *ImmutableIntIntMap) ToSlice() ImmutableIntIntTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *ImmutableIntIntMap) OrderedSlice(keys []int) ImmutableIntIntTuples {
	if mm == nil {
		return nil
	}

	s := make(ImmutableIntIntTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, ImmutableIntIntTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *ImmutableIntIntMap) Get(k int) (int, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *ImmutableIntIntMap) Put(k int, v int) *ImmutableIntIntMap {
	if mm == nil {
		return NewImmutableIntIntMap1(k, v)
	}

	result := NewImmutableIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *ImmutableIntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *ImmutableIntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *ImmutableIntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *ImmutableIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *ImmutableIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *ImmutableIntIntMap) Foreach(f func(int, int)) {
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
func (mm *ImmutableIntIntMap) Forall(f func(int, int) bool) bool {
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
func (mm *ImmutableIntIntMap) Exists(p func(int, int) bool) bool {
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

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
func (mm *ImmutableIntIntMap) Find(p func(int, int) bool) (ImmutableIntIntTuple, bool) {
	if mm == nil {
		return ImmutableIntIntTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return ImmutableIntIntTuple{k, v}, true
		}
	}

	return ImmutableIntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *ImmutableIntIntMap) Filter(p func(int, int) bool) *ImmutableIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableIntIntMap()

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
func (mm *ImmutableIntIntMap) Partition(p func(int, int) bool) (matching *ImmutableIntIntMap, others *ImmutableIntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewImmutableIntIntMap()
	others = NewImmutableIntIntMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new ImmutableIntMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *ImmutableIntIntMap) Map(f func(int, int) (int, int)) *ImmutableIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableIntIntMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new ImmutableIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *ImmutableIntIntMap) FlatMap(f func(int, int) []ImmutableIntIntTuple) *ImmutableIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewImmutableIntIntMap()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *ImmutableIntIntMap) Equals(other *ImmutableIntIntMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns the same map, which is immutable.
func (mm *ImmutableIntIntMap) Clone() *ImmutableIntIntMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *ImmutableIntIntMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *ImmutableIntIntMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *ImmutableIntIntMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *ImmutableIntIntMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

func (ts ImmutableIntIntTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts ImmutableIntIntTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts ImmutableIntIntTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts ImmutableIntIntTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	sep := before
	for _, t := range ts {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", t.Key, equals, t.Val)
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t ImmutableIntIntTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t ImmutableIntIntTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
