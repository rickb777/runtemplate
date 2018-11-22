package support

import (
	"fmt"
	"strings"
)

type Pair struct {
	Key, Val string
}

func NewPair(a string) Pair {
	co := strings.IndexByte(a, ':')
	if co < 0 {
		return Pair{"", ""}
	}
	k, v := a[:co], a[co+1:]
	return Pair{Key: k, Val: v}
}

func (p Pair) Valid() bool {
	return p.Key != ""
}

type Pairs []Pair

//-------------------------------------------------------------------------------------------------

type Type struct {
	Key string
	Val []string
}

func NewType(a string) Type {
	eq := strings.IndexByte(a, '=')
	if eq < 0 {
		return Type{"", nil}
	}
	k, v := a[:eq], a[eq+1:]
	return Type{Key: k, Val: strings.Split(v, "/")}
}

func (t Type) Valid() bool {
	return t.Key != "" && len(t.Val) > 0 && t.Val[0] != ""
}

func (t Type) Ptr() bool {
	return len(t.Val) > 0 && t.Val[0][0] == '*'
}

func (t Type) Elem() string {
	v := t.Val[0]
	if len(v) > 0 && v[0] == '*' {
		return v[1:]
	}
	return v
}

func (t Type) Ident() string {
	if len(t.Val) > 1 {
		return t.Val[1]
	}
	return t.Elem()
}

func (t Type) Zero() string {
	if len(t.Val) > 2 {
		return t.Val[2]
	}
	// this assumes Go code generation
	v := t.Val[0]
	switch v {
	case "string":
		return `""`
	case "bool":
		return `false`
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "byte", "rune":
		return `0`
	case "interface{}":
		return "nil"
	}
	return fmt.Sprintf("*(new(%s))", v)
}

//-------------------------------------------------------------------------------------------------

type Types []Type

func (triples Types) TValues() []string {
	var list []string
	for _, p := range triples {
		list = append(list, RichString(p.Ident()).NoDots().String())
	}
	return list
}
