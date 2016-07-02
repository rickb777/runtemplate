# runtemplate

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/runtemplate)
[![Build Status](https://travis-ci.org/rickb777/runtemplate.svg?branch=master)](https://travis-ci.org/rickb777/runtemplate)

This application provides a simple way of executing standard Go templates from the command line. The obvious
use-case is for source code generation, amongst many others.

You can install it with:

```
go get github.com/rickb777/runtemplate
```

It is intended to be used directly from the command-line and also with `go generate`.

It supports light-weight dependency checking, i.e. less work is done when the generated output file
already exists and is up to date.

## Command-Line

Flexible option parsing is provided. Example

```
runtemplate -tpl filename.tpl -output outfile.go -deps foo.go,bar.go Type=MyStruct Option1=Value1 Option2=Value2
```

 * `-tpl <name>`
   - (required) the name of the input template.

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
   - (optional) supply a (list of) key/value pairs that are passed in to the template. `true` and `false` are
     converted to booleans, allowing conditional blocks within your templates.

The option parser will also infer the template and output file names, so it is also permitted to use

```
runtemplate outfile.go filename.tpl Type=MyStruct Option1=Value1 Option2=true
```

i.e. to omit the explicit flags `-tpl` and `-output`.

Furthermore, the output file may be completely omitted

```
runtemplate filename.tpl Type=MyStruct Option1=Value1 Option2=true
```

in which case a name will be computed from all the values of the key/value pairs excluding true/false,
plus the name of the template. For the example above, it will be `mystruct_value1_filename.go`.

## Go Generate

Easy. Just put the `go generate` comment in your code like this:

```
//go:generate runtemplate -tpl filename.tpl -output outfile.go Option1=Value1 Option2=Value2
```

When you run `go generate`, it will find these marked comments and execute their commands. This will
`runtemplate` against the specified template, passing in whatever options have have been specified
on the command line as a map.

## Template

In the template file, you can access the key=value pairs simply by their keys. For instance:

`{{ .Option1 }}`

Boolean true/false key-values are available for `{{if .Flag}} ... {{end}}` conditional use.

For all other key-values, the values are supplemented by additional entries in the template's context. These are:

 * `.Type`  - the type name (without any '*' prefix)
 * `.PType` - the type name (prefixed by '*' if supplied)
 * `.UType` - the type name having its first character converted to uppercase - useful for exported identifiers.
 * `.LType` - the type name having its first character converted to lowercase - useful for internal identifiers.
 * `.TypeStar` - a '*' if the type is a pointer type, otherwise blank
 * `.TypeAmp` - a '&' if the type is a pointer type, otherwise blank

This table shows two examples of context symbols defined for Type=Foo and Type=*Foo.

|              |  Type=Foo    |  Type=*Foo   |
| ------------ | ------------ | ------------ |
| `.Type`      |  `Foo`       |  `Foo`       |
| `.PType`     |  `Foo`       |  `*Foo`      |
| `.UType`     |  `Foo`       |  `Foo`       |
| `.LType`     |  `foo`       |  `foo`       |
| `.TypeStar`  |              |  `*`         |
| `.TypeAmp`   |              |  `&`         |

Be aware that your shell might expand * so you may need suitable quote marks, such as `'Type=*Foo'`. This
is not needed when using go:generate comment lines.

Additional keys are also made available:

 * `.OutFile` - the name of the output file
 * `.TemplateFile` - the template name as specified
 * `.TemplatePath` - the location and name of the actual template file used
 * `.Package` - the name of the directory of the output file (often the current directory)
 * `.GOARCH`, `.GOOS`, `.GOPATH`, `GOROOT` - the value of Go environment variables.

Some filters are also included that may be helpful.

 * title - Converts the input to Title Case.
 * upper - Converts the input to UPPER CASE.
 * lower - Converts the input to lower case.
 * firstUpper - Converts the first character of input to upper case.
 * firstLower - Converts the first character of input to lower case.
 * splitDotFirst - Given an input that has a '.' separator, returns the part before the first '.'.
 * splitDotLast - Given an input that has a '.' separator, returns the part after the last '.'.

The last two are useful for getting only the package name or only the type name if passed an input of `package.Type`.

## Template Path

Templates are located by following `TEMPLATEPATH`, an optional environment variable. If it is defined, it
is used like `PATH`, i.e. a colon-separate list of directories to be searched.

If `TEMPLATEPATH` is absent, its default is `TEMPLATEPATH=.`, i.e. templates are relative to the current directory.

If available, the location of the builtin templates is also added to `TEMPLATEPATH`; this is found at
`$GOPATH/src/github/rickb777/runtemplate/builtin`.

# Built-in Templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types.
Their API style has been loosely influenced by other similar Go types and the excellent Scala collection classes.

[See BUILTIN](BUILTIN.md)
