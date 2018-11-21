// runtemplate is a command-line tool to facilitate using Go templates for a whole range for text file generation.
// It is particularly good at generating Go source code.

//go:generate statics -i=builtin -o=app/statics.go -pkg=app -group=Builtins

package main

import (
	"flag"
	"fmt"
	"github.com/rickb777/runtemplate/app"
	. "github.com/rickb777/runtemplate/app/support"
	"strings"
)

// Declare a user-defined flag type.
type Strings []string

func (f *Strings) String() string {
	return fmt.Sprint([]string(*f))
}

func (f *Strings) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func failIfLeftoversExist(leftover []string) {
	if len(leftover) > 1 {
		Fail(fmt.Sprintf("Unexpected parameters %v", leftover))
	} else if len(leftover) == 1 {
		Fail(fmt.Sprintf("Unexpected parameter %s", leftover[0]))
	}
}

func main() {
	var tpl, output1, output2 string
	flag.StringVar(&tpl, "tpl", "", "Name of template file; this must be available locally or be on TEMPLATEPATH.")
	flag.StringVar(&output1, "output", "", "Name of the output file.")
	flag.StringVar(&output2, "o", "", "Alias for -output.")

	var depsList Strings
	flag.Var(&depsList, "deps", "List of other dependent files (separated by commas). May appear several times.")

	var force bool
	flag.BoolVar(&Verbose, "v", false, "Verbose progress messages.")
	flag.BoolVar(&force, "f", false, "Force output generation, even if up to date.")
	flag.BoolVar(&Dbg, "z", false, "Debug messages.")

	flag.Parse()

	tpl, args := FindTemplateArg(tpl, flag.Args())

	var deps []string
	for _, s := range depsList {
		deps = append(deps, strings.Split(s, ",")...)
	}

	if len(output1) == 0 && len(output2) > 0 {
		output1 = output2
	}

	types, others, leftover := SplitKeyValArgs(args)
	failIfLeftoversExist(leftover)
	app.Generate(tpl, output1, force, deps, types, others, appVersion)
}
