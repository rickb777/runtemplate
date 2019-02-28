// An encapsulated []{{.Type}}.
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}}
// GobEncode:{{.GobEncode}} Mutable:always ToList:always ToSet:{{.ToSet}} MapTo:{{.MapTo}}
// by runtemplate {{.AppVersion}}
// See https://github.com/johanbrandhorst/runtemplate/blob/master/BUILTIN.md

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

// {{.Prefix.U}}{{.Type.U}}List contains a slice of type {{.Type}}.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.Prefix.U}}{{.Type.U}}List struct {
	m []{{.Type}}
}

//-------------------------------------------------------------------------------------------------

// Make{{.Prefix.U}}{{.Type.U}}List makes an empty list with both length and capacity initialised.
func Make{{.Prefix.U}}{{.Type.U}}List(length, capacity int) *{{.Prefix.U}}{{.Type.U}}List {
	return &{{.Prefix.U}}{{.Type.U}}List{
		m: make([]{{.Type}}, length, capacity),
	}
}

// New{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
func New{{.Prefix.U}}{{.Type.U}}List(values ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	list := Make{{.Prefix.U}}{{.Type.U}}List(len(values), len(values))
	copy(list.m, values)
	return list
}

// Convert{{.Prefix.U}}{{.Type.U}}List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func Convert{{.Prefix.U}}{{.Type.U}}List(values ...interface{}) (*{{.Prefix.U}}{{.Type.U}}List, bool) {
	list := Make{{.Prefix.U}}{{.Type.U}}List(0, len(values))

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
		{{- if .Type.IsPtr}}
        case {{.Type.Name}}:
			list.m = append(list.m, &j)
        case *{{.Type.Name}}:
			list.m = append(list.m, j)
		{{- else}}
        case {{.Type}}:
			list.m = append(list.m, j)
        case *{{.Type}}:
			list.m = append(list.m, *j)
		{{- end}}
{{- end}}
		}
	}

	return list, len(list.m) == len(values)
}

// Build{{.Prefix.U}}{{.Type.U}}ListFromChan constructs a new {{.Prefix.U}}{{.Type.U}}List from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}ListFromChan(source <-chan {{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	list := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
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
func (list *{{.Prefix.U}}{{.Type.U}}List) slice() []{{.Type}} {
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
func (list *{{.Prefix.U}}{{.Type.U}}List) ToSlice() []{{.Type}} {
	if list == nil {
		return nil
	}

	s := make([]{{.Type}}, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *{{.Prefix.U}}{{.Type.U}}List) ToInterfaceSlice() []interface{} {
	if list == nil {
		return nil
	}

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) Clone() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	return New{{.Prefix.U}}{{.Type.U}}List(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Get(i int) {{.Type}} {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Head() {{.Type}} {

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}}.
func (list *{{.Prefix.U}}{{.Type.U}}List) HeadOption() {{.Type}} {
	if list == nil {
		return {{.Type.Zero}}
	}

	if len(list.m) == 0 {
		return {{.Type.Zero}}
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Last() {{.Type}} {

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}}.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastOption() {{.Type}} {
	if list == nil {
		return {{.Type.Zero}}
	}

	if len(list.m) == 0 {
		return {{.Type.Zero}}
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Tail() *{{.Prefix.U}}{{.Type.U}}List {

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *{{.Prefix.U}}{{.Type.U}}List) Init() *{{.Prefix.U}}{{.Type.U}}List {

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
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
// This is one of the three methods in the standard sort.Interface.
func (list *{{.Prefix.U}}{{.Type.U}}List) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *{{.Prefix.U}}{{.Type.U}}List) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------
{{- if .Comparable}}

// Contains determines whether a given item is already in the list, returning true if so.
func (list *{{.Prefix.U}}{{.Type.U}}List) Contains(v {{.Type}}) bool {
	return list.Exists(func(x {{.Type}}) bool {
		return {{.Type.Star}}v == {{.Type.Star}}x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *{{.Prefix.U}}{{.Type.U}}List) ContainsAll(i ...{{.Type}}) bool {
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
func (list *{{.Prefix.U}}{{.Type.U}}List) Exists(p func({{.Type}}) bool) bool {
	if list == nil {
		return false
	}

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.Prefix.U}}{{.Type.U}}List return true for the predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) Forall(p func({{.Type}}) bool) bool {
	if list == nil {
		return true
	}

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.Prefix.U}}{{.Type.U}}List and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *{{.Prefix.U}}{{.Type.U}}List) Foreach(f func({{.Type}})) {
	if list == nil {
		return
	}

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list *{{.Prefix.U}}{{.Type.U}}List) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
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
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) Reverse() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := Make{{.Prefix.U}}{{.Type.U}}List(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a {{.Prefix.U}}{{.Type.U}}List with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoReverse() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

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

// Shuffle returns a shuffled copy of {{.Prefix.U}}{{.Type.U}}List, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) Shuffle() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled {{.Prefix.U}}{{.Type.U}}List, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoShuffle() *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	return list.doShuffle()
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doShuffle() *{{.Prefix.U}}{{.Type.U}}List {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *{{.Prefix.U}}{{.Type.U}}List) Clear() {
	if list != nil {
	    list.m = list.m[:]
    }
}

// Add adds items to the current list. This is a synonym for Append.
func (list *{{.Prefix.U}}{{.Type.U}}List) Add(more ...{{.Type}}) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) Append(more ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = Make{{.Prefix.U}}{{.Type.U}}List(0, len(more))
	}

	return list.doAppend(more...)
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doAppend(more ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a {{.Prefix.U}}{{.Type.U}}List by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoInsertAt(index int, more ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = Make{{.Prefix.U}}{{.Type.U}}List(0, len(more))
		return list.doInsertAt(index, more...)
	}

	return list.doInsertAt(index, more...)
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doInsertAt(index int, more ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]{{.Type}}, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a {{.Prefix.U}}{{.Type.U}}List by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoDeleteFirst(n int) *{{.Prefix.U}}{{.Type.U}}List {
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a {{.Prefix.U}}{{.Type.U}}List by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoDeleteLast(n int) *{{.Prefix.U}}{{.Type.U}}List {
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a {{.Prefix.U}}{{.Type.U}}List by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoDeleteAt(index, n int) *{{.Prefix.U}}{{.Type.U}}List {
	return list.doDeleteAt(index, n)
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doDeleteAt(index, n int) *{{.Prefix.U}}{{.Type.U}}List {
	if n == 0 {
		return list
	}

	newlist := make([]{{.Type}}, 0, len(list.m)-n)

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

// DoKeepWhere modifies a {{.Prefix.U}}{{.Type.U}}List by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) DoKeepWhere(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	return list.doKeepWhere(p)
}

func (list *{{.Prefix.U}}{{.Type.U}}List) doKeepWhere(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	result := make([]{{.Type}}, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) Take(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	if n >= len(list.m) {
		return list
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of {{.Prefix.U}}{{.Type.U}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) Drop(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of {{.Prefix.U}}{{.Type.U}}List containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) TakeLast(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of {{.Prefix.U}}{{.Type.U}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) DropLast(n int) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new {{.Prefix.U}}{{.Type.U}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) TakeWhile(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
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
//
// The original list is not modified.
func (list *{{.Prefix.U}}{{.Type.U}}List) DropWhile(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, 0)
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
func (list *{{.Prefix.U}}{{.Type.U}}List) Find(p func({{.Type}}) bool) ({{.Type}}, bool) {
	if list == nil {
		return {{.Type.Zero}}, false
	}

	for _, v := range list.m {
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
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *{{.Prefix.U}}{{.Type.U}}List) Filter(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new {{.Prefix.U}}{{.Type.U}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *{{.Prefix.U}}{{.Type.U}}List) Partition(p func({{.Type}}) bool) (*{{.Prefix.U}}{{.Type.U}}List, *{{.Prefix.U}}{{.Type.U}}List) {
	if list == nil {
		return nil, nil
	}

	matching := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))
	others := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))

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
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *{{.Prefix.U}}{{.Type.U}}List) Map(f func({{.Type}}) {{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
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
func (list *{{$.Prefix.U}}{{$.Type.U}}List) MapTo{{.U}}(f func({{$.Type}}) {{.}}) []{{.}} {
	if list == nil {
		return nil
	}

	result := make([]{{.}}, len(list.m))

	for i, v := range list.m {
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
func (list *{{.Prefix.U}}{{.Type.U}}List) FlatMap(f func({{.Type}}) []{{.Type}}) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
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
func (list *{{$.Prefix.U}}{{$.Type.U}}List) FlatMapTo{{.U}}(f func({{$.Type}}) []{{.}}) []{{.}} {
	if list == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(list.m))

	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}
{{- end}}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}List that return true for the predicate p.
func (list *{{.Prefix.U}}{{.Type.U}}List) CountBy(p func({{.Type}}) bool) (result int) {
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
func (list *{{.Prefix.U}}{{.Type.U}}List) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {

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
func (list *{{.Prefix.U}}{{.Type.U}}List) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {

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
func (list *{{.Prefix.U}}{{.Type.U}}List) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	result := Make{{.Prefix.U}}{{.Type.U}}List(0, len(list.m))
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
func (list *{{.Prefix.U}}{{.Type.U}}List) IndexWhere(p func({{.Type}}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) IndexWhere2(p func({{.Type}}) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastIndexWhere(p func({{.Type}}) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) LastIndexWhere2(p func({{.Type}}) bool, before int) int {

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
		if other == nil {
			return true
		}
		return len(other.m) == 0
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
	less func(i, j {{.Type}}) bool
	m []{{.Type}}
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
func (list *{{.Prefix.U}}{{.Type.U}}List) SortBy(less func(i, j {{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	sort.Sort(sortable{{.Prefix.U}}{{.Type.U}}List{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *{{.Prefix.U}}{{.Type.U}}List) StableSortBy(less func(i, j {{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}List {
	if list == nil {
		return nil
	}

	sort.Stable(sortable{{.Prefix.U}}{{.Type.U}}List{less, list.m})
	return list
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) Sorted() *{{.Prefix.U}}{{.Type.U}}List {
	return list.SortBy(func(a, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *{{.Prefix.U}}{{.Type.U}}List) StableSorted() *{{.Prefix.U}}{{.Type.U}}List {
	return list.StableSortBy(func(a, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
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
{{- if eq .Type.String "string"}}
	return list.ToSlice()
{{- else}}
	if list == nil {
		return nil
	}

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
{{- end}}
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
