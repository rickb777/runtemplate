// A queue or fifo that holds Apple, implemented via a ring buffer.
// Not thread-safe.
//
// Generated from fast/queue.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> ToList:<no value>
// by runtemplate v2.2.3
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import ()

// FastAppleQueue is a ring buffer containing a slice of type Apple. It is optimised
// for FIFO operations.
type FastAppleQueue struct {
	buffer    []Apple
	read      int
	write     int
	length    int
	cap       int
	overwrite bool
}

// NewFastAppleQueue returns a new queue of Apple. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func NewFastAppleQueue(size int, overwrite bool) *FastAppleQueue {
	if size < 1 {
		panic("size must be at least 1")
	}
	return &FastAppleQueue{
		buffer:    make([]Apple, size),
		read:      0,
		write:     0,
		length:    0,
		cap:       size,
		overwrite: overwrite,
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *FastAppleQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *FastAppleQueue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue FastAppleQueue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue FastAppleQueue) IsEmpty() bool {
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue FastAppleQueue) NonEmpty() bool {
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue FastAppleQueue) IsFull() bool {
	return queue.length == queue.cap
}

// Space returns the space available in the queue.
func (queue FastAppleQueue) Space() int {
	return queue.cap - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue FastAppleQueue) Size() int {
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue FastAppleQueue) Len() int {
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue FastAppleQueue) Cap() int {
	return queue.cap
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *FastAppleQueue) frontAndBack() ([]Apple, []Apple) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.buffer[queue.read:queue.write], nil
	}
	return queue.buffer[queue.read:], queue.buffer[:queue.write]
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *FastAppleQueue) ToSlice() []Apple {
	if queue == nil {
		return nil
	}

	return queue.toSlice(make([]Apple, queue.length))
}

func (queue *FastAppleQueue) toSlice(s []Apple) []Apple {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *FastAppleQueue) ToInterfaceSlice() []interface{} {
	if queue == nil {
		return nil
	}

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
func (queue *FastAppleQueue) Clone() *FastAppleQueue {
	if queue == nil {
		return nil
	}

	buffer := queue.toSlice(make([]Apple, queue.cap))

	return &FastAppleQueue{
		buffer:    buffer,
		read:      0,
		write:     queue.length,
		length:    queue.length,
		cap:       queue.cap,
		overwrite: queue.overwrite,
	}
}

// Resize adjusts the allocated capacity of the queue and allows the overwriting behaviour to be changed.
// It does not clone the underlying elements.
//func (queue *FastAppleQueue) Resize(newSize int, overwrite bool) *FastAppleQueue {
//	if queue == nil {
//		return NewFastAppleQueue(newSize, overwrite)
//	}
//
//
//	queue.overwrite = overwrite
//
//	if newSize != queue.cap {
//		queue.buffer = queue.toSlice(make([]Apple, newSize))
//		queue.read = 0
//		queue.write = len(queue.buffer)
//		queue.length = len(queue.buffer)
//		queue.cap = newSize
//	}
//
//	return queue
//}

//-------------------------------------------------------------------------------------------------

// Push appends items to the end of the queue.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
func (queue *FastAppleQueue) Push(items ...Apple) []Apple {
	return queue.doPush(items...)
}

func (queue *FastAppleQueue) doPush(items ...Apple) []Apple {
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
func (queue *FastAppleQueue) Pop1() (Apple, bool) {

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
func (queue *FastAppleQueue) Pop(n int) []Apple {

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

// HeadOption returns the oldest item in the queue without removing it. If the queue
// is empty, it returns the zero value instead.
func (queue *FastAppleQueue) HeadOption() Apple {

	if queue.length == 0 {
		return *(new(Apple))
	}

	return queue.buffer[queue.read]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is empty, it returns the zero value instead.
func (queue *FastAppleQueue) LastOption() Apple {

	if queue.length == 0 {
		return *(new(Apple))
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.cap - 1
	}

	return queue.buffer[i]
}
