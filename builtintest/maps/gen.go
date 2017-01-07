package maps

// Code generation with non-pointer values

//go:generate runtemplate -tpl maps/simple.tpl Prefix=SX Key=string Type=string Comparable=true Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SX Key=string Type=Apple
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SX Key=Apple  Type=string
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SX Key=Apple  Type=Pear

//go:generate runtemplate -tpl maps/simple.tpl Prefix=SP Key=*string Type=*string Comparable=true Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SP Key=*string Type=*Apple
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SP Key=*Apple  Type=*string
//go:generate runtemplate -tpl maps/simple.tpl Prefix=SP Key=*Apple  Type=*Pear


//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=string Type=string Comparable=true Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=string Type=Apple
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=Apple  Type=string
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=Apple  Type=Pear

//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*string Type=*string Comparable=true Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*string Type=*Apple
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*Apple  Type=*string
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*Apple  Type=*Pear


type Apple struct{}
type Pear struct{}
