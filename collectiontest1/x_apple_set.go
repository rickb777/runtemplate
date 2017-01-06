// Generated from set.tpl with Type=Apple
// options: Numeric=<no value> Ordered=<no value> Stringer=false Mutable=true

package collectiontest1


// Stringer is not supported.


// AppleSet is the primary type that represents a set
type AppleSet map[Apple]struct{}

// NewAppleSet creates and returns a reference to an empty set.
func NewAppleSet(a ...Apple) AppleSet {
	set := make(AppleSet)
	for _, i := range a {
		set[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set AppleSet) ToSlice() []Apple {
	var s []Apple
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set AppleSet) Clone() AppleSet {
	clonedSet := NewAppleSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set AppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set AppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set AppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set AppleSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set AppleSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set AppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------


// Add adds an item to the current set if it doesn't already exist in the set.
func (set AppleSet) Add(i Apple) bool {
	_, found := set[i]
	set.doAdd(i)
	return !found //False if it existed already
}


func (set AppleSet) doAdd(i Apple) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set AppleSet) Contains(i Apple) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set AppleSet) ContainsAll(i ...Apple) bool {
	for _, v := range i {
		_, found := set[v]
		if !found {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set AppleSet) IsSubset(other AppleSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set AppleSet) IsSuperset(other AppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set AppleSet) Append(more ...Apple) AppleSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set AppleSet) Union(other AppleSet) AppleSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set AppleSet) Intersect(other AppleSet) AppleSet {
	intersection := NewAppleSet()
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
func (set AppleSet) Difference(other AppleSet) AppleSet {
	differencedSet := NewAppleSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set AppleSet) SymmetricDifference(other AppleSet) AppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}


// Clear clears the entire set to be the empty set.
func (set *AppleSet) Clear() {
	*set = NewAppleSet()
}

// Remove allows the removal of a single item from the set.
func (set AppleSet) Remove(i Apple) {
	delete(set, i)
}


//-------------------------------------------------------------------------------------------------

// Send returns a channel of type Apple that you can range over.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set AppleSet) Send() <-chan Apple {
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
func (set AppleSet) Forall(fn func(Apple) bool) bool {
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
func (set AppleSet) Exists(fn func(Apple) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the passed func against each element.
func (set AppleSet) Foreach(fn func(Apple)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new AppleSet whose elements return true for func.
func (set AppleSet) Filter(fn func(Apple) bool) AppleSet {
	result := NewAppleSet()
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
func (set AppleSet) Partition(p func(Apple) bool) (AppleSet, AppleSet) {
	matching := NewAppleSet()
	others := NewAppleSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of AppleSet that return true for the passed predicate.
func (set AppleSet) CountBy(predicate func(Apple) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of AppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set AppleSet) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of AppleSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set AppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
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
func (set AppleSet) Equals(other AppleSet) bool {
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


