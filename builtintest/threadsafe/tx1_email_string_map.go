// An encapsulated map[Email]string.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=Email Type=string
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package threadsafe

import (

	"bytes"
	"fmt"
	"sync"
)

// TX1EmailStringMap is the primary type that represents a thread-safe map
type TX1EmailStringMap struct {
	s *sync.RWMutex
	m map[Email]string
}

// TX1EmailStringTuple represents a key/value pair.
type TX1EmailStringTuple struct {
	Key Email
	Val string
}

// TX1EmailStringTuples can be used as a builder for unmodifiable maps.
type TX1EmailStringTuples []TX1EmailStringTuple

func (ts TX1EmailStringTuples) Append1(k Email, v string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k, v})
}

func (ts TX1EmailStringTuples) Append2(k1 Email, v1 string, k2 Email, v2 string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k1, v1}, TX1EmailStringTuple{k2, v2})
}

func (ts TX1EmailStringTuples) Append3(k1 Email, v1 string, k2 Email, v2 string, k3 Email, v3 string) TX1EmailStringTuples {
	return append(ts, TX1EmailStringTuple{k1, v1}, TX1EmailStringTuple{k2, v2}, TX1EmailStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1EmailStringMap() TX1EmailStringMap {
	return TX1EmailStringMap{
		s: &sync.RWMutex{},
		m: make(map[Email]string),
	}
}

// NewTX1EmailStringMap creates and returns a reference to a map containing one item.
func NewTX1EmailStringMap1(k Email, v string) TX1EmailStringMap {
	mm := newTX1EmailStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1EmailStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1EmailStringMap(kv ...TX1EmailStringTuple) TX1EmailStringMap {
	mm := newTX1EmailStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1EmailStringMap) Keys() []Email {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Email
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1EmailStringMap) Values() []string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []string
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1EmailStringMap) ToSlice() []TX1EmailStringTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TX1EmailStringTuple
	for k, v := range mm.m {
		s = append(s, TX1EmailStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1EmailStringMap) Get(k Email) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1EmailStringMap) Put(k Email, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1EmailStringMap) ContainsKey(k Email) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1EmailStringMap) ContainsAllKeys(kk ...Email) bool {
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
func (mm *TX1EmailStringMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[Email]string)
}

// Remove allows the removal of a single item from the map.
func (mm TX1EmailStringMap) Remove(k Email) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1EmailStringMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1EmailStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1EmailStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TX1EmailStringMap) DropWhere(fn func(Email, string) bool) TX1EmailStringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TX1EmailStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TX1EmailStringTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TX1EmailStringMap) Foreach(fn func(Email, string)) {
	mm.s.Lock()
	defer mm.s.Unlock()

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
func (mm TX1EmailStringMap) Forall(fn func(Email, string) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm TX1EmailStringMap) Exists(fn func(Email, string) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (mm TX1EmailStringMap) Find(fn func(Email, string) bool) (TX1EmailStringTuple, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return TX1EmailStringTuple{k, v}, true
		}
	}

	return TX1EmailStringTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1EmailStringMap) Filter(fn func(Email, string) bool) TX1EmailStringMap {
	result := NewTX1EmailStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm TX1EmailStringMap) Partition(fn func(Email, string) bool) (matching TX1EmailStringMap, others TX1EmailStringMap) {
	matching = NewTX1EmailStringMap()
	others = NewTX1EmailStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm TX1EmailStringMap) Map(fn func(Email, string) (Email, string)) TX1EmailStringMap {
	result := NewTX1EmailStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

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
func (mm TX1EmailStringMap) FlatMap(fn func(Email, string) []TX1EmailStringTuple) TX1EmailStringMap {
	result := NewTX1EmailStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1EmailStringMap) Clone() TX1EmailStringMap {
	result := NewTX1EmailStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TX1EmailStringMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1EmailStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1EmailStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
// The map entries are sorted by their keys.
func (mm TX1EmailStringMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1EmailStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

