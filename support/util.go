package support

import (
	"fmt"
	"os"
	"strings"
)

var Verbose = false
var Dbg = false

func Fail(args ...interface{}) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func Info(msg string, args ...interface{}) {
	if Verbose {
		fmt.Printf(msg, args...)
	}
}

func Debug(msg string, args ...interface{}) {
	if Dbg {
		fmt.Printf("-- "+msg, args...)
	}
}

type RichString string

func (rs RichString) NoDots() RichString {
	if rs == "" {
		return ""
	}
	s := string(rs)
	d := strings.IndexByte(s, '.')
	for d >= 0 {
		s = s[:d] + s[d+1:]
		d = strings.IndexByte(s, '.')
	}
	return RichString(s)
}

func (rs RichString) FirstUpper() string {
	if rs == "" {
		return ""
	}
	s := string(rs)
	return strings.ToUpper(s[:1]) + s[1:]
}

func (rs RichString) FirstLower() string {
	if rs == "" {
		return ""
	}
	s := string(rs)
	return strings.ToLower(s[:1]) + s[1:]
}

func (rs RichString) DivideOr0(c byte) (string, string) {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s, ""
	}
	return s[:p], s[p+1:]
}

func (rs RichString) DivideOr1(c byte) (string, string) {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return "", s
	}
	return s[:p], s[p+1:]
}

func (rs RichString) RemoveBefore(c byte) string {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s
	}
	return s[p+1:]
}

//-------------------------------------------------------------------------------------------------

func FindTemplateArg(tpl string, args []string) (string, []string) {
	if tpl != "" {
		return tpl, args
	}
	var left []string
	for _, a := range args {
		if strings.HasSuffix(a, ".tpl") {
			tpl = a
		} else {
			left = append(left, a)
		}
	}
	return tpl, left
}

func SplitKeyValArgs(args []string) (Pairs, Pairs, []string) {
	var types []Pair
	var others []Pair
	var leftover []string
	for _, a := range args {
		found := false
		k, v := "", ""
		eq := strings.LastIndexByte(a, '=')
		co := strings.LastIndexByte(a, ':')
		if eq >= 0 {
			k, v = a[:eq], a[eq+1:]
			if k != "" && v != "" {
				p := Pair{a[:eq], a[eq+1:]}
				types = append(types, p)
				found = true
			}
		} else if co >= 0 {
			k, v = a[:co], a[co+1:]
			if k != "" && v != "" {
				p := Pair{a[:co], a[co+1:]}
				others = append(others, p)
				found = true
			}
		}
		if !found {
			leftover = append(leftover, a)
		}
	}
	return Pairs(types), Pairs(others), leftover
}
