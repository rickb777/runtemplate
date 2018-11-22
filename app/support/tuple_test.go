package support

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestPairsTValues(t *testing.T) {
	g := NewGomegaWithT(t)

	triples := Types([]Type{NewType("x=1/a"), NewType("y=*big.Int/b"), NewType("z=3/c")})
	keys := triples.TValues()
	g.Expect(keys).To(Equal([]string{"1", "bigInt", "3"}))
}
