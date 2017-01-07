// A simple map[Apple]Pear.
// Not thread-safe.
//
// Generated from simple.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Stringer=<no value> Mutable=false

package maps




// SPApplePearMap is the primary type that represents a map
type SPApplePearMap map[*Apple]*Pear

// SPApplePearTuple represents a key/value pair.
type SPApplePearTuple struct {
	Key *Apple
	Val *Pear
}

// SPApplePearTuples can be used as a builder for unmodifiable maps.
type SPApplePearTuples []SPApplePearTuple

func (ts SPApplePearTuples) Append1(k *Apple, v *Pear) SPApplePearTuples {
	return append(ts, SPApplePearTuple{k, v})
}

func (ts SPApplePearTuples) Append2(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear) SPApplePearTuples {
	return append(ts, SPApplePearTuple{k1, v1}, SPApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSPApplePearMap() SPApplePearMap {
	return SPApplePearMap(make(map[*Apple]*Pear))
}

// NewSPApplePearMap creates and returns a reference to a map containing one item.
func NewSPApplePearMap1(k *Apple, v *Pear) SPApplePearMap {
	mm := newSPApplePearMap()
	mm[k] = v
	return mm
}

// NewSPApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewSPApplePearMap(kv ...SPApplePearTuple) SPApplePearMap {
	mm := newSPApplePearMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SPApplePearMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SPApplePearMap) ToSlice() []SPApplePearTuple {
	var s []SPApplePearTuple
	for k, v := range mm {
		s = append(s, SPApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SPApplePearMap) Get(k *Apple) (*Pear, bool) {
	v, found := mm[k]
	return v, found
}


// ContainsKey determines if a given item is already in the map.
func (mm SPApplePearMap) ContainsKey(k *Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SPApplePearMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SPApplePearMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SPApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SPApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SPApplePearMap) Forall(fn func(*Apple, *Pear) bool) bool {
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
func (mm SPApplePearMap) Exists(fn func(*Apple, *Pear) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SPApplePearMap) Filter(fn func(*Apple, *Pear) bool) SPApplePearMap {
	result := NewSPApplePearMap()
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
func (mm SPApplePearMap) Partition(fn func(*Apple, *Pear) bool) (matching SPApplePearMap, others SPApplePearMap) {
	matching = NewSPApplePearMap()
	others = NewSPApplePearMap()
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
func (mm SPApplePearMap) Clone() SPApplePearMap {
	result := NewSPApplePearMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


