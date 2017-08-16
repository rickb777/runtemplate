// A simple type derived from map[string]string.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:<no value> KeyList:<no value> Mutable:always

package simple

// SP1StringStringMap is the primary type that represents a map
type SP1StringStringMap map[*string]*string

// SP1StringStringTuple represents a key/value pair.
type SP1StringStringTuple struct {
	Key *string
	Val *string
}

// SP1StringStringTuples can be used as a builder for unmodifiable maps.
type SP1StringStringTuples []SP1StringStringTuple

func (ts SP1StringStringTuples) Append1(k *string, v *string) SP1StringStringTuples {
	return append(ts, SP1StringStringTuple{k, v})
}

func (ts SP1StringStringTuples) Append2(k1 *string, v1 *string, k2 *string, v2 *string) SP1StringStringTuples {
	return append(ts, SP1StringStringTuple{k1, v1}, SP1StringStringTuple{k2, v2})
}

func (ts SP1StringStringTuples) Append3(k1 *string, v1 *string, k2 *string, v2 *string, k3 *string, v3 *string) SP1StringStringTuples {
	return append(ts, SP1StringStringTuple{k1, v1}, SP1StringStringTuple{k2, v2}, SP1StringStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1StringStringMap() SP1StringStringMap {
	return SP1StringStringMap(make(map[*string]*string))
}

// NewSP1StringStringMap creates and returns a reference to a map containing one item.
func NewSP1StringStringMap1(k *string, v *string) SP1StringStringMap {
	mm := newSP1StringStringMap()
	mm[k] = v
	return mm
}

// NewSP1StringStringMap creates and returns a reference to a map, optionally containing some items.
func NewSP1StringStringMap(kv ...SP1StringStringTuple) SP1StringStringMap {
	mm := newSP1StringStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1StringStringMap) Keys() []*string {
	var s []*string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1StringStringMap) Values() []*string {
	var s []*string
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1StringStringMap) ToSlice() []SP1StringStringTuple {
	var s []SP1StringStringTuple
	for k, v := range mm {
		s = append(s, SP1StringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1StringStringMap) Get(k *string) (*string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1StringStringMap) Put(k *string, v *string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1StringStringMap) ContainsKey(k *string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1StringStringMap) ContainsAllKeys(kk ...*string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1StringStringMap) Clear() {
	*mm = make(map[*string]*string)
}

// Remove allows the removal of a single item from the map.
func (mm SP1StringStringMap) Remove(k *string) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1StringStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1StringStringMap) Forall(fn func(*string, *string) bool) bool {
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
func (mm SP1StringStringMap) Exists(fn func(*string, *string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SP1StringStringMap) Filter(fn func(*string, *string) bool) SP1StringStringMap {
	result := NewSP1StringStringMap()
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
func (mm SP1StringStringMap) Partition(fn func(*string, *string) bool) (matching SP1StringStringMap, others SP1StringStringMap) {
	matching = NewSP1StringStringMap()
	others = NewSP1StringStringMap()
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
func (mm SP1StringStringMap) Equals(other SP1StringStringMap) bool {
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
func (mm SP1StringStringMap) Clone() SP1StringStringMap {
	result := NewSP1StringStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


