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
			if k != "" {
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
