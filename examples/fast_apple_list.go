// An encapsulated []Apple.
// Not thread-safe.
//
// Generated from fast/list.tpl with Type=Apple
// options: Comparable:true Numeric:<no value> Integer:<no value> Ordered:<no value>
//          StringLike:<no value> StringParser:<no value> Stringer:false
// GobEncode:true Mutable:always ToList:always ToSet:<no value> MapTo:<no value>
// by runtemplate v3.10.2
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"sort"
)

// FastAppleList contains a slice of type Apple.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FastAppleList struct {
	m []Apple
}

//-------------------------------------------------------------------------------------------------

// MakeFastAppleList makes an empty list with both length and capacity initialised.
func MakeFastAppleList(length, capacity int) *FastAppleList {
	return &FastAppleList{
		m: make([]Apple, length, capacity),
	}
}

// NewFastAppleList constructs a new list containing the supplied values, if any.
func NewFastAppleList(values ...Apple) *FastAppleList {
	list := MakeFastAppleList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertFastAppleList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertFastAppleList(values ...interface{}) (*FastAppleList, bool) {
	list := MakeFastAppleList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case Apple:
			list.m = append(list.m, j)
		case *Apple:
			list.m = append(list.m, *j)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildFastAppleListFromChan constructs a new FastAppleList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildFastAppleListFromChan(source <-chan Apple) *FastAppleList {
	list := MakeFastAppleList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *FastAppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *FastAppleList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *FastAppleList) slice() []Apple {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *FastAppleList) ToList() *FastAppleList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *FastAppleList) ToSlice() []Apple {
	if list == nil {
		return nil
	}

	s := make([]Apple, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *FastAppleList) ToInterfaceSlice() []interface{} {
	if list == nil {
		return nil
	}

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list *FastAppleList) Clone() *FastAppleList {
	if list == nil {
		return nil
	}

	return NewFastAppleList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *FastAppleList) Get(i int) Apple {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *FastAppleList) Head() Apple {

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *FastAppleList) HeadOption() (Apple, bool) {
	if list == nil {
		return *(new(Apple)), false
	}

	if len(list.m) == 0 {
		return *(new(Apple)), false
	}
	return list.m[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *FastAppleList) Last() Apple {

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *FastAppleList) LastOption() (Apple, bool) {
	if list == nil {
		return *(new(Apple)), false
	}

	if len(list.m) == 0 {
		return *(new(Apple)), false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *FastAppleList) Tail() *FastAppleList {

	result := MakeFastAppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *FastAppleList) Init() *FastAppleList {

	result := MakeFastAppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether FastAppleList is empty.
func (list *FastAppleList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether FastAppleList is empty.
func (list *FastAppleList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *FastAppleList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *FastAppleList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *FastAppleList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *FastAppleList) Contains(v Apple) bool {
	return list.Exists(func(x Apple) bool {
		return v == x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *FastAppleList) ContainsAll(i ...Apple) bool {
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

// Exists verifies that one or more elements of FastAppleList return true for the predicate p.
func (list *FastAppleList) Exists(p func(Apple) bool) bool {
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

// Forall verifies that all elements of FastAppleList return true for the predicate p.
func (list *FastAppleList) Forall(p func(Apple) bool) bool {
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

// Foreach iterates over FastAppleList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *FastAppleList) Foreach(f func(Apple)) {
	if list == nil {
		return
	}

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list *FastAppleList) Send() <-chan Apple {
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

// Reverse returns a copy of FastAppleList with all elements in the reverse order.
//
// The original list is not modified.
func (list *FastAppleList) Reverse() *FastAppleList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := MakeFastAppleList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a FastAppleList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *FastAppleList) DoReverse() *FastAppleList {
	if list == nil {
		return nil
	}

	mid := (len(list.m) + 1) / 2
	last := len(list.m) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list.m[i], list.m[r] = list.m[r], list.m[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of FastAppleList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *FastAppleList) Shuffle() *FastAppleList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled FastAppleList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *FastAppleList) DoShuffle() *FastAppleList {
	if list == nil {
		return nil
	}

	return list.doShuffle()
}

func (list *FastAppleList) doShuffle() *FastAppleList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *FastAppleList) Clear() {
	if list != nil {
		// could use list.m[:0] here but it may not free the dropped elements
		// until their array elements are overwritten
		list.m = make([]Apple, 0, cap(list.m))
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *FastAppleList) Add(more ...Apple) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *FastAppleList) Append(more ...Apple) *FastAppleList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeFastAppleList(0, len(more))
	}

	return list.doAppend(more...)
}

func (list *FastAppleList) doAppend(more ...Apple) *FastAppleList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a FastAppleList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *FastAppleList) DoInsertAt(index int, more ...Apple) *FastAppleList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeFastAppleList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	return list.doInsertAt(index, more...)
}

func (list *FastAppleList) doInsertAt(index int, more ...Apple) *FastAppleList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]Apple, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a FastAppleList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastAppleList) DoDeleteFirst(n int) *FastAppleList {
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a FastAppleList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastAppleList) DoDeleteLast(n int) *FastAppleList {
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a FastAppleList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *FastAppleList) DoDeleteAt(index, n int) *FastAppleList {
	return list.doDeleteAt(index, n)
}

func (list *FastAppleList) doDeleteAt(index, n int) *FastAppleList {
	if n == 0 {
		return list
	}

	newlist := make([]Apple, 0, len(list.m)-n)

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

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a FastAppleList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *FastAppleList) DoKeepWhere(p func(Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	return list.doKeepWhere(p)
}

func (list *FastAppleList) doKeepWhere(p func(Apple) bool) *FastAppleList {
	result := make([]Apple, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of FastAppleList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *FastAppleList) Take(n int) *FastAppleList {
	if list == nil {
		return nil
	}

	if n >= len(list.m) {
		return list
	}

	result := MakeFastAppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of FastAppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastAppleList) Drop(n int) *FastAppleList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := MakeFastAppleList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of FastAppleList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *FastAppleList) TakeLast(n int) *FastAppleList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeFastAppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of FastAppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastAppleList) DropLast(n int) *FastAppleList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeFastAppleList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new FastAppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *FastAppleList) TakeWhile(p func(Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new FastAppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *FastAppleList) DropWhile(p func(Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(0, 0)
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
func (list *FastAppleList) Find(p func(Apple) bool) (Apple, bool) {
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

// Filter returns a new FastAppleList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *FastAppleList) Filter(p func(Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new FastAppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *FastAppleList) Partition(p func(Apple) bool) (*FastAppleList, *FastAppleList) {
	if list == nil {
		return nil, nil
	}

	matching := MakeFastAppleList(0, len(list.m))
	others := MakeFastAppleList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new FastAppleList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastAppleList) Map(f func(Apple) Apple) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new FastAppleList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastAppleList) FlatMap(f func(Apple) []Apple) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of FastAppleList that return true for the predicate p.
func (list *FastAppleList) CountBy(p func(Apple) bool) (result int) {
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

// Fold aggregates all the values in the list using a supplied function, starting from some initial value.
func (list *FastAppleList) Fold(initial Apple, fn func(Apple, Apple) Apple) Apple {

	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of FastAppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *FastAppleList) MinBy(less func(Apple, Apple) bool) Apple {

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

// MaxBy returns an element of FastAppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *FastAppleList) MaxBy(less func(Apple, Apple) bool) Apple {

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

// DistinctBy returns a new FastAppleList whose elements are unique, where equality is defined by the equal function.
func (list *FastAppleList) DistinctBy(equal func(Apple, Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	result := MakeFastAppleList(0, len(list.m))
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
func (list *FastAppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *FastAppleList) IndexWhere2(p func(Apple) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *FastAppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *FastAppleList) LastIndexWhere2(p func(Apple) bool, before int) int {

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
func (list *FastAppleList) Equals(other *FastAppleList) bool {
	if list == nil {
		if other == nil {
			return true
		}
		return len(other.m) == 0
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

type sortableFastAppleList struct {
	less func(i, j Apple) bool
	m    []Apple
}

func (sl sortableFastAppleList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableFastAppleList) Len() int {
	return len(sl.m)
}

func (sl sortableFastAppleList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastAppleList) SortBy(less func(i, j Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	sort.Sort(sortableFastAppleList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *FastAppleList) StableSortBy(less func(i, j Apple) bool) *FastAppleList {
	if list == nil {
		return nil
	}

	sort.Stable(sortableFastAppleList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (list *FastAppleList) GobDecode(b []byte) error {

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (list FastAppleList) GobEncode() ([]byte, error) {

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list.m)
	return buf.Bytes(), err
}
