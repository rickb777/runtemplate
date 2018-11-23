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

func noStar(s string) string {
	if s != "" && s[0] == '*' {
		return s[1:]
	}
	return s
}

type Type struct {
	s     string
	ident RichString
	zero  string
}

func NewType(s string) Type {
	if strings.HasPrefix(s, "/") {
		return Type{}
	}
	ss := strings.Split(s, "/")
	switch len(ss) {
	case 0:
		return Type{}
	case 1:
		return Type{s: ss[0], ident: RichString(noStar(ss[0])).NoDots()}
	case 2:
		return Type{s: ss[0], ident: RichString(ss[1]).NoDots()}
	}
	return Type{s: ss[0], ident: RichString(ss[1]), zero: ss[2]}
}

func (t Type) NonBlank() bool {
	return t.s != ""
}

func (t Type) IsPtr() bool {
	return len(t.s) > 0 && t.s[0] == '*'
}

func (t Type) P() string {
	return t.s
}

func (t Type) String() string {
	if t.IsPtr() {
		return t.s[1:]
	}
	return t.s
}

func (t Type) Star() string {
	if t.IsPtr() {
		return "*"
	}
	return ""
}

func (t Type) Amp() string {
	if t.IsPtr() {
		return "&"
	}
	return ""
}

func (t Type) Zero() string {
	if t.zero != "" {
		return t.zero
	}

	if t.IsPtr() {
		return "nil"
	}

	// this assumes Go code generation
	switch t.s {
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
	return fmt.Sprintf("*(new(%s))", t.s)
}

func (t Type) Ident() RichString {
	return t.ident
}

//-------------------------------------------------------------------------------------------------

type Tuple struct {
	Key string
	Type
}

func NewTuple(a string) Tuple {
	eq := strings.IndexByte(a, '=')
	if eq < 0 {
		return Tuple{}
	}
	k, v := a[:eq], a[eq+1:]
	return Tuple{Key: k, Type: NewType(v)}
}

func (t Tuple) Valid() bool {
	return t.Key != "" && t.NonBlank()
}

//-------------------------------------------------------------------------------------------------

type Tuples []Tuple

func (triples Tuples) TValues() []string {
	var list []string
	for _, p := range triples {
		list = append(list, p.Ident().NoDots().String())
	}
	return list
}
