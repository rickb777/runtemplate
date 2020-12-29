// An encapsulated []big.Int.
// Thread-safe.
//
// Generated from threadsafe/list.tpl with Type=big.Int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> StringLike:<no value> Stringer:true
// GobEncode:<no value> Mutable:always ToList:always ToSet:false MapTo:string,int
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"sync"
)

// X1IntegerList contains a slice of type big.Int.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1IntegerList struct {
	s *sync.RWMutex
	m []big.Int
}

//-------------------------------------------------------------------------------------------------

// MakeX1IntegerList makes an empty list with both length and capacity initialised.
func MakeX1IntegerList(length, capacity int) *X1IntegerList {
	return &X1IntegerList{
		s: &sync.RWMutex{},
		m: make([]big.Int, length, capacity),
	}
}

// NewX1IntegerList constructs a new list containing the supplied values, if any.
func NewX1IntegerList(values ...big.Int) *X1IntegerList {
	list := MakeX1IntegerList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertX1IntegerList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1IntegerList(values ...interface{}) (*X1IntegerList, bool) {
	list := MakeX1IntegerList(0, len(values))

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
	list := MakeX1IntegerList(0, 0)
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
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]big.Int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1IntegerList) ToInterfaceSlice() []interface{} {
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
func (list *X1IntegerList) Clone() *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewX1IntegerList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *X1IntegerList) Get(i int) big.Int {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *X1IntegerList) Head() big.Int {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1IntegerList) HeadOption() (big.Int, bool) {
	if list == nil {
		return *(new(big.Int)), false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return *(new(big.Int)), false
	}
	return list.m[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *X1IntegerList) Last() big.Int {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *X1IntegerList) LastOption() (big.Int, bool) {
	if list == nil {
		return *(new(big.Int)), false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return *(new(big.Int)), false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *X1IntegerList) Tail() *X1IntegerList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *X1IntegerList) Init() *X1IntegerList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, 0)
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

	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *X1IntegerList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *X1IntegerList) Swap(i, j int) {
	list.s.Lock()
	defer list.s.Unlock()

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of X1IntegerList return true for the predicate p.
func (list *X1IntegerList) Exists(p func(big.Int) bool) bool {
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

// Forall verifies that all elements of X1IntegerList return true for the predicate p.
func (list *X1IntegerList) Forall(p func(big.Int) bool) bool {
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

// Foreach iterates over X1IntegerList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *X1IntegerList) Foreach(f func(big.Int)) {
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
func (list *X1IntegerList) Send() <-chan big.Int {
	ch := make(chan big.Int)
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

// Reverse returns a copy of X1IntegerList with all elements in the reverse order.
//
// The original list is not modified.
func (list *X1IntegerList) Reverse() *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	n := len(list.m)
	result := MakeX1IntegerList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a X1IntegerList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *X1IntegerList) DoReverse() *X1IntegerList {
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

// Shuffle returns a shuffled copy of X1IntegerList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *X1IntegerList) Shuffle() *X1IntegerList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled X1IntegerList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *X1IntegerList) DoShuffle() *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doShuffle()
}

func (list *X1IntegerList) doShuffle() *X1IntegerList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *X1IntegerList) Clear() {
	if list != nil {
		list.s.Lock()
		defer list.s.Unlock()
		list.m = list.m[:]
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *X1IntegerList) Add(more ...big.Int) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *X1IntegerList) Append(more ...big.Int) *X1IntegerList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeX1IntegerList(0, len(more))
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doAppend(more...)
}

func (list *X1IntegerList) doAppend(more ...big.Int) *X1IntegerList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a X1IntegerList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *X1IntegerList) DoInsertAt(index int, more ...big.Int) *X1IntegerList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeX1IntegerList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doInsertAt(index, more...)
}

func (list *X1IntegerList) doInsertAt(index int, more ...big.Int) *X1IntegerList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]big.Int, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a X1IntegerList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *X1IntegerList) DoDeleteFirst(n int) *X1IntegerList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a X1IntegerList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *X1IntegerList) DoDeleteLast(n int) *X1IntegerList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a X1IntegerList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *X1IntegerList) DoDeleteAt(index, n int) *X1IntegerList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(index, n)
}

func (list *X1IntegerList) doDeleteAt(index, n int) *X1IntegerList {
	if n == 0 {
		return list
	}

	newlist := make([]big.Int, 0, len(list.m)-n)

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

// DoKeepWhere modifies a X1IntegerList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *X1IntegerList) DoKeepWhere(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doKeepWhere(p)
}

func (list *X1IntegerList) doKeepWhere(p func(big.Int) bool) *X1IntegerList {
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

// Take returns a slice of X1IntegerList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *X1IntegerList) Take(n int) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return list
	}

	result := MakeX1IntegerList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1IntegerList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *X1IntegerList) Drop(n int) *X1IntegerList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return nil
	}

	result := MakeX1IntegerList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of X1IntegerList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *X1IntegerList) TakeLast(n int) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeX1IntegerList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1IntegerList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *X1IntegerList) DropLast(n int) *X1IntegerList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeX1IntegerList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new X1IntegerList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *X1IntegerList) TakeWhile(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, 0)
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
//
// The original list is not modified.
func (list *X1IntegerList) DropWhile(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, 0)
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

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty big.Int
	return empty, false
}

// Filter returns a new X1IntegerList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *X1IntegerList) Filter(p func(big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new X1IntegerLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *X1IntegerList) Partition(p func(big.Int) bool) (*X1IntegerList, *X1IntegerList) {
	if list == nil {
		return nil, nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	matching := MakeX1IntegerList(0, len(list.m))
	others := MakeX1IntegerList(0, len(list.m))

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
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) Map(f func(big.Int) big.Int) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := MakeX1IntegerList(len(list.m), len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) MapToString(f func(big.Int) string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result[i] = f(v)
	}

	return result
}

// MapToInt returns a new []int by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) MapToInt(f func(big.Int) int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result[i] = f(v)
	}

	return result
}

// FlatMap returns a new X1IntegerList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) FlatMap(f func(big.Int) []big.Int) *X1IntegerList {
	if list == nil {
		return nil
	}

	result := MakeX1IntegerList(0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) FlatMapToString(f func(big.Int) []string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, 0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}

// FlatMapToInt returns a new []int by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntegerList) FlatMapToInt(f func(big.Int) []int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, 0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X1IntegerList that return true for the predicate p.
func (list *X1IntegerList) CountBy(p func(big.Int) bool) (result int) {
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

// Fold aggregates all the values in the list using a supplied function, starting from some initial value.
func (list *X1IntegerList) Fold(initial big.Int, fn func(big.Int, big.Int) big.Int) big.Int {
	list.s.RLock()
	defer list.s.RUnlock()

	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of X1IntegerList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntegerList) MinBy(less func(big.Int, big.Int) bool) big.Int {
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

// MaxBy returns an element of X1IntegerList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntegerList) MaxBy(less func(big.Int, big.Int) bool) big.Int {
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

// DistinctBy returns a new X1IntegerList whose elements are unique, where equality is defined by the equal function.
func (list *X1IntegerList) DistinctBy(equal func(big.Int, big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeX1IntegerList(0, len(list.m))
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
func (list *X1IntegerList) LastIndexWhere(p func(big.Int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *X1IntegerList) LastIndexWhere2(p func(big.Int) bool, before int) int {
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

type sortableX1IntegerList struct {
	less func(i, j big.Int) bool
	m    []big.Int
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

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *X1IntegerList) SortBy(less func(i, j big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Sort(sortableX1IntegerList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *X1IntegerList) StableSortBy(less func(i, j big.Int) bool) *X1IntegerList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Stable(sortableX1IntegerList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *X1IntegerList) StringList() []string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	ss := make([]string, len(list.m))
	for i, v := range list.m {
		ss[i] = fmt.Sprintf("%v", v)
	}
	return ss
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *X1IntegerList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *X1IntegerList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *X1IntegerList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list X1IntegerList) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
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
func (list *X1IntegerList) UnmarshalJSON(b []byte) error {
	list.s.Lock()
	defer list.s.Unlock()

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list X1IntegerList) MarshalJSON() ([]byte, error) {
	list.s.RLock()
	defer list.s.RUnlock()

	buf, err := json.Marshal(list.m)
	return buf, err
}
