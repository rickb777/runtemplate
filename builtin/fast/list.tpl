// An encapsulated []{{.Type}}.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} Mutable:always

package {{.Package}}

import (
{{if .Stringer}}
	"bytes"
	"fmt" {{- end}}
	"math/rand"
	"sort"
{{- if .HasImport}}
	{{.Import}}
{{end}}
)

// {{.UPrefix}}{{.UType}}List contains a slice of type {{.PType}}. Use it where you would use []{{.PType}}.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.UPrefix}}{{.UType}}List struct {
	m []{{.PType}}
}


//-------------------------------------------------------------------------------------------------

func new{{.UPrefix}}{{.UType}}List(len, cap int) *{{.UPrefix}}{{.UType}}List {
	return &{{.UPrefix}}{{.UType}}List {
		m: make([]{{.PType}}, len, cap),
	}
}

// New{{.UPrefix}}{{.UType}}List constructs a new list containing the supplied values, if any.
func New{{.UPrefix}}{{.UType}}List(values ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
	result := new{{.UPrefix}}{{.UType}}List(len(values), len(values))
	copy(result.m, values)
	return result
}

// Convert{{.UPrefix}}{{.UType}}List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func Convert{{.UPrefix}}{{.UType}}List(values ...interface{}) (*{{.UPrefix}}{{.UType}}List, bool) {
	result := new{{.UPrefix}}{{.UType}}List(0, len(values))
{{if and .Numeric (eq .Type .PType)}}
	for _, i := range values {
		switch i.(type) {
		case int:
			result.m = append(result.m, {{.PType}}(i.(int)))
		case int8:
			result.m = append(result.m, {{.PType}}(i.(int8)))
		case int16:
			result.m = append(result.m, {{.PType}}(i.(int16)))
		case int32:
			result.m = append(result.m, {{.PType}}(i.(int32)))
		case int64:
			result.m = append(result.m, {{.PType}}(i.(int64)))
		case uint:
			result.m = append(result.m, {{.PType}}(i.(uint)))
		case uint8:
			result.m = append(result.m, {{.PType}}(i.(uint8)))
		case uint16:
			result.m = append(result.m, {{.PType}}(i.(uint16)))
		case uint32:
			result.m = append(result.m, {{.PType}}(i.(uint32)))
		case uint64:
			result.m = append(result.m, {{.PType}}(i.(uint64)))
		case float32:
			result.m = append(result.m, {{.PType}}(i.(float32)))
		case float64:
			result.m = append(result.m, {{.PType}}(i.(float64)))
		}
	}
{{else}}
	for _, i := range values {
		v, ok := i.({{.PType}})
		if ok {
			result.m = append(result.m, v)
		}
	}
{{end}}
	return result, len(result.m) == len(values)
}

// Build{{.UPrefix}}{{.UType}}ListFromChan constructs a new {{.UPrefix}}{{.UType}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}ListFromChan(source <-chan {{.PType}}) *{{.UPrefix}}{{.UType}}List {
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *{{.UPrefix}}{{.UType}}List) ToSlice() []{{.PType}} {

	s := make([]{{.PType}}, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *{{.UPrefix}}{{.UType}}List) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *{{.UPrefix}}{{.UType}}List) Clone() *{{.UPrefix}}{{.UType}}List {

	return New{{.UPrefix}}{{.UType}}List(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *{{.UPrefix}}{{.UType}}List) Get(i int) {{.PType}} {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *{{.UPrefix}}{{.UType}}List) Head() {{.PType}} {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *{{.UPrefix}}{{.UType}}List) Last() {{.PType}} {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *{{.UPrefix}}{{.UType}}List) Tail() *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *{{.UPrefix}}{{.UType}}List) Init() *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether {{.UPrefix}}{{.UType}}List is empty.
func (list *{{.UPrefix}}{{.UType}}List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether {{.UPrefix}}{{.UType}}List is empty.
func (list *{{.UPrefix}}{{.UType}}List) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *{{.UPrefix}}{{.UType}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *{{.UPrefix}}{{.UType}}List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list.
func (list *{{.UPrefix}}{{.UType}}List) Size() int {

	return len(list.m)
}

// Swap exchanges two elements.
func (list *{{.UPrefix}}{{.UType}}List) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

{{if .Comparable}}
// Contains determines if a given item is already in the list.
func (list *{{.UPrefix}}{{.UType}}List) Contains(v {{.Type}}) bool {
	return list.Exists(func (x {{.PType}}) bool {
		return {{.TypeStar}}x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *{{.UPrefix}}{{.UType}}List) ContainsAll(i ...{{.Type}}) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

{{end -}}
// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}List return true for the predicate p.
func (list *{{.UPrefix}}{{.UType}}List) Exists(p func({{.PType}}) bool) bool {

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.UPrefix}}{{.UType}}List return true for the predicate p.
func (list *{{.UPrefix}}{{.UType}}List) Forall(p func({{.PType}}) bool) bool {

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.UPrefix}}{{.UType}}List and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (list *{{.UPrefix}}{{.UType}}List) Foreach(fn func({{.PType}})) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *{{.UPrefix}}{{.UType}}List) Send() <-chan {{.PType}} {
	ch := make(chan {{.PType}})
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of {{.UPrefix}}{{.UType}}List with all elements in the reverse order.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) Reverse() *{{.UPrefix}}{{.UType}}List {
	return list.Clone().doReverse()
}

// DoReverse alters a {{.UPrefix}}{{.UType}}List with all elements in the reverse order.
//
// The modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) DoReverse() *{{.UPrefix}}{{.UType}}List {
	return list.doReverse()
}

func (list *{{.UPrefix}}{{.UType}}List) doReverse() *{{.UPrefix}}{{.UType}}List {
	mid := (len(list.m) + 1) / 2
	last := len(list.m) - 1
	for i := 0; i < mid; i++ {
	    r := last - i
	    if i != r {
		    list.m[i], list.m[r] = list.m[r], list.m[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of {{.UPrefix}}{{.UType}}List, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) Shuffle() *{{.UPrefix}}{{.UType}}List {
	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled {{.UPrefix}}{{.UType}}List, using a version of the Fisher-Yates shuffle.
//
// The modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) DoShuffle() *{{.UPrefix}}{{.UType}}List {
	return list.doShuffle()
}

func (list *{{.UPrefix}}{{.UType}}List) doShuffle() *{{.UPrefix}}{{.UType}}List {
	numItems := len(list.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
        list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current list. This is a synonym for Append.
func (list *{{.UPrefix}}{{.UType}}List) Add(more ...{{.PType}}) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *{{.UPrefix}}{{.UType}}List) Append(more ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
	return list.doAppend(more...)
}

func (list *{{.UPrefix}}{{.UType}}List) doAppend(more ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a {{.UPrefix}}{{.UType}}List by inserting elements at a given index.
// This is a generalised version of Append.
//
// The modified list is returned.
// Panics if the index is out of range.
func (list *{{.UPrefix}}{{.UType}}List) DoInsertAt(index int, more ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
    return list.doInsertAt(index, more...)
}

func (list *{{.UPrefix}}{{.UType}}List) doInsertAt(index int, more ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
    if len(more) == 0 {
        return list
    }

    if index == len(list.m) {
        // appending is an easy special case
    	return list.doAppend(more...)
    }

	newlist := make([]{{.PType}}, 0, len(list.m) + len(more))

    if index != 0 {
        newlist = append(newlist, list.m[:index]...)
    }

    newlist = append(newlist, more...)

    newlist = append(newlist, list.m[index:]...)

    list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a {{.UPrefix}}{{.UType}}List by deleting n elements from the start of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *{{.UPrefix}}{{.UType}}List) DoDeleteFirst(n int) *{{.UPrefix}}{{.UType}}List {
    return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a {{.UPrefix}}{{.UType}}List by deleting n elements from the end of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *{{.UPrefix}}{{.UType}}List) DoDeleteLast(n int) *{{.UPrefix}}{{.UType}}List {
    return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a {{.UPrefix}}{{.UType}}List by deleting n elements from a given index.
//
// The modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *{{.UPrefix}}{{.UType}}List) DoDeleteAt(index, n int) *{{.UPrefix}}{{.UType}}List {
    return list.doDeleteAt(index, n)
}

func (list *{{.UPrefix}}{{.UType}}List) doDeleteAt(index, n int) *{{.UPrefix}}{{.UType}}List {
    if n == 0 {
        return list
    }

	newlist := make([]{{.PType}}, 0, len(list.m) - n)

    if index != 0 {
        newlist = append(newlist, list.m[:index]...)
    }

    index += n

    if index != len(list.m) {
        newlist = append(newlist, list.m[index:]...)
    }

    list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a {{.UPrefix}}{{.UType}}List by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) DoKeepWhere(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {
    return list.doKeepWhere(p)
}

func (list *{{.UPrefix}}{{.UType}}List) doKeepWhere(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {
	result := make([]{{.PType}}, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

    list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.UPrefix}}{{.UType}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *{{.UPrefix}}{{.UType}}List) Take(n int) *{{.UPrefix}}{{.UType}}List {

	if n > len(list.m) {
		return list
	}
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of {{.UPrefix}}{{.UType}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) Drop(n int) *{{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}


	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of {{.UPrefix}}{{.UType}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) TakeLast(n int) *{{.UPrefix}}{{.UType}}List {

	l := len(list.m)
	if n > l {
		return list
	}
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of {{.UPrefix}}{{.UType}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) DropLast(n int) *{{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}


	l := len(list.m)
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new {{.UPrefix}}{{.UType}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) TakeWhile(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new {{.UPrefix}}{{.UType}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *{{.UPrefix}}{{.UType}}List) DropWhile(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	adding := false

	for _, v := range list.m {
		if !p(v) || adding {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type}} that returns true for predicate p.
// False is returned if none match.
func (list {{.UPrefix}}{{.UType}}List) Find(p func({{.PType}}) bool) ({{.PType}}, bool) {

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

{{if eq .TypeStar "*"}}
	return nil, false
{{else}}
	var empty {{.Type}}
	return empty, false
{{end}}
}

// Filter returns a new {{.UPrefix}}{{.UType}}List whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *{{.UPrefix}}{{.UType}}List) Filter(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *{{.UPrefix}}{{.UType}}List) Partition(p func({{.PType}}) bool) (*{{.UPrefix}}{{.UType}}List, *{{.UPrefix}}{{.UType}}List) {

	matching := new{{.UPrefix}}{{.UType}}List(0, len(list.m)/2)
	others := new{{.UPrefix}}{{.UType}}List(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
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
func (list *{{.UPrefix}}{{.UType}}List) Map(fn func({{.PType}}) {{.PType}}) *{{.UPrefix}}{{.UType}}List {
	result := new{{.UPrefix}}{{.UType}}List(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new {{.UPrefix}}{{.UType}}List by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *{{.UPrefix}}{{.UType}}List) FlatMap(fn func({{.PType}}) []{{.PType}}) *{{.UPrefix}}{{.UType}}List {
	result := new{{.UPrefix}}{{.UType}}List(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}List that return true for the passed predicate.
func (list *{{.UPrefix}}{{.UType}}List) CountBy(predicate func({{.PType}}) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.UPrefix}}{{.UType}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *{{.UPrefix}}{{.UType}}List) MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}} {

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

// MaxBy returns an element of {{.UPrefix}}{{.UType}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *{{.UPrefix}}{{.UType}}List) MaxBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}} {

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

// DistinctBy returns a new {{.UPrefix}}{{.UType}}List whose elements are unique, where equality is defined by a passed func.
func (list *{{.UPrefix}}{{.UType}}List) DistinctBy(equal func({{.PType}}, {{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, len(list.m))
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

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list *{{.UPrefix}}{{.UType}}List) IndexWhere(p func({{.PType}}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *{{.UPrefix}}{{.UType}}List) IndexWhere2(p func({{.PType}}) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *{{.UPrefix}}{{.UType}}List) LastIndexWhere(p func({{.PType}}) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *{{.UPrefix}}{{.UType}}List) LastIndexWhere2(p func({{.PType}}) bool, before int) int {

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

{{if .Numeric}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the list.
func (list *{{.UPrefix}}{{.UType}}List) Sum() {{.Type}} {

	sum := {{.Type}}(0)
	for _, v := range list.m {
		sum = sum + {{.TypeStar}}v
	}
	return sum
}

{{end -}}
{{if .Comparable}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *{{.UPrefix}}{{.UType}}List) Equals(other *{{.UPrefix}}{{.UType}}List) bool {

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

{{end -}}
//-------------------------------------------------------------------------------------------------

type sortable{{.UPrefix}}{{.UType}}List struct {
	less func(i, j {{.Type}}) bool
	m []{{.PType}}
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Less(i, j int) bool {
	return sl.less({{.TypeStar}}sl.m[i], {{.TypeStar}}sl.m[j])
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Len() int {
	return len(sl.m)
}

func (sl sortable{{.UPrefix}}{{.UType}}List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) SortBy(less func(i, j {{.Type}}) bool) *{{.UPrefix}}{{.UType}}List {

	sort.Sort(sortable{{.UPrefix}}{{.UType}}List{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *{{.UPrefix}}{{.UType}}List) StableSortBy(less func(i, j {{.Type}}) bool) *{{.UPrefix}}{{.UType}}List {

	sort.Stable(sortable{{.UPrefix}}{{.UType}}List{less, list.m})
	return list
}

{{if .Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) Sorted() *{{.UPrefix}}{{.UType}}List {
	return list.SortBy(func(a, b {{.Type}}) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *{{.UPrefix}}{{.UType}}List) StableSorted() *{{.UPrefix}}{{.UType}}List {
	return list.StableSortBy(func(a, b {{.Type}}) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *{{.UPrefix}}{{.UType}}List) Min() {{.Type}} {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := {{.TypeStar}}v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if {{.TypeStar}}v < m {
			m = {{.TypeStar}}v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *{{.UPrefix}}{{.UType}}List) Max() (result {{.Type}}) {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	v := list.m[0]
	m := {{.TypeStar}}v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if {{.TypeStar}}v > m {
			m = {{.TypeStar}}v
		}
	}
	return m
}

{{end -}}
{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list {{.UPrefix}}{{.UType}}List) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *{{.UPrefix}}{{.UType}}List) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list {{.UPrefix}}{{.UType}}List) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *{{.UPrefix}}{{.UType}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *{{.UPrefix}}{{.UType}}List) MkString3(before, between, after string) string {
	return list.mkString3Bytes(before, between, after).String()
}

func (list {{.UPrefix}}{{.UType}}List) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
{{end}}
