// An encapsulated immutable []big.Int.
// Thread-safe.
//
//
// Generated from immutable/list.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> GobEncode:<no value> Mutable:disabled
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"math/rand"
	"sort"
	"math/big"
)

// X1IntegerList contains a slice of type big.Int. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1IntegerList struct {
	m []big.Int
}

//-------------------------------------------------------------------------------------------------

func newX1IntegerList(length, capacity int) *X1IntegerList {
	return &X1IntegerList {
		m: make([]big.Int, length, capacity),
	}
}

// NewX1IntegerList constructs a new list containing the supplied values, if any.
func NewX1IntegerList(values ...big.Int) *X1IntegerList {
	list := newX1IntegerList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertX1IntegerList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1IntegerList(values ...interface{}) (*X1IntegerList, bool) {
	list := newX1IntegerList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case big.Int:
			list.m = append(list.m, j)
		case *big.Int:
			list.m = append(list.m, *j)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildX1IntegerListFromChan constructs a new X1IntegerList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildX1IntegerListFromChan(source <-chan big.Int) *X1IntegerList {
	list := newX1IntegerList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *X1IntegerList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *X1IntegerList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *X1IntegerList) slice() []big.Int {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *X1IntegerList) ToList() *X1IntegerList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *X1IntegerList) ToSlice() []big.Int {

	s := make([]big.Int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1IntegerList) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *X1IntegerList) Clone() *X1IntegerList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *X1IntegerList) Get(i int) big.Int {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *X1IntegerList) Head() big.Int {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1IntegerList) HeadOption() big.Int {
	if list == nil || len(list.m) == 0 {
		var v big.Int
		return v
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *X1IntegerList) Last() big.Int {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1IntegerList) LastOption() big.Int {
	if list == nil || len(list.m) == 0 {
		var v big.Int
		return v
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *X1IntegerList) Tail() *X1IntegerList {
	result := newX1IntegerList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *X1IntegerList) Init() *X1IntegerList {
	result := newX1IntegerList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1IntegerList is empty.
func (list *X1IntegerList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1IntegerList is empty.
func (list *X1IntegerList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *X1IntegerList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *X1IntegerList) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of X1IntegerList return true for the predicate p.
func (list *X1IntegerList) Exists(p func(big.Int) bool) bool {
	if list == nil {
		return false
	}

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1IntegerList return true for the predicate p.
func (list *X1IntegerList) Forall(p func(big.Int) bool) bool {
	if list == nil {
		return true
	}

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1IntegerList and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *X1IntegerList) Foreach(f func(big.Int)) {
	if list == nil {
		return
	}

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order. A goroutine is created to
// send the elements; this only terminates when all the elements have been consumed. The
// channel will be closed when all the elements have been sent.
func (list *X1IntegerList) Send() <-chan big.Int {
	ch := make(chan big.Int)
	go func() {
		if list != nil {
			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of X1IntegerList with all elements in the reverse order.
func (list *X1IntegerList) Reverse() *X1IntegerList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := newX1IntegerList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of X1IntegerList, using a version of the Fisher-Yates shuffle.
func (list *X1IntegerList) Shuffle() *X1IntegerList {
	if list == nil {
		return nil
	}

	result := NewX1IntegerList(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *X1IntegerList) Append(more ...big.Int) *X1IntegerList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return NewX1IntegerList(more...)
	}

	newList := NewX1IntegerList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *X1IntegerList) doAppend(more ...big.Int) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1IntegerList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *X1IntegerList) Take(n int) *X1IntegerList {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := newX1IntegerList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1IntegerList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntegerList) Drop(n int) *X1IntegerList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := newX1IntegerList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of X1IntegerList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *X1IntegerList) TakeLast(n int) *X1IntegerList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := newX1IntegerList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1IntegerList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntegerList) DropLast(n int) *X1IntegerList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := newX1IntegerList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new X1IntegerList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *X1IntegerList) TakeWhile(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1IntegerList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *X1IntegerList) DropWhile(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(0, 0)
	adding := false

	for _, v := range list.m {
		if adding || !p(v) {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first big.Int that returns true for predicate p.
// False is returned if none match.
func (list *X1IntegerList) Find(p func(big.Int) bool) (big.Int, bool) {
	if list == nil {
		return *(new(big.Int)), false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}


	var empty big.Int
	return empty, false
}

// Filter returns a new X1IntegerList whose elements return true for predicate p.
func (list *X1IntegerList) Filter(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(0, len(list.m)/2)

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
func (list *X1IntegerList) Partition(p func(big.Int) bool) (*X1IntegerList, *X1IntegerList) {
	if list == nil {
		return nil, nil
	}

	matching := newX1IntegerList(0, len(list.m)/2)
	others := newX1IntegerList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new X1IntegerList by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) Map(f func(big.Int) big.Int) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new X1IntegerList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) FlatMap(f func(big.Int) []big.Int) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X1IntegerList that return true for the predicate p.
func (list *X1IntegerList) CountBy(p func(big.Int) bool) (result int) {
	if list == nil {
		return 0
	}

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1IntegerList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntegerList) MinBy(less func(big.Int, big.Int) bool) big.Int {
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

// MaxBy returns an element of X1IntegerList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntegerList) MaxBy(less func(big.Int, big.Int) bool) big.Int {
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

// DistinctBy returns a new X1IntegerList whose elements are unique, where equality is defined by the equal function.
func (list *X1IntegerList) DistinctBy(equal func(big.Int, big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := newX1IntegerList(0, len(list.m))
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

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list *X1IntegerList) IndexWhere(p func(big.Int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *X1IntegerList) IndexWhere2(p func(big.Int) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *X1IntegerList) LastIndexWhere(p func(big.Int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *X1IntegerList) LastIndexWhere2(p func(big.Int) bool, before int) int {

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

type sortableX1IntegerList struct {
	less func(i, j big.Int) bool
	m []big.Int
}

func (sl sortableX1IntegerList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1IntegerList) Len() int {
	return len(sl.m)
}

func (sl sortableX1IntegerList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *X1IntegerList) SortBy(less func(i, j big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := NewX1IntegerList(list.m...)
	sort.Sort(sortableX1IntegerList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *X1IntegerList) StableSortBy(less func(i, j big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := NewX1IntegerList(list.m...)
	sort.Stable(sortableX1IntegerList{less, result.m})
	return result
}
