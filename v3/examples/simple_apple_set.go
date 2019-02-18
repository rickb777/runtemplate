// A simple type derived from map[Apple]struct{}
//
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=Apple
// options: Numeric:<no value> Stringer:false Mutable:always
// by runtemplate v3.3.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

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
		switch j := i.(type) {
		case Apple:
			set[j] = struct{}{}
		case *Apple:
			set[*j] = struct{}{}
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

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set SimpleAppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set SimpleAppleSet) IsSet() bool {
	return true
}

// ToSet returns the set; this is an identity operation in this case.
func (set SimpleAppleSet) ToSet() SimpleAppleSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set SimpleAppleSet) ToSlice() []Apple {
	s := make([]Apple, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set SimpleAppleSet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
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
func (set SimpleAppleSet) Add(more ...Apple) SimpleAppleSet {
	for _, v := range more {
		set.doAdd(v)
	}
	return set
}

func (set SimpleAppleSet) doAdd(i Apple) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set SimpleAppleSet) Contains(i Apple) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set SimpleAppleSet) ContainsAll(i ...Apple) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set SimpleAppleSet) IsSubset(other SimpleAppleSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set SimpleAppleSet) IsSuperset(other SimpleAppleSet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
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

// Clear the entire set. Aterwards, it will be an empty set.
func (set *SimpleAppleSet) Clear() {
	*set = NewSimpleAppleSet()
}

// Remove a single item from the set.
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

// Forall applies a predicate function p to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set SimpleAppleSet) Forall(p func(Apple) bool) bool {
	for v := range set {
		if !p(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set SimpleAppleSet) Exists(p func(Apple) bool) bool {
	for v := range set {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set SimpleAppleSet) Foreach(f func(Apple)) {
	for v := range set {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set SimpleAppleSet) Find(p func(Apple) bool) (Apple, bool) {

	for v := range set {
		if p(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false
}

// Filter returns a new SimpleAppleSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set SimpleAppleSet) Filter(p func(Apple) bool) SimpleAppleSet {
	result := NewSimpleAppleSet()
	for v := range set {
		if p(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new SimpleAppleSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
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

// Map returns a new SimpleAppleSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) Map(f func(Apple) Apple) SimpleAppleSet {
	result := NewSimpleAppleSet()

	for v := range set {
		k := f(v)
		result[k] = struct{}{}
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) MapToString(f func(Apple) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new SimpleAppleSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) FlatMap(f func(Apple) []Apple) SimpleAppleSet {
	result := NewSimpleAppleSet()

	for v := range set {
		for _, x := range f(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleAppleSet) FlatMapToString(f func(Apple) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of SimpleAppleSet that return true for the predicate p.
func (set SimpleAppleSet) CountBy(p func(Apple) bool) (result int) {
	for v := range set {
		if p(v) {
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
		panic("Cannot determine the minimum of an empty set.")
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
		panic("Cannot determine the minimum of an empty set.")
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

// Equals determines whether two sets are equal to each other, returning true if so.
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
