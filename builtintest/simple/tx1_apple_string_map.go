// A simple type derived from map[Apple]string.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package simple

// TX1AppleStringMap is the primary type that represents a map
type TX1AppleStringMap map[Apple]string

// TX1AppleStringTuple represents a key/value pair.
type TX1AppleStringTuple struct {
	Key Apple
	Val string
}

// TX1AppleStringTuples can be used as a builder for unmodifiable maps.
type TX1AppleStringTuples []TX1AppleStringTuple

func (ts TX1AppleStringTuples) Append1(k Apple, v string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k, v})
}

func (ts TX1AppleStringTuples) Append2(k1 Apple, v1 string, k2 Apple, v2 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2})
}

func (ts TX1AppleStringTuples) Append3(k1 Apple, v1 string, k2 Apple, v2 string, k3 Apple, v3 string) TX1AppleStringTuples {
	return append(ts, TX1AppleStringTuple{k1, v1}, TX1AppleStringTuple{k2, v2}, TX1AppleStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1AppleStringMap() TX1AppleStringMap {
	return TX1AppleStringMap(make(map[Apple]string))
}

// NewTX1AppleStringMap creates and returns a reference to a map containing one item.
func NewTX1AppleStringMap1(k Apple, v string) TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	mm[k] = v
	return mm
}

// NewTX1AppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1AppleStringMap(kv ...TX1AppleStringTuple) TX1AppleStringMap {
	mm := newTX1AppleStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1AppleStringMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1AppleStringMap) ToSlice() []TX1AppleStringTuple {
	var s []TX1AppleStringTuple
	for k, v := range mm {
		s = append(s, TX1AppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1AppleStringMap) Get(k Apple) (string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1AppleStringMap) Put(k Apple, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1AppleStringMap) ContainsKey(k Apple) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1AppleStringMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1AppleStringMap) Clear() {
	*mm = make(map[Apple]string)
}

// Remove allows the removal of a single item from the map.
func (mm TX1AppleStringMap) Remove(k Apple) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1AppleStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX1AppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1AppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1AppleStringMap) Forall(fn func(Apple, string) bool) bool {
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
func (mm TX1AppleStringMap) Exists(fn func(Apple, string) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1AppleStringMap) Filter(fn func(Apple, string) bool) TX1AppleStringMap {
	result := NewTX1AppleStringMap()
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
func (mm TX1AppleStringMap) Partition(fn func(Apple, string) bool) (matching TX1AppleStringMap, others TX1AppleStringMap) {
	matching = NewTX1AppleStringMap()
	others = NewTX1AppleStringMap()
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
func (mm TX1AppleStringMap) Clone() TX1AppleStringMap {
	result := NewTX1AppleStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


