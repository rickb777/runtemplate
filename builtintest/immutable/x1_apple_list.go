// An encapsulated []Apple.
// Thread-safe.
//
// Generated from immutable/list.tpl with Type=Apple
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:false Mutable:disabled

package immutable

import (

	"math/rand"
    "sort"
)

// X1AppleList contains a slice of type Apple. Use it where you would use []Apple.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1AppleList struct {
	m []Apple
}


//-------------------------------------------------------------------------------------------------

func newX1AppleList(len, cap int) *X1AppleList {
	return &X1AppleList {
		m: make([]Apple, len, cap),
	}
}

// NewX1AppleList constructs a new list containing the supplied values, if any.
func NewX1AppleList(values ...Apple) *X1AppleList {
	result := newX1AppleList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertX1AppleList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1AppleList(values ...interface{}) (*X1AppleList, bool) {
	result := newX1AppleList(0, len(values))

	for _, i := range values {
		v, ok := i.(Apple)
		if ok {
	    	result.m = append(result.m, v)
	    }
	}

	return result, len(result.m) == len(values)
}

// BuildX1AppleListFromChan constructs a new X1AppleList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1AppleListFromChan(source <-chan Apple) *X1AppleList {
	result := newX1AppleList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *X1AppleList) ToSlice() []Apple {

	s := make([]Apple, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1AppleList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *X1AppleList) Clone() *X1AppleList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *X1AppleList) Get(i int) Apple {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *X1AppleList) Head() Apple {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *X1AppleList) Last() Apple {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *X1AppleList) Tail() *X1AppleList {

	result := newX1AppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *X1AppleList) Init() *X1AppleList {

	result := newX1AppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1AppleList is empty.
func (list *X1AppleList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1AppleList is empty.
func (list *X1AppleList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *X1AppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *X1AppleList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *X1AppleList) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *X1AppleList) Len() int {

	return len(list.m)
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *X1AppleList) Contains(v Apple) bool {
	return list.Exists(func (x Apple) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *X1AppleList) ContainsAll(i ...Apple) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1AppleList return true for the passed func.
func (list *X1AppleList) Exists(fn func(Apple) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1AppleList return true for the passed func.
func (list *X1AppleList) Forall(fn func(Apple) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1AppleList and executes the passed func against each element.
func (list *X1AppleList) Foreach(fn func(Apple)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *X1AppleList) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X1AppleList with all elements in the reverse order.
func (list *X1AppleList) Reverse() *X1AppleList {

	numItems := len(list.m)
	result := newX1AppleList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X1AppleList, using a version of the Fisher-Yates shuffle.
func (list *X1AppleList) Shuffle() *X1AppleList {
	result := NewX1AppleList(list.m...)
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *X1AppleList) Append(more ...Apple) *X1AppleList {
	newList := NewX1AppleList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *X1AppleList) doAppend(more ...Apple) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1AppleList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1AppleList) Take(n int) *X1AppleList {

	if n > len(list.m) {
		return list
	}
	result := newX1AppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1AppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1AppleList) Drop(n int) *X1AppleList {
	if n == 0 {
		return list
	}


	result := newX1AppleList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of X1AppleList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1AppleList) TakeLast(n int) *X1AppleList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newX1AppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1AppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1AppleList) DropLast(n int) *X1AppleList {
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

// TakeWhile returns a new X1AppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *X1AppleList) TakeWhile(p func(Apple) bool) *X1AppleList {

	result := newX1AppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1AppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *X1AppleList) DropWhile(p func(Apple) bool) *X1AppleList {

	result := newX1AppleList(0, 0)
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

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (list X1AppleList) Find(fn func(Apple) bool) (Apple, bool) {

	for _, v := range list.m {
		if fn(v) {
			return v, true
		}
	}


    var empty Apple
	return empty, false

}

// Filter returns a new X1AppleList whose elements return true for func.
func (list *X1AppleList) Filter(fn func(Apple) bool) *X1AppleList {

	result := newX1AppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *X1AppleList) Partition(p func(Apple) bool) (*X1AppleList, *X1AppleList) {

	matching := newX1AppleList(0, len(list.m)/2)
	others := newX1AppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new X1AppleList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1AppleList) Map(fn func(Apple) Apple) *X1AppleList {
	result := newX1AppleList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new X1AppleList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *X1AppleList) FlatMap(fn func(Apple) []Apple) *X1AppleList {
	result := newX1AppleList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of X1AppleList that return true for the passed predicate.
func (list *X1AppleList) CountBy(predicate func(Apple) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1AppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1AppleList) MinBy(less func(Apple, Apple) bool) Apple {

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

// MaxBy returns an element of X1AppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1AppleList) MaxBy(less func(Apple, Apple) bool) Apple {

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

// DistinctBy returns a new X1AppleList whose elements are unique, where equality is defined by a passed func.
func (list *X1AppleList) DistinctBy(equal func(Apple, Apple) bool) *X1AppleList {

	result := newX1AppleList(0, len(list.m))
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
func (list *X1AppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *X1AppleList) IndexWhere2(p func(Apple) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *X1AppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *X1AppleList) LastIndexWhere2(p func(Apple) bool, before int) int {

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
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *X1AppleList) Equals(other *X1AppleList) bool {

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

type sortableX1AppleList struct {
    less func(i, j Apple) bool
    m []Apple
}

func (sl sortableX1AppleList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1AppleList) Len() int {
	return len(sl.m)
}

func (sl sortableX1AppleList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *X1AppleList) SortBy(less func(i, j Apple) bool) *X1AppleList {

	result := NewX1AppleList(list.m...)
    sort.Sort(sortableX1AppleList{less, result.m})
    return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *X1AppleList) StableSortBy(less func(i, j Apple) bool) *X1AppleList {

	result := NewX1AppleList(list.m...)
    sort.Stable(sortableX1AppleList{less, result.m})
    return result
}


