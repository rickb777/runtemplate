// A simple type derived from []big.Int
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=*big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value>

package simple

import (

	"math/rand"
	"sort"
	"math/big"

)

// P2IntList is a slice of type *big.Int. Use it where you would use []*big.Int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type P2IntList []*big.Int

//-------------------------------------------------------------------------------------------------

func newP2IntList(len, cap int) P2IntList {
	return make(P2IntList, len, cap)
}

// NewP2IntList constructs a new list containing the supplied values, if any.
func NewP2IntList(values ...*big.Int) P2IntList {
	result := newP2IntList(len(values), len(values))
	copy(result, values)
	return result
}

// ConvertP2IntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertP2IntList(values ...interface{}) (P2IntList, bool) {
	result := newP2IntList(0, len(values))

	for _, i := range values {
		v, ok := i.(*big.Int)
		if ok {
			result = append(result, v)
		}
	}

	return result, len(result) == len(values)
}

// BuildP2IntListFromChan constructs a new P2IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildP2IntListFromChan(source <-chan *big.Int) P2IntList {
	result := newP2IntList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list P2IntList) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list P2IntList) Clone() P2IntList {
	return NewP2IntList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list P2IntList) Get(i int) *big.Int {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list P2IntList) Head() *big.Int {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list P2IntList) Last() *big.Int {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list P2IntList) Tail() P2IntList {
	return P2IntList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list P2IntList) Init() P2IntList {
	return P2IntList(list[:len(list)-1])
}

// IsEmpty tests whether P2IntList is empty.
func (list P2IntList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether P2IntList is empty.
func (list P2IntList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list P2IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list P2IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list P2IntList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list P2IntList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list P2IntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of P2IntList return true for the passed func.
func (list P2IntList) Exists(fn func(*big.Int) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of P2IntList return true for the passed func.
func (list P2IntList) Forall(fn func(*big.Int) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over P2IntList and executes the passed func against each element.
func (list P2IntList) Foreach(fn func(*big.Int)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list P2IntList) Send() <-chan *big.Int {
	ch := make(chan *big.Int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of P2IntList with all elements in the reverse order.
func (list P2IntList) Reverse() P2IntList {
	numItems := len(list)
	result := newP2IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of P2IntList, using a version of the Fisher-Yates shuffle.
func (list P2IntList) Shuffle() P2IntList {
	result := list.Clone()
	numItems := len(list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of P2IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list P2IntList) Take(n int) P2IntList {
	if n > len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of P2IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list P2IntList) Drop(n int) P2IntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of P2IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list P2IntList) TakeLast(n int) P2IntList {
	l := len(list)
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of P2IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list P2IntList) DropLast(n int) P2IntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new P2IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list P2IntList) TakeWhile(p func(*big.Int) bool) P2IntList {
	result := newP2IntList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new P2IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list P2IntList) DropWhile(p func(*big.Int) bool) P2IntList {
	result := newP2IntList(0, 0)
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
func (list P2IntList) Find(fn func(*big.Int) bool) (*big.Int, bool) {

	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}


	return nil, false

}

// Filter returns a new P2IntList whose elements return true for func.
func (list P2IntList) Filter(fn func(*big.Int) bool) P2IntList {
	result := newP2IntList(0, len(list)/2)

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
func (list P2IntList) Partition(p func(*big.Int) bool) (P2IntList, P2IntList) {
	matching := newP2IntList(0, len(list)/2)
	others := newP2IntList(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of P2IntList that return true for the passed predicate.
func (list P2IntList) CountBy(predicate func(*big.Int) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of P2IntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list P2IntList) MinBy(less func(*big.Int, *big.Int) bool) *big.Int {
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

	return list[m]
}

// MaxBy returns an element of P2IntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list P2IntList) MaxBy(less func(*big.Int, *big.Int) bool) *big.Int {
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

	return list[m]
}

// DistinctBy returns a new P2IntList whose elements are unique, where equality is defined by a passed func.
func (list P2IntList) DistinctBy(equal func(*big.Int, *big.Int) bool) P2IntList {
	result := newP2IntList(0, len(list))
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
func (list P2IntList) IndexWhere(p func(*big.Int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list P2IntList) IndexWhere2(p func(*big.Int) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list P2IntList) LastIndexWhere(p func(*big.Int) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list P2IntList) LastIndexWhere2(p func(*big.Int) bool, before int) int {
	if before < 0 {
		before = len(list)
	}
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------

type sortableP2IntList struct {
	less func(i, j big.Int) bool
	m []*big.Int
}

func (sl sortableP2IntList) Less(i, j int) bool {
	return sl.less(*sl.m[i], *sl.m[j])
}

func (sl sortableP2IntList) Len() int {
	return len(sl.m)
}

func (sl sortableP2IntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list P2IntList) SortBy(less func(i, j big.Int) bool) P2IntList {

	sort.Sort(sortableP2IntList{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list P2IntList) StableSortBy(less func(i, j big.Int) bool) P2IntList {

	sort.Stable(sortableP2IntList{less, list})
	return list
}


