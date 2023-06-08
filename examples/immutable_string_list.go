// An encapsulated immutable []string.
// Thread-safe.
//
//
// Generated from immutable/list.tpl with Type=string
// options: Comparable:true Numeric:<no value> Integer:<no value> Ordered:true
//          StringLike:<no value> StringParser:<no value> Stringer:true
// by runtemplate v3.10.2
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// ImmutableStringList contains a slice of type string. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type ImmutableStringList struct {
	m []string
}

//-------------------------------------------------------------------------------------------------

func newImmutableStringList(length, capacity int) *ImmutableStringList {
	return &ImmutableStringList{
		m: make([]string, length, capacity),
	}
}

// NewImmutableStringList constructs a new list containing the supplied values, if any.
func NewImmutableStringList(values ...string) *ImmutableStringList {
	list := newImmutableStringList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertImmutableStringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertImmutableStringList(values ...interface{}) (*ImmutableStringList, bool) {
	list := newImmutableStringList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case string:
			k := string(j)
			list.m = append(list.m, k)
		case *string:
			k := string(*j)
			list.m = append(list.m, k)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildImmutableStringListFromChan constructs a new ImmutableStringList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildImmutableStringListFromChan(source <-chan string) *ImmutableStringList {
	list := newImmutableStringList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *ImmutableStringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *ImmutableStringList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *ImmutableStringList) slice() []string {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *ImmutableStringList) ToList() *ImmutableStringList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *ImmutableStringList) ToSlice() []string {

	s := make([]string, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *ImmutableStringList) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *ImmutableStringList) Clone() *ImmutableStringList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *ImmutableStringList) Get(i int) string {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *ImmutableStringList) Head() string {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *ImmutableStringList) HeadOption() (string, bool) {
	if list == nil || len(list.m) == 0 {
		var v string
		return v, false
	}
	return list.m[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *ImmutableStringList) Last() string {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *ImmutableStringList) LastOption() (string, bool) {
	if list == nil || len(list.m) == 0 {
		var v string
		return v, false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *ImmutableStringList) Tail() *ImmutableStringList {
	result := newImmutableStringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *ImmutableStringList) Init() *ImmutableStringList {
	result := newImmutableStringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether ImmutableStringList is empty.
func (list *ImmutableStringList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether ImmutableStringList is empty.
func (list *ImmutableStringList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *ImmutableStringList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *ImmutableStringList) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *ImmutableStringList) Contains(v string) bool {
	return list.Exists(func(x string) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *ImmutableStringList) ContainsAll(i ...string) bool {
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

// Exists verifies that one or more elements of ImmutableStringList return true for the predicate p.
func (list *ImmutableStringList) Exists(p func(string) bool) bool {
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

// Forall verifies that all elements of ImmutableStringList return true for the predicate p.
func (list *ImmutableStringList) Forall(p func(string) bool) bool {
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

// Foreach iterates over ImmutableStringList and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *ImmutableStringList) Foreach(f func(string)) {
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
func (list *ImmutableStringList) Send() <-chan string {
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

// Reverse returns a copy of ImmutableStringList with all elements in the reverse order.
func (list *ImmutableStringList) Reverse() *ImmutableStringList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := newImmutableStringList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of ImmutableStringList, using a version of the Fisher-Yates shuffle.
func (list *ImmutableStringList) Shuffle() *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := NewImmutableStringList(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *ImmutableStringList) Append(more ...string) *ImmutableStringList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return NewImmutableStringList(more...)
	}

	newList := NewImmutableStringList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *ImmutableStringList) doAppend(more ...string) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of ImmutableStringList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *ImmutableStringList) Take(n int) *ImmutableStringList {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := newImmutableStringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of ImmutableStringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableStringList) Drop(n int) *ImmutableStringList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := newImmutableStringList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of ImmutableStringList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *ImmutableStringList) TakeLast(n int) *ImmutableStringList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := newImmutableStringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of ImmutableStringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableStringList) DropLast(n int) *ImmutableStringList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := newImmutableStringList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new ImmutableStringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *ImmutableStringList) TakeWhile(p func(string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new ImmutableStringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *ImmutableStringList) DropWhile(p func(string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(0, 0)
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
func (list *ImmutableStringList) Find(p func(string) bool) (string, bool) {
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

// Filter returns a new ImmutableStringList whose elements return true for predicate p.
func (list *ImmutableStringList) Filter(p func(string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *ImmutableStringList) Partition(p func(string) bool) (*ImmutableStringList, *ImmutableStringList) {
	if list == nil {
		return nil, nil
	}

	matching := newImmutableStringList(0, len(list.m)/2)
	others := newImmutableStringList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new ImmutableStringList by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableStringList) Map(f func(string) string) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// MapToInt returns a new []int by transforming every element with function f.
// The resulting slice is the same size as the list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableStringList) MapToInt(f func(string) int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, len(list.m))
	for i, v := range list.m {
		result[i] = f(v)
	}

	return result
}

// FlatMap returns a new ImmutableStringList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableStringList) FlatMap(f func(string) []string) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// FlatMapToInt returns a new []int by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableStringList) FlatMapToInt(f func(string) []int) []int {
	if list == nil {
		return nil
	}

	result := make([]int, 0, len(list.m))
	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of ImmutableStringList that return true for the predicate p.
func (list *ImmutableStringList) CountBy(p func(string) bool) (result int) {
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
func (list *ImmutableStringList) Fold(initial string, fn func(string, string) string) string {
	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of ImmutableStringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableStringList) MinBy(less func(string, string) bool) string {
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

// MaxBy returns an element of ImmutableStringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableStringList) MaxBy(less func(string, string) bool) string {
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

// DistinctBy returns a new ImmutableStringList whose elements are unique, where equality is defined by the equal function.
func (list *ImmutableStringList) DistinctBy(equal func(string, string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := newImmutableStringList(0, len(list.m))
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
func (list *ImmutableStringList) IndexWhere(p func(string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *ImmutableStringList) IndexWhere2(p func(string) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *ImmutableStringList) LastIndexWhere(p func(string) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *ImmutableStringList) LastIndexWhere2(p func(string) bool, before int) int {

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
func (list *ImmutableStringList) Equals(other *ImmutableStringList) bool {
	if list == nil {
		return other == nil || len(other.m) == 0
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

type sortableImmutableStringList struct {
	less func(i, j string) bool
	m    []string
}

func (sl sortableImmutableStringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableImmutableStringList) Len() int {
	return len(sl.m)
}

func (sl sortableImmutableStringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *ImmutableStringList) SortBy(less func(i, j string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := NewImmutableStringList(list.m...)
	sort.Sort(sortableImmutableStringList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *ImmutableStringList) StableSortBy(less func(i, j string) bool) *ImmutableStringList {
	if list == nil {
		return nil
	}

	result := NewImmutableStringList(list.m...)
	sort.Stable(sortableImmutableStringList{less, result.m})
	return result
}

//-------------------------------------------------------------------------------------------------
// These methods are included when string is ordered.

// Sorted returns a new list in which the elements are sorted by their natural ordering.
func (list *ImmutableStringList) Sorted() *ImmutableStringList {
	return list.SortBy(func(a, b string) bool {
		return a < b
	})
}

// StableSorted returns a new list in which the elements are sorted by their natural ordering.
func (list *ImmutableStringList) StableSorted() *ImmutableStringList {
	return list.StableSortBy(func(a, b string) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *ImmutableStringList) Min() string {

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
func (list *ImmutableStringList) Max() (result string) {

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
func (list *ImmutableStringList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *ImmutableStringList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *ImmutableStringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *ImmutableStringList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list ImmutableStringList) mkString3Bytes(before, between, after string) *strings.Builder {
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
func (list *ImmutableStringList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list ImmutableStringList) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(list.m)
	return buf, err
}
