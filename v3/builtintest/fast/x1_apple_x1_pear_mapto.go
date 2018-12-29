// Generated from plumbing/mapTo.tpl with Type=Apple ToType=Pear
// by runtemplate v3.1.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package fast

// X1AppleMapToX1Pear transforms a stream of Apple to a stream of Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func X1AppleMapToX1Pear(in <-chan Apple, out chan<- Pear, fn func(Apple) Pear) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// X1AppleFlatMapToX1Pear transforms a stream of Apple to a stream of Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func X1AppleFlatMapToX1Pear(in <-chan Apple, out chan<- Pear, fn func(Apple) X1PearCollection) {
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
