package maps

// Code generation with non-pointer values

//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SX Key=int    Type=int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SX Key=string Type=string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SX Key=string Type=Apple  Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SX Key=Apple  Type=string Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SX Key=Apple  Type=Pear   Mutable=false

//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EX Key=int    Type=int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EX Key=string Type=string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EX Key=string Type=Apple  Mutable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EX Key=Apple  Type=string Mutable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EX Key=Apple  Type=Pear   Mutable=false

//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=int    Type=int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=string Type=string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=string Type=Apple  Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=Apple  Type=string Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TX Key=Apple  Type=Pear   Mutable=false


// Code generation with pointer values

//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SP Key=*int    Type=*int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SP Key=*string Type=*string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SP Key=*string Type=*Apple  Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SP Key=*Apple  Type=*string Mutable=true
//go:generate runtemplate -tpl maps/simple.tpl     Prefix=SP Key=*Apple  Type=*Pear   Mutable=false

//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EP Key=*int    Type=*int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EP Key=*string Type=*string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EP Key=*string Type=*Apple  Mutable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EP Key=*Apple  Type=*string Mutable=true
//go:generate runtemplate -tpl maps/encap.tpl      Prefix=EP Key=*Apple  Type=*Pear   Mutable=false

//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*int    Type=*int    Mutable=true Comparable=true Stringer=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*string Type=*string Mutable=true Comparable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*string Type=*Apple  Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*Apple  Type=*string Mutable=true
//go:generate runtemplate -tpl maps/threadsafe.tpl Prefix=TP Key=*Apple  Type=*Pear   Mutable=false


type Apple struct{}
type Pear struct{}
