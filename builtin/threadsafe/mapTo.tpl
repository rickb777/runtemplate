// Generated from {{.TemplateFile}} with Type={{.Type}} ToType={{.ToType}}

package {{.Package}}

// {{.UType}}MapTo{{.UToType}} transforms a stream of {{.PType}} to a stream of {{.PToType}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.UType}}MapTo{{.UToType}}(in <-chan {{.PType}}, out chan<- {{.PToType}}, fn func({{.PType}}) {{.PToType}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.UType}}FlatMapTo{{.UToType}} transforms a stream of {{.PType}} to a stream of {{.ToType}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.UType}}FlatMapTo{{.UToType}}(in <-chan {{.PType}}, out chan<- {{.PToType}}, fn func({{.PType}}) {{.ToType}}Collection) {
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

