// A simple type derived from map[url.URL]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=url.URL
// options: Numeric:<no value> Stringer:true Mutable:always

package simple


import (

	"bytes"
	"fmt"
	"net/url"
)

// X2URLSet is the primary type that represents a set
type X2URLSet map[url.URL]struct{}

// NewX2URLSet creates and returns a reference to an empty set.
func NewX2URLSet(values ...url.URL) X2URLSet {
	set := make(X2URLSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertX2URLSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertX2URLSet(values ...interface{}) (X2URLSet, bool) {
	set := make(X2URLSet)

	for _, i := range values {
		v, ok := i.(url.URL)
		if ok {
		    set[v] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildX2URLSetFromChan constructs a new X2URLSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2URLSetFromChan(source <-chan url.URL) X2URLSet {
	set := make(X2URLSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set X2URLSet) ToSlice() []url.URL {
	var s []url.URL
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set X2URLSet) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v, _ := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set X2URLSet) Clone() X2URLSet {
	clonedSet := NewX2URLSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X2URLSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X2URLSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set X2URLSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set X2URLSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X2URLSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X2URLSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set X2URLSet) Add(i ...url.URL) X2URLSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set X2URLSet) doAdd(i url.URL) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set X2URLSet) Contains(i url.URL) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set X2URLSet) ContainsAll(i ...url.URL) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set X2URLSet) IsSubset(other X2URLSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set X2URLSet) IsSuperset(other X2URLSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set X2URLSet) Append(more ...url.URL) X2URLSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set X2URLSet) Union(other X2URLSet) X2URLSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X2URLSet) Intersect(other X2URLSet) X2URLSet {
	intersection := NewX2URLSet()
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
func (set X2URLSet) Difference(other X2URLSet) X2URLSet {
	differencedSet := NewX2URLSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X2URLSet) SymmetricDifference(other X2URLSet) X2URLSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *X2URLSet) Clear() {
	*set = NewX2URLSet()
}

// Remove allows the removal of a single item from the set.
func (set X2URLSet) Remove(i url.URL) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X2URLSet) Send() <-chan url.URL {
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

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set X2URLSet) Forall(fn func(url.URL) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set X2URLSet) Exists(fn func(url.URL) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over url.URLSet and executes the passed func against each element.
func (set X2URLSet) Foreach(fn func(url.URL)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new X2URLSet whose elements return true for func.
// The original set is not modified
func (set X2URLSet) Filter(fn func(url.URL) bool) X2URLSet {
	result := NewX2URLSet()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new url.URLSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
// The original set is not modified
func (set X2URLSet) Partition(p func(url.URL) bool) (X2URLSet, X2URLSet) {
	matching := NewX2URLSet()
	others := NewX2URLSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new X2URLSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X2URLSet) Map(fn func(url.URL) url.URL) X2URLSet {
	result := NewX2URLSet()

	for v := range set {
        result[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new X2URLSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X2URLSet) FlatMap(fn func(url.URL) []url.URL) X2URLSet {
	result := NewX2URLSet()

	for v, _ := range set {
	    for _, x := range fn(v) {
            result[x] = struct{}{}
	    }
	}

	return result
}

// CountBy gives the number elements of X2URLSet that return true for the passed predicate.
func (set X2URLSet) CountBy(predicate func(url.URL) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2URLSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set X2URLSet) MinBy(less func(url.URL, url.URL) bool) url.URL {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
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

// MaxBy returns an element of X2URLSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set X2URLSet) MaxBy(less func(url.URL, url.URL) bool) url.URL {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
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

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set X2URLSet) Equals(other X2URLSet) bool {
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

func (set X2URLSet) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set X2URLSet) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set X2URLSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set X2URLSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set X2URLSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set X2URLSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set X2URLSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

