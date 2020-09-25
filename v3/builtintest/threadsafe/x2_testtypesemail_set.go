// An encapsulated map[testtypes.Email]struct{} used as a set.
//
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=testtypes.Email
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:<no value> ToList:<no value>
// by runtemplate v3.6.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"github.com/rickb777/runtemplate/v3/builtintest/testtypes"
	"sync"
)

// X2TesttypesEmailSet is the primary type that represents a set.
type X2TesttypesEmailSet struct {
	s *sync.RWMutex
	m map[testtypes.Email]struct{}
}

// NewX2TesttypesEmailSet creates and returns a reference to an empty set.
func NewX2TesttypesEmailSet(values ...testtypes.Email) *X2TesttypesEmailSet {
	set := &X2TesttypesEmailSet{
		s: &sync.RWMutex{},
		m: make(map[testtypes.Email]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertX2TesttypesEmailSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertX2TesttypesEmailSet(values ...interface{}) (*X2TesttypesEmailSet, bool) {
	set := NewX2TesttypesEmailSet()

	for _, i := range values {
		switch j := i.(type) {
		case testtypes.Email:
			k := testtypes.Email(j)
			set.m[k] = struct{}{}
		case *testtypes.Email:
			k := testtypes.Email(*j)
			set.m[k] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildX2TesttypesEmailSetFromChan constructs a new X2TesttypesEmailSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildX2TesttypesEmailSetFromChan(source <-chan testtypes.Email) *X2TesttypesEmailSet {
	set := NewX2TesttypesEmailSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *X2TesttypesEmailSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *X2TesttypesEmailSet) IsSet() bool {
	return true
}

// ToSet returns the set; this is an identity operation in this case.
func (set *X2TesttypesEmailSet) ToSet() *X2TesttypesEmailSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *X2TesttypesEmailSet) slice() []testtypes.Email {
	if set == nil {
		return nil
	}

	s := make([]testtypes.Email, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *X2TesttypesEmailSet) ToSlice() []testtypes.Email {
	set.s.RLock()
	defer set.s.RUnlock()

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *X2TesttypesEmailSet) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *X2TesttypesEmailSet) Clone() *X2TesttypesEmailSet {
	if set == nil {
		return nil
	}

	clonedSet := NewX2TesttypesEmailSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *X2TesttypesEmailSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *X2TesttypesEmailSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *X2TesttypesEmailSet) Size() int {
	if set == nil {
		return 0
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *X2TesttypesEmailSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *X2TesttypesEmailSet) Add(more ...testtypes.Email) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *X2TesttypesEmailSet) doAdd(i testtypes.Email) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *X2TesttypesEmailSet) Contains(i testtypes.Email) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *X2TesttypesEmailSet) ContainsAll(i ...testtypes.Email) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *X2TesttypesEmailSet) IsSubset(other *X2TesttypesEmailSet) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *X2TesttypesEmailSet) IsSuperset(other *X2TesttypesEmailSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *X2TesttypesEmailSet) Union(other *X2TesttypesEmailSet) *X2TesttypesEmailSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *X2TesttypesEmailSet) Intersect(other *X2TesttypesEmailSet) *X2TesttypesEmailSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewX2TesttypesEmailSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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
func (set *X2TesttypesEmailSet) Difference(other *X2TesttypesEmailSet) *X2TesttypesEmailSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewX2TesttypesEmailSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *X2TesttypesEmailSet) SymmetricDifference(other *X2TesttypesEmailSet) *X2TesttypesEmailSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *X2TesttypesEmailSet) Clear() {
	if set != nil {
		set.s.Lock()
		defer set.s.Unlock()

		set.m = make(map[testtypes.Email]struct{})
	}
}

// Remove a single item from the set.
func (set *X2TesttypesEmailSet) Remove(i testtypes.Email) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *X2TesttypesEmailSet) Send() <-chan testtypes.Email {
	ch := make(chan testtypes.Email)
	go func() {
		if set != nil {
			set.s.RLock()
			defer set.s.RUnlock()

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
func (set *X2TesttypesEmailSet) Forall(p func(testtypes.Email) bool) bool {
	if set == nil {
		return true
	}

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *X2TesttypesEmailSet) Exists(p func(testtypes.Email) bool) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *X2TesttypesEmailSet) Foreach(f func(testtypes.Email)) {
	if set == nil {
		return
	}

	set.s.Lock()
	defer set.s.Unlock()

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first testtypes.Email that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *X2TesttypesEmailSet) Find(p func(testtypes.Email) bool) (testtypes.Email, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty testtypes.Email
	return empty, false
}

// Filter returns a new X2TesttypesEmailSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *X2TesttypesEmailSet) Filter(p func(testtypes.Email) bool) *X2TesttypesEmailSet {
	if set == nil {
		return nil
	}

	result := NewX2TesttypesEmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new X2TesttypesEmailSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *X2TesttypesEmailSet) Partition(p func(testtypes.Email) bool) (*X2TesttypesEmailSet, *X2TesttypesEmailSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewX2TesttypesEmailSet()
	others := NewX2TesttypesEmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new X2TesttypesEmailSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X2TesttypesEmailSet) Map(f func(testtypes.Email) testtypes.Email) *X2TesttypesEmailSet {
	if set == nil {
		return nil
	}

	result := NewX2TesttypesEmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		k := f(v)
		result.m[k] = struct{}{}
	}

	return result
}

// FlatMap returns a new X2TesttypesEmailSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X2TesttypesEmailSet) FlatMap(f func(testtypes.Email) []testtypes.Email) *X2TesttypesEmailSet {
	if set == nil {
		return nil
	}

	result := NewX2TesttypesEmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of X2TesttypesEmailSet that return true for the predicate p.
func (set *X2TesttypesEmailSet) CountBy(p func(testtypes.Email) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2TesttypesEmailSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *X2TesttypesEmailSet) MinBy(less func(testtypes.Email, testtypes.Email) bool) testtypes.Email {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m testtypes.Email
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

// MaxBy returns an element of X2TesttypesEmailSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *X2TesttypesEmailSet) MaxBy(less func(testtypes.Email, testtypes.Email) bool) testtypes.Email {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m testtypes.Email
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
func (set *X2TesttypesEmailSet) Equals(other *X2TesttypesEmailSet) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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
