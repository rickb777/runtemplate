// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable={{.Comparable}} Numeric={{.Numeric}} Ordered={{.Ordered}} Stringer={{.Stringer}}

package {{.Package}}

import (
{{if .Stringer}}
    "bytes"
    "fmt"
{{end}}
	"math/rand"
)

// {{.UType}}List is a slice of type {{.PType}}. Use it where you would use []{{.PType}}.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.UType}}List []{{.PType}}

//-------------------------------------------------------------------------------------------------

// New{{.UType}}List constructs a new list containing the supplied values, if any.
func New{{.UType}}List(values ...{{.PType}}) {{.UType}}List {
	list := make({{.UType}}List, len(values))
	for i, v := range values {
		list[i] = v
	}
	return list
}

// Build{{.Type}}ListFromChan constructs a new {{.UType}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.Type}}ListFromChan(source <-chan {{.PType}}) {{.UType}}List {
	result := make({{.UType}}List, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list {{.UType}}List) Head() {{.PType}} {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list {{.UType}}List) Last() {{.PType}} {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list {{.UType}}List) Tail() {{.UType}}List {
	return {{.UType}}List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list {{.UType}}List) Init() {{.UType}}List {
	return {{.UType}}List(list[:len(list)-1])
}

// IsEmpty tests whether {{.UType}}List is empty.
func (list {{.UType}}List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.UType}}List is empty.
func (list {{.UType}}List) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list {{.UType}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list {{.UType}}List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list {{.UType}}List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list {{.UType}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.UType}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of {{.UType}}List return true for the passed func.
func (list {{.UType}}List) Exists(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.UType}}List return true for the passed func.
func (list {{.UType}}List) Forall(fn func({{.PType}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.UType}}List and executes the passed func against each element.
func (list {{.UType}}List) Foreach(fn func({{.PType}})) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
func (list {{.UType}}List) Send() <-chan {{.PType}} {
	ch := make(chan {{.PType}})
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of {{.UType}}List with all elements in the reverse order.
func (list {{.UType}}List) Reverse() {{.UType}}List {
	numItems := len(list)
	result := make({{.UType}}List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of {{.UType}}List, using a version of the Fisher-Yates shuffle.
func (list {{.UType}}List) Shuffle() {{.UType}}List {
	numItems := len(list)
	result := make({{.UType}}List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a new {{.UType}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.UType}}List) Take(n int) {{.UType}}List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new {{.UType}}List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.UType}}List) Drop(n int) {{.UType}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new {{.UType}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.UType}}List) TakeLast(n int) {{.UType}}List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new {{.UType}}List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.UType}}List) DropLast(n int) {{.UType}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new {{.UType}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list {{.UType}}List) TakeWhile(p func({{.PType}}) bool) (result {{.UType}}List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.UType}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list {{.UType}}List) DropWhile(p func({{.PType}}) bool) (result {{.UType}}List) {
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

// Filter returns a new {{.UType}}List whose elements return true for func.
func (list {{.UType}}List) Filter(fn func({{.PType}}) bool) {{.UType}}List {
	result := make({{.UType}}List, 0, len(list)/2)
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
func (list {{.UType}}List) Partition(p func({{.PType}}) bool) ({{.UType}}List, {{.UType}}List) {
	matching := make({{.UType}}List, 0, len(list)/2)
	others := make({{.UType}}List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of {{.UType}}List that return true for the passed predicate.
func (list {{.UType}}List) CountBy(predicate func({{.PType}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.UType}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list {{.UType}}List) MinBy(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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

// MaxBy returns an element of {{.UType}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list {{.UType}}List) MaxBy(less func({{.PType}}, {{.PType}}) bool) (result {{.PType}}) {
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

// DistinctBy returns a new {{.UType}}List whose elements are unique, where equality is defined by a passed func.
func (list {{.UType}}List) DistinctBy(equal func({{.PType}}, {{.PType}}) bool) (result {{.UType}}List) {
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
func (list {{.UType}}List) IndexWhere(p func({{.PType}}) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list {{.UType}}List) IndexWhere2(p func({{.PType}}) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list {{.UType}}List) LastIndexWhere(p func({{.PType}}) bool) int {
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
func (list {{.UType}}List) LastIndexWhere2(p func({{.PType}}) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
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
func (list {{.UType}}List) Sum() {{.Type}} {
    sum := {{.Type}}(0)
	for _, v := range list {
		sum = sum + v
	}
	return sum
}

{{end}}
{{if .Comparable}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list {{.UType}}List) Equals(other {{.UType}}List) bool {
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

{{end}}
{{if .Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UType}}List) Min() {{.PType}} {
	return list.MinBy(func(a {{.PType}}, b {{.PType}}) bool {
	    return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UType}}List) Max() (result {{.PType}}) {
	return list.MaxBy(func(a {{.PType}}, b {{.PType}}) bool {
	    return a < b
	})
}

{{end}}
{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list {{.UType}}List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list {{.UType}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list {{.UType}}List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}
{{end}}
