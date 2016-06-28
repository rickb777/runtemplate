# runtemplate built-in templates

A selection of built-in templates is included with `runtemplate`. These provide type-safe collection types.
Their API style has been loosely influenced by the excellent Scala collection classes.

There are two categories: threadsafe collections and non-threadsafe collections. These are generally not
intended to be used together directly, although it is possible to transfer content from either collections
of one category to collections of the other category.

To mix both categories, you will need to generate the output code in different packages. To do this,
specify different directories for the generated code via `-output`.

## Simple Collections
### collections/collection.tpl

This template generates a <Type>Collection interface for some specified type. Use the `-type` parameter.
The type can be a pointer to a type if preferred.

The list and set templates (below) both implement this interface.

### collections/list.tpl

This template generates a <Type>List for some specified type. Use the `-type` parameter. The type can be
a pointer to a type if preferred.

The generated code is a simple wrapper around a slice of the type. It is not suitable for access by more
than one goroutine at a time.

### collections/set.tpl

This template generates a <Type>Set for some specified type. Use the `-type` parameter. It accepts both
user-defined and built-in Go types.

The generated code is not suitable for access by more than one goroutine at a time.

## Thread-safe Collections
### threadasafe/set.tpl

This template generates a <Type>Set for some specified type. Use the `-type` parameter. It accepts both
user-defined and built-in Go types.

The generated code includes the necessary read and write locking to ensure it is suitable for access by
any number of concurrent goroutines.

### threadasafe/map.tpl

This template generates a <Key><Type>Map for some specified key-type and content type. It uses Go's
built-in `map` internally. Use the `-type` parameter and set the key type using `Key=type`. The key-type
and content type can be user-defined or built-in types as needed.

The generated code includes the necessary read and write locking to ensure it is suitable for access by
any number of concurrent goroutines.

