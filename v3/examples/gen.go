// package examples contains example collection types using a selection of the built-in templates. They are built for
// int and for Apple, the latter being a simple dummy struct. The types are prefixed according to their template group.
// These are
//
//   * 'Fast...' for the fast collections,
//   * 'Immutable...' for the imutable collections,
//   * 'Simple...' for the simple collections,
//   * 'Sync...' for the threadsafe collections.
//
// In each group, there are two list types, two set types and two map types. There are also two collection types
// in all except the simple group.
//
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md for a fuller description.
package examples

// Simple Examples
//go:generate runtemplate -tpl simple/list.tpl       Prefix=Simple Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl simple/list.tpl       Prefix=Simple Type=string Stringer:true  Comparable:true Ordered:true  String:true  MapTo:int
//go:generate runtemplate -tpl simple/list.tpl       Prefix=Simple Type=Apple  Stringer:false Comparable:true                            MapTo:string
//go:generate runtemplate -tpl simple/set.tpl        Prefix=Simple Type=int    Stringer:true  Ordered:true  Numeric:true                 MapTo:string MapTo:int64
//go:generate runtemplate -tpl simple/set.tpl        Prefix=Simple Type=Apple  Stringer:false                                            MapTo:string
//go:generate runtemplate -tpl simple/map.tpl        Prefix=Simple Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl        Prefix=Simple Key=string Type=Apple                   Stringer:true

// Immutable Examples
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=Immutable Type=int    Stringer:true  Comparable:true Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=Immutable Type=Apple  Stringer:false                                           MapTo:string
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=Immutable Type=int    Stringer:true  Comparable:true Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=Immutable Type=string Stringer:true  Comparable:true Ordered:true String:true  MapTo:int
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=Immutable Type=Apple  Stringer:false GobEncode:true                            MapTo:string
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=Immutable Type=int    Stringer:true                  Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=Immutable Type=Apple  Stringer:false GobEncode:true                            MapTo:string
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=Immutable Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=Immutable Key=string Type=Apple   GobEncode:true

// Fast Examples
//go:generate runtemplate -tpl fast/collection.tpl Prefix=Fast Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=Fast Type=Apple  Stringer:false
//go:generate runtemplate -tpl fast/list.tpl       Prefix=Fast Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl fast/list.tpl       Prefix=Fast Type=string Stringer:true  Comparable:true Ordered:true  String:true  MapTo:int
//go:generate runtemplate -tpl fast/list.tpl       Prefix=Fast Type=Apple  Stringer:false Comparable:true GobEncode:true
//go:generate runtemplate -tpl fast/set.tpl        Prefix=Fast Type=int    Stringer:true                  Ordered:true  Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl fast/set.tpl        Prefix=Fast Type=Apple  Stringer:false  GobEncode:true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=Fast Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=Fast Key=string Type=Apple   GobEncode:true
//go:generate runtemplate -tpl fast/queue.tpl      Prefix=Fast Type=int    MapTo:string MapTo:int64
//go:generate runtemplate -tpl fast/queue.tpl      Prefix=Fast Type=Apple

// Threadsafe Examples
//go:generate runtemplate -tpl threadsafe/collection.tpl Type=int    Stringer:true  Comparable:true Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl threadsafe/collection.tpl Type=Apple  Stringer:false                                           MapTo:string
//go:generate runtemplate -tpl threadsafe/list.tpl       Type=int    Stringer:true  Comparable:true Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl threadsafe/list.tpl       Type=string Stringer:true  Comparable:true Ordered:true  String:true  MapTo:int
//go:generate runtemplate -tpl threadsafe/list.tpl       Type=Apple  Stringer:false Comparable:true GobEncode:true            MapTo:string
//go:generate runtemplate -tpl threadsafe/set.tpl        Type=int    Stringer:true                  Ordered:true Numeric:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl threadsafe/set.tpl        Type=Apple  Stringer:false GobEncode:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Key=string Type=Apple   GobEncode:true
//go:generate runtemplate -tpl threadsafe/queue.tpl      Type=int                                                             MapTo:string MapTo:int64
//go:generate runtemplate -tpl threadsafe/queue.tpl      Type=Apple                                                           MapTo:string

// Other Categories
//go:generate runtemplate -tpl types/stringy.tpl         Type=Email SortableSlice:true
//go:generate runtemplate -tpl plumbing/plumbing.tpl     Type=Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl        Type=Apple ToType=int

// Apple is an empty placeholder used for the examples. "Insert your own type!"
type Apple struct{}
