# runtemplate

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/runtemplate/examples)
[![Build Status](https://travis-ci.org/rickb777/runtemplate.svg?branch=master)](https://travis-ci.org/rickb777/runtemplate)
[![Code Coverage](https://img.shields.io/coveralls/rickb777/runtemplate.svg)](https://coveralls.io/r/rickb777/runtemplate)
[![Issues](https://img.shields.io/github/issues/rickb777/runtemplate.svg)](https://github.com/rickb777/runtemplate/issues)

This application provides a simple way of executing standard Go templates from the command line. The obvious use-case is for source code generation, amongst many others.

You can install it with:

```
go get github.com/rickb777/runtemplate
```

It is intended to be used directly from the command-line and also with `go generate`.

It supports light-weight dependency checking, i.e. less work is done when the generated output file already exists and is up to date.

A selection of [built-in templates](BUILTIN.md) is included with `runtemplate`. These provide type-safe collection types etc.

## Command-Line

Flexible option parsing is provided. Example

```
runtemplate -tpl filename.tpl -output outfile.go -deps foo.go,bar.go Type=MyStruct Option1:Value1 Option2:Value2
```

 * `-tpl <name>`
   - (required) the name of the input template.

 * `-output <name>`
   `-o <name>`
   - the name of the output file. If `-tpl` is not specifed, `-output` is required, otherwise it is optional. The name `-` causes writing to standard out instead of a file. Standard out is also used if this flag is unspecified and there are no `key=value` types.

 * `-deps <name>,<name>,...`
   - adds more dependencies to be checked in addition to the template itself and the 'type' file (if any).

 * `-f`
   - force output generation; if this is not set the output file is only produced when it is older than the
     dependencies

 * `-v`
   - verbose info messages

 * key:value ...
   - (optional) supply a (list of) simple key/value pairs that are passed in to the template. `true` and `false` are converted to booleans, allowing conditional blocks within your templates.

 * key=value ...
   - (optional) supply a (list of) key/value pairs that are passed in to the template. These are often Go types; extra synthetic values are also added, making it really easy to generate source code. This is described further below. `true` and `false` are converted to booleans, allowing conditional blocks within your templates.

The option parser will also infer the template and output file names, so it is also permitted to use either

```
runtemplate -output outfile.go -tpl filename.tpl Type=MyStruct Option1:Value1 Option2:true
runtemplate outfile.go filename.tpl Type=MyStruct Option1:Value1 Option2:true
```

i.e. to omit the explicit flags `-tpl` and `-output` provided the files are named.

Furthermore, the output file may be completely omitted:

```
runtemplate filename.tpl Type=MyStruct Option1=Value1 Option2=true Option3:foo
```

in which case a name will be computed from all the values of the key=value pairs (excluding true/false) in the order they are specified, plus the name of the template, conjoined with underscores, plus the extension '.go'. All the key:value settings are excluded.

For the example above, it will be `mystruct_value1_filename.go` because Option2 and Option3 are ignored for the reasons above.

## Go Generate

Easy. Just put the `go generate` comment in your code like this:

```
//go:generate runtemplate -tpl filename.tpl -output outfile.go Option1=Value1 Option2:Value2
```

When you run `go generate`, it will find these marked comments and execute their commands. This will `runtemplate` against the specified template, passing in whatever options have have been specified on the command line as a map.

## Template

In the template file, you can access the key=value or key:value pairs simply by their keys. For instance:

`{{ .Option1 }}`

Boolean true/false key-values are available for `{{if .Flag}} ... {{end}}` conditional use. Undefined values default to false.

### Simple Keys

The key:value syntax (using colon) defines simple values. These can be repeated to supply a slice of values, which is useful for the template `range` operator.

The values `true` and `false` are converted to booleans.

|              |  Foo:Bar     |  Foo:Ban Foo:Bar Foo:Baz  |
| ------------ | ------------ | ------------------------- |
| `.Foo`       |  `Bar`       |  {`Ban`, `Bar`, `Baz`}    |
| `.HasFoo`    |  `true`      |  `true`                   |

### Keys for Types

The key=value syntax (using equals) does more and is intended for identifiers in the programming language of the generated code (usually Go).

The values are supplemented by additional entries in the template's context. For example, given `Type=SomeValue` or `Type=*SomeValue`, these are:

 * `.Type`  - the type name (without any '*' prefix)
 * `.PType` - the type name (prefixed by '*' if supplied)
 * `.UType` - the type name having its first character converted to uppercase - useful for exported identifiers; dots are removed.
 * `.LType` - the type name having its first character converted to lowercase - useful for internal identifiers; dots are removed.
 * `.TypeStar` - a '*' if the type is a pointer type, otherwise blank
 * `.TypeAmp` - a '&' if the type is a pointer type, otherwise blank
 * `.HasType` - set to `true` to allow conditional expressions (it defaults to false if undefined)

This table shows two examples of context symbols defined for Type=Foo and Type=*Foo.

|              |  `Type=big.Int`  |  `Type=*big.Int`  |
| ------------ | ---------------- | ----------------- |
| `.Type`      |  `big.Int`       |  `big.Int`        |
| `.PType`     |  `big.Int`       |  `*big.Int`       |
| `.UType`     |  `BigInt`        |  `BigInt`         |
| `.LType`     |  `bigInt`        |  `bigInt`         |
| `.TypeStar`  |  blank           |  `*`              |
| `.TypeAmp`   |  blank           |  `&`              |
| `.HasType`   |  `true`          |  `true`           |

Be aware that your shell might expand `*` so you may need suitable quote marks, such as `'Type=*Foo'`. This is not needed when using go:generate comment lines.

### Prefix

If you need to generate code for several generated types and they need to co-exist within the same package, you can easily define a prefix to differentiate their names.

For every `<X>Type` template value that you specify (for some `<X>`), there is a corresponding special value `<X>Prefix` that is always predefined with a *blank* default value. But you can set it to something else, and if you do the generated types can use this to prefix their names. This only happens for keys that end in `Type`.

In short, the key's suffix `Type` is replaced with `Prefix`.

As well as `<X>Prefix`, there will be `<X>UPrefix`  and `<X>LPrefix` as above.

### Other settings

Additional settings are also made available:

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
 * condFirstUpper - Given a string and a boolean, converts the first character of input to upper case when the boolean is true. 
 * splitDotFirst - Given an input that has a '.' separator, returns the part before the first '.'.
 * splitDotLast - Given an input that has a '.' separator, returns the part after the last '.'.

The last two are useful for getting only the package name or only the type name if passed an input of `package.Type`.

## Template Path

Templates are located by following `TEMPLATEPATH`, an optional environment variable. If it is defined, it is used like `PATH`, i.e. a colon-separate list of directories to be searched.

If `TEMPLATEPATH` is absent, its default is `TEMPLATEPATH=.`, i.e. templates are relative to the current directory.

The builtin templates are also available and are searched if no other match is found. For example, template "types/stringy.tpl" will resolve to the built-in template of that name unless the TEMPLATEPATH contains another file with the same path.

# Built-in Templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types. Their API style has been loosely influenced by other similar Go types and the excellent Scala collection classes.

[See BUILTIN](BUILTIN.md)
