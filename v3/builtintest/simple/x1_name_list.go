// A simple type derived from []Name
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=Name
// options: Comparable:true Numeric:<no value> Ordered:true StringLike:true Stringer:true
// GobEncode:<no value> Mutable:always ToList:always ToSet:true MapTo:<no value>
// by runtemplate v3.5.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package simple

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

// X1NameList is a slice of type Name. Use it where you would use []Name.
// To add items to the list, simply use the normal built-in append function.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1NameList []Name

//-------------------------------------------------------------------------------------------------

// MakeX1NameList makes an empty list with both length and capacity initialised.
func MakeX1NameList(length, capacity int) X1NameList {
	return make(X1NameList, length, capacity)
}

// NewX1NameList constructs a new list containing the supplied values, if any.
func NewX1NameList(values ...Name) X1NameList {
	list := MakeX1NameList(len(values), len(values))
	copy(list, values)
	return list
}

// ConvertX1NameList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertX1NameList(values ...interface{}) (X1NameList, bool) {
	list := MakeX1NameList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case Name:
			list = append(list, j)
		case *Name:
			list = append(list, *j)
		default:
			if s, ok := i.(fmt.Stringer); ok {
				list = append(list, Name(s.String()))
			}
		}
	}

	return list, len(list) == len(values)
}

// BuildX1NameListFromChan constructs a new X1NameList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1NameListFromChan(source <-chan Name) X1NameList {
	list := MakeX1NameList(0, 0)
	for v := range source {
		list = append(list, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list X1NameList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list X1NameList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list X1NameList) slice() []Name {
	return list
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list X1NameList) ToList() X1NameList {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list X1NameList) ToSet() X1NameSet {
	if list == nil {
		return nil
	}

	return NewX1NameSet(list...)
}

// ToSlice returns the elements of the list as a slice, which is an identity operation in this case,
// because the simple list is merely a dressed-up slice.
func (list X1NameList) ToSlice() []Name {
	return list
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list X1NameList) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(list))
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list X1NameList) Clone() X1NameList {
	return NewX1NameList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list X1NameList) Get(i int) Name {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list X1NameList) Head() Name {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list X1NameList) HeadOption() Name {
	if list.IsEmpty() {
		return *(new(Name))
	}
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list X1NameList) Last() Name {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list X1NameList) LastOption() Name {
	if list.IsEmpty() {
		return *(new(Name))
	}
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list X1NameList) Tail() X1NameList {
	return list[1:]
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list X1NameList) Init() X1NameList {
	return list[:len(list)-1]
}

// IsEmpty tests whether X1NameList is empty.
func (list X1NameList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether X1NameList is empty.
func (list X1NameList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list X1NameList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list X1NameList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list X1NameList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list X1NameList) Contains(v Name) bool {
	return list.Exists(func(x Name) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list X1NameList) ContainsAll(i ...Name) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1NameList return true for the predicate p.
func (list X1NameList) Exists(p func(Name) bool) bool {
	for _, v := range list {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1NameList return true for the predicate p.
func (list X1NameList) Forall(p func(Name) bool) bool {
	for _, v := range list {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1NameList and executes function f against each element.
func (list X1NameList) Foreach(f func(Name)) {
	for _, v := range list {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list X1NameList) Send() <-chan Name {
	ch := make(chan Name)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of X1NameList with all elements in the reverse order.
//
// The original list is not modified.
func (list X1NameList) Reverse() X1NameList {
	n := len(list)
	result := MakeX1NameList(n, n)
	last := n - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse alters a X1NameList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list X1NameList) DoReverse() X1NameList {
	mid := (len(list) + 1) / 2
	last := len(list) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list[i], list[r] = list[r], list[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of X1NameList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list X1NameList) Shuffle() X1NameList {
	if list == nil {
		return nil
	}

	return list.Clone().DoShuffle()
}

// DoShuffle returns a shuffled X1NameList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list X1NameList) DoShuffle() X1NameList {
	if list == nil {
		return nil
	}

	n := len(list)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list[i], list[r] = list[r], list[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1NameList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list X1NameList) Take(n int) X1NameList {
	if n >= len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of X1NameList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list X1NameList) Drop(n int) X1NameList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of X1NameList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list X1NameList) TakeLast(n int) X1NameList {
	l := len(list)
	if n >= l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of X1NameList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list X1NameList) DropLast(n int) X1NameList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	}
	return list[0 : l-n]
}

// TakeWhile returns a new X1NameList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list X1NameList) TakeWhile(p func(Name) bool) X1NameList {
	result := MakeX1NameList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1NameList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list X1NameList) DropWhile(p func(Name) bool) X1NameList {
	result := MakeX1NameList(0, 0)
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

// Find returns the first Name that returns true for predicate p.
// False is returned if none match.
func (list X1NameList) Find(p func(Name) bool) (Name, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}

	var empty Name
	return empty, false
}

// Filter returns a new X1NameList whose elements return true for predicate p.
//
// The original list is not modified.
func (list X1NameList) Filter(p func(Name) bool) X1NameList {
	result := MakeX1NameList(0, len(list))

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new NameLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list X1NameList) Partition(p func(Name) bool) (X1NameList, X1NameList) {
	matching := MakeX1NameList(0, len(list))
	others := MakeX1NameList(0, len(list))

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new X1NameList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list X1NameList) Map(f func(Name) Name) X1NameList {
	result := MakeX1NameList(0, len(list))

	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new X1NameList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list X1NameList) FlatMap(f func(Name) []Name) X1NameList {
	result := MakeX1NameList(0, len(list))

	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X1NameList that return true for the predicate p.
func (list X1NameList) CountBy(p func(Name) bool) (result int) {
	for _, v := range list {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1NameList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list X1NameList) MinBy(less func(Name, Name) bool) Name {
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

// MaxBy returns an element of X1NameList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list X1NameList) MaxBy(less func(Name, Name) bool) Name {
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

// DistinctBy returns a new X1NameList whose elements are unique, where equality is defined by the equal function.
func (list X1NameList) DistinctBy(equal func(Name, Name) bool) X1NameList {
	result := MakeX1NameList(0, len(list))
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

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list X1NameList) IndexWhere(p func(Name) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list X1NameList) IndexWhere2(p func(Name) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list X1NameList) LastIndexWhere(p func(Name) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list X1NameList) LastIndexWhere2(p func(Name) bool, before int) int {
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
// These methods are included when Name is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list X1NameList) Equals(other X1NameList) bool {
	if list == nil {
		return len(other) == 0
	}

	if other == nil {
		return len(list) == 0
	}

	if len(list) != len(other) {
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

type sortableX1NameList struct {
	less func(i, j Name) bool
	m    []Name
}

func (sl sortableX1NameList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableX1NameList) Len() int {
	return len(sl.m)
}

func (sl sortableX1NameList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list X1NameList) SortBy(less func(i, j Name) bool) X1NameList {
	sort.Sort(sortableX1NameList{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list X1NameList) StableSortBy(less func(i, j Name) bool) X1NameList {
	sort.Stable(sortableX1NameList{less, list})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when Name is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list X1NameList) Sorted() X1NameList {
	return list.SortBy(func(a, b Name) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list X1NameList) StableSorted() X1NameList {
	return list.StableSortBy(func(a, b Name) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list X1NameList) Min() Name {
	m := list.MinBy(func(a Name, b Name) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list X1NameList) Max() (result Name) {
	m := list.MaxBy(func(a Name, b Name) bool {
		return a < b
	})
	return m
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list X1NameList) StringList() []string {
	strings := make([]string, len(list))
	for i, v := range list {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list X1NameList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list X1NameList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list X1NameList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list X1NameList) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	for _, v := range list {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------
