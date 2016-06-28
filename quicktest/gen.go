package quicktest

//go:generate runtemplate -tpl collections/collection.tpl -output x_string_collection.go -type string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -output x_int32_collection.go -type int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/collection.tpl -output x_apple_collection.go -type Apple Stringer=true Comparable=true

//go:generate runtemplate -tpl collections/list.tpl -output x_string_list.go -type string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/list.tpl -output x_int32_list.go -type int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/list.tpl -output x_apple_list.go -type Apple Stringer=true Comparable=true

//go:generate runtemplate -tpl collections/set.tpl -output x_string_set.go -type string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/set.tpl -output x_int32_set.go -type int32 Stringer=true Comparable=true Ordered=true Numeric=true
//go:generate runtemplate -tpl collections/set.tpl -output x_apple_set.go -type Apple Stringer=true Comparable=true
type Apple struct{}
