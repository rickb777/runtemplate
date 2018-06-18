# runtemplate built-in templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types. Their API style has been loosely influenced by the excellent Scala collection classes.

There are several main categories:

 * simple types using Go's slices and maps
 * encapsulated types for lists, sets and maps - both simple and thread-safe variants
 * channel plumbing

The threadsafe collections and non-threadsafe collections are generally not intended to be used together directly, although it is possible to transfer content from either collections of one category to collections of the other category.

If you want to mix both categories, you will need to generate the output code with different prefixes or in different packages. To set a prefix, pass `Prefix=Abc` or similar to `runtemplate`. To do the latter, specify different directories for the generated code via `-output`.


## General Flags

The built-in collections support flags that allow you to control the generated names.

 * `Type=T` - use this to specify the name of the generated type
 * `Prefix=X` - use this to prefix the name of the generated type
 * `Key=K` - use this to specify the name of the generated key type (for maps only, see below)

The built-in collections support a small number of flags that allow you to control whether extra methods are generated or not.

 * `Comparable:true` - use this for types that are comparable (== and !=), such as strings, ints, floats and structs.
 * `Ordered:true` - use this for types that are ordered (<, <=, >=, >), such as strings, ints and floats (but not structs).
 * `Numeric:true` - use this for types that support arithmetic operations, such as ints and floats (but not structs).
 * `Stringer:true` - use this to include the `String()` method (and related); omit this if you prefer to provide your own.
 * `KeyList:<type>` - for maps only, this provides a slice type for the keys in this map. This is returned from the `Keys()` method. It is also used for sorting the output of the stringer methods by the keys, which affects `MkString3(...)`, `MkString()` and `String()`.
 * `ValueList:<type>` - for maps only, this provides a slice type for the values in this map. This is returned from the `Values()` method.
 * `Import:<imports>` - extra Go imports; the literals `\n` and `\t` are replaced with their character equivalent, allowing multiple imports. It's likely that single quotes will be needed to enclose the entire Import parameter, because double-quotes are also needed around the import string itself in Go syntax.

The choice of flags is up to you and needs to be done with the language specification in mind - see [Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) and
[Comparison operators](https://golang.org/ref/spec#Comparison_operators). If you set a flag that is impossible for the chosen data type, the generated code won't compile, but no other bad thing will happen; so it will soon become obvious.


## 1. Direct Use of Go Slices and Maps

There are three categories of collection. The simplest category, described first, directly use Go's slices and maps.

There is a List type derived from Go slices: this template produces a type-safe slice for your chosen type and provides useful methods for handling its data. Similarly, the Map and Set types are derived from Go maps.

Because their base types are Go reference types, the generated types are also reference types, so you will never *need* to create pointers to them.

This category of collections are always mutable and do not attempt to encapsulate the underlying Go slice/map. So feel free to access these via their base type slice/map when necessary.

Be careful, however, that these collections cannot be safely shared between goroutines and should never be sent via channels (unless the sending end stops using things it has sent).


### simple/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.

The supported options are: Comparable, Ordered, Numeric, Stringer. The generated types are always mutable.

Example use:

```
//go:generate runtemplate -tpl simple/list.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more than one goroutine at a time.

Examples: [IntList](builtintest/simple/x_int_list.go), [StringList](builtintest/simple/x_string_list.go)


### simple/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would rarely be of any value).

The supported options are: Numeric, Stringer. The generated types are always mutable.

Example use:
```
//go:generate runtemplate -tpl simple/set.tpl Type=int Stringer:true Numeric:true
```

The generated code is simple wrapper around a map of the type: the set elements are used as map keys; the map uses zero-size values. It is not suitable for access by more than one goroutine at a time.

Examples: [IntSet](builtintest/simple/x_int_set.go), [StringSet](builtintest/simple/x_string_set.go)


### simple/map.tpl

This template generates a `<Key><Type>Map` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Numeric, Stringer. The generated types are always mutable.

Example use:
```
//go:generate runtemplate -tpl simple/map.tpl Key=string Type=int Stringer:true Comparable:true Numeric:true KeySlice:sort.StringSlice
```

The generated code is a simple wrapper around a map of the key and type. It not suitable for access by more than one goroutine at a time.

A tuple type is also generated: this is a struct that pairs up the key and value. A slice of such structs can be converted to and from the map type (assuming there are no duplicates), so the generated methods provide for this.

Examples: [IntIntMap](builtintest/simple/sx_int_int_map.go), [StringStringMap](builtintest/simple/sx_string_string_map.go)


## 2. Encapsulated Collections - Mutable

The second kind of collection encapsulate their data within structs. The unexported fields can only be accessed via the methods provided. It is therefore possible to make meaningful differentiation between mutable and immutable flavours, both of which are available to you. The mutable flavour is described first.

There are two variants:

 * *fast* in which there is no thread locking
 * *thread-safe* in which all accesses have the necessary lock (read accesses have a read lock, write accesses have a write lock).

The lack of locking in the fast variant requires less CPU work, but this means you cannot share these mutable collections between goroutines unless you add your own locking. It is even unwise to transmit these collection via channels because it is hard to avoid unwanted aliases that lead to race conditions.

Conversely, collections of the threadsafe variant can be shared between goroutines always, and can be sent via channels. Standard Go locks are used to allow (only) multiple concurrent read accesses, or a single write access, at any one time.


### fast/list.tpl and threadsafe/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.

The supported options are: Comparable, Ordered, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl fast/list.tpl       Type=int Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl threadsafe/list.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```

Examples: 
Fast:       [IntList](builtintest/fast/x_int_list.go),       [StringList](builtintest/fast/x_string_list.go).
Threadsafe: [IntList](builtintest/threadsafe/x_int_list.go), [StringList](builtintest/threadsafe/x_string_list.go).


### fast/set.tpl and threadsafe/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would be of little value).

The supported options are: Comparable, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl fast/set.tpl       Type=int Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl threadsafe/set.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```

Examples:
Fast:       [IntSet](builtintest/fast/x_int_set.go),       [StringSet](builtintest/fast/x_string_set.go).
Threadsafe: [IntSet](builtintest/threadsafe/x_int_set.go), [StringSet](builtintest/threadsafe/x_string_set.go).


### fast/map.tpl and threadsafe/map.tpl

This template generates a `<Key><Type>Map` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl fast/map.tpl       Key=string Type=int Stringer:true Comparable:true Numeric:true KeySlice:sort.StringSlice
//go:generate runtemplate -tpl threadsafe/map.tpl Key=string Type=int Stringer:true Comparable:true Numeric:true KeySlice:sort.StringSlice
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
//go:generate runtemplate -tpl fast/collection.tpl       Type=int Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl threadsafe/collection.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```
Examples:
Fast:       [IntCollection](builtintest/fast/x_int_collection.go),       [StringCollection](builtintest/fast/x_string_collection.go),
Threadsafe: [IntCollection](builtintest/threadsafe/x_int_collection.go), [StringCollection](builtintest/threadsafe/x_string_collection.go),


## 3. Encapsulated Collections - Immutable

The third kind of collection also encapsulate their data within structs, but in this case the access methods do not allow the internal data to be altered. Such immutable data structures have benefits in many use-cases, such as being easy to share between goroutines without any need for locking.

These immutable collections are constructed via the `NewXxxXxx` functions, all of which accept the input data. 


### immutable/list.tpl

This template generates a `<Type>List` for some specified type. The type can be a pointer to a type if preferred.

The supported options are: Comparable, Ordered, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl immutable/list.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```

Examples: [IntList](builtintest/immutable/x_int_list.go), [StringList](builtintest/immutable/x_string_list.go).


### immutable/set.tpl

This template generates a `<Type>Set` for some specified type. It accepts both user-defined and built-in Go types. However, these should not be pointer types (a set of pointers would be of little value).

The supported options are: Comparable, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl immutable/set.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```

Examples: [IntSet](builtintest/immutable/x_int_set.go), [StringSet](builtintest/immutable/x_string_set.go).


### immutable/map.tpl

This template generates a `<Key><Type>Map` for some specified type. It accepts both user-defined and built-in Go types.

The supported options are: Comparable, Numeric, Stringer.

Example use:
```
//go:generate runtemplate -tpl immutable/map.tpl Key=string Type=int Stringer:true Comparable:true Numeric:true KeySlice:sort.StringSlice
```

A tuple type is also generated: this is a struct that pairs up the key and value. A slice of such structs can be converted to and from the map type (assuming there are no duplicates), so the generated methods provide for this.

Examples: [IntIntMap](builtintest/immutable/x_int_int_list.go), [StringStringMap](builtintest/immutable/x_string_string_map.go).


### immutable/collection.tpl

This template generates a `<Type>Collection` interface for some specified type. The type can be a pointer to a type if preferred. All options (above) are supported.

The **list** and **set** templates (above) both implement this interface.

Example use:
```
//go:generate runtemplate -tpl immutable/collection.tpl Type=int Stringer:true Comparable:true Ordered:true Numeric:true
```
Examples: [IntCollection](builtintest/immutable/x_int_collection.go), [StringCollection](builtintest/immutable/x_string_collection.go),


## 4. Channel-based Plumbing

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
 * `<Type>FlatMap` transforms a stream of `Type` by applying to each item in the stream a function that yields zero or more `Type` values, all of which are sent out.

All of these terminate their loops when their input channels get closed. They then close their output channels.

Example use:
```
//go:generate runtemplate -tpl plumbing/plumbing.tpl Type=int
```
Example: [ApplePlumbing](builtintest/threadsafe/x_apple_plumbing.go).


### plumbing/mapTo.tpl

This template provides two stream-based inline converters using Go channels.

 * `<Type>Map<ToType>` transforms a stream of `Type` to a stream of `ToType` using a given transformation function.
 * `<Type>FlatMap<ToType>` transforms a stream of `Type` to a stream of `ToType` using a given transformation function that yields zero or more `ToType` values, all of which are sent out.

Both of them terminate their loops when their input channels get closed. They then close their output channels.

Example use:
```
//go:generate runtemplate -tpl plumbing/mapTo.tpl Type=Apple ToType=Pear
```
Example: [ApplePearMapTo](builtintest/threadsafe/x_apple_x_pear_mapto.go).


## 5. Stringy Types

### types/stringy.tpl

This template allows you to name different kinds of string types. You might want to do this because it can greatly increase the type-safety of any code that manipulates a lot of string values. It prevents incompatible string values being assigned to each other.

The stringy types provide some of the stanrdard `strings` functions as methods. 

They also provide methods to make them compatible with the SQL and marshalling APIs, just like ordinary strings.

Finally, they can provide a sorting utility `sort.Interface` implementation. This is only generated when `SortableSlice:true` is specified.
