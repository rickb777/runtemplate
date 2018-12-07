package support

import "strings"

type RichString string

func (rs RichString) String() string {
	return string(rs)
}

// IndexByte returns the index of the first instance of c in rs, or -1 if c is not present in rs.
func (rs RichString) IndexByte(b byte) int {
	return strings.IndexByte(string(rs), b)
}

// LastIndexByte returns the index of the last instance of c in rs, or -1 if c is not present in rs.
func (rs RichString) LastIndexByte(b byte) int {
	return strings.LastIndexByte(string(rs), b)
}

// NoDots returns a copy of rs in which all dots have been removed.
func (rs RichString) NoDots() RichString {
	if rs == "" {
		return ""
	}
	d := rs.IndexByte('.')
	for d >= 0 {
		if d < len(rs) {
			rs = rs[:d] + rs[d+1:]
			d = rs.IndexByte('.')
		} else {
			rs = rs[:d]
			d = -1
		}
	}
	return rs
}

// ToUpper returns a copy of the string rs with all Unicode letters mapped to their upper case.
func (rs RichString) ToUpper() RichString {
	return RichString(strings.ToUpper(string(rs)))
}

// ToLower returns a copy of the string rs with all Unicode letters mapped to their lower case.
func (rs RichString) ToLower() RichString {
	return RichString(strings.ToLower(string(rs)))
}

// ToTitle returns a copy of the string rs with all Unicode letters mapped to their title case.
func (rs RichString) ToTitle() RichString {
	return RichString(strings.ToTitle(string(rs)))
}

// UL is an alias for FirstUpper.
func (rs RichString) U() RichString {
	return rs.FirstUpper()
}

// FirstUpper returns a copy of the string rs with the first letter mapped to upper case.
func (rs RichString) FirstUpper() RichString {
	if rs == "" {
		return ""
	}
	return rs[:1].ToUpper() + rs[1:]
}

// L is an alias for FirstLower
func (rs RichString) L() RichString {
	return rs.FirstLower()
}

// FirstLower returns a copy of the string rs with the first letter mapped to lower case.
func (rs RichString) FirstLower() RichString {
	if rs == "" {
		return ""
	}
	return rs[:1].ToLower() + rs[1:]
}

func (rs RichString) DivideLastOr0(c byte) (RichString, RichString) {
	p := rs.LastIndexByte(c)
	if p < 0 {
		return rs, ""
	}
	return rs[:p], rs[p+1:]
}

func (rs RichString) DivideLastOr1(c byte) (RichString, RichString) {
	p := rs.LastIndexByte(c)
	if p < 0 {
		return "", rs
	}
	return rs[:p], rs[p+1:]
}

// RemoveBeforeLast finds the last occurrence of byte c and retains everything after it.
func (rs RichString) RemoveBeforeLast(c byte) RichString {
	_, rem := rs.DivideLastOr1(c)
	return rem
}
