// Generated from list.tpl with Type=*Apple
// options: Comparable=true Numeric=<no value> Ordered=<no value> Stringer=false

package collectiontest2

import (

	"math/rand"
)

// AppleList is a slice of type *Apple. Use it where you would use []*Apple.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type AppleList []*Apple

//-------------------------------------------------------------------------------------------------

func newAppleList(len, cap int) AppleList {
	return make(AppleList, len, cap)
}

// NewAppleList constructs a new list containing the supplied values, if any.
func NewAppleList(values ...*Apple) AppleList {
	result := newAppleList(len(values), len(values))
	for i, v := range values {
		result[i] = v
	}
	return result
}

// BuildAppleListFromChan constructs a new AppleList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildAppleListFromChan(source <-chan *Apple) AppleList {
	result := newAppleList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list AppleList) Clone() AppleList {
	return NewAppleList(list...)
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list AppleList) Head() *Apple {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list AppleList) Last() *Apple {
	return list[list.Len()-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list AppleList) Tail() AppleList {
	return AppleList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list AppleList) Init() AppleList {
	return AppleList(list[:list.Len()-1])
}

// IsEmpty tests whether AppleList is empty.
func (list AppleList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether AppleList is empty.
func (list AppleList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list AppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list AppleList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list AppleList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list AppleList) Len() int {
	return len(list)
}


//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of AppleList return true for the passed func.
func (list AppleList) Exists(fn func(*Apple) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of AppleList return true for the passed func.
func (list AppleList) Forall(fn func(*Apple) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over AppleList and executes the passed func against each element.
func (list AppleList) Foreach(fn func(*Apple)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list AppleList) Send() <-chan *Apple {
	ch := make(chan *Apple)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of AppleList with all elements in the reverse order.
func (list AppleList) Reverse() AppleList {
	numItems := list.Len()
	result := make(AppleList, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of AppleList, using a version of the Fisher-Yates shuffle.
func (list AppleList) Shuffle() AppleList {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
    	result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of AppleList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list AppleList) Take(n int) AppleList {
	if n > list.Len() {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of AppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list AppleList) Drop(n int) AppleList {
	if n == 0 {
		return list
	}

	l := list.Len()
	if n < l {
		return list[l:]
	}
	return list[n:]
}

// TakeLast returns a slice of AppleList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list AppleList) TakeLast(n int) AppleList {
	l := list.Len()
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of AppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list AppleList) DropLast(n int) AppleList {
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

// TakeWhile returns a new AppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list AppleList) TakeWhile(p func(*Apple) bool) (result AppleList) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new AppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list AppleList) DropWhile(p func(*Apple) bool) (result AppleList) {
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

// Filter returns a new AppleList whose elements return true for func.
func (list AppleList) Filter(fn func(*Apple) bool) AppleList {
	result := make(AppleList, 0, list.Len()/2)

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list AppleList) Partition(p func(*Apple) bool) (AppleList, AppleList) {
	matching := make(AppleList, 0, list.Len()/2)
	others := make(AppleList, 0, list.Len()/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of AppleList that return true for the passed predicate.
func (list AppleList) CountBy(predicate func(*Apple) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of AppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list AppleList) MinBy(less func(*Apple, *Apple) bool) *Apple {
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

// MaxBy returns an element of AppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list AppleList) MaxBy(less func(*Apple, *Apple) bool) *Apple {
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

// DistinctBy returns a new AppleList whose elements are unique, where equality is defined by a passed func.
func (list AppleList) DistinctBy(equal func(*Apple, *Apple) bool) (result AppleList) {
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
func (list AppleList) IndexWhere(p func(*Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list AppleList) IndexWhere2(p func(*Apple) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list AppleList) LastIndexWhere(p func(*Apple) bool) int {
	return list.LastIndexWhere2(p, 0)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list AppleList) LastIndexWhere2(p func(*Apple) bool, before int) int {
	for i := list.Len() - 1; i >= 0; i-- {
		v := list[i]
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
func (list AppleList) Equals(other AppleList) bool {
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




