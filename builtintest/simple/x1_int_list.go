// A simple type derived from []int
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true

package simple

import (
	"math/rand"
	"bytes"
	"fmt"
)

// X1IntList is a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1IntList []int

//-------------------------------------------------------------------------------------------------

func newX1IntList(len, cap int) X1IntList {
	return make(X1IntList, len, cap)
}

// NewX1IntList constructs a new list containing the supplied values, if any.
func NewX1IntList(values ...int) X1IntList {
	result := newX1IntList(len(values), len(values))
	for i, v := range values {
		result[i] = v
	}
	return result
}

// BuildX1IntListFromChan constructs a new X1IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1IntListFromChan(source <-chan int) X1IntList {
	result := newX1IntList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list X1IntList) Clone() X1IntList {
	return NewX1IntList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list X1IntList) Get(i int) int {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list X1IntList) Head() int {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list X1IntList) Last() int {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list X1IntList) Tail() X1IntList {
	return X1IntList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list X1IntList) Init() X1IntList {
	return X1IntList(list[:list.Len()-1])
}

// IsEmpty tests whether X1IntList is empty.
func (list X1IntList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether X1IntList is empty.
func (list X1IntList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list X1IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list X1IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list X1IntList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list X1IntList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list X1IntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list X1IntList) Contains(v int) bool {
	return list.Exists(func (x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list X1IntList) ContainsAll(i ...int) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1IntList return true for the passed func.
func (list X1IntList) Exists(fn func(int) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1IntList return true for the passed func.
func (list X1IntList) Forall(fn func(int) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1IntList and executes the passed func against each element.
func (list X1IntList) Foreach(fn func(int)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list X1IntList) Send() <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X1IntList with all elements in the reverse order.
func (list X1IntList) Reverse() X1IntList {
	numItems := list.Len()
	result := newX1IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X1IntList, using a version of the Fisher-Yates shuffle.
func (list X1IntList) Shuffle() X1IntList {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list X1IntList) Take(n int) X1IntList {
	if n > list.Len() {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of X1IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list X1IntList) Drop(n int) X1IntList {
	if n == 0 {
		return list
	}

	l := list.Len()
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of X1IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list X1IntList) TakeLast(n int) X1IntList {
	l := list.Len()
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of X1IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list X1IntList) DropLast(n int) X1IntList {
	if n == 0 {
		return list
	}

	l := list.Len()
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new X1IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list X1IntList) TakeWhile(p func(int) bool) X1IntList {
	result := newX1IntList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list X1IntList) DropWhile(p func(int) bool) X1IntList {
	result := newX1IntList(0, 0)
	adding := false

	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (list X1IntList) Find(fn func(int) bool) (int, bool) {

	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}


    var empty int
	return empty, false

}

// Filter returns a new X1IntList whose elements return true for func.
func (list X1IntList) Filter(fn func(int) bool) X1IntList {
	result := newX1IntList(0, list.Len()/2)

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list X1IntList) Partition(p func(int) bool) (X1IntList, X1IntList) {
	matching := newX1IntList(0, list.Len()/2)
	others := newX1IntList(0, list.Len()/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of X1IntList that return true for the passed predicate.
func (list X1IntList) CountBy(predicate func(int) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1IntList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list X1IntList) MinBy(less func(int, int) bool) int {
	l := list.Len()
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

// MaxBy returns an element of X1IntList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list X1IntList) MaxBy(less func(int, int) bool) int {
	l := list.Len()
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

// DistinctBy returns a new X1IntList whose elements are unique, where equality is defined by a passed func.
func (list X1IntList) DistinctBy(equal func(int, int) bool) X1IntList {
	result := newX1IntList(0, list.Len())
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
func (list X1IntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list X1IntList) IndexWhere2(p func(int) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list X1IntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, list.Len())
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list X1IntList) LastIndexWhere2(p func(int) bool, before int) int {
	if before < 0 {
		before = list.Len()
	}
	for i := list.Len() - 1; i >= 0; i-- {
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
func (list X1IntList) Sum() int {
	sum := int(0)
	for _, v := range list {
		sum = sum + v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list X1IntList) Equals(other X1IntList) bool {
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
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list X1IntList) Min() int {
	m := list.MinBy(func(a int, b int) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list X1IntList) Max() (result int) {
	m := list.MaxBy(func(a int, b int) bool {
		return a < b
	})
	return m
}

// Less returns true if the element at index i is less than the element at index j. This implements
// one of the methods needed by sort.Interface.
// Panics if i or j is out of range.
func (list X1IntList) Less(i, j int) bool {
	return list[i] < list[j]
}


//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list X1IntList) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list X1IntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list X1IntList) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := list.Len()
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}
