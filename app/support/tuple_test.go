package support

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestPairsTValues(t *testing.T) {
	triples := Tuples([]Tuple{NewTuple("a=x/z"), NewTuple("b=*big.Int"), NewTuple("c=interface{}/Any/nil"), NewTuple("d=3/Cho")})

	expect.Slice(triples.TValues()).ToBe(t, "z", "bigInt", "Any", "Cho")

	expect.Bool(triples[0].IsPtr()).ToBeFalse(t)
	expect.String(triples[0].String()).ToBe(t, "x")
	expect.String(triples[0].Name()).ToBe(t, "x")
	expect.String(triples[0].Ident()).ToBe(t, RichString("z"))
	expect.String(triples[0].Zero()).ToBe(t, "*(new(x))")

	expect.Bool(triples[1].IsPtr()).ToBeTrue(t)
	expect.String(triples[1].String()).ToBe(t, "*big.Int")
	expect.String(triples[1].Name()).ToBe(t, "big.Int")
	expect.String(triples[1].Ident()).ToBe(t, RichString("bigInt"))
	expect.String(triples[1].Zero()).ToBe(t, "nil")

	expect.Bool(triples[2].IsPtr()).ToBeFalse(t)
	expect.String(triples[2].String()).ToBe(t, "interface{}")
	expect.String(triples[2].Name()).ToBe(t, "interface{}")
	expect.String(triples[2].Ident()).ToBe(t, RichString("Any"))
	expect.String(triples[2].Zero()).ToBe(t, "nil")
}

func TestNewType(t *testing.T) {
	cases := []struct{ in, exp Type }{
		{in: NewType("/x/y/z"), exp: Type{}},
		{in: NewType(""), exp: Type{}},
		{in: NewType("foo"), exp: Type{s: "foo", ident: "foo"}},
		{in: NewType("*foo"), exp: Type{s: "*foo", ident: "foo"}},
		{in: NewType("*big.Int"), exp: Type{s: "*big.Int", ident: "bigInt"}},
		{in: NewType("*big.Int/Integer"), exp: Type{s: "*big.Int", ident: "Integer"}},
		{in: NewType("*big.Int/Integer/nil"), exp: Type{s: "*big.Int", ident: "Integer", zero: "nil"}},
	}

	for _, c := range cases {
		expect.Any(c.in).ToBe(t, c.exp)
	}
}
