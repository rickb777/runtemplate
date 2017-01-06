# runtemplate built-in templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types.
Their API style has been loosely influenced by the excellent Scala collection classes.

There are two categories: threadsafe collections and non-threadsafe collections. These are generally not
intended to be used together directly, although it is possible to transfer content from either collections
of one category to collections of the other category.

If you want to mix both categories, you will need to generate the output code in different packages. To do this,
specify different directories for the generated code via `-output`.

## General Flags

The built-in collections support a small number of flags that allow you to control whether extra methods are
generated or not.

 * `Comparable=true` - use this for types that are comparable (== and !=), such as strings, ints and floats.
 * `Ordered=true` - use this for types that are ordered (<, <=, >=, >), such as ints and floats.
 * `Numeric=true` - use this for types that support arithmetic operations, such as ints and floats.
 * `Stringer=true` - use this to include the `String()` method (and related); omit this if you prefer to provide your own.
 * `Mutable=true` - use this to include mutation methods; omit this if you need immutable collections.

See [Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) and
[Comparison operators](https://golang.org/ref/spec#Comparison_operators).

## Simple Collections
### collections/collection.tpl

This template generates a `<Type>Collection` interface for some specified type.
The type can be a pointer to a type if preferred.
All four options (above) are supported.

The list and set templates (below) both implement this interface.

Example use:
```
//go:generate runtemplate -tpl collections/collection.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

### collections/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.
All four options (above) are supported.

Example use:
```
//go:generate runtemplate -tpl collections/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more
than one goroutine at a time.

### collections/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types.
However, these should not be pointer types (a set of pointers would be of little value).

Other options: `Ordered` (see above)

Example use:
```
//go:generate runtemplate -tpl collections/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

The generated code is not suitable for access by more than one goroutine at a time.

## Thread-safe Collections
### threadasafe/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types.

The generated code includes the necessary read and write locking to ensure it is suitable for access by
any number of concurrent goroutines.

### threadasafe/map.tpl

This template generates a `<Key><Type>Map` for some specified key-type and content type. It uses Go's
built-in `map` internally. Specify the key type using `Key=type`. The key type
and content type can be user-defined or built-in types as needed.

The generated code includes the necessary read and write locking to ensure it is suitable for access by
any number of concurrent goroutines.

