// An encapsulated immutable map[Apple]struct{} used as a set.
// Thread-safe.
//
//
// Generated from immutable/set.tpl with Type=Apple
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:false Mutable:disabled
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

// X1AppleSet is the primary type that represents a set.
type X1AppleSet struct {
	m map[Apple]struct{}
}

// NewX1AppleSet creates and returns a reference to an empty set.
func NewX1AppleSet(values ...Apple) *X1AppleSet {
	set := &X1AppleSet{
		m: make(map[Apple]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertX1AppleSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertX1AppleSet(values ...interface{}) (*X1AppleSet, bool) {
	set := NewX1AppleSet()

	for _, i := range values {
		switch j := i.(type) {
		case Apple:
			set.m[j] = struct{}{}
		case *Apple:
			set.m[*j] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildX1AppleSetFromChan constructs a new X1AppleSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1AppleSetFromChan(source <-chan Apple) *X1AppleSet {
	set := NewX1AppleSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *X1AppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *X1AppleSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *X1AppleSet) ToList() *X1AppleList {
	if set == nil {
		return nil
	}

	return &X1AppleList{
		m: set.ToSlice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *X1AppleSet) ToSet() *X1AppleSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *X1AppleSet) ToSlice() []Apple {
	if set == nil {
		return nil
	}

	s := make([]Apple, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *X1AppleSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set *X1AppleSet) Clone() *X1AppleSet {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *X1AppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *X1AppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *X1AppleSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *X1AppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set *X1AppleSet) Add(more ...Apple) *X1AppleSet {
	newSet := NewX1AppleSet()

	for v := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set *X1AppleSet) doAdd(i Apple) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *X1AppleSet) Contains(i Apple) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set *X1AppleSet) ContainsAll(i ...Apple) bool {
	if set == nil {
		return false
	}

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *X1AppleSet) IsSubset(other *X1AppleSet) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *X1AppleSet) IsSuperset(other *X1AppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *X1AppleSet) Union(other *X1AppleSet) *X1AppleSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := NewX1AppleSet()

	for v := range set.m {
		unionedSet.doAdd(v)
	}

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *X1AppleSet) Intersect(other *X1AppleSet) *X1AppleSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewX1AppleSet()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *X1AppleSet) Difference(other *X1AppleSet) *X1AppleSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewX1AppleSet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *X1AppleSet) SymmetricDifference(other *X1AppleSet) *X1AppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set *X1AppleSet) Remove(i Apple) *X1AppleSet {
	if set == nil {
		return nil
	}

	clonedSet := NewX1AppleSet()

	for v := range set.m {
		if i != v {
			clonedSet.doAdd(v)
		}
	}

	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *X1AppleSet) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {
		if set != nil {
			for v := range set.m {
				ch <- v
			}
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
func (set *X1AppleSet) Forall(p func(Apple) bool) bool {
	if set == nil {
		return true
	}

	for v := range set.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *X1AppleSet) Exists(p func(Apple) bool) bool {
	if set == nil {
		return false
	}

	for v := range set.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the function f against each element.
func (set *X1AppleSet) Foreach(f func(Apple)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *X1AppleSet) Find(p func(Apple) bool) (Apple, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false
}

// Filter returns a new X1AppleSet whose elements return true for the predicate p.
func (set *X1AppleSet) Filter(p func(Apple) bool) *X1AppleSet {
	if set == nil {
		return nil
	}

	result := NewX1AppleSet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new AppleSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *X1AppleSet) Partition(p func(Apple) bool) (*X1AppleSet, *X1AppleSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewX1AppleSet()
	others := NewX1AppleSet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new X1AppleSet by transforming every element with a function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X1AppleSet) Map(f func(Apple) Apple) *X1AppleSet {
	if set == nil {
		return nil
	}

	result := NewX1AppleSet()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new X1AppleSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X1AppleSet) FlatMap(f func(Apple) []Apple) *X1AppleSet {
	if set == nil {
		return nil
	}

	result := NewX1AppleSet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of X1AppleSet that return true for the predicate p.
func (set *X1AppleSet) CountBy(p func(Apple) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// Fold aggregates all the values in the set using a supplied function, starting from some initial value.
func (set *X1AppleSet) Fold(initial Apple, fn func(Apple, Apple) Apple) Apple {

	m := initial
	for v := range set.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of X1AppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *X1AppleSet) MinBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m Apple
	first := true
	for v := range set.m {
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
func (set *X1AppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m Apple
	first := true
	for v := range set.m {
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
func (set *X1AppleSet) Equals(other *X1AppleSet) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	if set.Size() != other.Size() {
		return false
	}

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}
