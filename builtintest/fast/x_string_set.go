// An encapsulated map[string]struct{} used as a set.
// Not thread-safe.
//
// Generated from set.tpl with Type=string
// options: Numeric=false Ordered=false Stringer=true Mutable=true

package fast


import (
	"bytes"
	"fmt"
)

// XStringSet is the primary type that represents a set
type XStringSet struct {
    m map[string]struct{}
}

// NewXStringSet creates and returns a reference to an empty set.
func NewXStringSet(a ...string) *XStringSet {
	set := &XStringSet{
		m: make(map[string]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set *XStringSet) ToSlice() []string {
	var s []string
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set *XStringSet) Clone() *XStringSet {
	clonedSet := NewXStringSet()
	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *XStringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *XStringSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set *XStringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set *XStringSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *XStringSet) Size() int {
	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *XStringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------


// Add adds items to the current set, returning the modified set.
func (set *XStringSet) Add(i ...string) *XStringSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set *XStringSet) doAdd(i string) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set *XStringSet) Contains(i string) bool {
	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set *XStringSet) ContainsAll(i ...string) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set *XStringSet) IsSubset(other *XStringSet) bool {
	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set *XStringSet) IsSuperset(other *XStringSet) bool {
	return other.IsSubset(set)
}

// Append returns a new set with all original items and all in `more`.
func (set *XStringSet) Append(more ...string) *XStringSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set *XStringSet) Union(other *XStringSet) *XStringSet {
	unionedSet := set.Clone()
	for v := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *XStringSet) Intersect(other *XStringSet) *XStringSet {
	intersection := NewXStringSet()
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
func (set *XStringSet) Difference(other *XStringSet) *XStringSet {
	differencedSet := NewXStringSet()
	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *XStringSet) SymmetricDifference(other *XStringSet) *XStringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}


// Clear clears the entire set to be the empty set.
func (set *XStringSet) Clear() {
	set.m = make(map[string]struct{})
}

// Remove allows the removal of a single item from the set.
func (set *XStringSet) Remove(i string) {
	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *XStringSet) Send() <-chan string {
	ch := make(chan string)
	go func() {
		for v := range set.m {
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
func (set *XStringSet) Forall(fn func(string) bool) bool {
	for v := range set.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *XStringSet) Exists(fn func(string) bool) bool {
	for v := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over stringSet and executes the passed func against each element.
func (set *XStringSet) Foreach(fn func(string)) {
	for v := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new XStringSet whose elements return true for func.
func (set *XStringSet) Filter(fn func(string) bool) *XStringSet {
	result := NewXStringSet()
	for v := range set.m {
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
func (set *XStringSet) Partition(p func(string) bool) (*XStringSet, *XStringSet) {
	matching := NewXStringSet()
	others := NewXStringSet()
	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of XStringSet that return true for the passed predicate.
func (set *XStringSet) CountBy(predicate func(string) bool) (result int) {
	for v := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XStringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *XStringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m string
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

// MaxBy returns an element of XStringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *XStringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m string
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

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *XStringSet) Equals(other *XStringSet) bool {
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

func (set *XStringSet) StringList() []string {
	strings := make([]string, 0)
	for v := range set.m {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set *XStringSet) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set *XStringSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *XStringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *XStringSet) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set *XStringSet) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

