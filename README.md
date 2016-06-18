# runtemplate

This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
use-case is for source code generation.

You can install it with:

```
go get github.com/rickb777/runtemplate
```

It is intended to be used directly from the command-line and also with `go generate`.

It supports light-weight dependency checking, i.e. much less work is done when the generated output file
already exists and is up to date.

## Command-Line

Flexible option parsing is provided. Example

```
runtemplate -tpl filename.tpl -output outfile.go -type MyStruct -deps foo.go,bar.go Option1=Value1 Option2=Value2
```

 * `-tpl <name>`
   - (required) the name of the input template.

 * `-type <name>`
   - (optional) the name of the primary Go type for which code generation is being used; the file `<name>.go`
     (in lowercase) is checked for modification timestamp and treated as a dependency, if it exists.

 * `-output <name>`
   - the name of the output file to be written. If `-tpl` is not specifed, `-output` is required,
     otherwise it is optional.

 * `-deps <name>,<name>,...`
   - adds more dependencies to be checked in addition to the template itself and the 'type' file (if any).

 * `-f`
   - force output generation; if this is not set the output file is only produced when it is older than the
     dependencies

 * `-v`
   - verbose info messages

 * key=value ...
   - (optional) supply a (list of) key/value pairs that re passed in to the template. `true` and `false` are
     converted to booleans, allowing conditional blocks within your templates.

The option parser will also infer the template and output file names, so it is also permitted to use

```
runtemplate outfile.go filename.tpl Type=MyStruct Option1=Value1 Option2=Value2
```

i.e. to omit the explicit flags `-tpl` and `-output`.

## Go Generate

Simply put the `go generate` comment in your code like this:

```
//go:generate runtemplate -tpl filename.tpl -output outfile.go Option1=Value1 Option2=Value2
```

When you run `go generate`, it will find these marked comments and execute their commands. This will
`runtemplate` against the specified template, passing in whatever options have have been specified
on the command line as a map.

The explicit `-tpl` and `-output` flags can be omitted if preferred, but both the template file and
output file must be specified.

## Template

In the template file, you can access the key=value pairs simply by their keys. For instance:

`{{ .Option1 }}`

If `-type` is specified, its value is provided in several variants:

 * `.Type`  - the type name (without '*' prefix)
 * `.PType` - the type name (prefixed by '*' if supplied)
 * `.LType` - the type name having its first character converted to lowercase.

Additional keys are also made available:

 * `.OutFile` - the name of the output file
 * `.TemplateFile` - the template name as specified
 * `.TemplatePath` - the location and name of the actual template file used
 * `.Package` - the name of the directory of the output file (often the current directory)
 * `.GOARCH`, `.GOOS`, `.GOPATH`, `GOROOT` - the value of Go environment variables.

Also included are some filters that may be helpful.

 * title - Converts the input to Title Case.
 * upper - Converts the input to UPPER CASE.
 * lower - Converts the input to lower case.
 * splitDotFirst - Given an input that has a '.' separator, returns the part before the first '.'.
 * splitDotLast - Given an input that has a '.' separator, returns the part after the last '.'.

The last two are useful for getting only the package name or only the type name if passed an input of `package.Type`.

## Template Path

Templates are located by following `TEMPLATEPATH`, an optional environment variable. If it is defined, it
is used like `PATH`, i.e. a colon-separate list of directories to be searched.

If `TEMPLATEPATH` is absent, its default is `TEMPLATEPATH=.`, i.e. templates are relative to the current directory.

If available, the location of the builtin templates is also added to `TEMPLATEPATH`; this is found at
`$GOPATH/src/github/rickb777/runtemplate/builtin`.
