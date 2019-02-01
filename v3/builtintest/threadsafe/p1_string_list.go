// An encapsulated []*string.
// Thread-safe.
//
// Generated from threadsafe/list.tpl with Type=*string
// options: Comparable:true Numeric:false Ordered:false Stringer:true
// GobEncode:<no value> Mutable:always ToList:always ToSet:true MapTo:<no value>
// by runtemplate v3.2.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

// P1StringList contains a slice of type *string.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type P1StringList struct {
	s *sync.RWMutex
	m []*string
}

//-------------------------------------------------------------------------------------------------

// MakeP1StringList makes an empty list with both length and capacity initialised.
func MakeP1StringList(length, capacity int) *P1StringList {
	return &P1StringList{
		s: &sync.RWMutex{},
		m: make([]*string, length, capacity),
	}
}

// NewP1StringList constructs a new list containing the supplied values, if any.
func NewP1StringList(values ...*string) *P1StringList {
	list := MakeP1StringList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertP1StringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertP1StringList(values ...interface{}) (*P1StringList, bool) {
	list := MakeP1StringList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case string:
			list.m = append(list.m, &j)
		case *string:
			list.m = append(list.m, j)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildP1StringListFromChan constructs a new P1StringList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildP1StringListFromChan(source <-chan *string) *P1StringList {
	list := MakeP1StringList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *P1StringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *P1StringList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *P1StringList) slice() []*string {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *P1StringList) ToList() *P1StringList {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *P1StringList) ToSet() *P1StringSet {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewP1StringSet(list.m...)
}

// ToSlice returns the elements of the current list as a slice.
func (list *P1StringList) ToSlice() []*string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]*string, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *P1StringList) ToInterfaceSlice() []interface{} {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list *P1StringList) Clone() *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewP1StringList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *P1StringList) Get(i int) *string {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *P1StringList) Head() *string {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns nil.
func (list *P1StringList) HeadOption() *string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *P1StringList) Last() *string {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns nil.
func (list *P1StringList) LastOption() *string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *P1StringList) Tail() *P1StringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *P1StringList) Init() *P1StringList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether P1StringList is empty.
func (list *P1StringList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether P1StringList is empty.
func (list *P1StringList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *P1StringList) Size() int {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *P1StringList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *P1StringList) Swap(i, j int) {
	list.s.Lock()
	defer list.s.Unlock()

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *P1StringList) Contains(v *string) bool {
	return list.Exists(func(x *string) bool {
		return *v == *x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *P1StringList) ContainsAll(i ...*string) bool {
	if list == nil {
		return len(i) == 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of P1StringList return true for the predicate p.
func (list *P1StringList) Exists(p func(*string) bool) bool {
	if list == nil {
		return false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of P1StringList return true for the predicate p.
func (list *P1StringList) Forall(p func(*string) bool) bool {
	if list == nil {
		return true
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over P1StringList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *P1StringList) Foreach(f func(*string)) {
	if list == nil {
		return
	}

	list.s.Lock()
	defer list.s.Unlock()

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list *P1StringList) Send() <-chan *string {
	ch := make(chan *string)
	go func() {
		if list != nil {
			list.s.RLock()
			defer list.s.RUnlock()

			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of P1StringList with all elements in the reverse order.
//
// The original list is not modified.
func (list *P1StringList) Reverse() *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	n := len(list.m)
	result := MakeP1StringList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a P1StringList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *P1StringList) DoReverse() *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

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

// Shuffle returns a shuffled copy of P1StringList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *P1StringList) Shuffle() *P1StringList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled P1StringList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *P1StringList) DoShuffle() *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doShuffle()
}

func (list *P1StringList) doShuffle() *P1StringList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *P1StringList) Clear() {
	if list != nil {
		list.s.Lock()
		defer list.s.Unlock()
		list.m = list.m[:]
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *P1StringList) Add(more ...*string) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *P1StringList) Append(more ...*string) *P1StringList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeP1StringList(0, len(more))
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doAppend(more...)
}

func (list *P1StringList) doAppend(more ...*string) *P1StringList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a P1StringList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *P1StringList) DoInsertAt(index int, more ...*string) *P1StringList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeP1StringList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doInsertAt(index, more...)
}

func (list *P1StringList) doInsertAt(index int, more ...*string) *P1StringList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]*string, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a P1StringList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteFirst(n int) *P1StringList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a P1StringList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteLast(n int) *P1StringList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a P1StringList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteAt(index, n int) *P1StringList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(index, n)
}

func (list *P1StringList) doDeleteAt(index, n int) *P1StringList {
	if n == 0 {
		return list
	}

	newlist := make([]*string, 0, len(list.m)-n)

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

// DoKeepWhere modifies a P1StringList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *P1StringList) DoKeepWhere(p func(*string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doKeepWhere(p)
}

func (list *P1StringList) doKeepWhere(p func(*string) bool) *P1StringList {
	result := make([]*string, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of P1StringList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *P1StringList) Take(n int) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return list
	}

	result := MakeP1StringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of P1StringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1StringList) Drop(n int) *P1StringList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return nil
	}

	result := MakeP1StringList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of P1StringList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *P1StringList) TakeLast(n int) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeP1StringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of P1StringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1StringList) DropLast(n int) *P1StringList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeP1StringList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new P1StringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *P1StringList) TakeWhile(p func(*string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new P1StringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *P1StringList) DropWhile(p func(*string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, 0)
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
func (list *P1StringList) Find(p func(*string) bool) (*string, bool) {
	if list == nil {
		return nil, false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	return nil, false
}

// Filter returns a new P1StringList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *P1StringList) Filter(p func(*string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new P1StringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *P1StringList) Partition(p func(*string) bool) (*P1StringList, *P1StringList) {
	if list == nil {
		return nil, nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	matching := MakeP1StringList(0, len(list.m))
	others := MakeP1StringList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new P1StringList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1StringList) Map(f func(*string) *string) *P1StringList {
	if list == nil {
		return nil
	}

	result := MakeP1StringList(len(list.m), len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new P1StringList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1StringList) FlatMap(f func(*string) []*string) *P1StringList {
	if list == nil {
		return nil
	}

	result := MakeP1StringList(0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of P1StringList that return true for the predicate p.
func (list *P1StringList) CountBy(p func(*string) bool) (result int) {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of P1StringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *P1StringList) MinBy(less func(*string, *string) bool) *string {
	list.s.RLock()
	defer list.s.RUnlock()

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

// MaxBy returns an element of P1StringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *P1StringList) MaxBy(less func(*string, *string) bool) *string {
	list.s.RLock()
	defer list.s.RUnlock()

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

// DistinctBy returns a new P1StringList whose elements are unique, where equality is defined by the equal function.
func (list *P1StringList) DistinctBy(equal func(*string, *string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeP1StringList(0, len(list.m))
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
func (list *P1StringList) IndexWhere(p func(*string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *P1StringList) IndexWhere2(p func(*string) bool, from int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *P1StringList) LastIndexWhere(p func(*string) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *P1StringList) LastIndexWhere2(p func(*string) bool, before int) int {
	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *P1StringList) Equals(other *P1StringList) bool {
	if list == nil {
		if other == nil {
			return true
		}
		other.s.RLock()
		defer other.s.RUnlock()
		return len(other.m) == 0
	}

	if other == nil {
		list.s.RLock()
		defer list.s.RUnlock()
		return len(list.m) == 0
	}

	list.s.RLock()
	other.s.RLock()
	defer list.s.RUnlock()
	defer other.s.RUnlock()

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

type sortableP1StringList struct {
	less func(i, j *string) bool
	m []*string
}

func (sl sortableP1StringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableP1StringList) Len() int {
	return len(sl.m)
}

func (sl sortableP1StringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *P1StringList) SortBy(less func(i, j *string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Sort(sortableP1StringList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *P1StringList) StableSortBy(less func(i, j *string) bool) *P1StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Stable(sortableP1StringList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *P1StringList) StringList() []string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *P1StringList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *P1StringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *P1StringList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list P1StringList) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *P1StringList) UnmarshalJSON(b []byte) error {
	list.s.Lock()
	defer list.s.Unlock()

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list P1StringList) MarshalJSON() ([]byte, error) {
	list.s.RLock()
	defer list.s.RUnlock()

	buf, err := json.Marshal(list.m)
	return buf, err
}