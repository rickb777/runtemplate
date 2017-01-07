// Generated from mapTo.tpl with Type=Apple ToType=Pear

package threadsafe

// PAppleMapToPPear transforms a stream of *Apple to a stream of *Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func PAppleMapToPPear(in <-chan *Apple, out chan<- *Pear, fn func(*Apple) *Pear) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// PAppleFlatMapToPPear transforms a stream of *Apple to a stream of Pear.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func PAppleFlatMapToPPear(in <-chan *Apple, out chan<- *Pear, fn func(*Apple) PPearCollection) {
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

