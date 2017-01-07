// Generated from list.tpl with Type=string
// options: Comparable=true Numeric=false Ordered=false Stringer=true

package threadsafe

import (

	"bytes"
	"fmt"

	"sync"
	"math/rand"
)

// XStringList contains a slice of type string. Use it where you would use []string.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type XStringList struct {
	s *sync.RWMutex
	m []string
}


//-------------------------------------------------------------------------------------------------

func newXStringList(len, cap int) XStringList {
	return XStringList{
		s: &sync.RWMutex{},
		m: make([]string, len, cap),
	}
}

// NewXStringList constructs a new list containing the supplied values, if any.
func NewXStringList(values ...string) XStringList {
	result := newXStringList(len(values), len(values))
	for i, v := range values {
		result.m[i] = v
	}
	return result
}

// BuildXStringListFromChan constructs a new XStringList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildXStringListFromChan(source <-chan string) XStringList {
	result := newXStringList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list XStringList) Clone() XStringList {
	return NewXStringList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list XStringList) Head() string {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list XStringList) Last() string {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list XStringList) Tail() XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list XStringList) Init() XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether XStringList is empty.
func (list XStringList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether XStringList is empty.
func (list XStringList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list XStringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list XStringList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list XStringList) Size() int {
	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list XStringList) Len() int {
	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}


//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of XStringList return true for the passed func.
func (list XStringList) Exists(fn func(string) bool) bool {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of XStringList return true for the passed func.
func (list XStringList) Forall(fn func(string) bool) bool {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over XStringList and executes the passed func against each element.
func (list XStringList) Foreach(fn func(string)) {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list XStringList) Send() <-chan string {
	ch := make(chan string)
	go func() {
		list.s.RLock()
		defer list.s.RUnlock()

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of XStringList with all elements in the reverse order.
func (list XStringList) Reverse() XStringList {
	list.s.Lock()
	defer list.s.Unlock()

	numItems := list.Len()
	result := newXStringList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of XStringList, using a version of the Fisher-Yates shuffle.
func (list XStringList) Shuffle() XStringList {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of XStringList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list XStringList) Take(n int) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	if n > list.Len() {
		return list
	}
	result := newXStringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of XStringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list XStringList) Drop(n int) XStringList {
	if n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of XStringList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list XStringList) TakeLast(n int) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	l := list.Len()
	if n > l {
		return list
	}
	result := newXStringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of XStringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list XStringList) DropLast(n int) XStringList {
	if n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := list.Len()
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new XStringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list XStringList) TakeWhile(p func(string) bool) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new XStringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list XStringList) DropWhile(p func(string) bool) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, 0)
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

// Filter returns a new XStringList whose elements return true for func.
func (list XStringList) Filter(fn func(string) bool) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, list.Len()/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list XStringList) Partition(p func(string) bool) (XStringList, XStringList) {
	list.s.RLock()
	defer list.s.RUnlock()

	matching := newXStringList(0, list.Len()/2)
	others := newXStringList(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of XStringList that return true for the passed predicate.
func (list XStringList) CountBy(predicate func(string) bool) (result int) {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of XStringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list XStringList) MinBy(less func(string, string) bool) string {
	list.s.RLock()
	defer list.s.RUnlock()

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

// MaxBy returns an element of XStringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list XStringList) MaxBy(less func(string, string) bool) string {
	list.s.RLock()
	defer list.s.RUnlock()

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

// DistinctBy returns a new XStringList whose elements are unique, where equality is defined by a passed func.
func (list XStringList) DistinctBy(equal func(string, string) bool) XStringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newXStringList(0, list.Len())
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
func (list XStringList) IndexWhere(p func(string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list XStringList) IndexWhere2(p func(string) bool, from int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list XStringList) LastIndexWhere(p func(string) bool) int {
	return list.LastIndexWhere2(p, 0)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list XStringList) LastIndexWhere2(p func(string) bool, before int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i := list.Len() - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}



//-------------------------------------------------------------------------------------------------
// These methods are included when string is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list XStringList) Equals(other XStringList) bool {
	list.s.RLock()
	other.s.RLock()
	defer list.s.RUnlock()
	defer other.s.RUnlock()

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




//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list XStringList) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list XStringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list XStringList) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)

	list.s.RLock()
	defer list.s.RUnlock()

	l := list.Len()
	if l > 0 {
		v := list.m[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list.m[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}

