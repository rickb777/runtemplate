// This package contains example collection types using the simple templates. These are simple types, not structs.
// There is no encapsulation of the underlying data. The type is a reference type (as per the underlying type).
// The types must not be shared between goroutines.
package simple

// Code generation with non-pointer values

//go:generate runtemplate -tpl simple/list.tpl  Prefix=X Type=string  Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X Type=int     Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X Type=Apple   Stringer:false Comparable:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X Type=big.Int Import:"math/big"


//go:generate runtemplate -tpl simple/set.tpl   Prefix=X Type=string  Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X Type=int     Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X Type=Apple   Stringer:false
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X Type=url.URL Stringer:true  Comparable:true Import:"net/url"

//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=int     Type=int       Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=string  Type=string    Comparable:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=string  Type=Apple
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=Apple   Type=string
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=Apple   Type=Pear
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX Key=Apple   Type=big.Int   Import:"math/big"

//go:generate runtemplate -tpl ../list_test.tpl         Type=int Mutable:true M:
//go:generate runtemplate -tpl ../set_test.tpl          Type=int Mutable:true M:
//go:generate runtemplate -tpl ../map_test.tpl  Key=int Type=int Mutable:true M:.m

// Code generation with pointer values

//go:generate runtemplate -tpl simple/list.tpl  Prefix=P Type=*string  Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P Type=*int     Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P Type=*Apple   Stringer:false Comparable:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P Type=*big.Int Import:"math/big"

//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*int    Type=*int      Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*string Type=*string   Comparable:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*string Type=*Apple
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*string Type=*big.Int  Import:"math/big"
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*Apple  Type=*string
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP Key=*Apple  Type=*Pear

type Apple struct{}
type Pear struct{}
