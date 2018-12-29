// A simple type derived from map[string]struct{}
// Note that the api uses *string but the set uses <no value> keys.
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=string
// options: Numeric:false Stringer:true Mutable:always
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple


import (
	"bytes"
	"encoding/json"
	"fmt")

// P1StringSet is the primary type that represents a set
type P1StringSet map[string]struct{}

// NewP1StringSet creates and returns a reference to an empty set.
func NewP1StringSet(values ...*string) P1StringSet {
	set := make(P1StringSet)
	for _, i := range values {
		set[*i] = struct{}{}
	}
	return set
}

// ConvertP1StringSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertP1StringSet(values ...interface{}) (P1StringSet, bool) {
	set := make(P1StringSet)

	for _, i := range values {
		switch j := i.(type) {
        case string:
			set[j] = struct{}{}
        case *string:
			set[*j] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildP1StringSetFromChan constructs a new P1StringSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildP1StringSetFromChan(source <-chan *string) P1StringSet {
	set := make(P1StringSet)
	for v := range source {
		set[*v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set P1StringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set P1StringSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set P1StringSet) ToList() P1StringList {
	if set == nil {
		return nil
	}

	return P1StringList(set.ToSlice())
}

// ToSet returns the set; this is an identity operation in this case.
func (set P1StringSet) ToSet() P1StringSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set P1StringSet) ToSlice() []*string {
	s := make([]*string, 0, len(set))
	for v := range set {
		s = append(s, &v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set P1StringSet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set P1StringSet) Clone() P1StringSet {
	clonedSet := NewP1StringSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set P1StringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set P1StringSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set P1StringSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set P1StringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set P1StringSet) Add(more ...*string) P1StringSet {
	for _, v := range more {
		set.doAdd(*v)
	}
	return set
}

func (set P1StringSet) doAdd(i string) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set P1StringSet) Contains(i *string) bool {
	_, found := set[*i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set P1StringSet) ContainsAll(i ...*string) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set P1StringSet) IsSubset(other P1StringSet) bool {
	for v := range set {
		if !other.Contains(&v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set P1StringSet) IsSuperset(other P1StringSet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set P1StringSet) Append(more ...string) P1StringSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set P1StringSet) Union(other P1StringSet) P1StringSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set P1StringSet) Intersect(other P1StringSet) P1StringSet {
	intersection := NewP1StringSet()
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
func (set P1StringSet) Difference(other P1StringSet) P1StringSet {
	differencedSet := NewP1StringSet()
	for v := range set {
		if !other.Contains(&v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set P1StringSet) SymmetricDifference(other P1StringSet) P1StringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *P1StringSet) Clear() {
	*set = NewP1StringSet()
}

// Remove a single item from the set.
func (set P1StringSet) Remove(i *string) {
	delete(set, *i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set P1StringSet) Send() <-chan *string {
	ch := make(chan *string)
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
func (set P1StringSet) Forall(p func(*string) bool) bool {
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
func (set P1StringSet) Exists(p func(*string) bool) bool {
	for v := range set {
		if p(&v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set P1StringSet) Foreach(f func(*string)) {
	for v := range set {
		f(&v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first string that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set P1StringSet) Find(p func(*string) bool) (*string, bool) {

	for v := range set {
		if p(&v) {
			return &v, true
		}
	}

	return nil, false
}

// Filter returns a new P1StringSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set P1StringSet) Filter(p func(*string) bool) P1StringSet {
	result := NewP1StringSet()
	for v := range set {
		if p(&v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new P1StringSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set P1StringSet) Partition(p func(*string) bool) (P1StringSet, P1StringSet) {
	matching := NewP1StringSet()
	others := NewP1StringSet()
	for v := range set {
		if p(&v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new P1StringSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set P1StringSet) Map(f func(*string) *string) P1StringSet {
	result := NewP1StringSet()

	for v := range set {
	    k := f(&v)
		result[*k] = struct{}{}
	}

	return result
}

// FlatMap returns a new P1StringSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set P1StringSet) FlatMap(f func(*string) []*string) P1StringSet {
	result := NewP1StringSet()

	for v := range set {
		for _, x := range f(&v) {
			result[*x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of P1StringSet that return true for the predicate p.
func (set P1StringSet) CountBy(p func(*string) bool) (result int) {
	for v := range set {
		if p(&v) {
			result++
		}
	}
	return
}

// MinBy returns an element of P1StringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set P1StringSet) MinBy(less func(*string, *string) bool) *string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m string
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

// MaxBy returns an element of P1StringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set P1StringSet) MaxBy(less func(*string, *string) bool) *string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m string
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

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set P1StringSet) Equals(other P1StringSet) bool {
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
func (set P1StringSet) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set P1StringSet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set P1StringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set P1StringSet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set P1StringSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (set P1StringSet) UnmarshalJSON(b []byte) error {
	values := make([]*string, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set P1StringSet) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set P1StringSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
