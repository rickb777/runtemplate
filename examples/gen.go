// This package contains example collection types using a selection of built-in templates.
package examples

// Simple Examples
//go:generate runtemplate -f -tpl simple/list.tpl       Prefix=Simple Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl simple/list.tpl       Prefix=Simple Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -f -tpl simple/set.tpl        Prefix=Simple Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl simple/set.tpl        Prefix=Simple Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl simple/map.tpl        Prefix=Simple Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -f -tpl simple/map.tpl        Prefix=Simple Key=string Type=Apple                   Stringer:true

// Threadsafe / Fast Examples
//go:generate runtemplate -f -tpl immutable/collection.tpl Prefix=Immutable Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -f -tpl immutable/collection.tpl Prefix=Immutable Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl immutable/list.tpl       Prefix=Immutable Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl immutable/list.tpl       Prefix=Immutable Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -f -tpl immutable/set.tpl        Prefix=Immutable Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl immutable/set.tpl        Prefix=Immutable Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl immutable/map.tpl        Prefix=Immutable Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -f -tpl immutable/map.tpl        Prefix=Immutable Key=string Type=Apple                   Stringer:true

// Fast Examples
//go:generate runtemplate -f -tpl fast/collection.tpl Prefix=Fast Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -f -tpl fast/collection.tpl Prefix=Fast Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl fast/list.tpl       Prefix=Fast Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl fast/list.tpl       Prefix=Fast Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -f -tpl fast/set.tpl        Prefix=Fast Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl fast/set.tpl        Prefix=Fast Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl fast/map.tpl        Prefix=Fast Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -f -tpl fast/map.tpl        Prefix=Fast Key=string Type=Apple                   Stringer:true

// Threadsafe Examples
//go:generate runtemplate -f -tpl threadsafe/collection.tpl Prefix=Sync Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -f -tpl threadsafe/collection.tpl Prefix=Sync Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl threadsafe/list.tpl       Prefix=Sync Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl threadsafe/list.tpl       Prefix=Sync Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -f -tpl threadsafe/set.tpl        Prefix=Sync Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -f -tpl threadsafe/set.tpl        Prefix=Sync Type=Apple  Stringer:false
//go:generate runtemplate -f -tpl threadsafe/map.tpl        Prefix=Sync Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -f -tpl threadsafe/map.tpl        Prefix=Sync Key=string Type=Apple                   Stringer:true

//go:generate runtemplate -f -tpl types/stringy.tpl Type=Email SortableSlice:true

type Apple struct {
	N int
}
