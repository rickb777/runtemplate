package collectiontest2

// Code generation with pointer values

//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=*string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=*int32  Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=*Apple Stringer=false
//go:generate runtemplate -tpl collections/collection.tpl -prefix x_ Type=*Pear

//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=*string Stringer=true Comparable=true Ordered=false Numeric=false
//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=*int32  Stringer=true Comparable=true Ordered=true  Numeric=true
//go:generate runtemplate -tpl collections/list.tpl -prefix x_ Type=*Apple  Stringer=false Comparable=true

type Apple struct{}
type Pear struct{}
