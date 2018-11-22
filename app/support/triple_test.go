package support

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestPairsKeys(t *testing.T) {
	g := NewGomegaWithT(t)

	triples := Triples([]Triple{{"x", "1", "a"}, {"y", "*Foo", "b"}, {"z", "3", "c"}})
	keys := triples.Keys()
	g.Expect(keys).To(Equal([]string{"x", "y", "z"}))
}

func TestPairsPValues(t *testing.T) {
	g := NewGomegaWithT(t)

	triples := Triples([]Triple{{"x", "1", "a"}, {"y", "*Foo", "b"}, {"z", "3", "c"}})
	keys := triples.PValues()
	g.Expect(keys).To(Equal([]string{"1", "*Foo", "3"}))
}

func TestPairsTValues(t *testing.T) {
	g := NewGomegaWithT(t)

	triples := Triples([]Triple{{"x", "1", "a"}, {"y", "*big.Int", "b"}, {"z", "3", "c"}})
	keys := triples.TValues()
	g.Expect(keys).To(Equal([]string{"1", "bigInt", "3"}))
}
