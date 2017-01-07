// An encapsulated map[Apple]Pear
// Not thread-safe.
//
// Generated from encap.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Stringer=<no value> Mutable=false

package maps



// EXApplePearMap is the primary type that represents a map
type EXApplePearMap struct {
	m map[Apple]Pear
}

// EXApplePearTuple represents a key/value pair.
type EXApplePearTuple struct {
	Key Apple
	Val Pear
}

// EXApplePearTuples can be used as a builder for unmodifiable maps.
type EXApplePearTuples []EXApplePearTuple

func (ts EXApplePearTuples) Append1(k Apple, v Pear) EXApplePearTuples {
	return append(ts, EXApplePearTuple{k, v})
}

func (ts EXApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) EXApplePearTuples {
	return append(ts, EXApplePearTuple{k1, v1}, EXApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newEXApplePearMap() EXApplePearMap {
	return EXApplePearMap{
		make(map[Apple]Pear),
	}
}

// NewEXApplePearMap creates and returns a reference to a map containing one item.
func NewEXApplePearMap1(k Apple, v Pear) EXApplePearMap {
	mm := newEXApplePearMap()
	mm.m[k] = v
	return mm
}

// NewEXApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewEXApplePearMap(kv ...EXApplePearTuple) EXApplePearMap {
	mm := newEXApplePearMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *EXApplePearMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *EXApplePearMap) ToSlice() []EXApplePearTuple {
	var s []EXApplePearTuple
	for k, v := range mm.m {
		s = append(s, EXApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *EXApplePearMap) Get(k Apple) (Pear, bool) {
	v, found := mm.m[k]
	return v, found
}


// ContainsKey determines if a given item is already in the map.
func (mm *EXApplePearMap) ContainsKey(k Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *EXApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *EXApplePearMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *EXApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *EXApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *EXApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm *EXApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *EXApplePearMap) Filter(fn func(Apple, Pear) bool) EXApplePearMap {
	result := NewEXApplePearMap()
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
func (mm *EXApplePearMap) Partition(fn func(Apple, Pear) bool) (matching EXApplePearMap, others EXApplePearMap) {
	matching = NewEXApplePearMap()
	others = NewEXApplePearMap()
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
func (mm *EXApplePearMap) Clone() EXApplePearMap {
	result := NewEXApplePearMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


