package fast

import (
	"testing"
	"sort"
)

func rangeOf(from, to int32) *XInt32List {
	a := NewXInt32List()
	for i := from; i <= to; i++ {
		a.Add(int32(i))
	}
	return a
}

func TestNewMutableList(t *testing.T) {
	a := NewXInt32List(1, 2, 3)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if a.Get(1) != 2 {
		t.Errorf("Expected 2 but got %d", a.Get(1))
	}

	if a.IsSet() {
		t.Error("Expected not a set")
	}

	if !a.IsSequence() {
		t.Error("Expected a sequence")
	}
}

func TestMutableListAppend(t *testing.T) {
	a := NewXInt32List(1, 2, 3)

	b := a.Append(4, 5).Append(6, 7)

	if a.Size() != 7 {
		t.Errorf("Expected 5 but got %d", a.Size())
	}

	if b != a {
		t.Errorf("Expected b but got %v", a)
	}

	if b.Get(3) != 4 {
		t.Errorf("Expected 4 but got %d", b.Get(3))
	}

	if b.Last() != 7 {
		t.Errorf("Expected 7 but got %d", b.Last())
	}
}

func TestMutableListClone(t *testing.T) {
	a := NewXInt32List(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}

	a.Append(3)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, b.m)
	}

	c := a.Clone().Tail()

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, c.m)
	}
}

func TestMutableListSend(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b := BuildXInt32ListFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}
}

func TestMutableListHeadTailLastInit(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	if a.Head() != 1 {
		t.Errorf("Expected 1 but got %d", a.Head())
	}

	if a.Last() != 4 {
		t.Errorf("Expected 4 but got %d", a.Last())
	}

	tail := a.Tail()
	if !tail.Equals(NewXInt32List(2, 3, 4)) {
		t.Errorf("Expected '2, 3, 4' but got '%+v'", tail.m)
	}

	init := a.Init()
	if !init.Equals(NewXInt32List(1, 2, 3)) {
		t.Errorf("Expected '1, 2, 3' but got '%+v'", init.m)
	}
}

func TestMutableListFilter(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b := a.Filter(func(v int32) bool {
		return v > 2
	})

	if !b.Equals(NewXInt32List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b.m)
	}
}

func TestMutableListPartition(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b, c := a.Partition(func(v int32) bool {
		return v > 2
	})

	if !b.Equals(NewXInt32List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b.m)
	}

	if !c.Equals(NewXInt32List(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c.m)
	}
}

func TestMutableListSort(t *testing.T) {
	a := NewXInt32List(13, 4, 7, -2, 9)

	sort.Sort(a)

	if !a.Equals(NewXInt32List(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", a.m)
	}
}

func TestMutableListReverse(t *testing.T) {
	a := NewXInt32List(13, 4, 7, -2, 9)

	b := a.Reverse()

	if b.Equals(a) {
		t.Errorf("Expected reverse of '%+v' but got '%+v'", a.m, b.m)
	}

	c := b.Reverse()

	if !c.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a.m, c.m)
	}
}

func TestMutableListShuffle(t *testing.T) {
	a := rangeOf(1, 1000)

	b := a.Shuffle()

	if b.Equals(a) {
		t.Errorf("Expected shuffle of '%+v' but got '%+v'", a.m, b.m)
	}

	// prove that the same set of numbers is present
	sort.Sort(b)

	if !b.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a.m, b.m)
	}
}

func TestMutableListTake(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.Take(30)

	if b.Len() != 30 || b.Head() != 1 || b.Last() != 30 {
		t.Errorf("Expected list from 1 to 30 but got '%+v'", b.m)
	}

	c := a.TakeLast(30)

	if c.Len() != 30 || c.Head() != 71 || c.Last() != 100 {
		t.Errorf("Expected list from 71 to 100 but got '%+v'", c.m)
	}

	d := a.Take(101)

	if d.Len() != 100 || d.Head() != 1 || d.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", d.m)
	}

	e := a.TakeLast(101)

	if e.Len() != 100 || e.Head() != 1 || e.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", e.m)
	}
}

func TestMutableListDrop(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.Drop(70)

	if b.Len() != 30 || b.Head() != 71 || b.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b.m)
	}

	c := a.DropLast(75)

	if c.Len() != 25 || c.Head() != 1 || c.Last() != 25 {
		t.Errorf("Expected list from 1 to 25 but got '%+v'", c.m)
	}

	d := a.Drop(101)

	if d.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", d.m)
	}

	e := a.DropLast(101)

	if e.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", e.m)
	}
}

func TestMutableListTakeWhile(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.TakeWhile(func(v int32) bool {
		return v <= 20
	})

	if b.Len() != 20 || b.Head() != 1 || b.Last() != 20 {
		t.Errorf("Expected list from 1 to 20 but got '%+v'", b.m)
	}

	c := a.TakeWhile(func(v int32) bool {
		return true
	})

	if c.Len() != 100 || c.Head() != 1 || c.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b.m)
	}
}

func TestMutableListDropWhile(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.DropWhile(func(v int32) bool {
		return v <= 80
	})

	if b.Len() != 20 || b.Head() != 81 || b.Last() != 100 {
		t.Errorf("Expected list from 81 to 100 but got '%+v'", b.m)
	}

	c := a.DropWhile(func(v int32) bool {
		return true
	})

	if c.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", b.m)
	}
}

func TestMutableListDistinctBy(t *testing.T) {
	a := rangeOf(1, 10)
	b := a.Clone().Append(a.ToSlice()...).Append(a.ToSlice()...)

	if b.Len() != 30 {
		t.Errorf("Expected list of 30 but got '%+v'", b.m)
	}

	c := b.DistinctBy(func(v1, v2 int32) bool {
		return v1 == v2
	})

	if !c.Equals(NewXInt32List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) {
		t.Errorf("Expected 1 to 10 but got '%+v'", c.m)
	}
}

func TestMutableListIndexWhere(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.IndexWhere(func(v int32) bool {
		return v >= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.IndexWhere(func(v int32) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.IndexWhere2(func(v int32) bool {
		return v % 3 == 0
	}, 10)

	if d != 11 {
		t.Errorf("Expected 11 but got %d", d)
	}
}

func TestMutableListLastIndexWhere(t *testing.T) {
	a := rangeOf(1, 100)

	b := a.LastIndexWhere(func(v int32) bool {
		return v <= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.LastIndexWhere(func(v int32) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.LastIndexWhere2(func(v int32) bool {
		return v % 3 == 0
	}, 61)

	if d != 59 {
		t.Errorf("Expected 59 but got %d", d)
	}
}

