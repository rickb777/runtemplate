// An encapsulated immutable []Apple.
// Thread-safe.
//
//
// Generated from immutable/list.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:false GobEncode:true Mutable:disabled
// by runtemplate v3.3.3
// See https://github.com/johanbrandhorst/runtemplate/blob/master/v3/BUILTIN.md

package examples

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"sort"
)

// ImmutableAppleList contains a slice of type Apple. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type ImmutableAppleList struct {
	m []Apple
}

//-------------------------------------------------------------------------------------------------

func newImmutableAppleList(length, capacity int) *ImmutableAppleList {
	return &ImmutableAppleList{
		m: make([]Apple, length, capacity),
	}
}

// NewImmutableAppleList constructs a new list containing the supplied values, if any.
func NewImmutableAppleList(values ...Apple) *ImmutableAppleList {
	list := newImmutableAppleList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertImmutableAppleList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertImmutableAppleList(values ...interface{}) (*ImmutableAppleList, bool) {
	list := newImmutableAppleList(0, len(values))

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

// BuildImmutableAppleListFromChan constructs a new ImmutableAppleList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildImmutableAppleListFromChan(source <-chan Apple) *ImmutableAppleList {
	list := newImmutableAppleList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *ImmutableAppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *ImmutableAppleList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *ImmutableAppleList) slice() []Apple {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *ImmutableAppleList) ToList() *ImmutableAppleList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *ImmutableAppleList) ToSlice() []Apple {

	s := make([]Apple, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *ImmutableAppleList) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *ImmutableAppleList) Clone() *ImmutableAppleList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *ImmutableAppleList) Get(i int) Apple {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *ImmutableAppleList) Head() Apple {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *ImmutableAppleList) HeadOption() Apple {
	if list == nil || len(list.m) == 0 {
		var v Apple
		return v
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *ImmutableAppleList) Last() Apple {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *ImmutableAppleList) LastOption() Apple {
	if list == nil || len(list.m) == 0 {
		var v Apple
		return v
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *ImmutableAppleList) Tail() *ImmutableAppleList {
	result := newImmutableAppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *ImmutableAppleList) Init() *ImmutableAppleList {
	result := newImmutableAppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether ImmutableAppleList is empty.
func (list *ImmutableAppleList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether ImmutableAppleList is empty.
func (list *ImmutableAppleList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *ImmutableAppleList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *ImmutableAppleList) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of ImmutableAppleList return true for the predicate p.
func (list *ImmutableAppleList) Exists(p func(Apple) bool) bool {
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

// Forall verifies that all elements of ImmutableAppleList return true for the predicate p.
func (list *ImmutableAppleList) Forall(p func(Apple) bool) bool {
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

// Foreach iterates over ImmutableAppleList and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *ImmutableAppleList) Foreach(f func(Apple)) {
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
func (list *ImmutableAppleList) Send() <-chan Apple {
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

// Reverse returns a copy of ImmutableAppleList with all elements in the reverse order.
func (list *ImmutableAppleList) Reverse() *ImmutableAppleList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := newImmutableAppleList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of ImmutableAppleList, using a version of the Fisher-Yates shuffle.
func (list *ImmutableAppleList) Shuffle() *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := NewImmutableAppleList(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *ImmutableAppleList) Append(more ...Apple) *ImmutableAppleList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return NewImmutableAppleList(more...)
	}

	newList := NewImmutableAppleList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *ImmutableAppleList) doAppend(more ...Apple) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of ImmutableAppleList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *ImmutableAppleList) Take(n int) *ImmutableAppleList {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := newImmutableAppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of ImmutableAppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableAppleList) Drop(n int) *ImmutableAppleList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := newImmutableAppleList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of ImmutableAppleList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *ImmutableAppleList) TakeLast(n int) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := newImmutableAppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of ImmutableAppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableAppleList) DropLast(n int) *ImmutableAppleList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := newImmutableAppleList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new ImmutableAppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *ImmutableAppleList) TakeWhile(p func(Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new ImmutableAppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *ImmutableAppleList) DropWhile(p func(Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(0, 0)
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
func (list *ImmutableAppleList) Find(p func(Apple) bool) (Apple, bool) {
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

// Filter returns a new ImmutableAppleList whose elements return true for predicate p.
func (list *ImmutableAppleList) Filter(p func(Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(0, len(list.m)/2)

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
func (list *ImmutableAppleList) Partition(p func(Apple) bool) (*ImmutableAppleList, *ImmutableAppleList) {
	if list == nil {
		return nil, nil
	}

	matching := newImmutableAppleList(0, len(list.m)/2)
	others := newImmutableAppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new ImmutableAppleList by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableAppleList) Map(f func(Apple) Apple) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new ImmutableAppleList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableAppleList) FlatMap(f func(Apple) []Apple) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of ImmutableAppleList that return true for the predicate p.
func (list *ImmutableAppleList) CountBy(p func(Apple) bool) (result int) {
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

// MinBy returns an element of ImmutableAppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableAppleList) MinBy(less func(Apple, Apple) bool) Apple {
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

// MaxBy returns an element of ImmutableAppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableAppleList) MaxBy(less func(Apple, Apple) bool) Apple {
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

// DistinctBy returns a new ImmutableAppleList whose elements are unique, where equality is defined by the equal function.
func (list *ImmutableAppleList) DistinctBy(equal func(Apple, Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := newImmutableAppleList(0, len(list.m))
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
func (list *ImmutableAppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) IndexWhere2(p func(Apple) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) LastIndexWhere2(p func(Apple) bool, before int) int {

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

type sortableImmutableAppleList struct {
	less func(i, j Apple) bool
	m    []Apple
}

func (sl sortableImmutableAppleList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableImmutableAppleList) Len() int {
	return len(sl.m)
}

func (sl sortableImmutableAppleList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *ImmutableAppleList) SortBy(less func(i, j Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := NewImmutableAppleList(list.m...)
	sort.Sort(sortableImmutableAppleList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *ImmutableAppleList) StableSortBy(less func(i, j Apple) bool) *ImmutableAppleList {
	if list == nil {
		return nil
	}

	result := NewImmutableAppleList(list.m...)
	sort.Stable(sortableImmutableAppleList{less, result.m})
	return result
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (list *ImmutableAppleList) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register Apple with the 'gob' package before this method is used.
func (list ImmutableAppleList) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list.m)
	return buf.Bytes(), err
}
