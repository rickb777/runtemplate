// This package contains example collection types using the simple templates. These are simple types, not structs.
// There is no encapsulation of the underlying data. The type is a reference type (as per the underlying type).
// The types must not be shared between goroutines.
package simple

// Code generation with non-pointer values

//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl simple/set.tpl  Prefix=X Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl simple/set.tpl  Prefix=X Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl simple/set.tpl  Prefix=X Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl simple/set.tpl  Prefix=X Type=Apple  Stringer=false Mutable=true

//go:generate runtemplate -tpl simple/map.tpl  Prefix=SX Key=int    Type=int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SX Key=string Type=string Mutable=true Comparable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SX Key=string Type=Apple  Mutable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SX Key=Apple  Type=string Mutable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SX Key=Apple  Type=Pear   Mutable=false


// Code generation with pointer values

//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl simple/map.tpl  Prefix=SP Key=*int    Type=*int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SP Key=*string Type=*string Mutable=true Comparable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SP Key=*string Type=*Apple  Mutable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SP Key=*Apple  Type=*string Mutable=true
//go:generate runtemplate -tpl simple/map.tpl  Prefix=SP Key=*Apple  Type=*Pear   Mutable=false

type Apple struct{}
type Pear struct{}
