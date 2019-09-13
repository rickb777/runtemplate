// Generated from plumbing/plumbing.tpl with Type=Apple
// No other options are available.
// by runtemplate v3.5.3
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package threadsafe

//-------------------------------------------------------------------------------------------------

// This plumbing suite provides standard functions for piping data between goroutines.
// All these functions run until the input channel is closed (or all input channels are closed, if
// multiple). They then close their output channel(s).
// None of these functions create a goroutine - this must be done at the call site.

//-------------------------------------------------------------------------------------------------

// X1AppleGenerator produces a stream of Apple based on a supplied generator function.
// The function fn is invoked N times with the integers from 0 to N-1. Each result is sent out.
// Finally, the output channel is closed and the generator terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleGenerator(out chan<- Apple, iterations int, fn func(int) Apple) {
	X1AppleGenerator3(out, 0, iterations-1, 1, fn)
}

// X1AppleGenerator produces a stream of Apple based on a supplied generator function.
// The function fn is invoked *(|to - from|) / |stride|* times with the integers in the range specified by
// from, to and stride. If stride is negative, from should be greater than to.
// For each iteration, the computed function result is sent out.
// If stride is zero, the loop never terminates. Otherwise, after the generator has reached the
// loop end, the output channel is closed and the generator terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleGenerator3(out chan<- Apple, from, to, stride int, fn func(int) Apple) {
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

// X1AppleDelta duplicates a stream of Apple to two output channels.
// When the sender closes the input channel, both output channels are closed then the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleDelta(in <-chan Apple, out1, out2 chan<- Apple) {
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

// X1AppleZip2 interleaves two streams of Apple.
// Each input channel is used in turn, alternating between them.
// The function terminates when *both* input channels have been closed by their senders.
// The output channel is then closed also.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleZip2(in1, in2 <-chan Apple, out chan<- Apple) {
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

// X1AppleMux2 multiplexes two streams of Apple into a single output channel.
// Each input channel is used as soon as it is ready.
// When a signal is received from the closer channel, the output channel is then closed.
// Concurrently, both input channels are then passed into blackholes that comsume them until they too are closed,
// and the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleMux2(in1, in2 <-chan Apple, closer <-chan bool, out chan<- Apple) {
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
	go X1AppleBlackHole(in1)
	go X1AppleBlackHole(in2)
	close(out)
}

// X1AppleBlackHole silently consumes a stream of Apple.
// It terminates when the sender closes the channel.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleBlackHole(in <-chan Apple) {
	for _ = range in {
		// om nom nom
	}
}

// X1AppleFilter filters a stream of Apple, silently dropping elements that do not match the predicate p.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleFilter(in <-chan Apple, out chan<- Apple, p func(Apple) bool) {
	for v := range in {
		if p(v) {
			out <- v
		}
	}
	close(out)
}

// X1ApplePartition filters a stream of Apple into two output streams using a predicate p, those that
// match and all others.
// When the sender closes the input channel, both output channels are closed then the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1ApplePartition(in <-chan Apple, matching, others chan<- Apple, p func(Apple) bool) {
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

// X1AppleMap transforms a stream of Apple by applying a function fn to each item in the stream.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleMap(in <-chan Apple, out chan<- Apple, fn func(Apple) Apple) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// X1AppleFlatMap transforms a stream of Apple by applying a function fn to each item in the stream that
// gives zero or more results, all of which are sent out.
// When the sender closes the input channel, the output channel is closed then the function terminates.
//
// It is part of the Plumbing function suite for Apple.
func X1AppleFlatMap(in <-chan Apple, out chan<- Apple, fn func(Apple) X1AppleCollection) {
	for vi := range in {
		c := fn(vi)
		c.Foreach(func(vo Apple) {
			out <- vo
		})
	}
	close(out)
}
