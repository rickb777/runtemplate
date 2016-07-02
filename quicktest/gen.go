package quicktest

//go:generate runtemplate -tpl collections/collection.tpl -output x_string_collection.go Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -output x_int32_collection.go Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/collection.tpl Type=*Apple Stringer=false

//go:generate runtemplate -tpl collections/list.tpl -output x_string_list.go Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/list.tpl -output x_int32_list.go Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/list.tpl Type=*Apple Stringer=false Comparable=true

//go:generate runtemplate -tpl collections/set.tpl -output x_string_set.go Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/set.tpl -output x_int32_set.go Type=int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/set.tpl Type=Apple Stringer=false

type Apple struct{}
