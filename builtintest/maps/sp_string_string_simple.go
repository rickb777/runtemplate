// A simple map[string]string.
// Not thread-safe.
//
// Generated from simple.tpl with Key=string Type=string
// options: Comparable=true Stringer=<no value> Mutable=true

package maps




// SPStringStringMap is the primary type that represents a map
type SPStringStringMap map[*string]*string

// SPStringStringTuple represents a key/value pair.
type SPStringStringTuple struct {
	Key *string
	Val *string
}

// SPStringStringTuples can be used as a builder for unmodifiable maps.
type SPStringStringTuples []SPStringStringTuple

func (ts SPStringStringTuples) Append1(k *string, v *string) SPStringStringTuples {
	return append(ts, SPStringStringTuple{k, v})
}

func (ts SPStringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) SPStringStringTuples {
	return append(ts, SPStringStringTuple{k1, v1}, SPStringStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSPStringStringMap() SPStringStringMap {
	return SPStringStringMap(make(map[*string]*string))
}

// NewSPStringStringMap creates and returns a reference to a map containing one item.
func NewSPStringStringMap1(k *string, v *string) SPStringStringMap {
	mm := newSPStringStringMap()
	mm[k] = v
	return mm
}

// NewSPStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewSPStringStringMap(kv ...SPStringStringTuple) SPStringStringMap {
	mm := newSPStringStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SPStringStringMap) Keys() []*string {
	var s []*string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SPStringStringMap) ToSlice() []SPStringStringTuple {
	var s []SPStringStringTuple
	for k, v := range mm {
		s = append(s, SPStringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SPStringStringMap) Get(k *string) (*string, bool) {
	v, found := mm[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm SPStringStringMap) Put(k *string, v *string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm SPStringStringMap) ContainsKey(k *string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SPStringStringMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm SPStringStringMap) Clear() {
	mm = make(map[*string]*string)
}

// Remove allows the removal of a single item from the map.
func (mm SPStringStringMap) Remove(k *string) {
	delete(mm, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SPStringStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SPStringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SPStringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SPStringStringMap) Forall(fn func(*string, *string) bool) bool {
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
func (mm SPStringStringMap) Exists(fn func(*string, *string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SPStringStringMap) Filter(fn func(*string, *string) bool) SPStringStringMap {
	result := NewSPStringStringMap()
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
func (mm SPStringStringMap) Partition(fn func(*string, *string) bool) (matching SPStringStringMap, others SPStringStringMap) {
	matching = NewSPStringStringMap()
	others = NewSPStringStringMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm SPStringStringMap) Equals(other SPStringStringMap) bool {
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
func (mm SPStringStringMap) Clone() SPStringStringMap {
	result := NewSPStringStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


