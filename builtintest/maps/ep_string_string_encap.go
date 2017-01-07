// An encapsulated map[string]string
// Not thread-safe.
//
// Generated from encap.tpl with Key=string Type=string
// options: Comparable=true Stringer=<no value> Mutable=true

package maps



// EPStringStringMap is the primary type that represents a map
type EPStringStringMap struct {
	m map[*string]*string
}

// EPStringStringTuple represents a key/value pair.
type EPStringStringTuple struct {
	Key *string
	Val *string
}

// EPStringStringTuples can be used as a builder for unmodifiable maps.
type EPStringStringTuples []EPStringStringTuple

func (ts EPStringStringTuples) Append1(k *string, v *string) EPStringStringTuples {
	return append(ts, EPStringStringTuple{k, v})
}

func (ts EPStringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) EPStringStringTuples {
	return append(ts, EPStringStringTuple{k1, v1}, EPStringStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newEPStringStringMap() EPStringStringMap {
	return EPStringStringMap{
		make(map[*string]*string),
	}
}

// NewEPStringStringMap creates and returns a reference to a map containing one item.
func NewEPStringStringMap1(k *string, v *string) EPStringStringMap {
	mm := newEPStringStringMap()
	mm.m[k] = v
	return mm
}

// NewEPStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewEPStringStringMap(kv ...EPStringStringTuple) EPStringStringMap {
	mm := newEPStringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *EPStringStringMap) Keys() []*string {
	var s []*string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *EPStringStringMap) ToSlice() []EPStringStringTuple {
	var s []EPStringStringTuple
	for k, v := range mm.m {
		s = append(s, EPStringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *EPStringStringMap) Get(k *string) (*string, bool) {
	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *EPStringStringMap) Put(k *string, v *string) bool {
	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *EPStringStringMap) ContainsKey(k *string) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *EPStringStringMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *EPStringStringMap) Clear() {
	mm.m = make(map[*string]*string)
}

// Remove allows the removal of a single item from the map.
func (mm *EPStringStringMap) Remove(k *string) {
	delete(mm.m, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *EPStringStringMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *EPStringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *EPStringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *EPStringStringMap) Forall(fn func(*string, *string) bool) bool {
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
func (mm *EPStringStringMap) Exists(fn func(*string, *string) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *EPStringStringMap) Filter(fn func(*string, *string) bool) EPStringStringMap {
	result := NewEPStringStringMap()
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
func (mm *EPStringStringMap) Partition(fn func(*string, *string) bool) (matching EPStringStringMap, others EPStringStringMap) {
	matching = NewEPStringStringMap()
	others = NewEPStringStringMap()
	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *EPStringStringMap) Equals(other EPStringStringMap) bool {
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
func (mm *EPStringStringMap) Clone() EPStringStringMap {
	result := NewEPStringStringMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


