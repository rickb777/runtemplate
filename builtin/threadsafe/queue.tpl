// A queue or fifo that holds {{.Type}}, implemented via a ring buffer. Unlike the list collections, these
// have a fixed size (although this can be changed when needed). For mutable collection that need frequent
// appending, the fixed size is a benefit because the memory footprint is constrained. However, this is
// not usable unless the rate of removing items from the queue is, over time, the same as the rate of addition.
// For similar reasons, there is no immutable variant of a queue.
//
// The queue provides a method to sort its elements.
//
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Sorted:{{.Sorted}} Stringer:{{.Stringer}} ToList:{{.ToList}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if or .Stringer .GobEncode}}
	"bytes"
{{- end}}
//{{- if .GobEncode}}
//	"encoding/gob"
//{{- end}}
{{- if .Stringer}}
//	"encoding/json"
	"fmt"
{{- end}}
	"sort"
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.UPrefix}}{{.UType}}Queue is a ring buffer containing a slice of type {{.PType}}. It is optimised
// for FIFO operations.
type {{.UPrefix}}{{.UType}}Queue struct {
	m         []{{.PType}}
	read      int
	write     int
	length    int
	capacity  int
	overwrite bool
	less      func(i, j {{.PType}}) bool
	s         *sync.RWMutex
}

// New{{.UPrefix}}{{.UType}}Queue returns a new queue of {{.PType}}. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
// If the 'less' comparison function is not nil, elements can be easily sorted.
func New{{.UPrefix}}{{.UType}}Queue(capacity int, overwrite bool, less func(i, j {{.PType}}) bool) *{{.UPrefix}}{{.UType}}Queue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &{{.UPrefix}}{{.UType}}Queue{
		m:    make([]{{.PType}}, capacity),
		read:      0,
		write:     0,
		length:    0,
		capacity:  capacity,
		overwrite: overwrite,
		less:      less,
		s:         &sync.RWMutex{},
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *{{.UPrefix}}{{.UType}}Queue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *{{.UPrefix}}{{.UType}}Queue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue {{.UPrefix}}{{.UType}}Queue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue *{{.UPrefix}}{{.UType}}Queue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *{{.UPrefix}}{{.UType}}Queue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue *{{.UPrefix}}{{.UType}}Queue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.capacity
}

// Space returns the space available in the queue.
func (queue *{{.UPrefix}}{{.UType}}Queue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.capacity - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *{{.UPrefix}}{{.UType}}Queue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *{{.UPrefix}}{{.UType}}Queue) Len() int {
	if queue == nil {
		return 0
	}
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue *{{.UPrefix}}{{.UType}}Queue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.capacity
}

// Less reports whether the element with index i should sort before the element with index j.
// The queue must have been created with a non-nil 'less' comparison function and it must not
// be empty.
func (queue *{{.UPrefix}}{{.UType}}Queue) Less(i, j int) bool {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	return queue.less(queue.m[ri], queue.m[rj])
}

// Swap swaps the elements with indexes i and j.
// The queue must not be empty.
func (queue *{{.UPrefix}}{{.UType}}Queue) Swap(i, j int) {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	queue.m[ri], queue.m[rj] = queue.m[rj], queue.m[ri]
}

// Sort sorts the queue using the 'less' comparison function, which must not be nil.
func (queue *{{.UPrefix}}{{.UType}}Queue) Sort() {
	sort.Sort(queue)
}

// StableSort sorts the queue using the 'less' comparison function, which must not be nil.
// The result is stable so that repeated calls will not arbtrarily swap equal items.
func (queue *{{.UPrefix}}{{.UType}}Queue) StableSort() {
	sort.Stable(queue)
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *{{.UPrefix}}{{.UType}}Queue) frontAndBack() ([]{{.PType}}, []{{.PType}}) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.m[queue.read:queue.write], nil
	}
	return queue.m[queue.read:], queue.m[:queue.write]
}
{{- if .ToList}}

// ToList returns the elements of the queue as a list. The returned list is a shallow
// copy; the queue is not altered.
func (queue *{{.UPrefix}}{{.UType}}Queue) ToList() *{{.UPrefix}}{{.UType}}List {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	list := Make{{.UPrefix}}{{.UType}}List(queue.length, queue.length)
	queue.toSlice(list.m)
	return list
}
{{- end}}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *{{.UPrefix}}{{.UType}}Queue) ToSlice() []{{.PType}} {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]{{.PType}}, queue.length))
}

func (queue *{{.UPrefix}}{{.UType}}Queue) toSlice(s []{{.PType}}) []{{.PType}} {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *{{.UPrefix}}{{.UType}}Queue) ToInterfaceSlice() []interface{} {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	var s []interface{}
	for _, v := range front {
		s = append(s, v)
	}

	for _, v := range back {
		s = append(s, v)
	}

	return s
}

// Clone returns a shallow copy of the queue. It does not clone the underlying elements.
func (queue *{{.UPrefix}}{{.UType}}Queue) Clone() *{{.UPrefix}}{{.UType}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]{{.PType}}, queue.capacity))

	return &{{.UPrefix}}{{.UType}}Queue{
		m:    buffer,
		read:      0,
		write:     queue.length,
		length:    queue.length,
		capacity:  queue.capacity,
		overwrite: queue.overwrite,
		s:         &sync.RWMutex{},
	}
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the queue.
// Panics if the index is out of range or the queue is nil.
func (queue *{{.UPrefix}}{{.UType}}Queue) Get(i int) {{.PType}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	ri := (queue.read + i) % queue.capacity
	return queue.m[ri]
}

// Head gets the first element in the queue. Head is the opposite of Last.
// Panics if queue is empty or nil.
func (queue *{{.UPrefix}}{{.UType}}Queue) Head() {{.PType}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.m[queue.read]
}


// HeadOption returns the oldest item in the queue without removing it. If the queue
// is nil or empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.UPrefix}}{{.UType}}Queue) HeadOption() {{.PType}} {
	if queue == nil {
		return {{.TypeZero}}
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.TypeZero}}
	}

	return queue.m[queue.read]
}

// Last gets the the newest item in the queue (i.e. last element pushed) without removing it.
// Last is the opposite of Head.
// Panics if queue is empty or nil.
func (queue *{{.UPrefix}}{{.UType}}Queue) Last() {{.PType}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is nil empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.UPrefix}}{{.UType}}Queue) LastOption() {{.PType}} {
	if queue == nil {
		return {{.TypeZero}}
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.TypeZero}}
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

//-------------------------------------------------------------------------------------------------

// Reallocate adjusts the allocated capacity of the queue and allows the overwriting behaviour to be changed.
// If the new queue capacity is less than the old capacity, the oldest items in the queue are discarded so
// that the remaining data can fit in the space available.
//
// If the new queue capacity is the same as the old capacity, the queue is not altered except for adopting
// the new overwrite flag's value. Therefore this is the means to change the overwriting behaviour.
//
// Reallocate adjusts the storage space but does not clone the underlying elements.
//
// The queue must not be nil.
func (queue *{{.UPrefix}}{{.UType}}Queue) Reallocate(capacity int, overwrite bool) *{{.UPrefix}}{{.UType}}Queue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.s.Lock()
	defer queue.s.Unlock()

	queue.overwrite = overwrite

	if capacity < queue.length {
		// existing data is too big and has to be trimmed to fit
		n := queue.length - capacity
		queue.read = (queue.read + n) % queue.capacity
		queue.length -= n
	}

	if capacity != queue.capacity {
		oldLength := queue.length
		queue.m = queue.toSlice(make([]{{.PType}}, capacity))
		if oldLength > len(queue.m) {
			oldLength = len(queue.m)
		}
		queue.read = 0
		queue.write = oldLength
		queue.length = oldLength
		queue.capacity = capacity
	}

	return queue
}

//-------------------------------------------------------------------------------------------------

// Insert adds items to the queue in sorted order.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
//func (queue *{{.UPrefix}}{{.UType}}Queue) Insert(items ...{{.PType}}) []{{.PType}} {
//	queue.s.Lock()
//	defer queue.s.Unlock()
//	return queue.doInsert(items...)
//}
//
//func (queue *{{.UPrefix}}{{.UType}}Queue) doInsert(items ...{{.PType}}) []{{.PType}} {
//	n := len(items)
//
//	space := queue.capacity - queue.length
//	overwritten := n - space
//
//	if queue.overwrite {
//		space = queue.capacity
//	}
//
//	if space < n {
//		// there is too little space; reject surplus elements
//		surplus := items[space:]
//		queue.doInsert(items[:space]...)
//		return surplus
//	}
//
//  for _, item := range items {
//      queue.doInsertItem(item)
//  }
//	return nil
//}
//
//func (queue *{{.UPrefix}}{{.UType}}Queue) doInsertOne(item {{.PType}}) {
//	if queue.write < queue.capacity {
//		// easy case: enough space at end for the item
//		queue.m[queue.write] = item
//		queue.write = (queue.write + n) % queue.capacity
//		queue.length++
//		return
//	}
//
//	end := queue.capacity - queue.write
//	queue.m[queue.write] = items
//	//copy(queue.m, items[end:])
//	queue.write = n - end
//	queue.length++
//	if queue.length > queue.capacity {
//		queue.length = queue.capacity
//	}
//	if overwritten > 0 {
//		queue.read = (queue.read + overwritten) % queue.capacity
//	}
//}

// Push appends items to the end of the queue.
// This panics if the queue does not have enough space.
func (queue *{{.UPrefix}}{{.UType}}Queue) Push(items ...{{.PType}}) *{{.UPrefix}}{{.UType}}Queue {
	queue.s.Lock()
	defer queue.s.Unlock()

	overflow := queue.doPush(items...)
	if len(overflow) > 0 {
	    panic(len(overflow))
	}
	return queue
}

// Offer appends as many items to the end of the queue as it can.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
func (queue *{{.UPrefix}}{{.UType}}Queue) Offer(items ...{{.PType}}) []{{.PType}} {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *{{.UPrefix}}{{.UType}}Queue) doPush(items ...{{.PType}}) []{{.PType}} {
	n := len(items)

	space := queue.capacity - queue.length
	overwritten := n - space

	if queue.overwrite {
		space = queue.capacity
	}

	if space < n {
		// there is too little space; reject surplus elements
		surplus := items[space:]
		queue.doPush(items[:space]...)
		return surplus
	}

	if n <= queue.capacity - queue.write {
		// easy case: enough space at end for all items
		copy(queue.m[queue.write:], items)
		queue.write = (queue.write + n) % queue.capacity
		queue.length += n
		return nil
	}

	// not yet full
	end := queue.capacity - queue.write
	copy(queue.m[queue.write:], items[:end])
	copy(queue.m, items[end:])
	queue.write = n - end
	queue.length += n
	if queue.length > queue.capacity {
		queue.length = queue.capacity
	}
	if overwritten > 0 {
		queue.read = (queue.read + overwritten) % queue.capacity
	}
	return nil
}

// Pop1 removes and returns the oldest item from the queue. If the queue is
// empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
// The boolean is true only if the element was available.
func (queue *{{.UPrefix}}{{.UType}}Queue) Pop1() ({{.PType}}, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return {{.TypeZero}}, false
	}

	v := queue.m[queue.read]
	queue.read = (queue.read + 1) % queue.capacity
	queue.length--

	return v, true
}

// Pop removes and returns the oldest items from the queue. If the queue is
// empty, it returns a nil slice. If n is larger than the current queue length,
// it returns all the available elements, so in this case the returned slice
// will be shorter than n.
func (queue *{{.UPrefix}}{{.UType}}Queue) Pop(n int) []{{.PType}} {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *{{.UPrefix}}{{.UType}}Queue) doPop(n int) []{{.PType}} {
	if queue.length == 0 {
		return nil
	}

	if n > queue.length {
		n = queue.length
	}

	s := make([]{{.PType}}, n)
	front, back := queue.frontAndBack()
	// note the length copied is whichever is shorter
	copy(s, front)
	if n > len(front) {
		copy(s[len(front):], back)
	}

	queue.read = (queue.read + n) % queue.capacity
	queue.length -= n

	return s
}

//-------------------------------------------------------------------------------------------------
{{- if .Comparable}}

// Contains determines if a given item is already in the queue.
func (queue *{{.UPrefix}}{{.UType}}Queue) Contains(v {{.Type}}) bool {
	return queue.Exists(func(x {{.PType}}) bool {
		return {{.TypeStar}}x == v
	})
}

// ContainsAll determines if the given items are all in the queue.
// This is potentially a slow method and should only be used rarely.
func (queue *{{.UPrefix}}{{.UType}}Queue) ContainsAll(i ...{{.Type}}) bool {
	if queue == nil {
		return len(i) == 0
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	for _, v := range i {
		if !queue.Contains(v) {
			return false
		}
	}
	return true
}
{{- end}}

// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}Queue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *{{.UPrefix}}{{.UType}}Queue) Exists(p func({{.PType}}) bool) bool {
	if queue == nil {
		return false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			return true
		}
	}
	for _, v := range back {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.UPrefix}}{{.UType}}Queue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *{{.UPrefix}}{{.UType}}Queue) Forall(p func({{.PType}}) bool) bool {
	if queue == nil {
		return true
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		if !p(v) {
			return false
		}
	}
	for _, v := range back {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.UPrefix}}{{.UType}}Queue and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (queue *{{.UPrefix}}{{.UType}}Queue) Foreach(fn func({{.PType}})) {
	if queue == nil {
		return
	}

	queue.s.Lock()
	defer queue.s.Unlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		fn(v)
	}
	for _, v := range back {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type}} that returns true for predicate p.
// False is returned if none match.
func (queue *{{.UPrefix}}{{.UType}}Queue) Find(p func({{.PType}}) bool) ({{.PType}}, bool) {
	if queue == nil {
		return {{.TypeZero}}, false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			return v, true
		}
	}
	for _, v := range back {
		if p(v) {
			return v, true
		}
	}
{{if eq .TypeStar "*"}}

	return nil, false
{{else}}

	var empty {{.Type}}
	return empty, false
{{end -}}
}

// Filter returns a new {{.UPrefix}}{{.UType}}Queue whose elements return true for predicate p.
//
// The original queue is not modified. See also DoKeepWhere (which does modify the original queue).
func (queue *{{.UPrefix}}{{.UType}}Queue) Filter(p func({{.PType}}) bool) *{{.UPrefix}}{{.UType}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := New{{.UPrefix}}{{.UType}}Queue(queue.length, queue.overwrite, queue.less)
	i := 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			result.m[i] = v
			i++
		}
	}
	for _, v := range back {
		if p(v) {
			result.m[i] = v
			i++
		}
	}
	result.length = i
	result.write = i

	return result
}

// Partition returns two new {{.Type}}Queues whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original queue.
//
// The original queue is not modified
func (queue *{{.UPrefix}}{{.UType}}Queue) Partition(p func({{.PType}}) bool) (*{{.UPrefix}}{{.UType}}Queue, *{{.UPrefix}}{{.UType}}Queue) {
	if queue == nil {
		return nil, nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	matching := New{{.UPrefix}}{{.UType}}Queue(queue.length, queue.overwrite, queue.less)
	others := New{{.UPrefix}}{{.UType}}Queue(queue.length, queue.overwrite, queue.less)
	m, o := 0, 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			matching.m[m] = v
			m++
		} else {
			others.m[o] = v
			o++
		}
	}
	for _, v := range back {
		if p(v) {
			matching.m[m] = v
			m++
		} else {
			others.m[o] = v
			o++
		}
	}
	matching.length = m
	matching.write = m
	others.length = o
	others.write = o

	return matching, others
}

// Map returns a new {{.UPrefix}}{{.UType}}Queue by transforming every element with a function fn.
// The resulting queue is the same size as the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *{{.UPrefix}}{{.UType}}Queue) Map(fn func({{.PType}}) {{.PType}}) *{{.UPrefix}}{{.UType}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := New{{.UPrefix}}{{.UType}}Queue(queue.length, queue.overwrite, queue.less)
	i := 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		result.m[i] = fn(v)
		i++
	}
	for _, v := range back {
		result.m[i] = fn(v)
		i++
	}
	result.length = i
	result.write = i

	return result
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Queue that return true for the passed predicate.
//func (queue *{{.UPrefix}}{{.UType}}Queue) CountBy(predicate func({{.PType}}) bool) (result int) {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
//
//	front, back := queue.frontAndBack()
//	for _, v := range front {
//		if predicate(v) {
//			result++
//		}
//	}
//	for _, v := range back {
//		if predicate(v) {
//			result++
//		}
//	}
//	return
//}

{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the queue.
//func (queue *{{.UPrefix}}{{.UType}}Queue) Sum() {{.Type}} {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
//
//	sum := {{.Type}}(0)
//	front, back := queue.frontAndBack()
//	for _, v := range front {
//		sum = sum + {{.TypeStar}}v
//	}
//	for _, v := range back {
//		sum = sum + {{.TypeStar}}v
//	}
//	return sum
//}
{{- end}}
{{- if .Comparable}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is comparable.

// Equals determines if two queues are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil queues are considered to be empty.
func (queue *{{.UPrefix}}{{.UType}}Queue) Equals(other *{{.UPrefix}}{{.UType}}Queue) bool {
	if queue == nil {
		return other == nil || other.length == 0
	}

	queue.s.RLock()
	other.s.RLock()
	defer queue.s.RUnlock()
	defer other.s.RUnlock()

	if queue.length != other.length {
		return false
	}

	for i := 0; i < queue.length; i++ {
		qi := (queue.read + i) % queue.capacity
		oi := (other.read + i) % queue.capacity
		if queue.m[qi] != other.m[oi] {
			return false
		}
	}

	return true
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
//func (queue {{.UPrefix}}{{.UType}}Queue) StringList() []string {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
//
//	strings := make([]string, queue.length)
//	i := 0
//	front, back := queue.frontAndBack()
//	for _, v := range front {
//		strings[i] = fmt.Sprintf("%v", v)
//		i++
//	}
//	for _, v := range back {
//		strings[i] = fmt.Sprintf("%v", v)
//		i++
//	}
//	return strings
//}

// String implements the Stringer interface to render the queue as a comma-separated string enclosed in square brackets.
func (queue *{{.UPrefix}}{{.UType}}Queue) String() string {
	return queue.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (queue *{{.UPrefix}}{{.UType}}Queue) MkString(sep string) string {
	return queue.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (queue *{{.UPrefix}}{{.UType}}Queue) MkString3(before, between, after string) string {
	if queue == nil {
		return ""
	}

	return queue.mkString3Bytes(before, between, after).String()
}

func (queue {{.UPrefix}}{{.UType}}Queue) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	for _, v := range back {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

//// UnmarshalJSON implements JSON decoding for this queue type.
//func (queue *{{.UPrefix}}{{.UType}}Queue) UnmarshalJSON(b []byte) error {
//	queue.s.Lock()
//	defer queue.s.Unlock()
//
//	return json.Unmarshal(b, &queue.m)
//}
//
//// MarshalJSON implements JSON encoding for this queue type.
//func (queue {{.UPrefix}}{{.UType}}Queue) MarshalJSON() ([]byte, error) {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
//
//	buf, err := json.Marshal(queue.m)
//	return buf, err
//}
//{{- end}}
//{{- if .GobEncode}}
//
////-------------------------------------------------------------------------------------------------
//
//// GobDecode implements 'gob' decoding for this queue type.
//// You must register {{.Type}} with the 'gob' package before this method is used.
//func (queue *{{.UPrefix}}{{.UType}}Queue) GobDecode(b []byte) error {
//	queue.s.Lock()
//	defer queue.s.Unlock()
//
//	buf := bytes.NewBuffer(b)
//	return gob.NewDecoder(buf).Decode(&queue.m)
//}
//
//// GobDecode implements 'gob' encoding for this queue type.
//// You must register {{.Type}} with the 'gob' package before this method is used.
//func (queue {{.UPrefix}}{{.UType}}Queue) GobEncode() ([]byte, error) {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
//
//	buf := &bytes.Buffer{}
//	err := gob.NewEncoder(buf).Encode(queue.m)
//	return buf.Bytes(), err
//}
//{{- end}}
