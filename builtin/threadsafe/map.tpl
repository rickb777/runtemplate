// An encapsulated map[{{.Key.Name}}]{{.Type}}.
//{{if .Key.IsPtr}} Note that the api uses {{.Key}} but the map uses {{.Key.Name}} keys.{{end}}
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Key={{.Key}} Type={{.Type}}
// options: Comparable:{{.Comparable}} Stringer:{{.Stringer}} KeyList:{{.KeyList}} ValueList:{{.ValueList}} Mutable:always
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if or .Stringer .GobEncode}}
	"bytes"
{{- end}}
{{- if .GobEncode}}
	"encoding/gob"
{{- end}}
{{- if .Stringer}}
	"encoding/json"
{{- end}}
	"fmt"
{{- if .Stringer}}
	"strings"
{{- end}}
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map is the primary type that represents a thread-safe map
type {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map struct {
	s *sync.RWMutex
	m map[{{.Key.Name}}]{{.Type}}
}

// {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple represents a key/value pair.
type {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple struct {
	Key {{.Key}}
	Val {{.Type}}
}

// {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples can be used as a builder for unmodifiable maps.
type {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples []{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple

// Append1 adds one item.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) Append1(k {{.Key}}, v {{.Type}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	return append(ts, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k, v})
}

// Append2 adds two items.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) Append2(k1 {{.Key}}, v1 {{.Type}}, k2 {{.Key}}, v2 {{.Type}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	return append(ts, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k1, v1}, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k2, v2})
}

// Append3 adds three items.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) Append3(k1 {{.Key}}, v1 {{.Type}}, k2 {{.Key}}, v2 {{.Type}}, k3 {{.Key}}, v3 {{.Type}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	return append(ts, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k1, v1}, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k2, v2}, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k3, v3})
}

// {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Zip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map
// constructor function.
func {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Zip(keys ...{{.Key}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	ts := make({{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Zip.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) Values(values ...{{.Type}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) ToMap() *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	return New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map(ts...)
}

//-------------------------------------------------------------------------------------------------

func new{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map() *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	return &{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map{
		s: &sync.RWMutex{},
		m: make(map[{{.Key.Name}}]{{.Type}}),
	}
}

// New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map1 creates and returns a reference to a map containing one item.
func New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map1(k {{.Key}}, v {{.Type}}) *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	mm := new{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.m[{{.Key.Star}}k] = v
	return mm
}

// New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map creates and returns a reference to a map, optionally containing some items.
func New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map(kv ...{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple) *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	mm := new{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	for _, t := range kv {
		mm.m[{{.Key.Star}}t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Keys() {{if .KeyList}}{{.KeyList}}{{else}}[]{{.Key}}{{end}} {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make({{if .KeyList}}{{.KeyList}}{{else}}[]{{.Key}}{{end}}, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, {{.Key.Amp}}k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Values() {{if .ValueList}}{{.ValueList}}{{else}}[]{{.Type}}{{end}} {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make({{if .ValueList}}{{.ValueList}}{{else}}[]{{.Type}}{{end}}, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) slice() {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	if mm == nil {
		return nil
	}

	s := make({{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{({{.Key.Amp}}k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) ToSlice() {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) OrderedSlice(keys {{if .KeyList}}{{.KeyList}}{{else}}[]{{.Key}}{{end}}) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make({{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[{{.Key.Star}}k]
		if found {
			s = append(s, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Get(k {{.Key}}) ({{.Type}}, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[{{.Key.Star}}k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Put(k {{.Key}}, v {{.Type}}) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[{{.Key.Star}}k]
	mm.m[{{.Key.Star}}k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) ContainsKey(k {{.Key}}) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[{{.Key.Star}}k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) ContainsAllKeys(kk ...{{.Key}}) bool {
	if mm == nil {
		return len(kk) == 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[{{.Key.Name}}]{{.Type}})
	}
}

// Remove a single item from the map.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Remove(k {{.Key}}) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, {{.Key.Star}}k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Pop(k {{.Key}}) ({{.Type}}, bool) {
	if mm == nil {
		return {{.Type.Zero}}, false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[{{.Key.Star}}k]
	delete(mm.m, {{.Key.Star}}k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) DropWhere(fn func({{.Key}}, {{.Type}}) bool) {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make({{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples, 0)
	for k, v := range mm.m {
		if fn({{.Key.Amp}}k, v) {
			removed = append(removed, {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{({{.Key.Amp}}k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Foreach(f func({{.Key}}, {{.Type}})) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		for k, v := range mm.m {
			f({{.Key.Amp}}k, v)
		}
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Forall(p func({{.Key}}, {{.Type}}) bool) bool {
	if mm == nil {
		return true
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !p({{.Key.Amp}}k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Exists(p func({{.Key}}, {{.Type}}) bool) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p({{.Key.Amp}}k, v) {
			return true
		}
	}

	return false
}

// Find returns the first {{.Type}} that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Find(p func({{.Key}}, {{.Type}}) bool) ({{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple, bool) {
	if mm == nil {
		return {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p({{.Key.Amp}}k, v) {
			return {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{({{.Key.Amp}}k), v}, true
		}
	}

	return {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Filter(p func({{.Key}}, {{.Type}}) bool) *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	if mm == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p({{.Key.Amp}}k, v) {
			result.m[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Partition(p func({{.Key}}, {{.Type}}) bool) (matching *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map, others *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) {
	if mm == nil {
		return nil, nil
	}

	matching = New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	others = New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p({{.Key.Amp}}k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new {{.Prefix.U}}{{.Type.U}}Map by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Map(f func({{.Key}}, {{.Type}}) ({{.Key}}, {{.Type}})) *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	if mm == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f({{.Key.Amp}}k1, v1)
		result.m[{{.Key.Star}}k2] = v2
	}

	return result
}

// FlatMap returns a new {{.Prefix.U}}{{.Type.U}}Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) FlatMap(f func({{.Key}}, {{.Type}}) []{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple) *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	if mm == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		ts := f({{.Key.Amp}}k1, v1)
		for _, t := range ts {
			result.m[{{.Key.Star}}t.Key] = t.Val
		}
	}

	return result
}
{{- if .Comparable}}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Equals(other *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

	mm.s.RLock()
	other.s.RLock()
	defer mm.s.RUnlock()
	defer other.s.RUnlock()

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || {{.Type.Star}}v1 != {{.Type.Star}}v2 {
			return false
		}
	}
	return true
}
{{- end}}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) Clone() *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map {
	if mm == nil {
		return nil
	}

	result := New{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
{{- if .HasKeySlice}}
// The map entries are sorted by their keys.{{- end}}
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()
{{if .HasKeyList}}
	keys := make({{.KeyList}}, 0, len(mm.m))
	for k, _ := range mm.m {
		keys  = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm.m[k]
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}
{{else}}
	for k, v := range mm.m {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}
{{end}}
	b.WriteString(after)
	return b
}
{{- if eq .Key.String "string"}}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) UnmarshalJSON(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) MarshalJSON() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return json.Marshal(mm.m)
}
{{- end}}
{{- end}}
{{- if .GobEncode}}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register {{.Type.Name}} with the 'gob' package before this method is used.
func (mm *{{.Prefix.U}}{{.Key.U}}{{.Type.U}}Map) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
{{- end}}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
{{- if .HasKeySlice}}
// The map entries are sorted by their keys.{{- end}}
func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	sep := before
	for _, t := range ts {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", t.Key, equals, t.Val)
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t {{.Prefix.U}}{{.Key.U}}{{.Type.U}}Tuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
{{- end}}
