// An encapsulated map[{{.Type}}]struct{} used as a set.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable=always Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}} Mutable={{.Mutable}}

package {{.Package}}

{{if .Stringer}}
import (
	"bytes"
	"fmt"
)

{{else}}
// Stringer is not supported.

import (
)

{{end -}}
// {{.UPrefix}}{{.UType}}Set is the primary type that represents a set
type {{.UPrefix}}{{.UType}}Set struct {
	m map[{{.Type}}]struct{}
}

// New{{.UPrefix}}{{.UType}}Set creates and returns a reference to an empty set.
func New{{.UPrefix}}{{.UType}}Set(a ...{{.Type}}) {{.UPrefix}}{{.UType}}Set {
	set := {{.UPrefix}}{{.UType}}Set{
		m: make(map[{{.Type}}]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// Build{{.UPrefix}}{{.UType}}SetFromChan constructs a new {{.UPrefix}}{{.UType}}Set from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}SetFromChan(source <-chan {{.PType}}) {{.UPrefix}}{{.UType}}Set {
	result := New{{.UPrefix}}{{.UType}}Set()
	for v := range source {
		result.m[v] = struct{}{}
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (set {{.UPrefix}}{{.UType}}Set) ToSlice() []{{.Type}} {

	var s []{{.Type}}
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set {{.UPrefix}}{{.UType}}Set) Clone() {{.UPrefix}}{{.UType}}Set {
	clonedSet := New{{.UPrefix}}{{.UType}}Set()


	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set {{.UPrefix}}{{.UType}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set {{.UPrefix}}{{.UType}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set {{.UPrefix}}{{.UType}}Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set {{.UPrefix}}{{.UType}}Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set {{.UPrefix}}{{.UType}}Set) Size() int {

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set {{.UPrefix}}{{.UType}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

{{if .Mutable}}
// Add adds items to the current set.
func (set {{.UPrefix}}{{.UType}}Set) Add(more ...{{.Type}}) {

	for _, v := range more {
		set.doAdd(v)
	}
}

{{else}}
// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set {{.UPrefix}}{{.UType}}Set) Add(more ...{{.Type}}) {{.UPrefix}}{{.UType}}Set {
	newSet := set.Clone()
	for _, v := range more {
		newSet.doAdd(v)
	}
	return newSet
}

{{end -}}
func (set {{.UPrefix}}{{.UType}}Set) doAdd(i {{.Type}}) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set {{.UPrefix}}{{.UType}}Set) Contains(i {{.Type}}) bool {

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set {{.UPrefix}}{{.UType}}Set) ContainsAll(i ...{{.Type}}) bool {

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set {{.UPrefix}}{{.UType}}Set) IsSubset(other {{.UPrefix}}{{.UType}}Set) bool {

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set {{.UPrefix}}{{.UType}}Set) IsSuperset(other {{.UPrefix}}{{.UType}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set {{.UPrefix}}{{.UType}}Set) Union(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	unionedSet := set.Clone()


	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.UPrefix}}{{.UType}}Set) Intersect(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	intersection := New{{.UPrefix}}{{.UType}}Set()


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
func (set {{.UPrefix}}{{.UType}}Set) Difference(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	differencedSet := New{{.UPrefix}}{{.UType}}Set()


	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set {{.UPrefix}}{{.UType}}Set) SymmetricDifference(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

{{if .Mutable}}
// Clear clears the entire set to be the empty set.
func (set *{{.UPrefix}}{{.UType}}Set) Clear() {

	set.m = make(map[{{.Type}}]struct{})
}

// Remove removes a single item from the set.
func (set {{.UPrefix}}{{.UType}}Set) Remove(i {{.Type}}) {

	delete(set.m, i)
}

{{end -}}
//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set {{.UPrefix}}{{.UType}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
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
func (set {{.UPrefix}}{{.UType}}Set) Forall(fn func({{.Type}}) bool) bool {

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
func (set {{.UPrefix}}{{.UType}}Set) Exists(fn func({{.Type}}) bool) bool {

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over {{.Type}}Set and executes the passed func against each element.
func (set {{.UPrefix}}{{.UType}}Set) Foreach(fn func({{.Type}})) {

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new {{.UPrefix}}{{.UType}}Set whose elements return true for func.
func (set {{.UPrefix}}{{.UType}}Set) Filter(fn func({{.Type}}) bool) {{.UPrefix}}{{.UType}}Set {
	result := New{{.UPrefix}}{{.UType}}Set()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set {{.UPrefix}}{{.UType}}Set) Partition(p func({{.Type}}) bool) ({{.UPrefix}}{{.UType}}Set, {{.UPrefix}}{{.UType}}Set) {
	matching := New{{.UPrefix}}{{.UType}}Set()
	others := New{{.UPrefix}}{{.UType}}Set()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Set that return true for the passed predicate.
func (set {{.UPrefix}}{{.UType}}Set) CountBy(predicate func({{.Type}}) bool) (result int) {

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

{{if .Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set {{.UPrefix}}{{.UType}}Set) Min() {{.Type}} {

	var m {{.Type}}
	first := true
	for v, _ := range set.m {
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
func (set {{.UPrefix}}{{.UType}}Set) Max() (result {{.Type}}) {

	var m {{.Type}}
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}

{{else -}}
// MinBy returns an element of {{.UPrefix}}{{.UType}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UPrefix}}{{.UType}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}


	var m {{.Type}}
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

// MaxBy returns an element of {{.UPrefix}}{{.UType}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UPrefix}}{{.UType}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}


	var m {{.Type}}
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

{{end -}}
{{if .Numeric}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the set.
func (set {{.UPrefix}}{{.UType}}Set) Sum() {{.Type}} {

	sum := {{.Type}}(0)
	for v, _ := range set.m {
		sum = sum + {{.TypeStar}}v
	}
	return sum
}

{{end -}}
//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set {{.UPrefix}}{{.UType}}Set) Equals(other {{.UPrefix}}{{.UType}}Set) bool {

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

{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set {{.UPrefix}}{{.UType}}Set) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v, _ := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (set {{.UPrefix}}{{.UType}}Set) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// implements json.Marshaler interface {
func (set {{.UPrefix}}{{.UType}}Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set {{.UPrefix}}{{.UType}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set {{.UPrefix}}{{.UType}}Set) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set {{.UPrefix}}{{.UType}}Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for v, _ := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
{{end}}
