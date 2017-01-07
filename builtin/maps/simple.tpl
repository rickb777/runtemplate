// Generated from {{.TemplateFile}} with Key={{.Key}} Type={{.Type}}
// options: Comparable={{.Comparable}} Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}}

package {{.Package}}

// {{.UKey}}{{.UType}}Map is the primary type that represents a thread-safe map
type {{.UKey}}{{.UType}}Map struct {
	m map[{{.Key}}]{{.Type}}
}

// {{.UKey}}{{.UType}}Tuple represents a key/value pair.
type {{.UKey}}{{.UType}}Tuple struct {
	Key {{.Key}}
	Val {{.Type}}
}

// New{{.UKey}}{{.UType}}Map creates and returns a reference to an empty map.
func New{{.UKey}}{{.UType}}Map(kv ...{{.UKey}}{{.UType}}Tuple) {{.UKey}}{{.UType}}Map {
	mm := {{.UKey}}{{.UType}}Map{
		m: make(map[{{.Key}}]{{.Type}}),
	}
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *{{.UKey}}{{.UType}}Map) Keys() []{{.Key}} {
	var s []{{.Key}}
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *{{.UKey}}{{.UType}}Map) ToSlice() []{{.UKey}}{{.UType}}Tuple {
	var s []{{.UKey}}{{.UType}}Tuple
	for k, v := range mm.m {
		s = append(s, {{.UKey}}{{.UType}}Tuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *{{.UKey}}{{.UType}}Map) Get(k {{.Key}}) ({{.Type}}, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *{{.UKey}}{{.UType}}Map) Put(k {{.Key}}, v {{.Type}}) bool {
	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *{{.UKey}}{{.UType}}Map) ContainsKey(k {{.Key}}) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *{{.UKey}}{{.UType}}Map) ContainsAllKeys(kk ...{{.Key}}) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *{{.UKey}}{{.UType}}Map) Clear() {
	mm.m = make(map[{{.Key}}]{{.Type}})
}

// Remove allows the removal of a single item from the map.
func (mm *{{.UKey}}{{.UType}}Map) Remove(k {{.Key}}) {
	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *{{.UKey}}{{.UType}}Map) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *{{.UKey}}{{.UType}}Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *{{.UKey}}{{.UType}}Map) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *{{.UKey}}{{.UType}}Map) Forall(fn func({{.Key}}, {{.Type}}) bool) bool {
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
func (mm *{{.UKey}}{{.UType}}Map) Exists(fn func({{.Key}}, {{.Type}}) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *{{.UKey}}{{.UType}}Map) Filter(fn func({{.Key}}, {{.Type}}) bool) {{.UKey}}{{.UType}}Map {
	result := New{{.UKey}}{{.UType}}Map()
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
func (mm *{{.UKey}}{{.UType}}Map) Partition(fn func({{.Key}}, {{.Type}}) bool) (matching {{.UKey}}{{.UType}}Map, others {{.UKey}}{{.UType}}Map) {
	matching = New{{.UKey}}{{.UType}}Map()
	others = New{{.UKey}}{{.UType}}Map()
	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (mm *{{.UKey}}{{.UType}}Map) Equals(other {{.UKey}}{{.UType}}Map) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *{{.UKey}}{{.UType}}Map) Clone() {{.UKey}}{{.UType}}Map {
	result := New{{.UKey}}{{.UType}}Map()
	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
