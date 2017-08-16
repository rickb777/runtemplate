// An encapsulated []string.
// Thread-safe.
//
// Generated from immutable/list.tpl with Type=string
// options: Comparable:true Numeric:false Ordered:false Stringer:true Mutable:disabled

package immutable

import (

	"bytes"
	"fmt"
	"math/rand"
    "sort"
)

// X1StringList contains a slice of type string. Use it where you would use []string.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1StringList struct {
	m []string
}


//-------------------------------------------------------------------------------------------------

func newX1StringList(len, cap int) *X1StringList {
	return &X1StringList {
		m: make([]string, len, cap),
	}
}

// NewX1StringList constructs a new list containing the supplied values, if any.
func NewX1StringList(values ...string) *X1StringList {
	result := newX1StringList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertX1StringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1StringList(values ...interface{}) (*X1StringList, bool) {
	result := newX1StringList(0, len(values))
	good := true

	for _, i := range values {
		v, ok := i.(string)
		if !ok {
		    good = false
		} else {
	    	result.m = append(result.m, v)
	    }
	}

	return result, good
}

// BuildX1StringListFromChan constructs a new X1StringList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1StringListFromChan(source <-chan string) *X1StringList {
	result := newX1StringList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *X1StringList) ToSlice() []string {

	s := make([]string, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *X1StringList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *X1StringList) Clone() *X1StringList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *X1StringList) Get(i int) string {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *X1StringList) Head() string {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *X1StringList) Last() string {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *X1StringList) Tail() *X1StringList {

	result := newX1StringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *X1StringList) Init() *X1StringList {

	result := newX1StringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1StringList is empty.
func (list *X1StringList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1StringList is empty.
func (list *X1StringList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *X1StringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *X1StringList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *X1StringList) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *X1StringList) Len() int {

	return len(list.m)
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *X1StringList) Contains(v string) bool {
	return list.Exists(func (x string) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *X1StringList) ContainsAll(i ...string) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1StringList return true for the passed func.
func (list *X1StringList) Exists(fn func(string) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1StringList return true for the passed func.
func (list *X1StringList) Forall(fn func(string) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1StringList and executes the passed func against each element.
func (list *X1StringList) Foreach(fn func(string)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *X1StringList) Send() <-chan string {
	ch := make(chan string)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X1StringList with all elements in the reverse order.
func (list *X1StringList) Reverse() *X1StringList {

	numItems := len(list.m)
	result := newX1StringList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X1StringList, using a version of the Fisher-Yates shuffle.
func (list *X1StringList) Shuffle() *X1StringList {
	result := NewX1StringList(list.m...)
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *X1StringList) Append(more ...string) *X1StringList {
	newList := NewX1StringList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *X1StringList) doAppend(more ...string) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1StringList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1StringList) Take(n int) *X1StringList {

	if n > len(list.m) {
		return list
	}
	result := newX1StringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1StringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1StringList) Drop(n int) *X1StringList {
	if n == 0 {
		return list
	}


	result := newX1StringList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of X1StringList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1StringList) TakeLast(n int) *X1StringList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newX1StringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1StringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1StringList) DropLast(n int) *X1StringList {
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

// TakeWhile returns a new X1StringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *X1StringList) TakeWhile(p func(string) bool) *X1StringList {

	result := newX1StringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1StringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *X1StringList) DropWhile(p func(string) bool) *X1StringList {

	result := newX1StringList(0, 0)
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

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (list X1StringList) Find(fn func(string) bool) (string, bool) {

	for _, v := range list.m {
		if fn(v) {
			return v, true
		}
	}


    var empty string
	return empty, false

}

// Filter returns a new X1StringList whose elements return true for func.
func (list *X1StringList) Filter(fn func(string) bool) *X1StringList {

	result := newX1StringList(0, len(list.m)/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *X1StringList) Partition(p func(string) bool) (*X1StringList, *X1StringList) {

	matching := newX1StringList(0, len(list.m)/2)
	others := newX1StringList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of X1StringList that return true for the passed predicate.
func (list *X1StringList) CountBy(predicate func(string) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1StringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *X1StringList) MinBy(less func(string, string) bool) string {

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

// MaxBy returns an element of X1StringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *X1StringList) MaxBy(less func(string, string) bool) string {

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

// DistinctBy returns a new X1StringList whose elements are unique, where equality is defined by a passed func.
func (list *X1StringList) DistinctBy(equal func(string, string) bool) *X1StringList {

	result := newX1StringList(0, len(list.m))
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
func (list *X1StringList) IndexWhere(p func(string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *X1StringList) IndexWhere2(p func(string) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *X1StringList) LastIndexWhere(p func(string) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *X1StringList) LastIndexWhere2(p func(string) bool, before int) int {

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
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *X1StringList) Equals(other *X1StringList) bool {

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

type sortableX1StringList struct {
    less func(i, j string) bool
    m []string
}

func (sl sortableX1StringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1StringList) Len() int {
	return len(sl.m)
}

func (sl sortableX1StringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *X1StringList) SortBy(less func(i, j string) bool) *X1StringList {

	result := NewX1StringList(list.m...)
    sort.Sort(sortableX1StringList{less, result.m})
    return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *X1StringList) StableSortBy(less func(i, j string) bool) *X1StringList {

	result := NewX1StringList(list.m...)
    sort.Stable(sortableX1StringList{less, result.m})
    return result
}


//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list X1StringList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *X1StringList) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list X1StringList) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *X1StringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *X1StringList) MkString3(pfx, mid, sfx string) string {
	return list.mkString3Bytes(pfx, mid, sfx).String()
}

func (list X1StringList) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

