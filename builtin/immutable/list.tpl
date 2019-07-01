// An encapsulated immutable []{{.Type.Name}}.
// Thread-safe.
//{{if .Type.IsPtr}}
// Warning: THIS COLLECTION IS NOT DESIGNED TO BE USED WITH POINTER TYPES.
//{{end}}
//
// Generated from {{.TemplateFile}} with Type={{.Type.Name}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} GobEncode:{{.GobEncode}} Mutable:disabled
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

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
	"math/rand"
	"sort"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.Prefix.U}}{{.Type.U}}List contains a slice of type {{.Type.Name}}. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.Prefix.U}}{{.Type.U}}List struct {
	m []{{.Type.Name}}
}

//-------------------------------------------------------------------------------------------------

func new{{.Prefix.U}}{{.Type.U}}List(length, capacity int) *{{.Prefix.U}}{{.Type.U}}List {
	return &{{.Prefix.U}}{{.Type.U}}List{
		m: make([]{{.Type.Name}}, length, capacity),
	}
}

// New{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
func New{{.Prefix.U}}{{.Type.U}}List(values ...{{.Type.Name}}) *{{.Prefix.U}}{{.Type.U}}List {
	list := new{{.Prefix.U}}{{.Type.U}}List(len(values), len(values))
	copy(list.m, values)
	return list
}

// Convert{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func Convert{{.Prefix.U}}{{.Type.U}}List(values ...interface{}) (*{{.Prefix.U}}{{.Type.U}}List, bool) {
	list := new{{.Prefix.U}}{{.Type.U}}List(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
{{- if .Numeric}}
		case int:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *int:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case int8:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *int8:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case int16:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *int16:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case int32:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *int32:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case int64:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *int64:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case uint:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *uint:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case uint8:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *uint8:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case uint16:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *uint16:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case uint32:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *uint32:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case uint64:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *uint64:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case float32:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *float32:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case float64:
			k := {{.Type.Name}}(j)
			list.m = append(list.m, {{.Type.Amp}}k)
		case *float64:
			k := {{.Type.Name}}(*j)
			list.m = append(list.m, {{.Type.Amp}}k)
{{- else}}
		case {{.Type.Name}}:
			list.m = append(list.m, j)
		case *{{.Type.Name}}:
			list.m = append(list.m, *j)
{{- end}}
		}
	}

	return list, len(list.m) == len(values)
}

// Build{{.Prefix.U}}{{.Type.U}}ListFromChan constructs a new {{.Prefix.U}}{{.Type.U}}List from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}ListFromChan(source <-chan {{.Type.Name}}) *{{.Prefix.U}}{{.Type.U}}List {
	list := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *{{.Prefix.U}}{{.Type.U}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *{{.Prefix.U}}{{.Type.U}}List) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *{{.Prefix.U}}{{.Type.U}}List) slice() []{{.Type.Name}} {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *{{.Prefix.U}}{{.Type.U}}List) ToList() *{{.Prefix.U}}{{.Type.U}}List {
	return list
}
{{- if .ToSet}}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *{{.Prefix.U}}{{.Type.U}}List) ToSet() *{{.Prefix.U}}{{.Type.U}}Set {
	if list == nil {
		return nil
	}

	return New{{.Prefix.U}}{{.Type.U}}Set(list.m...)
}
{{- end}}

// ToSlice returns the elements of the current list as a slice.
func (list *{{.Prefix.U}}{{.Type.U}}List) ToSlice() []{{.Type.Name}} {

	s := make([]{{.Type.Name}}, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *{{.Prefix.U}}{{.Type.U}}List) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *{{.Prefix.U}}{{.Type.U}}List) Clone() *{{.Prefix.U}}{{.Type.U}}List {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Get(i int) {{.Type.Name}} {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Head() {{.Type.Name}} {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}}.
func (list *{{.Prefix.U}}{{.Type.U}}List) HeadOption() {{.Type.Name}} {
	if list == nil || len(list.m) == 0 {
		var v {{.Type.Name}}
		return v
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Last() {{.Type.Name}} {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}}.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastOption() {{.Type.Name}} {
	if list == nil || len(list.m) == 0 {
		var v {{.Type.Name}}
		return v
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Tail() *{{.Prefix.U}}{{.Type.U}}List {
	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Init() *{{.Prefix.U}}{{.Type.U}}List {
	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether {{.Prefix.U}}{{.Type.U}}List is empty.
func (list *{{.Prefix.U}}{{.Type.U}}List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether {{.Prefix.U}}{{.Type.U}}List is empty.
func (list *{{.Prefix.U}}{{.Type.U}}List) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *{{.Prefix.U}}{{.Type.U}}List) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *{{.Prefix.U}}{{.Type.U}}List) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------
{{- if .Comparable}}

// Contains determines whether a given item is already in the list, returning true if so.
func (list *{{.Prefix.U}}{{.Type.U}}List) Contains(v {{.Type.Name}}) bool {
	return list.Exists(func(x {{.Type.Name}}) bool {
		return {{.Type.Star}}x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *{{.Prefix.U}}{{.Type.U}}List) ContainsAll(i ...{{.Type.Name}}) bool {
	if list == nil {
		return len(i) == 0
	}

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}
{{- end}}

// Exists verifies that one or more elements of {{.Prefix.U}}{{.Type.U}}List return true for the predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) Exists(p func({{.Type.Name}}) bool) bool {
	if list == nil {
		return false
	}

	for _, v := range list.m {
		if p({{.Type.Star}}v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.Prefix.U}}{{.Type.U}}List return true for the predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) Forall(p func({{.Type.Name}}) bool) bool {
	if list == nil {
		return true
	}

	for _, v := range list.m {
		if !p({{.Type.Star}}v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.Prefix.U}}{{.Type.U}}List and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *{{.Prefix.U}}{{.Type.U}}List) Foreach(f func({{.Type.Name}})) {
	if list == nil {
		return
	}

	for _, v := range list.m {
		f({{.Type.Star}}v)
	}
}

// Send returns a channel that will send all the elements in order. A goroutine is created to
// send the elements; this only terminates when all the elements have been consumed. The
// channel will be closed when all the elements have been sent.
func (list *{{.Prefix.U}}{{.Type.U}}List) Send() <-chan {{.Type.Name}} {
	ch := make(chan {{.Type.Name}})
	go func() {
		if list != nil {
			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of {{.Prefix.U}}{{.Type.U}}List with all elements in the reverse order.
func (list *{{.Prefix.U}}{{.Type.U}}List) Reverse() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := new{{.Prefix.U}}{{.Type.U}}List(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of {{.Prefix.U}}{{.Type.U}}List, using a version of the Fisher-Yates shuffle.
func (list *{{.Prefix.U}}{{.Type.U}}List) Shuffle() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}List(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *{{.Prefix.U}}{{.Type.U}}List) Append(more ...{{.Type.Name}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return New{{.Prefix.U}}{{.Type.U}}List(more...)
	}

	newList := New{{.Prefix.U}}{{.Type.U}}List(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doAppend(more ...{{.Type.Name}}) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) Take(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of {{.Prefix.U}}{{.Type.U}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) Drop(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) TakeLast(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of {{.Prefix.U}}{{.Type.U}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) DropLast(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new {{.Prefix.U}}{{.Type.U}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *{{.Prefix.U}}{{.Type.U}}List) TakeWhile(p func({{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new {{.Prefix.U}}{{.Type.U}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *{{.Prefix.U}}{{.Type.U}}List) DropWhile(p func({{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, 0)
	adding := false

	for _, v := range list.m {
		if adding || !p(v) {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type.Name}} that returns true for predicate p.
// False is returned if none match.
func (list *{{.Prefix.U}}{{.Type.U}}List) Find(p func({{.Type.Name}}) bool) ({{.Type.Name}}, bool) {
	if list == nil {
		return {{.Type.Zero}}, false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}
{{if eq .Type.Star "*"}}

	return nil, false
{{else}}

	var empty {{.Type.Name}}
	return empty, false
{{end -}}
}

// Filter returns a new {{.Prefix.U}}{{.Type.U}}List whose elements return true for predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) Filter(p func({{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new {{.Type.Name}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *{{.Prefix.U}}{{.Type.U}}List) Partition(p func({{.Type.Name}}) bool) (*{{.Prefix.U}}{{.Type.U}}List, *{{.Prefix.U}}{{.Type.U}}List) {
	if list == nil {
		return nil, nil
	}

	matching := new{{.Prefix.U}}{{.Type.U}}List(0, len(list.m)/2)
	others := new{{.Prefix.U}}{{.Type.U}}List(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new {{.Prefix.U}}{{.Type.U}}List by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *{{.Prefix.U}}{{.Type.U}}List) Map(f func({{.Type.Name}}) {{.Type.Name}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}List by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *{{.Prefix.U}}{{.Type.U}}List) FlatMap(f func({{.Type.Name}}) []{{.Type.Name}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}List that return true for the predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) CountBy(p func({{.Type.Name}}) bool) (result int) {
	if list == nil {
		return 0
	}

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) MinBy(less func({{.Type.Name}}, {{.Type.Name}}) bool) {{.Type.Name}} {
	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[i], list.m[m]) {
			m = i
		}
	}
	return list.m[m]
}

// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) MaxBy(less func({{.Type.Name}}, {{.Type.Name}}) bool) {{.Type.Name}} {
	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[m], list.m[i]) {
			m = i
		}
	}

	return list.m[m]
}

// DistinctBy returns a new {{.Prefix.U}}{{.Type.U}}List whose elements are unique, where equality is defined by the equal function.
func (list *{{.Prefix.U}}{{.Type.U}}List) DistinctBy(equal func({{.Type.Name}}, {{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := new{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))
Outer:
	for _, v := range list.m {
		for _, r := range result.m {
			if equal(v, r) {
				continue Outer
			}
		}
		result.m = append(result.m, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) IndexWhere(p func({{.Type.Name}}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) IndexWhere2(p func({{.Type.Name}}) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastIndexWhere(p func({{.Type.Name}}) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastIndexWhere2(p func({{.Type.Name}}) bool, before int) int {

	if before < 0 {
		before = len(list.m)
	}
	for i := len(list.m) - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}
{{- if .Comparable}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *{{.Prefix.U}}{{.Type.U}}List) Equals(other *{{.Prefix.U}}{{.Type.U}}List) bool {
	if list == nil {
		return other == nil || len(other.m) == 0
	}

	if other == nil {
		return len(list.m) == 0
	}

	if len(list.m) != len(other.m) {
		return false
	}

	for i, v := range list.m {
		if v != other.m[i] {
			return false
		}
	}

	return true
}
{{- end}}

//-------------------------------------------------------------------------------------------------

type sortable{{.Prefix.U}}{{.Type.U}}List struct {
	less func(i, j {{.Type.Name}}) bool
	m []{{.Type.Name}}
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Less(i, j int) bool {
	return sl.less({{.Type.Star}}sl.m[i], {{.Type.Star}}sl.m[j])
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Len() int {
	return len(sl.m)
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *{{.Prefix.U}}{{.Type.U}}List) SortBy(less func(i, j {{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}List(list.m...)
	sort.Sort(sortable{{.Prefix.U}}{{.Type.U}}List{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) StableSortBy(less func(i, j {{.Type.Name}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Type.U}}List(list.m...)
	sort.Stable(sortable{{.Prefix.U}}{{.Type.U}}List{less, result.m})
	return result
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Sorted returns a new list in which the elements are sorted by their natural ordering.
func (list *{{.Prefix.U}}{{.Type.U}}List) Sorted() *{{.Prefix.U}}{{.Type.U}}List {
	return list.SortBy(func(a, b {{.Type.Name}}) bool {
		return a < b
	})
}

// StableSorted returns a new list in which the elements are sorted by their natural ordering.
func (list *{{.Prefix.U}}{{.Type.U}}List) StableSorted() *{{.Prefix.U}}{{.Type.U}}List {
	return list.StableSortBy(func(a, b {{.Type.Name}}) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *{{.Prefix.U}}{{.Type.U}}List) Min() {{.Type.Name}} {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := {{.Type.Star}}v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if {{.Type.Star}}v < m {
			m = {{.Type.Star}}v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *{{.Prefix.U}}{{.Type.U}}List) Max() (result {{.Type.Name}}) {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	v := list.m[0]
	m := {{.Type.Star}}v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if {{.Type.Star}}v > m {
			m = {{.Type.Star}}v
		}
	}
	return m
}
{{- end}}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is numeric.

// Sum returns the sum of all the elements in the list.
func (list *{{.Prefix.U}}{{.Type.U}}List) Sum() {{.Type.Name}} {

	sum := {{.Type.Name}}(0)
	for _, v := range list.m {
		sum = sum + {{.Type.Star}}v
	}
	return sum
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *{{.Prefix.U}}{{.Type.U}}List) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *{{.Prefix.U}}{{.Type.U}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *{{.Prefix.U}}{{.Type.U}}List) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list {{.Prefix.U}}{{.Type.U}}List) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""


	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this list type.
func (list *{{.Prefix.U}}{{.Type.U}}List) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list {{.Prefix.U}}{{.Type.U}}List) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(list.m)
	return buf, err
}
{{- end}}
{{- if .GobEncode}}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (list *{{.Prefix.U}}{{.Type.U}}List) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (list {{.Prefix.U}}{{.Type.U}}List) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list.m)
	return buf.Bytes(), err
}
{{- end}}
