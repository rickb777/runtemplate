// A simple map[string]string.
// Not thread-safe.
//
// Generated from simple.tpl with Key=string Type=string
// options: Comparable=true Stringer=<no value> Mutable=true

package maps




// SXStringStringMap is the primary type that represents a map
type SXStringStringMap map[string]string

// SXStringStringTuple represents a key/value pair.
type SXStringStringTuple struct {
	Key string
	Val string
}

// SXStringStringTuples can be used as a builder for unmodifiable maps.
type SXStringStringTuples []SXStringStringTuple

func (ts SXStringStringTuples) Append1(k string, v string) SXStringStringTuples {
	return append(ts, SXStringStringTuple{k, v})
}

func (ts SXStringStringTuples) Append2(k1 string, v1 string, k2 string, v2 string) SXStringStringTuples {
	return append(ts, SXStringStringTuple{k1, v1}, SXStringStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSXStringStringMap() SXStringStringMap {
	return SXStringStringMap(make(map[string]string))
}

// NewSXStringStringMap creates and returns a reference to a map containing one item.
func NewSXStringStringMap1(k string, v string) SXStringStringMap {
	mm := newSXStringStringMap()
	mm[k] = v
	return mm
}

// NewSXStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewSXStringStringMap(kv ...SXStringStringTuple) SXStringStringMap {
	mm := newSXStringStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SXStringStringMap) Keys() []string {
	var s []string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SXStringStringMap) ToSlice() []SXStringStringTuple {
	var s []SXStringStringTuple
	for k, v := range mm {
		s = append(s, SXStringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SXStringStringMap) Get(k string) (string, bool) {
	v, found := mm[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm SXStringStringMap) Put(k string, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm SXStringStringMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SXStringStringMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm SXStringStringMap) Clear() {
	mm = make(map[string]string)
}

// Remove allows the removal of a single item from the map.
func (mm SXStringStringMap) Remove(k string) {
	delete(mm, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SXStringStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SXStringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SXStringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SXStringStringMap) Forall(fn func(string, string) bool) bool {
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
func (mm SXStringStringMap) Exists(fn func(string, string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SXStringStringMap) Filter(fn func(string, string) bool) SXStringStringMap {
	result := NewSXStringStringMap()
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
func (mm SXStringStringMap) Partition(fn func(string, string) bool) (matching SXStringStringMap, others SXStringStringMap) {
	matching = NewSXStringStringMap()
	others = NewSXStringStringMap()
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
func (mm SXStringStringMap) Equals(other SXStringStringMap) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}


// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SXStringStringMap) Clone() SXStringStringMap {
	result := NewSXStringStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


