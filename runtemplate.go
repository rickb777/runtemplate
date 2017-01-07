// Runtemplate
// This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
// use-case is for source code generation.
//
// Please see the [README](https://github.com/rickb777/runtemplate/blob/master/README.md).

package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"runtime"
	. "github.com/rickb777/runtemplate/support"
)

const defaultTplPath = "/src/github.com/rickb777/runtemplate/builtin"
const Prefix = "Prefix"

func findTemplateFileFromPath(templateFile string) FileMeta {
	Debug("findTemplateFileFromPath '%s'\n", templateFile)
	templatePath := os.Getenv("TEMPLATEPATH")
	Debug("TEMPLATEPATH=%s\n", templatePath)

	goPath := os.Getenv("GOPATH")
	if goPath != "" {
		if templatePath != "" {
			templatePath = templatePath + ":"
		}
		templatePath = templatePath + goPath + defaultTplPath
	}

	x := strings.Split(templatePath, ":")
	Debug("searching template path %+v\n", x)
	for _, p := range x {
		path := p + "/" + templateFile
		file := NewFileMeta(true, path)[0]
		if file.Exists() {
			return file
		}
	}

	return NewFileMeta(true, templateFile)[0]
}

// Set up some text munging functions that will be available in the templates.
func makeFuncMap() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"condFirstUpper": func(s string, condition interface{}) string {
			if v, ok := condition.(string); ok && v == "true" {
				return RichString(s).FirstUpper()
			}
			return s
		},
		"firstLower": func(s string) string {
			return RichString(s).FirstLower()
		},
		"firstUpper": func(s string) string {
			return RichString(s).FirstUpper()
		},
		// splitDotFirst returns the first part of a string split on a "."
		// Useful for the case in which you want the package name from a passed value
		// like "package.Type"
		"splitDotFirst": func(s string) string {
			first, _ := RichString(s).DivideOr0('.')
			return first
		},
		// splitDotLast returns the last part of a string split on a "."
		// Useful for the case in which you want the type name from a passed value
		// like "package.Type"
		"splitDotLast": func(s string) string {
			_, second := RichString(s).DivideOr0('.')
			return second
		},
	}
}

func choosePackage(outputFile string) (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		Fail(err)
	}

	pkg := RichString(wd).RemoveBefore('/')

	if strings.IndexByte(outputFile, '/') > 0 {
		dir, _ := RichString(outputFile).DivideOr0('/')
		if strings.IndexByte(dir, '/') > 0 {
			dir = RichString(dir).RemoveBefore('/')
		}
		if dir != "." {
			pkg = dir
		}
	}

	return wd, pkg
}

func setTypeInContext(k, v string, context map[string]interface{}) {
	p := v
	star := ""
	amp := ""

	if len(v) > 0 && v[0] == '*' {
		v = v[1:]
		star = "*"
		amp = "&"
	}

	Debug("setTypeInContext %s=%s for %s, star=%s, amp=%s\n", k, v, p, star, amp)

	context[k] = v
	context["U" + k] = RichString(v).FirstUpper()
	context["L" + k] = RichString(v).FirstLower()

	if !strings.HasSuffix(k, Prefix) {
		context[k + "Star"] = star
		context[k + "Amp"] = amp
		context["P" + k] = p
	}
}

func setPairInContext(pp Pair, context map[string]interface{}) {
	k := pp.Key
	v := pp.Val
	switch v {
	case "true":
		context[k] = true
	case "false":
		context[k] = false
	default:
		setTypeInContext(k, v, context)
	}
}

func createContext(foundTemplate FileMeta, outputFile string, vals Pairs) map[string]interface{} {
	// Context will be passed to the template as a map.
	context := make(map[string]interface{})
	context["GOARCH"] = runtime.GOARCH
	context["GOOS"] = runtime.GOOS
	context["GOPATH"] = os.Getenv("GOPATH")
	context["GOROOT"] = os.Getenv("GOROOT")

	context["PWD"], context["Package"] = choosePackage(outputFile)

	// set up some special context values just in case they are wanted.
	context["OutFile"] = outputFile
	context["TemplateFile"] = foundTemplate.Name
	context["TemplatePath"] = foundTemplate.Path

	// define automatic prefix template values with default blank value.
	for _, p := range vals {
		if strings.HasSuffix(p.Key, "Type") {
			l := len(p.Key)
			k := p.Key[:l-4]
			setTypeInContext(k + Prefix, "", context)
		}
	}

	// copy the key/vals to template values
	for _, p := range vals {
		setPairInContext(p, context)
	}

	return context
}

func runTheTemplate(foundTemplate, outputFile string, context map[string]interface{}) {
	Debug("ReadFile %s\n", foundTemplate)
	b, err := ioutil.ReadFile(foundTemplate)
	if err != nil {
		Fail(err)
	}

	funcMap := makeFuncMap()
	Debug("Parse template\n")
	tmpl, err := template.New(foundTemplate).Funcs(funcMap).Parse(string(b))
	if err != nil {
		Fail(err)
	}

	Debug("Create %s\n", outputFile)
	f, err := os.Create(outputFile)
	if err != nil {
		Fail(err)
	}
	defer f.Close()

	Debug("Execute template\n")
	err = tmpl.Execute(f, context)
	if err != nil {
		Fail(err)
	}
}

func generate(templateFile, outputFile string, force bool, deps []string, vals Pairs) {
	Debug("generate %s %s %v %+v %+v\n", templateFile, outputFile, force, deps, vals)

	foundTemplate := findTemplateFileFromPath(templateFile)
	than := templateFile

	youngestDep := foundTemplate

	if outputFile == "" {
		keys := strings.Join(vals.TValues(), "_")
		tf, _ := RichString(templateFile).DivideOr0('.')
		tf = RichString(tf).RemoveBefore('/')
		outputFile = strings.ToLower(keys + "_" + tf) + ".go"
		Debug("default output now '%s'\n", outputFile)
	}

	otherDeps := NewFileMeta(false, deps...)
	youngestDep = youngestDep.Younger(YoungestFile(otherDeps...))

	outputInfo := NewFileMeta(true, outputFile)[0]

	Debug("output=%s %v, youngest=%s %v\n", outputInfo.Name, outputInfo.ModTime, youngestDep.Name, youngestDep.ModTime)
	if outputInfo.ModTime.After(youngestDep.ModTime) {
		if !force {
			if len(deps) > 0 {
				than = than + ", " + strings.Join(deps, ", ")
			}
			Info("%s is already newer than %s.\n", outputFile, than)
			return
		}
	}

	context := createContext(foundTemplate, outputFile, vals)
	Debug("context %+v\n", context)

	runTheTemplate(foundTemplate.Path, outputFile, context)
	Info("Generated %s.\n", outputFile)
}
