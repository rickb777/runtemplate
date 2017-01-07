// Generated from simple.tpl with Key=string Type=Apple
// options: Comparable=<no value> Stringer=<no value> Mutable=true

package maps




// SXStringAppleMap is the primary type that represents a map
type SXStringAppleMap struct {
	M map[string]Apple
}

// SXStringAppleTuple represents a key/value pair.
type SXStringAppleTuple struct {
	Key string
	Val Apple
}

// SXStringAppleTuples can be used as a builder for unmodifiable maps.
type SXStringAppleTuples []SXStringAppleTuple

func (ts SXStringAppleTuples) Append1(k string, v Apple) SXStringAppleTuples {
	return append(ts, SXStringAppleTuple{k, v})
}

func (ts SXStringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) SXStringAppleTuples {
	return append(ts, SXStringAppleTuple{k1, v1}, SXStringAppleTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

// NewSXStringAppleMap creates and returns a reference to a map containing one item.
func NewSXStringAppleMap1(k string, v Apple) SXStringAppleMap {
	mm := SXStringAppleMap{
		M: make(map[string]Apple),
	}
	mm.M[k] = v
	return mm
}

// NewSXStringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewSXStringAppleMap(kv ...SXStringAppleTuple) SXStringAppleMap {
	mm := SXStringAppleMap{
		M: make(map[string]Apple),
	}
	for _, t := range kv {
		mm.M[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SXStringAppleMap) Keys() []string {
	var s []string
	for k, _ := range mm.M {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SXStringAppleMap) ToSlice() []SXStringAppleTuple {
	var s []SXStringAppleTuple
	for k, v := range mm.M {
		s = append(s, SXStringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *SXStringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm.M[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *SXStringAppleMap) Put(k string, v Apple) bool {
	_, found := mm.M[k]
	mm.M[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *SXStringAppleMap) ContainsKey(k string) bool {
	_, found := mm.M[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SXStringAppleMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *SXStringAppleMap) Clear() {
	mm.M = make(map[string]Apple)
}

// Remove allows the removal of a single item from the map.
func (mm *SXStringAppleMap) Remove(k string) {
	delete(mm.M, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *SXStringAppleMap) Size() int {
	return len(mm.M)
}

// IsEmpty returns true if the map is empty.
func (mm *SXStringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SXStringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *SXStringAppleMap) Forall(fn func(string, Apple) bool) bool {
	for k, v := range mm.M {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *SXStringAppleMap) Exists(fn func(string, Apple) bool) bool {
	for k, v := range mm.M {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *SXStringAppleMap) Filter(fn func(string, Apple) bool) SXStringAppleMap {
	result := NewSXStringAppleMap()
	for k, v := range mm.M {
		if fn(k, v) {
			result.M[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm *SXStringAppleMap) Partition(fn func(string, Apple) bool) (matching SXStringAppleMap, others SXStringAppleMap) {
	matching = NewSXStringAppleMap()
	others = NewSXStringAppleMap()
	for k, v := range mm.M {
		if fn(k, v) {
			matching.M[k] = v
		} else {
			others.M[k] = v
		}
	}
	return
}


// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *SXStringAppleMap) Clone() SXStringAppleMap {
	result := NewSXStringAppleMap()
	for k, v := range mm.M {
		result.M[k] = v
	}
	return result
}


