package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"flag"
)

var tpl = flag.String("tpl", "", "Name of template file; this must be available locally or be on TEMPLATEPATH.")
var output = flag.String("output", "", "Name of the output file.")
var mainType = flag.String("type", "", "Name of the main type.")
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
	return s[:p], s[p + 1:]
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

	if len(templatePath) > 0 {
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

func runTheTemplate(templateFile, outputFile string, context map[string]string) {
	debug("ReadFile %s\n", templateFile)
	b, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fail(err)
	}

	for _, arg := range flag.Args() {
		if strings.Contains(arg, "=") {
			key, val := divide(arg, '=')
			context[strings.TrimSpace(key)] = strings.TrimSpace(val)
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

func generate() {
	// Context will be passed to the template as a map.
	context := make(map[string]string)

	templateFile := chooseArg(tpl, ".tpl")
	outputFile := chooseArg(output, ".go")
	if outputFile == "" {
		fail("Output file must be specified.")
	}

	var outputInfo os.FileInfo
	var err error
	if !*force {
		debug("stat %s\n", outputFile)
		outputInfo, err = os.Stat(outputFile)
		if os.IsNotExist(err) {
			outputInfo = nil
		}
	}

	var mainTypeInfo os.FileInfo
	var mainTypeGo string
	if mainType != nil && *mainType != "" {
		t := *mainType
		mainTypeGo = strings.ToLower(t) + ".go"
		context["Type"] = t
		context["LType"] = strings.ToLower(t[:1]) + t[1:]
		debug("stat %s\n", mainTypeGo)
		mainTypeInfo, err = os.Stat(mainTypeGo)
		if os.IsNotExist(err) {
			mainTypeInfo = nil
			mainTypeGo = ""
		}
		if mainTypeGo == outputFile {
			fail(mainTypeGo, "is specified as both an input dependency and the output file.")
		}
	}

	foundTemplate, templateInfo := findTemplateFileFromPath(templateFile)

	// set up some special context values just in case they are wanted.
	context["OutFile"] = outputFile
	context["TemplateFile"] = templateFile
	context["TemplatePath"] = foundTemplate
	//context["GOARCH"] = os.Getenv("GOARCH")
	//context["GOOS"] = os.Getenv("GOOS")
	context["GOPATH"] = os.Getenv("GOPATH")
	context["GOROOT"] = os.Getenv("GOROOT")

	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}
	context["PWD"] = wd
	_, context["Package"] = divide(wd, '/')

	youngestDep := templateInfo
	if mainTypeInfo != nil {
		debug("mainType=%v, template=%v\n", mainTypeInfo.ModTime(), templateInfo.ModTime())
		if mainTypeInfo.ModTime().After(templateInfo.ModTime()) {
			youngestDep = mainTypeInfo
		}
	}

	if outputInfo != nil {
		debug("output=%v, youngest=%v\n", outputInfo.ModTime(), youngestDep.ModTime())
		if outputInfo.ModTime().After(youngestDep.ModTime()) {
			info("%s is already newer than %s.\n", outputFile, templateFile)
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
