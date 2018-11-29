// A queue or fifo that holds Apple, implemented via a ring buffer.
// Thread-safe.
//
// Generated from threadsafe/queue.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> ToList:<no value>
// by runtemplate v2.2.7
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	"sync"
)

// AppleQueue is a ring buffer containing a slice of type Apple. It is optimised
// for FIFO operations.
type AppleQueue struct {
	buffer    []Apple
	read      int
	write     int
	length    int
	cap       int
	overwrite bool
	s         *sync.RWMutex
}

// NewAppleQueue returns a new queue of Apple. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func NewAppleQueue(capacity int, overwrite bool) *AppleQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &AppleQueue{
		buffer:    make([]Apple, capacity),
		read:      0,
		write:     0,
		length:    0,
		cap:       capacity,
		overwrite: overwrite,
		s:         &sync.RWMutex{},
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *AppleQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *AppleQueue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue AppleQueue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue *AppleQueue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *AppleQueue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue *AppleQueue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.cap
}

// Space returns the space available in the queue.
func (queue *AppleQueue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.cap - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *AppleQueue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *AppleQueue) Len() int {
	if queue == nil {
		return 0
	}
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue *AppleQueue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.cap
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *AppleQueue) frontAndBack() ([]Apple, []Apple) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.buffer[queue.read:queue.write], nil
	}
	return queue.buffer[queue.read:], queue.buffer[:queue.write]
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *AppleQueue) ToSlice() []Apple {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]Apple, queue.length))
}

func (queue *AppleQueue) toSlice(s []Apple) []Apple {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *AppleQueue) ToInterfaceSlice() []interface{} {
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
func (queue *AppleQueue) Clone() *AppleQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]Apple, queue.cap))

	return &AppleQueue{
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
func (queue *AppleQueue) Reallocate(capacity int, overwrite bool) *AppleQueue {
	if queue == nil {
		return NewAppleQueue(capacity, overwrite)
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
		queue.buffer = queue.toSlice(make([]Apple, capacity))
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
func (queue *AppleQueue) Push(items ...Apple) []Apple {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *AppleQueue) doPush(items ...Apple) []Apple {
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
func (queue *AppleQueue) Pop1() (Apple, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return *(new(Apple)), false
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
func (queue *AppleQueue) Pop(n int) []Apple {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *AppleQueue) doPop(n int) []Apple {
	if queue.length == 0 {
		return nil
	}

	if n > queue.length {
		n = queue.length
	}

	s := make([]Apple, n)
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
func (queue *AppleQueue) HeadOption() Apple {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return *(new(Apple))
	}

	return queue.buffer[queue.read]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is empty, it returns the zero value instead.
func (queue *AppleQueue) LastOption() Apple {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return *(new(Apple))
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.cap - 1
	}

	return queue.buffer[i]
}
