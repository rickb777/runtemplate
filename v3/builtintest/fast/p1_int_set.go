// An encapsulated map[*int]struct{} used as a set.
// Note that the api uses *int but the set uses <no value> keys.
// Not thread-safe.
//
// Generated from fast/set.tpl with Type=*int
// options: Comparable:always Numeric:true Ordered:true Stringer:true ToList:true
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// P1IntSet is the primary type that represents a set.
type P1IntSet struct {
	m map[int]struct{}
}

// NewP1IntSet creates and returns a reference to an empty set.
func NewP1IntSet(values ...*int) *P1IntSet {
	set := &P1IntSet{
		m: make(map[int]struct{}),
	}
	for _, i := range values {
		set.m[*i] = struct{}{}
	}
	return set
}

// ConvertP1IntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertP1IntSet(values ...interface{}) (*P1IntSet, bool) {
	set := NewP1IntSet()

	for _, i := range values {
		switch j := i.(type) {
		case int:
			k := int(j)
			set.m[k] = struct{}{}
		case *int:
			k := int(*j)
			set.m[k] = struct{}{}
		case int8:
			k := int(j)
			set.m[k] = struct{}{}
		case *int8:
			k := int(*j)
			set.m[k] = struct{}{}
		case int16:
			k := int(j)
			set.m[k] = struct{}{}
		case *int16:
			k := int(*j)
			set.m[k] = struct{}{}
		case int32:
			k := int(j)
			set.m[k] = struct{}{}
		case *int32:
			k := int(*j)
			set.m[k] = struct{}{}
		case int64:
			k := int(j)
			set.m[k] = struct{}{}
		case *int64:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint8:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint8:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint16:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint16:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint32:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint32:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint64:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint64:
			k := int(*j)
			set.m[k] = struct{}{}
		case float32:
			k := int(j)
			set.m[k] = struct{}{}
		case *float32:
			k := int(*j)
			set.m[k] = struct{}{}
		case float64:
			k := int(j)
			set.m[k] = struct{}{}
		case *float64:
			k := int(*j)
			set.m[k] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildP1IntSetFromChan constructs a new P1IntSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildP1IntSetFromChan(source <-chan *int) *P1IntSet {
	set := NewP1IntSet()
	for v := range source {
		set.m[*v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *P1IntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *P1IntSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *P1IntSet) ToList() *P1IntList {
	if set == nil {
		return nil
	}

	return &P1IntList{
		m: set.slice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *P1IntSet) ToSet() *P1IntSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *P1IntSet) slice() []*int {
	if set == nil {
		return nil
	}

	s := make([]*int, 0, len(set.m))
	for v := range set.m {
		s = append(s, &v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *P1IntSet) ToSlice() []*int {

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *P1IntSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *P1IntSet) Clone() *P1IntSet {
	if set == nil {
		return nil
	}

	clonedSet := NewP1IntSet()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *P1IntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *P1IntSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *P1IntSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *P1IntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *P1IntSet) Add(more ...*int) {

	for _, v := range more {
		set.doAdd(*v)
	}
}

func (set *P1IntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *P1IntSet) Contains(i *int) bool {
	if set == nil {
		return false
	}

	_, found := set.m[*i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *P1IntSet) ContainsAll(i ...*int) bool {
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
func (set *P1IntSet) IsSubset(other *P1IntSet) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	for v := range set.m {
		if !other.Contains(&v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *P1IntSet) IsSuperset(other *P1IntSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *P1IntSet) Union(other *P1IntSet) *P1IntSet {
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
func (set *P1IntSet) Intersect(other *P1IntSet) *P1IntSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewP1IntSet()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(&v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(&v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *P1IntSet) Difference(other *P1IntSet) *P1IntSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewP1IntSet()

	for v := range set.m {
		if !other.Contains(&v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *P1IntSet) SymmetricDifference(other *P1IntSet) *P1IntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *P1IntSet) Clear() {
	if set != nil {

		set.m = make(map[int]struct{})
	}
}

// Remove a single item from the set.
func (set *P1IntSet) Remove(i *int) {

	delete(set.m, *i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *P1IntSet) Send() <-chan *int {
	ch := make(chan *int)
	go func() {
		if set != nil {

			for v := range set.m {
				ch <- &v
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
func (set *P1IntSet) Forall(p func(*int) bool) bool {
	if set == nil {
		return true
	}

	for v := range set.m {
		if !p(&v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *P1IntSet) Exists(p func(*int) bool) bool {
	if set == nil {
		return false
	}

	for v := range set.m {
		if p(&v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *P1IntSet) Foreach(f func(*int)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(&v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *P1IntSet) Find(p func(*int) bool) (*int, bool) {

	for v := range set.m {
		if p(&v) {
			return &v, true
		}
	}

	return nil, false
}

// Filter returns a new P1IntSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *P1IntSet) Filter(p func(*int) bool) *P1IntSet {
	if set == nil {
		return nil
	}

	result := NewP1IntSet()

	for v := range set.m {
		if p(&v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new P1IntSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *P1IntSet) Partition(p func(*int) bool) (*P1IntSet, *P1IntSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewP1IntSet()
	others := NewP1IntSet()

	for v := range set.m {
		if p(&v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new P1IntSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1IntSet) Map(f func(*int) *int) *P1IntSet {
	if set == nil {
		return nil
	}

	result := NewP1IntSet()

	for v := range set.m {
	    k := f(&v)
		result.m[*k] = struct{}{}
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1IntSet) MapToString(f func(*int) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(&v))
	}

	return result
}

// MapToInt64 returns a new []int64 by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1IntSet) MapToInt64(f func(*int) int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(&v))
	}

	return result
}

// FlatMap returns a new P1IntSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1IntSet) FlatMap(f func(*int) []*int) *P1IntSet {
	if set == nil {
		return nil
	}

	result := NewP1IntSet()

	for v := range set.m {
		for _, x := range f(&v) {
			result.m[*x] = struct{}{}
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
func (set *P1IntSet) FlatMapToString(f func(*int) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(&v)...)
	}

	return result
}

// FlatMapToInt64 returns a new []int64 by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1IntSet) FlatMapToInt64(f func(*int) []int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(&v)...)
	}

	return result
}

// CountBy gives the number elements of P1IntSet that return true for the predicate p.
func (set *P1IntSet) CountBy(p func(*int) bool) (result int) {

	for v := range set.m {
		if p(&v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *P1IntSet) Min() int {

	var m int
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
func (set *P1IntSet) Max() (result int) {

	var m int
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

// MinBy returns an element of P1IntSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *P1IntSet) MinBy(less func(*int, *int) bool) *int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(&v, &m) {
			m = v
		}
	}
	return &m
}

// MaxBy returns an element of P1IntSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *P1IntSet) MaxBy(less func(*int, *int) bool) *int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(&m, &v) {
			m = v
		}
	}
	return &m
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the set.
func (set *P1IntSet) Sum() int {

	sum := int(0)
	for v := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *P1IntSet) Equals(other *P1IntSet) bool {
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
		if !other.Contains(&v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// StringSet gets a list of strings that depicts all the elements.
func (set *P1IntSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *P1IntSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *P1IntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *P1IntSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *P1IntSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (set *P1IntSet) UnmarshalJSON(b []byte) error {

	values := make([]*int, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewP1IntSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *P1IntSet) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *P1IntSet) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this set type.
// You must register *int with the 'gob' package before this method is used.
func (set *P1IntSet) GobDecode(b []byte) error {

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register *int with the 'gob' package before this method is used.
func (set P1IntSet) GobEncode() ([]byte, error) {

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
