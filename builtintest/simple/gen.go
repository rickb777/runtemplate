package simple

// Code generation with non-pointer values

//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl simple/list.tpl Prefix=X Type=Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl simple/set.tpl Prefix=X Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl simple/set.tpl Prefix=X Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl simple/set.tpl Prefix=X Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl simple/set.tpl Prefix=X Type=Apple  Stringer=false Mutable=true

// Code generation with pointer values

//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl simple/list.tpl Prefix=P Type=*Apple  Stringer=false Comparable=true

type Apple struct{}
type Pear struct{}
