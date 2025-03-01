// A simple type derived from map[int]struct{}
//
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=int
// options: Numeric:<no value> Integer:true Stringer:true Mutable:always
// by runtemplate v3.10.2
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// SimpleIntSet is the primary type that represents a set
type SimpleIntSet map[int]struct{}

// NewSimpleIntSet creates and returns a reference to an empty set.
func NewSimpleIntSet(values ...int) SimpleIntSet {
	set := make(SimpleIntSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertSimpleIntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertSimpleIntSet(values ...interface{}) (SimpleIntSet, bool) {
	set := make(SimpleIntSet)

	for _, i := range values {
		switch s := i.(type) {
		case string:
			k, e := strconv.ParseInt(s, 10, 64)
			if e == nil {
				i = k
			}
		case *string:
			k, e := strconv.ParseInt(*s, 10, 64)
			if e == nil {
				i = k
			}
		}
		switch j := i.(type) {
		case int:
			k := int(j)
			set[k] = struct{}{}
		case *int:
			k := int(*j)
			set[k] = struct{}{}
		case int8:
			k := int(j)
			set[k] = struct{}{}
		case *int8:
			k := int(*j)
			set[k] = struct{}{}
		case int16:
			k := int(j)
			set[k] = struct{}{}
		case *int16:
			k := int(*j)
			set[k] = struct{}{}
		case int32:
			k := int(j)
			set[k] = struct{}{}
		case *int32:
			k := int(*j)
			set[k] = struct{}{}
		case int64:
			k := int(j)
			set[k] = struct{}{}
		case *int64:
			k := int(*j)
			set[k] = struct{}{}
		case uint:
			k := int(j)
			set[k] = struct{}{}
		case *uint:
			k := int(*j)
			set[k] = struct{}{}
		case uint8:
			k := int(j)
			set[k] = struct{}{}
		case *uint8:
			k := int(*j)
			set[k] = struct{}{}
		case uint16:
			k := int(j)
			set[k] = struct{}{}
		case *uint16:
			k := int(*j)
			set[k] = struct{}{}
		case uint32:
			k := int(j)
			set[k] = struct{}{}
		case *uint32:
			k := int(*j)
			set[k] = struct{}{}
		case uint64:
			k := int(j)
			set[k] = struct{}{}
		case *uint64:
			k := int(*j)
			set[k] = struct{}{}
		case float32:
			k := int(j)
			set[k] = struct{}{}
		case *float32:
			k := int(*j)
			set[k] = struct{}{}
		case float64:
			k := int(j)
			set[k] = struct{}{}
		case *float64:
			k := int(*j)
			set[k] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildSimpleIntSetFromChan constructs a new SimpleIntSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildSimpleIntSetFromChan(source <-chan int) SimpleIntSet {
	set := make(SimpleIntSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set SimpleIntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set SimpleIntSet) IsSet() bool {
	return true
}

// ToSet returns the set; this is an identity operation in this case.
func (set SimpleIntSet) ToSet() SimpleIntSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set SimpleIntSet) ToSlice() []int {
	s := make([]int, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set SimpleIntSet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set SimpleIntSet) Clone() SimpleIntSet {
	clonedSet := NewSimpleIntSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set SimpleIntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set SimpleIntSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set SimpleIntSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set SimpleIntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set SimpleIntSet) Add(more ...int) SimpleIntSet {
	for _, v := range more {
		set.doAdd(v)
	}
	return set
}

func (set SimpleIntSet) doAdd(i int) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set SimpleIntSet) Contains(i int) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set SimpleIntSet) ContainsAll(i ...int) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set SimpleIntSet) IsSubset(other SimpleIntSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set SimpleIntSet) IsSuperset(other SimpleIntSet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set SimpleIntSet) Append(more ...int) SimpleIntSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set SimpleIntSet) Union(other SimpleIntSet) SimpleIntSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set SimpleIntSet) Intersect(other SimpleIntSet) SimpleIntSet {
	intersection := NewSimpleIntSet()
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
func (set SimpleIntSet) Difference(other SimpleIntSet) SimpleIntSet {
	differencedSet := NewSimpleIntSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set SimpleIntSet) SymmetricDifference(other SimpleIntSet) SimpleIntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *SimpleIntSet) Clear() {
	*set = NewSimpleIntSet()
}

// Remove a single item from the set.
func (set SimpleIntSet) Remove(i int) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set SimpleIntSet) Send() <-chan int {
	ch := make(chan int)
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
func (set SimpleIntSet) Forall(p func(int) bool) bool {
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
func (set SimpleIntSet) Exists(p func(int) bool) bool {
	for v := range set {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set SimpleIntSet) Foreach(f func(int)) {
	for v := range set {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set SimpleIntSet) Find(p func(int) bool) (int, bool) {

	for v := range set {
		if p(v) {
			return v, true
		}
	}

	var empty int
	return empty, false
}

// Filter returns a new SimpleIntSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set SimpleIntSet) Filter(p func(int) bool) SimpleIntSet {
	result := NewSimpleIntSet()
	for v := range set {
		if p(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new SimpleIntSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set SimpleIntSet) Partition(p func(int) bool) (SimpleIntSet, SimpleIntSet) {
	matching := NewSimpleIntSet()
	others := NewSimpleIntSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new SimpleIntSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleIntSet) Map(f func(int) int) SimpleIntSet {
	result := NewSimpleIntSet()

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
func (set SimpleIntSet) MapToString(f func(int) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v))
	}

	return result
}

// MapToInt64 returns a new []int64 by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleIntSet) MapToInt64(f func(int) int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set))
	for v := range set {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new SimpleIntSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleIntSet) FlatMap(f func(int) []int) SimpleIntSet {
	result := NewSimpleIntSet()

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
func (set SimpleIntSet) FlatMapToString(f func(int) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v)...)
	}

	return result
}

// FlatMapToInt64 returns a new []int64 by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set SimpleIntSet) FlatMapToInt64(f func(int) []int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set))
	for v := range set {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of SimpleIntSet that return true for the predicate p.
func (set SimpleIntSet) CountBy(p func(int) bool) (result int) {
	for v := range set {
		if p(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set SimpleIntSet) Min() int {
	v := set.MinBy(func(a int, b int) bool {
		return a < b
	})
	return v
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set SimpleIntSet) Max() int {
	v := set.MaxBy(func(a int, b int) bool {
		return a < b
	})
	return v
}

// Fold aggregates all the values in the set using a supplied function, starting from some initial value.
func (set SimpleIntSet) Fold(initial int, fn func(int, int) int) int {
	m := initial
	for v := range set {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of SimpleIntSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set SimpleIntSet) MinBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m int
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

// MaxBy returns an element of SimpleIntSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set SimpleIntSet) MaxBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m int
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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the set.
func (set SimpleIntSet) Sum() int {
	sum := int(0)
	for v := range set {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set SimpleIntSet) Equals(other SimpleIntSet) bool {
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

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set SimpleIntSet) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set SimpleIntSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set SimpleIntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set SimpleIntSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set SimpleIntSet) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set SimpleIntSet) UnmarshalJSON(b []byte) error {
	values := make([]int, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set SimpleIntSet) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set SimpleIntSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
