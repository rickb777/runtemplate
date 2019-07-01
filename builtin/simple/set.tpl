// A simple type derived from map[{{.Type.Name}}]struct{}
//{{if .Type.IsPtr}} Note that the api uses {{.Type}} but the set uses {{.Key.Name}} keys.{{end}}
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type.Name}}
// options: Numeric:{{.Numeric}} Stringer:{{.Stringer}} Mutable:always
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

{{if or .Stringer .HasImport -}}
import (
{{- if .Stringer}}
	"bytes"
	"encoding/json"
	"fmt"
{{- end}}
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

{{end -}}
// {{.Prefix.U}}{{.Type.U}}Set is the primary type that represents a set
type {{.Prefix.U}}{{.Type.U}}Set map[{{.Type.Name}}]struct{}

// New{{.Prefix.U}}{{.Type.U}}Set creates and returns a reference to an empty set.
func New{{.Prefix.U}}{{.Type.U}}Set(values ...{{.Type}}) {{.Prefix.U}}{{.Type.U}}Set {
	set := make({{.Prefix.U}}{{.Type.U}}Set)
	for _, i := range values {
		set[{{.Type.Star}}i] = struct{}{}
	}
	return set
}

// Convert{{.Prefix.U}}{{.Type.U}}Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func Convert{{.Prefix.U}}{{.Type.U}}Set(values ...interface{}) ({{.Prefix.U}}{{.Type.U}}Set, bool) {
	set := make({{.Prefix.U}}{{.Type.U}}Set)

	for _, i := range values {
		switch j := i.(type) {
{{- if .Numeric}}
		case int:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *int:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case int8:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *int8:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case int16:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *int16:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case int32:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *int32:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case int64:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *int64:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case uint:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *uint:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case uint8:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *uint8:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case uint16:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *uint16:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case uint32:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *uint32:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case uint64:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *uint64:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case float32:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *float32:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
		case float64:
			k := {{.Type.Name}}(j)
			set[k] = struct{}{}
		case *float64:
			k := {{.Type.Name}}(*j)
			set[k] = struct{}{}
{{- else}}
		case {{.Type.Name}}:
			set[j] = struct{}{}
		case *{{.Type.Name}}:
			set[*j] = struct{}{}
{{- end}}
		}
	}

	return set, len(set) == len(values)
}

// Build{{.Prefix.U}}{{.Type.U}}SetFromChan constructs a new {{.Prefix.U}}{{.Type.U}}Set from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.Prefix.U}}{{.Type.U}}SetFromChan(source <-chan {{.Type}}) {{.Prefix.U}}{{.Type.U}}Set {
	set := make({{.Prefix.U}}{{.Type.U}}Set)
	for v := range source {
		set[{{.Type.Star}}v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set {{.Prefix.U}}{{.Type.U}}Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set {{.Prefix.U}}{{.Type.U}}Set) IsSet() bool {
	return true
}

{{- if .ToList}}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set {{.Prefix.U}}{{.Type.U}}Set) ToList() {{.Prefix.U}}{{.Type.U}}List {
	if set == nil {
		return nil
	}

	return {{.Prefix.U}}{{.Type.U}}List(set.ToSlice())
}
{{- end}}

// ToSet returns the set; this is an identity operation in this case.
func (set {{.Prefix.U}}{{.Type.U}}Set) ToSet() {{.Prefix.U}}{{.Type.U}}Set {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set {{.Prefix.U}}{{.Type.U}}Set) ToSlice() []{{.Type}} {
	s := make([]{{.Type}}, 0, len(set))
	for v := range set {
		s = append(s, {{.Type.Amp}}v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set {{.Prefix.U}}{{.Type.U}}Set) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set {{.Prefix.U}}{{.Type.U}}Set) Clone() {{.Prefix.U}}{{.Type.U}}Set {
	clonedSet := New{{.Prefix.U}}{{.Type.U}}Set()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set {{.Prefix.U}}{{.Type.U}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set {{.Prefix.U}}{{.Type.U}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set {{.Prefix.U}}{{.Type.U}}Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set {{.Prefix.U}}{{.Type.U}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set {{.Prefix.U}}{{.Type.U}}Set) Add(more ...{{.Type}}) {{.Prefix.U}}{{.Type.U}}Set {
	for _, v := range more {
		set.doAdd({{.Type.Star}}v)
	}
	return set
}

func (set {{.Prefix.U}}{{.Type.U}}Set) doAdd(i {{.Type.Name}}) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set {{.Prefix.U}}{{.Type.U}}Set) Contains(i {{.Type}}) bool {
	_, found := set[{{.Type.Star}}i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set {{.Prefix.U}}{{.Type.U}}Set) ContainsAll(i ...{{.Type}}) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set {{.Prefix.U}}{{.Type.U}}Set) IsSubset(other {{.Prefix.U}}{{.Type.U}}Set) bool {
	for v := range set {
		if !other.Contains({{.Type.Amp}}v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set {{.Prefix.U}}{{.Type.U}}Set) IsSuperset(other {{.Prefix.U}}{{.Type.U}}Set) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set {{.Prefix.U}}{{.Type.U}}Set) Append(more ...{{.Type.Name}}) {{.Prefix.U}}{{.Type.U}}Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set {{.Prefix.U}}{{.Type.U}}Set) Union(other {{.Prefix.U}}{{.Type.U}}Set) {{.Prefix.U}}{{.Type.U}}Set {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set {{.Prefix.U}}{{.Type.U}}Set) Intersect(other {{.Prefix.U}}{{.Type.U}}Set) {{.Prefix.U}}{{.Type.U}}Set {
	intersection := New{{.Prefix.U}}{{.Type.U}}Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains({{.Type.Amp}}v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains({{.Type.Amp}}v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.Prefix.U}}{{.Type.U}}Set) Difference(other {{.Prefix.U}}{{.Type.U}}Set) {{.Prefix.U}}{{.Type.U}}Set {
	differencedSet := New{{.Prefix.U}}{{.Type.U}}Set()
	for v := range set {
		if !other.Contains({{.Type.Amp}}v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set {{.Prefix.U}}{{.Type.U}}Set) SymmetricDifference(other {{.Prefix.U}}{{.Type.U}}Set) {{.Prefix.U}}{{.Type.U}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *{{.Prefix.U}}{{.Type.U}}Set) Clear() {
	*set = New{{.Prefix.U}}{{.Type.U}}Set()
}

// Remove a single item from the set.
func (set {{.Prefix.U}}{{.Type.U}}Set) Remove(i {{.Type}}) {
	delete(set, {{.Type.Star}}i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set {{.Prefix.U}}{{.Type.U}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		for v := range set {
			ch <- {{.Type.Amp}}v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function p to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set {{.Prefix.U}}{{.Type.U}}Set) Forall(p func({{.Type}}) bool) bool {
	for v := range set {
		if !p({{.Type.Amp}}v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set {{.Prefix.U}}{{.Type.U}}Set) Exists(p func({{.Type}}) bool) bool {
	for v := range set {
		if p({{.Type.Amp}}v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set {{.Prefix.U}}{{.Type.U}}Set) Foreach(f func({{.Type}})) {
	for v := range set {
		f({{.Type.Amp}}v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type.Name}} that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set {{.Prefix.U}}{{.Type.U}}Set) Find(p func({{.Type}}) bool) ({{.Type}}, bool) {

	for v := range set {
		if p({{.Type.Amp}}v) {
			return {{.Type.Amp}}v, true
		}
	}
{{- if .Type.IsPtr}}

	return nil, false
{{- else}}

	var empty {{.Type.Name}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.Prefix.U}}{{.Type.U}}Set whose elements return true for the predicate p.
//
// The original set is not modified
func (set {{.Prefix.U}}{{.Type.U}}Set) Filter(p func({{.Type}}) bool) {{.Prefix.U}}{{.Type.U}}Set {
	result := New{{.Prefix.U}}{{.Type.U}}Set()
	for v := range set {
		if p({{.Type.Amp}}v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new {{.Prefix.U}}{{.Type.U}}Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set {{.Prefix.U}}{{.Type.U}}Set) Partition(p func({{.Type}}) bool) ({{.Prefix.U}}{{.Type.U}}Set, {{.Prefix.U}}{{.Type.U}}Set) {
	matching := New{{.Prefix.U}}{{.Type.U}}Set()
	others := New{{.Prefix.U}}{{.Type.U}}Set()
	for v := range set {
		if p({{.Type.Amp}}v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new {{.Prefix.U}}{{.Type.U}}Set by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{.Prefix.U}}{{.Type.U}}Set) Map(f func({{.Type}}) {{.Type}}) {{.Prefix.U}}{{.Type.U}}Set {
	result := New{{.Prefix.U}}{{.Type.U}}Set()

	for v := range set {
		k := f({{.Type.Amp}}v)
		result[{{.Type.Star}}k] = struct{}{}
	}

	return result
}
{{- range .MapTo}}

// MapTo{{.U}} returns a new []{{.}} by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{$.Prefix.U}}{{$.Type.U}}Set) MapTo{{.U}}(f func({{$.Type}}) {{.}}) []{{.}} {
	if set == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(set))
	for v := range set {
		result = append(result, f({{$.Type.Amp}}v))
	}

	return result
}
{{- end}}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}Set by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{.Prefix.U}}{{.Type.U}}Set) FlatMap(f func({{.Type}}) []{{.Type}}) {{.Prefix.U}}{{.Type.U}}Set {
	result := New{{.Prefix.U}}{{.Type.U}}Set()

	for v := range set {
		for _, x := range f({{.Type.Amp}}v) {
			result[{{.Type.Star}}x] = struct{}{}
		}
	}

	return result
}
{{- range .MapTo}}

// FlatMapTo{{.U}} returns a new []{{.}} by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set {{$.Prefix.U}}{{$.Type.U}}Set) FlatMapTo{{.U}}(f func({{$.Type}}) []{{.}}) []{{.}} {
	if set == nil {
		return nil
	}

	result := make([]{{.}}, 0, len(set))
	for v := range set {
		result = append(result, f({{$.Type.Amp}}v)...)
	}

	return result
}
{{- end}}

// CountBy gives the number elements of {{.Prefix.U}}{{.Type.U}}Set that return true for the predicate p.
func (set {{.Prefix.U}}{{.Type.U}}Set) CountBy(p func({{.Type}}) bool) (result int) {
	for v := range set {
		if p({{.Type.Amp}}v) {
			result++
		}
	}
	return
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set {{.Prefix.U}}{{.Type.U}}Set) Min() {{.Type.Name}} {
	v := set.MinBy(func(a {{.Type}}, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
	return {{.Type.Star}}v
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set {{.Prefix.U}}{{.Type.U}}Set) Max() {{.Type.Name}} {
	v := set.MaxBy(func(a {{.Type}}, b {{.Type}}) bool {
		return {{.Type.Star}}a < {{.Type.Star}}b
	})
	return {{.Type.Star}}v
}
{{- end}}

// MinBy returns an element of {{.Prefix.U}}{{.Type.U}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set {{.Prefix.U}}{{.Type.U}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m {{.Type.Name}}
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less({{.Type.Amp}}v, {{.Type.Amp}}m) {
			m = v
		}
	}
	return {{.Type.Amp}}m
}

// MaxBy returns an element of {{.Prefix.U}}{{.Type.U}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set {{.Prefix.U}}{{.Type.U}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m {{.Type.Name}}
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less({{.Type.Amp}}m, {{.Type.Amp}}v) {
			m = v
		}
	}
	return {{.Type.Amp}}m
}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type.Name}} is numeric.

// Sum returns the sum of all the elements in the set.
func (set {{.Prefix.U}}{{.Type.U}}Set) Sum() {{.Type.Name}} {
	sum := {{.Type.Name}}(0)
	for v := range set {
		sum = sum + v
	}
	return sum
}
{{- end}}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set {{.Prefix.U}}{{.Type.U}}Set) Equals(other {{.Prefix.U}}{{.Type.U}}Set) bool {
	if set.Size() != other.Size() {
		return false
	}

	for v := range set {
		if !other.Contains({{.Type.Amp}}v) {
			return false
		}
	}

	return true
}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set {{.Prefix.U}}{{.Type.U}}Set) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set {{.Prefix.U}}{{.Type.U}}Set) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set {{.Prefix.U}}{{.Type.U}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set {{.Prefix.U}}{{.Type.U}}Set) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set {{.Prefix.U}}{{.Type.U}}Set) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (set {{.Prefix.U}}{{.Type.U}}Set) UnmarshalJSON(b []byte) error {
	values := make([]{{.Type}}, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set {{.Prefix.U}}{{.Type.U}}Set) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set {{.Prefix.U}}{{.Type.U}}Set) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
{{- end}}
