// A queue or fifo that holds int, implemented via a ring buffer.
// Thread-safe.
//
// Generated from threadsafe/queue.tpl with Type=int
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> ToList:<no value>
// by runtemplate v2.2.6
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"sync"
)

// IntQueue is a ring buffer containing a slice of type int. It is optimised
// for FIFO operations.
type IntQueue struct {
	buffer    []int
	read      int
	write     int
	length    int
	cap       int
	overwrite bool
	s         *sync.RWMutex
}

// NewIntQueue returns a new queue of int. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func NewIntQueue(capacity int, overwrite bool) *IntQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &IntQueue{
		buffer:    make([]int, capacity),
		read:      0,
		write:     0,
		length:    0,
		cap:       capacity,
		overwrite: overwrite,
		s:         &sync.RWMutex{},
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *IntQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *IntQueue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue IntQueue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue *IntQueue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *IntQueue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue *IntQueue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.cap
}

// Space returns the space available in the queue.
func (queue *IntQueue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.cap - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *IntQueue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *IntQueue) Len() int {
	if queue == nil {
		return 0
	}
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue *IntQueue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.cap
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *IntQueue) frontAndBack() ([]int, []int) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.buffer[queue.read:queue.write], nil
	}
	return queue.buffer[queue.read:], queue.buffer[:queue.write]
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *IntQueue) ToSlice() []int {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]int, queue.length))
}

func (queue *IntQueue) toSlice(s []int) []int {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *IntQueue) ToInterfaceSlice() []interface{} {
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
func (queue *IntQueue) Clone() *IntQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]int, queue.cap))

	return &IntQueue{
		buffer:    buffer,
		read:      0,
		write:     queue.length,
		length:    queue.length,
		cap:       queue.cap,
		overwrite: queue.overwrite,
		s:         &sync.RWMutex{},
	}
}

// Reallocate adjusts the allocated capacity of the queue and allows the overwriting behaviour to be changed.
// If the new queue capacity is less than the old capacity, the oldest items in the queue are discarded so
// that the remaining data can fit in the space available.
//
// If the new queue capacity is the same as the old capacity, the queue is not altered except for adopting
// the new overwrite flag's value. Therefore this is the means to change the overwriting behaviour.
//
// Reallocate adjusts the storage space but does not clone the underlying elements.
func (queue *IntQueue) Reallocate(capacity int, overwrite bool) *IntQueue {
	if queue == nil {
		return NewIntQueue(capacity, overwrite)
	}

	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.s.Lock()
	defer queue.s.Unlock()

	queue.overwrite = overwrite

	if capacity < queue.length {
		// existing data is too big and has to be trimmed to fit
		n := queue.length - capacity
		queue.read = (queue.read + n) % queue.cap
		queue.length -= n
	}

	if capacity != queue.cap {
		oldLength := queue.length
		queue.buffer = queue.toSlice(make([]int, capacity))
		if oldLength > len(queue.buffer) {
			oldLength = len(queue.buffer)
		}
		queue.read = 0
		queue.write = oldLength
		queue.length = oldLength
		queue.cap = capacity
	}

	return queue
}

//-------------------------------------------------------------------------------------------------

// Push appends items to the end of the queue.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
func (queue *IntQueue) Push(items ...int) []int {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *IntQueue) doPush(items ...int) []int {
	n := len(items)

	space := queue.cap - queue.length
	overwritten := n - space

	if queue.overwrite {
		space = queue.cap
	}

	if space < n {
		// there is too little space; reject surplus elements
		surplus := items[space:]
		queue.doPush(items[:space]...)
		return surplus
	}

	if n <= queue.cap-queue.write {
		// easy case: enough space at end for all items
		copy(queue.buffer[queue.write:], items)
		queue.write = (queue.write + n) % queue.cap
		queue.length += n
		return nil
	}

	// not yet full
	end := queue.cap - queue.write
	copy(queue.buffer[queue.write:], items[:end])
	copy(queue.buffer, items[end:])
	queue.write = n - end
	queue.length += n
	if queue.length > queue.cap {
		queue.length = queue.cap
	}
	if overwritten > 0 {
		queue.read = (queue.read + overwritten) % queue.cap
	}
	return nil
}

// Pop1 removes and returns the oldest item from the queue. If the queue is
// empty, it returns the zero value instead.
// The boolean is true only if the element was available.
func (queue *IntQueue) Pop1() (int, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return 0, false
	}

	v := queue.buffer[queue.read]
	queue.read = (queue.read + 1) % queue.cap
	queue.length--

	return v, true
}

// Pop removes and returns the oldest items from the queue. If the queue is
// empty, it returns a nil slice. If n is larger than the current queue length,
// it returns all the available elements, so in this case the returned slice
// will be shorter than n.
func (queue *IntQueue) Pop(n int) []int {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *IntQueue) doPop(n int) []int {
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

	queue.read = (queue.read + n) % queue.cap
	queue.length -= n

	return s
}

//-------------------------------------------------------------------------------------------------

// HeadOption returns the oldest item in the queue without removing it. If the queue
// is empty, it returns the zero value instead.
func (queue *IntQueue) HeadOption() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return 0
	}

	return queue.buffer[queue.read]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is empty, it returns the zero value instead.
func (queue *IntQueue) LastOption() int {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return 0
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.cap - 1
	}

	return queue.buffer[i]
}
