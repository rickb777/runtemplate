package support

import (
	"testing"
	"reflect"
)

func TestPairsKeys(t *testing.T) {
	pairs := Pairs([]Pair{{"x", "1"}, {"y", "*Foo"}, {"z", "3"}})
	keys := pairs.Keys()
	if !reflect.DeepEqual(keys, []string{"x", "y", "z"}) {
		t.Fatalf("Got %#v", keys)
	}
}

func TestPairsPValues(t *testing.T) {
	pairs := Pairs([]Pair{{"x", "1"}, {"y", "*Foo"}, {"z", "3"}})
	keys := pairs.PValues()
	if !reflect.DeepEqual(keys, []string{"1", "*Foo", "3"}) {
		t.Fatalf("Got %#v", keys)
	}
}

func TestPairsTValues(t *testing.T) {
	pairs := Pairs([]Pair{{"x", "1"}, {"y", "*big.Int"}, {"z", "3"}})
	keys := pairs.TValues()
	if !reflect.DeepEqual(keys, []string{"1", "bigInt", "3"}) {
		t.Fatalf("Got %#v", keys)
	}
}

