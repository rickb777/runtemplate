// An encapsulated []big.Int.
// Thread-safe.
//
// Generated from fast/list.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:always

package fast

import (

	"math/rand"
	"sort"
	"math/big"

)

// X2IntList contains a slice of type big.Int. Use it where you would use []big.Int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X2IntList struct {
	m []big.Int
}


//-------------------------------------------------------------------------------------------------

func newX2IntList(len, cap int) *X2IntList {
	return &X2IntList {
		m: make([]big.Int, len, cap),
	}
}

// NewX2IntList constructs a new list containing the supplied values, if any.
func NewX2IntList(values ...big.Int) *X2IntList {
	result := newX2IntList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertX2IntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX2IntList(values ...interface{}) (*X2IntList, bool) {
	result := newX2IntList(0, len(values))

	for _, i := range values {
		v, ok := i.(big.Int)
		if ok {
			result.m = append(result.m, v)
		}
	}

	return result, len(result.m) == len(values)
}

// BuildX2IntListFromChan constructs a new X2IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2IntListFromChan(source <-chan big.Int) *X2IntList {
	result := newX2IntList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *X2IntList) ToSlice() []big.Int {

	s := make([]big.Int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X2IntList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *X2IntList) Clone() *X2IntList {
	return NewX2IntList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *X2IntList) Get(i int) big.Int {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *X2IntList) Head() big.Int {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *X2IntList) Last() big.Int {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *X2IntList) Tail() *X2IntList {

	result := newX2IntList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *X2IntList) Init() *X2IntList {

	result := newX2IntList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X2IntList is empty.
func (list *X2IntList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X2IntList is empty.
func (list *X2IntList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *X2IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *X2IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list.
func (list *X2IntList) Size() int {

	return len(list.m)
}

// Swap exchanges two elements.
func (list *X2IntList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of X2IntList return true for the predicate p.
func (list *X2IntList) Exists(p func(big.Int) bool) bool {

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X2IntList return true for the predicate p.
func (list *X2IntList) Forall(p func(big.Int) bool) bool {

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X2IntList and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (list *X2IntList) Foreach(fn func(big.Int)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *X2IntList) Send() <-chan big.Int {
	ch := make(chan big.Int)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X2IntList with all elements in the reverse order.
func (list *X2IntList) Reverse() *X2IntList {

	numItems := len(list.m)
	result := newX2IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X2IntList, using a version of the Fisher-Yates shuffle.
func (list *X2IntList) Shuffle() *X2IntList {
	result := list.Clone()
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Add adds items to the current list. This is a synonym for Append.
func (list *X2IntList) Add(more ...big.Int) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *X2IntList) Append(more ...big.Int) *X2IntList {

	list.doAppend(more...)
	return list
}

func (list *X2IntList) doAppend(more ...big.Int) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X2IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X2IntList) Take(n int) *X2IntList {

	if n > len(list.m) {
		return list
	}
	result := newX2IntList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X2IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X2IntList) Drop(n int) *X2IntList {
	if n == 0 {
		return list
	}


	result := newX2IntList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of X2IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X2IntList) TakeLast(n int) *X2IntList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newX2IntList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X2IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X2IntList) DropLast(n int) *X2IntList {
	if n == 0 {
		return list
	}


	l := len(list.m)
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new X2IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *X2IntList) TakeWhile(p func(big.Int) bool) *X2IntList {

	result := newX2IntList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X2IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *X2IntList) DropWhile(p func(big.Int) bool) *X2IntList {

	result := newX2IntList(0, 0)
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

// InsertAt modifies a X2IntList by inserting elements at a given index.
// This is a generalised version of Append.
//
// The modified list is returned.
// Panics if the index is out of range.
func (list *X2IntList) InsertAt(index int, more ...big.Int) *X2IntList {

    if len(more) == 0 {
        return list
    }

	newlist := make([]big.Int, 0, len(list.m) + len(more))

    if index != 0 {
        newlist = append(newlist, list.m[:index]...)
    }

    newlist = append(newlist, more...)

    if index != len(list.m) {
        newlist = append(newlist, list.m[index:]...)
    }

    list.m = newlist
	return list
}

// DeleteAt modifies a X2IntList by deleting n elements from a given index.
//
// The modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *X2IntList) DeleteAt(index, n int) *X2IntList {

    if n == 0 {
        return list
    }

	newlist := make([]big.Int, 0, len(list.m) - n)

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

// KeepWhere modifies a X2IntList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The modified list is returned.
func (list *X2IntList) KeepWhere(p func(big.Int) bool) *X2IntList {

	result := make([]big.Int, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

    list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Find returns the first big.Int that returns true for predicate p.
// False is returned if none match.
func (list X2IntList) Find(p func(big.Int) bool) (big.Int, bool) {

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}


	var empty big.Int
	return empty, false

}

// Filter returns a new X2IntList whose elements return true for predicate p.
//
// The original list is not modified. See also KeepWhere (which does modify the original list).
func (list *X2IntList) Filter(p func(big.Int) bool) *X2IntList {

	result := newX2IntList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new big.IntLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *X2IntList) Partition(p func(big.Int) bool) (*X2IntList, *X2IntList) {

	matching := newX2IntList(0, len(list.m)/2)
	others := newX2IntList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new X2IntList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X2IntList) Map(fn func(big.Int) big.Int) *X2IntList {
	result := newX2IntList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new X2IntList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X2IntList) FlatMap(fn func(big.Int) []big.Int) *X2IntList {
	result := newX2IntList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of X2IntList that return true for the passed predicate.
func (list *X2IntList) CountBy(predicate func(big.Int) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2IntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X2IntList) MinBy(less func(big.Int, big.Int) bool) big.Int {

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

// MaxBy returns an element of X2IntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X2IntList) MaxBy(less func(big.Int, big.Int) bool) big.Int {

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

// DistinctBy returns a new X2IntList whose elements are unique, where equality is defined by a passed func.
func (list *X2IntList) DistinctBy(equal func(big.Int, big.Int) bool) *X2IntList {

	result := newX2IntList(0, len(list.m))
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
func (list *X2IntList) IndexWhere(p func(big.Int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *X2IntList) IndexWhere2(p func(big.Int) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *X2IntList) LastIndexWhere(p func(big.Int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *X2IntList) LastIndexWhere2(p func(big.Int) bool, before int) int {

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

//-------------------------------------------------------------------------------------------------

type sortableX2IntList struct {
	less func(i, j big.Int) bool
	m []big.Int
}

func (sl sortableX2IntList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX2IntList) Len() int {
	return len(sl.m)
}

func (sl sortableX2IntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *X2IntList) SortBy(less func(i, j big.Int) bool) *X2IntList {

	sort.Sort(sortableX2IntList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *X2IntList) StableSortBy(less func(i, j big.Int) bool) *X2IntList {

	sort.Stable(sortableX2IntList{less, list.m})
	return list
}


