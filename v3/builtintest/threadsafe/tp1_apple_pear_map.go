// An encapsulated map[Apple]*Pear.
// Note that the api uses *Apple but the map uses Apple keys.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=*Apple Type=*Pear
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"fmt"
	"sync"
)

// TP1ApplePearMap is the primary type that represents a thread-safe map
type TP1ApplePearMap struct {
	s *sync.RWMutex
	m map[Apple]*Pear
}

// TP1ApplePearTuple represents a key/value pair.
type TP1ApplePearTuple struct {
	Key *Apple
	Val *Pear
}

// TP1ApplePearTuples can be used as a builder for unmodifiable maps.
type TP1ApplePearTuples []TP1ApplePearTuple

// Append1 adds one item.
func (ts TP1ApplePearTuples) Append1(k *Apple, v *Pear) TP1ApplePearTuples {
	return append(ts, TP1ApplePearTuple{k, v})
}

// Append2 adds two items.
func (ts TP1ApplePearTuples) Append2(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear) TP1ApplePearTuples {
	return append(ts, TP1ApplePearTuple{k1, v1}, TP1ApplePearTuple{k2, v2})
}

// Append3 adds three items.
func (ts TP1ApplePearTuples) Append3(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear, k3 *Apple, v3 *Pear) TP1ApplePearTuples {
	return append(ts, TP1ApplePearTuple{k1, v1}, TP1ApplePearTuple{k2, v2}, TP1ApplePearTuple{k3, v3})
}

// TP1ApplePearZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewTP1ApplePearMap
// constructor function.
func TP1ApplePearZip(keys ...*Apple) TP1ApplePearTuples {
	ts := make(TP1ApplePearTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with TP1ApplePearZip.
func (ts TP1ApplePearTuples) Values(values ...*Pear) TP1ApplePearTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newTP1ApplePearMap() *TP1ApplePearMap {
	return &TP1ApplePearMap{
		s: &sync.RWMutex{},
		m: make(map[Apple]*Pear),
	}
}

// NewTP1ApplePearMap1 creates and returns a reference to a map containing one item.
func NewTP1ApplePearMap1(k *Apple, v *Pear) *TP1ApplePearMap {
	mm := newTP1ApplePearMap()
	mm.m[*k] = v
	return mm
}

// NewTP1ApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTP1ApplePearMap(kv ...TP1ApplePearTuple) *TP1ApplePearMap {
	mm := newTP1ApplePearMap()
	for _, t := range kv {
		mm.m[*t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *TP1ApplePearMap) Keys() []*Apple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*Apple, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, &k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *TP1ApplePearMap) Values() []*Pear {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]*Pear, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *TP1ApplePearMap) slice() []TP1ApplePearTuple {
	if mm == nil {
		return nil
	}

	s := make([]TP1ApplePearTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, TP1ApplePearTuple{(&k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *TP1ApplePearMap) ToSlice() []TP1ApplePearTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *TP1ApplePearMap) Get(k *Apple) (*Pear, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[*k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *TP1ApplePearMap) Put(k *Apple, v *Pear) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[*k]
	mm.m[*k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *TP1ApplePearMap) ContainsKey(k *Apple) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[*k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *TP1ApplePearMap) ContainsAllKeys(kk ...*Apple) bool {
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
func (mm *TP1ApplePearMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[Apple]*Pear)
	}
}

// Remove a single item from the map.
func (mm *TP1ApplePearMap) Remove(k *Apple) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, *k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *TP1ApplePearMap) Pop(k *Apple) (*Pear, bool) {
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
func (mm *TP1ApplePearMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *TP1ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *TP1ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *TP1ApplePearMap) DropWhere(fn func(*Apple, *Pear) bool) TP1ApplePearTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TP1ApplePearTuples, 0)
	for k, v := range mm.m {
		if fn(&k, v) {
			removed = append(removed, TP1ApplePearTuple{(&k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *TP1ApplePearMap) Foreach(f func(*Apple, *Pear)) {
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
func (mm *TP1ApplePearMap) Forall(p func(*Apple, *Pear) bool) bool {
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
func (mm *TP1ApplePearMap) Exists(p func(*Apple, *Pear) bool) bool {
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

// Find returns the first *Pear that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *TP1ApplePearMap) Find(p func(*Apple, *Pear) bool) (TP1ApplePearTuple, bool) {
	if mm == nil {
		return TP1ApplePearTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(&k, v) {
			return TP1ApplePearTuple{(&k), v}, true
		}
	}

	return TP1ApplePearTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *TP1ApplePearMap) Filter(p func(*Apple, *Pear) bool) *TP1ApplePearMap {
	if mm == nil {
		return nil
	}

	result := NewTP1ApplePearMap()
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
func (mm *TP1ApplePearMap) Partition(p func(*Apple, *Pear) bool) (matching *TP1ApplePearMap, others *TP1ApplePearMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewTP1ApplePearMap()
	others = NewTP1ApplePearMap()
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

// Map returns a new TP1PearMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1ApplePearMap) Map(f func(*Apple, *Pear) (*Apple, *Pear)) *TP1ApplePearMap {
	if mm == nil {
		return nil
	}

	result := NewTP1ApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(&k1, v1)
		result.m[*k2] = v2
	}

	return result
}

// FlatMap returns a new TP1PearMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *TP1ApplePearMap) FlatMap(f func(*Apple, *Pear) []TP1ApplePearTuple) *TP1ApplePearMap {
	if mm == nil {
		return nil
	}

	result := NewTP1ApplePearMap()
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
func (mm *TP1ApplePearMap) Clone() *TP1ApplePearMap {
	if mm == nil {
		return nil
	}

	result := NewTP1ApplePearMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
