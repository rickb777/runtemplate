package threadsafetest

//go:generate runtemplate -tpl collections/collection.tpl -output x_string_collection.go Type=string Stringer=true Comparable=true
//go:generate runtemplate -tpl collections/collection.tpl Type=Apple
//go:generate runtemplate -tpl collections/collection.tpl -output x_pear_collection.go Type=Pear

//go:generate runtemplate -tpl threadsafe/set.tpl -output x_string_set.go Type=string Stringer=true

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_string_map.go Type=string Key=string

//go:generate runtemplate -tpl threadsafe/map.tpl -output x_string_apple_map.go Key=string Type=Apple
//go:generate runtemplate -tpl threadsafe/map.tpl -output x_apple_string_map.go Key=Apple Type=string

//go:generate runtemplate -tpl threadsafe/plumbing.tpl -output x_apple_plumbing.go Type=Apple
//go:generate runtemplate -tpl threadsafe/mapTo.tpl -output x_apple_mapTo.go Type=Apple ToType=Pear

type Apple struct{}
type Pear struct{}
