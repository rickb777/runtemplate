// This package contains example collection types using the non-thread-safe templates. Encapsulation of the underlying data is provided.
package fast

// Code generation with non-pointer values

//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=Apple  Stringer=false
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=Pear

//go:generate runtemplate -tpl fast/list.tpl       Prefix=X Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl fast/list.tpl       Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl fast/list.tpl       Prefix=X Type=Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl fast/set.tpl        Prefix=X Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl fast/set.tpl        Prefix=X Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl fast/set.tpl        Prefix=X Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl fast/set.tpl        Prefix=X Type=Apple  Stringer=false Mutable=true

//go:generate runtemplate -tpl fast/map.tpl        Prefix=X Key=int    Type=int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=X Key=string Type=string Mutable=true Comparable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=X Key=string Type=Apple  Mutable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=X Key=Apple  Type=string Mutable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=X Key=Apple  Type=Pear   Mutable=false


// Code generation with pointer values

//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*string Stringer=true Comparable=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*Apple  Stringer=false
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*Pear

//go:generate runtemplate -tpl fast/list.tpl       Prefix=P Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl fast/list.tpl       Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl fast/list.tpl       Prefix=P Type=*Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl fast/map.tpl        Prefix=P Key=*int    Type=*int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=P Key=*string Type=*string Mutable=true Comparable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=P Key=*string Type=*Apple  Mutable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=P Key=*Apple  Type=*string Mutable=true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=P Key=*Apple  Type=*Pear   Mutable=false


type Apple struct{}
type Pear struct{}

//var _ StringCollection = NewStringList()
//var _ Int32Collection = NewInt32List()
//var _ AppleCollection = NewAppleList()
//
//var _ StringCollection = NewStringSet()
//var _ Int32Collection = NewInt32Set()
//var _ AppleCollection = NewAppleSet()
