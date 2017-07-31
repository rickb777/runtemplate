// A simple type derived from map[Apple]string.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package simple

// SP1AppleStringMap is the primary type that represents a map
type SP1AppleStringMap map[*Apple]*string

// SP1AppleStringTuple represents a key/value pair.
type SP1AppleStringTuple struct {
	Key *Apple
	Val *string
}

// SP1AppleStringTuples can be used as a builder for unmodifiable maps.
type SP1AppleStringTuples []SP1AppleStringTuple

func (ts SP1AppleStringTuples) Append1(k *Apple, v *string) SP1AppleStringTuples {
	return append(ts, SP1AppleStringTuple{k, v})
}

func (ts SP1AppleStringTuples) Append2(k1 *Apple, v1 *string, k2 *Apple, v2 *string) SP1AppleStringTuples {
	return append(ts, SP1AppleStringTuple{k1, v1}, SP1AppleStringTuple{k2, v2})
}

func (ts SP1AppleStringTuples) Append3(k1 *Apple, v1 *string, k2 *Apple, v2 *string, k3 *Apple, v3 *string) SP1AppleStringTuples {
	return append(ts, SP1AppleStringTuple{k1, v1}, SP1AppleStringTuple{k2, v2}, SP1AppleStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1AppleStringMap() SP1AppleStringMap {
	return SP1AppleStringMap(make(map[*Apple]*string))
}

// NewSP1AppleStringMap creates and returns a reference to a map containing one item.
func NewSP1AppleStringMap1(k *Apple, v *string) SP1AppleStringMap {
	mm := newSP1AppleStringMap()
	mm[k] = v
	return mm
}

// NewSP1AppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewSP1AppleStringMap(kv ...SP1AppleStringTuple) SP1AppleStringMap {
	mm := newSP1AppleStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1AppleStringMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1AppleStringMap) Values() []*string {
	var s []*string
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1AppleStringMap) ToSlice() []SP1AppleStringTuple {
	var s []SP1AppleStringTuple
	for k, v := range mm {
		s = append(s, SP1AppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1AppleStringMap) Get(k *Apple) (*string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1AppleStringMap) Put(k *Apple, v *string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1AppleStringMap) ContainsKey(k *Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1AppleStringMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1AppleStringMap) Clear() {
	*mm = make(map[*Apple]*string)
}

// Remove allows the removal of a single item from the map.
func (mm SP1AppleStringMap) Remove(k *Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1AppleStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1AppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1AppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1AppleStringMap) Forall(fn func(*Apple, *string) bool) bool {
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
func (mm SP1AppleStringMap) Exists(fn func(*Apple, *string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SP1AppleStringMap) Filter(fn func(*Apple, *string) bool) SP1AppleStringMap {
	result := NewSP1AppleStringMap()
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
func (mm SP1AppleStringMap) Partition(fn func(*Apple, *string) bool) (matching SP1AppleStringMap, others SP1AppleStringMap) {
	matching = NewSP1AppleStringMap()
	others = NewSP1AppleStringMap()
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
func (mm SP1AppleStringMap) Clone() SP1AppleStringMap {
	result := NewSP1AppleStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


