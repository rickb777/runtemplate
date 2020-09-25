// An encapsulated immutable []Apple.
// Thread-safe.
//
//
// Generated from immutable/list.tpl with Type=Apple
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:false GobEncode:<no value> Mutable:disabled
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"math/rand"
	"sort"
)

// X1AppleList contains a slice of type Apple. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1AppleList struct {
	m []Apple
}

//-------------------------------------------------------------------------------------------------

func newX1AppleList(length, capacity int) *X1AppleList {
	return &X1AppleList{
		m: make([]Apple, length, capacity),
	}
}

// NewX1AppleList constructs a new list containing the supplied values, if any.
func NewX1AppleList(values ...Apple) *X1AppleList {
	list := newX1AppleList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertX1AppleList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1AppleList(values ...interface{}) (*X1AppleList, bool) {
	list := newX1AppleList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case Apple:
			k := Apple(j)
			list.m = append(list.m, k)
		case *Apple:
			k := Apple(*j)
			list.m = append(list.m, k)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildX1AppleListFromChan constructs a new X1AppleList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildX1AppleListFromChan(source <-chan Apple) *X1AppleList {
	list := newX1AppleList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *X1AppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *X1AppleList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *X1AppleList) slice() []Apple {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *X1AppleList) ToList() *X1AppleList {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *X1AppleList) ToSet() *X1AppleSet {
	if list == nil {
		return nil
	}

	return NewX1AppleSet(list.m...)
}

// ToSlice returns the elements of the current list as a slice.
func (list *X1AppleList) ToSlice() []Apple {

	s := make([]Apple, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1AppleList) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *X1AppleList) Clone() *X1AppleList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *X1AppleList) Get(i int) Apple {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *X1AppleList) Head() Apple {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1AppleList) HeadOption() Apple {
	if list == nil || len(list.m) == 0 {
		var v Apple
		return v
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *X1AppleList) Last() Apple {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1AppleList) LastOption() Apple {
	if list == nil || len(list.m) == 0 {
		var v Apple
		return v
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *X1AppleList) Tail() *X1AppleList {
	result := newX1AppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *X1AppleList) Init() *X1AppleList {
	result := newX1AppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1AppleList is empty.
func (list *X1AppleList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1AppleList is empty.
func (list *X1AppleList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *X1AppleList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *X1AppleList) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *X1AppleList) Contains(v Apple) bool {
	return list.Exists(func(x Apple) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *X1AppleList) ContainsAll(i ...Apple) bool {
	if list == nil {
		return len(i) == 0
	}

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1AppleList return true for the predicate p.
func (list *X1AppleList) Exists(p func(Apple) bool) bool {
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

// Forall verifies that all elements of X1AppleList return true for the predicate p.
func (list *X1AppleList) Forall(p func(Apple) bool) bool {
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

// Foreach iterates over X1AppleList and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *X1AppleList) Foreach(f func(Apple)) {
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
func (list *X1AppleList) Send() <-chan Apple {
	ch := make(chan Apple)
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

// Reverse returns a copy of X1AppleList with all elements in the reverse order.
func (list *X1AppleList) Reverse() *X1AppleList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := newX1AppleList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of X1AppleList, using a version of the Fisher-Yates shuffle.
func (list *X1AppleList) Shuffle() *X1AppleList {
	if list == nil {
		return nil
	}

	result := NewX1AppleList(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *X1AppleList) Append(more ...Apple) *X1AppleList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return NewX1AppleList(more...)
	}

	newList := NewX1AppleList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *X1AppleList) doAppend(more ...Apple) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1AppleList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *X1AppleList) Take(n int) *X1AppleList {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := newX1AppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1AppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1AppleList) Drop(n int) *X1AppleList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := newX1AppleList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of X1AppleList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *X1AppleList) TakeLast(n int) *X1AppleList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := newX1AppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1AppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1AppleList) DropLast(n int) *X1AppleList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := newX1AppleList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new X1AppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *X1AppleList) TakeWhile(p func(Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1AppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *X1AppleList) DropWhile(p func(Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(0, 0)
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

// Find returns the first Apple that returns true for predicate p.
// False is returned if none match.
func (list *X1AppleList) Find(p func(Apple) bool) (Apple, bool) {
	if list == nil {
		return *(new(Apple)), false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false
}

// Filter returns a new X1AppleList whose elements return true for predicate p.
func (list *X1AppleList) Filter(p func(Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *X1AppleList) Partition(p func(Apple) bool) (*X1AppleList, *X1AppleList) {
	if list == nil {
		return nil, nil
	}

	matching := newX1AppleList(0, len(list.m)/2)
	others := newX1AppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new X1AppleList by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1AppleList) Map(f func(Apple) Apple) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new X1AppleList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1AppleList) FlatMap(f func(Apple) []Apple) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X1AppleList that return true for the predicate p.
func (list *X1AppleList) CountBy(p func(Apple) bool) (result int) {
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

// MinBy returns an element of X1AppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1AppleList) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of X1AppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1AppleList) MaxBy(less func(Apple, Apple) bool) Apple {
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

// DistinctBy returns a new X1AppleList whose elements are unique, where equality is defined by the equal function.
func (list *X1AppleList) DistinctBy(equal func(Apple, Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := newX1AppleList(0, len(list.m))
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
func (list *X1AppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *X1AppleList) IndexWhere2(p func(Apple) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *X1AppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *X1AppleList) LastIndexWhere2(p func(Apple) bool, before int) int {

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
// These methods are included when Apple is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *X1AppleList) Equals(other *X1AppleList) bool {
	if list == nil {
		return other == nil || len(other.m) == 0
	}

	if other == nil {
		return len(list.m) == 0
	}

	if len(list.m) != len(other.m) {
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

type sortableX1AppleList struct {
	less func(i, j Apple) bool
	m    []Apple
}

func (sl sortableX1AppleList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1AppleList) Len() int {
	return len(sl.m)
}

func (sl sortableX1AppleList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *X1AppleList) SortBy(less func(i, j Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := NewX1AppleList(list.m...)
	sort.Sort(sortableX1AppleList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *X1AppleList) StableSortBy(less func(i, j Apple) bool) *X1AppleList {
	if list == nil {
		return nil
	}

	result := NewX1AppleList(list.m...)
	sort.Stable(sortableX1AppleList{less, result.m})
	return result
}
