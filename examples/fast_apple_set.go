// An encapsulated map[Apple]struct{} used as a set.
//
// Generated from fast/set.tpl with Type=Apple
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:false

package examples

import (
	"bytes"
	"encoding/gob"
)

// FastAppleSet is the primary type that represents a set
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
		v, ok := i.(Apple)
		if ok {
			set.m[v] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildFastAppleSetFromChan constructs a new FastAppleSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildFastAppleSetFromChan(source <-chan Apple) *FastAppleSet {
	set := NewFastAppleSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *FastAppleSet) ToSlice() []Apple {

	var s []Apple
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *FastAppleSet) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set *FastAppleSet) Clone() *FastAppleSet {
	clonedSet := NewFastAppleSet()

	for v, _ := range set.m {
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

// IsSequence returns true for lists.
func (set *FastAppleSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set *FastAppleSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *FastAppleSet) Size() int {

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

// Contains determines if a given item is already in the set.
func (set *FastAppleSet) Contains(i Apple) bool {

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set *FastAppleSet) ContainsAll(i ...Apple) bool {

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set *FastAppleSet) IsSubset(other *FastAppleSet) bool {

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set *FastAppleSet) IsSuperset(other *FastAppleSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *FastAppleSet) Union(other *FastAppleSet) *FastAppleSet {
	unionedSet := set.Clone()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *FastAppleSet) Intersect(other *FastAppleSet) *FastAppleSet {
	intersection := NewFastAppleSet()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v, _ := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v, _ := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *FastAppleSet) Difference(other *FastAppleSet) *FastAppleSet {
	differencedSet := NewFastAppleSet()

	for v, _ := range set.m {
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

// Clear clears the entire set to be the empty set.
func (set *FastAppleSet) Clear() {

	set.m = make(map[Apple]struct{})
}

// Remove removes a single item from the set.
func (set *FastAppleSet) Remove(i Apple) {

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *FastAppleSet) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {

		for v, _ := range set.m {
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
func (set *FastAppleSet) Forall(fn func(Apple) bool) bool {

	for v, _ := range set.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *FastAppleSet) Exists(fn func(Apple) bool) bool {

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over AppleSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set *FastAppleSet) Foreach(fn func(Apple)) {

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (set *FastAppleSet) Find(fn func(Apple) bool) (Apple, bool) {

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false

}

// Filter returns a new FastAppleSet whose elements return true for func.
//
// The original set is not modified
func (set *FastAppleSet) Filter(fn func(Apple) bool) *FastAppleSet {
	result := NewFastAppleSet()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *FastAppleSet) Partition(p func(Apple) bool) (*FastAppleSet, *FastAppleSet) {
	matching := NewFastAppleSet()
	others := NewFastAppleSet()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new FastAppleSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastAppleSet) Map(fn func(Apple) Apple) *FastAppleSet {
	result := NewFastAppleSet()

	for v, _ := range set.m {
		result.m[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new FastAppleSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastAppleSet) FlatMap(fn func(Apple) []Apple) *FastAppleSet {
	result := NewFastAppleSet()

	for v, _ := range set.m {
		for _, x := range fn(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of FastAppleSet that return true for the passed predicate.
func (set *FastAppleSet) CountBy(predicate func(Apple) bool) (result int) {

	for v, _ := range set.m {
		if predicate(v) {
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
	for v, _ := range set.m {
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
	for v, _ := range set.m {
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
func (set *FastAppleSet) Equals(other *FastAppleSet) bool {

	if set.Size() != other.Size() {
		return false
	}
	for v, _ := range set.m {
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

// GobDecode implements 'gob' encoding for this set type.
// You must register Apple with the 'gob' package before this method is used.
func (set FastAppleSet) GobEncode() ([]byte, error) {

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
