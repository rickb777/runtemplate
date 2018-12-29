// An encapsulated map[{{.Type}}]struct{} used as a set.
//{{if .Type.IsPtr}} Note that the api uses {{.Type}} but the set uses {{.Key.Name}} keys.{{end}}
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable:always Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} ToList:{{.ToList}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package {{.Package}}

import (
{{- if or .Stringer .GobEncode}}
	"bytes"
{{- end}}
{{- if .GobEncode}}
	"encoding/gob"
{{- end}}
{{- if .Stringer}}
	"encoding/json"
	"fmt"
{{- end}}
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{end}}
)

// {{.Prefix.U}}{{.Type.U}}Set is the primary type that represents a set.
type {{.Prefix.U}}{{.Type.U}}Set struct {
	s *sync.RWMutex
	m map[{{.Type.Name}}]struct{}
}

// New{{.Prefix.U}}{{.Type.U}}Set creates and returns a reference to an empty set.
func New{{.Prefix.U}}{{.Type.U}}Set(values ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}Set {
	set := &{{.Prefix.U}}{{.Type.U}}Set{
		s: &sync.RWMutex{},
		m: make(map[{{.Type.Name}}]struct{}),
	}
	for _, i := range values {
		set.m[{{.Type.Star}}i] = struct{}{}
	}
	return set
}

// Convert{{.Prefix.U}}{{.Type.U}}Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func Convert{{.Prefix.U}}{{.Type.U}}Set(values ...interface{}) (*{{.Prefix.U}}{{.Type.U}}Set, bool) {
	set := New{{.Prefix.U}}{{.Type.U}}Set()

	for _, i := range values {
		switch j := i.(type) {
{{- if .Numeric}}
		case int:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *int:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case int8:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *int8:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case int16:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *int16:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case int32:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *int32:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case int64:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *int64:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case uint:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *uint:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case uint8:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *uint8:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case uint16:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *uint16:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case uint32:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *uint32:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case uint64:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *uint64:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case float32:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *float32:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
		case float64:
			k := {{.Type.Name}}(j)
			set.m[k] = struct{}{}
		case *float64:
			k := {{.Type.Name}}(*j)
			set.m[k] = struct{}{}
{{- else}}
        case {{.Type.Name}}:
			set.m[j] = struct{}{}
        case *{{.Type.Name}}:
			set.m[*j] = struct{}{}
{{- end}}
		}
	}

	return set, len(set.m) == len(values)
}

// Build{{.Prefix.U}}{{.Type.U}}SetFromChan constructs a new {{.Prefix.U}}{{.Type.U}}Set from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}SetFromChan(source <-chan {{.Type}}) *{{.Prefix.U}}{{.Type.U}}Set {
	set := New{{.Prefix.U}}{{.Type.U}}Set()
	for v := range source {
		set.m[{{.Type.Star}}v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *{{.Prefix.U}}{{.Type.U}}Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *{{.Prefix.U}}{{.Type.U}}Set) IsSet() bool {
	return true
}
{{- if .ToList}}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *{{.Prefix.U}}{{.Type.U}}Set) ToList() *{{.Prefix.U}}{{.Type.U}}List {
	if set == nil {
		return nil
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return &{{.Prefix.U}}{{.Type.U}}List{
		s: &sync.RWMutex{},
		m: set.slice(),
	}
}
{{- end}}

// ToSet returns the set; this is an identity operation in this case.
func (set *{{.Prefix.U}}{{.Type.U}}Set) ToSet() *{{.Prefix.U}}{{.Type.U}}Set {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *{{.Prefix.U}}{{.Type.U}}Set) slice() []{{.Type}} {
	if set == nil {
		return nil
	}

	s := make([]{{.Type}}, 0, len(set.m))
	for v := range set.m {
		s = append(s, {{.Type.Amp}}v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *{{.Prefix.U}}{{.Type.U}}Set) ToSlice() []{{.Type}} {
	set.s.RLock()
	defer set.s.RUnlock()

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *{{.Prefix.U}}{{.Type.U}}Set) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Clone() *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return nil
	}

	clonedSet := New{{.Prefix.U}}{{.Type.U}}Set()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *{{.Prefix.U}}{{.Type.U}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *{{.Prefix.U}}{{.Type.U}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Size() int {
	if set == nil {
		return 0
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Add(more ...{{.Type}}) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd({{.Type.Star}}v)
	}
}

func (set *{{.Prefix.U}}{{.Type.U}}Set) doAdd(i {{.Type.Name}}) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Contains(i {{.Type}}) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[{{.Type.Star}}i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *{{.Prefix.U}}{{.Type.U}}Set) ContainsAll(i ...{{.Type}}) bool {
	if set == nil {
		return false
	}

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

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *{{.Prefix.U}}{{.Type.U}}Set) IsSubset(other *{{.Prefix.U}}{{.Type.U}}Set) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains({{.Type.Amp}}v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *{{.Prefix.U}}{{.Type.U}}Set) IsSuperset(other *{{.Prefix.U}}{{.Type.U}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Union(other *{{.Prefix.U}}{{.Type.U}}Set) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Intersect(other *{{.Prefix.U}}{{.Type.U}}Set) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil || other == nil {
		return nil
	}

	intersection := New{{.Prefix.U}}{{.Type.U}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains({{.Type.Amp}}v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains({{.Type.Amp}}v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *{{.Prefix.U}}{{.Type.U}}Set) Difference(other *{{.Prefix.U}}{{.Type.U}}Set) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := New{{.Prefix.U}}{{.Type.U}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains({{.Type.Amp}}v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *{{.Prefix.U}}{{.Type.U}}Set) SymmetricDifference(other *{{.Prefix.U}}{{.Type.U}}Set) *{{.Prefix.U}}{{.Type.U}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Clear() {
	if set != nil {
		set.s.Lock()
		defer set.s.Unlock()

		set.m = make(map[{{.Type.Name}}]struct{})
	}
}

// Remove a single item from the set.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Remove(i {{.Type}}) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, {{.Type.Star}}i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *{{.Prefix.U}}{{.Type.U}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		if set != nil {
			set.s.RLock()
			defer set.s.RUnlock()

			for v := range set.m {
				ch <- {{.Type.Amp}}v
			}
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
func (set *{{.Prefix.U}}{{.Type.U}}Set) Forall(p func({{.Type}}) bool) bool {
	if set == nil {
		return true
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if !p({{.Type.Amp}}v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Exists(p func({{.Type}}) bool) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p({{.Type.Amp}}v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Foreach(f func({{.Type}})) {
	if set == nil {
		return
	}

	set.s.Lock()
	defer set.s.Unlock()

	for v := range set.m {
		f({{.Type.Amp}}v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type.Name}} that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Find(p func({{.Type}}) bool) ({{.Type}}, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p({{.Type.Amp}}v) {
			return {{.Type.Amp}}v, true
		}
	}
{{- if .Type.IsPtr}}

	return nil, false
{{- else}}

	var empty {{.Type.Name}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.Prefix.U}}{{.Type.U}}Set whose elements return true for the predicate p.
//
// The original set is not modified
func (set *{{.Prefix.U}}{{.Type.U}}Set) Filter(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p({{.Type.Amp}}v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new {{.Prefix.U}}{{.Type.U}}Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *{{.Prefix.U}}{{.Type.U}}Set) Partition(p func({{.Type}}) bool) (*{{.Prefix.U}}{{.Type.U}}Set, *{{.Prefix.U}}{{.Type.U}}Set) {
	if set == nil {
		return nil, nil
	}

	matching := New{{.Prefix.U}}{{.Type.U}}Set()
	others := New{{.Prefix.U}}{{.Type.U}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p({{.Type.Amp}}v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new {{.Prefix.U}}{{.Type.U}}Set by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Map(f func({{.Type}}) {{.Type}}) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
	    k := f({{.Type.Amp}}v)
		result.m[{{.Type.Star}}k] = struct{}{}
	}

	return result
}
{{- range .MapTo}}

// MapTo{{.U}} returns a new []{{.}} by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{$.Prefix.U}}{{$.Type.U}}Set) MapTo{{.U}}(f func({{$.Type}}) {{.}}) []{{.}} {
	if set == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(set.m))
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		result = append(result, f({{$.Type.Amp}}v))
	}

	return result
}
{{- end}}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}Set by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{.Prefix.U}}{{.Type.U}}Set) FlatMap(f func({{.Type}}) []{{.Type}}) *{{.Prefix.U}}{{.Type.U}}Set {
	if set == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		for _, x := range f({{.Type.Amp}}v) {
			result.m[{{.Type.Star}}x] = struct{}{}
		}
	}

	return result
}
{{- range .MapTo}}

// FlatMapTo{{.U}} returns a new []{{.}} by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{$.Prefix.U}}{{$.Type.U}}Set) FlatMapTo{{.U}}(f func({{$.Type}}) []{{.}}) []{{.}} {
	if set == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(set.m))
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		result = append(result, f({{$.Type.Amp}}v)...)
	}

	return result
}
{{- end}}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}Set that return true for the predicate p.
func (set *{{.Prefix.U}}{{.Type.U}}Set) CountBy(p func({{.Type}}) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p({{.Type.Amp}}v) {
			result++
		}
	}
	return
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Min() {{.Type.Name}} {
	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type.Name}}
	first := true
	for v := range set.m {
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
func (set *{{.Prefix.U}}{{.Type.U}}Set) Max() (result {{.Type.Name}}) {
	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type.Name}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}
{{- end}}

// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *{{.Prefix.U}}{{.Type.U}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type.Name}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less({{.Type.Amp}}v, {{.Type.Amp}}m) {
			m = v
		}
	}
	return {{.Type.Amp}}m
}

// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *{{.Prefix.U}}{{.Type.U}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type.Name}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less({{.Type.Amp}}m, {{.Type.Amp}}v) {
			m = v
		}
	}
	return {{.Type.Amp}}m
}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is numeric.

// Sum returns the sum of all the elements in the set.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Sum() {{.Type.Name}} {
	set.s.RLock()
	defer set.s.RUnlock()

	sum := {{.Type.Name}}(0)
	for v := range set.m {
		sum = sum + v
	}
	return sum
}
{{- end}}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Equals(other *{{.Prefix.U}}{{.Type.U}}Set) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	if set.Size() != other.Size() {
		return false
	}

	for v := range set.m {
		if !other.Contains({{.Type.Amp}}v) {
			return false
		}
	}

	return true
}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringSet gets a list of strings that depicts all the elements.
func (set *{{.Prefix.U}}{{.Type.U}}Set) StringList() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *{{.Prefix.U}}{{.Type.U}}Set) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *{{.Prefix.U}}{{.Type.U}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *{{.Prefix.U}}{{.Type.U}}Set) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *{{.Prefix.U}}{{.Type.U}}Set) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set *{{.Prefix.U}}{{.Type.U}}Set) UnmarshalJSON(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	values := make([]{{.Type}}, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := New{{.Prefix.U}}{{.Type.U}}Set(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *{{.Prefix.U}}{{.Type.U}}Set) MarshalJSON() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *{{.Prefix.U}}{{.Type.U}}Set) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
{{- end}}
{{- if .GobEncode}}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this set type.
// You must register {{.Type}} with the 'gob' package before this method is used.
func (set *{{.Prefix.U}}{{.Type.U}}Set) GobDecode(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register {{.Type}} with the 'gob' package before this method is used.
func (set {{.Prefix.U}}{{.Type.U}}Set) GobEncode() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
{{- end}}
