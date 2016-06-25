// Generated from {{.TemplateFile}} with Type={{.PType}}

package {{.Package}}

import (
	"bytes"
	"fmt"
	"math/rand"
)

// {{.Type}}List is a slice of type {{.PType}}. Use it where you would use []{{.PType}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.Type}}List []{{.PType}}

//-------------------------------------------------------------------------------------------------

// New{{.Type}}List constructs a new list containing the supplied values, if any.
func New{{.Type}}List(values ...{{.PType}}) {{.Type}}List {
	list := make({{.Type}}List, len(values))
	for i, v := range values {
		list[i] = v
	}
	return list
}

// Build{{.Type}}ListFromChan constructs a new {{.Type}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.Type}}ListFromChan(source <-chan {{.PType}}) {{.Type}}List {
	result := make({{.Type}}List, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list {{.Type}}List) Head() {{.PType}} {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list {{.Type}}List) Last() {{.PType}} {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list {{.Type}}List) Tail() {{.Type}}List {
	return {{.Type}}List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list {{.Type}}List) Init() {{.Type}}List {
	return {{.Type}}List(list[:len(list)-1])
}

// IsEmpty tests whether {{.Type}}List is empty.
func (list {{.Type}}List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.Type}}List is empty.
func (list {{.Type}}List) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list {{.Type}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list {{.Type}}List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list {{.Type}}List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list {{.Type}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.Type}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of {{.Type}}List return true for the passed func.
func (list {{.Type}}List) Exists(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.Type}}List return true for the passed func.
func (list {{.Type}}List) Forall(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.Type}}List and executes the passed func against each element.
func (list {{.Type}}List) Foreach(fn func({{.PType}})) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
func (list {{.Type}}List) Send() <-chan {{.PType}} {
	ch := make(chan {{.PType}})
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of {{.Type}}List with all elements in the reverse order.
func (list {{.Type}}List) Reverse() {{.Type}}List {
	numItems := len(list)
	result := make({{.Type}}List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of {{.Type}}List, using a version of the Fisher-Yates shuffle.
func (list {{.Type}}List) Shuffle() {{.Type}}List {
	numItems := len(list)
	result := make({{.Type}}List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a new {{.Type}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) Take(n int) {{.Type}}List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new {{.Type}}List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) Drop(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new {{.Type}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) TakeLast(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new {{.Type}}List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) DropLast(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new {{.Type}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list {{.Type}}List) TakeWhile(p func({{.PType}}) bool) (result {{.Type}}List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.Type}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list {{.Type}}List) DropWhile(p func({{.PType}}) bool) (result {{.Type}}List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new {{.Type}}List whose elements return true for func.
func (list {{.Type}}List) Filter(fn func({{.PType}}) bool) {{.Type}}List {
	result := make({{.Type}}List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list {{.Type}}List) Partition(p func({{.PType}}) bool) ({{.Type}}List, {{.Type}}List) {
	matching := make({{.Type}}List, 0, len(list)/2)
	others := make({{.Type}}List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of {{.Type}}List that return true for the passed predicate.
func (list {{.Type}}List) CountBy(predicate func({{.PType}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.Type}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list {{.Type}}List) MinBy(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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
	result = list[m]
	return
}

// MaxBy returns an element of {{.Type}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list {{.Type}}List) MaxBy(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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
	result = list[m]
	return
}

// DistinctBy returns a new {{.Type}}List whose elements are unique, where equality is defined by a passed func.
func (list {{.Type}}List) DistinctBy(equal func({{.PType}}, {{.PType}}) bool) (result {{.Type}}List) {
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
func (list {{.Type}}List) IndexWhere(p func({{.PType}}) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list {{.Type}}List) IndexWhere2(p func({{.PType}}) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list {{.Type}}List) LastIndexWhere(p func({{.PType}}) bool) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list {{.Type}}List) LastIndexWhere2(p func({{.PType}}) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is not ordered.

// Min returns the first element containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (list {{.Type}}List) Min(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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
	result = list[m]
	return
}

// Max returns the first element containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (list {{.Type}}List) Max(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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
	result = list[m]
	return
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list {{.Type}}List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list {{.Type}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list {{.Type}}List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", *v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", *v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}
