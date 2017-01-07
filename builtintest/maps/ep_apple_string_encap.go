// An encapsulated map[Apple]string
// Not thread-safe.
//
// Generated from encap.tpl with Key=Apple Type=string
// options: Comparable=<no value> Stringer=<no value> Mutable=true

package maps



// EPAppleStringMap is the primary type that represents a map
type EPAppleStringMap struct {
	m map[*Apple]*string
}

// EPAppleStringTuple represents a key/value pair.
type EPAppleStringTuple struct {
	Key *Apple
	Val *string
}

// EPAppleStringTuples can be used as a builder for unmodifiable maps.
type EPAppleStringTuples []EPAppleStringTuple

func (ts EPAppleStringTuples) Append1(k *Apple, v *string) EPAppleStringTuples {
	return append(ts, EPAppleStringTuple{k, v})
}

func (ts EPAppleStringTuples) Append2(k1 *Apple, v1 *string, k2 *Apple, v2 *string) EPAppleStringTuples {
	return append(ts, EPAppleStringTuple{k1, v1}, EPAppleStringTuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func newEPAppleStringMap() EPAppleStringMap {
	return EPAppleStringMap{
		make(map[*Apple]*string),
	}
}

// NewEPAppleStringMap creates and returns a reference to a map containing one item.
func NewEPAppleStringMap1(k *Apple, v *string) EPAppleStringMap {
	mm := newEPAppleStringMap()
	mm.m[k] = v
	return mm
}

// NewEPAppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewEPAppleStringMap(kv ...EPAppleStringTuple) EPAppleStringMap {
	mm := newEPAppleStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *EPAppleStringMap) Keys() []*Apple {
	var s []*Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *EPAppleStringMap) ToSlice() []EPAppleStringTuple {
	var s []EPAppleStringTuple
	for k, v := range mm.m {
		s = append(s, EPAppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *EPAppleStringMap) Get(k *Apple) (*string, bool) {
	v, found := mm.m[k]
	return v, found
}


// Put adds an item to the current map, replacing any prior value.
func (mm *EPAppleStringMap) Put(k *Apple, v *string) bool {
	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}


// ContainsKey determines if a given item is already in the map.
func (mm *EPAppleStringMap) ContainsKey(k *Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *EPAppleStringMap) ContainsAllKeys(kk ...*Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}


// Clear clears the entire map.
func (mm *EPAppleStringMap) Clear() {
	mm.m = make(map[*Apple]*string)
}

// Remove allows the removal of a single item from the map.
func (mm *EPAppleStringMap) Remove(k *Apple) {
	delete(mm.m, k)
}


// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *EPAppleStringMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *EPAppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *EPAppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *EPAppleStringMap) Forall(fn func(*Apple, *string) bool) bool {
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
func (mm *EPAppleStringMap) Exists(fn func(*Apple, *string) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *EPAppleStringMap) Filter(fn func(*Apple, *string) bool) EPAppleStringMap {
	result := NewEPAppleStringMap()
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
func (mm *EPAppleStringMap) Partition(fn func(*Apple, *string) bool) (matching EPAppleStringMap, others EPAppleStringMap) {
	matching = NewEPAppleStringMap()
	others = NewEPAppleStringMap()
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
func (mm *EPAppleStringMap) Clone() EPAppleStringMap {
	result := NewEPAppleStringMap()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


