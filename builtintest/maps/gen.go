package maps

// Code generation with non-pointer values

//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=X Key=string Type=string
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=X Key=string Type=Apple
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=X Key=Apple  Type=string
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=X Key=Apple  Type=Pear


type Apple struct{}
type Pear struct{}
