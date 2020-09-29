// An encapsulated map[*Name]struct{} used as a set.
// Note that the api uses *Name but the set uses <no value> keys.
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=*Name
// options: Comparable:always Numeric:<no value> Ordered:false Stringer:true ToList:true
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// P1NameSet is the primary type that represents a set.
type P1NameSet struct {
	s *sync.RWMutex
	m map[Name]struct{}
}

// NewP1NameSet creates and returns a reference to an empty set.
func NewP1NameSet(values ...*Name) *P1NameSet {
	set := &P1NameSet{
		s: &sync.RWMutex{},
		m: make(map[Name]struct{}),
	}
	for _, i := range values {
		set.m[*i] = struct{}{}
	}
	return set
}

// ConvertP1NameSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertP1NameSet(values ...interface{}) (*P1NameSet, bool) {
	set := NewP1NameSet()

	for _, i := range values {
		switch j := i.(type) {
		case Name:
			k := Name(j)
			set.m[k] = struct{}{}
		case *Name:
			k := Name(*j)
			set.m[k] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildP1NameSetFromChan constructs a new P1NameSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildP1NameSetFromChan(source <-chan *Name) *P1NameSet {
	set := NewP1NameSet()
	for v := range source {
		set.m[*v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *P1NameSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *P1NameSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *P1NameSet) ToList() *P1NameList {
	if set == nil {
		return nil
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return &P1NameList{
		s: &sync.RWMutex{},
		m: set.slice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *P1NameSet) ToSet() *P1NameSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *P1NameSet) slice() []*Name {
	if set == nil {
		return nil
	}

	s := make([]*Name, 0, len(set.m))
	for v := range set.m {
		s = append(s, &v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *P1NameSet) ToSlice() []*Name {
	set.s.RLock()
	defer set.s.RUnlock()

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *P1NameSet) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *P1NameSet) Clone() *P1NameSet {
	if set == nil {
		return nil
	}

	clonedSet := NewP1NameSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *P1NameSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *P1NameSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *P1NameSet) Size() int {
	if set == nil {
		return 0
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *P1NameSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *P1NameSet) Add(more ...*Name) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(*v)
	}
}

func (set *P1NameSet) doAdd(i Name) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *P1NameSet) Contains(i *Name) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[*i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *P1NameSet) ContainsAll(i ...*Name) bool {
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
func (set *P1NameSet) IsSubset(other *P1NameSet) bool {
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
		if !other.Contains(&v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *P1NameSet) IsSuperset(other *P1NameSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *P1NameSet) Union(other *P1NameSet) *P1NameSet {
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
func (set *P1NameSet) Intersect(other *P1NameSet) *P1NameSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewP1NameSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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
func (set *P1NameSet) Difference(other *P1NameSet) *P1NameSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewP1NameSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(&v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *P1NameSet) SymmetricDifference(other *P1NameSet) *P1NameSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *P1NameSet) Clear() {
	if set != nil {
		set.s.Lock()
		defer set.s.Unlock()

		set.m = make(map[Name]struct{})
	}
}

// Remove a single item from the set.
func (set *P1NameSet) Remove(i *Name) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, *i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *P1NameSet) Send() <-chan *Name {
	ch := make(chan *Name)
	go func() {
		if set != nil {
			set.s.RLock()
			defer set.s.RUnlock()

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
func (set *P1NameSet) Forall(p func(*Name) bool) bool {
	if set == nil {
		return true
	}

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *P1NameSet) Exists(p func(*Name) bool) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(&v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *P1NameSet) Foreach(f func(*Name)) {
	if set == nil {
		return
	}

	set.s.Lock()
	defer set.s.Unlock()

	for v := range set.m {
		f(&v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Name that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *P1NameSet) Find(p func(*Name) bool) (*Name, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(&v) {
			return &v, true
		}
	}

	return nil, false
}

// Filter returns a new P1NameSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *P1NameSet) Filter(p func(*Name) bool) *P1NameSet {
	if set == nil {
		return nil
	}

	result := NewP1NameSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(&v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new P1NameSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *P1NameSet) Partition(p func(*Name) bool) (*P1NameSet, *P1NameSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewP1NameSet()
	others := NewP1NameSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(&v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new P1NameSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1NameSet) Map(f func(*Name) *Name) *P1NameSet {
	if set == nil {
		return nil
	}

	result := NewP1NameSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		k := f(&v)
		result.m[*k] = struct{}{}
	}

	return result
}

// FlatMap returns a new P1NameSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *P1NameSet) FlatMap(f func(*Name) []*Name) *P1NameSet {
	if set == nil {
		return nil
	}

	result := NewP1NameSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		for _, x := range f(&v) {
			result.m[*x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of P1NameSet that return true for the predicate p.
func (set *P1NameSet) CountBy(p func(*Name) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(&v) {
			result++
		}
	}
	return
}

// MinBy returns an element of P1NameSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *P1NameSet) MinBy(less func(*Name, *Name) bool) *Name {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m Name
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

// MaxBy returns an element of P1NameSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *P1NameSet) MaxBy(less func(*Name, *Name) bool) *Name {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m Name
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

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *P1NameSet) Equals(other *P1NameSet) bool {
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
		if !other.Contains(&v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// StringSet gets a list of strings that depicts all the elements.
func (set *P1NameSet) StringList() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *P1NameSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *P1NameSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *P1NameSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *P1NameSet) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *P1NameSet) UnmarshalJSON(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	values := make([]*Name, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewP1NameSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *P1NameSet) MarshalJSON() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *P1NameSet) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
