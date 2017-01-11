# runtemplate built-in templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types. Their API style has been loosely influenced by the excellent Scala collection classes.

There are several main categories:

 * simple types using Go's slices and maps
 * encapsulated types for lists, sets and maps - both simple and thread-safe variants
 * channel plumbing

The threadsafe collections and non-threadsafe collections are generally not intended to be used together directly, although it is possible to transfer content from either collections of one category to collections of the other category.

If you want to mix both categories, you will need to generate the output code with different prefixes or in different packages. To set a prefix, pass `Prefix=Abc` or simiilar to `runtemplate`. To do the latter, specify different directories for the generated code via `-output`.

## General Flags

The built-in collections support a flag that allow you to control the generated names.

 * `Prefix=X` - use this to prefix the name of the generated type

The built-in collections support a small number of flags that allow you to control whether extra methods are generated or not.

 * `Comparable=true` - use this for types that are comparable (== and !=), such as strings, ints, floats and structs.
 * `Ordered=true` - use this for types that are ordered (<, <=, >=, >), such as strings, ints and floats (but not structs).
 * `Numeric=true` - use this for types that support arithmetic operations, such as ints and floats (but not structs).
 * `Stringer=true` - use this to include the `String()` method (and related); omit this if you prefer to provide your own.
 * `Mutable=true` - use this to include mutation methods; omit this if you need immutable collections. (Note that the simple collections are inherently mutable.)

The choice of flags is up to you and needs to be done with the language specification in mind - see [Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) and
[Comparison operators](https://golang.org/ref/spec#Comparison_operators). If you set a flag that is impossible for the chosen data type, the generated code won't compile, but no other bad thing will happen; so it will soon become obvious.

## 1. Direct Use of Go Slices and Maps

There are three categories of collection. The simplest category, described first, directly use Go's slices and maps.

There is a List type derived from Go slices: this template produces a type-safe slice for your chosen type and provides useful methods for handling its data. Similarly, the Map and Set types are derived from Go maps.

Because their base types are Go reference types, the generated types are also reference types, so you will never *need* to create pointers to them.

This category of collections are always mutable and do not attempt to encapsulate the underlying Go slice/map. So feel free to access these via their base type slice/map when necessary.


### simple/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.

The supported options are: Comparable, Ordered, Numeric, Stringer. The generated types are always mutable.

Example use:

```
//go:generate runtemplate -tpl simple/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true
```

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more than one goroutine at a time.

Examples: [Int32List](builtintest/simple/x_int32_list.go), [StringList](builtintest/simple/x_string_list.go)


### simple/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would rarely be of any value).

The supported options are: Numeric, Stringer. The generated types are always mutable.

Example use:
```
//go:generate runtemplate -tpl simple/set.tpl Type=int32 Stringer=true Numeric=true
```

The generated code is simple wrapper around a map of the type: the set elements are used as map keys; the map uses zero-size values. It is not suitable for access by more than one goroutine at a time.

Examples: [Int32Set](builtintest/simple/x_int32_set.go), [StringSet](builtintest/simple/x_string_set.go)


### simple/map.tpl

This template generates a `<Key><Type>Map` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Numeric, Stringer. The generated types are always mutable.

Example use:
```
//go:generate runtemplate -tpl simple/map.tpl Key=string Type=int32 Stringer=true Comparable=true Numeric=true
```

The generated code is a simple wrapper around a map of the key and type. It not suitable for access by more than one goroutine at a time.

A tuple type is also generated: this is a struct that pairs up the key and value. A slice of such structs can be converted to and from the map type (assuming there are no duplicates), so the generated methods provide for this.

Examples: [IntIntMap](builtintest/simple/sx_int_int_map.go), [StringStringMap](builtintest/simple/sx_string_string_map.go)


## 2. Encapsulated Collections

The second kind of collection encapsulate their data within structs. The unexported fields can only be accessed via the methods provided. It is therefore possible to make meaningful differentiation between mutable and immutable variants, both of which are available to you.

There are two variants:

 * *fast* in which there is no thread locking
 * *thread-safe* in which all accesses have the necessary lock (read accesses have a read lock, write accesses have a write lock).

The lack of locking in the fast variant requires less CPU work, but this means you cannot share mutable collections between goroutines unless you add your own locking or use channels (in the latter case be *very* careful to avoid unwanted aliases). However, if you restrict yourself to immutable methods (i.e. set `Mutable=false`), then sharing is fine.

Collections of the threadsafe variant can be shared between goroutines always. Standard Go locks are used to allow (only) multiple concurrent read accesses, or a single write access, at any one time.


### fast/list.tpl and threadsafe/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.

The supported options are: Comparable, Ordered, Numeric, Stringer, Mutable.

Example use:
```
//go:generate runtemplate -tpl fast/list.tpl       Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
//go:generate runtemplate -tpl threadsafe/list.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

Examples: 
Fast:       [Int32List](builtintest/fast/x_int32_list.go),       [StringList](builtintest/fast/x_string_list.go).
Threadsafe: [Int32List](builtintest/threadsafe/x_int32_list.go), [StringList](builtintest/threadsafe/x_string_list.go).


### fast/set.tpl and threadsafe/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would be of little value).

The supported options are: Comparable, Numeric, Stringer, Mutable.

Example use:
```
//go:generate runtemplate -tpl fast/set.tpl       Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
//go:generate runtemplate -tpl threadsafe/set.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```

Examples:
Fast:       [Int32Set](builtintest/fast/x_int32_set.go),       [StringSet](builtintest/fast/x_string_set.go).
Threadsafe: [Int32Set](builtintest/threadsafe/x_int32_set.go), [StringSet](builtintest/threadsafe/x_string_set.go).


### fast/map.tpl and threadsafe/map.tpl

This template generates a `<Key><Type>Map` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Numeric, Stringer, Mutable.

Example use:
```
//go:generate runtemplate -tpl fast/map.tpl       Type=int32 Stringer=true Comparable=true Numeric=true Mutable=true
//go:generate runtemplate -tpl threadsafe/map.tpl Type=int32 Stringer=true Comparable=true Numeric=true Mutable=true
```

A tuple type is also generated: this is a struct that pairs up the key and value. A slice of such structs can be converted to and from the map type (assuming there are no duplicates), so the generated methods provide for this.

Examples:
Fast:       [IntIntMap](builtintest/fast/x_int_int_list.go),       [StringStringMap](builtintest/fast/x_string_string_map.go).
Threadsafe: [IntIntMap](builtintest/threadsafe/x_int_int_list.go), [StringStringMap](builtintest/threadsafe/x_string_string_map.go).


### fast/collection.tpl and threadsafe/collection.tpl

This template generates a `<Type>Collection` interface for some specified type. The type can be a pointer to a type if preferred. All options (above) are supported.

The **list** and **set** templates (above) both implement this interface.

Example use:
```
//go:generate runtemplate -tpl fast/collection.tpl       Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
//go:generate runtemplate -tpl threadsafe/collection.tpl Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true Mutable=true
```
Examples:
Fast:       [Int32Collection](builtintest/fast/x_int32_collection.go),       [StringCollection](builtintest/fast/x_string_collection.go),
Threadsafe: [Int32Collection](builtintest/threadsafe/x_int32_collection.go), [StringCollection](builtintest/threadsafe/x_string_collection.go),


## 3. Channel-based Plumbing

### plumbing/core.tpl

This template provides stream-based processing primitives using Go channels.

 * `<Type>Generator` and `<Type>Generator3` create a stream of `Type` values using a generator function. 
 * `<Type>Delta` duplicates a stream of `Type` to two output channels.
 * `<Type>Zip2` interleaves two streams of `Type`.
 * `<Type>Mux2` multiplexes two streams of `Type` into a single output channel.
 * `<Type>BlackHole` silently consumes a stream of `Type`.
 * `<Type>Filter` filters a stream of `Type`, silently dropping elements that do not match a predicate function.
 * `<Type>Partition` filters a stream of `Type` into two output streams using a predicate function.
 * `<Type>Map` alters a stream of `Type` by applying a function to each item in the stream.
 * `<Type>FlatMap` transforms a stream of `Type` by applying to each item in the stream a function that yields zero or more `Type`s, all of which are sent out.

All of these terminate their loops when their input channels get closed. They then close their output channels.

### plumbing/mapTo.tpl

This template provides two stream-based inline converters using Go channels.

 * `<Type>Map<ToType>` transforms a stream of `Type` to a stream of `ToType` using a given transformation function.
 * `<Type>FlatMap<ToType>` transforms a stream of `Type` to a stream of `ToType` using a given transformation function that returns a `<ToType>Collection` for each `Type` value.

Both of them terminate their loops when their input channels get closed. They then close their output channels.
