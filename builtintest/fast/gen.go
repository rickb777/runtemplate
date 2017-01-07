package fast

// Code generation with non-pointer values

//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=Apple  Stringer=false
//go:generate runtemplate -tpl fast/collection.tpl Prefix=X Type=Pear

//go:generate runtemplate -tpl fast/list.tpl Prefix=X Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl fast/list.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl fast/list.tpl Prefix=X Type=Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl fast/set.tpl Prefix=X Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl fast/set.tpl Prefix=X Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl fast/set.tpl Prefix=X Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl fast/set.tpl Prefix=X Type=Apple  Stringer=false Mutable=true

// Code generation with pointer values

//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*string Stringer=true Comparable=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*Apple  Stringer=false
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P Type=*Pear

//go:generate runtemplate -tpl fast/list.tpl Prefix=P Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl fast/list.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl fast/list.tpl Prefix=P Type=*Apple  Stringer=false Comparable=true

type Apple struct{}
type Pear struct{}

//var _ StringCollection = NewStringList()
//var _ Int32Collection = NewInt32List()
//var _ AppleCollection = NewAppleList()
//
//var _ StringCollection = NewStringSet()
//var _ Int32Collection = NewInt32Set()
//var _ AppleCollection = NewAppleSet()
