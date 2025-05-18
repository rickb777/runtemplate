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

// {{.Type.U}}Slice attaches the methods of sort.Interface to []{{.Type.Name}}, sorting in increasing order.
type {{.Type.U}}Slice []{{.Type.Name}}
