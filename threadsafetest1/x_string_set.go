// Generated from set.tpl with Type=string
// options: Numeric=false Ordered=false Stringer=true Mutable=always-true

package threadsafetest1


import (
	"bytes"
	"fmt"
	"sync"
)

// StringSet is the primary type that represents a set
type StringSet struct {
	s *sync.RWMutex
	m map[string]struct{}
}

// NewStringSet creates and returns a reference to an empty set.
func NewStringSet(a ...string) StringSet {
	set := StringSet{
		s: &sync.RWMutex{},
		m: make(map[string]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set StringSet) ToSlice() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []string
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set StringSet) Clone() StringSet {
	clonedSet := NewStringSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set StringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set StringSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set StringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set StringSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set StringSet) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set StringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds an item to the current set if it doesn't already exist in the set.
func (set StringSet) Add(i string) bool {
	set.s.Lock()
	defer set.s.Unlock()

	_, found := set.m[i]
	set.m[i] = struct{}{}
	return !found //False if it existed already
}

func (set StringSet) doAdd(i string) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set StringSet) Contains(i string) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set StringSet) ContainsAll(i ...string) bool {
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

// IsSubset determines if every item in the other set is in this set.
func (set StringSet) IsSubset(other StringSet) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set StringSet) IsSuperset(other StringSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set StringSet) Append(more ...string) StringSet {
	set.s.Lock()
	defer set.s.Unlock()

	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set StringSet) Union(other StringSet) StringSet {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.m[v] = struct{}{}
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set StringSet) Intersect(other StringSet) StringSet {
	intersection := NewStringSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(v) {
				intersection.Add(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set StringSet) Difference(other StringSet) StringSet {
	differencedSet := NewStringSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.Add(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set StringSet) SymmetricDifference(other StringSet) StringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *StringSet) Clear() {
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[string]struct{})
}

// Remove allows the removal of a single item from the set.
func (set StringSet) Remove(i string) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set StringSet) Send() <-chan string {
	ch := make(chan string)
	go func() {
		set.s.RLock()
		defer set.s.RUnlock()

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
func (set StringSet) Forall(fn func(string) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

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
func (set StringSet) Exists(fn func(string) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over stringSet and executes the passed func against each element.
func (set StringSet) Foreach(fn func(string)) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new StringSet whose elements return true for func.
func (set StringSet) Filter(fn func(string) bool) StringSet {
	result := NewStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if fn(v) {
			result.m[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set StringSet) Partition(p func(string) bool) (StringSet, StringSet) {
	matching := NewStringSet()
	others := NewStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			matching.m[v] = struct{}{}
		} else {
			others.m[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of StringSet that return true for the passed predicate.
func (set StringSet) CountBy(predicate func(string) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of StringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set StringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

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

// MaxBy returns an element of StringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set StringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set StringSet) Equals(other StringSet) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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

func (set StringSet) StringList() []string {
	strings := make([]string, 0)
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set StringSet) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set StringSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set StringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set StringSet) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set StringSet) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

