// A simple type derived from map[Apple]Pear.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=Pear
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> Mutable:always

package simple

// SP1ApplePearMap is the primary type that represents a map
type SP1ApplePearMap map[*Apple]*Pear

// SP1ApplePearTuple represents a key/value pair.
type SP1ApplePearTuple struct {
	Key *Apple
	Val *Pear
}

// SP1ApplePearTuples can be used as a builder for unmodifiable maps.
type SP1ApplePearTuples []SP1ApplePearTuple

func (ts SP1ApplePearTuples) Append1(k *Apple, v *Pear) SP1ApplePearTuples {
	return append(ts, SP1ApplePearTuple{k, v})
}

func (ts SP1ApplePearTuples) Append2(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear) SP1ApplePearTuples {
	return append(ts, SP1ApplePearTuple{k1, v1}, SP1ApplePearTuple{k2, v2})
}

func (ts SP1ApplePearTuples) Append3(k1 *Apple, v1 *Pear, k2 *Apple, v2 *Pear, k3 *Apple, v3 *Pear) SP1ApplePearTuples {
	return append(ts, SP1ApplePearTuple{k1, v1}, SP1ApplePearTuple{k2, v2}, SP1ApplePearTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newSP1ApplePearMap() SP1ApplePearMap {
	return SP1ApplePearMap(make(map[*Apple]*Pear))
}

// NewSP1ApplePearMap creates and returns a reference to a map containing one item.
func NewSP1ApplePearMap1(k *Apple, v *Pear) SP1ApplePearMap {
	mm := newSP1ApplePearMap()
	mm[k] = v
	return mm
}

// NewSP1ApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewSP1ApplePearMap(kv ...SP1ApplePearTuple) SP1ApplePearMap {
	mm := newSP1ApplePearMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm SP1ApplePearMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm SP1ApplePearMap) Values() []*Pear {
	var s []*Pear
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm SP1ApplePearMap) ToSlice() []SP1ApplePearTuple {
	var s []SP1ApplePearTuple
	for k, v := range mm {
		s = append(s, SP1ApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm SP1ApplePearMap) Get(k *Apple) (*Pear, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm SP1ApplePearMap) Put(k *Apple, v *Pear) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm SP1ApplePearMap) ContainsKey(k *Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm SP1ApplePearMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SP1ApplePearMap) Clear() {
	*mm = make(map[*Apple]*Pear)
}

// Remove allows the removal of a single item from the map.
func (mm SP1ApplePearMap) Remove(k *Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm SP1ApplePearMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm SP1ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm SP1ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm SP1ApplePearMap) Forall(fn func(*Apple, *Pear) bool) bool {
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
func (mm SP1ApplePearMap) Exists(fn func(*Apple, *Pear) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm SP1ApplePearMap) Filter(fn func(*Apple, *Pear) bool) SP1ApplePearMap {
	result := NewSP1ApplePearMap()
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
// The original map is not modified
func (mm SP1ApplePearMap) Partition(fn func(*Apple, *Pear) bool) (matching SP1ApplePearMap, others SP1ApplePearMap) {
	matching = NewSP1ApplePearMap()
	others = NewSP1ApplePearMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new SP1PearMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SP1ApplePearMap) Map(fn func(*Apple, *Pear) (*Apple, *Pear)) SP1ApplePearMap {
	result := NewSP1ApplePearMap()

	for k1, v1 := range mm {
	    k2, v2 := fn(k1, v1)
	    result[k2] = v2
	}

	return result
}

// FlatMap returns a new SP1PearMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm SP1ApplePearMap) FlatMap(fn func(*Apple, *Pear) []SP1ApplePearTuple) SP1ApplePearMap {
	result := NewSP1ApplePearMap()

	for k1, v1 := range mm {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm SP1ApplePearMap) Clone() SP1ApplePearMap {
	result := NewSP1ApplePearMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


