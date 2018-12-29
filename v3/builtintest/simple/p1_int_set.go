// A simple type derived from map[int]struct{}
// Note that the api uses *int but the set uses <no value> keys.
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=int
// options: Numeric:true Stringer:true Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple


import (
	"bytes"
	"encoding/json"
	"fmt")

// P1IntSet is the primary type that represents a set
type P1IntSet map[int]struct{}

// NewP1IntSet creates and returns a reference to an empty set.
func NewP1IntSet(values ...*int) P1IntSet {
	set := make(P1IntSet)
	for _, i := range values {
		set[*i] = struct{}{}
	}
	return set
}

// ConvertP1IntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertP1IntSet(values ...interface{}) (P1IntSet, bool) {
	set := make(P1IntSet)

	for _, i := range values {
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

// BuildP1IntSetFromChan constructs a new P1IntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildP1IntSetFromChan(source <-chan *int) P1IntSet {
	set := make(P1IntSet)
	for v := range source {
		set[*v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set P1IntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set P1IntSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set P1IntSet) ToList() P1IntList {
	if set == nil {
		return nil
	}

	return P1IntList(set.ToSlice())
}

// ToSet returns the set; this is an identity operation in this case.
func (set P1IntSet) ToSet() P1IntSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set P1IntSet) ToSlice() []*int {
	s := make([]*int, 0, len(set))
	for v := range set {
		s = append(s, &v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set P1IntSet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set P1IntSet) Clone() P1IntSet {
	clonedSet := NewP1IntSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set P1IntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set P1IntSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set P1IntSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set P1IntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set P1IntSet) Add(more ...*int) P1IntSet {
	for _, v := range more {
		set.doAdd(*v)
	}
	return set
}

func (set P1IntSet) doAdd(i int) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set P1IntSet) Contains(i *int) bool {
	_, found := set[*i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set P1IntSet) ContainsAll(i ...*int) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set P1IntSet) IsSubset(other P1IntSet) bool {
	for v := range set {
		if !other.Contains(&v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set P1IntSet) IsSuperset(other P1IntSet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set P1IntSet) Append(more ...int) P1IntSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set P1IntSet) Union(other P1IntSet) P1IntSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set P1IntSet) Intersect(other P1IntSet) P1IntSet {
	intersection := NewP1IntSet()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(&v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(&v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set P1IntSet) Difference(other P1IntSet) P1IntSet {
	differencedSet := NewP1IntSet()
	for v := range set {
		if !other.Contains(&v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set P1IntSet) SymmetricDifference(other P1IntSet) P1IntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *P1IntSet) Clear() {
	*set = NewP1IntSet()
}

// Remove a single item from the set.
func (set P1IntSet) Remove(i *int) {
	delete(set, *i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set P1IntSet) Send() <-chan *int {
	ch := make(chan *int)
	go func() {
		for v := range set {
			ch <- &v
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
func (set P1IntSet) Forall(p func(*int) bool) bool {
	for v := range set {
		if !p(&v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set P1IntSet) Exists(p func(*int) bool) bool {
	for v := range set {
		if p(&v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set P1IntSet) Foreach(f func(*int)) {
	for v := range set {
		f(&v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set P1IntSet) Find(p func(*int) bool) (*int, bool) {

	for v := range set {
		if p(&v) {
			return &v, true
		}
	}

	return nil, false
}

// Filter returns a new P1IntSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set P1IntSet) Filter(p func(*int) bool) P1IntSet {
	result := NewP1IntSet()
	for v := range set {
		if p(&v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new P1IntSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set P1IntSet) Partition(p func(*int) bool) (P1IntSet, P1IntSet) {
	matching := NewP1IntSet()
	others := NewP1IntSet()
	for v := range set {
		if p(&v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new P1IntSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set P1IntSet) Map(f func(*int) *int) P1IntSet {
	result := NewP1IntSet()

	for v := range set {
	    k := f(&v)
		result[*k] = struct{}{}
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set P1IntSet) MapToString(f func(*int) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
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
func (set P1IntSet) MapToInt64(f func(*int) int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set))
	for v := range set {
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
func (set P1IntSet) FlatMap(f func(*int) []*int) P1IntSet {
	result := NewP1IntSet()

	for v := range set {
		for _, x := range f(&v) {
			result[*x] = struct{}{}
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
func (set P1IntSet) FlatMapToString(f func(*int) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
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
func (set P1IntSet) FlatMapToInt64(f func(*int) []int64) []int64 {
	if set == nil {
		return nil
	}

	result := make([]int64, 0, len(set))
	for v := range set {
		result = append(result, f(&v)...)
	}

	return result
}

// CountBy gives the number elements of P1IntSet that return true for the predicate p.
func (set P1IntSet) CountBy(p func(*int) bool) (result int) {
	for v := range set {
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
func (set P1IntSet) Min() int {
	v := set.MinBy(func(a *int, b *int) bool {
		return *a < *b
	})
	return *v
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set P1IntSet) Max() int {
	v := set.MaxBy(func(a *int, b *int) bool {
		return *a < *b
	})
	return *v
}

// MinBy returns an element of P1IntSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set P1IntSet) MinBy(less func(*int, *int) bool) *int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m int
	first := true
	for v := range set {
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
func (set P1IntSet) MaxBy(less func(*int, *int) bool) *int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m int
	first := true
	for v := range set {
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
func (set P1IntSet) Sum() int {
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
func (set P1IntSet) Equals(other P1IntSet) bool {
	if set.Size() != other.Size() {
		return false
	}

	for v := range set {
		if !other.Contains(&v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set P1IntSet) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set P1IntSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set P1IntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set P1IntSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set P1IntSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (set P1IntSet) UnmarshalJSON(b []byte) error {
	values := make([]*int, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set P1IntSet) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set P1IntSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
