// An encapsulated []{{.Type}}.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} Mutable:disabled

package {{.Package}}

import (
{{if .Stringer}}
	"bytes"
	"fmt" {{- end}}
	"math/rand"
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

// Build{{.UPrefix}}{{.UType}}ListFromChan constructs a new {{.UPrefix}}{{.UType}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}ListFromChan(source <-chan {{.PType}}) *{{.UPrefix}}{{.UType}}List {
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (list *{{.UPrefix}}{{.UType}}List) ToSlice() []{{.PType}} {
	s := make([]{{.PType}}, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// Clone returns the same list, which is immutable.
func (list *{{.UPrefix}}{{.UType}}List) Clone() *{{.UPrefix}}{{.UType}}List {
	return list
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
	return list.Len() == 0
}

// NonEmpty tests whether {{.UPrefix}}{{.UType}}List is empty.
func (list *{{.UPrefix}}{{.UType}}List) NonEmpty() bool {
	return list.Len() > 0
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

// Size returns the number of items in the list - an alias of Len().
func (list *{{.UPrefix}}{{.UType}}List) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This implements one of the methods needed by sort.Interface (along with Less and Swap).
func (list *{{.UPrefix}}{{.UType}}List) Len() int {

	return len(list.m)
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
// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}List return true for the passed func.
func (list *{{.UPrefix}}{{.UType}}List) Exists(fn func({{.PType}}) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.UPrefix}}{{.UType}}List return true for the passed func.
func (list *{{.UPrefix}}{{.UType}}List) Forall(fn func({{.PType}}) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.UPrefix}}{{.UType}}List and executes the passed func against each element.
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

// Reverse returns a copy of {{.UPrefix}}{{.UType}}List with all elements in the reverse order.
func (list *{{.UPrefix}}{{.UType}}List) Reverse() *{{.UPrefix}}{{.UType}}List {

	numItems := len(list.m)
	result := new{{.UPrefix}}{{.UType}}List(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of {{.UPrefix}}{{.UType}}List, using a version of the Fisher-Yates shuffle.
func (list *{{.UPrefix}}{{.UType}}List) Shuffle() *{{.UPrefix}}{{.UType}}List {
	result := New{{.UPrefix}}{{.UType}}List(list.m...)
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *{{.UPrefix}}{{.UType}}List) Append(more ...{{.PType}}) *{{.UPrefix}}{{.UType}}List {
	newList := New{{.UPrefix}}{{.UType}}List(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *{{.UPrefix}}{{.UType}}List) doAppend(more ...{{.PType}}) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of {{.UPrefix}}{{.UType}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *{{.UPrefix}}{{.UType}}List) Take(n int) *{{.UPrefix}}{{.UType}}List {

	if n > list.Len() {
		return list
	}
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of {{.UPrefix}}{{.UType}}List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *{{.UPrefix}}{{.UType}}List) Drop(n int) *{{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}


	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of {{.UPrefix}}{{.UType}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *{{.UPrefix}}{{.UType}}List) TakeLast(n int) *{{.UPrefix}}{{.UType}}List {

	l := list.Len()
	if n > l {
		return list
	}
	result := new{{.UPrefix}}{{.UType}}List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of {{.UPrefix}}{{.UType}}List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *{{.UPrefix}}{{.UType}}List) DropLast(n int) *{{.UPrefix}}{{.UType}}List {
	if n == 0 {
		return list
	}


	l := list.Len()
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new {{.UPrefix}}{{.UType}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
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
// elemense are added.
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

// Find returns the first {{.Type}} that returns true for some function.
// False is returned if none match.
func (list {{.UPrefix}}{{.UType}}List) Find(fn func({{.PType}}) bool) ({{.PType}}, bool) {

	for _, v := range list.m {
		if fn(v) {
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

// Filter returns a new {{.UPrefix}}{{.UType}}List whose elements return true for func.
func (list *{{.UPrefix}}{{.UType}}List) Filter(fn func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, list.Len()/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *{{.UPrefix}}{{.UType}}List) Partition(p func({{.PType}}) bool) (*{{.UPrefix}}{{.UType}}List, *{{.UPrefix}}{{.UType}}List) {

	matching := new{{.UPrefix}}{{.UType}}List(0, list.Len()/2)
	others := new{{.UPrefix}}{{.UType}}List(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
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

{{if .Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Less returns true if the element at index i is less than the element at index j.
// This implements one of the methods needed by sort.Interface (along with Len and Swap).
// Panics if i or j is out of range.
func (list *{{.UPrefix}}{{.UType}}List) Less(i, j int) bool {
	return {{.TypeStar}}list.m[i] < {{.TypeStar}}list.m[j]
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *{{.UPrefix}}{{.UType}}List) Min() {{.Type}} {

	l := list.Len()
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

	l := list.Len()
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

{{else -}}
// MinBy returns an element of {{.UPrefix}}{{.UType}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *{{.UPrefix}}{{.UType}}List) MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}} {

	l := list.Len()
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

	l := list.Len()
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

{{end -}}
// DistinctBy returns a new {{.UPrefix}}{{.UType}}List whose elements are unique, where equality is defined by a passed func.
func (list *{{.UPrefix}}{{.UType}}List) DistinctBy(equal func({{.PType}}, {{.PType}}) bool) *{{.UPrefix}}{{.UType}}List {

	result := new{{.UPrefix}}{{.UType}}List(0, list.Len())
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
	return list.LastIndexWhere2(p, len(list.m))
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

	if list.Size() != other.Size() {
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
func (list *{{.UPrefix}}{{.UType}}List) MkString3(pfx, mid, sfx string) string {
	return list.mkString3Bytes(pfx, mid, sfx).String()
}

func (list {{.UPrefix}}{{.UType}}List) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
{{end}}