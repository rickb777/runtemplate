// An encapsulated []Apple
// Not thread-safe.
//
// Generated from list.tpl with Type=Apple
// options: Comparable=true Numeric=<no value> Ordered=<no value> Stringer=false

package fast

import (
"math/rand"
)

// XAppleList contains a slice of type Apple. Use it where you would use []Apple.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type XAppleList struct {
	m []Apple
}


//-------------------------------------------------------------------------------------------------

func newXAppleList(len, cap int) *XAppleList {
	return &XAppleList{
		m: make([]Apple, len, cap),
	}
}

// NewXAppleList constructs a new list containing the supplied values, if any.
func NewXAppleList(values ...Apple) *XAppleList {
	result := newXAppleList(len(values), len(values))
	for i, v := range values {
		result.m[i] = v
	}
	return result
}

// BuildXAppleListFromChan constructs a new XAppleList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildXAppleListFromChan(source <-chan Apple) *XAppleList {
	result := newXAppleList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *XAppleList) Clone() *XAppleList {
	return NewXAppleList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *XAppleList) Head() Apple {
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *XAppleList) Last() Apple {
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *XAppleList) Tail() *XAppleList {
	result := newXAppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *XAppleList) Init() *XAppleList {
	result := newXAppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether XAppleList is empty.
func (list *XAppleList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether XAppleList is empty.
func (list *XAppleList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list *XAppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *XAppleList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *XAppleList) Size() int {
	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *XAppleList) Len() int {
	return len(list.m)
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of XAppleList return true for the passed func.
func (list *XAppleList) Exists(fn func(Apple) bool) bool {
	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of XAppleList return true for the passed func.
func (list *XAppleList) Forall(fn func(Apple) bool) bool {
	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over XAppleList and executes the passed func against each element.
func (list *XAppleList) Foreach(fn func(Apple)) {
	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *XAppleList) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {
		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of XAppleList with all elements in the reverse order.
func (list *XAppleList) Reverse() *XAppleList {
	numItems := list.Len()
	result := newXAppleList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of XAppleList, using a version of the Fisher-Yates shuffle.
func (list *XAppleList) Shuffle() *XAppleList {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of XAppleList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *XAppleList) Take(n int) *XAppleList {
	if n > list.Len() {
		return list
	}
	result := newXAppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of XAppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *XAppleList) Drop(n int) *XAppleList {
	if n == 0 {
		return list
	}

	result := newXAppleList(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of XAppleList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *XAppleList) TakeLast(n int) *XAppleList {
	l := list.Len()
	if n > l {
		return list
	}
	result := newXAppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of XAppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *XAppleList) DropLast(n int) *XAppleList {
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

// TakeWhile returns a new XAppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *XAppleList) TakeWhile(p func(Apple) bool) *XAppleList {
	result := newXAppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new XAppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *XAppleList) DropWhile(p func(Apple) bool) *XAppleList {
	result := newXAppleList(0, 0)
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

// Filter returns a new XAppleList whose elements return true for func.
func (list *XAppleList) Filter(fn func(Apple) bool) *XAppleList {
	result := newXAppleList(0, list.Len()/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *XAppleList) Partition(p func(Apple) bool) (*XAppleList, *XAppleList) {
	matching := newXAppleList(0, list.Len()/2)
	others := newXAppleList(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of XAppleList that return true for the passed predicate.
func (list *XAppleList) CountBy(predicate func(Apple) bool) (result int) {
	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XAppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *XAppleList) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of XAppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *XAppleList) MaxBy(less func(Apple, Apple) bool) Apple {
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

// DistinctBy returns a new XAppleList whose elements are unique, where equality is defined by a passed func.
func (list *XAppleList) DistinctBy(equal func(Apple, Apple) bool) *XAppleList {
	result := newXAppleList(0, list.Len())
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
func (list *XAppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *XAppleList) IndexWhere2(p func(Apple) bool, from int) int {
	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *XAppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, 0)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *XAppleList) LastIndexWhere2(p func(Apple) bool, before int) int {
	for i := list.Len() - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}


//-------------------------------------------------------------------------------------------------
// These methods are included when Apple is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *XAppleList) Equals(other *XAppleList) bool {
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


