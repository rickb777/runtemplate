// An encapsulated []int32.
// Thread-safe.
//
// Generated from list.tpl with Type=*int32
// options: Comparable=true Numeric=true Ordered=true Stringer=true Mutable=true

package fast

import (

	"bytes"
	"fmt"
"math/rand"
)

// PInt32List contains a slice of type *int32. Use it where you would use []*int32.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type PInt32List struct {
	m []*int32
}


//-------------------------------------------------------------------------------------------------

func newPInt32List(len, cap int) *PInt32List {
	return &PInt32List {
		m: make([]*int32, len, cap),
	}
}

// NewPInt32List constructs a new list containing the supplied values, if any.
func NewPInt32List(values ...*int32) *PInt32List {
	result := newPInt32List(len(values), len(values))
	for i, v := range values {
		result.m[i] = v
	}
	return result
}

// BuildPInt32ListFromChan constructs a new PInt32List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildPInt32ListFromChan(source <-chan *int32) *PInt32List {
	result := newPInt32List(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (list *PInt32List) ToSlice() []*int32 {

	var s []*int32
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *PInt32List) Clone() *PInt32List {
	return NewPInt32List(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *PInt32List) Get(i int) *int32 {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *PInt32List) Head() *int32 {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *PInt32List) Last() *int32 {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *PInt32List) Tail() *PInt32List {

	result := newPInt32List(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *PInt32List) Init() *PInt32List {

	result := newPInt32List(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether PInt32List is empty.
func (list *PInt32List) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether PInt32List is empty.
func (list *PInt32List) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list *PInt32List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *PInt32List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *PInt32List) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *PInt32List) Len() int {

	return len(list.m)
}


// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *PInt32List) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *PInt32List) Contains(v int32) bool {
	return list.Exists(func (x *int32) bool {
	    return *x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *PInt32List) ContainsAll(i ...int32) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of PInt32List return true for the passed func.
func (list *PInt32List) Exists(fn func(*int32) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of PInt32List return true for the passed func.
func (list *PInt32List) Forall(fn func(*int32) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over PInt32List and executes the passed func against each element.
func (list *PInt32List) Foreach(fn func(*int32)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *PInt32List) Send() <-chan *int32 {
	ch := make(chan *int32)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of PInt32List with all elements in the reverse order.
func (list *PInt32List) Reverse() *PInt32List {

	numItems := list.Len()
	result := newPInt32List(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of PInt32List, using a version of the Fisher-Yates shuffle.
func (list *PInt32List) Shuffle() *PInt32List {
	numItems := list.Len()
	result := list.Clone()
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}


// Add adds items to the current list. This is a synonym for Append.
func (list *PInt32List) Add(more ...*int32) {
    list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *PInt32List) Append(more ...*int32) *PInt32List {

	for _, v := range more {
		list.doAppend(v)
	}
	return list
}

func (list *PInt32List) doAppend(i *int32) {
	list.m = append(list.m, i)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of PInt32List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *PInt32List) Take(n int) *PInt32List {

	if n > list.Len() {
		return list
	}
	result := newPInt32List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of PInt32List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *PInt32List) Drop(n int) *PInt32List {
	if n == 0 {
		return list
	}


	result := newPInt32List(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of PInt32List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *PInt32List) TakeLast(n int) *PInt32List {

	l := list.Len()
	if n > l {
		return list
	}
	result := newPInt32List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of PInt32List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *PInt32List) DropLast(n int) *PInt32List {
	if n == 0 {
		return list
	}


	l := list.Len()
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new PInt32List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *PInt32List) TakeWhile(p func(*int32) bool) *PInt32List {

	result := newPInt32List(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new PInt32List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *PInt32List) DropWhile(p func(*int32) bool) *PInt32List {

	result := newPInt32List(0, 0)
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

// Filter returns a new PInt32List whose elements return true for func.
func (list *PInt32List) Filter(fn func(*int32) bool) *PInt32List {

	result := newPInt32List(0, list.Len()/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new int32Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *PInt32List) Partition(p func(*int32) bool) (*PInt32List, *PInt32List) {

	matching := newPInt32List(0, list.Len()/2)
	others := newPInt32List(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of PInt32List that return true for the passed predicate.
func (list *PInt32List) CountBy(predicate func(*int32) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is ordered.

// Less returns true if the element at index i is less than the element at index j. This implements
// one of the methods needed by sort.Interface.
// Panics if i or j is out of range.
func (list *PInt32List) Less(i, j int) bool {
	return *list.m[i] < *list.m[j]
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *PInt32List) Min() int32 {

	l := list.Len()
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := *v
	for i := 1; i < l; i++ {
    	v := list.m[i]
		if *v < m {
			m = *v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *PInt32List) Max() (result int32) {

    l := list.Len()
    if l == 0 {
        panic("Cannot determine the maximum of an empty list.")
    }

    v := list.m[0]
    m := *v
    for i := 1; i < l; i++ {
        v := list.m[i]
        if *v > m {
            m = *v
        }
    }
    return m
}

// DistinctBy returns a new PInt32List whose elements are unique, where equality is defined by a passed func.
func (list *PInt32List) DistinctBy(equal func(*int32, *int32) bool) *PInt32List {

	result := newPInt32List(0, list.Len())
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
func (list *PInt32List) IndexWhere(p func(*int32) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *PInt32List) IndexWhere2(p func(*int32) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *PInt32List) LastIndexWhere(p func(*int32) bool) int {
	return list.LastIndexWhere2(p, 0)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *PInt32List) LastIndexWhere2(p func(*int32) bool, before int) int {

	for i := list.Len() - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is numeric.

// Sum returns the sum of all the elements in the list.
func (list *PInt32List) Sum() int32 {

	sum := int32(0)
	for _, v := range list.m {
		sum = sum + *v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int32 is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *PInt32List) Equals(other *PInt32List) bool {

	if list.Size() != other.Size() {
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

// StringList gets a list of strings that depicts all the elements.
func (list PInt32List) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *PInt32List) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list PInt32List) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *PInt32List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *PInt32List) MkString3(pfx, mid, sfx string) string {
	return list.mkString3Bytes(pfx, mid, sfx).String()
}

func (list PInt32List) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
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

