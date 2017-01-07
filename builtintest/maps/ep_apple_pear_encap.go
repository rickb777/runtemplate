// An encapsulated map[Apple]Pear
// Not thread-safe.
//
// Generated from encap.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Stringer=<no value> Mutable=false

package maps



// EPApplePearMap is the primary type that represents a map
type EPApplePearMap struct {
	m map[*Apple]*Pear
}

// EPApplePearTuple represents a key/value pair.
type EPApplePearTuple struct {
	Key *Apple
	Val *Pear
}

// EPApplePearTuples can be used as a builder for unmodifiable maps.
type EPApplePearTuples []EPApplePearTuple

func (ts EPApplePearTuples) Append1(k *Apple, v *Pear) EPApplePearTuples {
	return append(ts, EPApplePearTuple{k, v})
}

func (ts EPApplePearTuples) Append2(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear) EPApplePearTuples {
	return append(ts, EPApplePearTuple{k1, v1}, EPApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newEPApplePearMap() EPApplePearMap {
	return EPApplePearMap{
		make(map[*Apple]*Pear),
	}
}

// NewEPApplePearMap creates and returns a reference to a map containing one item.
func NewEPApplePearMap1(k *Apple, v *Pear) EPApplePearMap {
	mm := newEPApplePearMap()
	mm.m[k] = v
	return mm
}

// NewEPApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewEPApplePearMap(kv ...EPApplePearTuple) EPApplePearMap {
	mm := newEPApplePearMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *EPApplePearMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *EPApplePearMap) ToSlice() []EPApplePearTuple {
	var s []EPApplePearTuple
	for k, v := range mm.m {
		s = append(s, EPApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *EPApplePearMap) Get(k *Apple) (*Pear, bool) {
	v, found := mm.m[k]
	return v, found
}


// ContainsKey determines if a given item is already in the map.
func (mm *EPApplePearMap) ContainsKey(k *Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *EPApplePearMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *EPApplePearMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *EPApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *EPApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *EPApplePearMap) Forall(fn func(*Apple, *Pear) bool) bool {
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
func (mm *EPApplePearMap) Exists(fn func(*Apple, *Pear) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *EPApplePearMap) Filter(fn func(*Apple, *Pear) bool) EPApplePearMap {
	result := NewEPApplePearMap()
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
func (mm *EPApplePearMap) Partition(fn func(*Apple, *Pear) bool) (matching EPApplePearMap, others EPApplePearMap) {
	matching = NewEPApplePearMap()
	others = NewEPApplePearMap()
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
func (mm *EPApplePearMap) Clone() EPApplePearMap {
	result := NewEPApplePearMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


