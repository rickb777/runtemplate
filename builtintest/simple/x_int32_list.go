// A simple type derived from []int32
// Not thread-safe.
//
// Generated from list.tpl with Type=int32
// options: Comparable=true Numeric=true Ordered=true Stringer=true Mutable=always

package simple

import (
	"math/rand"

	"bytes"
	"fmt"
)

// XInt32List is a slice of type int32. Use it where you would use []int32.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type XInt32List []int32

//-------------------------------------------------------------------------------------------------

func newXInt32List(len, cap int) XInt32List {
	return make(XInt32List, len, cap)
}

// NewXInt32List constructs a new list containing the supplied values, if any.
func NewXInt32List(values ...int32) XInt32List {
	result := newXInt32List(len(values), len(values))
	for i, v := range values {
		result[i] = v
	}
	return result
}

// BuildXInt32ListFromChan constructs a new XInt32List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildXInt32ListFromChan(source <-chan int32) XInt32List {
	result := newXInt32List(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list XInt32List) Clone() XInt32List {
	return NewXInt32List(list...)
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list XInt32List) Head() int32 {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list XInt32List) Last() int32 {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list XInt32List) Tail() XInt32List {
	return XInt32List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list XInt32List) Init() XInt32List {
	return XInt32List(list[:list.Len()-1])
}

// IsEmpty tests whether XInt32List is empty.
func (list XInt32List) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether XInt32List is empty.
func (list XInt32List) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list XInt32List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list XInt32List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list XInt32List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list XInt32List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list XInt32List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of XInt32List return true for the passed func.
func (list XInt32List) Exists(fn func(int32) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of XInt32List return true for the passed func.
func (list XInt32List) Forall(fn func(int32) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over XInt32List and executes the passed func against each element.
func (list XInt32List) Foreach(fn func(int32)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list XInt32List) Send() <-chan int32 {
	ch := make(chan int32)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of XInt32List with all elements in the reverse order.
func (list XInt32List) Reverse() XInt32List {
	numItems := list.Len()
	result := newXInt32List(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of XInt32List, using a version of the Fisher-Yates shuffle.
func (list XInt32List) Shuffle() XInt32List {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of XInt32List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list XInt32List) Take(n int) XInt32List {
	if n > list.Len() {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of XInt32List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list XInt32List) Drop(n int) XInt32List {
	if n == 0 {
		return list
	}

	result := list
	l := list.Len()
	if n < l {
		result = list[n:]
	}
	return result
}

// TakeLast returns a slice of XInt32List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list XInt32List) TakeLast(n int) XInt32List {
	l := list.Len()
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of XInt32List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list XInt32List) DropLast(n int) XInt32List {
	if n == 0 {
		return list
	}

	l := list.Len()
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new XInt32List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list XInt32List) TakeWhile(p func(int32) bool) XInt32List {
	result := newXInt32List(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new XInt32List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list XInt32List) DropWhile(p func(int32) bool) XInt32List {
	result := newXInt32List(0, 0)
	adding := false

	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new XInt32List whose elements return true for func.
func (list XInt32List) Filter(fn func(int32) bool) XInt32List {
	result := newXInt32List(0, list.Len()/2)

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new int32Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list XInt32List) Partition(p func(int32) bool) (XInt32List, XInt32List) {
	matching := newXInt32List(0, list.Len()/2)
	others := newXInt32List(0, list.Len()/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of XInt32List that return true for the passed predicate.
func (list XInt32List) CountBy(predicate func(int32) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XInt32List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list XInt32List) MinBy(less func(int32, int32) bool) int32 {
	l := list.Len()
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

// MaxBy returns an element of XInt32List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list XInt32List) MaxBy(less func(int32, int32) bool) int32 {
	l := list.Len()
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

// DistinctBy returns a new XInt32List whose elements are unique, where equality is defined by a passed func.
func (list XInt32List) DistinctBy(equal func(int32, int32) bool) XInt32List {
	result := newXInt32List(0, list.Len())
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
func (list XInt32List) IndexWhere(p func(int32) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list XInt32List) IndexWhere2(p func(int32) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list XInt32List) LastIndexWhere(p func(int32) bool) int {
	return list.LastIndexWhere2(p, 0)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list XInt32List) LastIndexWhere2(p func(int32) bool, before int) int {
	for i := list.Len() - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is numeric.

// Sum returns the sum of all the elements in the list.
func (list XInt32List) Sum() int32 {
	sum := int32(0)
	for _, v := range list {
		sum = sum + v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list XInt32List) Equals(other XInt32List) bool {
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


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list XInt32List) Min() int32 {
	m := list.MinBy(func(a int32, b int32) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list XInt32List) Max() (result int32) {
	m := list.MaxBy(func(a int32, b int32) bool {
		return a < b
	})
	return m
}

// Less returns true if the element at index i is less than the element at index j. This implements
// one of the methods needed by sort.Interface.
// Panics if i or j is out of range.
func (list XInt32List) Less(i, j int) bool {
	return list[i] < list[j]
}


//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list XInt32List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list XInt32List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list XInt32List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := list.Len()
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

