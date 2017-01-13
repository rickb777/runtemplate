// Generated from plumbing/mapTo.tpl with Type=Apple ToType=Pear

package threadsafe

// XAppleMapToXPear transforms a stream of Apple to a stream of Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func XAppleMapToXPear(in <-chan Apple, out chan<- Pear, fn func(Apple) Pear) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// XAppleFlatMapToXPear transforms a stream of Apple to a stream of Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func XAppleFlatMapToXPear(in <-chan Apple, out chan<- Pear, fn func(Apple) XPearCollection) {
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

