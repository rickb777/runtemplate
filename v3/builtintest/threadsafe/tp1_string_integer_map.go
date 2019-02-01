// An encapsulated map[string]*big.Int.
// Note that the api uses *string but the map uses string keys.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=*string Type=*big.Int
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.2.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"fmt"
	"sync"
	"math/big"

)

// TP1StringIntegerMap is the primary type that represents a thread-safe map
type TP1StringIntegerMap struct {
	s *sync.RWMutex
	m map[string]*big.Int
}

// TP1StringIntegerTuple represents a key/value pair.
type TP1StringIntegerTuple struct {
	Key *string
	Val *big.Int
}

// TP1StringIntegerTuples can be used as a builder for unmodifiable maps.
type TP1StringIntegerTuples []TP1StringIntegerTuple

// Append1 adds one item.
func (ts TP1StringIntegerTuples) Append1(k *string, v *big.Int) TP1StringIntegerTuples {
	return append(ts, TP1StringIntegerTuple{k, v})
}

// Append2 adds two items.
func (ts TP1StringIntegerTuples) Append2(k1 *string, v1 *big.Int, k2 *string, v2 *big.Int) TP1StringIntegerTuples {
	return append(ts, TP1StringIntegerTuple{k1, v1}, TP1StringIntegerTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1StringIntegerTuples) Append3(k1 *string, v1 *big.Int, k2 *string, v2 *big.Int, k3 *string, v3 *big.Int) TP1StringIntegerTuples {
	return append(ts, TP1StringIntegerTuple{k1, v1}, TP1StringIntegerTuple{k2, v2}, TP1StringIntegerTuple{k3, v3})
}

// TP1StringIntegerZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1StringIntegerMap
// constructor function.
func TP1StringIntegerZip(keys ...*string) TP1StringIntegerTuples {
	ts := make(TP1StringIntegerTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1StringIntegerZip.
func (ts TP1StringIntegerTuples) Values(values ...*big.Int) TP1StringIntegerTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTP1StringIntegerMap() *TP1StringIntegerMap {
	return &TP1StringIntegerMap{
		s: &sync.RWMutex{},
		m: make(map[string]*big.Int),
	}
}

// NewTP1StringIntegerMap1 creates and returns a reference to a map containing one item.
func NewTP1StringIntegerMap1(k *string, v *big.Int) *TP1StringIntegerMap {
	mm := newTP1StringIntegerMap()
	mm.m[*k] = v
	return mm
}

// NewTP1StringIntegerMap creates and returns a reference to a map, optionally containing some items.
func NewTP1StringIntegerMap(kv ...TP1StringIntegerTuple) *TP1StringIntegerMap {
	mm := newTP1StringIntegerMap()
	for _, t := range kv {
		mm.m[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TP1StringIntegerMap) Keys() []*string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, &k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TP1StringIntegerMap) Values() []*big.Int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*big.Int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TP1StringIntegerMap) slice() []TP1StringIntegerTuple {
	if mm == nil {
		return nil
	}

	s := make([]TP1StringIntegerTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TP1StringIntegerTuple{(&k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TP1StringIntegerMap) ToSlice() []TP1StringIntegerTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TP1StringIntegerMap) Get(k *string) (*big.Int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TP1StringIntegerMap) Put(k *string, v *big.Int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[*k]
	mm.m[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TP1StringIntegerMap) ContainsKey(k *string) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TP1StringIntegerMap) ContainsAllKeys(kk ...*string) bool {
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
func (mm *TP1StringIntegerMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[string]*big.Int)
	}
}

// Remove a single item from the map.
func (mm *TP1StringIntegerMap) Remove(k *string) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, *k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TP1StringIntegerMap) Pop(k *string) (*big.Int, bool) {
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
func (mm *TP1StringIntegerMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TP1StringIntegerMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TP1StringIntegerMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TP1StringIntegerMap) DropWhere(fn func(*string, *big.Int) bool) TP1StringIntegerTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TP1StringIntegerTuples, 0)
	for k, v := range mm.m {
		if fn(&k, v) {
			removed = append(removed, TP1StringIntegerTuple{(&k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TP1StringIntegerMap) Foreach(f func(*string, *big.Int)) {
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
func (mm *TP1StringIntegerMap) Forall(p func(*string, *big.Int) bool) bool {
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
func (mm *TP1StringIntegerMap) Exists(p func(*string, *big.Int) bool) bool {
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

// Find returns the first *big.Int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *TP1StringIntegerMap) Find(p func(*string, *big.Int) bool) (TP1StringIntegerTuple, bool) {
	if mm == nil {
		return TP1StringIntegerTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(&k, v) {
			return TP1StringIntegerTuple{(&k), v}, true
		}
	}

	return TP1StringIntegerTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TP1StringIntegerMap) Filter(p func(*string, *big.Int) bool) *TP1StringIntegerMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringIntegerMap()
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
func (mm *TP1StringIntegerMap) Partition(p func(*string, *big.Int) bool) (matching *TP1StringIntegerMap, others *TP1StringIntegerMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTP1StringIntegerMap()
	others = NewTP1StringIntegerMap()
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

// Map returns a new TP1IntegerMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1StringIntegerMap) Map(f func(*string, *big.Int) (*string, *big.Int)) *TP1StringIntegerMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringIntegerMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(&k1, v1)
		result.m[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1IntegerMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1StringIntegerMap) FlatMap(f func(*string, *big.Int) []TP1StringIntegerTuple) *TP1StringIntegerMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringIntegerMap()
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

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *TP1StringIntegerMap) Clone() *TP1StringIntegerMap {
	if mm == nil {
		return nil
	}

	result := NewTP1StringIntegerMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
