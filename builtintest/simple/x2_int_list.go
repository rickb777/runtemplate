// A simple type derived from []big.Int
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value>

package simple

import (
	"math/rand"
    "math/big"

)

// X2IntList is a slice of type big.Int. Use it where you would use []big.Int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X2IntList []big.Int

//-------------------------------------------------------------------------------------------------

func newX2IntList(len, cap int) X2IntList {
	return make(X2IntList, len, cap)
}

// NewX2IntList constructs a new list containing the supplied values, if any.
func NewX2IntList(values ...big.Int) X2IntList {
	result := newX2IntList(len(values), len(values))
	for i, v := range values {
		result[i] = v
	}
	return result
}

// BuildX2IntListFromChan constructs a new X2IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2IntListFromChan(source <-chan big.Int) X2IntList {
	result := newX2IntList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list X2IntList) Clone() X2IntList {
	return NewX2IntList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list X2IntList) Get(i int) big.Int {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list X2IntList) Head() big.Int {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list X2IntList) Last() big.Int {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list X2IntList) Tail() X2IntList {
	return X2IntList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list X2IntList) Init() X2IntList {
	return X2IntList(list[:list.Len()-1])
}

// IsEmpty tests whether X2IntList is empty.
func (list X2IntList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether X2IntList is empty.
func (list X2IntList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list X2IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list X2IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list X2IntList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list X2IntList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list X2IntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of X2IntList return true for the passed func.
func (list X2IntList) Exists(fn func(big.Int) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X2IntList return true for the passed func.
func (list X2IntList) Forall(fn func(big.Int) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X2IntList and executes the passed func against each element.
func (list X2IntList) Foreach(fn func(big.Int)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list X2IntList) Send() <-chan big.Int {
	ch := make(chan big.Int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X2IntList with all elements in the reverse order.
func (list X2IntList) Reverse() X2IntList {
	numItems := list.Len()
	result := newX2IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X2IntList, using a version of the Fisher-Yates shuffle.
func (list X2IntList) Shuffle() X2IntList {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X2IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list X2IntList) Take(n int) X2IntList {
	if n > list.Len() {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of X2IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list X2IntList) Drop(n int) X2IntList {
	if n == 0 {
		return list
	}

	l := list.Len()
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of X2IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list X2IntList) TakeLast(n int) X2IntList {
	l := list.Len()
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of X2IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list X2IntList) DropLast(n int) X2IntList {
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

// TakeWhile returns a new X2IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list X2IntList) TakeWhile(p func(big.Int) bool) X2IntList {
	result := newX2IntList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X2IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list X2IntList) DropWhile(p func(big.Int) bool) X2IntList {
	result := newX2IntList(0, 0)
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

// Find returns the first big.Int that returns true for some function.
// False is returned if none match.
func (list X2IntList) Find(fn func(big.Int) bool) (big.Int, bool) {

	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}


    var empty big.Int
	return empty, false

}

// Filter returns a new X2IntList whose elements return true for func.
func (list X2IntList) Filter(fn func(big.Int) bool) X2IntList {
	result := newX2IntList(0, list.Len()/2)

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new big.IntLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list X2IntList) Partition(p func(big.Int) bool) (X2IntList, X2IntList) {
	matching := newX2IntList(0, list.Len()/2)
	others := newX2IntList(0, list.Len()/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of X2IntList that return true for the passed predicate.
func (list X2IntList) CountBy(predicate func(big.Int) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2IntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list X2IntList) MinBy(less func(big.Int, big.Int) bool) big.Int {
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

// MaxBy returns an element of X2IntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list X2IntList) MaxBy(less func(big.Int, big.Int) bool) big.Int {
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

// DistinctBy returns a new X2IntList whose elements are unique, where equality is defined by a passed func.
func (list X2IntList) DistinctBy(equal func(big.Int, big.Int) bool) X2IntList {
	result := newX2IntList(0, list.Len())
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
func (list X2IntList) IndexWhere(p func(big.Int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list X2IntList) IndexWhere2(p func(big.Int) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list X2IntList) LastIndexWhere(p func(big.Int) bool) int {
	return list.LastIndexWhere2(p, list.Len())
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list X2IntList) LastIndexWhere2(p func(big.Int) bool, before int) int {
	if before < 0 {
		before = list.Len()
	}
	for i := list.Len() - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}


