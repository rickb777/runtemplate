// A simple type derived from map[string]Apple.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package simple

// SPStringAppleMap is the primary type that represents a map
type SPStringAppleMap map[*string]*Apple

// SPStringAppleTuple represents a key/value pair.
type SPStringAppleTuple struct {
	Key *string
	Val *Apple
}

// SPStringAppleTuples can be used as a builder for unmodifiable maps.
type SPStringAppleTuples []SPStringAppleTuple

func (ts SPStringAppleTuples) Append1(k *string, v *Apple) SPStringAppleTuples {
	return append(ts, SPStringAppleTuple{k, v})
}

func (ts SPStringAppleTuples) Append2(k1 *string, v1 *Apple, k2 *string, v2 *Apple) SPStringAppleTuples {
	return append(ts, SPStringAppleTuple{k1, v1}, SPStringAppleTuple{k2, v2})
}

func (ts SPStringAppleTuples) Append3(k1 *string, v1 *Apple, k2 *string, v2 *Apple, k3 *string, v3 *Apple) SPStringAppleTuples {
	return append(ts, SPStringAppleTuple{k1, v1}, SPStringAppleTuple{k2, v2}, SPStringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSPStringAppleMap() SPStringAppleMap {
	return SPStringAppleMap(make(map[*string]*Apple))
}

// NewSPStringAppleMap creates and returns a reference to a map containing one item.
func NewSPStringAppleMap1(k *string, v *Apple) SPStringAppleMap {
	mm := newSPStringAppleMap()
	mm[k] = v
	return mm
}

// NewSPStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewSPStringAppleMap(kv ...SPStringAppleTuple) SPStringAppleMap {
	mm := newSPStringAppleMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SPStringAppleMap) Keys() []*string {
	var s []*string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SPStringAppleMap) ToSlice() []SPStringAppleTuple {
	var s []SPStringAppleTuple
	for k, v := range mm {
		s = append(s, SPStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SPStringAppleMap) Get(k *string) (*Apple, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SPStringAppleMap) Put(k *string, v *Apple) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SPStringAppleMap) ContainsKey(k *string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SPStringAppleMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SPStringAppleMap) Clear() {
	*mm = make(map[*string]*Apple)
}

// Remove allows the removal of a single item from the map.
func (mm SPStringAppleMap) Remove(k *string) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SPStringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SPStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SPStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SPStringAppleMap) Forall(fn func(*string, *Apple) bool) bool {
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
func (mm SPStringAppleMap) Exists(fn func(*string, *Apple) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SPStringAppleMap) Filter(fn func(*string, *Apple) bool) SPStringAppleMap {
	result := NewSPStringAppleMap()
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
func (mm SPStringAppleMap) Partition(fn func(*string, *Apple) bool) (matching SPStringAppleMap, others SPStringAppleMap) {
	matching = NewSPStringAppleMap()
	others = NewSPStringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SPStringAppleMap) Clone() SPStringAppleMap {
	result := NewSPStringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


