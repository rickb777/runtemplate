package threadsafetest

//go:generate runtemplate -tpl collections/collection.tpl -output x_string_collection.go -type string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl -output x_apple_collection.go -type Apple
//go:generate runtemplate -tpl collections/collection.tpl -output x_pear_collection.go -type Pear

//go:generate runtemplate -tpl threadsafe/set.tpl -output x_string_set.go -type string Stringer=true

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_string_map.go -type string Key=string

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_apple_map.go -type Apple Key=string
//go:generate runtemplate -tpl threadsafe/map.tpl -output x_apple_string_map.go -type string Key=Apple

//go:generate runtemplate -tpl threadsafe/plumbing.tpl -output x_apple_plumbing.go -type Apple
//go:generate runtemplate -tpl threadsafe/mapTo.tpl -output x_apple_mapTo.go -type Apple ToType=Pear

type Apple struct{}
type Pear struct{}
