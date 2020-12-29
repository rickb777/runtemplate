// An encapsulated map[Name]struct{} used as a set.
//
// Not thread-safe.
//
// Generated from fast/set.tpl with Type=Name
// options: Comparable:always Numeric:<no value> Ordered:true Stringer:true ToList:true
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

import (
	"encoding/json"
	"fmt"
	"strings"
)

// X1NameSet is the primary type that represents a set.
type X1NameSet struct {
	m map[Name]struct{}
}

// NewX1NameSet creates and returns a reference to an empty set.
func NewX1NameSet(values ...Name) *X1NameSet {
	set := &X1NameSet{
		m: make(map[Name]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertX1NameSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertX1NameSet(values ...interface{}) (*X1NameSet, bool) {
	set := NewX1NameSet()

	for _, i := range values {
		switch j := i.(type) {
		case Name:
			k := Name(j)
			set.m[k] = struct{}{}
		case *Name:
			k := Name(*j)
			set.m[k] = struct{}{}
		default:
			if s, ok := i.(fmt.Stringer); ok {
				k := Name(s.String())
				set.m[k] = struct{}{}
			}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildX1NameSetFromChan constructs a new X1NameSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildX1NameSetFromChan(source <-chan Name) *X1NameSet {
	set := NewX1NameSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *X1NameSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *X1NameSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *X1NameSet) ToList() *X1NameList {
	if set == nil {
		return nil
	}

	return &X1NameList{
		m: set.slice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *X1NameSet) ToSet() *X1NameSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *X1NameSet) slice() []Name {
	if set == nil {
		return nil
	}

	s := make([]Name, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *X1NameSet) ToSlice() []Name {

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *X1NameSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *X1NameSet) Clone() *X1NameSet {
	if set == nil {
		return nil
	}

	clonedSet := NewX1NameSet()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *X1NameSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *X1NameSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *X1NameSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *X1NameSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *X1NameSet) Add(more ...Name) {

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *X1NameSet) doAdd(i Name) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *X1NameSet) Contains(i Name) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *X1NameSet) ContainsAll(i ...Name) bool {
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
func (set *X1NameSet) IsSubset(other *X1NameSet) bool {
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
func (set *X1NameSet) IsSuperset(other *X1NameSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *X1NameSet) Union(other *X1NameSet) *X1NameSet {
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
func (set *X1NameSet) Intersect(other *X1NameSet) *X1NameSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewX1NameSet()

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
func (set *X1NameSet) Difference(other *X1NameSet) *X1NameSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewX1NameSet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *X1NameSet) SymmetricDifference(other *X1NameSet) *X1NameSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *X1NameSet) Clear() {
	if set != nil {

		set.m = make(map[Name]struct{})
	}
}

// Remove a single item from the set.
func (set *X1NameSet) Remove(i Name) {

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *X1NameSet) Send() <-chan Name {
	ch := make(chan Name)
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
func (set *X1NameSet) Forall(p func(Name) bool) bool {
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
func (set *X1NameSet) Exists(p func(Name) bool) bool {
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
func (set *X1NameSet) Foreach(f func(Name)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Name that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *X1NameSet) Find(p func(Name) bool) (Name, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty Name
	return empty, false
}

// Filter returns a new X1NameSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *X1NameSet) Filter(p func(Name) bool) *X1NameSet {
	if set == nil {
		return nil
	}

	result := NewX1NameSet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new X1NameSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *X1NameSet) Partition(p func(Name) bool) (*X1NameSet, *X1NameSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewX1NameSet()
	others := NewX1NameSet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new X1NameSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X1NameSet) Map(f func(Name) Name) *X1NameSet {
	if set == nil {
		return nil
	}

	result := NewX1NameSet()

	for v := range set.m {
		k := f(v)
		result.m[k] = struct{}{}
	}

	return result
}

// FlatMap returns a new X1NameSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *X1NameSet) FlatMap(f func(Name) []Name) *X1NameSet {
	if set == nil {
		return nil
	}

	result := NewX1NameSet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of X1NameSet that return true for the predicate p.
func (set *X1NameSet) CountBy(p func(Name) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when Name is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *X1NameSet) Min() Name {

	var m Name
	first := true
	for v := range set.m {
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
func (set *X1NameSet) Max() (result Name) {

	var m Name
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}

// Fold aggregates all the values in the set using a supplied function, starting from some initial value.
func (set *X1NameSet) Fold(initial Name, fn func(Name, Name) Name) Name {

	m := initial
	for v := range set.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of X1NameSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *X1NameSet) MinBy(less func(Name, Name) bool) Name {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m Name
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

// MaxBy returns an element of X1NameSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *X1NameSet) MaxBy(less func(Name, Name) bool) Name {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m Name
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
func (set *X1NameSet) Equals(other *X1NameSet) bool {
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

// StringSet gets a list of strings that depicts all the elements.
func (set *X1NameSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *X1NameSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *X1NameSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *X1NameSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *X1NameSet) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set *X1NameSet) UnmarshalJSON(b []byte) error {

	values := make([]Name, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewX1NameSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *X1NameSet) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *X1NameSet) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
