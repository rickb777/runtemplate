// This package contains example collection types using the simple templates. These are simple types, not structs.
// There is no encapsulation of the underlying data. The type is a reference type (as per the underlying type).
// The types must not be shared between goroutines.
package simple

//-------------------------------------------------------------------------------------------------
// Code generation with non-pointer values

//go:generate runtemplate -tpl simple/list.tpl  Prefix=X1 Type=string  ToSet:true   Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X1 Type=int     ToSet:true   Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X1 Type=Apple   ToSet:true   Stringer:false Comparable:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl simple/set.tpl   Prefix=X1 Type=string  ToList:true  Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X1 Type=int     ToList:true  Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X1 Type=Apple   ToList:true  Stringer:false
//go:generate runtemplate -tpl simple/set.tpl   Prefix=X2 Type=url.URL ToList:false Stringer:true  Comparable:true Import:"net/url"

//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=int     Type=int       Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=string  Type=string    Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=string  Type=Apple                     Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=Email   Type=string                  Stringer:true KeySlice:EmailSlice
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=Apple   Type=string
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=Apple   Type=Pear                      Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=TX1 Key=Apple   Type=big.Int   Import:"math/big"

//go:generate runtemplate -tpl types/stringy.tpl Prefix=X1 Type=Email SortableSlice:true

//go:generate runtemplate -tpl ../list_test.tpl         Type=int Mutable:true M:
//go:generate runtemplate -tpl ../set_test.tpl          Type=int Mutable:true M:
//go:generate runtemplate -tpl ../map_test.tpl  Key=int Type=int Mutable:true M:

//-------------------------------------------------------------------------------------------------
// Code generation with pointer values

//go:generate runtemplate -tpl simple/list.tpl  Prefix=P1 Type=*string  Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P1 Type=*int     Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P1 Type=*Apple   Stringer:false Comparable:true
//go:generate runtemplate -tpl simple/list.tpl  Prefix=P2 Type=*big.Int Import:"math/big"

//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*int    Type=*int      Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*string Type=*string   Comparable:true
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*string Type=*Apple
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*string Type=*big.Int  Import:"math/big"
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*Apple  Type=*string
//go:generate runtemplate -tpl simple/map.tpl   Prefix=SP1 Key=*Apple  Type=*Pear

//-------------------------------------------------------------------------------------------------

type Apple struct{}
type Pear struct{}
