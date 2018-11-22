// A simple type derived from map[Apple]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=Apple
// options: Numeric:<no value> Stringer:false Mutable:always
// by runtemplate v2.2.4
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

// SimpleAppleSet is the primary type that represents a set
type SimpleAppleSet map[Apple]struct{}

// NewSimpleAppleSet creates and returns a reference to an empty set.
func NewSimpleAppleSet(values ...Apple) SimpleAppleSet {
	set := make(SimpleAppleSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertSimpleAppleSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertSimpleAppleSet(values ...interface{}) (SimpleAppleSet, bool) {
	set := make(SimpleAppleSet)

	for _, i := range values {
		v, ok := i.(Apple)
		if ok {
			set[v] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildSimpleAppleSetFromChan constructs a new SimpleAppleSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildSimpleAppleSetFromChan(source <-chan Apple) SimpleAppleSet {
	set := make(SimpleAppleSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set SimpleAppleSet) ToSlice() []Apple {
	var s []Apple
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set SimpleAppleSet) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v, _ := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set SimpleAppleSet) Clone() SimpleAppleSet {
	clonedSet := NewSimpleAppleSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set SimpleAppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set SimpleAppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set SimpleAppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set SimpleAppleSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set SimpleAppleSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set SimpleAppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set SimpleAppleSet) Add(i ...Apple) SimpleAppleSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set SimpleAppleSet) doAdd(i Apple) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set SimpleAppleSet) Contains(i Apple) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set SimpleAppleSet) ContainsAll(i ...Apple) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set SimpleAppleSet) IsSubset(other SimpleAppleSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set SimpleAppleSet) IsSuperset(other SimpleAppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set SimpleAppleSet) Append(more ...Apple) SimpleAppleSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set SimpleAppleSet) Union(other SimpleAppleSet) SimpleAppleSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set SimpleAppleSet) Intersect(other SimpleAppleSet) SimpleAppleSet {
	intersection := NewSimpleAppleSet()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set SimpleAppleSet) Difference(other SimpleAppleSet) SimpleAppleSet {
	differencedSet := NewSimpleAppleSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set SimpleAppleSet) SymmetricDifference(other SimpleAppleSet) SimpleAppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *SimpleAppleSet) Clear() {
	*set = NewSimpleAppleSet()
}

// Remove allows the removal of a single item from the set.
func (set SimpleAppleSet) Remove(i Apple) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set SimpleAppleSet) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set SimpleAppleSet) Forall(fn func(Apple) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set SimpleAppleSet) Exists(fn func(Apple) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the passed func against each element.
func (set SimpleAppleSet) Foreach(fn func(Apple)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new SimpleAppleSet whose elements return true for func.
// The original set is not modified
func (set SimpleAppleSet) Filter(fn func(Apple) bool) SimpleAppleSet {
	result := NewSimpleAppleSet()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new AppleSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
// The original set is not modified
func (set SimpleAppleSet) Partition(p func(Apple) bool) (SimpleAppleSet, SimpleAppleSet) {
	matching := NewSimpleAppleSet()
	others := NewSimpleAppleSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new SimpleAppleSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) Map(fn func(Apple) Apple) SimpleAppleSet {
	result := NewSimpleAppleSet()

	for v := range set {
		result[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new SimpleAppleSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) FlatMap(fn func(Apple) []Apple) SimpleAppleSet {
	result := NewSimpleAppleSet()

	for v, _ := range set {
		for _, x := range fn(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of SimpleAppleSet that return true for the passed predicate.
func (set SimpleAppleSet) CountBy(predicate func(Apple) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of SimpleAppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set SimpleAppleSet) MinBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m Apple
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of SimpleAppleSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set SimpleAppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m Apple
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set SimpleAppleSet) Equals(other SimpleAppleSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}
