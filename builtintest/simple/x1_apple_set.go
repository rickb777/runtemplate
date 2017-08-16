// A simple type derived from map[Apple]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=Apple
// options: Numeric:<no value> Stringer:false Mutable:always

package simple

// X1AppleSet is the primary type that represents a set
type X1AppleSet map[Apple]struct{}

// NewX1AppleSet creates and returns a reference to an empty set.
func NewX1AppleSet(values ...Apple) X1AppleSet {
	set := make(X1AppleSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertX1AppleSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertX1AppleSet(values ...interface{}) (X1AppleSet, bool) {
	set := make(X1AppleSet)
	for _, i := range values {
		v, ok := i.(Apple)
		if !ok {
		    return set, false
		}
		set[v] = struct{}{}
	}
	return set, true
}

// BuildX1AppleSetFromChan constructs a new X1AppleSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1AppleSetFromChan(source <-chan Apple) X1AppleSet {
	set := make(X1AppleSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set X1AppleSet) ToSlice() []Apple {
	var s []Apple
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set X1AppleSet) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set X1AppleSet) Clone() X1AppleSet {
	clonedSet := NewX1AppleSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X1AppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X1AppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set X1AppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set X1AppleSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X1AppleSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X1AppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set X1AppleSet) Add(i ...Apple) X1AppleSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set X1AppleSet) doAdd(i Apple) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set X1AppleSet) Contains(i Apple) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set X1AppleSet) ContainsAll(i ...Apple) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set X1AppleSet) IsSubset(other X1AppleSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set X1AppleSet) IsSuperset(other X1AppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set X1AppleSet) Append(more ...Apple) X1AppleSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set X1AppleSet) Union(other X1AppleSet) X1AppleSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X1AppleSet) Intersect(other X1AppleSet) X1AppleSet {
	intersection := NewX1AppleSet()
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
func (set X1AppleSet) Difference(other X1AppleSet) X1AppleSet {
	differencedSet := NewX1AppleSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X1AppleSet) SymmetricDifference(other X1AppleSet) X1AppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *X1AppleSet) Clear() {
	*set = NewX1AppleSet()
}

// Remove allows the removal of a single item from the set.
func (set X1AppleSet) Remove(i Apple) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X1AppleSet) Send() <-chan Apple {
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
func (set X1AppleSet) Forall(fn func(Apple) bool) bool {
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
func (set X1AppleSet) Exists(fn func(Apple) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the passed func against each element.
func (set X1AppleSet) Foreach(fn func(Apple)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new X1AppleSet whose elements return true for func.
func (set X1AppleSet) Filter(fn func(Apple) bool) X1AppleSet {
	result := NewX1AppleSet()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set X1AppleSet) Partition(p func(Apple) bool) (X1AppleSet, X1AppleSet) {
	matching := NewX1AppleSet()
	others := NewX1AppleSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of X1AppleSet that return true for the passed predicate.
func (set X1AppleSet) CountBy(predicate func(Apple) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1AppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set X1AppleSet) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of X1AppleSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set X1AppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
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
func (set X1AppleSet) Equals(other X1AppleSet) bool {
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


