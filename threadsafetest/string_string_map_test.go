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

package threadsafetest

import (
	"testing"
)

//func assertEqual(a, b interface{}, t *testing.T) {
//	if !reflect.DeepEqual(a, b) {
//		_, file, line, _ := runtime.Caller(1)
//		s := strings.LastIndexByte(file, '/')
//		if s >= 0 {
//			file = file[s + 1:]
//		}
//		t.Errorf("\n%s:%d\n%v != %v", file, line, a, b)
//	}
//}

func Test_NewStringStringMap(t *testing.T) {
	a := NewStringStringMap()

	assertEqual(a.Size(), 0, t)
	assertEqual(a.Len(), 0, t)
	assertEqual(a.IsEmpty(), true, t)
	assertEqual(a.NonEmpty(), false, t)
}

func Test_MapPut(t *testing.T) {
	a := NewStringStringMap()
	a.Put("a", "z")
	a.Put("b", "y")
	a.Put("c", "x")

	assertEqual(a.Len(), 3, t)
	assertEqual(a.IsEmpty(), false, t)
	assertEqual(a.NonEmpty(), true, t)
}

func Test_RemoveFromMap(t *testing.T) {
	a := NewStringStringMap()
	a.Put("a", "z")
	a.Put("b", "y")
	a.Put("c", "x")

	a.Remove("b")

	assertEqual(a.Len(), 2, t)

	if !(a.ContainsKey("a") && a.ContainsKey("c")) {
		t.Error("RemoveSet should have only items a and c in the set")
	}

	a.Remove("a")
	a.Remove("c")

	assertEqual(a.Len(), 0, t)
}

func Test_MapContainsKey(t *testing.T) {
	a := NewStringStringMap(StringStringTuple{"foo", "1"})

	if !a.ContainsKey("foo") {
		t.Error("ContainsSet should contain foo")
	}

	a.Remove("foo")

	if a.ContainsKey("foo") {
		t.Error("ContainsSet should not contain foo")
	}

	a.Put("a", "z")
	a.Put("b", "y")
	a.Put("c", "x")

	if !(a.ContainsKey("a") && a.ContainsKey("b") && a.ContainsKey("c")) {
		t.Error("ContainsSet should contain a, b, c")
	}
}

func Test_MapContainsMultipleSet(t *testing.T) {
	a := NewStringStringMap()
	a.Put("a", "z")
	a.Put("b", "y")
	a.Put("c", "x")
	a.Put("d", "w")
	a.Put("e", "v")
	a.Put("f", "u")
	a.Put("g", "t")

	if !a.ContainsAllKeys("e", "a", "d", "g", "f", "e", "b") {
		t.Error("ContainsAll should contain all 7")
	}

	if a.ContainsAllKeys("e", "a", "d", "g", "k", "z") {
		t.Error("ContainsAll should not have all of these")
	}
}

func Test_MapClearSet(t *testing.T) {
	a := NewStringStringMap()
	a.Put("a", "z")
	a.Put("b", "y")
	a.Put("c", "x")

	assertEqual(a.Len(), 3, t)

	a.Clear()

	assertEqual(a.Len(), 0, t)
}

func Test_MapSetEquals(t *testing.T) {
	a := NewStringStringMap()
	b := NewStringStringMap()

	if !a.Equals(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Put("a", "z")

	if a.Equals(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Put("a", "z")

	if !a.Equals(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Put("b", "y")
	b.Put("c", "x")
	b.Put("d", "w")

	if a.Equals(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Put("d", "w")
	a.Put("c", "x")
	a.Put("b", "y")

	if !a.Equals(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_MapSetClone(t *testing.T) {
	a := NewStringStringMap()
	a.Put("a", "z")
	a.Put("b", "y")

	b := a.Clone()

	if !a.Equals(b) {
		t.Error("Clones should be equal")
	}

	a.Put("c", "x")
	if a.Equals(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove("a")

	if a.Equals(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_MapExists(t *testing.T) {
	a := NewStringStringMap()
	a.Put("aa", "z")
	a.Put("ab", "y")
	a.Put("ac", "x")
	a.Put("ad", "w")

	e1 := a.Exists(func(k, v string) bool {
		return k[1] == 'c'
	})
	if !e1 {
		t.Fail()
	}

	e2 := a.Exists(func(k, v string) bool {
		return k[1] == 'e'
	})
	if e2 {
		t.Fail()
	}
}

func Test_MapForall(t *testing.T) {
	a := NewStringStringMap()
	a.Put("aa", "z")
	a.Put("ab", "y")
	a.Put("ac", "x")
	a.Put("ad", "w")
	e1 := a.Forall(func(k, v string) bool {
		return k[0] == 'a'
	})
	if !e1 {
		t.Fail()
	}

	e2 := a.Forall(func(k, v string) bool {
		return k[1] == 'a'
	})
	if e2 {
		t.Fail()
	}
}

func Test_MapToSlice(t *testing.T) {
	a := NewStringStringMap()
	a.Put("aa", "z")
	a.Put("ab", "y")
	a.Put("ac", "x")
	a.Put("ad", "w")
	setAsSlice := a.ToSlice()
	if len(setAsSlice) != a.Size() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !a.ContainsKey(i.Key) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}
