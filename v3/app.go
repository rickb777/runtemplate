// runtemplate is a command-line tool to facilitate using Go templates for a whole range for text file generation.
// It is particularly good at generating Go source code.

//go:generate packr -z

package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/rickb777/runtemplate/v3/app"
	"github.com/rickb777/runtemplate/v3/app/support"
	"os"
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
		support.Fail(fmt.Sprintf("Unexpected parameters %v", leftover))
	} else if len(leftover) == 1 {
		support.Fail(fmt.Sprintf("Unexpected parameter %s", leftover[0]))
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "A template file must be specified.")
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s [options] [template] [Type=Name ...] [Flag:Value ...]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "\nOptions:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nSee https://github.com/rickb777/runtemplate/blob/master/README.md")
	fmt.Fprintln(os.Stderr, "and https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md")
	fmt.Fprintln(os.Stderr, "Version", appVersion)
	os.Exit(1)
}

func main() {
	var tpl, output1, output2 string
	flag.StringVar(&tpl, "tpl", "", "Name of template file; this must be available locally or be on TEMPLATEPATH.")
	flag.StringVar(&output1, "output", "", "Name of the output file. Default is computed from template name plus types.")
	flag.StringVar(&output2, "o", "", "Alias for -output.")

	var depsList Strings
	flag.Var(&depsList, "deps", "List of other dependent files (separated by commas) to avoid unnecessary output file change. May appear several times.")

	var force, showVersion bool
	flag.BoolVar(&support.Verbose, "v", false, "Verbose progress messages.")
	flag.BoolVar(&support.ShowContextInfo, "i", false, "Show the context.")
	flag.BoolVar(&force, "f", false, "Force output generation, even if up to date.")
	flag.BoolVar(&support.Dbg, "z", false, "Debug messages.")
	flag.BoolVar(&showVersion, "version", false, "Show the version.")

	flag.Parse()

	tpl, args := support.FindTemplateArg(tpl, flag.Args())

	if tpl == "" {
		usage()
	}

	if showVersion {
		fmt.Println(appVersion)
	}

	var deps []string
	for _, s := range depsList {
		deps = append(deps, strings.Split(s, ",")...)
	}

	if len(output1) == 0 && len(output2) > 0 {
		output1 = output2
	}

	types, others, leftover := support.SplitKeyValArgs(args)
	failIfLeftoversExist(leftover)
	app.Generate(tpl, output1, force, deps, types, others, builtins, appVersion)
}

var builtins = []packr.Box{
	packr.NewBox("builtin"),
	//packr.NewBox("builtin/fast"),
	//packr.NewBox("builtin/immutable"),
	//packr.NewBox("builtin/plumbing"),
	//packr.NewBox("builtin/simple"),
	//packr.NewBox("builtin/threadsafe"),
	//packr.NewBox("builtin/types"),
}
