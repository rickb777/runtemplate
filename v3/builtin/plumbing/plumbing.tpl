// Generated from {{.TemplateFile}} with Type={{.Type}}
// No other options are available.
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package {{.Package}}

{{if .HasImport}}
import (
	{{.Import}}
)

{{end -}}
//-------------------------------------------------------------------------------------------------

// This plumbing suite provides standard functions for piping data between goroutines.
// All these functions run until the input channel is closed (or all input channels are closed, if
// multiple). They then close their output channel(s).
// None of these functions create a goroutine - this must be done at the call site.

//-------------------------------------------------------------------------------------------------

// {{.Prefix.U}}{{.Type.U}}Generator produces a stream of {{.Type}} based on a supplied generator function.
// The function fn is invoked N times with the integers from 0 to N-1. Each result is sent out.
// Finally, the output channel is closed and the generator terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Generator(out chan<- {{.Type}}, iterations int, fn func(int) {{.Type}}) {
	{{.Prefix.U}}{{.Type.U}}Generator3(out, 0, iterations-1, 1, fn)
}

// {{.Prefix.U}}{{.Type.U}}Generator produces a stream of {{.Type}} based on a supplied generator function.
// The function fn is invoked *(|to - from|) / |stride|* times with the integers in the range specified by
// from, to and stride. If stride is negative, from should be greater than to.
// For each iteration, the computed function result is sent out.
// If stride is zero, the loop never terminates. Otherwise, after the generator has reached the
// loop end, the output channel is closed and the generator terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Generator3(out chan<- {{.Type}}, from, to, stride int, fn func(int) {{.Type}}) {
	if (from > to && stride > 0) || (from < to && stride < 0) {
		panic("Loop conditions are divergent.")
	}
	if from > to && stride < 0 {
		for i := from; i >= to; i += stride {
			out <- fn(i)
		}
	} else {
		for i := from; i <= to; i += stride {
			out <- fn(i)
		}
	}
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}Delta duplicates a stream of {{.Type}} to two output channels.
// When the sender closes the input channel, both output channels are closed then the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Delta(in <-chan {{.Type}}, out1, out2 chan<- {{.Type}}) {
	for v := range in {
		select {
		case out1 <- v:
			out2 <- v
		case out2 <- v:
			out1 <- v
		}
	}
	close(out1)
	close(out2)
}

// {{.Prefix.U}}{{.Type.U}}Zip2 interleaves two streams of {{.Type}}.
// Each input channel is used in turn, alternating between them.
// The function terminates when *both* input channels have been closed by their senders.
// The output channel is then closed also.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Zip2(in1, in2 <-chan {{.Type}}, out chan<- {{.Type}}) {
	closed2 := false
	for v := range in1 {
		out <- v
		v, ok := <-in2
		if ok {
			out <- v
		} else {
			closed2 = true
		}
	}
	// need to drain in2 as well?
	if !closed2 {
		for _ = range in2 {
		}
	}
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}Mux2 multiplexes two streams of {{.Type}} into a single output channel.
// Each input channel is used as soon as it is ready.
// When a signal is received from the closer channel, the output channel is then closed.
// Concurrently, both input channels are then passed into blackholes that comsume them until they too are closed,
// and the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Mux2(in1, in2 <-chan {{.Type}}, closer <-chan bool, out chan<- {{.Type}}) {
	running := true
	for running {
		select {
		case v := <-in1:
			out <- v
		case v := <-in2:
			out <- v
		case _ = <-closer:
			running = false
		}
	}
	go {{.Prefix.U}}{{.Type.U}}BlackHole(in1)
	go {{.Prefix.U}}{{.Type.U}}BlackHole(in2)
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}BlackHole silently consumes a stream of {{.Type}}.
// It terminates when the sender closes the channel.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}BlackHole(in <-chan {{.Type}}) {
	for _ = range in {
		// om nom nom
	}
}

// {{.Prefix.U}}{{.Type.U}}Filter filters a stream of {{.Type}}, silently dropping elements that do not match the predicate p.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Filter(in <-chan {{.Type}}, out chan<- {{.Type}}, p func({{.Type}}) bool) {
	for v := range in {
		if p(v) {
			out <- v
		}
	}
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}Partition filters a stream of {{.Type}} into two output streams using a predicate p, those that
// match and all others.
// When the sender closes the input channel, both output channels are closed then the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Partition(in <-chan {{.Type}}, matching, others chan<- {{.Type}}, p func({{.Type}}) bool) {
	for v := range in {
		if p(v) {
			matching <- v
		} else {
			others <- v
		}
	}
	close(matching)
	close(others)
}

// {{.Prefix.U}}{{.Type.U}}Map transforms a stream of {{.Type}} by applying a function fn to each item in the stream.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}Map(in <-chan {{.Type}}, out chan<- {{.Type}}, fn func({{.Type}}) {{.Type}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}FlatMap transforms a stream of {{.Type}} by applying a function fn to each item in the stream that
// gives zero or more results, all of which are sent out.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for {{.Type.U}}.
func {{.Prefix.U}}{{.Type.U}}FlatMap(in <-chan {{.Type}}, out chan<- {{.Type}}, fn func({{.Type}}) {{.Prefix.U}}{{.Type.U}}Collection) {
	for vi := range in {
		c := fn(vi)
		c.Foreach(func(vo {{.Type}}) {
			out <- vo
		})
	}
	close(out)
}
