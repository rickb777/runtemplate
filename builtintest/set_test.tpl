// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
{{- if .GobEncode}}
    "bytes"
    "encoding/gob"
{{- end}}
    "encoding/json"
	"fmt"
	"testing"
	. "github.com/onsi/gomega"
)

func TestNew{{.UType}}Set(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3)

	g.Expect(a.Size()).To(Equal(3))

	g.Expect(a.IsSet()).To(BeTrue())
	g.Expect(a.IsSequence()).To(BeFalse())
}

func TestNew{{.UType}}SetNoDuplicate(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(7, 5, 3, 7)

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(a.Contains(7)).To(BeTrue())
	g.Expect(a.Contains(5)).To(BeTrue())
	g.Expect(a.Contains(3)).To(BeTrue())
}

func TestConvert{{.UType}}Set(t *testing.T) {
	g := NewGomegaWithT(t)

	a, ok := ConvertX1{{.UType}}Set(1, 5.1, uint8(2), 7, 3)

	g.Expect(ok).To(BeTrue())
	g.Expect(a.Equals(NewX1{{.UType}}Set(1, 5, 2, 7, 3))).To(BeTrue())

    b, ok := ConvertX1{{.UType}}Set(a.ToInterfaceSlice()...)

	g.Expect(ok).To(BeTrue())
	g.Expect(a.Equals(b)).To(BeTrue())
}
{{- if .Mutable}}

func TestMutable{{.UType}}SetCardinality(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set()

	g.Expect(a.Size()).To(Equal(0))

	a.Add(1)

	g.Expect(a.Size()).To(Equal(1))
	g.Expect(a.Cardinality()).To(Equal(1))

	a.Remove(1)

	g.Expect(a.Size()).To(Equal(0))

	a.Add(9)

	g.Expect(a.Size()).To(Equal(1))
}

func TestMutable{{.UType}}SetRemove(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(6, 3, 1)

	a.Remove(3)

	g.Expect(a.Size()).To(Equal(2))
	g.Expect(a.Contains(6) && a.Contains(1)).To(BeTrue())

	a.Remove(6)
	a.Remove(1)

	g.Expect(a.Size()).To(Equal(0))
}
{{- end}}

func Test{{.UType}}SetContainsAll(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(8, 6, 7, 5, 3, 0, 9)

	g.Expect(a.ContainsAll(8, 6, 7, 5, 3, 0, 9)).To(BeTrue())
	g.Expect(a.ContainsAll(8, 6, 11, 5, 3, 0, 9)).To(BeFalse())
}

func Test{{.UType}}SetIsSubset(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 5, 7)
	b := NewX1{{.UType}}Set(3, 5, 7)
	c := NewX1{{.UType}}Set(3, 5, 7, 72)

	g.Expect(b.IsSubset(a)).To(BeTrue())
	g.Expect(c.IsSubset(a)).To(BeFalse())

    // check correct nil handling
    a = nil
	g.Expect(b.IsSubset(a)).To(BeFalse())
	g.Expect(a.IsSubset(b)).To(BeTrue())
}

func Test{{.UType}}SetIsSuperSet(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(9, 5, 2, 1, 11)
	b := NewX1{{.UType}}Set(5, 2, 11)
	c := NewX1{{.UType}}Set(5, 2, 11, 42)

	g.Expect(a.IsSuperset(b)).To(BeTrue())
	g.Expect(a.IsSuperset(c)).To(BeFalse())

    // check correct nil handling
    a = nil
	g.Expect(b.IsSuperset(a)).To(BeTrue())
	g.Expect(a.IsSuperset(b)).To(BeFalse())
}

func Test{{.UType}}SetUnion(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set()

	b := NewX1{{.UType}}Set(1, 2, 3, 4, 5)

	c := a.Union(b)

	g.Expect(c.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5))).To(BeTrue())

	d := NewX1{{.UType}}Set(10, 14, 0)

	e := c.Union(d)
	g.Expect(e.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5, 10, 14, 0))).To(BeTrue())

	a = NewX1{{.UType}}Set(14, 3)

	f := a.Union(e)
	g.Expect(f.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5, 10, 14, 0))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Union(b)
	d = b.Union(a)

	g.Expect(c.Equals(b)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func Test{{.UType}}SetIntersection(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewX1{{.UType}}Set(1, 3, 5, 7)
	a2 := NewX1{{.UType}}Set(1, 3, 5, 7, 10)

	b1 := NewX1{{.UType}}Set(0, 2, 4, 6)
	b2 := NewX1{{.UType}}Set(2, 4, 6, 10)

	c1 := a1.Intersect(a2)
	c2 := b1.Intersect(a1)

	g.Expect(c1.Equals(NewX1IntSet(1, 3, 5, 7))).To(BeTrue())
	g.Expect(c2.NonEmpty()).To(BeFalse())

	d1 := a1.Intersect(a2)
	d2 := b2.Intersect(a2)

	g.Expect(d1.Equals(NewX1IntSet(1, 3, 5, 7))).To(BeTrue(), d1.String())
	g.Expect(d2.Equals(NewX1IntSet(10))).To(BeTrue(), d2.String())

    // check correct nil handling
    a1 = nil
	c1 = a1.Intersect(b1)
	d1 = b1.Intersect(a1)

	g.Expect(c1.NonEmpty()).To(BeFalse())
	g.Expect(d1.NonEmpty()).To(BeFalse())
}

func Test{{.UType}}SetDifference(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3)
	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)
	d := b.Difference(a)

	g.Expect(c.Equals(NewX1IntSet(2))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	g.Expect(c.Equals(a)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func Test{{.UType}}SetSymmetricDifference(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 50)
	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.SymmetricDifference(b)
	d := b.SymmetricDifference(a)

	g.Expect(c.Equals(NewX1IntSet(2, 4, 5, 6, 50, 99))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	g.Expect(c.Equals(a)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func Test{{.UType}}SetEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set()
	b := NewX1{{.UType}}Set()

	g.Expect(a.Equals(b)).To(BeTrue())

	c := NewX1{{.UType}}Set(1, 3, 5, 6, 8)
	d := NewX1{{.UType}}Set(1, 3, 5, 6, 9)

	g.Expect(c.Equals(d)).To(BeFalse())

    // check correct nil handling
    a = nil
    g.Expect(a.Equals(b)).To(BeTrue())
    g.Expect(b.Equals(a)).To(BeTrue())
}
{{- if .Append}}

func Test{{.UType}}SetToList(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	b := a.ToList()

	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.Contains(1)).To(BeTrue())
	g.Expect(a.Contains(2)).To(BeTrue())
	g.Expect(a.Contains(3)).To(BeTrue())
	g.Expect(a.Contains(4)).To(BeTrue())

    c := a.ToSet()
	g.Expect(c.Size()).To(Equal(4))
	g.Expect(c.Contains(1)).To(BeTrue())
	g.Expect(c.Contains(2)).To(BeTrue())
	g.Expect(c.Contains(3)).To(BeTrue())
	g.Expect(c.Contains(4)).To(BeTrue())
{{- if .Mutable}}

	a = nil
	b = a.ToList()

	g.Expect(b.IsEmpty()).To(BeTrue())
{{- end}}
}
{{- end}}
{{- if .Append}}

func Test{{.UType}}SetToSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}List(1, 2, 3, 4)
	b1 := a.ToSlice()
	b2 := a.ToInterfaceSlice()

	g.Expect(b1).To(Equal([]{{.PType}}{1, 2, 3, 4}))
	g.Expect(b2).To(Equal([]interface{}{1, 2, 3, 4}))
{{- if .Mutable}}

	a = nil
	b1 = a.ToSlice()
	b2 = a.ToInterfaceSlice()

	g.Expect(b1).To(HaveLen(0))
	g.Expect(b2).To(HaveLen(0))
{{- end}}
}
{{- end}}

func Test{{.UType}}SetSend(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	b := BuildX1{{.UType}}SetFromChan(a.Send())

	g.Expect(a.Equals(b)).To(BeTrue())
{{- if .Mutable}}

    // check correct nil handling
	a = nil
	b = BuildX1{{.UType}}SetFromChan(a.Send())

	g.Expect(a.Equals(b)).To(BeTrue())
{{- end}}
}

func Test{{.UType}}SetForall(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	found := a.Forall(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeTrue())

	found = a.Forall(func(v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	found = a.Forall(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeTrue())
}

func Test{{.UType}}SetExists(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	found := a.Exists(func(v int) bool {
		return v > 2
	})

	g.Expect(found).To(BeTrue())

	found = a.Exists(func(v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	found = a.Exists(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeFalse())
}

func Test{{.UType}}SetForeach(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	s := 0

	a.Foreach(func(v int) {
		s += v
	})

	g.Expect(s).To(Equal(10))

    // check correct nil handling
    a = nil
	a.Foreach(func(v int) {
		s += v
	})

	g.Expect(s).To(Equal(10))
}

func Test{{.UType}}SetFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1{{.UType}}Set(3, 4))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Filter(func(v int) bool {
		return v > 2
	})
}

func Test{{.UType}}SetPartition(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1{{.UType}}Set(3, 4))).To(BeTrue())
	g.Expect(c.Equals(NewX1{{.UType}}Set(1, 2))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Partition(func(v int) bool {
		return v > 2
	})
}

func Test{{.UType}}SetTransform(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.Map(func(v {{.Type}}) {{.Type}} {
		return v * v
	})

	g.Expect(b.Equals(NewX1{{.UType}}Set(1, 4, 9, 16))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Map(func(v {{.Type}}) {{.Type}} {
		return v * v
	})
}

func Test{{.UType}}SetFlatMap(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.FlatMap(func(v {{.Type}}) []{{.Type}} {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})

    exp := NewX1{{.UType}}Set(2, 3, 4, 6, 6, 9)
	g.Expect(b.Equals(exp)).To(BeTrue())

    // check correct nil handling
	a = nil
	a.FlatMap(func(v {{.Type}}) []{{.Type}} {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})
}

func Test{{.UType}}SetStringMap(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.StringMap()

	for _, c := range a.ToSlice() {
		s := fmt.Sprintf("%d", c)
		_, ok := b[s]
		g.Expect(ok).To(BeTrue())
	}

    // check correct nil handling
	a = nil
	a.StringMap()
}
{{- if .Mutable}}

func TestMutable{{.UType}}SetClear(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(2, 5, 9, 10)

	a.Clear()

	g.Expect(a.Size()).To(Equal(0))

    // check correct nil handling
	a = nil
	a.Clear()
}

func TestMutable{{.UType}}SetClone(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(1, 2)

	b := a.Clone()

	g.Expect(a.Equals(b)).To(BeTrue())

	a.Add(3)
	g.Expect(a.Equals(b)).To(BeFalse())

	c := a.Clone()
	c.Remove(1)

	g.Expect(a.Equals(c)).To(BeFalse())

    // check correct nil handling
	a = nil
	a.Clone()
}
{{- end}}

func Test{{.UType}}SetMkString(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(13, 4)

	c := a.MkString("|")

	g.Expect(c).To(Or(Equal("13|4"), Equal("4|13")))

    // check correct nil handling
	a = nil
	a.MkString("|")
}

func Test{{.UType}}SetMkString3(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(13, 4)

	c := a.MkString3("<", ", ", ">")

	g.Expect(c).To(Or(Equal("<13, 4>"), Equal("<4, 13>")))

    // check correct nil handling
	a = nil
	a.MkString3("<", ",", ">")
}

{{if .GobEncode}}
func Test{{.UType}}SetGobEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(13, 4, 7, -2, 9)
	b := NewX1{{.UType}}Set()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

	g.Expect(err).NotTo(HaveOccurred())

    err = gob.NewDecoder(buf).Decode(&b)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(a.Equals(b)).To(BeTrue())
}

{{end}}
func Test{{.UType}}SetJsonEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Set(13, 4, 7, -2, 9)
	b := NewX1{{.UType}}Set()

    buf, err := json.Marshal(a)

	g.Expect(err).NotTo(HaveOccurred())

    err = json.Unmarshal(buf, &b)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(a.Equals(b)).To(BeTrue())
}
