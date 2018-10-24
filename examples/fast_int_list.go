// An encapsulated []int.
//
// Generated from fast/list.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true GobEncode:<no value> Mutable:always

package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

// FastIntList contains a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FastIntList struct {
	m []int
}

//-------------------------------------------------------------------------------------------------

// MakeFastIntList makes an empty list with both length and capacity initialised.
func MakeFastIntList(length, capacity int) *FastIntList {
	return &FastIntList{
		m: make([]int, length, capacity),
	}
}

// NewFastIntList constructs a new list containing the supplied values, if any.
func NewFastIntList(values ...int) *FastIntList {
	result := MakeFastIntList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertFastIntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertFastIntList(values ...interface{}) (*FastIntList, bool) {
	result := MakeFastIntList(0, len(values))

	for _, i := range values {
		switch i.(type) {
		case int:
			result.m = append(result.m, int(i.(int)))
		case int8:
			result.m = append(result.m, int(i.(int8)))
		case int16:
			result.m = append(result.m, int(i.(int16)))
		case int32:
			result.m = append(result.m, int(i.(int32)))
		case int64:
			result.m = append(result.m, int(i.(int64)))
		case uint:
			result.m = append(result.m, int(i.(uint)))
		case uint8:
			result.m = append(result.m, int(i.(uint8)))
		case uint16:
			result.m = append(result.m, int(i.(uint16)))
		case uint32:
			result.m = append(result.m, int(i.(uint32)))
		case uint64:
			result.m = append(result.m, int(i.(uint64)))
		case float32:
			result.m = append(result.m, int(i.(float32)))
		case float64:
			result.m = append(result.m, int(i.(float64)))
		}
	}

	return result, len(result.m) == len(values)
}

// BuildFastIntListFromChan constructs a new FastIntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildFastIntListFromChan(source <-chan int) *FastIntList {
	result := MakeFastIntList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *FastIntList) ToSlice() []int {

	s := make([]int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *FastIntList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *FastIntList) Clone() *FastIntList {

	return NewFastIntList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *FastIntList) Get(i int) int {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *FastIntList) Head() int {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *FastIntList) Last() int {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *FastIntList) Tail() *FastIntList {

	result := MakeFastIntList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *FastIntList) Init() *FastIntList {

	result := MakeFastIntList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether FastIntList is empty.
func (list *FastIntList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether FastIntList is empty.
func (list *FastIntList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *FastIntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *FastIntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list.
func (list *FastIntList) Size() int {

	return len(list.m)
}

// Swap exchanges two elements.
func (list *FastIntList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines if a given item is already in the list.
func (list *FastIntList) Contains(v int) bool {
	return list.Exists(func(x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *FastIntList) ContainsAll(i ...int) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of FastIntList return true for the predicate p.
func (list *FastIntList) Exists(p func(int) bool) bool {

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of FastIntList return true for the predicate p.
func (list *FastIntList) Forall(p func(int) bool) bool {

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over FastIntList and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (list *FastIntList) Foreach(fn func(int)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *FastIntList) Send() <-chan int {
	ch := make(chan int)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of FastIntList with all elements in the reverse order.
//
// The original list is not modified.
func (list *FastIntList) Reverse() *FastIntList {
	return list.Clone().doReverse()
}

// DoReverse alters a FastIntList with all elements in the reverse order.
//
// The modified list is returned.
func (list *FastIntList) DoReverse() *FastIntList {
	return list.doReverse()
}

func (list *FastIntList) doReverse() *FastIntList {
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

// Shuffle returns a shuffled copy of FastIntList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *FastIntList) Shuffle() *FastIntList {
	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled FastIntList, using a version of the Fisher-Yates shuffle.
//
// The modified list is returned.
func (list *FastIntList) DoShuffle() *FastIntList {
	return list.doShuffle()
}

func (list *FastIntList) doShuffle() *FastIntList {
	numItems := len(list.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current list. This is a synonym for Append.
func (list *FastIntList) Add(more ...int) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *FastIntList) Append(more ...int) *FastIntList {
	return list.doAppend(more...)
}

func (list *FastIntList) doAppend(more ...int) *FastIntList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a FastIntList by inserting elements at a given index.
// This is a generalised version of Append.
//
// The modified list is returned.
// Panics if the index is out of range.
func (list *FastIntList) DoInsertAt(index int, more ...int) *FastIntList {
	return list.doInsertAt(index, more...)
}

func (list *FastIntList) doInsertAt(index int, more ...int) *FastIntList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]int, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a FastIntList by deleting n elements from the start of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastIntList) DoDeleteFirst(n int) *FastIntList {
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a FastIntList by deleting n elements from the end of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *FastIntList) DoDeleteLast(n int) *FastIntList {
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a FastIntList by deleting n elements from a given index.
//
// The modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *FastIntList) DoDeleteAt(index, n int) *FastIntList {
	return list.doDeleteAt(index, n)
}

func (list *FastIntList) doDeleteAt(index, n int) *FastIntList {
	if n == 0 {
		return list
	}

	newlist := make([]int, 0, len(list.m)-n)

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

// DoKeepWhere modifies a FastIntList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The modified list is returned.
func (list *FastIntList) DoKeepWhere(p func(int) bool) *FastIntList {
	return list.doKeepWhere(p)
}

func (list *FastIntList) doKeepWhere(p func(int) bool) *FastIntList {
	result := make([]int, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of FastIntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *FastIntList) Take(n int) *FastIntList {

	if n > len(list.m) {
		return list
	}
	result := MakeFastIntList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of FastIntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastIntList) Drop(n int) *FastIntList {
	if n == 0 {
		return list
	}

	result := MakeFastIntList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of FastIntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *FastIntList) TakeLast(n int) *FastIntList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := MakeFastIntList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of FastIntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *FastIntList) DropLast(n int) *FastIntList {
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

// TakeWhile returns a new FastIntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *FastIntList) TakeWhile(p func(int) bool) *FastIntList {

	result := MakeFastIntList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new FastIntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *FastIntList) DropWhile(p func(int) bool) *FastIntList {

	result := MakeFastIntList(0, 0)
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

// Find returns the first int that returns true for predicate p.
// False is returned if none match.
func (list *FastIntList) Find(p func(int) bool) (int, bool) {

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty int
	return empty, false
}

// Filter returns a new FastIntList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *FastIntList) Filter(p func(int) bool) *FastIntList {

	result := MakeFastIntList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *FastIntList) Partition(p func(int) bool) (*FastIntList, *FastIntList) {

	matching := MakeFastIntList(0, len(list.m)/2)
	others := MakeFastIntList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new FastIntList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastIntList) Map(fn func(int) int) *FastIntList {
	result := MakeFastIntList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new FastIntList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *FastIntList) FlatMap(fn func(int) []int) *FastIntList {
	result := MakeFastIntList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of FastIntList that return true for the passed predicate.
func (list *FastIntList) CountBy(predicate func(int) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of FastIntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *FastIntList) MinBy(less func(int, int) bool) int {

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

// MaxBy returns an element of FastIntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *FastIntList) MaxBy(less func(int, int) bool) int {

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

// DistinctBy returns a new FastIntList whose elements are unique, where equality is defined by a passed func.
func (list *FastIntList) DistinctBy(equal func(int, int) bool) *FastIntList {

	result := MakeFastIntList(0, len(list.m))
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
func (list *FastIntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *FastIntList) IndexWhere2(p func(int) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *FastIntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *FastIntList) LastIndexWhere2(p func(int) bool, before int) int {

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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the list.
func (list *FastIntList) Sum() int {

	sum := int(0)
	for _, v := range list.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *FastIntList) Equals(other *FastIntList) bool {

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

type sortableFastIntList struct {
	less func(i, j int) bool
	m    []int
}

func (sl sortableFastIntList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableFastIntList) Len() int {
	return len(sl.m)
}

func (sl sortableFastIntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastIntList) SortBy(less func(i, j int) bool) *FastIntList {

	sort.Sort(sortableFastIntList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *FastIntList) StableSortBy(less func(i, j int) bool) *FastIntList {

	sort.Stable(sortableFastIntList{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastIntList) Sorted() *FastIntList {
	return list.SortBy(func(a, b int) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *FastIntList) StableSorted() *FastIntList {
	return list.StableSortBy(func(a, b int) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *FastIntList) Min() int {

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
func (list *FastIntList) Max() (result int) {

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
func (list FastIntList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *FastIntList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *FastIntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *FastIntList) MkString3(before, between, after string) string {
	return list.mkString3Bytes(before, between, after).String()
}

func (list FastIntList) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (list *FastIntList) UnmarshalJSON(b []byte) error {

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list FastIntList) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(list.m)
	return buf, err
}
