// An encapsulated map[string]Apple
// Not thread-safe.
//
// Generated from encap.tpl with Key=string Type=Apple
// options: Comparable=<no value> Stringer=<no value> Mutable=true

package maps



// EPStringAppleMap is the primary type that represents a map
type EPStringAppleMap struct {
	m map[*string]*Apple
}

// EPStringAppleTuple represents a key/value pair.
type EPStringAppleTuple struct {
	Key *string
	Val *Apple
}

// EPStringAppleTuples can be used as a builder for unmodifiable maps.
type EPStringAppleTuples []EPStringAppleTuple

func (ts EPStringAppleTuples) Append1(k *string, v *Apple) EPStringAppleTuples {
	return append(ts, EPStringAppleTuple{k, v})
}

func (ts EPStringAppleTuples) Append2(k1 *string, v1 *Apple, k2 *string, v2 *Apple) EPStringAppleTuples {
	return append(ts, EPStringAppleTuple{k1, v1}, EPStringAppleTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newEPStringAppleMap() EPStringAppleMap {
	return EPStringAppleMap{
		make(map[*string]*Apple),
	}
}

// NewEPStringAppleMap creates and returns a reference to a map containing one item.
func NewEPStringAppleMap1(k *string, v *Apple) EPStringAppleMap {
	mm := newEPStringAppleMap()
	mm.m[k] = v
	return mm
}

// NewEPStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewEPStringAppleMap(kv ...EPStringAppleTuple) EPStringAppleMap {
	mm := newEPStringAppleMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *EPStringAppleMap) Keys() []*string {
	var s []*string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *EPStringAppleMap) ToSlice() []EPStringAppleTuple {
	var s []EPStringAppleTuple
	for k, v := range mm.m {
		s = append(s, EPStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *EPStringAppleMap) Get(k *string) (*Apple, bool) {
	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *EPStringAppleMap) Put(k *string, v *Apple) bool {
	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *EPStringAppleMap) ContainsKey(k *string) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *EPStringAppleMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *EPStringAppleMap) Clear() {
	mm.m = make(map[*string]*Apple)
}

// Remove allows the removal of a single item from the map.
func (mm *EPStringAppleMap) Remove(k *string) {
	delete(mm.m, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *EPStringAppleMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *EPStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *EPStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *EPStringAppleMap) Forall(fn func(*string, *Apple) bool) bool {
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
func (mm *EPStringAppleMap) Exists(fn func(*string, *Apple) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *EPStringAppleMap) Filter(fn func(*string, *Apple) bool) EPStringAppleMap {
	result := NewEPStringAppleMap()
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
func (mm *EPStringAppleMap) Partition(fn func(*string, *Apple) bool) (matching EPStringAppleMap, others EPStringAppleMap) {
	matching = NewEPStringAppleMap()
	others = NewEPStringAppleMap()
	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}


// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *EPStringAppleMap) Clone() EPStringAppleMap {
	result := NewEPStringAppleMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


