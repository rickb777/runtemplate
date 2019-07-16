// An encapsulated map[int]*int.
// Note that the api uses *int but the map uses int keys.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=*int Type=*int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"bytes"
	"fmt"
	"sync"
)

// TP1IntIntMap is the primary type that represents a thread-safe map
type TP1IntIntMap struct {
	s *sync.RWMutex
	m map[int]*int
}

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

//-------------------------------------------------------------------------------------------------

func newTP1IntIntMap() *TP1IntIntMap {
	return &TP1IntIntMap{
		s: &sync.RWMutex{},
		m: make(map[int]*int),
	}
}

// NewTP1IntIntMap1 creates and returns a reference to a map containing one item.
func NewTP1IntIntMap1(k *int, v *int) *TP1IntIntMap {
	mm := newTP1IntIntMap()
	mm.m[*k] = v
	return mm
}

// NewTP1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTP1IntIntMap(kv ...TP1IntIntTuple) *TP1IntIntMap {
	mm := newTP1IntIntMap()
	for _, t := range kv {
		mm.m[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TP1IntIntMap) Keys() []*int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*int, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, &k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TP1IntIntMap) Values() []*int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TP1IntIntMap) slice() []TP1IntIntTuple {
	if mm == nil {
		return nil
	}

	s := make([]TP1IntIntTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TP1IntIntTuple{(&k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TP1IntIntMap) ToSlice() []TP1IntIntTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TP1IntIntMap) Get(k *int) (*int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TP1IntIntMap) Put(k *int, v *int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[*k]
	mm.m[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TP1IntIntMap) ContainsKey(k *int) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TP1IntIntMap) ContainsAllKeys(kk ...*int) bool {
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
func (mm *TP1IntIntMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int]*int)
	}
}

// Remove a single item from the map.
func (mm *TP1IntIntMap) Remove(k *int) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, *k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TP1IntIntMap) Pop(k *int) (*int, bool) {
	if mm == nil {
		return nil, false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[*k]
	delete(mm.m, *k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *TP1IntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TP1IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TP1IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TP1IntIntMap) DropWhere(fn func(*int, *int) bool) TP1IntIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TP1IntIntTuples, 0)
	for k, v := range mm.m {
		if fn(&k, v) {
			removed = append(removed, TP1IntIntTuple{(&k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TP1IntIntMap) Foreach(f func(*int, *int)) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

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
func (mm *TP1IntIntMap) Forall(p func(*int, *int) bool) bool {
	if mm == nil {
		return true
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm *TP1IntIntMap) Exists(p func(*int, *int) bool) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(&k, v) {
			return true
		}
	}

	return false
}

// Find returns the first *int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *TP1IntIntMap) Find(p func(*int, *int) bool) (TP1IntIntTuple, bool) {
	if mm == nil {
		return TP1IntIntTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(&k, v) {
			return TP1IntIntTuple{(&k), v}, true
		}
	}

	return TP1IntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TP1IntIntMap) Filter(p func(*int, *int) bool) *TP1IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewTP1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm *TP1IntIntMap) Partition(p func(*int, *int) bool) (matching *TP1IntIntMap, others *TP1IntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTP1IntIntMap()
	others = NewTP1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(&k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TP1IntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1IntIntMap) Map(f func(*int, *int) (*int, *int)) *TP1IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewTP1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(&k1, v1)
		result.m[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1IntIntMap) FlatMap(f func(*int, *int) []TP1IntIntTuple) *TP1IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewTP1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm *TP1IntIntMap) Equals(other *TP1IntIntMap) bool {
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
		if !found || *v1 != *v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TP1IntIntMap) Clone() *TP1IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewTP1IntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *TP1IntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *TP1IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *TP1IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *TP1IntIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *TP1IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
