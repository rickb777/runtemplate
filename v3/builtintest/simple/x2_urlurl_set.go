// A simple type derived from map[url.URL]struct{}
//
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=url.URL
// options: Numeric:<no value> Stringer:true Mutable:always
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// X2UrlURLSet is the primary type that represents a set
type X2UrlURLSet map[url.URL]struct{}

// NewX2UrlURLSet creates and returns a reference to an empty set.
func NewX2UrlURLSet(values ...url.URL) X2UrlURLSet {
	set := make(X2UrlURLSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertX2UrlURLSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertX2UrlURLSet(values ...interface{}) (X2UrlURLSet, bool) {
	set := make(X2UrlURLSet)

	for _, i := range values {
		switch j := i.(type) {
		case url.URL:
			k := url.URL(j)
			set[k] = struct{}{}
		case *url.URL:
			k := url.URL(*j)
			set[k] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildX2UrlURLSetFromChan constructs a new X2UrlURLSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2UrlURLSetFromChan(source <-chan url.URL) X2UrlURLSet {
	set := make(X2UrlURLSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set X2UrlURLSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set X2UrlURLSet) IsSet() bool {
	return true
}

// ToSet returns the set; this is an identity operation in this case.
func (set X2UrlURLSet) ToSet() X2UrlURLSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set X2UrlURLSet) ToSlice() []url.URL {
	s := make([]url.URL, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set X2UrlURLSet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set X2UrlURLSet) Clone() X2UrlURLSet {
	clonedSet := NewX2UrlURLSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X2UrlURLSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X2UrlURLSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X2UrlURLSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X2UrlURLSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set X2UrlURLSet) Add(more ...url.URL) X2UrlURLSet {
	for _, v := range more {
		set.doAdd(v)
	}
	return set
}

func (set X2UrlURLSet) doAdd(i url.URL) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set X2UrlURLSet) Contains(i url.URL) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set X2UrlURLSet) ContainsAll(i ...url.URL) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set X2UrlURLSet) IsSubset(other X2UrlURLSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set X2UrlURLSet) IsSuperset(other X2UrlURLSet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set X2UrlURLSet) Append(more ...url.URL) X2UrlURLSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set X2UrlURLSet) Union(other X2UrlURLSet) X2UrlURLSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X2UrlURLSet) Intersect(other X2UrlURLSet) X2UrlURLSet {
	intersection := NewX2UrlURLSet()
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
func (set X2UrlURLSet) Difference(other X2UrlURLSet) X2UrlURLSet {
	differencedSet := NewX2UrlURLSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X2UrlURLSet) SymmetricDifference(other X2UrlURLSet) X2UrlURLSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *X2UrlURLSet) Clear() {
	*set = NewX2UrlURLSet()
}

// Remove a single item from the set.
func (set X2UrlURLSet) Remove(i url.URL) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X2UrlURLSet) Send() <-chan url.URL {
	ch := make(chan url.URL)
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
func (set X2UrlURLSet) Forall(p func(url.URL) bool) bool {
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
func (set X2UrlURLSet) Exists(p func(url.URL) bool) bool {
	for v := range set {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set X2UrlURLSet) Foreach(f func(url.URL)) {
	for v := range set {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first url.URL that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set X2UrlURLSet) Find(p func(url.URL) bool) (url.URL, bool) {

	for v := range set {
		if p(v) {
			return v, true
		}
	}

	var empty url.URL
	return empty, false
}

// Filter returns a new X2UrlURLSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set X2UrlURLSet) Filter(p func(url.URL) bool) X2UrlURLSet {
	result := NewX2UrlURLSet()
	for v := range set {
		if p(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new X2UrlURLSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set X2UrlURLSet) Partition(p func(url.URL) bool) (X2UrlURLSet, X2UrlURLSet) {
	matching := NewX2UrlURLSet()
	others := NewX2UrlURLSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new X2UrlURLSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X2UrlURLSet) Map(f func(url.URL) url.URL) X2UrlURLSet {
	result := NewX2UrlURLSet()

	for v := range set {
		k := f(v)
		result[k] = struct{}{}
	}

	return result
}

// FlatMap returns a new X2UrlURLSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X2UrlURLSet) FlatMap(f func(url.URL) []url.URL) X2UrlURLSet {
	result := NewX2UrlURLSet()

	for v := range set {
		for _, x := range f(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of X2UrlURLSet that return true for the predicate p.
func (set X2UrlURLSet) CountBy(p func(url.URL) bool) (result int) {
	for v := range set {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2UrlURLSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set X2UrlURLSet) MinBy(less func(url.URL, url.URL) bool) url.URL {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m url.URL
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

// MaxBy returns an element of X2UrlURLSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set X2UrlURLSet) MaxBy(less func(url.URL, url.URL) bool) url.URL {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m url.URL
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

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set X2UrlURLSet) Equals(other X2UrlURLSet) bool {
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
func (set X2UrlURLSet) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set X2UrlURLSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set X2UrlURLSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set X2UrlURLSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set X2UrlURLSet) mkString3Bytes(before, between, after string) *strings.Builder {
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
func (set X2UrlURLSet) UnmarshalJSON(b []byte) error {
	values := make([]url.URL, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set X2UrlURLSet) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set X2UrlURLSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
