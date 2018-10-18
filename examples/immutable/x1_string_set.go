// An encapsulated map[string]struct{} used as a set.
// Thread-safe.
//
// Generated from immutable/set.tpl with Type=string
// options: Comparable:always Numeric:false Ordered:false Stringer:true Mutable:disabled

package immutable


import (

	"bytes"
	"fmt"
)

// X1StringSet is the primary type that represents a set
type X1StringSet struct {
	m map[string]struct{}
}

// NewX1StringSet creates and returns a reference to an empty set.
func NewX1StringSet(values ...string) X1StringSet {
	set := X1StringSet{
		m: make(map[string]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertX1StringSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertX1StringSet(values ...interface{}) (X1StringSet, bool) {
	set := NewX1StringSet()

	for _, i := range values {
		v, ok := i.(string)
		if ok {
			set.m[v] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildX1StringSetFromChan constructs a new X1StringSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1StringSetFromChan(source <-chan string) X1StringSet {
	set := NewX1StringSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set X1StringSet) ToSlice() []string {

	var s []string
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set X1StringSet) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set X1StringSet) Clone() X1StringSet {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X1StringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X1StringSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set X1StringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set X1StringSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X1StringSet) Size() int {
	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X1StringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set X1StringSet) Add(more ...string) X1StringSet {
	newSet := NewX1StringSet()

	for v, _ := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set X1StringSet) doAdd(i string) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set X1StringSet) Contains(i string) bool {

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set X1StringSet) ContainsAll(i ...string) bool {

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set X1StringSet) IsSubset(other X1StringSet) bool {

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set X1StringSet) IsSuperset(other X1StringSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set X1StringSet) Union(other X1StringSet) X1StringSet {
	unionedSet := NewX1StringSet()

	for v, _ := range set.m {
		unionedSet.doAdd(v)
	}

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X1StringSet) Intersect(other X1StringSet) X1StringSet {
	intersection := NewX1StringSet()

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
func (set X1StringSet) Difference(other X1StringSet) X1StringSet {
	differencedSet := NewX1StringSet()

	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X1StringSet) SymmetricDifference(other X1StringSet) X1StringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set X1StringSet) Remove(i string) X1StringSet {
	clonedSet := NewX1StringSet()

	for v, _ := range set.m {
		if i != v {
			clonedSet.doAdd(v)
		}
	}

	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X1StringSet) Send() <-chan string {
	ch := make(chan string)
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
func (set X1StringSet) Forall(fn func(string) bool) bool {

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
func (set X1StringSet) Exists(fn func(string) bool) bool {

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over stringSet and executes the passed func against each element.
func (set X1StringSet) Foreach(fn func(string)) {

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (set X1StringSet) Find(fn func(string) bool) (string, bool) {

	for v, _ := range set.m {
		if fn(v) {
			return v, true
		}
	}


	var empty string
	return empty, false

}

// Filter returns a new X1StringSet whose elements return true for func.
func (set X1StringSet) Filter(fn func(string) bool) X1StringSet {
	result := NewX1StringSet()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set X1StringSet) Partition(p func(string) bool) (X1StringSet, X1StringSet) {
	matching := NewX1StringSet()
	others := NewX1StringSet()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new X1StringSet by transforming every element with a function fn.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X1StringSet) Map(fn func(string) string) X1StringSet {
	result := NewX1StringSet()

	for v := range set.m {
        result.m[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new X1StringSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set X1StringSet) FlatMap(fn func(string) []string) X1StringSet {
	result := NewX1StringSet()

	for v, _ := range set.m {
	    for _, x := range fn(v) {
            result.m[x] = struct{}{}
	    }
	}

	return result
}

// CountBy gives the number elements of X1StringSet that return true for the passed predicate.
func (set X1StringSet) CountBy(predicate func(string) bool) (result int) {

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1StringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set X1StringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}


	var m string
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

// MaxBy returns an element of X1StringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set X1StringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}


	var m string
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
func (set X1StringSet) Equals(other X1StringSet) bool {

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
func (set X1StringSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (set X1StringSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// implements json.Marshaler interface {
func (set X1StringSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set X1StringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set X1StringSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set X1StringSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set X1StringSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

