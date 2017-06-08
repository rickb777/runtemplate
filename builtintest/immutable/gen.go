// This package contains example collection types using the non-thread-safe templates. Encapsulation of the underlying data is provided.
package immutable

// Code generation with non-pointer values

//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X Type=string Stringer:true Comparable:true
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X Type=Apple  Stringer:false
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X Type=Pear
//go:generate runtemplate -tpl immutable/collection.tpl Prefix=X Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X Type=string Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -tpl immutable/list.tpl       Prefix=X Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X Type=string Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X Type=int    Stringer:true  Ordered:true  Numeric:true
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X Type=Apple  Stringer:false
//go:generate runtemplate -tpl immutable/set.tpl        Prefix=X Type=url.URL Stringer:true  Comparable:true Import:"net/url"

//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=int    Type=int     Comparable:true Stringer:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=string Type=string  Comparable:true
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=string Type=Apple
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=Apple  Type=string
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=Apple  Type=Pear
//go:generate runtemplate -tpl immutable/map.tpl        Prefix=TX Key=Apple  Type=big.Int  Import:"math/big"


// Code generation with pointer values

//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P Type=*string Stringer:true Comparable:true
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P Type=*int    Stringer:true Comparable:true Ordered:true Numeric:true
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P Type=*Apple  Stringer:false
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P Type=*Pear
//zz:generate runtemplate -tpl immutable/collection.tpl Prefix=P Type=*big.Int Import:"math/big"

//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P Type=*string Stringer:true  Comparable:true Ordered:false Numeric:false
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P Type=*int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P Type=*Apple  Stringer:false Comparable:true
//zz:generate runtemplate -tpl immutable/list.tpl       Prefix=P Type=*big.Int Import:"math/big"

//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP Key=*int    Type=*int     Comparable:true Stringer:true
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP Key=*string Type=*string  Comparable:true
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP Key=*string Type=*Apple
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP Key=*Apple  Type=*string
//zz:generate runtemplate -tpl immutable/map.tpl        Prefix=TP Key=*Apple  Type=*Pear


type Apple struct {
	N int
}

type Pear struct{
	K int
}

var _ XStringCollection = NewXStringList()
var _ XIntCollection = NewXIntList()
var _ XAppleCollection = NewXAppleList()

var _ XStringCollection = NewXStringSet()
var _ XIntCollection = NewXIntSet()
var _ XAppleCollection = NewXAppleSet()
