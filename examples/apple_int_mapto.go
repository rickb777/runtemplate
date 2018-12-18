// Generated from plumbing/mapTo.tpl with Type=Apple ToType=int
// by runtemplate v3.0.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

// AppleMapToInt transforms a stream of Apple to a stream of int.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func AppleMapToInt(in <-chan Apple, out chan<- int, fn func(Apple) int) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// AppleFlatMapToInt transforms a stream of Apple to a stream of int.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func AppleFlatMapToInt(in <-chan Apple, out chan<- int, fn func(Apple) IntCollection) {
	for vi := range in {
		c := fn(vi)
		if c.NonEmpty() {
			for vo := range c.Send() {
				out <- vo
			}
		}
	}
	close(out)
}
