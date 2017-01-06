package collectiontest1

// Code generation with non-pointer values

//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=Apple  Stringer=false
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=Pear

//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=Apple  Stringer=false Comparable=true

//go:generate runtemplate -tpl collections/set.tpl -prefix x_ Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl collections/set.tpl -prefix x_ Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl collections/set.tpl -prefix x_ Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl collections/set.tpl -prefix x_ Type=Apple  Stringer=false Mutable=true

type Apple struct{}
type Pear struct{}

var _ StringCollection = NewStringList()
var _ Int32Collection = NewInt32List()
var _ AppleCollection = NewAppleList()

var _ StringCollection = NewStringSet()
var _ Int32Collection = NewInt32Set()
var _ AppleCollection = NewAppleSet()
