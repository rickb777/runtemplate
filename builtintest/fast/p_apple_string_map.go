// An encapsulated map[Apple]string
// Not thread-safe.
//
// Generated from map.tpl with Key=Apple Type=string
// options: Comparable=<no value> Stringer=<no value> Mutable=true

package fast

// PAppleStringMap is the primary type that represents a map
type PAppleStringMap struct {
	m map[*Apple]*string
}

// PAppleStringTuple represents a key/value pair.
type PAppleStringTuple struct {
	Key *Apple
	Val *string
}

// PAppleStringTuples can be used as a builder for unmodifiable maps.
type PAppleStringTuples []PAppleStringTuple

func (ts PAppleStringTuples) Append1(k *Apple, v *string) PAppleStringTuples {
	return append(ts, PAppleStringTuple{k, v})
}

func (ts PAppleStringTuples) Append2(k1 *Apple, v1 *string, k2 *Apple, v2 *string) PAppleStringTuples {
	return append(ts, PAppleStringTuple{k1, v1}, PAppleStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newPAppleStringMap() *PAppleStringMap {
	return &PAppleStringMap{
		make(map[*Apple]*string),
	}
}

// NewPAppleStringMap creates and returns a reference to a map containing one item.
func NewPAppleStringMap1(k *Apple, v *string) *PAppleStringMap {
	mm := newPAppleStringMap()
	mm.m[k] = v
	return mm
}

// NewPAppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewPAppleStringMap(kv ...PAppleStringTuple) *PAppleStringMap {
	mm := newPAppleStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *PAppleStringMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *PAppleStringMap) ToSlice() []PAppleStringTuple {
	var s []PAppleStringTuple
	for k, v := range mm.m {
		s = append(s, PAppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *PAppleStringMap) Get(k *Apple) (*string, bool) {
	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *PAppleStringMap) Put(k *Apple, v *string) bool {
	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *PAppleStringMap) ContainsKey(k *Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *PAppleStringMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *PAppleStringMap) Clear() {
	mm.m = make(map[*Apple]*string)
}

// Remove allows the removal of a single item from the map.
func (mm *PAppleStringMap) Remove(k *Apple) {
	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *PAppleStringMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *PAppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *PAppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *PAppleStringMap) Forall(fn func(*Apple, *string) bool) bool {
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
func (mm *PAppleStringMap) Exists(fn func(*Apple, *string) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *PAppleStringMap) Filter(fn func(*Apple, *string) bool) *PAppleStringMap {
	result := NewPAppleStringMap()
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
func (mm *PAppleStringMap) Partition(fn func(*Apple, *string) bool) (matching *PAppleStringMap, others *PAppleStringMap) {
	matching = NewPAppleStringMap()
	others = NewPAppleStringMap()
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
func (mm *PAppleStringMap) Clone() *PAppleStringMap {
	result := NewPAppleStringMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


