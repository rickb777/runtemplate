package support

import (
	"fmt"
	"os"
	"strings"
)

var ShowContextInfo = false
var Verbose = false
var Dbg = false

func Fail(args ...interface{}) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func Progress(msg string, args ...interface{}) {
	if Verbose {
		fmt.Printf(msg, args...)
	}
}

func Debug(msg string, args ...interface{}) {
	if Dbg {
		fmt.Printf("-- "+msg, args...)
	}
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

func expandSpecialChars(s string) string {
	s2 := strings.Replace(s, `\n`, "\n", -1)
	return strings.Replace(s2, `\t`, "\t", -1)
}

func SplitKeyValArgs(args []string) (Triples, Triples, []string) {
	var types []Triple
	var others []Triple
	var leftover []string
	for _, a := range args {
		found := false
		k, v := "", ""
		eq := strings.LastIndexByte(a, '=')
		co := strings.LastIndexByte(a, ':')
		sl := strings.LastIndexByte(a, '/')
		if eq >= 0 {
			k, v = a[:eq], a[eq+1:]
			if k != "" && v != "" {
				if sl > eq {
					p := Triple{Key: a[:eq], Val: a[eq+1 : sl], Alt: a[sl+1:]}
					types = append(types, p)
				} else {
					p := Triple{Key: a[:eq], Val: a[eq+1:], Alt: ""}
					types = append(types, p)
				}
				found = true
			}
		} else if co >= 0 {
			k, v = a[:co], a[co+1:]
			if k != "" {
				p := Triple{Key: a[:co], Val: expandSpecialChars(a[co+1:]), Alt: ""}
				others = append(others, p)
				found = true
			}
		}
		if !found {
			leftover = append(leftover, a)
		}
	}
	return Triples(types), Triples(others), leftover
}
