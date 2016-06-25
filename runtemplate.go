// Runtemplate
// This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
// use-case is for source code generation.
//
// Please see the [README](https://github.com/rickb777/runtemplate/blob/master/README.md).

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const defaultTplPath = "/src/github.com/rickb777/runtemplate/builtin"

var tpl = flag.String("tpl", "", "Name of template file; this must be available locally or be on TEMPLATEPATH.")
var output = flag.String("output", "", "Name of the output file.")
var mainType = flag.String("type", "", "Name of the main type.")
var deps = flag.String("deps", "", "List of other dependent files (separated by commas).")
var force = flag.Bool("f", false, "Force output generation, even if up to date.")
var verbose = flag.Bool("v", false, "Verbose progress messages.")
var dbg = flag.Bool("z", false, "Debug messages.")

func fail(args ...interface{}) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func info(msg string, args ...interface{}) {
	if *verbose {
		fmt.Printf(msg, args...)
	}
}

func debug(msg string, args ...interface{}) {
	if *dbg {
		fmt.Printf(msg, args...)
	}
}

func divide(s string, c byte) (string, string) {
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s, ""
	}
	return s[:p], s[p+1:]
}

func removeBefore(s string, c byte) string {
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s
	}
	return s[p+1:]
}

func chooseArg(flagValue *string, suffix string) string {
	args := flag.Args()

	var arg string
	if flagValue != nil {
		arg = *flagValue
	} else {
		for i, a := range args {
			if strings.HasSuffix(a, suffix) {
				arg = a
				args[i] = ""
			}
		}
	}
	return arg
}

func findTemplateFileFromPath(templateFile string) (string, os.FileInfo) {
	debug("findTemplateFileFromPath '%s'\n", templateFile)
	templatePath := os.Getenv("TEMPLATEPATH")
	debug("TEMPLATEPATH=%s\n", templatePath)

	goPath := os.Getenv("GOPATH")
	if goPath != "" {
		if templatePath != "" {
			templatePath = templatePath + ":"
		}
		templatePath = templatePath + goPath + defaultTplPath
	}

	x := strings.Split(templatePath, ":")
	debug("searching path %+v\n", x)
	for _, p := range x {
		path := p + "/" + templateFile
		debug("stat '%s'\n", path)
		info, err := os.Stat(path)
		if err == nil {
			if info.IsDir() {
				fail(fmt.Errorf("%s is a directory.", path))
			}
			return path, info
		}
		if !os.IsNotExist(err) {
			fail(path, err)
		}
	}

	debug("stat '%s'\n", templateFile)
	info, err := os.Stat(templateFile)
	if os.IsNotExist(err) {
		fail(templateFile, err)
	}
	return templateFile, info
}

// Set up some text munging functions that will be available in the templates.
func makeFuncMap() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		// splitDotFirst returns the first part of a string split on a "."
		// Useful for the case in which you want the package name from a passed value
		// like "package.Type"
		"splitDotFirst": func(s string) string {
			first, _ := divide(s, '.')
			return first
		},
		// splitDotLast returns the last part of a string split on a "."
		// Useful for the case in which you want the type name from a passed value
		// like "package.Type"
		"splitDotLast": func(s string) string {
			_, second := divide(s, '.')
			return second
		},
	}
}

func choosePackage(outputFile string) (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	pkg := removeBefore(wd, '/')

	if strings.IndexByte(outputFile, '/') > 0 {
		dir, _ := divide(outputFile, '/')
		if strings.IndexByte(dir, '/') > 0 {
			dir = removeBefore(dir, '/')
		}
		if dir != "." {
			pkg = dir
		}
	}

	return wd, pkg
}

func setTypeInContext(kind, t string, context map[string]interface{}) string {
	p := t
	if t[0] == '*' {
		t = t[1:]
	}
	lt := strings.ToLower(t)
	context[kind] = t
	context["P"+kind] = p
	context["U"+kind] = strings.ToUpper(t[:1]) + t[1:]
	context["L"+kind] = strings.ToLower(t[:1]) + t[1:]
	return lt
}

func runTheTemplate(templateFile, outputFile string, context map[string]interface{}) {
	debug("ReadFile %s\n", templateFile)
	b, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fail(err)
	}

	//context["GOARCH"] = os.Getenv("GOARCH")
	//context["GOOS"] = os.Getenv("GOOS")
	context["GOPATH"] = os.Getenv("GOPATH")
	context["GOROOT"] = os.Getenv("GOROOT")

	context["PWD"], context["Package"] = choosePackage(outputFile)

	for _, arg := range flag.Args() {
		if strings.Contains(arg, "=") {
			key1, val1 := divide(arg, '=')
			key2, val2 := strings.TrimSpace(key1), strings.TrimSpace(val1)
			switch val2 {
			case "true":
				context[key2] = true
			case "false":
				context[key2] = false
			default:
				setTypeInContext(key2, val2, context)

				context[key2] = val2
			}
		}
	}
	debug("context %+v\n", context)

	funcMap := makeFuncMap()
	debug("Parse template\n")
	tmpl, err := template.New(templateFile).Funcs(funcMap).Parse(string(b))
	if err != nil {
		fail(err)
	}

	debug("Create %s\n", outputFile)
	f, err := os.Create(outputFile)
	if err != nil {
		fail(err)
	}
	defer f.Close()

	debug("Execute template\n")
	err = tmpl.Execute(f, context)
	if err != nil {
		fail(err)
	}
}

func youngestDependency(main ...os.FileInfo) os.FileInfo {
	result := main[0]
	for _, m := range main {
		if m != nil && m.ModTime().After(result.ModTime()) {
			debug("change dep1 %s %v -> %s %v\n", result.Name(), result.ModTime(), m.Name(), m.ModTime())
			result = m
		}
	}

	if deps == nil || *deps == "" {
		return result
	}

	list := strings.Split(*deps, ",")
	for _, f := range list {
		fi, err := os.Stat(f)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Warn: %s does not exist.\n", f)
			} else {
				fail(err)
			}
		} else {
			if fi.ModTime().After(result.ModTime()) {
				debug("change dep2 %s %v -> %s %v\n", result.Name(), result.ModTime(), fi.Name(), fi.ModTime())
				result = fi
			}
		}
	}

	return result
}

func generate() {
	// Context will be passed to the template as a map.
	var err error
	context := make(map[string]interface{})

	templateFile := chooseArg(tpl, ".tpl")
	outputFile := chooseArg(output, ".go")

	var mainTypeInfo os.FileInfo
	var mainTypeGo string
	if mainType != nil && *mainType != "" {
		lt := setTypeInContext("Type", *mainType, context)
		mainTypeGo = lt + ".go"

		debug("stat %s\n", mainTypeGo)
		mainTypeInfo, err = os.Stat(mainTypeGo)
		if os.IsNotExist(err) {
			mainTypeInfo = nil
			mainTypeGo = ""
		}

		if mainTypeGo == outputFile {
			fail(mainTypeGo, "is specified as both an input dependency and the output file.")
		} else if outputFile == "" {
			tf, _ := divide(templateFile, '.')
			tf = removeBefore(tf, '/')
			outputFile = fmt.Sprintf("%s_%s.go", lt, tf)
			debug("default output now '%s'\n", outputFile)
		}
	}

	if outputFile == "" {
		fail("Output file must be specified.")
	}

	foundTemplate, templateInfo := findTemplateFileFromPath(templateFile)

	// set up some special context values just in case they are wanted.
	context["OutFile"] = outputFile
	context["TemplateFile"] = templateFile
	context["TemplatePath"] = foundTemplate

	youngestDep := youngestDependency(templateInfo, mainTypeInfo)

	var outputInfo os.FileInfo
	if !*force {
		debug("stat %s\n", outputFile)
		outputInfo, err = os.Stat(outputFile)
		if os.IsNotExist(err) {
			outputInfo = nil
		}
	}

	if outputInfo != nil {
		debug("output=%s %v, youngest=%s %v\n", outputInfo.Name(), outputInfo.ModTime(), youngestDep.Name(), youngestDep.ModTime())
		if outputInfo.ModTime().After(youngestDep.ModTime()) {
			than := templateFile
			if mainTypeInfo != nil {
				than = than + "," + mainTypeGo
			}
			if deps != nil && *deps != "" {
				than = than + "," + *deps
			}
			info("%s is already newer than %s.\n", outputFile, than)
			return
		}
	}

	runTheTemplate(foundTemplate, outputFile, context)
	info("Generated %s.\n", outputFile)
}

func main() {
	flag.Parse()
	generate()
}
