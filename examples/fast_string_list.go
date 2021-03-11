// An encapsulated []string.
// Not thread-safe.
//
// Generated from fast/list.tpl with Type=string
// options: Comparable:true Numeric:<no value> Integer:<no value> Ordered:true
//          StringLike:<no value> StringParser:<no value> Stringer:true
// GobEncode:<no value> Mutable:always ToList:always ToSet:<no value> MapTo:int
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// FastStringList contains a slice of type string.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FastStringList struct {
	m []string
}

//-------------------------------------------------------------------------------------------------

// MakeFastStringList makes an empty list with both length and capacity initialised.
func MakeFastStringList(length, capacity int) *FastStringList {
	return &FastStringList{
		m: make([]string, length, capacity),
	}
}

// NewFastStringList constructs a new list containing the supplied values, if any.
func NewFastStringList(values ...string) *FastStringList {
	list := MakeFastStringList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertFastStringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertFastStringList(values ...interface{}) (*FastStringList, bool) {
	list := MakeFastStringList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case string:
			list.m = append(list.m, j)
		case *string:
			list.m = append(list.m, *j)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildFastStringListFromChan constructs a new FastStringList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildFastStringListFromChan(source <-chan string) *FastStringList {
	list := MakeFastStringList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *FastStringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *FastStringList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *FastStringList) slice() []string {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *FastStringList) ToList() *FastStringList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *FastStringList) ToSlice() []string {
	if list == nil {
		return nil
	}

	s := make([]string, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *FastStringList) ToInterfaceSlice() []interface{} {
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
func (list *FastStringList) Clone() *FastStringList {
	if list == nil {
		return nil
	}

	return NewFastStringList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *FastStringList) Get(i int) string {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *FastStringList) Head() string {

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *FastStringList) HeadOption() (string, bool) {
	if list == nil {
		return "", false
	}

	if len(list.m) == 0 {
		return "", false
	}
	return list.m[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *FastStringList) Last() string {

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *FastStringList) LastOption() (string, bool) {
	if list == nil {
		return "", false
	}

	if len(list.m) == 0 {
		return "", false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *FastStringList) Tail() *FastStringList {

	result := MakeFastStringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *FastStringList) Init() *FastStringList {

	result := MakeFastStringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether FastStringList is empty.
func (list *FastStringList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether FastStringList is empty.
func (list *FastStringList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *FastStringList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *FastStringList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *FastStringList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *FastStringList) Contains(v string) bool {
	return list.Exists(func(x string) bool {
		return v == x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *FastStringList) ContainsAll(i ...string) bool {
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

// Exists verifies that one or more elements of FastStringList return true for the predicate p.
func (list *FastStringList) Exists(p func(string) bool) bool {
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

// Forall verifies that all elements of FastStringList return true for the predicate p.
func (list *FastStringList) Forall(p func(string) bool) bool {
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

// Foreach iterates over FastStringList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *FastStringList) Foreach(f func(string)) {
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
func (list *FastStringList) Send() <-chan string {
	ch := make(chan string)
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

// Reverse returns a copy of FastStringList with all elements in the reverse order.
//
// The original list is not modified.
func (list *FastStringList) Reverse() *FastStringList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := MakeFastStringList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a FastStringList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *FastStringList) DoReverse() *FastStringList {
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

// Shuffle returns a shuffled copy of FastStringList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *FastStringList) Shuffle() *FastStringList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled FastStringList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *FastStringList) DoShuffle() *FastStringList {
	if list == nil {
		return nil
	}

	return list.doShuffle()
}

func (list *FastStringList) doShuffle() *FastStringList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *FastStringList) Clear() {
	if list != nil {
		// could use list.m[:0] here but it may not free the dropped elements
		// until their array elements are overwritten
		list.m = make([]string, 0, cap(list.m))
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *FastStringList) Add(more ...string) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *FastStringList) Append(more ...string) *FastStringList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeFastStringList(0, len(more))
	}

	return list.doAppend(more...)
}

func (list *FastStringList) doAppend(more ...string) *FastStringList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a FastStringList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *FastStringList) DoInsertAt(index int, more ...string) *FastStringList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeFastStringList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	return list.doInsertAt(index, more...)
}

func (list *FastStringList) doInsertAt(index int, more ...string) *FastStringList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]string, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a FastStringList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastStringList) DoDeleteFirst(n int) *FastStringList {
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a FastStringList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastStringList) DoDeleteLast(n int) *FastStringList {
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a FastStringList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *FastStringList) DoDeleteAt(index, n int) *FastStringList {
	return list.doDeleteAt(index, n)
}

func (list *FastStringList) doDeleteAt(index, n int) *FastStringList {
	if n == 0 {
		return list
	}

	newlist := make([]string, 0, len(list.m)-n)

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

// DoKeepWhere modifies a FastStringList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *FastStringList) DoKeepWhere(p func(string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	return list.doKeepWhere(p)
}

func (list *FastStringList) doKeepWhere(p func(string) bool) *FastStringList {
	result := make([]string, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of FastStringList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *FastStringList) Take(n int) *FastStringList {
	if list == nil {
		return nil
	}

	if n >= len(list.m) {
		return list
	}

	result := MakeFastStringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of FastStringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastStringList) Drop(n int) *FastStringList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := MakeFastStringList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of FastStringList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *FastStringList) TakeLast(n int) *FastStringList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeFastStringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of FastStringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastStringList) DropLast(n int) *FastStringList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeFastStringList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new FastStringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *FastStringList) TakeWhile(p func(string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new FastStringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *FastStringList) DropWhile(p func(string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(0, 0)
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

// Find returns the first string that returns true for predicate p.
// False is returned if none match.
func (list *FastStringList) Find(p func(string) bool) (string, bool) {
	if list == nil {
		return "", false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty string
	return empty, false
}

// Filter returns a new FastStringList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *FastStringList) Filter(p func(string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new FastStringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *FastStringList) Partition(p func(string) bool) (*FastStringList, *FastStringList) {
	if list == nil {
		return nil, nil
	}

	matching := MakeFastStringList(0, len(list.m))
	others := MakeFastStringList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new FastStringList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastStringList) Map(f func(string) string) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// MapToInt returns a new []int by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastStringList) MapToInt(f func(string) int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, len(list.m))

	for i, v := range list.m {
		result[i] = f(v)
	}

	return result
}

// FlatMap returns a new FastStringList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastStringList) FlatMap(f func(string) []string) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// FlatMapToInt returns a new []int by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastStringList) FlatMapToInt(f func(string) []int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, 0, len(list.m))

	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of FastStringList that return true for the predicate p.
func (list *FastStringList) CountBy(p func(string) bool) (result int) {
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
func (list *FastStringList) Fold(initial string, fn func(string, string) string) string {

	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of FastStringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *FastStringList) MinBy(less func(string, string) bool) string {

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

// MaxBy returns an element of FastStringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *FastStringList) MaxBy(less func(string, string) bool) string {

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

// DistinctBy returns a new FastStringList whose elements are unique, where equality is defined by the equal function.
func (list *FastStringList) DistinctBy(equal func(string, string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	result := MakeFastStringList(0, len(list.m))
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
func (list *FastStringList) IndexWhere(p func(string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *FastStringList) IndexWhere2(p func(string) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *FastStringList) LastIndexWhere(p func(string) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *FastStringList) LastIndexWhere2(p func(string) bool, before int) int {

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
// These methods are included when string is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *FastStringList) Equals(other *FastStringList) bool {
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

type sortableFastStringList struct {
	less func(i, j string) bool
	m    []string
}

func (sl sortableFastStringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableFastStringList) Len() int {
	return len(sl.m)
}

func (sl sortableFastStringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastStringList) SortBy(less func(i, j string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	sort.Sort(sortableFastStringList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *FastStringList) StableSortBy(less func(i, j string) bool) *FastStringList {
	if list == nil {
		return nil
	}

	sort.Stable(sortableFastStringList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when string is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastStringList) Sorted() *FastStringList {
	return list.SortBy(func(a, b string) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastStringList) StableSorted() *FastStringList {
	return list.StableSortBy(func(a, b string) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *FastStringList) Min() string {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *FastStringList) Max() (result string) {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v > m {
			m = v
		}
	}
	return m
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *FastStringList) StringList() []string {
	return list.ToSlice()
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *FastStringList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *FastStringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *FastStringList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list FastStringList) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this list type.
func (list *FastStringList) UnmarshalJSON(b []byte) error {

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list FastStringList) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(list.m)
	return buf, err
}
