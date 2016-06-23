// Generated from {{.TemplateFile}} with Type={{.Type}}

package {{.Package}}

{{if .Stringer}}
import (
    "bytes"
)
{{else}}
// Stringer is not supported.

{{end}}
// {{.Type}}Set is the primary type that represents a set
type {{.Type}}Set map[{{.Type}}]struct{}

// New{{.Type}}Set creates and returns a reference to an empty set.
func New{{.Type}}Set(a ...{{.Type}}) {{.Type}}Set {
	s := make({{.Type}}Set)
	for _, i := range a {
		s.Add(i)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice
func (set {{.Type}}Set) ToSlice() []{{.Type}} {
	var s []{{.Type}}
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Add adds an item to the current set if it doesn't already exist in the set.
func (set {{.Type}}Set) Add(i {{.Type}}) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Contains determines if a given item is already in the set.
func (set {{.Type}}Set) Contains(i {{.Type}}) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set {{.Type}}Set) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// IsSubset determines if every item in the other set is in this set.
func (set {{.Type}}Set) IsSubset(other {{.Type}}Set) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set {{.Type}}Set) IsSuperset(other {{.Type}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set {{.Type}}Set) Union(other {{.Type}}Set) {{.Type}}Set {
	unionedSet := New{{.Type}}Set()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.Type}}Set) Intersect(other {{.Type}}Set) {{.Type}}Set {
	intersection := New{{.Type}}Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.Type}}Set) Difference(other {{.Type}}Set) {{.Type}}Set {
	differencedSet := New{{.Type}}Set()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set {{.Type}}Set) SymmetricDifference(other {{.Type}}Set) {{.Type}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *{{.Type}}Set) Clear() {
	*set = make({{.Type}}Set)
}

// Remove allows the removal of a single item in the set.
func (set {{.Type}}Set) Remove(i {{.Type}}) {
	delete(set, i)
}

// Size returns how many items are currently in the set.
func (set {{.Type}}Set) Size() int {
	return len(set)
}

// Iter returns a channel of type {{.Type}} that you can range over.
func (set {{.Type}}Set) Iter() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set {{.Type}}Set) Forall(fn func({{.Type}}) bool) bool {
    for elem := range set {
        if !fn(elem) {
            return false
        }
    }
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set {{.Type}}Set) Exists(fn func({{.Type}}) bool) bool {
    for elem := range set {
        if fn(elem) {
            return true
        }
    }
	return false
}

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set {{.Type}}Set) Equals(other {{.Type}}Set) bool {
	if set.Size() != other.Size() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Clone returns a clone of the set.
// Does NOT clone the underlying elements.
func (set {{.Type}}Set) Clone() {{.Type}}Set {
	clonedSet := New{{.Type}}Set()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}

{{if .Stringer}}
func (set {{.Type}}Set) StringList() []string {
	strings := make([]string, 0)
	for elem := range set {
		strings = append(strings, elem.String())
	}
	return strings
}

func (set {{.Type}}Set) String() string {
	return string(set.toBytes("", ", "))
}

// implements encoding.Marshaler interface {
func (set {{.Type}}Set) MarshalJSON() ([]byte, error) {
	return set.toBytes("\"", "\", \""), nil
}

func (set {{.Type}}Set) toBytes(quote, sep string) []byte {
	b := bytes.Buffer{}
	b.WriteString("[")
	comma := quote
	for elem := range set {
		b.WriteString(comma)
		b.WriteString(elem.String())
		comma = "\", \""
	}
    b.WriteString(quote)
    b.WriteString("]")
	return b.Bytes()
}
{{end}}
