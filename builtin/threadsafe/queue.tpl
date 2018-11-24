// A queue or fifo that holds {{.Type}}, implemented via a ring buffer.
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}}
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

// Len gets the number of elements currently in this queue.
func (queue {{.UPrefix}}{{.UType}}Queue) Len() int {
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Cap gets the capacity of this queue.
func (queue {{.UPrefix}}{{.UType}}Queue) Cap() int {
	return queue.cap
}

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
    if queue.overwrite {
        space = queue.cap
    }

	if space < n {
        surplus := items[space:]
        queue.doPush(items[:space]...)
    	return surplus
	}

    if n <= queue.cap - queue.write {
		// easy case - enough space at end for all items
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
    return nil
}

// Pop removes and returns the oldest item in the queue. If the queue is
// empty, it returns {{if .TypeIsPtr}}nil{{else}}the zero value{{end}} instead.
func (queue *{{.UPrefix}}{{.UType}}Queue) Pop() {{.PType}} {
	queue.s.Lock()
	defer queue.s.Unlock()

	if queue.length == 0 {
		return {{.TypeZero}}
	}

	element := queue.buffer[queue.read]
	queue.read = (queue.read + 1) % queue.cap
	queue.length--

	return element
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
