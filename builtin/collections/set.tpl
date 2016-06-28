// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable={{.Comparable}} Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}}

package {{.Package}}

{{if .Stringer}}
import (
    "bytes"
    "fmt"
)
{{else}}
// Stringer is not supported.

{{end}}
// {{.UType}}Set is the primary type that represents a set
type {{.UType}}Set map[{{.Type}}]struct{}

// New{{.UType}}Set creates and returns a reference to an empty set.
func New{{.UType}}Set(a ...{{.Type}}) {{.UType}}Set {
	set := make({{.UType}}Set)
	for _, i := range a {
		set[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set {{.UType}}Set) ToSlice() []{{.Type}} {
	var s []{{.Type}}
	for v := range set {
		s = append(s, v)
	}
	return s
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set {{.UType}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set {{.UType}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set {{.UType}}Set) IsSequence() bool {
    return false
}

// IsSet returns false for lists.
func (set {{.UType}}Set) IsSet() bool {
    return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set {{.UType}}Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set {{.UType}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds an item to the current set if it doesn't already exist in the set.
func (set {{.UType}}Set) Add(i {{.Type}}) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Contains determines if a given item is already in the set.
func (set {{.UType}}Set) Contains(i {{.Type}}) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set {{.UType}}Set) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		_, found := set[v]
		if !found {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set {{.UType}}Set) IsSubset(other {{.UType}}Set) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set {{.UType}}Set) IsSuperset(other {{.UType}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set {{.UType}}Set) Union(other {{.UType}}Set) {{.UType}}Set {
	unionedSet := New{{.UType}}Set()

	for v := range set {
		unionedSet.Add(v)
	}
	for v := range other {
		unionedSet.Add(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.UType}}Set) Intersect(other {{.UType}}Set) {{.UType}}Set {
	intersection := New{{.UType}}Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.Add(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.UType}}Set) Difference(other {{.UType}}Set) {{.UType}}Set {
	differencedSet := New{{.UType}}Set()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.Add(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set {{.UType}}Set) SymmetricDifference(other {{.UType}}Set) {{.UType}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *{{.UType}}Set) Clear() {
	*set = New{{.UType}}Set()
}

// Remove allows the removal of a single item from the set.
func (set {{.UType}}Set) Remove(i {{.Type}}) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel of type {{.Type}} that you can range over.
func (set {{.UType}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

// Clone returns a clone of the set.
// Does NOT clone the underlying elements.
func (set {{.UType}}Set) Clone() {{.UType}}Set {
	clonedSet := New{{.UType}}Set()
	for v := range set {
		clonedSet.Add(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set {{.UType}}Set) Forall(fn func({{.Type}}) bool) bool {
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
func (set {{.UType}}Set) Exists(fn func({{.Type}}) bool) bool {
    for v := range set {
        if fn(v) {
            return true
        }
    }
	return false
}

// Foreach iterates over {{.Type}}Set and executes the passed func against each element.
func (set {{.UType}}Set) Foreach(fn func({{.Type}})) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new {{.UType}}Set whose elements return true for func.
func (set {{.UType}}Set) Filter(fn func({{.Type}}) bool) {{.UType}}Set {
	result := New{{.UType}}Set()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set {{.UType}}Set) Partition(p func({{.Type}}) bool) ({{.UType}}Set, {{.UType}}Set) {
	matching := New{{.UType}}Set()
	others := New{{.UType}}Set()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of {{.UType}}Set that return true for the passed predicate.
func (set {{.UType}}Set) CountBy(predicate func({{.Type}}) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.UType}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UType}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m {{.Type}}
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

// MaxBy returns an element of {{.UType}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UType}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m {{.Type}}
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
func (set {{.UType}}Set) Equals(other {{.UType}}Set) bool {
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

{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

func {{.LType}}ToString(v {{.Type}}) string {
    return fmt.Sprintf("%v", v)
}

func (set {{.UType}}Set) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, {{.LType}}ToString(v))
	}
	return strings
}

func (set {{.UType}}Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set {{.UType}}Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set {{.UType}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set {{.UType}}Set) MkString3(pfx, mid, sfx string) string {
    return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set {{.UType}}Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString({{.LType}}ToString(v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
{{end}}
