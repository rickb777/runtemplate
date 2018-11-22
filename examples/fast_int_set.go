// An encapsulated map[int]struct{} used as a set.
//
// Generated from fast/set.tpl with Type=int
// options: Comparable:always Numeric:true Ordered:true Stringer:true
// by runtemplate v2.2.4
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// FastIntSet is the primary type that represents a set
type FastIntSet struct {
	m map[int]struct{}
}

// NewFastIntSet creates and returns a reference to an empty set.
func NewFastIntSet(values ...int) *FastIntSet {
	set := &FastIntSet{
		m: make(map[int]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertFastIntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertFastIntSet(values ...interface{}) (*FastIntSet, bool) {
	set := NewFastIntSet()

	for _, i := range values {
		switch i.(type) {
		case int:
			set.m[int(i.(int))] = struct{}{}
		case int8:
			set.m[int(i.(int8))] = struct{}{}
		case int16:
			set.m[int(i.(int16))] = struct{}{}
		case int32:
			set.m[int(i.(int32))] = struct{}{}
		case int64:
			set.m[int(i.(int64))] = struct{}{}
		case uint:
			set.m[int(i.(uint))] = struct{}{}
		case uint8:
			set.m[int(i.(uint8))] = struct{}{}
		case uint16:
			set.m[int(i.(uint16))] = struct{}{}
		case uint32:
			set.m[int(i.(uint32))] = struct{}{}
		case uint64:
			set.m[int(i.(uint64))] = struct{}{}
		case float32:
			set.m[int(i.(float32))] = struct{}{}
		case float64:
			set.m[int(i.(float64))] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildFastIntSetFromChan constructs a new FastIntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildFastIntSetFromChan(source <-chan int) *FastIntSet {
	set := NewFastIntSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *FastIntSet) slice() []int {
	if set == nil {
		return nil
	}

	var s []int
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *FastIntSet) ToSlice() []int {

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *FastIntSet) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *FastIntSet) Clone() *FastIntSet {
	if set == nil {
		return nil
	}

	clonedSet := NewFastIntSet()

	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *FastIntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *FastIntSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set *FastIntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set *FastIntSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *FastIntSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *FastIntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *FastIntSet) Add(more ...int) {

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *FastIntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set *FastIntSet) Contains(i int) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set *FastIntSet) ContainsAll(i ...int) bool {
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

// IsSubset determines if every item in the other set is in this set.
func (set *FastIntSet) IsSubset(other *FastIntSet) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set *FastIntSet) IsSuperset(other *FastIntSet) bool {
	if set.IsEmpty() {
		return other.IsEmpty()
	}

	if other.IsEmpty() {
		return true
	}

	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *FastIntSet) Union(other *FastIntSet) *FastIntSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *FastIntSet) Intersect(other *FastIntSet) *FastIntSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewFastIntSet()

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
func (set *FastIntSet) Difference(other *FastIntSet) *FastIntSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewFastIntSet()

	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *FastIntSet) SymmetricDifference(other *FastIntSet) *FastIntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *FastIntSet) Clear() {
	if set != nil {

		set.m = make(map[int]struct{})
	}
}

// Remove removes a single item from the set.
func (set *FastIntSet) Remove(i int) {

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *FastIntSet) Send() <-chan int {
	ch := make(chan int)
	go func() {
		if set != nil {

			for v, _ := range set.m {
				ch <- v
			}
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
func (set *FastIntSet) Forall(fn func(int) bool) bool {
	if set == nil {
		return true
	}

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
func (set *FastIntSet) Exists(fn func(int) bool) bool {
	if set == nil {
		return false
	}

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over intSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set *FastIntSet) Foreach(fn func(int)) {
	if set == nil {
		return
	}

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (set *FastIntSet) Find(fn func(int) bool) (int, bool) {

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}

	var empty int
	return empty, false

}

// Filter returns a new FastIntSet whose elements return true for func.
//
// The original set is not modified
func (set *FastIntSet) Filter(fn func(int) bool) *FastIntSet {
	if set == nil {
		return nil
	}

	result := NewFastIntSet()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *FastIntSet) Partition(p func(int) bool) (*FastIntSet, *FastIntSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewFastIntSet()
	others := NewFastIntSet()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new FastIntSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastIntSet) Map(fn func(int) int) *FastIntSet {
	if set == nil {
		return nil
	}

	result := NewFastIntSet()

	for v, _ := range set.m {
		result.m[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new FastIntSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *FastIntSet) FlatMap(fn func(int) []int) *FastIntSet {
	if set == nil {
		return nil
	}

	result := NewFastIntSet()

	for v, _ := range set.m {
		for _, x := range fn(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of FastIntSet that return true for the passed predicate.
func (set *FastIntSet) CountBy(predicate func(int) bool) (result int) {

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *FastIntSet) Min() int {

	var m int
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set *FastIntSet) Max() (result int) {

	var m int
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}

// MinBy returns an element of FastIntSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *FastIntSet) MinBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
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

// MaxBy returns an element of FastIntSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *FastIntSet) MaxBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the set.
func (set *FastIntSet) Sum() int {

	sum := int(0)
	for v, _ := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *FastIntSet) Equals(other *FastIntSet) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

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

// StringList gets a list of strings that depicts all the elements.
func (set *FastIntSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *FastIntSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *FastIntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *FastIntSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *FastIntSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for v, _ := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set *FastIntSet) UnmarshalJSON(b []byte) error {

	values := make([]int, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewFastIntSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *FastIntSet) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *FastIntSet) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v, _ := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
