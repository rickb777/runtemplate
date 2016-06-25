package threadsafetest

//go:generate runtemplate -tpl threadsafe/set.tpl -output x_string_set.go -type string Stringer=true

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_string_map.go -type string Key=string

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_apple_map.go -type Apple Key=string
//go:generate runtemplate -tpl threadsafe/map.tpl -output x_apple_string_map.go -type string Key=Apple

type Apple struct{}
type Pear struct{}
