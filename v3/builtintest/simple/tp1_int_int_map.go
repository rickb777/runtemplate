// A simple type derived from map[int]*int.
// Note that the api uses *int but the map uses int keys.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=*int Type=*int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TP1IntIntMap is the primary type that represents a map
type TP1IntIntMap map[int]*int

// TP1IntIntTuple represents a key/value pair.
type TP1IntIntTuple struct {
	Key *int
	Val *int
}

// TP1IntIntTuples can be used as a builder for unmodifiable maps.
type TP1IntIntTuples []TP1IntIntTuple

// Append1 adds one item.
func (ts TP1IntIntTuples) Append1(k *int, v *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k, v})
}

// Append2 adds two items.
func (ts TP1IntIntTuples) Append2(k1 *int, v1 *int, k2 *int, v2 *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k1, v1}, TP1IntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1IntIntTuples) Append3(k1 *int, v1 *int, k2 *int, v2 *int, k3 *int, v3 *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k1, v1}, TP1IntIntTuple{k2, v2}, TP1IntIntTuple{k3, v3})
}

// TP1IntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1IntIntMap
// constructor function.
func TP1IntIntZip(keys ...*int) TP1IntIntTuples {
	ts := make(TP1IntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1IntIntZip.
func (ts TP1IntIntTuples) Values(values ...*int) TP1IntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts TP1IntIntTuples) ToMap() TP1IntIntMap {
	return NewTP1IntIntMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newTP1IntIntMap() TP1IntIntMap {
	return TP1IntIntMap(make(map[int]*int))
}

// NewTP1IntIntMap1 creates and returns a reference to a map containing one item.
func NewTP1IntIntMap1(k *int, v *int) TP1IntIntMap {
	mm := newTP1IntIntMap()
	mm[*k] = v
	return mm
}

// NewTP1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTP1IntIntMap(kv ...TP1IntIntTuple) TP1IntIntMap {
	mm := newTP1IntIntMap()
	for _, t := range kv {
		mm[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TP1IntIntMap) Keys() []*int {
	if mm == nil {
		return nil
	}

	s := make([]*int, 0, len(mm))
	for k := range mm {
		s = append(s, &k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TP1IntIntMap) Values() []*int {
	if mm == nil {
		return nil
	}

	s := make([]*int, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm TP1IntIntMap) slice() TP1IntIntTuples {
	s := make(TP1IntIntTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, TP1IntIntTuple{(&k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm TP1IntIntMap) ToSlice() TP1IntIntTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm TP1IntIntMap) OrderedSlice(keys []*int) TP1IntIntTuples {
	s := make(TP1IntIntTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[*k]
		if found {
			s = append(s, TP1IntIntTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TP1IntIntMap) Get(k *int) (*int, bool) {
	v, found := mm[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TP1IntIntMap) Put(k *int, v *int) bool {
	_, found := mm[*k]
	mm[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TP1IntIntMap) ContainsKey(k *int) bool {
	_, found := mm[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TP1IntIntMap) ContainsAllKeys(kk ...*int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TP1IntIntMap) Clear() {
	*mm = make(map[int]*int)
}

// Remove a single item from the map.
func (mm TP1IntIntMap) Remove(k *int) {
	delete(mm, *k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm TP1IntIntMap) Pop(k *int) (*int, bool) {
	v, found := mm[*k]
	delete(mm, *k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TP1IntIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TP1IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TP1IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm TP1IntIntMap) DropWhere(fn func(*int, *int) bool) TP1IntIntTuples {
	removed := make(TP1IntIntTuples, 0)
	for k, v := range mm {
		if fn(&k, v) {
			removed = append(removed, TP1IntIntTuple{(&k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TP1IntIntMap) Foreach(f func(*int, *int)) {
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
func (mm TP1IntIntMap) Forall(p func(*int, *int) bool) bool {
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
func (mm TP1IntIntMap) Exists(p func(*int, *int) bool) bool {
	for k, v := range mm {
		if p(&k, v) {
			return true
		}
	}

	return false
}

// Find returns the first *int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm TP1IntIntMap) Find(p func(*int, *int) bool) (TP1IntIntTuple, bool) {
	for k, v := range mm {
		if p(&k, v) {
			return TP1IntIntTuple{(&k), v}, true
		}
	}

	return TP1IntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm TP1IntIntMap) Filter(p func(*int, *int) bool) TP1IntIntMap {
	result := NewTP1IntIntMap()
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
func (mm TP1IntIntMap) Partition(p func(*int, *int) bool) (matching TP1IntIntMap, others TP1IntIntMap) {
	matching = NewTP1IntIntMap()
	others = NewTP1IntIntMap()
	for k, v := range mm {
		if p(&k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TP1IntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1IntIntMap) Map(f func(*int, *int) (*int, *int)) TP1IntIntMap {
	result := NewTP1IntIntMap()

	for k1, v1 := range mm {
		k2, v2 := f(&k1, v1)
		result[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1IntIntMap) FlatMap(f func(*int, *int) []TP1IntIntTuple) TP1IntIntMap {
	result := NewTP1IntIntMap()

	for k1, v1 := range mm {
		ts := f(&k1, v1)
		for _, t := range ts {
			result[*t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TP1IntIntMap) Equals(other TP1IntIntMap) bool {
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
func (mm TP1IntIntMap) Clone() TP1IntIntMap {
	result := NewTP1IntIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm TP1IntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TP1IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TP1IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TP1IntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TP1IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

//-------------------------------------------------------------------------------------------------

func (ts TP1IntIntTuples) String() string {
	return ts.MkString3("[", ", ", "]")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts TP1IntIntTuples) MkString(sep string) string {
	return ts.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts TP1IntIntTuples) MkString3(before, between, after string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString3Bytes(before, between, after).String()
}

func (ts TP1IntIntTuples) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	for _, t := range ts {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", t.Key, t.Val))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t TP1IntIntTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t TP1IntIntTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
