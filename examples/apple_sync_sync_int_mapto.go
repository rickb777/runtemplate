// Generated from plumbing/mapTo.tpl with Type=Apple ToType=int

package examples

// SyncAppleMapToSyncInt transforms a stream of Apple to a stream of int.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func SyncAppleMapToSyncInt(in <-chan Apple, out chan<- int, fn func(Apple) int) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// SyncAppleFlatMapToSyncInt transforms a stream of Apple to a stream of int.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func SyncAppleFlatMapToSyncInt(in <-chan Apple, out chan<- int, fn func(Apple) SyncIntCollection) {
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

