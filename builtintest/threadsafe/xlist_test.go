package threadsafe

import (
	"testing"
)

func TestNewList(t *testing.T) {
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

func TestListAppendMutable(t *testing.T) {
	a := NewXInt32List(1, 2, 3)

	b := a.Append(4, 5)

	if a.Size() != 5 {
		t.Errorf("Expected 5 but got %d", a.Size())
	}

	if b != a {
		t.Errorf("Expected b but got %v", a)
	}

	if b.Get(3) != 4 {
		t.Errorf("Expected 4 but got %d", b.Get(3))
	}
}

func TestListAppendImmutable(t *testing.T) {
	a := NewXAppleList(Apple{1}, Apple{2}, Apple{3})

	b := a.Append(Apple{4}, Apple{5})

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if b.Size() != 5 {
		t.Errorf("Expected 5 but got %d", b.Size())
	}

	apple4 := Apple{4}
	if b.Get(3) != apple4 {
		t.Errorf("Expected 4 but got %v", b.Get(3))
	}

}

func TestListClone(t *testing.T) {
	a := NewXInt32List(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	a.Append(3)
	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, b)
	}

	c := a.Clone().Tail()

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, c)
	}
}

func TestListSend(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b := BuildXInt32ListFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}

func TestListHeadTailLastInit(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	if a.Head() != 1 {
		t.Errorf("Expected 1 but got %d", a.Head())
	}

	if a.Last() != 4 {
		t.Errorf("Expected 4 but got %d", a.Last())
	}

	tail := a.Tail()
	if !tail.Equals(NewXInt32List(2, 3, 4)) {
		t.Errorf("Expected '2, 3, 4' but got '%+v'", tail)
	}

	init := a.Init()
	if !init.Equals(NewXInt32List(1, 2, 3)) {
		t.Errorf("Expected '1, 2, 3' but got '%+v'", init)
	}
}

func TestListFilter(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b := a.Filter(func(v int32) bool {
		return v > 2
	})

	if !b.Equals(NewXInt32List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}
}

func TestListPartition(t *testing.T) {
	a := NewXInt32List(1, 2, 3, 4)

	b, c := a.Partition(func(v int32) bool {
		return v > 2
	})

	if !b.Equals(NewXInt32List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	if !c.Equals(NewXInt32List(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c)
	}
}

