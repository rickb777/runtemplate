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
		fmt.Printf("-- " + msg, args...)
	}
}

type RichString string

func (rs RichString) FirstUpper() string {
	s := string(rs)
	return strings.ToUpper(s[:1]) + s[1:]
}

func (rs RichString) FirstLower() string {
	s := string(rs)
	return strings.ToLower(s[:1]) + s[1:]
}

func (rs RichString) DivideOr0(c byte) (string, string) {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s, ""
	}
	return s[:p], s[p + 1:]
}

func (rs RichString) DivideOr1(c byte) (string, string) {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return "", s
	}
	return s[:p], s[p + 1:]
}

func (rs RichString) RemoveBefore(c byte) string {
	s := string(rs)
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s
	}
	return s[p + 1:]
}

//-------------------------------------------------------------------------------------------------

type Pair struct {
	Key, Val string
}

type Pairs []Pair

func (pairs Pairs) Keys() []string {
	var list []string
	for _, p := range pairs {
		list = append(list, p.Key)
	}
	return list
}

func (pairs Pairs) Values() []string {
	var list []string
	for _, p := range pairs {
		switch p.Val {
		case "true", "false": // drop
		default:
			list = append(list, p.Val)
		}
	}
	return list
}

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

func SplitKeyValArgs(args []string) (Pairs, []string) {
	var pairs []Pair
	var leftover []string
	for _, a := range args {
		k, v := "", ""
		eq := strings.LastIndexByte(a, '=')
		if eq >= 0 {
			k, v = a[:eq], a[eq + 1:]
		}
		if k != "" && v != "" {
			p := Pair{a[:eq], a[eq + 1:]}
			pairs = append(pairs, p)
		} else {
			leftover = append(leftover, a)
		}
	}
	return Pairs(pairs), leftover
}

