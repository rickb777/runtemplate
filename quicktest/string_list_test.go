/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2016 Rick Beton

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
var _ StringCollection = NewStringList()

//func assertEqual(a, b interface{}, t *testing.T) {
//	if !reflect.DeepEqual(a, b) {
//		_, file, line, _ := runtime.Caller(1)
//		s := strings.LastIndexByte(file, '/')
//		if s >= 0 {
//			file = file[s+1:]
//		}
//		t.Errorf("\n%s:%d\n%v != %v", file, line, a, b)
//	}
//}

func xTest_NewSet(t *testing.T) {
	a := NewStringList()
	assertEqual(a.Size(), 0, t)
}

func Test_NewStringList1(t *testing.T) {
	a := NewStringList()

	assertEqual(a.Len(), 0, t)
	assertEqual(a.Size(), 0, t)
	assertEqual(a.IsEmpty(), true, t)
	assertEqual(a.NonEmpty(), false, t)
}

func Test_NewStringList2(t *testing.T) {
	a := NewStringList("a", "b", "c")

	assertEqual(a.Len(), 3, t)
	assertEqual(a.Size(), 3, t)
	assertEqual(a.IsEmpty(), false, t)
	assertEqual(a.NonEmpty(), true, t)
}

//func xTest_SetEquals(t *testing.T) {
//	a := NewStringList()
//	b := NewStringList()
//
//	if !a.Equals(b) {
//		t.Error("Both a and b are empty sets, and should be equal")
//	}
//
//	a.Add("a")
//
//	if a.Equals(b) {
//		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
//	}
//
//	b.Add("a")
//
//	if !a.Equals(b) {
//		t.Error("a is now equal again to b because both have the item 10 in them")
//	}
//
//	b.Add("b")
//	b.Add("c")
//	b.Add("d")
//
//	if a.Equals(b) {
//		t.Error("b has 3 more elements in it so therefore should not be equal to a")
//	}
//
//	a.Add("d")
//	a.Add("c")
//	a.Add("b")
//
//	if !a.Equals(b) {
//		t.Error("a and b should be equal with the same number of elements")
//	}
//}

//func xTest_SetClone(t *testing.T) {
//	a := NewStringList("a", "b")
//
//	b := a.Clone()
//
//	if !a.Equals(b) {
//		t.Error("Clones should be equal")
//	}
//
//	a.Add("c")
//	if a.Equals(b) {
//		t.Error("a contains one more element, they should not be equal")
//	}
//
//	c := a.Clone()
//	c.Remove("a")
//
//	if a.Equals(c) {
//		t.Error("C contains one element less, they should not be equal")
//	}
//}

func xTest_Iterator(t *testing.T) {
	a := NewStringList("z", "y", "x", "w")

	b := NewStringList()
	for val := range a.Send() {
		b = append(b, val)
	}

	if !a.Equals(b) {
		t.Error("The sets are not equal after iterating through the first set")
	}
}

func xTest_Exists(t *testing.T) {
	a := NewStringList("aa", "ab", "ac", "ad")
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

func xTest_Forall(t *testing.T) {
	a := NewStringList("aa", "ab", "ac", "ad")
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
