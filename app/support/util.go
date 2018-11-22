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

func SplitKeyValArgs(args []string) (Types, Pairs, []string) {
	var types Types
	var others Pairs
	var leftover []string

	for _, a := range args {
		found := false
		eq := strings.LastIndexByte(a, '=')
		co := strings.LastIndexByte(a, ':')
		if eq >= 0 {
			tr := NewType(a)
			if tr.Valid() {
				types = append(types, tr)
			}
		} else if co >= 0 {
			k, v := a[:co], a[co+1:]
			if k != "" {
				p := Pair{Key: k, Val: expandSpecialChars(v)}
				others = append(others, p)
				found = true
			}
		}
		if !found {
			leftover = append(leftover, a)
		}
	}
	return types, others, leftover
}
