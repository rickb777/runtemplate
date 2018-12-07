// Generated from {{.TemplateFile}} with Type={{.Type.Name}} ToType={{.ToType}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

{{if .HasImport}}
import (
	{{.Import}}
)

{{end -}}
// {{.Prefix.U}}{{.Type.U}}MapTo{{.ToPrefix.U}}{{.ToType.U}} transforms a stream of {{.Type.Name}} to a stream of {{.ToType.Name}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.Prefix.U}}{{.Type.U}}MapTo{{.ToPrefix.U}}{{.ToType.U}}(in <-chan {{.Type.Name}}, out chan<- {{.ToType}}, fn func({{.Type.Name}}) {{.ToType}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.Prefix.U}}{{.Type.U}}FlatMapTo{{.ToPrefix.U}}{{.ToType.U}} transforms a stream of {{.Type.Name}} to a stream of {{.ToType.Name}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.Prefix.U}}{{.Type.U}}FlatMapTo{{.ToPrefix.U}}{{.ToType.U}}(in <-chan {{.Type.Name}}, out chan<- {{.ToType}}, fn func({{.Type.Name}}) {{.ToPrefix.U}}{{.ToType.U}}Collection) {
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
