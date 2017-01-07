// An encapsulated map[Apple]Pear
// Not thread-safe.
//
// Generated from map.tpl with Key=Apple Type=Pear
// options: Comparable=<no value> Stringer=<no value> Mutable=false

package fast

// XApplePearMap is the primary type that represents a map
type XApplePearMap struct {
	m map[Apple]Pear
}

// XApplePearTuple represents a key/value pair.
type XApplePearTuple struct {
	Key Apple
	Val Pear
}

// XApplePearTuples can be used as a builder for unmodifiable maps.
type XApplePearTuples []XApplePearTuple

func (ts XApplePearTuples) Append1(k Apple, v Pear) XApplePearTuples {
	return append(ts, XApplePearTuple{k, v})
}

func (ts XApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) XApplePearTuples {
	return append(ts, XApplePearTuple{k1, v1}, XApplePearTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newXApplePearMap() *XApplePearMap {
	return &XApplePearMap{
		make(map[Apple]Pear),
	}
}

// NewXApplePearMap creates and returns a reference to a map containing one item.
func NewXApplePearMap1(k Apple, v Pear) *XApplePearMap {
	mm := newXApplePearMap()
	mm.m[k] = v
	return mm
}

// NewXApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewXApplePearMap(kv ...XApplePearTuple) *XApplePearMap {
	mm := newXApplePearMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *XApplePearMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *XApplePearMap) ToSlice() []XApplePearTuple {
	var s []XApplePearTuple
	for k, v := range mm.m {
		s = append(s, XApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *XApplePearMap) Get(k Apple) (Pear, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm *XApplePearMap) ContainsKey(k Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *XApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *XApplePearMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *XApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *XApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *XApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm *XApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *XApplePearMap) Filter(fn func(Apple, Pear) bool) *XApplePearMap {
	result := NewXApplePearMap()
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
func (mm *XApplePearMap) Partition(fn func(Apple, Pear) bool) (matching *XApplePearMap, others *XApplePearMap) {
	matching = NewXApplePearMap()
	others = NewXApplePearMap()
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
func (mm *XApplePearMap) Clone() *XApplePearMap {
	result := NewXApplePearMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


