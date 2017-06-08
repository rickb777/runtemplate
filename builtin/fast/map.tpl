// An encapsulated map[{{.Key}}]{{.Type}}.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Key={{.Key}} Type={{.Type}}
// options: Comparable:{{.Comparable}} Stringer:{{.Stringer}} Mutable:always

package {{.Package}}

{{if or .Stringer .HasImport}}
import (
{{if .Stringer}}
	"bytes"
	"fmt" {{- end}}
{{- if .HasImport}}
    {{.Import}}
{{end}}
)

{{end -}}
// {{.UPrefix}}{{.UKey}}{{.UType}}Map is the primary type that represents a thread-safe map
type {{.UPrefix}}{{.UKey}}{{.UType}}Map struct {
	m map[{{.PKey}}]{{.PType}}
}

// {{.UPrefix}}{{.UKey}}{{.UType}}Tuple represents a key/value pair.
type {{.UPrefix}}{{.UKey}}{{.UType}}Tuple struct {
	Key {{.PKey}}
	Val {{.PType}}
}

// {{.UPrefix}}{{.UKey}}{{.UType}}Tuples can be used as a builder for unmodifiable maps.
type {{.UPrefix}}{{.UKey}}{{.UType}}Tuples []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple

func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Append1(k {{.PKey}}, v {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	return append(ts, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v})
}

func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Append2(k1 {{.PKey}}, v1 {{.PType}}, k2 {{.PKey}}, v2 {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	return append(ts, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k1, v1}, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k2, v2})
}

//-------------------------------------------------------------------------------------------------

func new{{.UPrefix}}{{.UKey}}{{.UType}}Map() {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	return {{.UPrefix}}{{.UKey}}{{.UType}}Map{
		m: make(map[{{.PKey}}]{{.PType}}),
	}
}

// New{{.UPrefix}}{{.UKey}}{{.UType}}Map creates and returns a reference to a map containing one item.
func New{{.UPrefix}}{{.UKey}}{{.UType}}Map1(k {{.PKey}}, v {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	mm := new{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	mm.m[k] = v
	return mm
}

// New{{.UPrefix}}{{.UKey}}{{.UType}}Map creates and returns a reference to a map, optionally containing some items.
func New{{.UPrefix}}{{.UKey}}{{.UType}}Map(kv ...{{.UPrefix}}{{.UKey}}{{.UType}}Tuple) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	mm := new{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Keys() []{{.PKey}} {

	var s []{{.PKey}}
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ToSlice() []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple {

	var s []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple
	for k, v := range mm.m {
		s = append(s, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Get(k {{.PKey}}) ({{.PType}}, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Put(k {{.PKey}}, v {{.PType}}) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ContainsKey(k {{.PKey}}) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ContainsAllKeys(kk ...{{.PKey}}) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *{{.UPrefix}}{{.UKey}}{{.UType}}Map) Clear() {

	mm.m = make(map[{{.PKey}}]{{.PType}})
}

// Remove allows the removal of a single item from the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Remove(k {{.PKey}}) {

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Forall(fn func({{.PKey}}, {{.PType}}) bool) bool {

	for k, v := range mm.m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Exists(fn func({{.PKey}}, {{.PType}}) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Filter(fn func({{.PKey}}, {{.PType}}) bool) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()

	for k, v := range mm.m {
		if fn(k, v) {
			result.m[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Partition(fn func({{.PKey}}, {{.PType}}) bool) (matching {{.UPrefix}}{{.UKey}}{{.UType}}Map, others {{.UPrefix}}{{.UKey}}{{.UType}}Map) {
	matching = New{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	others = New{{.UPrefix}}{{.UKey}}{{.UType}}Map()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

{{if .Comparable}}
// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Equals(other {{.UPrefix}}{{.UKey}}{{.UType}}Map) bool {

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || {{.TypeStar}}v1 != {{.TypeStar}}v2 {
			return false
		}
	}
	return true
}

{{end -}}
// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Clone() {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

{{if .Stringer}}
//-------------------------------------------------------------------------------------------------

func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
{{end}}
