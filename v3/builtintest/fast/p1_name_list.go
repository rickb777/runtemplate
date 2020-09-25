// An encapsulated []*Name.
// Not thread-safe.
//
// Generated from fast/list.tpl with Type=*Name
// options: Comparable:true Numeric:<no value> Ordered:false StringLike:true Stringer:true
// GobEncode:<no value> Mutable:always ToList:always ToSet:true MapTo:<no value>
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

// P1NameList contains a slice of type *Name.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type P1NameList struct {
	m []*Name
}

//-------------------------------------------------------------------------------------------------

// MakeP1NameList makes an empty list with both length and capacity initialised.
func MakeP1NameList(length, capacity int) *P1NameList {
	return &P1NameList{
		m: make([]*Name, length, capacity),
	}
}

// NewP1NameList constructs a new list containing the supplied values, if any.
func NewP1NameList(values ...*Name) *P1NameList {
	list := MakeP1NameList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertP1NameList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertP1NameList(values ...interface{}) (*P1NameList, bool) {
	list := MakeP1NameList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case Name:
			list.m = append(list.m, &j)
		case *Name:
			list.m = append(list.m, j)
		default:
			if s, ok := i.(fmt.Stringer); ok {
				k := Name(s.String())
				list.m = append(list.m, &k)
			}
		}
	}

	return list, len(list.m) == len(values)
}

// BuildP1NameListFromChan constructs a new P1NameList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildP1NameListFromChan(source <-chan *Name) *P1NameList {
	list := MakeP1NameList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *P1NameList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *P1NameList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *P1NameList) slice() []*Name {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *P1NameList) ToList() *P1NameList {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *P1NameList) ToSet() *P1NameSet {
	if list == nil {
		return nil
	}

	return NewP1NameSet(list.m...)
}

// ToSlice returns the elements of the current list as a slice.
func (list *P1NameList) ToSlice() []*Name {
	if list == nil {
		return nil
	}

	s := make([]*Name, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *P1NameList) ToInterfaceSlice() []interface{} {
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
func (list *P1NameList) Clone() *P1NameList {
	if list == nil {
		return nil
	}

	return NewP1NameList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *P1NameList) Get(i int) *Name {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *P1NameList) Head() *Name {

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns nil.
func (list *P1NameList) HeadOption() *Name {
	if list == nil {
		return nil
	}

	if len(list.m) == 0 {
		return nil
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *P1NameList) Last() *Name {

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns nil.
func (list *P1NameList) LastOption() *Name {
	if list == nil {
		return nil
	}

	if len(list.m) == 0 {
		return nil
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *P1NameList) Tail() *P1NameList {

	result := MakeP1NameList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *P1NameList) Init() *P1NameList {

	result := MakeP1NameList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether P1NameList is empty.
func (list *P1NameList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether P1NameList is empty.
func (list *P1NameList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *P1NameList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *P1NameList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *P1NameList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *P1NameList) Contains(v *Name) bool {
	return list.Exists(func(x *Name) bool {
		return *v == *x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *P1NameList) ContainsAll(i ...*Name) bool {
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

// Exists verifies that one or more elements of P1NameList return true for the predicate p.
func (list *P1NameList) Exists(p func(*Name) bool) bool {
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

// Forall verifies that all elements of P1NameList return true for the predicate p.
func (list *P1NameList) Forall(p func(*Name) bool) bool {
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

// Foreach iterates over P1NameList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *P1NameList) Foreach(f func(*Name)) {
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
func (list *P1NameList) Send() <-chan *Name {
	ch := make(chan *Name)
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

// Reverse returns a copy of P1NameList with all elements in the reverse order.
//
// The original list is not modified.
func (list *P1NameList) Reverse() *P1NameList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := MakeP1NameList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a P1NameList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *P1NameList) DoReverse() *P1NameList {
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

// Shuffle returns a shuffled copy of P1NameList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *P1NameList) Shuffle() *P1NameList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled P1NameList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *P1NameList) DoShuffle() *P1NameList {
	if list == nil {
		return nil
	}

	return list.doShuffle()
}

func (list *P1NameList) doShuffle() *P1NameList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *P1NameList) Clear() {
	if list != nil {
		list.m = list.m[:]
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *P1NameList) Add(more ...*Name) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *P1NameList) Append(more ...*Name) *P1NameList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeP1NameList(0, len(more))
	}

	return list.doAppend(more...)
}

func (list *P1NameList) doAppend(more ...*Name) *P1NameList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a P1NameList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *P1NameList) DoInsertAt(index int, more ...*Name) *P1NameList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeP1NameList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	return list.doInsertAt(index, more...)
}

func (list *P1NameList) doInsertAt(index int, more ...*Name) *P1NameList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]*Name, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a P1NameList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1NameList) DoDeleteFirst(n int) *P1NameList {
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a P1NameList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1NameList) DoDeleteLast(n int) *P1NameList {
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a P1NameList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *P1NameList) DoDeleteAt(index, n int) *P1NameList {
	return list.doDeleteAt(index, n)
}

func (list *P1NameList) doDeleteAt(index, n int) *P1NameList {
	if n == 0 {
		return list
	}

	newlist := make([]*Name, 0, len(list.m)-n)

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

// DoKeepWhere modifies a P1NameList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *P1NameList) DoKeepWhere(p func(*Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	return list.doKeepWhere(p)
}

func (list *P1NameList) doKeepWhere(p func(*Name) bool) *P1NameList {
	result := make([]*Name, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of P1NameList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *P1NameList) Take(n int) *P1NameList {
	if list == nil {
		return nil
	}

	if n >= len(list.m) {
		return list
	}

	result := MakeP1NameList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of P1NameList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1NameList) Drop(n int) *P1NameList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := MakeP1NameList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of P1NameList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *P1NameList) TakeLast(n int) *P1NameList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeP1NameList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of P1NameList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1NameList) DropLast(n int) *P1NameList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeP1NameList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new P1NameList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *P1NameList) TakeWhile(p func(*Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new P1NameList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *P1NameList) DropWhile(p func(*Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(0, 0)
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

// Find returns the first Name that returns true for predicate p.
// False is returned if none match.
func (list *P1NameList) Find(p func(*Name) bool) (*Name, bool) {
	if list == nil {
		return nil, false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	return nil, false
}

// Filter returns a new P1NameList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *P1NameList) Filter(p func(*Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new P1NameLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *P1NameList) Partition(p func(*Name) bool) (*P1NameList, *P1NameList) {
	if list == nil {
		return nil, nil
	}

	matching := MakeP1NameList(0, len(list.m))
	others := MakeP1NameList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new P1NameList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1NameList) Map(f func(*Name) *Name) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new P1NameList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1NameList) FlatMap(f func(*Name) []*Name) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of P1NameList that return true for the predicate p.
func (list *P1NameList) CountBy(p func(*Name) bool) (result int) {
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

// MinBy returns an element of P1NameList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *P1NameList) MinBy(less func(*Name, *Name) bool) *Name {

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

// MaxBy returns an element of P1NameList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *P1NameList) MaxBy(less func(*Name, *Name) bool) *Name {

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

// DistinctBy returns a new P1NameList whose elements are unique, where equality is defined by the equal function.
func (list *P1NameList) DistinctBy(equal func(*Name, *Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	result := MakeP1NameList(0, len(list.m))
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
func (list *P1NameList) IndexWhere(p func(*Name) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *P1NameList) IndexWhere2(p func(*Name) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *P1NameList) LastIndexWhere(p func(*Name) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *P1NameList) LastIndexWhere2(p func(*Name) bool, before int) int {

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
// These methods are included when Name is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *P1NameList) Equals(other *P1NameList) bool {
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

type sortableP1NameList struct {
	less func(i, j *Name) bool
	m    []*Name
}

func (sl sortableP1NameList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableP1NameList) Len() int {
	return len(sl.m)
}

func (sl sortableP1NameList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *P1NameList) SortBy(less func(i, j *Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	sort.Sort(sortableP1NameList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *P1NameList) StableSortBy(less func(i, j *Name) bool) *P1NameList {
	if list == nil {
		return nil
	}

	sort.Stable(sortableP1NameList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *P1NameList) StringList() []string {
	if list == nil {
		return nil
	}

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *P1NameList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *P1NameList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *P1NameList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list P1NameList) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (list *P1NameList) UnmarshalJSON(b []byte) error {

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list P1NameList) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(list.m)
	return buf, err
}
