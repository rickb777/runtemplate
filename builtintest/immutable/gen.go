// This package contains example collection types using the non-thread-safe templates. Encapsulation of the underlying data is provided.
package immutable

// Code generation with non-pointer values

//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X1 Type=string Stringer:true Comparable:true
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X1 Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X1 Type=Apple  Stringer:false
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X1 Type=Pear
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X1 Type=string Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X1 Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X1 Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X1 Type=string Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X1 Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X1 Type=Apple  Stringer:false
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X2 Type=url.URL Stringer:true  Comparable:true Import:"net/url"
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X2 Type=testtypes.Email Import:"github.com/rickb777/runtemplate/builtintest/testtypes"

//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX1 Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX1 Key=string Type=string  Comparable:true Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX1 Key=string Type=Apple                   Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX1 Key=Apple  Type=string
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX1 Key=Apple  Type=Pear                    Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX2 Key=Apple  Type=big.Int  Import:"math/big"

//go:generate runtemplate -tpl ../collection_test.tpl   Type=int
//go:generate runtemplate -tpl ../list_test.tpl         Type=int M:.m Append:true
//go:generate runtemplate -tpl ../set_test.tpl          Type=int M:.m Append:true
//go:generate runtemplate -tpl ../map_test.tpl  Key=int Type=int M:.m

// Code generation with pointer values

//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P1 Type=*string Stringer:true Comparable:true
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P1 Type=*int    Stringer:true Comparable:true Ordered:true Numeric:true
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P1 Type=*Apple  Stringer:false
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P1 Type=*Pear
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P2 Type=*big.Int Import:"math/big"

//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P1 Type=*string Stringer:true  Comparable:true Ordered:false Numeric:false
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P1 Type=*int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P1 Type=*Apple  Stringer:false Comparable:true
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P2 Type=*big.Int Import:"math/big"

//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP1 Key=*int    Type=*int     Comparable:true Stringer:true
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP1 Key=*string Type=*string  Comparable:true
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP1 Key=*string Type=*Apple
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP1 Key=*Apple  Type=*string
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP1 Key=*Apple  Type=*Pear


type Apple struct {
	N int
}

type Pear struct{
	K int
}

var _ X1StringCollection = NewX1StringList()
var _ X1IntCollection = NewX1IntList()
var _ X1AppleCollection = NewX1AppleList()

var _ X1StringCollection = NewX1StringSet()
var _ X1IntCollection = NewX1IntSet()
var _ X1AppleCollection = NewX1AppleSet()
