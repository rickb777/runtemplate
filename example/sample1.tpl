// Generated from {{.TemplateFile}} with Type={{.Type}} PName={{.PType}} LName={{.LType}}
// Full path {{.TemplatePath}}
// PWD={{.PWD}}
// GOPATH={{.GOPATH}}
// GOROOT={{.GOROOT}}

package {{.Package}}

import "fmt"

{{if .Egg}}
// you found the easter egg!
{{end}}

func As{{.Type}}(s string) ({{.Type}}, error) {
	i0 := _{{.Type}}_index[0]
	for j := 1; j < len(_{{.Type}}_index); j++ {
		i1 := _{{.Type}}_index[j]
		p := _{{.Type}}_name[i0:i1]
		if s == p {
			return {{.Type}}(j-1), nil
		}
		i0 = i1
	}
	return 0, errors.New(s + ": unrecognised {{.Type}} name")
}
