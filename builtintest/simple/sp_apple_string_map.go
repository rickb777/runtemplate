// A simple type derived from map[Apple]string.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package simple

// SPAppleStringMap is the primary type that represents a map
type SPAppleStringMap map[*Apple]*string

// SPAppleStringTuple represents a key/value pair.
type SPAppleStringTuple struct {
	Key *Apple
	Val *string
}

// SPAppleStringTuples can be used as a builder for unmodifiable maps.
type SPAppleStringTuples []SPAppleStringTuple

func (ts SPAppleStringTuples) Append1(k *Apple, v *string) SPAppleStringTuples {
	return append(ts, SPAppleStringTuple{k, v})
}

func (ts SPAppleStringTuples) Append2(k1 *Apple, v1 *string, k2 *Apple, v2 *string) SPAppleStringTuples {
	return append(ts, SPAppleStringTuple{k1, v1}, SPAppleStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSPAppleStringMap() SPAppleStringMap {
	return SPAppleStringMap(make(map[*Apple]*string))
}

// NewSPAppleStringMap creates and returns a reference to a map containing one item.
func NewSPAppleStringMap1(k *Apple, v *string) SPAppleStringMap {
	mm := newSPAppleStringMap()
	mm[k] = v
	return mm
}

// NewSPAppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewSPAppleStringMap(kv ...SPAppleStringTuple) SPAppleStringMap {
	mm := newSPAppleStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SPAppleStringMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SPAppleStringMap) ToSlice() []SPAppleStringTuple {
	var s []SPAppleStringTuple
	for k, v := range mm {
		s = append(s, SPAppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SPAppleStringMap) Get(k *Apple) (*string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SPAppleStringMap) Put(k *Apple, v *string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SPAppleStringMap) ContainsKey(k *Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SPAppleStringMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Remove allows the removal of a single item from the map.
func (mm SPAppleStringMap) Remove(k *Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SPAppleStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SPAppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SPAppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SPAppleStringMap) Forall(fn func(*Apple, *string) bool) bool {
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
func (mm SPAppleStringMap) Exists(fn func(*Apple, *string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SPAppleStringMap) Filter(fn func(*Apple, *string) bool) SPAppleStringMap {
	result := NewSPAppleStringMap()
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
func (mm SPAppleStringMap) Partition(fn func(*Apple, *string) bool) (matching SPAppleStringMap, others SPAppleStringMap) {
	matching = NewSPAppleStringMap()
	others = NewSPAppleStringMap()
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
func (mm SPAppleStringMap) Clone() SPAppleStringMap {
	result := NewSPAppleStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


