package threadsafe

// Code generation with non-pointer values

//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X Type=Apple Stringer=false
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X Type=Pear

//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=X Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=X Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=X Type=Apple Stringer=false Comparable=true

//go:generate runtemplate -tpl threadsafe/set.tpl Prefix=X Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl threadsafe/set.tpl Prefix=X Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl threadsafe/set.tpl Prefix=X Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl threadsafe/set.tpl Prefix=X Type=Apple Stringer=false Mutable=true

//go:generate runtemplate -tpl plumbing/core.tpl  Prefix=X Type=Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl Prefix=X Type=Apple ToPrefix=X ToType=Pear


// Code generation with pointer values

//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P Pfx=x Type=*string Stringer=true Comparable=true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P Type=*Apple Stringer=false
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P Type=*Pear

//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=P Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=P Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl threadsafe/list.tpl Prefix=P Type=*Apple Stringer=false Comparable=true

//go:generate runtemplate -tpl plumbing/core.tpl  Prefix=P Type=*Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl Prefix=P Type=*Apple ToPrefix=P ToType=*Pear


type Apple struct{}
type Pear struct{}

var _ XStringCollection = NewXStringList()
var _ XInt32Collection = NewXInt32List()
var _ XAppleCollection = NewXAppleList()

var _ XStringCollection = NewXStringSet()
var _ XInt32Collection = NewXInt32Set()
var _ XAppleCollection = NewXAppleSet()

var _ PStringCollection = NewPStringList()
var _ PInt32Collection = NewPInt32List()
var _ PAppleCollection = NewPAppleList()
