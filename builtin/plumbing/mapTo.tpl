// Generated from {{.TemplateFile}} with Type={{.Type}} ToType={{.ToType}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md#simplelisttpl

package {{.Package}}

{{if .HasImport}}
import (
	{{.Import}}
)

{{end -}}
// {{.UPrefix}}{{.UType}}MapTo{{.UToPrefix}}{{.UToType}} transforms a stream of {{.PType}} to a stream of {{.PToType}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.UPrefix}}{{.UType}}MapTo{{.UToPrefix}}{{.UToType}}(in <-chan {{.PType}}, out chan<- {{.PToType}}, fn func({{.PType}}) {{.PToType}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.UPrefix}}{{.UType}}FlatMapTo{{.UToPrefix}}{{.UToType}} transforms a stream of {{.PType}} to a stream of {{.ToType}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.UPrefix}}{{.UType}}FlatMapTo{{.UToPrefix}}{{.UToType}}(in <-chan {{.PType}}, out chan<- {{.PToType}}, fn func({{.PType}}) {{.UToPrefix}}{{.UToType}}Collection) {
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
