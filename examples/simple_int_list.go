// A simple type derived from []int
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true
// by runtemplate v2.2.3
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

// SimpleIntList is a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
//
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type SimpleIntList []int

//-------------------------------------------------------------------------------------------------

// MakeSimpleIntList makes an empty list with both length and capacity initialised.
func MakeSimpleIntList(length, capacity int) SimpleIntList {
	return make(SimpleIntList, length, capacity)
}

// NewSimpleIntList constructs a new list containing the supplied values, if any.
func NewSimpleIntList(values ...int) SimpleIntList {
	result := MakeSimpleIntList(len(values), len(values))
	copy(result, values)
	return result
}

// ConvertSimpleIntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertSimpleIntList(values ...interface{}) (SimpleIntList, bool) {
	result := MakeSimpleIntList(0, len(values))

	for _, i := range values {
		switch i.(type) {
		case int:
			result = append(result, int(i.(int)))
		case int8:
			result = append(result, int(i.(int8)))
		case int16:
			result = append(result, int(i.(int16)))
		case int32:
			result = append(result, int(i.(int32)))
		case int64:
			result = append(result, int(i.(int64)))
		case uint:
			result = append(result, int(i.(uint)))
		case uint8:
			result = append(result, int(i.(uint8)))
		case uint16:
			result = append(result, int(i.(uint16)))
		case uint32:
			result = append(result, int(i.(uint32)))
		case uint64:
			result = append(result, int(i.(uint64)))
		case float32:
			result = append(result, int(i.(float32)))
		case float64:
			result = append(result, int(i.(float64)))
		}
	}

	return result, len(result) == len(values)
}

// BuildSimpleIntListFromChan constructs a new SimpleIntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildSimpleIntListFromChan(source <-chan int) SimpleIntList {
	result := MakeSimpleIntList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list SimpleIntList) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list SimpleIntList) Clone() SimpleIntList {
	return NewSimpleIntList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list SimpleIntList) Get(i int) int {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty.
func (list SimpleIntList) Head() int {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list SimpleIntList) HeadOption() int {
	if list.IsEmpty() {
		return 0
	}
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty.
func (list SimpleIntList) Last() int {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list SimpleIntList) LastOption() int {
	if list.IsEmpty() {
		return 0
	}
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty.
func (list SimpleIntList) Tail() SimpleIntList {
	return SimpleIntList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty.
func (list SimpleIntList) Init() SimpleIntList {
	return SimpleIntList(list[:len(list)-1])
}

// IsEmpty tests whether SimpleIntList is empty.
func (list SimpleIntList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether SimpleIntList is empty.
func (list SimpleIntList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for ordered lists and queues.
func (list SimpleIntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list SimpleIntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list SimpleIntList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list SimpleIntList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list SimpleIntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines if a given item is already in the list.
func (list SimpleIntList) Contains(v int) bool {
	return list.Exists(func(x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list SimpleIntList) ContainsAll(i ...int) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of SimpleIntList return true for the passed func.
func (list SimpleIntList) Exists(fn func(int) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of SimpleIntList return true for the passed func.
func (list SimpleIntList) Forall(fn func(int) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over SimpleIntList and executes function fn against each element.
func (list SimpleIntList) Foreach(fn func(int)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list SimpleIntList) Send() <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of SimpleIntList with all elements in the reverse order.
func (list SimpleIntList) Reverse() SimpleIntList {
	numItems := len(list)
	result := MakeSimpleIntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse returns a copy of SimpleIntList with all elements in the reverse order.
// This is an alias for Reverse.
func (list SimpleIntList) DoReverse() SimpleIntList {
	return list.Reverse()
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of SimpleIntList, using a version of the Fisher-Yates shuffle.
func (list SimpleIntList) Shuffle() SimpleIntList {
	result := list.Clone()
	numItems := len(list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of SimpleIntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list SimpleIntList) Take(n int) SimpleIntList {
	if n > len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of SimpleIntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list SimpleIntList) Drop(n int) SimpleIntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of SimpleIntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list SimpleIntList) TakeLast(n int) SimpleIntList {
	l := len(list)
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of SimpleIntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list SimpleIntList) DropLast(n int) SimpleIntList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new SimpleIntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list SimpleIntList) TakeWhile(p func(int) bool) SimpleIntList {
	result := MakeSimpleIntList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new SimpleIntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list SimpleIntList) DropWhile(p func(int) bool) SimpleIntList {
	result := MakeSimpleIntList(0, 0)
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

// Find returns the first int that returns true for predicate p.
// False is returned if none match.
func (list SimpleIntList) Find(p func(int) bool) (int, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}

	var empty int
	return empty, false

}

// Filter returns a new SimpleIntList whose elements return true for predicate p.
//
// The original list is not modified.
func (list SimpleIntList) Filter(p func(int) bool) SimpleIntList {
	result := MakeSimpleIntList(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list SimpleIntList) Partition(p func(int) bool) (SimpleIntList, SimpleIntList) {
	matching := MakeSimpleIntList(0, len(list)/2)
	others := MakeSimpleIntList(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new SimpleIntList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list SimpleIntList) Map(fn func(int) int) SimpleIntList {
	result := MakeSimpleIntList(0, len(list))

	for _, v := range list {
		result = append(result, fn(v))
	}

	return result
}

// FlatMap returns a new SimpleIntList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list SimpleIntList) FlatMap(fn func(int) []int) SimpleIntList {
	result := MakeSimpleIntList(0, len(list))

	for _, v := range list {
		result = append(result, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of SimpleIntList that return true for the passed predicate.
func (list SimpleIntList) CountBy(predicate func(int) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of SimpleIntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list SimpleIntList) MinBy(less func(int, int) bool) int {
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

// MaxBy returns an element of SimpleIntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list SimpleIntList) MaxBy(less func(int, int) bool) int {
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

// DistinctBy returns a new SimpleIntList whose elements are unique, where equality is defined by a passed func.
func (list SimpleIntList) DistinctBy(equal func(int, int) bool) SimpleIntList {
	result := MakeSimpleIntList(0, len(list))
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

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list SimpleIntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list SimpleIntList) IndexWhere2(p func(int) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list SimpleIntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list SimpleIntList) LastIndexWhere2(p func(int) bool, before int) int {
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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the list.
func (list SimpleIntList) Sum() int {
	sum := int(0)
	for _, v := range list {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list SimpleIntList) Equals(other SimpleIntList) bool {
	if list.Size() != other.Size() {
		return false
	}

	for i, v := range list {
		if v != other[i] {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

type sortableSimpleIntList struct {
	less func(i, j int) bool
	m    []int
}

func (sl sortableSimpleIntList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableSimpleIntList) Len() int {
	return len(sl.m)
}

func (sl sortableSimpleIntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list SimpleIntList) SortBy(less func(i, j int) bool) SimpleIntList {

	sort.Sort(sortableSimpleIntList{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list SimpleIntList) StableSortBy(less func(i, j int) bool) SimpleIntList {

	sort.Stable(sortableSimpleIntList{less, list})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
func (list SimpleIntList) Sorted() SimpleIntList {
	return list.SortBy(func(a, b int) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
func (list SimpleIntList) StableSorted() SimpleIntList {
	return list.StableSortBy(func(a, b int) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list SimpleIntList) Min() int {
	m := list.MinBy(func(a int, b int) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list SimpleIntList) Max() (result int) {
	m := list.MaxBy(func(a int, b int) bool {
		return a < b
	})
	return m
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list SimpleIntList) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list SimpleIntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list SimpleIntList) MkString3(before, between, after string) string {
	b := bytes.Buffer{}
	b.WriteString(before)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(between)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(after)
	return b.String()
}
