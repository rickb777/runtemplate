package threadsafetest1

// Code generation with non-pointer values

//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=Apple Stringer=false
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=Pear

//go:generate runtemplate -tpl threadsafe/list.tpl -prefix x_ Type=string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl threadsafe/list.tpl -prefix x_ Type=int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl threadsafe/list.tpl -prefix x_ Type=Apple Stringer=false Comparable=true

//go:generate runtemplate -tpl threadsafe/set.tpl -prefix x_ Type=string Stringer=true Ordered=false Numeric=false Mutable=true
//go:generate runtemplate -tpl threadsafe/set.tpl -prefix x_ Type=int32  Stringer=true Ordered=true  Numeric=true  Mutable=true
//go:generate runtemplate -tpl threadsafe/set.tpl -prefix x_ Type=int64  Stringer=true Ordered=true  Numeric=true  Mutable=false
//go:generate runtemplate -tpl threadsafe/set.tpl -prefix x_ Type=Apple Stringer=false Mutable=true

//go:generate runtemplate -tpl threadsafe/map.tpl -prefix x_ Type=string Key=string
//go:generate runtemplate -tpl threadsafe/map.tpl -prefix x_ Key=string Type=Apple
//go:generate runtemplate -tpl threadsafe/map.tpl -prefix x_ Key=Apple Type=string

//go:generate runtemplate -tpl threadsafe/plumbing.tpl -prefix x_ Type=Apple
//go:generate runtemplate -tpl threadsafe/mapTo.tpl    -prefix x_ Type=Apple ToType=Pear

type Apple struct{}
type Pear struct{}
