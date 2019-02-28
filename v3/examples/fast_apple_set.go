// An encapsulated map[Apple]struct{} used as a set.
//
// Not thread-safe.
//
// Generated from fast/set.tpl with Type=Apple
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:false ToList:<no value>
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/gob"
)

// FastAppleSet is the primary type that represents a set.
type FastAppleSet struct {
	m map[Apple]struct{}
}

// NewFastAppleSet creates and returns a reference to an empty set.
func NewFastAppleSet(values ...Apple) *FastAppleSet {
	set := &FastAppleSet{
		m: make(map[Apple]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertFastAppleSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertFastAppleSet(values ...interface{}) (*FastAppleSet, bool) {
	set := NewFastAppleSet()

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

// BuildFastAppleSetFromChan constructs a new FastAppleSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildFastAppleSetFromChan(source <-chan Apple) *FastAppleSet {
	set := NewFastAppleSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *FastAppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *FastAppleSet) IsSet() bool {
	return true
}

// ToSet returns the set; this is an identity operation in this case.
func (set *FastAppleSet) ToSet() *FastAppleSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *FastAppleSet) slice() []Apple {
	if set == nil {
		return nil
	}

	s := make([]Apple, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *FastAppleSet) ToSlice() []Apple {

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *FastAppleSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *FastAppleSet) Clone() *FastAppleSet {
	if set == nil {
		return nil
	}

	clonedSet := NewFastAppleSet()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *FastAppleSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *FastAppleSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *FastAppleSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *FastAppleSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *FastAppleSet) Add(more ...Apple) {

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *FastAppleSet) doAdd(i Apple) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *FastAppleSet) Contains(i Apple) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *FastAppleSet) ContainsAll(i ...Apple) bool {
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
func (set *FastAppleSet) IsSubset(other *FastAppleSet) bool {
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
func (set *FastAppleSet) IsSuperset(other *FastAppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *FastAppleSet) Union(other *FastAppleSet) *FastAppleSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *FastAppleSet) Intersect(other *FastAppleSet) *FastAppleSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewFastAppleSet()

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
func (set *FastAppleSet) Difference(other *FastAppleSet) *FastAppleSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewFastAppleSet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *FastAppleSet) SymmetricDifference(other *FastAppleSet) *FastAppleSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *FastAppleSet) Clear() {
	if set != nil {

		set.m = make(map[Apple]struct{})
	}
}

// Remove a single item from the set.
func (set *FastAppleSet) Remove(i Apple) {

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *FastAppleSet) Send() <-chan Apple {
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
func (set *FastAppleSet) Forall(p func(Apple) bool) bool {
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
func (set *FastAppleSet) Exists(p func(Apple) bool) bool {
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

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *FastAppleSet) Foreach(f func(Apple)) {
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
func (set *FastAppleSet) Find(p func(Apple) bool) (Apple, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false
}

// Filter returns a new FastAppleSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *FastAppleSet) Filter(p func(Apple) bool) *FastAppleSet {
	if set == nil {
		return nil
	}

	result := NewFastAppleSet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new FastAppleSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *FastAppleSet) Partition(p func(Apple) bool) (*FastAppleSet, *FastAppleSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewFastAppleSet()
	others := NewFastAppleSet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new FastAppleSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastAppleSet) Map(f func(Apple) Apple) *FastAppleSet {
	if set == nil {
		return nil
	}

	result := NewFastAppleSet()

	for v := range set.m {
		k := f(v)
		result.m[k] = struct{}{}
	}

	return result
}

// FlatMap returns a new FastAppleSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastAppleSet) FlatMap(f func(Apple) []Apple) *FastAppleSet {
	if set == nil {
		return nil
	}

	result := NewFastAppleSet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of FastAppleSet that return true for the predicate p.
func (set *FastAppleSet) CountBy(p func(Apple) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of FastAppleSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *FastAppleSet) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of FastAppleSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *FastAppleSet) MaxBy(less func(Apple, Apple) bool) Apple {
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
func (set *FastAppleSet) Equals(other *FastAppleSet) bool {
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

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this set type.
// You must register Apple with the 'gob' package before this method is used.
func (set *FastAppleSet) GobDecode(b []byte) error {

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (set FastAppleSet) GobEncode() ([]byte, error) {

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
