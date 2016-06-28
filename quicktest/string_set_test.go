/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com) & 2016 Rick Beton

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package quicktest

import (
	"testing"
)

// type match validation
var _ StringCollection = NewStringSet()

func Test_NewSet(t *testing.T) {
	a := NewStringSet()
	assertEqual(a.Cardinality(), 0, t)
}

func Test_NewStringSet(t *testing.T) {
	a := NewStringSet()

	assertEqual(a.Size(), 0, t)
	assertEqual(a.Cardinality(), 0, t)
	assertEqual(a.IsEmpty(), true, t)
	assertEqual(a.NonEmpty(), false, t)
}

func Test_AddSet(t *testing.T) {
	a := NewStringSet()
	a.Add("a")
	a.Add("b")
	a.Add("c")

	assertEqual(a.Cardinality(), 3, t)
	assertEqual(a.IsEmpty(), false, t)
	assertEqual(a.NonEmpty(), true, t)
}

func Test_AddSetNoDuplicate(t *testing.T) {
	a := NewStringSet()
	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("a")

	assertEqual(a.Cardinality(), 3, t)

	if !(a.Contains("a") && a.Contains("b") && a.Contains("c")) {
		t.Error("AddSetNoDuplicate set should have a, b, c in it.")
	}
}

func Test_RemoveSet(t *testing.T) {
	a := NewStringSet()
	a.Add("a")
	a.Add("b")
	a.Add("c")

	a.Remove("b")

	assertEqual(a.Cardinality(), 2, t)

	if !(a.Contains("a") && a.Contains("c")) {
		t.Error("RemoveSet should have only items a and c in the set")
	}

	a.Remove("a")
	a.Remove("c")

	assertEqual(a.Cardinality(), 0, t)
}

func Test_ContainsSet(t *testing.T) {
	a := NewStringSet("foo")

	if !a.Contains("foo") {
		t.Error("ContainsSet should contain foo")
	}

	a.Remove("foo")

	if a.Contains("foo") {
		t.Error("ContainsSet should not contain foo")
	}

	a.Add("a")
	a.Add("b")
	a.Add("c")

	if !(a.Contains("a") && a.Contains("b") && a.Contains("c")) {
		t.Error("ContainsSet should contain a, b, c")
	}
}

func Test_ContainsMultipleSet(t *testing.T) {
	a := NewStringSet("a", "b", "c", "d", "e", "f", "g")

	if !a.ContainsAll("e", "a", "d", "g", "f", "e", "b") {
		t.Error("ContainsAll should contain all 7")
	}

	if a.ContainsAll("e", "a", "d", "g", "k", "z") {
		t.Error("ContainsAll should not have all of these")
	}
}

func Test_ClearSet(t *testing.T) {
	a := NewStringSet("a", "b", "c")

	assertEqual(a.Cardinality(), 3, t)

	a.Clear()

	assertEqual(a.Cardinality(), 0, t)
}

func Test_SetIsSubset(t *testing.T) {
	a := NewStringSet("a", "b", "c", "d", "e", "f", "g")

	b := NewStringSet()
	b.Add("c")
	b.Add("a")
	b.Add("f")

	if !b.IsSubset(a) {
		t.Error("set b should be a subset of set a")
	}

	b.Add("z")

	if b.IsSubset(a) {
		t.Error("set b should not be a subset of set a")
	}
}

func Test_SetIsSuperSet(t *testing.T) {
	a := NewStringSet("a", "b", "c", "d", "e", "f", "g")

	b := NewStringSet("b", "d", "g")

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}

	b.Add("z")

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b")
	}
}

func Test_SetUnion(t *testing.T) {
	a := NewStringSet()

	b := NewStringSet("a", "b", "c", "d", "e", "f", "g")

	c := a.Union(b)

	assertEqual(c.Cardinality(), 7, t)

	d := NewStringSet("h", "i", "j")

	e := c.Union(d)

	assertEqual(e.Cardinality(), 10, t)

	f := NewStringSet("a", "g", "k", "l")

	g := f.Union(e)

	assertEqual(g.Cardinality(), 12, t)
}

func Test_SetIntersect(t *testing.T) {
	a := NewStringSet("a", "b", "c")

	b := NewStringSet("b", "c")

	c := a.Intersect(b)

	assertEqual(c.Cardinality(), 2, t)

	a.Add("x")
	b.Add("x")

	d := a.Intersect(b)

	assertEqual(d.Cardinality(), 3, t)
}

func Test_SetDifference(t *testing.T) {
	a := NewStringSet("a", "b", "c")

	b := NewStringSet("a", "c", "d", "e", "f", "z")

	c := a.Difference(b)

	if !(c.Cardinality() == 1 && c.Contains("b")) {
		t.Error("unexpected difference " + c.String())
	}
}

func Test_SetSymmetricDifference(t *testing.T) {
	a := NewStringSet("a", "b", "c", "x")

	b := NewStringSet("a", "c", "d", "e", "f", "z")

	c := a.SymmetricDifference(b)

	if !(c.Cardinality() == 6 && c.Contains("b") && c.Contains("d") && c.Contains("e") && c.Contains("f") && c.Contains("x") && c.Contains("z")) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: b, d, e, f, x, z")
	}
}

func Test_SetEquals(t *testing.T) {
	a := NewStringSet()
	b := NewStringSet()

	if !a.Equals(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add("a")

	if a.Equals(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add("a")

	if !a.Equals(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add("b")
	b.Add("c")
	b.Add("d")

	if a.Equals(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add("d")
	a.Add("c")
	a.Add("b")

	if !a.Equals(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_SetClone(t *testing.T) {
	a := NewStringSet("a", "b")

	b := a.Clone()

	if !a.Equals(b) {
		t.Error("Clones should be equal")
	}

	a.Add("c")
	if a.Equals(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove("a")

	if a.Equals(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_SetSend(t *testing.T) {
	a := NewStringSet("z", "y", "x", "w")

	b := NewStringSet()
	for val := range a.Send() {
		b.Add(val)
	}

	if !a.Equals(b) {
		t.Error("The sets are not equal after iterating through the first set")
	}
}

func Test_Exists(t *testing.T) {
	a := NewStringSet("aa", "ab", "ac", "ad")
	e1 := a.Exists(func(s string) bool {
		return s[1] == 'c'
	})
	if !e1 {
		t.Fail()
	}

	e2 := a.Exists(func(s string) bool {
		return s[1] == 'e'
	})
	if e2 {
		t.Fail()
	}
}

func Test_Forall(t *testing.T) {
	a := NewStringSet("aa", "ab", "ac", "ad")
	e1 := a.Forall(func(s string) bool {
		return s[0] == 'a'
	})
	if !e1 {
		t.Fail()
	}

	e2 := a.Forall(func(s string) bool {
		return s[1] == 'a'
	})
	if e2 {
		t.Fail()
	}
}

func Test_ToSlice(t *testing.T) {
	a := NewStringSet("z", "y", "x", "w")
	setAsSlice := a.ToSlice()
	if len(setAsSlice) != a.Size() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !a.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

func Test_StringList(t *testing.T) {
	a := NewStringSet("z", "y", "x", "w")
	list := a.StringList()
	if len(list) != a.Size() {
		t.Errorf("Set length is incorrect: %v", len(list))
	}

	// this is only possible because the set contains strings
	for _, i := range list {
		if !a.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}
