// A simple type derived from []big.Int
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value>
// GobEncode:<no value> Mutable:always ToList:always ToSet:false
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"math/rand"
	"sort"
	"math/big"
)

// X2BigIntList is a slice of type big.Int. Use it where you would use []big.Int.
// To add items to the list, simply use the normal built-in append function.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X2BigIntList []big.Int

//-------------------------------------------------------------------------------------------------

// MakeX2BigIntList makes an empty list with both length and capacity initialised.
func MakeX2BigIntList(length, capacity int) X2BigIntList {
	return make(X2BigIntList, length, capacity)
}

// NewX2BigIntList constructs a new list containing the supplied values, if any.
func NewX2BigIntList(values ...big.Int) X2BigIntList {
	list := MakeX2BigIntList(len(values), len(values))
	copy(list, values)
	return list
}

// ConvertX2BigIntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX2BigIntList(values ...interface{}) (X2BigIntList, bool) {
	list := MakeX2BigIntList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
        case big.Int:
			list = append(list, j)
        case *big.Int:
			list = append(list, *j)
		}
	}

	return list, len(list) == len(values)
}

// BuildX2BigIntListFromChan constructs a new X2BigIntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2BigIntListFromChan(source <-chan big.Int) X2BigIntList {
	list := MakeX2BigIntList(0, 0)
	for v := range source {
		list = append(list, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list X2BigIntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list X2BigIntList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list X2BigIntList) slice() []big.Int {
	return list
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list X2BigIntList) ToList() X2BigIntList {
	return list
}

// ToSlice returns the elements of the list as a slice, which is an identity operation in this case,
// because the simple list is merely a dressed-up slice.
func (list X2BigIntList) ToSlice() []big.Int {
	return list
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list X2BigIntList) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(list))
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list X2BigIntList) Clone() X2BigIntList {
	return NewX2BigIntList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list X2BigIntList) Get(i int) big.Int {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list X2BigIntList) Head() big.Int {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list X2BigIntList) HeadOption() big.Int {
	if list.IsEmpty() {
		return *(new(big.Int))
	}
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list X2BigIntList) Last() big.Int {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list X2BigIntList) LastOption() big.Int {
	if list.IsEmpty() {
		return *(new(big.Int))
	}
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list X2BigIntList) Tail() X2BigIntList {
	return list[1:]
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list X2BigIntList) Init() X2BigIntList {
	return list[:len(list)-1]
}

// IsEmpty tests whether X2BigIntList is empty.
func (list X2BigIntList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X2BigIntList is empty.
func (list X2BigIntList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list X2BigIntList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list X2BigIntList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list X2BigIntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of X2BigIntList return true for the predicate p.
func (list X2BigIntList) Exists(p func(big.Int) bool) bool {
	for _, v := range list {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X2BigIntList return true for the predicate p.
func (list X2BigIntList) Forall(p func(big.Int) bool) bool {
	for _, v := range list {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X2BigIntList and executes function f against each element.
func (list X2BigIntList) Foreach(f func(big.Int)) {
	for _, v := range list {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list X2BigIntList) Send() <-chan big.Int {
	ch := make(chan big.Int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of X2BigIntList with all elements in the reverse order.
//
// The original list is not modified.
func (list X2BigIntList) Reverse() X2BigIntList {
	n := len(list)
	result := MakeX2BigIntList(n, n)
	last := n - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse alters a X2BigIntList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list X2BigIntList) DoReverse() X2BigIntList {
	mid := (len(list) + 1) / 2
	last := len(list) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list[i], list[r] = list[r], list[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of X2BigIntList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list X2BigIntList) Shuffle() X2BigIntList {
	if list == nil {
		return nil
	}

	return list.Clone().DoShuffle()
}

// DoShuffle returns a shuffled X2BigIntList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list X2BigIntList) DoShuffle() X2BigIntList {
	if list == nil {
		return nil
	}

	n := len(list)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list[i], list[r] = list[r], list[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X2BigIntList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list X2BigIntList) Take(n int) X2BigIntList {
	if n >= len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of X2BigIntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list X2BigIntList) Drop(n int) X2BigIntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of X2BigIntList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list X2BigIntList) TakeLast(n int) X2BigIntList {
	l := len(list)
	if n >= l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of X2BigIntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list X2BigIntList) DropLast(n int) X2BigIntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	}
    return list[0:l-n]
}

// TakeWhile returns a new X2BigIntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list X2BigIntList) TakeWhile(p func(big.Int) bool) X2BigIntList {
	result := MakeX2BigIntList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X2BigIntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list X2BigIntList) DropWhile(p func(big.Int) bool) X2BigIntList {
	result := MakeX2BigIntList(0, 0)
	adding := false

	for _, v := range list {
		if adding || !p(v) {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first big.Int that returns true for predicate p.
// False is returned if none match.
func (list X2BigIntList) Find(p func(big.Int) bool) (big.Int, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}

	var empty big.Int
	return empty, false
}

// Filter returns a new X2BigIntList whose elements return true for predicate p.
//
// The original list is not modified.
func (list X2BigIntList) Filter(p func(big.Int) bool) X2BigIntList {
	result := MakeX2BigIntList(0, len(list))

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new BigIntLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list X2BigIntList) Partition(p func(big.Int) bool) (X2BigIntList, X2BigIntList) {
	matching := MakeX2BigIntList(0, len(list))
	others := MakeX2BigIntList(0, len(list))

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new X2BigIntList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list X2BigIntList) Map(f func(big.Int) big.Int) X2BigIntList {
	result := MakeX2BigIntList(0, len(list))

	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new X2BigIntList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list X2BigIntList) FlatMap(f func(big.Int) []big.Int) X2BigIntList {
	result := MakeX2BigIntList(0, len(list))

	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X2BigIntList that return true for the predicate p.
func (list X2BigIntList) CountBy(p func(big.Int) bool) (result int) {
	for _, v := range list {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2BigIntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list X2BigIntList) MinBy(less func(big.Int, big.Int) bool) big.Int {
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

// MaxBy returns an element of X2BigIntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list X2BigIntList) MaxBy(less func(big.Int, big.Int) bool) big.Int {
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

// DistinctBy returns a new X2BigIntList whose elements are unique, where equality is defined by the equal function.
func (list X2BigIntList) DistinctBy(equal func(big.Int, big.Int) bool) X2BigIntList {
	result := MakeX2BigIntList(0, len(list))
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

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list X2BigIntList) IndexWhere(p func(big.Int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list X2BigIntList) IndexWhere2(p func(big.Int) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list X2BigIntList) LastIndexWhere(p func(big.Int) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list X2BigIntList) LastIndexWhere2(p func(big.Int) bool, before int) int {
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

type sortableX2BigIntList struct {
	less func(i, j big.Int) bool
	m []big.Int
}

func (sl sortableX2BigIntList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX2BigIntList) Len() int {
	return len(sl.m)
}

func (sl sortableX2BigIntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list X2BigIntList) SortBy(less func(i, j big.Int) bool) X2BigIntList {

	sort.Sort(sortableX2BigIntList{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list X2BigIntList) StableSortBy(less func(i, j big.Int) bool) X2BigIntList {

	sort.Stable(sortableX2BigIntList{less, list})
	return list
}
