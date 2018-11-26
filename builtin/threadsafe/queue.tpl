// A queue or fifo that holds {{.Type}}, implemented via a ring buffer.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} ToList:{{.ToList}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.UPrefix}}{{.UType}}List is a slice of type {{.PType}}. Use it where you would use []{{.PType}}.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.UPrefix}}{{.UType}}Queue struct {
	buffer    []{{.PType}}
	read      int
	write     int
	length    int
	cap       int
	overwrite bool
	s         *sync.RWMutex
}

// New{{.UPrefix}}{{.UType}}Queue returns a new queue of {{.PType}}. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func New{{.UPrefix}}{{.UType}}Queue(size int, overwrite bool) *{{.UPrefix}}{{.UType}}Queue {
	if size < 1 {
		panic("size must be at least 1")
	}
	return &{{.UPrefix}}{{.UType}}Queue{
		buffer:    make([]{{.PType}}, size),
		read:      0,
		write:     0,
		length:    0,
		cap:       size,
		overwrite: overwrite,
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
func (queue {{.UPrefix}}{{.UType}}Queue) IsEmpty() bool {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue {{.UPrefix}}{{.UType}}Queue) NonEmpty() bool {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue {{.UPrefix}}{{.UType}}Queue) IsFull() bool {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.cap
}

// Space returns the space available in the queue.
func (queue {{.UPrefix}}{{.UType}}Queue) Space() int {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.cap - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue {{.UPrefix}}{{.UType}}Queue) Size() int {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue {{.UPrefix}}{{.UType}}Queue) Len() int {
    return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue {{.UPrefix}}{{.UType}}Queue) Cap() int {
	return queue.cap
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *{{.UPrefix}}{{.UType}}Queue) frontAndBack() ([]{{.PType}}, []{{.PType}}) {
	if queue == nil || queue.length == 0 {
	    return nil, nil
    }
	if queue.write > queue.read {
	    return queue.buffer[queue.read:queue.write], nil
	}
    return queue.buffer[queue.read:], queue.buffer[:queue.write]
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
    if len(back) > 0 {
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

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (queue *{{.UPrefix}}{{.UType}}Queue) Clone() *{{.UPrefix}}{{.UType}}Queue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]{{.PType}}, queue.cap))

	return &{{.UPrefix}}{{.UType}}Queue{
        buffer:    buffer,
        read:      0,
        write:     queue.length,
        length:    queue.length,
        cap:       queue.cap,
        overwrite: queue.overwrite,
        s:         &sync.RWMutex{},
    }
}

//-------------------------------------------------------------------------------------------------

// Push appends items to the end of the queue.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
func (queue *{{.UPrefix}}{{.UType}}Queue) Push(items ...{{.PType}}) []{{.PType}} {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *{{.UPrefix}}{{.UType}}Queue) doPush(items ...{{.PType}}) []{{.PType}} {
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

    if n <= queue.cap - queue.write {
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
// empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
// The boolean is true only if the element was available.
func (queue *{{.UPrefix}}{{.UType}}Queue) Pop1() ({{.PType}}, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return {{.TypeZero}}, false
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
func (queue *{{.UPrefix}}{{.UType}}Queue) Pop(n int) []{{.PType}} {
	queue.s.Lock()
	defer queue.s.Unlock()

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

	queue.read = (queue.read + n) % queue.cap
	queue.length -= n

	return s
}

// HeadOption returns the oldest item in the queue without removing it. If the queue
// is empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.UPrefix}}{{.UType}}Queue) HeadOption() {{.PType}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.TypeZero}}
	}

	return queue.buffer[queue.read]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.UPrefix}}{{.UType}}Queue) LastOption() {{.PType}} {
	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return {{.TypeZero}}
	}

    i := queue.write - 1
    if i < 0 {
        i = queue.cap - 1
    }

	return queue.buffer[i]
}
