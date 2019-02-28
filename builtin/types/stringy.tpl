// A derived string-based type compatible with marshalling and database APIs.
//
// Generated from {{.TemplateFile}} with Type={{.Type.Name}}
// options: SortableSlice:{{.SortableSlice}}
// by runtemplate {{.AppVersion}}
// See https://github.com/johanbrandhorst/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
	"errors"
{{- if .SortableSlice}}
	"sort"
{{- end}}
	"strings"
	"database/sql/driver"
	"fmt"
)

// {{.Type.Name}} is a specialised kind of string.
type {{.Type.Name}} string

// Ptr returns the address of a {{.Type.Name}}.
func ({{.Type.L}} {{.Type.Name}}) Ptr() *{{.Type.Name}} {
	return &{{.Type.L}}
}

// String converts to a string and implements fmt.Stringer.
func ({{.Type.L}} {{.Type.Name}}) String() string {
	return string({{.Type.L}})
}

// TrimSpace removes surrounding whitespace.
func ({{.Type.L}} {{.Type.Name}}) TrimSpace() {{.Type.Name}} {
	return {{.Type.Name}}(strings.TrimSpace({{.Type.L}}.String()))
}

// ToLower converts the value to lowercase.
func ({{.Type.L}} {{.Type.Name}}) ToLower() {{.Type.Name}} {
	return {{.Type.Name}}(strings.ToLower(string({{.Type.L}})))
}

// ToUpper converts the value to uppercase.
func ({{.Type.L}} {{.Type.Name}}) ToUpper() {{.Type.Name}} {
	return {{.Type.Name}}(strings.ToUpper(string({{.Type.L}})))
}

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func ({{.Type.L}} *{{.Type.Name}}) Scan(value interface{}) error {
	if value == nil {
		*{{.Type.L}} = {{.Type.Name}}("")
		return nil
	}

	switch value.(type) {
	case string:
		*{{.Type.L}} = {{.Type.Name}}(value.(string))
	case []byte:
		*{{.Type.L}} = {{.Type.Name}}(string(value.([]byte)))
	case nil:
	default:
		return errors.New(fmt.Sprintf("{{.Type.Name}}.Scan(%#v)", value))
	}
	return nil
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func ({{.Type.L}} {{.Type.Name}}) Value() (driver.Value, error) {
	return string({{.Type.L}}), nil
}

//-------------------------------------------------------------------------------------------------

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// https://golang.org/pkg/encoding/#TextMarshaler
func ({{.Type.L}} {{.Type.Name}}) MarshalText() (text []byte, err error) {
	return []byte({{.Type.L}}.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
// https://golang.org/pkg/encoding/#TextUnmarshaler
func ({{.Type.L}} *{{.Type.Name}}) UnmarshalText(text []byte) error {
	return {{.Type.L}}.Scan(text)
}

//-------------------------------------------------------------------------------------------------
{{if .SortableSlice}}
// {{.Type.U}}Slice attaches the methods of sort.Interface to []{{.Type.Name}}, sorting in increasing order.
type {{.Type.U}}Slice []{{.Type.Name}}

func (p {{.Type.U}}Slice) Len() int           { return len(p) }
func (p {{.Type.U}}Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p {{.Type.U}}Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// SortedN is a convenience method.
func (p {{.Type.U}}Slice) Sorted() { sort.Sort(p) }
{{end}}
