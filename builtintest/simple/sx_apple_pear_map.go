// A simple type derived from map[Apple]Pear.
// Not thread-safe.
//
// Generated from map.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Stringer=<no value> Mutable=always

package simple

// SXApplePearMap is the primary type that represents a map
type SXApplePearMap map[Apple]Pear

// SXApplePearTuple represents a key/value pair.
type SXApplePearTuple struct {
	Key Apple
	Val Pear
}

// SXApplePearTuples can be used as a builder for unmodifiable maps.
type SXApplePearTuples []SXApplePearTuple

func (ts SXApplePearTuples) Append1(k Apple, v Pear) SXApplePearTuples {
	return append(ts, SXApplePearTuple{k, v})
}

func (ts SXApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) SXApplePearTuples {
	return append(ts, SXApplePearTuple{k1, v1}, SXApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newSXApplePearMap() SXApplePearMap {
	return SXApplePearMap(make(map[Apple]Pear))
}

// NewSXApplePearMap creates and returns a reference to a map containing one item.
func NewSXApplePearMap1(k Apple, v Pear) SXApplePearMap {
	mm := newSXApplePearMap()
	mm[k] = v
	return mm
}

// NewSXApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewSXApplePearMap(kv ...SXApplePearTuple) SXApplePearMap {
	mm := newSXApplePearMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SXApplePearMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SXApplePearMap) ToSlice() []SXApplePearTuple {
	var s []SXApplePearTuple
	for k, v := range mm {
		s = append(s, SXApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SXApplePearMap) Get(k Apple) (Pear, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SXApplePearMap) Put(k Apple, v Pear) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SXApplePearMap) ContainsKey(k Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SXApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Remove allows the removal of a single item from the map.
func (mm SXApplePearMap) Remove(k Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SXApplePearMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SXApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SXApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SXApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm SXApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm SXApplePearMap) Filter(fn func(Apple, Pear) bool) SXApplePearMap {
	result := NewSXApplePearMap()
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
func (mm SXApplePearMap) Partition(fn func(Apple, Pear) bool) (matching SXApplePearMap, others SXApplePearMap) {
	matching = NewSXApplePearMap()
	others = NewSXApplePearMap()
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
func (mm SXApplePearMap) Clone() SXApplePearMap {
	result := NewSXApplePearMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


