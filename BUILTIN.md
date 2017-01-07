# runtemplate built-in templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types. Their API style has been loosely influenced by the excellent Scala collection classes.

There are several main categories:

 * simple types using Go's slices and maps
 * encapsulated types for lists and sets - both simple and thread-safe variants
 * encapsulated type-safe maps - both simple and thread-safe variants
 * channel plumbing

The threadsafe collections and non-threadsafe collections are generally not intended to be used together directly, although it is possible to transfer content from either collections of one category to collections of the other category.

If you want to mix both categories, you will need to generate the output code with different prefixes or in different packages. To do the latter, specify different directories for the generated code via `-output`.

## General Flags

The built-in collections support a flag that allow you to control the generated names.

 * `Prefix=X` - use this to prefix the name of the generated type

The built-in collections support a small number of flags that allow you to control whether extra methods are generated or not.

 * `Comparable=true` - use this for types that are comparable (== and !=), such as strings, ints, floats and structs.
 * `Ordered=true` - use this for types that are ordered (<, <=, >=, >), such as strings, ints and floats (but not structs).
 * `Numeric=true` - use this for types that support arithmetic operations, such as ints and floats (but not structs).
 * `Stringer=true` - use this to include the `String()` method (and related); omit this if you prefer to provide your own.
 * `Mutable=true` - use this to include mutation methods; omit this if you need immutable collections. (Note that the simple collections are inherently mutable.)

See [Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) and
[Comparison operators](https://golang.org/ref/spec#Comparison_operators).

## 1. Simple Lists and Sets

The simplest kind of collection directly use Go's slices and maps. These are mutable and do not attempt to encapsulate the underlying Go slice/map. Feel free to access these as slices/maps as appropriate.

### simple/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred. The supported options are: Comparable, Ordered, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl simple/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true
```

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more than one goroutine at a time.

### simple/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would be of little value).

The supported options are: Comparable, Ordered, Numeric, Stringer, Mutable.

Example use:
```
//go:generate runtemplate -tpl fast/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

The generated code is not suitable for access by more than one goroutine at a time.

## 2. Encapsulated Collections

### fast/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred. The supported options are: Comparable, Ordered, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl fast/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more than one goroutine at a time.

### fast/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would be of little value).

The supported options are: Comparable, Ordered, Numeric, Stringer, Mutable.

Example use:
```
//go:generate runtemplate -tpl fast/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

The generated code is not suitable for access by more than one goroutine at a time.

### fast/collection.tpl

This template generates a `<Type>Collection` interface for some specified type. The type can be a pointer to a type if preferred. All options (above) are supported.

The list and set templates (below) both implement this interface.

Example use:
```
//go:generate runtemplate -tpl fast/collection.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

### threadsafe/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Ordered, Numeric, Stringer, Mutable.

The generated code includes the necessary read and write locking to ensure it is suitable for access by any number of concurrent goroutines.

### threadsafe/map.tpl

This template generates a `<Key><Type>Map` for some specified key-type and content type. It uses Go's built-in `map` internally. Specify the key type using `Key=type`. The key type and content type can be user-defined or built-in types as needed.

The supported options are: Comparable, Ordered, Numeric, Stringer, Mutable.

The generated code includes the necessary read and write locking to ensure it is suitable for access by any number of concurrent goroutines.

### threadsafe/collection.tpl

This is the same as `fast/collection.tpl`.

## 3. Typesafe Maps

### map/simple.tpl

TODO

### map/threadsafe.tpl

TODO

## 4. Channel-based Plumbing

### plumbing/core.tpl

TODO

### plumbing/mapTo.tpl

TODO
