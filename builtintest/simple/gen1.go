// Generated code - do not alter

// This package contains example collection types using the simple templates`.
// Encapsulation of the underlying data is a feature.
package simple

//-------------------------------------------------------------------------------------------------
// Code generation with non-pointer values

//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=string          ToList:true ToSet:true Stringer:true Comparable:true
//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=int             ToList:true ToSet:true Stringer:true Comparable:true Ordered:true Integer:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=Apple
//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=Pear
//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=url.URL         ToList:true ToSet:true  Stringer:true Comparable:true Import:"net/url"
//go:generate runtemplate -tpl simple/collection.tpl Prefix=X1 Type=big.Int/Integer ToList:true ToSet:false Stringer:true                 Import:"math/big"

//go:generate runtemplate -tpl simple/list.tpl       Prefix=X1 Type=string          ToSet:true  Stringer:true  Comparable:true Ordered:true  StringLike:false`
//go:generate runtemplate -tpl simple/list.tpl       Prefix=X1 Type=Name            ToSet:true  Stringer:true  Comparable:true Ordered:true  StringLike:true
//go:generate runtemplate -tpl simple/list.tpl       Prefix=X1 Type=int             ToSet:true  Stringer:true  Comparable:true Ordered:true  Integer:true JsonEncode:true MapTo:string MapTo:int64
//go:generate runtemplate -tpl simple/list.tpl       Prefix=X1 Type=url.URL         ToSet:true  Stringer:true  Comparable:true StringParser:url.Parse Import:"net/url"
//go:generate runtemplate -tpl simple/list.tpl       Prefix=X1 Type=big.Int/Integer ToSet:false Stringer:true  Import:"math/big"                                          MapTo:string MapTo:int

//go:generate runtemplate -tpl simple/set.tpl        Prefix=X1 Type=string          ToList:true  Stringer:true  Ordered:true  StringLike:false
//go:generate runtemplate -tpl simple/set.tpl        Prefix=X1 Type=Name            ToList:true  Stringer:true  Ordered:true  StringLike:true
//go:generate runtemplate -tpl simple/set.tpl        Prefix=X1 Type=int             ToList:true  Stringer:true  Ordered:true  Integer:true JsonEncode:true                MapTo:string MapTo:int64
//go:generate runtemplate -tpl simple/set.tpl        Prefix=X1 Type=url.URL         ToList:true  Stringer:true  Comparable:true StringParser:url.Parse Import:"net/url"
//go:generate runtemplate -tpl simple/set.tpl        Prefix=X1 Type=testtypes.Email Import:"github.com/rickb777/runtemplate/v3/builtintest/testtypes"

//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=int             Type=int     Comparable:true Stringer:true Integer:true JsonEncode:true
//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=string          Type=string  Comparable:true Stringer:true
//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=string          Type=Apple                   Stringer:true KeySlice:sort.StringSlice
//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=Email           Type=string                  Stringer:true KeySlice:EmailSlice
//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=Apple           Type=string
//go:generate runtemplate -tpl simple/map.tpl        Prefix=TX1 Key=Apple           Type=Pear                    Stringer:true

//go:generate runtemplate -tpl types/stringy.tpl         Prefix=X1 Type=Email SortableSlice:true
//go:generate runtemplate -tpl plumbing/plumbing.tpl     Prefix=X1 Type=Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl        Prefix=X1 Type=Apple ToPrefix=X1 ToType=Pear

//go:generate runtemplate -tpl ../collection_test.tpl    Type=int Numeric:true Comparable:true
//go:generate runtemplate -tpl ../list_test.tpl          Type=int Numeric:true Comparable:true M:.slice() JsonEncode:true
//go:generate runtemplate -tpl ../set_test.tpl           Type=int Numeric:true Comparable:true M:.slice() JsonEncode:true
//go:generate runtemplate -tpl ../map_test.tpl   Key=int Type=int Numeric:true Comparable:true M:.slice() JsonEncode:true

//-------------------------------------------------------------------------------------------------

type Apple struct {
	N int
}

type Pear struct {
	K int
}

type Name string

var _ X1StringCollection = NewX1StringList()
var _ X1IntCollection = NewX1IntList()
var _ X1UrlURLCollection = NewX1UrlURLList()

var _ X1StringCollection = NewX1StringSet()
var _ X1IntCollection = NewX1IntSet()
var _ X1UrlURLCollection = NewX1UrlURLSet()
