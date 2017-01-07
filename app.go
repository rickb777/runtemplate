package main

import (
	"flag"
	"strings"
	. "github.com/rickb777/runtemplate/support"
	"fmt"
)

func failIfLeftoversExist(leftover []string) {
	if len(leftover) > 1 {
		Fail(fmt.Sprintf("Unexpected parameters %v", leftover))
	} else if len(leftover) == 1 {
		Fail(fmt.Sprintf("Unexpected parameter %s", leftover[0]))
	}
}

func main() {
	var tpl, output, depsList string
	flag.StringVar(&tpl, "tpl", "", "Name of template file; this must be available locally or be on TEMPLATEPATH.")
	flag.StringVar(&output, "output", "", "Name of the output file.")
	flag.StringVar(&depsList, "deps", "", "List of other dependent files (separated by commas).")

	var force bool
	flag.BoolVar(&Verbose, "v", false, "Verbose progress messages.")
	flag.BoolVar(&force, "f", false, "Force output generation, even if up to date.")
	flag.BoolVar(&Dbg, "z", false, "Debug messages.")

	flag.Parse()

	tpl, args := FindTemplateArg(tpl, flag.Args())

	deps := strings.Split(depsList, ",")
	keyVals, leftover := SplitKeyValArgs(args)
	failIfLeftoversExist(leftover)
	generate(tpl, output, force, deps, keyVals)
}

