// Generated from {{.TemplateFile}} with Type={{.Type}}

package {{.Package}}

{{if .Stringer}}
import (
    "bytes"
    "fmt"
    "sync"
)
{{else}}
// Stringer is not supported.

import (
    "sync"
)

{{end}}
// {{.UType}}Set is the primary type that represents a set
type {{.UType}}Set struct {
	s *sync.RWMutex
    m map[{{.Type}}]struct{}
}

// New{{.UType}}Set creates and returns a reference to an empty set.
func New{{.UType}}Set(a ...{{.Type}}) {{.UType}}Set {
	set := {{.UType}}Set{
	    s: &sync.RWMutex{},
	    m: make(map[string]struct{}),
	}
	for _, i := range a {
    	set.m[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set {{.UType}}Set) ToSlice() []{{.Type}} {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []{{.Type}}
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Add adds an item to the current set if it doesn't already exist in the set.
func (set {{.UType}}Set) Add(i {{.Type}}) bool {
	set.s.Lock()
	defer set.s.Unlock()

	_, found := set.m[i]
	set.m[i] = struct{}{}
	return !found //False if it existed already
}

// Contains determines if a given item is already in the set.
func (set {{.UType}}Set) Contains(i {{.Type}}) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set {{.UType}}Set) ContainsAll(i ...{{.Type}}) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// IsSubset determines if every item in the other set is in this set.
func (set {{.UType}}Set) IsSubset(other {{.UType}}Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for elem := range set.m {
		if !other.Contains(elem) {
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

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for elem := range set.m {
    	unionedSet.m[elem] = struct{}{}
	}
	for elem := range other.m {
    	unionedSet.m[elem] = struct{}{}
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.UType}}Set) Intersect(other {{.UType}}Set) {{.UType}}Set {
	intersection := New{{.UType}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for elem := range set.m {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other.m {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.UType}}Set) Difference(other {{.UType}}Set) {{.UType}}Set {
	differencedSet := New{{.UType}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for elem := range set.m {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
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
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[string]struct{})
}

// Remove allows the removal of a single item from the set.
func (set {{.UType}}Set) Remove(i {{.Type}}) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set {{.UType}}Set) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set {{.UType}}Set) Cardinality() int {
	return set.Size()
}

// IsEmpty returns true if the set is empty.
func (set {{.UType}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set {{.UType}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// Iter returns a channel of type {{.Type}} that you can range over.
func (set {{.UType}}Set) Iter() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
    	set.s.RLock()
	    defer set.s.RUnlock()

		for elem := range set.m {
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
func (set {{.UType}}Set) Forall(fn func({{.Type}}) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

    for elem := range set.m {
        if !fn(elem) {
            return false
        }
    }
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set {{.UType}}Set) Exists(fn func({{.Type}}) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

    for elem := range set.m {
        if fn(elem) {
            return true
        }
    }
	return false
}

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set {{.UType}}Set) Equals(other {{.UType}}Set) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	if set.Size() != other.Size() {
		return false
	}
	for elem := range set.m {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Clone returns a clone of the set.
// Does NOT clone the underlying elements.
func (set {{.UType}}Set) Clone() {{.UType}}Set {
	clonedSet := New{{.UType}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for elem := range set.m {
		clonedSet.m[elem] = struct{}{}
	}
	return clonedSet
}

{{if .Stringer}}
func {{.LType}}ToString(v {{.Type}}) string {
    return fmt.Sprintf("%v", v)
}

func (set {{.UType}}Set) StringList() []string {
	strings := make([]string, 0)
	set.s.RLock()
	defer set.s.RUnlock()

	for elem := range set.m {
		strings = append(strings, {{.LType}}ToString(elem))
	}
	return strings
}

func (set {{.UType}}Set) String() string {
	return string(set.toBytes("", ", "))
}

// implements encoding.Marshaler interface {
func (set {{.UType}}Set) MarshalJSON() ([]byte, error) {
	return set.toBytes("\"", "\", \""), nil
}

func (set {{.UType}}Set) toBytes(quote, sep string) []byte {
	b := bytes.Buffer{}
	b.WriteString("[")
	comma := quote

	set.s.RLock()
	defer set.s.RUnlock()

	for elem := range set.m {
		b.WriteString(comma)
		b.WriteString({{.LType}}ToString(elem))
		comma = "\", \""
	}
    b.WriteString(quote)
    b.WriteString("]")
	return b.Bytes()
}
{{end}}
