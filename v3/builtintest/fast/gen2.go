// Generated code - do not alter

// This package contains example collection types using the fast templates.
// Encapsulation of the underlying data is a feature.
package fast

//-------------------------------------------------------------------------------------------------
// Code generation with pointer values

//go:generate runtemplate -tpl fast/collection.tpl Prefix=P1 Type=*string          ToList:true ToSet:true Stringer:true Comparable:true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P1 Type=*int             ToList:true ToSet:true Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P1 Type=*Apple           Stringer:false
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P1 Type=*Pear
//go:generate runtemplate -tpl fast/collection.tpl Prefix=P2 Type=*big.Int/Integer Import:"math/big"

//go:generate runtemplate -tpl fast/list.tpl       Prefix=P1 Type=*string          ToSet:true Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl fast/list.tpl       Prefix=P1 Type=*int             ToSet:true Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl fast/list.tpl       Prefix=P1 Type=*Apple           ToSet:true Stringer:false Comparable:true
//go:generate runtemplate -tpl fast/list.tpl       Prefix=P2 Type=*big.Int/Integer Import:"math/big"

//go:generate runtemplate -tpl fast/set.tpl        Prefix=P1 Type=*string          ToList:true Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl fast/set.tpl        Prefix=P1 Type=*int             ToList:true Stringer:true  Ordered:true  Numeric:true GobEncode:true JsonEncode:true                MapTo:string MapTo:int64
//go:generate runtemplate -tpl fast/set.tpl        Prefix=P1 Type=*Apple           ToList:true Stringer:false

//go:generate runtemplate -tpl fast/queue.tpl      Prefix=P1 Type=*string          ToList:true ToSet:true  Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl fast/queue.tpl      Prefix=P1 Type=*int             ToList:true ToSet:true  Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl fast/queue.tpl      Prefix=P1 Type=*Apple           ToList:true ToSet:true  Stringer:false Comparable:true
//go:generate runtemplate -tpl fast/queue.tpl      Prefix=P2 Type=*big.Int/Integer ToList:true ToSet:false Import:"math/big"

//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*int    Type=*int     Comparable:true Stringer:true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*string Type=*string  Comparable:true
//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*string Type=*Apple
//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*Apple  Type=*string
//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*Apple  Type=*Pear
//go:generate runtemplate -tpl fast/map.tpl        Prefix=TP1 Key=*string Type=*big.Int/Integer Import:"math/big"

//go:generate runtemplate -tpl plumbing/plumbing.tpl     Prefix=P1 Type=*Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl        Prefix=P1 Type=*Apple ToPrefix=P1 ToType=*Pear

//-------------------------------------------------------------------------------------------------

var _ P1StringCollection = NewP1StringList()
var _ P1IntCollection = NewP1IntList()
var _ P1AppleCollection = NewP1AppleList()

var _ P1StringSizer = NewP1StringQueue(1, false)
var _ P1IntCollection = NewP1IntQueue(1, false)
var _ P1AppleSizer = NewP1AppleQueue(1, false)

var _ P1StringCollection = NewP1StringSet()
var _ P1IntCollection = NewP1IntSet()
var _ P1AppleCollection = NewP1AppleSet()
