// A queue or fifo that holds int, implemented via a ring buffer. Unlike the list collections, these
// have a fixed size (although this can be changed when needed). For mutable collection that need frequent
// appending, the fixed size is a benefit because the memory footprint is constrained. However, this is
// not usable unless the rate of removing items from the queue is, over time, the same as the rate of addition.
// For similar reasons, there is no immutable variant of a queue.
//
// The queue provides a method to sort its elements.
//
// Thread-safe.
//
// Generated from threadsafe/queue.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Sorted:<no value> Stringer:true
// ToList:true ToSet:true
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

// X1IntQueue is a ring buffer containing a slice of type int. It is optimised
// for FIFO operations.
type X1IntQueue struct {
	m         []int
	read      int
	write     int
	length    int
	capacity  int
	overwrite bool
	less      func(i, j int) bool
	s         *sync.RWMutex
}

// NewX1IntQueue returns a new queue of int. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func NewX1IntQueue(capacity int, overwrite bool) *X1IntQueue {
	return NewX1IntSortedQueue(capacity, overwrite, nil)
}

// NewX1IntSortedQueue returns a new queue of int. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
// If the 'less' comparison function is not nil, elements can be easily sorted.
func NewX1IntSortedQueue(capacity int, overwrite bool, less func(i, j int) bool) *X1IntQueue {
	return &X1IntQueue{
		m:         make([]int, capacity),
		read:      0,
		write:     0,
		length:    0,
		capacity:  capacity,
		overwrite: overwrite,
		less:      less,
		s:         &sync.RWMutex{},
	}
}

// BuildX1IntQueueFromChan constructs a new X1IntQueue from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildX1IntQueueFromChan(source <-chan int) *X1IntQueue {
	queue := NewX1IntQueue(0, false)
	for v := range source {
		queue.m = append(queue.m, v)
	}
	queue.length = len(queue.m)
	queue.capacity = cap(queue.m)
	return queue
}

//-------------------------------------------------------------------------------------------------

// Reallocate adjusts the allocated capacity of the queue and allows the overwriting behaviour to be changed.
//
// If the new queue capacity is different to the current capacity, the queue is re-allocated to the new
// capacity. If this is less than the current number of elements, the oldest items in the queue are
// discarded so that the remaining data can fit in the new space available.
//
// If the new queue capacity is the same as the current capacity, the queue is not altered except for adopting
// the new overwrite flag's value. Therefore this is the means to change the overwriting behaviour.
//
// Reallocate adjusts the storage space but does not clone the underlying elements.
//
// The queue must not be nil.
func (queue *X1IntQueue) Reallocate(capacity int, overwrite bool) *X1IntQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doReallocate(capacity, overwrite)
}

func (queue *X1IntQueue) doReallocate(capacity int, overwrite bool) *X1IntQueue {
	queue.overwrite = overwrite

	if capacity < queue.length {
		// existing data is too big and has to be trimmed to fit
		n := queue.length - capacity
		queue.read = (queue.read + n) % queue.capacity
		queue.length -= n
	}

	if capacity != queue.capacity {
		oldLength := queue.length
		queue.m = queue.toSlice(make([]int, capacity))
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

// Space returns the space available in the queue.
func (queue *X1IntQueue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.capacity - queue.length
}

// Cap gets the capacity of this queue.
func (queue *X1IntQueue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.capacity
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for ordered lists and queues.
func (queue *X1IntQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *X1IntQueue) IsSet() bool {
	return false
}

// ToList returns the elements of the queue as a list. The returned list is a shallow
// copy; the queue is not altered.
func (queue *X1IntQueue) ToList() *X1IntList {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	list := MakeX1IntList(queue.length, queue.length)
	queue.toSlice(list.m)
	return list
}

// ToSet returns the elements of the queue as a set. The returned set is a shallow
// copy; the queue is not altered.
func (queue *X1IntQueue) ToSet() *X1IntSet {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	slice := queue.toSlice(make([]int, queue.length))
	return NewX1IntSet(slice...)
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *X1IntQueue) ToSlice() []int {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]int, queue.length))
}

func (queue *X1IntQueue) toSlice(s []int) []int {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *X1IntQueue) ToInterfaceSlice() []interface{} {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	s := make([]interface{}, 0, queue.length)
	for _, v := range front {
		s = append(s, v)
	}

	for _, v := range back {
		s = append(s, v)
	}

	return s
}

// Clone returns a shallow copy of the queue. It does not clone the underlying elements.
func (queue *X1IntQueue) Clone() *X1IntQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]int, queue.capacity))
	return queue.doClone(buffer[:queue.length])
}

func (queue *X1IntQueue) doClone(buffer []int) *X1IntQueue {
	w := 0
	if len(buffer) < cap(buffer) {
		w = len(buffer)
	}
	return &X1IntQueue{
		m:         buffer,
		read:      0,
		write:     w,
		length:    len(buffer),
		capacity:  cap(buffer),
		overwrite: queue.overwrite,
		less:      queue.less,
		s:         &sync.RWMutex{},
	}
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the queue.
// Panics if the index is out of range or the queue is nil.
func (queue *X1IntQueue) Get(i int) int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	ri := (queue.read + i) % queue.capacity
	return queue.m[ri]
}

// Head gets the first element in the queue. Head is the opposite of Last.
// Panics if queue is empty or nil.
func (queue *X1IntQueue) Head() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.m[queue.read]
}


// HeadOption returns the oldest item in the queue without removing it. If the queue
// is nil or empty, it returns the zero value instead.
func (queue *X1IntQueue) HeadOption() int {
	if queue == nil {
		return 0
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return 0
	}

	return queue.m[queue.read]
}

// Last gets the the newest item in the queue (i.e. last element pushed) without removing it.
// Last is the opposite of Head.
// Panics if queue is empty or nil.
func (queue *X1IntQueue) Last() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is nil empty, it returns the zero value instead.
func (queue *X1IntQueue) LastOption() int {
	if queue == nil {
		return 0
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return 0
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

//-------------------------------------------------------------------------------------------------

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue *X1IntQueue) IsOverwriting() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.overwrite
}

// IsFull returns true if the queue is full.
func (queue *X1IntQueue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.capacity
}

// IsEmpty returns true if the queue is empty.
func (queue *X1IntQueue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *X1IntQueue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *X1IntQueue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *X1IntQueue) Len() int {
	return queue.Size()
}

// Swap swaps the elements with indexes i and j.
// The queue must not be empty.
func (queue *X1IntQueue) Swap(i, j int) {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	queue.m[ri], queue.m[rj] = queue.m[rj], queue.m[ri]
}

// Less reports whether the element with index i should sort before the element with index j.
// The queue must have been created with a non-nil 'less' comparison function and it must not
// be empty.
func (queue *X1IntQueue) Less(i, j int) bool {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	return queue.less(queue.m[ri], queue.m[rj])
}

// Sort sorts the queue using the 'less' comparison function, which must not be nil.
// This function will panic if the collection was created with a nil 'less' function
// (see NewX1IntSortedQueue).
func (queue *X1IntQueue) Sort() {
	sort.Sort(queue)
}

// StableSort sorts the queue using the 'less' comparison function, which must not be nil.
// The result is stable so that repeated calls will not arbitrarily swap equal items.
// This function will panic if the collection was created with a nil 'less' function
// (see NewX1IntSortedQueue).
func (queue *X1IntQueue) StableSort() {
	sort.Stable(queue)
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *X1IntQueue) frontAndBack() ([]int, []int) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.m[queue.read:queue.write], nil
	}
	return queue.m[queue.read:], queue.m[:queue.write]
}

// indexes gets the indexes for the front and back portions of the queue. The front
// portion starts from the read index. The back portion ends at the write index.
func (queue *X1IntQueue) indexes() []int {
	if queue == nil || queue.length == 0 {
		return nil
	}
	if queue.write > queue.read {
		return []int{queue.read, queue.write}
	}
	return []int{queue.read, queue.capacity, 0, queue.write}
}

//-------------------------------------------------------------------------------------------------

// Clear the entire queue.
func (queue *X1IntQueue) Clear() {
	if queue != nil {
		queue.s.Lock()
		defer queue.s.Unlock()
		queue.read = 0
		queue.write = 0
		queue.length = 0
	}
}

// Add adds items to the queue. This is a synonym for Push.
func (queue *X1IntQueue) Add(more ...int) {
	queue.Push(more...)
}

// Push appends items to the end of the queue. If the queue does not have enough space,
// more will be allocated: how this happens depends on the overwriting mode.
//
// When overwriting, the oldest items are overwritten with the new data; it expands the queue
// only if there is still not enough space.
//
// Otherwise, the queue might be reallocated if necessary, ensuring that all the data is pushed
// without any older items being affected.
//
// The modified queue is returned.
func (queue *X1IntQueue) Push(items ...int) *X1IntQueue {
	queue.s.Lock()
	defer queue.s.Unlock()

	n := queue.capacity
	if queue.overwrite && len(items) > queue.capacity {
		n = len(items)
		// no rounding in this case because the old items are expected to be overwritten

	} else if !queue.overwrite && len(items) > (queue.capacity - queue.length) {
		n = len(items) + queue.length
		// rounded up to multiple of 128 to reduce repeated reallocation
		n = ((n + 127) / 128) * 128
	}

	if n > queue.capacity {
		queue = queue.doReallocate(n, queue.overwrite)
	}

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
// The queue capacity is never altered.
func (queue *X1IntQueue) Offer(items ...int) []int {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *X1IntQueue) doPush(items ...int) []int {
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
// empty, it returns the zero value instead.
// The boolean is true only if the element was available.
func (queue *X1IntQueue) Pop1() (int, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return 0, false
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
func (queue *X1IntQueue) Pop(n int) []int {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *X1IntQueue) doPop(n int) []int {
	if queue.length == 0 {
		return nil
	}

	if n > queue.length {
		n = queue.length
	}

	s := make([]int, n)
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

// Contains determines whether a given item is already in the queue, returning true if so.
func (queue *X1IntQueue) Contains(v int) bool {
	return queue.Exists(func(x int) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the queue, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (queue *X1IntQueue) ContainsAll(i ...int) bool {
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

// Exists verifies that one or more elements of X1IntQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *X1IntQueue) Exists(p func(int) bool) bool {
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

// Forall verifies that all elements of X1IntQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *X1IntQueue) Forall(p func(int) bool) bool {
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

// Foreach iterates over X1IntQueue and executes function f against each element.
// The function can safely alter the values via side-effects.
func (queue *X1IntQueue) Foreach(f func(int)) {
	if queue == nil {
		return
	}

	queue.s.Lock()
	defer queue.s.Unlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		f(v)
	}
	for _, v := range back {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (queue *X1IntQueue) Send() <-chan int {
	ch := make(chan int)
	go func() {
		if queue != nil {
			queue.s.RLock()
			defer queue.s.RUnlock()

			front, back := queue.frontAndBack()
			for _, v := range front {
				ch <- v
			}
			for _, v := range back {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a X1IntQueue by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the queue in place.
//
// The queue is modified and the modified queue is returned.
func (queue *X1IntQueue) DoKeepWhere(p func(int) bool) *X1IntQueue {
	if queue == nil {
		return nil
	}

	queue.s.Lock()
	defer queue.s.Unlock()
	if queue.length == 0 {
		return queue
	}

	return queue.doKeepWhere(p)
}

func (queue *X1IntQueue) doKeepWhere(p func(int) bool) *X1IntQueue {
	last := queue.capacity

	if queue.write > queue.read {
		// only need to process the front of the queue
		last = queue.write
	}

	r := queue.read
	w := r
	n := 0

	// 1st loop: front of queue (from queue.read)
	for r < last {
		if p(queue.m[r]) {
			if w != r {
				queue.m[w] = queue.m[r]
			}
			w++
			n++
		}
		r++
	}

	w = w % queue.capacity

	if queue.write > queue.read {
		// only needed to process the front of the queue
		queue.write = w
		queue.length = n
		return queue
	}

	// 2nd loop: back of queue (from 0 to queue.write)
	r = 0
	for r < queue.write {
		if p(queue.m[r]) {
			if w != r {
				queue.m[w] = queue.m[r]
			}
			w = (w + 1) % queue.capacity
			n++
		}
		r++
	}

	queue.write = w
	queue.length = n

	return queue
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for predicate p.
// False is returned if none match.
func (queue *X1IntQueue) Find(p func(int) bool) (int, bool) {
	if queue == nil {
		return 0, false
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

	var empty int
	return empty, false
}

// Filter returns a new X1IntQueue whose elements return true for predicate p.
//
// The original queue is not modified. See also DoKeepWhere (which does modify the original queue).
func (queue *X1IntQueue) Filter(p func(int) bool) *X1IntQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := NewX1IntSortedQueue(queue.length, queue.overwrite, queue.less)
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

// Partition returns two new X1IntQueues whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original queue.
//
// The original queue is not modified
func (queue *X1IntQueue) Partition(p func(int) bool) (*X1IntQueue, *X1IntQueue) {
	if queue == nil {
		return nil, nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	matching := NewX1IntSortedQueue(queue.length, queue.overwrite, queue.less)
	others := NewX1IntSortedQueue(queue.length, queue.overwrite, queue.less)
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

// Map returns a new X1IntQueue by transforming every element with function f.
// The resulting queue is the same size as the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) Map(f func(int) int) *X1IntQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	slice := make([]int, queue.length)
	i := 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		slice[i] = f(v)
		i++
	}
	for _, v := range back {
		slice[i] = f(v)
		i++
	}

	return queue.doClone(slice)
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) MapToString(f func(int) string) []string {
	if queue == nil {
		return nil
	}

	result := make([]string, 0, queue.length)
	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		result = append(result, f(v))
	}
	for _, v := range back {
		result = append(result, f(v))
	}

	return result
}

// MapToInt64 returns a new []int64 by transforming every element with function f.
// The resulting slice is the same size as the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) MapToInt64(f func(int) int64) []int64 {
	if queue == nil {
		return nil
	}

	result := make([]int64, 0, queue.length)
	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		result = append(result, f(v))
	}
	for _, v := range back {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new X1IntQueue by transforming every element with function f that
// returns zero or more items in a slice. The resulting queue may have a different size to the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) FlatMap(f func(int) []int) *X1IntQueue {
	if queue == nil {
		return nil
	}

	slice := make([]int, 0, queue.length)

	front, back := queue.frontAndBack()
	for _, v := range front {
		slice = append(slice, f(v)...)
	}
	for _, v := range back {
		slice = append(slice, f(v)...)
	}

	return queue.doClone(slice)
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) FlatMapToString(f func(int) []string) []string {
	if queue == nil {
		return nil
	}

	result := make([]string, 0, 32)
	queue.s.RLock()
	defer queue.s.RUnlock()

	for _, v := range queue.m {
		result = append(result, f(v)...)
	}

	return result
}

// FlatMapToInt64 returns a new []int64 by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *X1IntQueue) FlatMapToInt64(f func(int) []int64) []int64 {
	if queue == nil {
		return nil
	}

	result := make([]int64, 0, 32)
	queue.s.RLock()
	defer queue.s.RUnlock()

	for _, v := range queue.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of X1IntQueue that return true for the predicate p.
func (queue *X1IntQueue) CountBy(p func(int) bool) (result int) {
	if queue == nil {
		return 0
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			result++
		}
	}
	for _, v := range back {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X1IntQueue containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (queue *X1IntQueue) MinBy(less func(int, int) bool) int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the minimum of an empty queue.")
	}

	indexes := queue.indexes()
	m := indexes[0]
	for len(indexes) > 1 {
		f := indexes[0]
		for i := f; i < indexes[1]; i++ {
			if i != m {
				if less(queue.m[i], queue.m[m]) {
					m = i
				}
			}
		}
		indexes = indexes[2:]
	}
	return queue.m[m]
}

// MaxBy returns an element of X1IntQueue containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (queue *X1IntQueue) MaxBy(less func(int, int) bool) int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the maximum of an empty queue.")
	}

	indexes := queue.indexes()
	m := indexes[0]
	for len(indexes) > 1 {
		f := indexes[0]
		for i := f; i < indexes[1]; i++ {
			if i != m {
				if less(queue.m[m], queue.m[i]) {
					m = i
				}
			}
		}
		indexes = indexes[2:]
	}
	return queue.m[m]
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is comparable.

// Equals determines if two queues are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil queues are considered to be empty.
func (queue *X1IntQueue) Equals(other *X1IntQueue) bool {
	if queue == nil {
		if other == nil {
			return true
		}
		other.s.RLock()
		defer other.s.RUnlock()
		return other.length == 0
	}

	if other == nil {
		queue.s.RLock()
		defer queue.s.RUnlock()
		return queue.length == 0
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

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (queue *X1IntQueue) Min() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the minimum of an empty queue.")
	}

	z := queue.m[queue.read]
	m := z
	front, back := queue.frontAndBack()
	for _, v := range front {
		if v < m {
			m = v
		}
	}
	for _, v := range back {
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (queue *X1IntQueue) Max() (result int) {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the maximum of an empty queue.")
	}

	z := queue.m[queue.read]
	m := z
	front, back := queue.frontAndBack()
	for _, v := range front {
		if v > m {
			m = v
		}
	}
	for _, v := range back {
		if v > m {
			m = v
		}
	}
	return m
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the queue.
func (queue *X1IntQueue) Sum() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	sum := int(0)
	front, back := queue.frontAndBack()
	for _, v := range front {
		sum = sum + v
	}
	for _, v := range back {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (queue *X1IntQueue) StringList() []string {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	strings := make([]string, queue.length)
	i := 0
	front, back := queue.frontAndBack()
	for _, v := range front {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	for _, v := range back {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the queue as a comma-separated string enclosed in square brackets.
func (queue *X1IntQueue) String() string {
	return queue.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (queue *X1IntQueue) MkString(sep string) string {
	return queue.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (queue *X1IntQueue) MkString3(before, between, after string) string {
	if queue == nil {
		return ""
	}

	return queue.mkString3Bytes(before, between, after).String()
}

func (queue X1IntQueue) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
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

// UnmarshalJSON implements JSON decoding for this queue type.
func (queue *X1IntQueue) UnmarshalJSON(b []byte) error {
	queue.s.Lock()
	defer queue.s.Unlock()

	return json.Unmarshal(b, &queue.m)
}

// MarshalJSON implements JSON encoding for this queue type.
func (queue X1IntQueue) MarshalJSON() ([]byte, error) {
	queue.s.RLock()
	defer queue.s.RUnlock()

	buf, err := json.Marshal(queue.m)
	return buf, err
}
