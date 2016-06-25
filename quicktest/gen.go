package quicktest

//go:generate runtemplate -tpl collections/set.tpl -output x_string_set.go -type string Stringer=true

//go:generate runtemplate -tpl collections/set.tpl -output x_apple_set.go -type Apple Stringer=true
type Apple struct{}
