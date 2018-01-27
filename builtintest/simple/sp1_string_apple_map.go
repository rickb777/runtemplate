// A simple type derived from map[string]Apple.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> Mutable:always

package simple

// SP1StringAppleMap is the primary type that represents a map
type SP1StringAppleMap map[*string]*Apple

// SP1StringAppleTuple represents a key/value pair.
type SP1StringAppleTuple struct {
	Key *string
	Val *Apple
}

// SP1StringAppleTuples can be used as a builder for unmodifiable maps.
type SP1StringAppleTuples []SP1StringAppleTuple

func (ts SP1StringAppleTuples) Append1(k *string, v *Apple) SP1StringAppleTuples {
	return append(ts, SP1StringAppleTuple{k, v})
}

func (ts SP1StringAppleTuples) Append2(k1 *string, v1 *Apple, k2 *string, v2 *Apple) SP1StringAppleTuples {
	return append(ts, SP1StringAppleTuple{k1, v1}, SP1StringAppleTuple{k2, v2})
}

func (ts SP1StringAppleTuples) Append3(k1 *string, v1 *Apple, k2 *string, v2 *Apple, k3 *string, v3 *Apple) SP1StringAppleTuples {
	return append(ts, SP1StringAppleTuple{k1, v1}, SP1StringAppleTuple{k2, v2}, SP1StringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1StringAppleMap() SP1StringAppleMap {
	return SP1StringAppleMap(make(map[*string]*Apple))
}

// NewSP1StringAppleMap creates and returns a reference to a map containing one item.
func NewSP1StringAppleMap1(k *string, v *Apple) SP1StringAppleMap {
	mm := newSP1StringAppleMap()
	mm[k] = v
	return mm
}

// NewSP1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewSP1StringAppleMap(kv ...SP1StringAppleTuple) SP1StringAppleMap {
	mm := newSP1StringAppleMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1StringAppleMap) Keys() []*string {
	var s []*string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1StringAppleMap) Values() []*Apple {
	var s []*Apple
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1StringAppleMap) ToSlice() []SP1StringAppleTuple {
	var s []SP1StringAppleTuple
	for k, v := range mm {
		s = append(s, SP1StringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1StringAppleMap) Get(k *string) (*Apple, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1StringAppleMap) Put(k *string, v *Apple) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1StringAppleMap) ContainsKey(k *string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1StringAppleMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1StringAppleMap) Clear() {
	*mm = make(map[*string]*Apple)
}

// Remove allows the removal of a single item from the map.
func (mm SP1StringAppleMap) Remove(k *string) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1StringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1StringAppleMap) Forall(fn func(*string, *Apple) bool) bool {
	for k, v := range mm {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm SP1StringAppleMap) Exists(fn func(*string, *Apple) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm SP1StringAppleMap) Filter(fn func(*string, *Apple) bool) SP1StringAppleMap {
	result := NewSP1StringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified
func (mm SP1StringAppleMap) Partition(fn func(*string, *Apple) bool) (matching SP1StringAppleMap, others SP1StringAppleMap) {
	matching = NewSP1StringAppleMap()
	others = NewSP1StringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Transform returns a new SP1AppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SP1StringAppleMap) Transform(fn func(*string, *Apple) (*string, *Apple)) SP1StringAppleMap {
	result := NewSP1StringAppleMap()

	for k1, v1 := range mm {
	    k2, v2 := fn(k1, v1)
	    result[k2] = v2
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SP1StringAppleMap) Clone() SP1StringAppleMap {
	result := NewSP1StringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


