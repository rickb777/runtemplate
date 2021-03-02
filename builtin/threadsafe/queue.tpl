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
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Sorted:{{.Sorted}} Stringer:{{.Stringer}}
// ToList:{{.ToList}} ToSet:{{.ToSet}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if .Stringer}}
	"encoding/json"
	"fmt"
{{- end}}
	"sort"
{{- if or .Stringer .GobEncode}}
	"strings"
{{- end}}
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.Prefix.U}}{{.Type.U}}Queue is a ring buffer containing a slice of type {{.Type}}. It is optimised
// for FIFO operations.
type {{.Prefix.U}}{{.Type.U}}Queue struct {
	m         []{{.Type}}
	read      int
	write     int
	length    int
	capacity  int
	overwrite bool
	less      func(i, j {{.Type}}) bool
	s         *sync.RWMutex
}

// New{{.Prefix.U}}{{.Type.U}}Queue returns a new queue of {{.Type}}. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func New{{.Prefix.U}}{{.Type.U}}Queue(capacity int, overwrite bool) *{{.Prefix.U}}{{.Type.U}}Queue {
	return New{{.Prefix.U}}{{.Type.U}}SortedQueue(capacity, overwrite, nil)
}

// New{{.Prefix.U}}{{.Type.U}}SortedQueue returns a new queue of {{.Type}}. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
// If the 'less' comparison function is not nil, elements can be easily sorted.
func New{{.Prefix.U}}{{.Type.U}}SortedQueue(capacity int, overwrite bool, less func(i, j {{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}Queue {
	return &{{.Prefix.U}}{{.Type.U}}Queue{
		m:         make([]{{.Type}}, capacity),
		read:      0,
		write:     0,
		length:    0,
		capacity:  capacity,
		overwrite: overwrite,
		less:      less,
		s:         &sync.RWMutex{},
	}
}

// Build{{.Prefix.U}}{{.Type.U}}QueueFromChan constructs a new {{.Prefix.U}}{{.Type.U}}Queue from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}QueueFromChan(source <-chan {{.Type}}) *{{.Prefix.U}}{{.Type.U}}Queue {
	queue := New{{.Prefix.U}}{{.Type.U}}Queue(0, false)
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Reallocate(capacity int, overwrite bool) *{{.Prefix.U}}{{.Type.U}}Queue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doReallocate(capacity, overwrite)
}

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) doReallocate(capacity int, overwrite bool) *{{.Prefix.U}}{{.Type.U}}Queue {
	queue.overwrite = overwrite

	if capacity < queue.length {
		// existing data is too big and has to be trimmed to fit
		n := queue.length - capacity
		queue.read = (queue.read + n) % queue.capacity
		queue.length -= n
	}

	if capacity != queue.capacity {
		oldLength := queue.length
		queue.m = queue.toSlice(make([]{{.Type}}, capacity))
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.capacity - queue.length
}

// Cap gets the capacity of this queue.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.capacity
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for ordered lists and queues.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) IsSet() bool {
	return false
}

{{- if .ToList}}

// ToList returns the elements of the queue as a list. The returned list is a shallow
// copy; the queue is not altered.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) ToList() *{{.Prefix.U}}{{.Type.U}}List {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	list := Make{{.Prefix.U}}{{.Type.U}}List(queue.length, queue.length)
	queue.toSlice(list.m)
	return list
}
{{- end}}
{{- if .ToSet}}

// ToSet returns the elements of the queue as a set. The returned set is a shallow
// copy; the queue is not altered.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) ToSet() *{{.Prefix.U}}{{.Type.U}}Set {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	slice := queue.toSlice(make([]{{.Type}}, queue.length))
	return New{{.Prefix.U}}{{.Type.U}}Set(slice...)
}
{{- end}}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) ToSlice() []{{.Type}} {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]{{.Type}}, queue.length))
}

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) toSlice(s []{{.Type}}) []{{.Type}} {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) ToInterfaceSlice() []interface{} {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Clone() *{{.Prefix.U}}{{.Type.U}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]{{.Type}}, queue.capacity))
	return queue.doClone(buffer[:queue.length])
}

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) doClone(buffer []{{.Type}}) *{{.Prefix.U}}{{.Type.U}}Queue {
	w := 0
	if len(buffer) < cap(buffer) {
		w = len(buffer)
	}
	return &{{.Prefix.U}}{{.Type.U}}Queue{
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Get(i int) {{.Type}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	ri := (queue.read + i) % queue.capacity
	return queue.m[ri]
}

// Head gets the first element in the queue. Head is the opposite of Last.
// Panics if queue is empty or nil.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Head() {{.Type}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.m[queue.read]
}


// HeadOption returns the oldest item in the queue without removing it. If the queue
// is nil or empty, it returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) HeadOption() ({{.Type}}, bool) {
	if queue == nil {
		return {{.Type.Zero}}, false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.Type.Zero}}, false
	}

	return queue.m[queue.read], true
}

// Last gets the the newest item in the queue (i.e. last element pushed) without removing it.
// Last is the opposite of Head.
// Panics if queue is empty or nil.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Last() {{.Type}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is nil empty, it returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) LastOption() ({{.Type}}, bool) {
	if queue == nil {
		return {{.Type.Zero}}, false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.Type.Zero}}, false
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i], true
}

//-------------------------------------------------------------------------------------------------

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) IsOverwriting() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.overwrite
}

// IsFull returns true if the queue is full.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.capacity
}

// IsEmpty returns true if the queue is empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Len() int {
	return queue.Size()
}

// Swap swaps the elements with indexes i and j.
// The queue must not be empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Swap(i, j int) {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	queue.m[ri], queue.m[rj] = queue.m[rj], queue.m[ri]
}

// Less reports whether the element with index i should sort before the element with index j.
// The queue must have been created with a non-nil 'less' comparison function and it must not
// be empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Less(i, j int) bool {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	return queue.less(queue.m[ri], queue.m[rj])
}

// Sort sorts the queue using the 'less' comparison function, which must not be nil.
// This function will panic if the collection was created with a nil 'less' function
// (see New{{.Prefix.U}}{{.Type.U}}SortedQueue).
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Sort() {
	sort.Sort(queue)
}

// StableSort sorts the queue using the 'less' comparison function, which must not be nil.
// The result is stable so that repeated calls will not arbitrarily swap equal items.
// This function will panic if the collection was created with a nil 'less' function
// (see New{{.Prefix.U}}{{.Type.U}}SortedQueue).
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) StableSort() {
	sort.Stable(queue)
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) frontAndBack() ([]{{.Type}}, []{{.Type}}) {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) indexes() []int {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Clear() {
	if queue != nil {
		queue.s.Lock()
		defer queue.s.Unlock()
		queue.read = 0
		queue.write = 0
		queue.length = 0
	}
}

// Add adds items to the queue. This is a synonym for Push.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Add(more ...{{.Type}}) {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Push(items ...{{.Type}}) *{{.Prefix.U}}{{.Type.U}}Queue {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Offer(items ...{{.Type}}) []{{.Type}} {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) doPush(items ...{{.Type}}) []{{.Type}} {
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
// empty, it returns {{if .Type.IsPtr}}nil{{else}}the zero value{{end}} instead.
// The boolean is true only if the element was available.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Pop1() ({{.Type}}, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return {{.Type.Zero}}, false
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Pop(n int) []{{.Type}} {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) doPop(n int) []{{.Type}} {
	if queue.length == 0 {
		return nil
	}

	if n > queue.length {
		n = queue.length
	}

	s := make([]{{.Type}}, n)
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

// Contains determines whether a given item is already in the queue, returning true if so.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Contains(v {{.Type}}) bool {
	return queue.Exists(func(x {{.Type}}) bool {
		return {{.Type.Star}}x == {{.Type.Star}}v
	})
}

// ContainsAll determines whether the given items are all in the queue, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) ContainsAll(i ...{{.Type}}) bool {
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

// Exists verifies that one or more elements of {{.Prefix.U}}{{.Type.U}}Queue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Exists(p func({{.Type}}) bool) bool {
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

// Forall verifies that all elements of {{.Prefix.U}}{{.Type.U}}Queue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Forall(p func({{.Type}}) bool) bool {
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

// Foreach iterates over {{.Prefix.U}}{{.Type.U}}Queue and executes function f against each element.
// The function can safely alter the values via side-effects.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Foreach(f func({{.Type}})) {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
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

// DoKeepWhere modifies a {{.Prefix.U}}{{.Type.U}}Queue by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the queue in place.
//
// The queue is modified and the modified queue is returned.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) DoKeepWhere(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}Queue {
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

func (queue *{{.Prefix.U}}{{.Type.U}}Queue) doKeepWhere(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}Queue {
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

// Find returns the first {{.Type}} that returns true for predicate p.
// False is returned if none match.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Find(p func({{.Type}}) bool) ({{.Type}}, bool) {
	if queue == nil {
		return {{.Type.Zero}}, false
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
{{- if .TypeIsPtr}}

	return nil, false
{{- else}}

	var empty {{.Type}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.Prefix.U}}{{.Type.U}}Queue whose elements return true for predicate p.
//
// The original queue is not modified. See also DoKeepWhere (which does modify the original queue).
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Filter(p func({{.Type}}) bool) *{{.Prefix.U}}{{.Type.U}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := New{{.Prefix.U}}{{.Type.U}}SortedQueue(queue.length, queue.overwrite, queue.less)
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

// Partition returns two new {{.Prefix.U}}{{.Type.U}}Queues whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original queue.
//
// The original queue is not modified
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Partition(p func({{.Type}}) bool) (*{{.Prefix.U}}{{.Type.U}}Queue, *{{.Prefix.U}}{{.Type.U}}Queue) {
	if queue == nil {
		return nil, nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	matching := New{{.Prefix.U}}{{.Type.U}}SortedQueue(queue.length, queue.overwrite, queue.less)
	others := New{{.Prefix.U}}{{.Type.U}}SortedQueue(queue.length, queue.overwrite, queue.less)
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

// Map returns a new {{.Prefix.U}}{{.Type.U}}Queue by transforming every element with function f.
// The resulting queue is the same size as the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Map(f func({{.Type}}) {{.Type}}) *{{.Prefix.U}}{{.Type.U}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	slice := make([]{{.Type}}, queue.length)
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
{{- range .MapTo}}

// MapTo{{.U}} returns a new []{{.}} by transforming every element with function f.
// The resulting slice is the same size as the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *{{$.Prefix.U}}{{$.Type.U}}Queue) MapTo{{.U}}(f func({{$.Type}}) {{.}}) []{{.}} {
	if queue == nil {
		return nil
	}

	result := make([]{{.}}, 0, queue.length)
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
{{- end}}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}Queue by transforming every element with function f that
// returns zero or more items in a slice. The resulting queue may have a different size to the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) FlatMap(f func({{.Type}}) []{{.Type}}) *{{.Prefix.U}}{{.Type.U}}Queue {
	if queue == nil {
		return nil
	}

	slice := make([]{{.Type}}, 0, queue.length)

	front, back := queue.frontAndBack()
	for _, v := range front {
		slice = append(slice, f(v)...)
	}
	for _, v := range back {
		slice = append(slice, f(v)...)
	}

	return queue.doClone(slice)
}
{{- range .MapTo}}

// FlatMapTo{{.U}} returns a new []{{.}} by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the queue.
// The queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *{{$.Prefix.U}}{{$.Type.U}}Queue) FlatMapTo{{.U}}(f func({{$.Type}}) []{{.}}) []{{.}} {
	if queue == nil {
		return nil
	}

	result := make([]{{.}}, 0, 32)
	queue.s.RLock()
	defer queue.s.RUnlock()

	for _, v := range queue.m {
		result = append(result, f(v)...)
	}

	return result
}
{{- end}}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}Queue that return true for the predicate p.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) CountBy(p func({{.Type}}) bool) (result int) {
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

// Fold aggregates all the values in the queue using a supplied function, starting from some initial value.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Fold(initial {{.Type}}, fn func({{.Type}}, {{.Type}}) {{.Type}}) {{.Type}} {
	if queue == nil {
		return initial
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

    m := initial
	front, back := queue.frontAndBack()
	for _, v := range front {
		m = fn(m, v)
	}
	for _, v := range back {
		m = fn(m, v)
	}
	return m
}

// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}Queue containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
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

// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}Queue containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
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
{{- if .Comparable}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is comparable.

// Equals determines if two queues are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil queues are considered to be empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Equals(other *{{.Prefix.U}}{{.Type.U}}Queue) bool {
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
{{- end}}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Min() {{.Type.Name}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the minimum of an empty queue.")
	}

	z := queue.m[queue.read]
	m := {{.Type.Star}}z
	front, back := queue.frontAndBack()
	for _, v := range front {
		if {{.Type.Star}}v < m {
			m = {{.Type.Star}}v
		}
	}
	for _, v := range back {
		if {{.Type.Star}}v < m {
			m = {{.Type.Star}}v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Max() (result {{.Type.Name}}) {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		panic("Cannot determine the maximum of an empty queue.")
	}

	z := queue.m[queue.read]
	m := {{.Type.Star}}z
	front, back := queue.frontAndBack()
	for _, v := range front {
		if {{.Type.Star}}v > m {
			m = {{.Type.Star}}v
		}
	}
	for _, v := range back {
		if {{.Type.Star}}v > m {
			m = {{.Type.Star}}v
		}
	}
	return m
}
{{- end}}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is numeric.

// Sum returns the sum of all the elements in the queue.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) Sum() {{.Type.Name}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	sum := {{.Type.Name}}(0)
	front, back := queue.frontAndBack()
	for _, v := range front {
		sum = sum + {{.Type.Star}}v
	}
	for _, v := range back {
		sum = sum + {{.Type.Star}}v
	}
	return sum
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) StringList() []string {
{{- if eq .Type.String "string"}}
	return queue.ToSlice()
{{- else}}
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
{{- end}}
}

// String implements the Stringer interface to render the queue as a comma-separated string enclosed in square brackets.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) String() string {
	return queue.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) MkString(sep string) string {
	return queue.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) MkString3(before, between, after string) string {
	if queue == nil {
		return ""
	}

	return queue.mkString3Bytes(before, between, after).String()
}

func (queue {{.Prefix.U}}{{.Type.U}}Queue) mkString3Bytes(before, between, after string) *strings.Builder {
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
func (queue *{{.Prefix.U}}{{.Type.U}}Queue) UnmarshalJSON(b []byte) error {
	queue.s.Lock()
	defer queue.s.Unlock()

	return json.Unmarshal(b, &queue.m)
}

// MarshalJSON implements JSON encoding for this queue type.
func (queue {{.Prefix.U}}{{.Type.U}}Queue) MarshalJSON() ([]byte, error) {
	queue.s.RLock()
	defer queue.s.RUnlock()

	buf, err := json.Marshal(queue.m)
	return buf, err
}
{{- end}}
