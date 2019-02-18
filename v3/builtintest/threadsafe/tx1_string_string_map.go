// An encapsulated map[string]string.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.3.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

// TX1StringStringMap is the primary type that represents a thread-safe map
type TX1StringStringMap struct {
	s *sync.RWMutex
	m map[string]string
}

// TX1StringStringTuple represents a key/value pair.
type TX1StringStringTuple struct {
	Key string
	Val string
}

// TX1StringStringTuples can be used as a builder for unmodifiable maps.
type TX1StringStringTuples []TX1StringStringTuple

// Append1 adds one item.
func (ts TX1StringStringTuples) Append1(k string, v string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k, v})
}

// Append2 adds two items.
func (ts TX1StringStringTuples) Append2(k1 string, v1 string, k2 string, v2 string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k1, v1}, TX1StringStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts TX1StringStringTuples) Append3(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k1, v1}, TX1StringStringTuple{k2, v2}, TX1StringStringTuple{k3, v3})
}

// TX1StringStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTX1StringStringMap
// constructor function.
func TX1StringStringZip(keys ...string) TX1StringStringTuples {
	ts := make(TX1StringStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TX1StringStringZip.
func (ts TX1StringStringTuples) Values(values ...string) TX1StringStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTX1StringStringMap() *TX1StringStringMap {
	return &TX1StringStringMap{
		s: &sync.RWMutex{},
		m: make(map[string]string),
	}
}

// NewTX1StringStringMap1 creates and returns a reference to a map containing one item.
func NewTX1StringStringMap1(k string, v string) *TX1StringStringMap {
	mm := newTX1StringStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1StringStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1StringStringMap(kv ...TX1StringStringTuple) *TX1StringStringMap {
	mm := newTX1StringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TX1StringStringMap) Keys() []string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TX1StringStringMap) Values() []string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TX1StringStringMap) slice() []TX1StringStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]TX1StringStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TX1StringStringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TX1StringStringMap) ToSlice() []TX1StringStringTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TX1StringStringMap) Get(k string) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TX1StringStringMap) Put(k string, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TX1StringStringMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TX1StringStringMap) ContainsAllKeys(kk ...string) bool {
	if mm == nil {
		return len(kk) == 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1StringStringMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[string]string)
	}
}

// Remove a single item from the map.
func (mm *TX1StringStringMap) Remove(k string) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TX1StringStringMap) Pop(k string) (string, bool) {
	if mm == nil {
		return "", false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TX1StringStringMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TX1StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TX1StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TX1StringStringMap) DropWhere(fn func(string, string) bool) TX1StringStringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TX1StringStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TX1StringStringTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TX1StringStringMap) Foreach(f func(string, string)) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

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
func (mm *TX1StringStringMap) Forall(p func(string, string) bool) bool {
	if mm == nil {
		return true
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *TX1StringStringMap) Exists(p func(string, string) bool) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *TX1StringStringMap) Find(p func(string, string) bool) (TX1StringStringTuple, bool) {
	if mm == nil {
		return TX1StringStringTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return TX1StringStringTuple{(k), v}, true
		}
	}

	return TX1StringStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TX1StringStringMap) Filter(p func(string, string) bool) *TX1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
// The original map is not modified.
func (mm *TX1StringStringMap) Partition(p func(string, string) bool) (matching *TX1StringStringMap, others *TX1StringStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTX1StringStringMap()
	others = NewTX1StringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1StringStringMap) Map(f func(string, string) (string, string)) *TX1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TX1StringStringMap) FlatMap(f func(string, string) []TX1StringStringTuple) *TX1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm *TX1StringStringMap) Equals(other *TX1StringStringMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

	mm.s.RLock()
	other.s.RLock()
	defer mm.s.RUnlock()
	defer other.s.RUnlock()

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

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TX1StringStringMap) Clone() *TX1StringStringMap {
	if mm == nil {
		return nil
	}

	result := NewTX1StringStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *TX1StringStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *TX1StringStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *TX1StringStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *TX1StringStringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *TX1StringStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *TX1StringStringMap) UnmarshalJSON(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *TX1StringStringMap) MarshalJSON() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return json.Marshal(mm.m)
}
