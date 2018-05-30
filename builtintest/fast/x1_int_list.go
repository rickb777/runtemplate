// An encapsulated []int.
// Thread-safe.
//
// Generated from fast/list.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always

package fast

import (

	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

// X1IntList contains a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1IntList struct {
	m []int
}


//-------------------------------------------------------------------------------------------------

func newX1IntList(len, cap int) *X1IntList {
	return &X1IntList {
		m: make([]int, len, cap),
	}
}

// NewX1IntList constructs a new list containing the supplied values, if any.
func NewX1IntList(values ...int) *X1IntList {
	result := newX1IntList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertX1IntList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1IntList(values ...interface{}) (*X1IntList, bool) {
	result := newX1IntList(0, len(values))

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

// BuildX1IntListFromChan constructs a new X1IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1IntListFromChan(source <-chan int) *X1IntList {
	result := newX1IntList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *X1IntList) ToSlice() []int {

	s := make([]int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1IntList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *X1IntList) Clone() *X1IntList {
	return NewX1IntList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *X1IntList) Get(i int) int {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *X1IntList) Head() int {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *X1IntList) Last() int {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *X1IntList) Tail() *X1IntList {

	result := newX1IntList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *X1IntList) Init() *X1IntList {

	result := newX1IntList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1IntList is empty.
func (list *X1IntList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1IntList is empty.
func (list *X1IntList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *X1IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *X1IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list.
func (list *X1IntList) Size() int {

	return len(list.m)
}

// Swap exchanges two elements.
func (list *X1IntList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *X1IntList) Contains(v int) bool {
	return list.Exists(func (x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *X1IntList) ContainsAll(i ...int) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1IntList return true for the predicate p.
func (list *X1IntList) Exists(p func(int) bool) bool {

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1IntList return true for the predicate p.
func (list *X1IntList) Forall(p func(int) bool) bool {

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1IntList and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (list *X1IntList) Foreach(fn func(int)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *X1IntList) Send() <-chan int {
	ch := make(chan int)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X1IntList with all elements in the reverse order.
func (list *X1IntList) Reverse() *X1IntList {

	numItems := len(list.m)
	result := newX1IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X1IntList, using a version of the Fisher-Yates shuffle.
func (list *X1IntList) Shuffle() *X1IntList {
	result := list.Clone()
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Add adds items to the current list. This is a synonym for Append.
func (list *X1IntList) Add(more ...int) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *X1IntList) Append(more ...int) *X1IntList {

	list.doAppend(more...)
	return list
}

func (list *X1IntList) doAppend(more ...int) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1IntList) Take(n int) *X1IntList {

	if n > len(list.m) {
		return list
	}
	result := newX1IntList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntList) Drop(n int) *X1IntList {
	if n == 0 {
		return list
	}


	result := newX1IntList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of X1IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1IntList) TakeLast(n int) *X1IntList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newX1IntList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntList) DropLast(n int) *X1IntList {
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

// TakeWhile returns a new X1IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *X1IntList) TakeWhile(p func(int) bool) *X1IntList {

	result := newX1IntList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *X1IntList) DropWhile(p func(int) bool) *X1IntList {

	result := newX1IntList(0, 0)
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

// DeleteAt modifies a X1IntList by deleting n elements from a given index.
// The modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *X1IntList) DeleteAt(index, n int) *X1IntList {

	newlist := make([]int, 0, len(list.m) - n)

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

// KeepWhere modifies a X1IntList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
func (list *X1IntList) KeepWhere(p func(int) bool) *X1IntList {

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

// Find returns the first int that returns true for predicate p.
// False is returned if none match.
func (list X1IntList) Find(p func(int) bool) (int, bool) {

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}


	var empty int
	return empty, false

}

// Filter returns a new X1IntList whose elements return true for predicate p.
//
// The original list is not modified. See also KeepWhere (which does modify the original list).
func (list *X1IntList) Filter(p func(int) bool) *X1IntList {

	result := newX1IntList(0, len(list.m)/2)

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
func (list *X1IntList) Partition(p func(int) bool) (*X1IntList, *X1IntList) {

	matching := newX1IntList(0, len(list.m)/2)
	others := newX1IntList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new X1IntList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntList) Map(fn func(int) int) *X1IntList {
	result := newX1IntList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new X1IntList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1IntList) FlatMap(fn func(int) []int) *X1IntList {
	result := newX1IntList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of X1IntList that return true for the passed predicate.
func (list *X1IntList) CountBy(predicate func(int) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1IntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntList) MinBy(less func(int, int) bool) int {

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

// MaxBy returns an element of X1IntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1IntList) MaxBy(less func(int, int) bool) int {

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

// DistinctBy returns a new X1IntList whose elements are unique, where equality is defined by a passed func.
func (list *X1IntList) DistinctBy(equal func(int, int) bool) *X1IntList {

	result := newX1IntList(0, len(list.m))
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
func (list *X1IntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *X1IntList) IndexWhere2(p func(int) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *X1IntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *X1IntList) LastIndexWhere2(p func(int) bool, before int) int {

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
func (list *X1IntList) Sum() int {

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
func (list *X1IntList) Equals(other *X1IntList) bool {

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

type sortableX1IntList struct {
	less func(i, j int) bool
	m []int
}

func (sl sortableX1IntList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1IntList) Len() int {
	return len(sl.m)
}

func (sl sortableX1IntList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *X1IntList) SortBy(less func(i, j int) bool) *X1IntList {

	sort.Sort(sortableX1IntList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *X1IntList) StableSortBy(less func(i, j int) bool) *X1IntList {

	sort.Stable(sortableX1IntList{less, list.m})
	return list
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *X1IntList) Sorted() *X1IntList {
	return list.SortBy(func(a, b int) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *X1IntList) StableSorted() *X1IntList {
	return list.StableSortBy(func(a, b int) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *X1IntList) Min() int {

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
func (list *X1IntList) Max() (result int) {

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
func (list X1IntList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *X1IntList) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list X1IntList) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *X1IntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *X1IntList) MkString3(before, between, after string) string {
	return list.mkString3Bytes(before, between, after).String()
}

func (list X1IntList) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

