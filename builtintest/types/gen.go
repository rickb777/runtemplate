// This package contains example collection types using the types templates. These are simple types, not structs.
package types

//go:generate runtemplate -tpl types/stringy.tpl -output x_category_stringy.go Type=Category

type Apple struct{}
type Pear struct{}
