// A simple type derived from []{{.Type}}
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if .Stringer}}
	"bytes"
	"fmt"
{{- end}}
	"math/rand"
	"sort"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.UPrefix}}{{.UType}}List is a slice of type {{.PType}}. Use it where you would use []{{.PType}}.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.UPrefix}}{{.UType}}List []{{.PType}}

//-------------------------------------------------------------------------------------------------

// Make{{.UPrefix}}{{.UType}}List makes an empty list with both length and capacity initialised.
func Make{{.UPrefix}}{{.UType}}List(length, capacity int) {{.UPrefix}}{{.UType}}List {
	return make({{.UPrefix}}{{.UType}}List, length, capacity)
}

// New{{.UPrefix}}{{.UType}}List constructs a new list containing the supplied values, if any.
func New{{.UPrefix}}{{.UType}}List(values ...{{.PType}}) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(len(values), len(values))
	copy(result, values)
	return result
}

// Convert{{.UPrefix}}{{.UType}}List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func Convert{{.UPrefix}}{{.UType}}List(values ...interface{}) ({{.UPrefix}}{{.UType}}List, bool) {
	result := Make{{.UPrefix}}{{.UType}}List(0, len(values))
{{if and .Numeric (eq .Type .PType)}}
	for _, i := range values {
		switch i.(type) {
		case int:
			result = append(result, {{.PType}}(i.(int)))
		case int8:
			result = append(result, {{.PType}}(i.(int8)))
		case int16:
			result = append(result, {{.PType}}(i.(int16)))
		case int32:
			result = append(result, {{.PType}}(i.(int32)))
		case int64:
			result = append(result, {{.PType}}(i.(int64)))
		case uint:
			result = append(result, {{.PType}}(i.(uint)))
		case uint8:
			result = append(result, {{.PType}}(i.(uint8)))
		case uint16:
			result = append(result, {{.PType}}(i.(uint16)))
		case uint32:
			result = append(result, {{.PType}}(i.(uint32)))
		case uint64:
			result = append(result, {{.PType}}(i.(uint64)))
		case float32:
			result = append(result, {{.PType}}(i.(float32)))
		case float64:
			result = append(result, {{.PType}}(i.(float64)))
		}
	}
{{else}}
	for _, i := range values {
		v, ok := i.({{.PType}})
		if ok {
			result = append(result, v)
		}
	}
{{end}}
	return result, len(result) == len(values)
}

// Build{{.UPrefix}}{{.UType}}ListFromChan constructs a new {{.UPrefix}}{{.UType}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}ListFromChan(source <-chan {{.PType}}) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list {{.UPrefix}}{{.UType}}List) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list {{.UPrefix}}{{.UType}}List) Clone() {{.UPrefix}}{{.UType}}List {
	return New{{.UPrefix}}{{.UType}}List(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list {{.UPrefix}}{{.UType}}List) Get(i int) {{.PType}} {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty.
func (list {{.UPrefix}}{{.UType}}List) Head() {{.PType}} {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}}.
func (list {{.UPrefix}}{{.UType}}List) HeadOption() {{.PType}} {
	if list.IsEmpty() {
		return {{.TypeZero}}
	}
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty.
func (list {{.UPrefix}}{{.UType}}List) Last() {{.PType}} {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}}.
func (list {{.UPrefix}}{{.UType}}List) LastOption() {{.PType}} {
	if list.IsEmpty() {
		return {{.TypeZero}}
	}
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty.
func (list {{.UPrefix}}{{.UType}}List) Tail() {{.UPrefix}}{{.UType}}List {
	return {{.UPrefix}}{{.UType}}List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty.
func (list {{.UPrefix}}{{.UType}}List) Init() {{.UPrefix}}{{.UType}}List {
	return {{.UPrefix}}{{.UType}}List(list[:len(list)-1])
}

// IsEmpty tests whether {{.UPrefix}}{{.UType}}List is empty.
func (list {{.UPrefix}}{{.UType}}List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether {{.UPrefix}}{{.UType}}List is empty.
func (list {{.UPrefix}}{{.UType}}List) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list {{.UPrefix}}{{.UType}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list {{.UPrefix}}{{.UType}}List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list {{.UPrefix}}{{.UType}}List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list {{.UPrefix}}{{.UType}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.UPrefix}}{{.UType}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------
{{- if .Comparable}}

// Contains determines if a given item is already in the list.
func (list {{.UPrefix}}{{.UType}}List) Contains(v {{.Type}}) bool {
	return list.Exists(func (x {{.PType}}) bool {
		return {{.TypeStar}}x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list {{.UPrefix}}{{.UType}}List) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}
{{- end}}

// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}List return true for the passed func.
func (list {{.UPrefix}}{{.UType}}List) Exists(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.UPrefix}}{{.UType}}List return true for the passed func.
func (list {{.UPrefix}}{{.UType}}List) Forall(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.UPrefix}}{{.UType}}List and executes function fn against each element.
func (list {{.UPrefix}}{{.UType}}List) Foreach(fn func({{.PType}})) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list {{.UPrefix}}{{.UType}}List) Send() <-chan {{.PType}} {
	ch := make(chan {{.PType}})
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of {{.UPrefix}}{{.UType}}List with all elements in the reverse order.
func (list {{.UPrefix}}{{.UType}}List) Reverse() {{.UPrefix}}{{.UType}}List {
	numItems := len(list)
	result := Make{{.UPrefix}}{{.UType}}List(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse returns a copy of {{.UPrefix}}{{.UType}}List with all elements in the reverse order.
// This is an alias for Reverse.
func (list {{.UPrefix}}{{.UType}}List) DoReverse() {{.UPrefix}}{{.UType}}List {
	return list.Reverse()
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of {{.UPrefix}}{{.UType}}List, using a version of the Fisher-Yates shuffle.
func (list {{.UPrefix}}{{.UType}}List) Shuffle() {{.UPrefix}}{{.UType}}List {
	result := list.Clone()
	numItems := len(list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.UPrefix}}{{.UType}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list {{.UPrefix}}{{.UType}}List) Take(n int) {{.UPrefix}}{{.UType}}List {
	if n > len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of {{.UPrefix}}{{.UType}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list {{.UPrefix}}{{.UType}}List) Drop(n int) {{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of {{.UPrefix}}{{.UType}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list {{.UPrefix}}{{.UType}}List) TakeLast(n int) {{.UPrefix}}{{.UType}}List {
	l := len(list)
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of {{.UPrefix}}{{.UType}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list {{.UPrefix}}{{.UType}}List) DropLast(n int) {{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new {{.UPrefix}}{{.UType}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list {{.UPrefix}}{{.UType}}List) TakeWhile(p func({{.PType}}) bool) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new {{.UPrefix}}{{.UType}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list {{.UPrefix}}{{.UType}}List) DropWhile(p func({{.PType}}) bool) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, 0)
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

// Find returns the first {{.Type}} that returns true for predicate p.
// False is returned if none match.
func (list {{.UPrefix}}{{.UType}}List) Find(p func({{.PType}}) bool) ({{.PType}}, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}

{{- if eq .TypeStar "*"}}
	return nil, false
{{- else}}
	var empty {{.Type}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.UPrefix}}{{.UType}}List whose elements return true for predicate p.
//
// The original list is not modified.
func (list {{.UPrefix}}{{.UType}}List) Filter(p func({{.PType}}) bool) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list {{.UPrefix}}{{.UType}}List) Partition(p func({{.PType}}) bool) ({{.UPrefix}}{{.UType}}List, {{.UPrefix}}{{.UType}}List) {
	matching := Make{{.UPrefix}}{{.UType}}List(0, len(list)/2)
	others := Make{{.UPrefix}}{{.UType}}List(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new {{.UPrefix}}{{.UType}}List by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{.UPrefix}}{{.UType}}List) Map(fn func({{.PType}}) {{.PType}}) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, len(list))

	for _, v := range list {
		result = append(result, fn(v))
	}

	return result
}

// FlatMap returns a new {{.UPrefix}}{{.UType}}List by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list {{.UPrefix}}{{.UType}}List) FlatMap(fn func({{.PType}}) []{{.PType}}) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, len(list))

	for _, v := range list {
		result = append(result, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}List that return true for the passed predicate.
func (list {{.UPrefix}}{{.UType}}List) CountBy(predicate func({{.PType}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.UPrefix}}{{.UType}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list {{.UPrefix}}{{.UType}}List) MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}} {
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

// MaxBy returns an element of {{.UPrefix}}{{.UType}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list {{.UPrefix}}{{.UType}}List) MaxBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}} {
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

// DistinctBy returns a new {{.UPrefix}}{{.UType}}List whose elements are unique, where equality is defined by a passed func.
func (list {{.UPrefix}}{{.UType}}List) DistinctBy(equal func({{.PType}}, {{.PType}}) bool) {{.UPrefix}}{{.UType}}List {
	result := Make{{.UPrefix}}{{.UType}}List(0, len(list))
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

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list {{.UPrefix}}{{.UType}}List) IndexWhere(p func({{.PType}}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list {{.UPrefix}}{{.UType}}List) IndexWhere2(p func({{.PType}}) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list {{.UPrefix}}{{.UType}}List) LastIndexWhere(p func({{.PType}}) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list {{.UPrefix}}{{.UType}}List) LastIndexWhere2(p func({{.PType}}) bool, before int) int {
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
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the list.
func (list {{.UPrefix}}{{.UType}}List) Sum() {{.Type}} {
	sum := {{.Type}}(0)
	for _, v := range list {
		sum = sum + {{.TypeStar}}v
	}
	return sum
}
{{- end}}
{{- if .Comparable}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list {{.UPrefix}}{{.UType}}List) Equals(other {{.UPrefix}}{{.UType}}List) bool {
	if list.Size() != other.Size() {
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

type sortable{{.UPrefix}}{{.UType}}List struct {
	less func(i, j {{.PType}}) bool
	m []{{.PType}}
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Len() int {
	return len(sl.m)
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list {{.UPrefix}}{{.UType}}List) SortBy(less func(i, j {{.PType}}) bool) {{.UPrefix}}{{.UType}}List {

	sort.Sort(sortable{{.UPrefix}}{{.UType}}List{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list {{.UPrefix}}{{.UType}}List) StableSortBy(less func(i, j {{.PType}}) bool) {{.UPrefix}}{{.UType}}List {

	sort.Stable(sortable{{.UPrefix}}{{.UType}}List{less, list})
	return list
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
func (list {{.UPrefix}}{{.UType}}List) Sorted() {{.UPrefix}}{{.UType}}List {
	return list.SortBy(func(a, b {{.PType}}) bool {
		return {{.TypeStar}}a < {{.TypeStar}}b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
func (list {{.UPrefix}}{{.UType}}List) StableSorted() {{.UPrefix}}{{.UType}}List {
	return list.StableSortBy(func(a, b {{.PType}}) bool {
		return {{.TypeStar}}a < {{.TypeStar}}b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UPrefix}}{{.UType}}List) Min() {{.Type}} {
	m := list.MinBy(func(a {{.PType}}, b {{.PType}}) bool {
		return {{.TypeStar}}a < {{.TypeStar}}b
	})
	return {{.TypeStar}}m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UPrefix}}{{.UType}}List) Max() (result {{.Type}}) {
	m := list.MaxBy(func(a {{.PType}}, b {{.PType}}) bool {
		return {{.TypeStar}}a < {{.TypeStar}}b
	})
	return {{.TypeStar}}m
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list {{.UPrefix}}{{.UType}}List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list {{.UPrefix}}{{.UType}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list {{.UPrefix}}{{.UType}}List) MkString3(before, between, after string) string {
	b := bytes.Buffer{}
	b.WriteString(before)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(between)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(after)
	return b.String()
}
{{- end}}
