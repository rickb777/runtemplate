// A simple type derived from []{{.Type}}
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} StringLike:{{.StringLike}} Stringer:{{.Stringer}}
// GobEncode:{{.GobEncode}} Mutable:always ToList:always ToSet:{{.ToSet}} MapTo:{{.MapTo}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package {{.Package}}

import (
{{- if .GobEncode}}
	"bytes"
{{- end}}
{{- if .GobEncode}}
	"encoding/gob"
{{- end}}
{{- if .Stringer}}
	"fmt"
{{- end}}
	"math/rand"
	"sort"
{{- if .Stringer}}
	"strings"
{{- end}}
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.Prefix.U}}{{.Type.U}}List is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
// To add items to the list, simply use the normal built-in append function.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.Prefix.U}}{{.Type.U}}List []{{.Type}}

//-------------------------------------------------------------------------------------------------

// Make{{.Prefix.U}}{{.Type.U}}List makes an empty list with both length and capacity initialised.
func Make{{.Prefix.U}}{{.Type.U}}List(length, capacity int) {{.Prefix.U}}{{.Type.U}}List {
	return make({{.Prefix.U}}{{.Type.U}}List, length, capacity)
}

// New{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
func New{{.Prefix.U}}{{.Type.U}}List(values ...{{.Type}}) {{.Prefix.U}}{{.Type.U}}List {
	list := Make{{.Prefix.U}}{{.Type.U}}List(len(values), len(values))
	copy(list, values)
	return list
}

// Convert{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func Convert{{.Prefix.U}}{{.Type.U}}List(values ...interface{}) ({{.Prefix.U}}{{.Type.U}}List, bool) {
	list := Make{{.Prefix.U}}{{.Type.U}}List(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
{{- if .Numeric}}
		case int:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *int:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case int8:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *int8:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case int16:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *int16:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case int32:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *int32:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case int64:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *int64:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case uint:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *uint:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case uint8:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *uint8:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case uint16:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *uint16:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case uint32:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *uint32:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case uint64:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *uint64:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case float32:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *float32:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
		case float64:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *float64:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
{{- else if .Type.IsPtr}}
		case {{.Type.Name}}:
			list = append(list, &j)
		case *{{.Type.Name}}:
			list = append(list, j)
{{- else}}
		case {{.Type.Name}}:
			list = append(list, j)
		case *{{.Type.Name}}:
			list = append(list, *j)
{{- end}}
{{- if and .StringLike .ne .Type.Name "string"}}
		case string:
			k := {{.Type.Name}}(j)
			list = append(list, {{.Type.Amp}}k)
		case *string:
			k := {{.Type.Name}}(*j)
			list = append(list, {{.Type.Amp}}k)
{{- end}}
{{- if .StringLike}}
		default:
			if s, ok := i.(fmt.Stringer); ok {
				k := {{.Type.Name}}(s.String())
				list = append(list, {{.Type.Amp}}k)
			}
{{- end}}
		}
	}

	return list, len(list) == len(values)
}

// Build{{.Prefix.U}}{{.Type.U}}ListFromChan constructs a new {{.Prefix.U}}{{.Type.U}}List from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}ListFromChan(source <-chan {{.Type}}) {{.Prefix.U}}{{.Type.U}}List {
	list := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	for v := range source {
		list = append(list, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list {{.Prefix.U}}{{.Type.U}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list {{.Prefix.U}}{{.Type.U}}List) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list {{.Prefix.U}}{{.Type.U}}List) slice() []{{.Type}} {
	return list
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list {{.Prefix.U}}{{.Type.U}}List) ToList() {{.Prefix.U}}{{.Type.U}}List {
	return list
}
{{- if .ToSet}}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list {{.Prefix.U}}{{.Type.U}}List) ToSet() {{.Prefix.U}}{{.Type.U}}Set {
	if list == nil {
		return nil
	}

	return New{{.Prefix.U}}{{.Type.U}}Set(list...)
}
{{- end}}

// ToSlice returns the elements of the list as a slice, which is an identity operation in this case,
// because the simple list is merely a dressed-up slice.
func (list {{.Prefix.U}}{{.Type.U}}List) ToSlice() []{{.Type}} {
	return list
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list {{.Prefix.U}}{{.Type.U}}List) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(list))
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list {{.Prefix.U}}{{.Type.U}}List) Clone() {{.Prefix.U}}{{.Type.U}}List {
	return New{{.Prefix.U}}{{.Type.U}}List(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list {{.Prefix.U}}{{.Type.U}}List) Get(i int) {{.Type}} {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list {{.Prefix.U}}{{.Type.U}}List) Head() {{.Type}} {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}}.
func (list {{.Prefix.U}}{{.Type.U}}List) HeadOption() ({{.Type}}, bool) {
	if list.IsEmpty() {
		return {{.Type.Zero}}, false
	}
	return list[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list {{.Prefix.U}}{{.Type.U}}List) Last() {{.Type}} {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}}.
func (list {{.Prefix.U}}{{.Type.U}}List) LastOption() ({{.Type}}, bool) {
	if list.IsEmpty() {
		return {{.Type.Zero}}, false
	}
	return list[len(list)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list {{.Prefix.U}}{{.Type.U}}List) Tail() {{.Prefix.U}}{{.Type.U}}List {
	return list[1:]
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list {{.Prefix.U}}{{.Type.U}}List) Init() {{.Prefix.U}}{{.Type.U}}List {
	return list[:len(list)-1]
}

// IsEmpty tests whether {{.Prefix.U}}{{.Type.U}}List is empty.
func (list {{.Prefix.U}}{{.Type.U}}List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether {{.Prefix.U}}{{.Type.U}}List is empty.
func (list {{.Prefix.U}}{{.Type.U}}List) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list {{.Prefix.U}}{{.Type.U}}List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list {{.Prefix.U}}{{.Type.U}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.Prefix.U}}{{.Type.U}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------
{{- if .Comparable}}

// Contains determines whether a given item is already in the list, returning true if so.
func (list {{.Prefix.U}}{{.Type.U}}List) Contains(v {{.Type}}) bool {
	return list.Exists(func(x {{.Type}}) bool {
		return {{.Type.Star}}x == {{.Type.Star}}v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list {{.Prefix.U}}{{.Type.U}}List) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}
{{- end}}

// Exists verifies that one or more elements of {{.Prefix.U}}{{.Type.U}}List return true for the predicate p.
func (list {{.Prefix.U}}{{.Type.U}}List) Exists(p func({{.Type}}) bool) bool {
	for _, v := range list {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.Prefix.U}}{{.Type.U}}List return true for the predicate p.
func (list {{.Prefix.U}}{{.Type.U}}List) Forall(p func({{.Type}}) bool) bool {
	for _, v := range list {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.Prefix.U}}{{.Type.U}}List and executes function f against each element.
func (list {{.Prefix.U}}{{.Type.U}}List) Foreach(f func({{.Type}})) {
	for _, v := range list {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list {{.Prefix.U}}{{.Type.U}}List) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of {{.Prefix.U}}{{.Type.U}}List with all elements in the reverse order.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) Reverse() {{.Prefix.U}}{{.Type.U}}List {
	n := len(list)
	result := Make{{.Prefix.U}}{{.Type.U}}List(n, n)
	last := n - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse alters a {{.Prefix.U}}{{.Type.U}}List with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) DoReverse() {{.Prefix.U}}{{.Type.U}}List {
	mid := (len(list) + 1) / 2
	last := len(list) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list[i], list[r] = list[r], list[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of {{.Prefix.U}}{{.Type.U}}List, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) Shuffle() {{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	return list.Clone().DoShuffle()
}

// DoShuffle returns a shuffled {{.Prefix.U}}{{.Type.U}}List, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) DoShuffle() {{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	n := len(list)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list[i], list[r] = list[r], list[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) Take(n int) {{.Prefix.U}}{{.Type.U}}List {
	if n >= len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of {{.Prefix.U}}{{.Type.U}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) Drop(n int) {{.Prefix.U}}{{.Type.U}}List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) TakeLast(n int) {{.Prefix.U}}{{.Type.U}}List {
	l := len(list)
	if n >= l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of {{.Prefix.U}}{{.Type.U}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) DropLast(n int) {{.Prefix.U}}{{.Type.U}}List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	}
	return list[0 : l-n]
}

// TakeWhile returns a new {{.Prefix.U}}{{.Type.U}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) TakeWhile(p func({{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new {{.Prefix.U}}{{.Type.U}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) DropWhile(p func({{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	adding := false

	for _, v := range list {
		if adding || !p(v) {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type.Name}} that returns true for predicate p.
// False is returned if none match.
func (list {{.Prefix.U}}{{.Type.U}}List) Find(p func({{.Type}}) bool) ({{.Type}}, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}
{{- if .Type.IsPtr}}

	return nil, false
{{- else}}

	var empty {{.Type}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.Prefix.U}}{{.Type.U}}List whose elements return true for predicate p.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) Filter(p func({{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new {{.Type.U}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list {{.Prefix.U}}{{.Type.U}}List) Partition(p func({{.Type}}) bool) ({{.Prefix.U}}{{.Type.U}}List, {{.Prefix.U}}{{.Type.U}}List) {
	matching := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))
	others := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new {{.Prefix.U}}{{.Type.U}}List by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{.Prefix.U}}{{.Type.U}}List) Map(f func({{.Type}}) {{.Type}}) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))

	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}
{{- range .MapTo}}

// MapTo{{.U}} returns a new []{{.}} by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{$.Prefix.U}}{{$.Type.U}}List) MapTo{{.U}}(f func({{$.Type}}) {{.}}) []{{.}} {
	if list == nil {
		return nil
	}

	result := make([]{{.}}, len(list))
	for i, v := range list {
		result[i] = f(v)
	}

	return result
}
{{- end}}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}List by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{.Prefix.U}}{{.Type.U}}List) FlatMap(f func({{.Type}}) []{{.Type}}) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))

	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}
{{- range .MapTo}}

// FlatMapTo{{.U}} returns a new []{{.}} by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{$.Prefix.U}}{{$.Type.U}}List) FlatMapTo{{.U}}(f func({{$.Type}}) []{{.}}) []{{.}} {
	if list == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(list))
	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}
{{- end}}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}List that return true for the predicate p.
func (list {{.Prefix.U}}{{.Type.U}}List) CountBy(p func({{.Type}}) bool) (result int) {
	for _, v := range list {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list {{.Prefix.U}}{{.Type.U}}List) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}

	return list[m]
}

// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list {{.Prefix.U}}{{.Type.U}}List) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[m], list[i]) {
			m = i
		}
	}

	return list[m]
}

// DistinctBy returns a new {{.Prefix.U}}{{.Type.U}}List whose elements are unique, where equality is defined by the equal function.
func (list {{.Prefix.U}}{{.Type.U}}List) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list))
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) IndexWhere(p func({{.Type}}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) IndexWhere2(p func({{.Type}}) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) LastIndexWhere(p func({{.Type}}) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) LastIndexWhere2(p func({{.Type}}) bool, before int) int {
	if before < 0 {
		before = len(list)
	}
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
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
func (list {{.Prefix.U}}{{.Type.U}}List) Equals(other {{.Prefix.U}}{{.Type.U}}List) bool {
	if list == nil {
		return len(other) == 0
	}

	if other == nil {
		return len(list) == 0
	}

	if len(list) != len(other) {
		return false
	}

	for i, v := range list {
		if v != other[i] {
			return false
		}
	}

	return true
}
{{- end}}

//-------------------------------------------------------------------------------------------------

type sortable{{.Prefix.U}}{{.Type.U}}List struct {
	less func(i, j {{.Type}}) bool
	m    []{{.Type}}
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Len() int {
	return len(sl.m)
}

func (sl sortable{{.Prefix.U}}{{.Type.U}}List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) SortBy(less func(i, j {{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	sort.Sort(sortable{{.Prefix.U}}{{.Type.U}}List{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list {{.Prefix.U}}{{.Type.U}}List) StableSortBy(less func(i, j {{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}List {
	sort.Stable(sortable{{.Prefix.U}}{{.Type.U}}List{less, list})
	return list
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) Sorted() {{.Prefix.U}}{{.Type.U}}List {
	return list.SortBy(func(a, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list {{.Prefix.U}}{{.Type.U}}List) StableSorted() {{.Prefix.U}}{{.Type.U}}List {
	return list.StableSortBy(func(a, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.Prefix.U}}{{.Type.U}}List) Min() {{.Type.Name}} {
	m := list.MinBy(func(a {{.Type}}, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
	return {{.Type.Star}}m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.Prefix.U}}{{.Type.U}}List) Max() (result {{.Type.Name}}) {
	m := list.MaxBy(func(a {{.Type}}, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
	return {{.Type.Star}}m
}
{{- end}}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is numeric.

// Sum returns the sum of all the elements in the list.
func (list {{.Prefix.U}}{{.Type.U}}List) Sum() {{.Type.Name}} {
	sum := {{.Type.Name}}(0)
	for _, v := range list {
		sum = sum + {{.Type.Star}}v
	}
	return sum
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list {{.Prefix.U}}{{.Type.U}}List) StringList() []string {
{{- if eq .Type.String "string"}}
	return list
{{- else}}
	strings := make([]string, len(list))
	for i, v := range list {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
{{- end}}
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list {{.Prefix.U}}{{.Type.U}}List) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list {{.Prefix.U}}{{.Type.U}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list {{.Prefix.U}}{{.Type.U}}List) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list {{.Prefix.U}}{{.Type.U}}List) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	for _, v := range list {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------
{{- end}}
{{- if .GobEncode}}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (list {{.Prefix.U}}{{.Type.U}}List) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (list {{.Prefix.U}}{{.Type.U}}List) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list)
	return buf.Bytes(), err
}
{{- end}}
