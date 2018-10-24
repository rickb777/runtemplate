// A simple type derived from map[{{.Type}}]struct{}
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Numeric:{{.Numeric}} Stringer:{{.Stringer}} Mutable:always

package {{.Package}}

{{if or .Stringer .HasImport}}
import (
{{if .Stringer}}
	"bytes"
	"encoding/json"
	"fmt"
{{- end}}
{{- if .HasImport}}
	{{.Import}}
{{end -}}
)

{{end -}}
// {{.UPrefix}}{{.UType}}Set is the primary type that represents a set
type {{.UPrefix}}{{.UType}}Set map[{{.Type}}]struct{}

// New{{.UPrefix}}{{.UType}}Set creates and returns a reference to an empty set.
func New{{.UPrefix}}{{.UType}}Set(values ...{{.Type}}) {{.UPrefix}}{{.UType}}Set {
	set := make({{.UPrefix}}{{.UType}}Set)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// Convert{{.UPrefix}}{{.UType}}Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func Convert{{.UPrefix}}{{.UType}}Set(values ...interface{}) ({{.UPrefix}}{{.UType}}Set, bool) {
	set := make({{.UPrefix}}{{.UType}}Set)
{{if and .Numeric (eq .Type .PType)}}
	for _, i := range values {
		switch i.(type) {
		case int:
			set[{{.PType}}(i.(int))] = struct{}{}
		case int8:
			set[{{.PType}}(i.(int8))] = struct{}{}
		case int16:
			set[{{.PType}}(i.(int16))] = struct{}{}
		case int32:
			set[{{.PType}}(i.(int32))] = struct{}{}
		case int64:
			set[{{.PType}}(i.(int64))] = struct{}{}
		case uint:
			set[{{.PType}}(i.(uint))] = struct{}{}
		case uint8:
			set[{{.PType}}(i.(uint8))] = struct{}{}
		case uint16:
			set[{{.PType}}(i.(uint16))] = struct{}{}
		case uint32:
			set[{{.PType}}(i.(uint32))] = struct{}{}
		case uint64:
			set[{{.PType}}(i.(uint64))] = struct{}{}
		case float32:
			set[{{.PType}}(i.(float32))] = struct{}{}
		case float64:
			set[{{.PType}}(i.(float64))] = struct{}{}
		}
	}
{{else}}
	for _, i := range values {
		v, ok := i.({{.PType}})
		if ok {
			set[v] = struct{}{}
		}
	}
{{end}}
	return set, len(set) == len(values)
}

// Build{{.UPrefix}}{{.UType}}SetFromChan constructs a new {{.UPrefix}}{{.UType}}Set from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}SetFromChan(source <-chan {{.PType}}) {{.UPrefix}}{{.UType}}Set {
	set := make({{.UPrefix}}{{.UType}}Set)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set {{.UPrefix}}{{.UType}}Set) ToSlice() []{{.Type}} {
	var s []{{.Type}}
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set {{.UPrefix}}{{.UType}}Set) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v, _ := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set {{.UPrefix}}{{.UType}}Set) Clone() {{.UPrefix}}{{.UType}}Set {
	clonedSet := New{{.UPrefix}}{{.UType}}Set()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set {{.UPrefix}}{{.UType}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set {{.UPrefix}}{{.UType}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set {{.UPrefix}}{{.UType}}Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set {{.UPrefix}}{{.UType}}Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set {{.UPrefix}}{{.UType}}Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set {{.UPrefix}}{{.UType}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set {{.UPrefix}}{{.UType}}Set) Add(i ...{{.Type}}) {{.UPrefix}}{{.UType}}Set {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set {{.UPrefix}}{{.UType}}Set) doAdd(i {{.Type}}) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set {{.UPrefix}}{{.UType}}Set) Contains(i {{.Type}}) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set {{.UPrefix}}{{.UType}}Set) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set {{.UPrefix}}{{.UType}}Set) IsSubset(other {{.UPrefix}}{{.UType}}Set) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set {{.UPrefix}}{{.UType}}Set) IsSuperset(other {{.UPrefix}}{{.UType}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set {{.UPrefix}}{{.UType}}Set) Append(more ...{{.Type}}) {{.UPrefix}}{{.UType}}Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set {{.UPrefix}}{{.UType}}Set) Union(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.UPrefix}}{{.UType}}Set) Intersect(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	intersection := New{{.UPrefix}}{{.UType}}Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.UPrefix}}{{.UType}}Set) Difference(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	differencedSet := New{{.UPrefix}}{{.UType}}Set()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set {{.UPrefix}}{{.UType}}Set) SymmetricDifference(other {{.UPrefix}}{{.UType}}Set) {{.UPrefix}}{{.UType}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *{{.UPrefix}}{{.UType}}Set) Clear() {
	*set = New{{.UPrefix}}{{.UType}}Set()
}

// Remove allows the removal of a single item from the set.
func (set {{.UPrefix}}{{.UType}}Set) Remove(i {{.Type}}) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set {{.UPrefix}}{{.UType}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set {{.UPrefix}}{{.UType}}Set) Forall(fn func({{.Type}}) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set {{.UPrefix}}{{.UType}}Set) Exists(fn func({{.Type}}) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over {{.Type}}Set and executes the passed func against each element.
func (set {{.UPrefix}}{{.UType}}Set) Foreach(fn func({{.Type}})) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new {{.UPrefix}}{{.UType}}Set whose elements return true for func.
// The original set is not modified
func (set {{.UPrefix}}{{.UType}}Set) Filter(fn func({{.Type}}) bool) {{.UPrefix}}{{.UType}}Set {
	result := New{{.UPrefix}}{{.UType}}Set()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new {{.Type}}Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
// The original set is not modified
func (set {{.UPrefix}}{{.UType}}Set) Partition(p func({{.Type}}) bool) ({{.UPrefix}}{{.UType}}Set, {{.UPrefix}}{{.UType}}Set) {
	matching := New{{.UPrefix}}{{.UType}}Set()
	others := New{{.UPrefix}}{{.UType}}Set()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new {{.UPrefix}}{{.UType}}Set by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{.UPrefix}}{{.UType}}Set) Map(fn func({{.PType}}) {{.PType}}) {{.UPrefix}}{{.UType}}Set {
	result := New{{.UPrefix}}{{.UType}}Set()

	for v := range set {
		result[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new {{.UPrefix}}{{.UType}}Set by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{.UPrefix}}{{.UType}}Set) FlatMap(fn func({{.PType}}) []{{.PType}}) {{.UPrefix}}{{.UType}}Set {
	result := New{{.UPrefix}}{{.UType}}Set()

	for v, _ := range set {
		for _, x := range fn(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Set that return true for the passed predicate.
func (set {{.UPrefix}}{{.UType}}Set) CountBy(predicate func({{.Type}}) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

{{if .Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UPrefix}}{{.UType}}Set) Min() {{.PType}} {
	return list.MinBy(func(a {{.PType}}, b {{.PType}}) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list {{.UPrefix}}{{.UType}}Set) Max() (result {{.PType}}) {
	return list.MaxBy(func(a {{.PType}}, b {{.PType}}) bool {
		return a < b
	})
}

{{end -}}
// MinBy returns an element of {{.UPrefix}}{{.UType}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UPrefix}}{{.UType}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m {{.Type}}
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of {{.UPrefix}}{{.UType}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set {{.UPrefix}}{{.UType}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m {{.Type}}
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}

{{if .Numeric}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the set.
func (set {{.UPrefix}}{{.UType}}Set) Sum() {{.Type}} {
	sum := {{.Type}}(0)
	for v, _ := range set {
		sum = sum + {{.TypeStar}}v
	}
	return sum
}

{{end -}}
//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set {{.UPrefix}}{{.UType}}Set) Equals(other {{.UPrefix}}{{.UType}}Set) bool {
	if set.Size() != other.Size() {
		return false
	}
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

func (set {{.UPrefix}}{{.UType}}Set) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set {{.UPrefix}}{{.UType}}Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set {{.UPrefix}}{{.UType}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set {{.UPrefix}}{{.UType}}Set) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set {{.UPrefix}}{{.UType}}Set) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set {{.UPrefix}}{{.UType}}Set) UnmarshalJSON(b []byte) error {
    values := make([]{{.PType}}, 0)
    buf := bytes.NewBuffer(b)
    err := json.NewDecoder(buf).Decode(&values)
    if err != nil {
        return err
    }
    set.Add(values...)
    return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set {{.UPrefix}}{{.UType}}Set) MarshalJSON() ([]byte, error) {
    buf := &bytes.Buffer{}
    err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set {{.UPrefix}}{{.UType}}Set) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v, _ := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
{{end}}
