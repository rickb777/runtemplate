// Runtemplate
// This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
// use-case is for source code generation.
//
// Please see the [README](https://github.com/rickb777/runtemplate/blob/master/README.md).

package main

import (
	. "github.com/rickb777/runtemplate/support"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"io"
)

const defaultTplPath = "/src/github.com/rickb777/runtemplate/builtin"

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
		file := SingleFileMeta(path, templateFile)
		if file.Exists() {
			return file
		}
	}

	return SingleFileMeta(templateFile, templateFile)
}

// Set up some text munging functions that will be available in the templates.
func makeFuncMap() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"yesno": func(yes string, no string, value ...bool) string {
			if len(value) > 0 && value[0] {
				return yes
			}
			return no
		},
		"firstLower": func(s string) RichString {
			return RichString(s).FirstLower()
		},
		"firstUpper": func(s string) RichString {
			return RichString(s).FirstUpper()
		},
		// splitDotFirst returns the first part of a string split on a "."
		// Useful for the case in which you want the package name from a passed value
		// like "package.Type"
		"splitDotFirst": func(s string) RichString {
			first, _ := RichString(s).DivideLastOr0('.')
			return first
		},
		// splitDotLast returns the last part of a string split on a "."
		// Useful for the case in which you want the type name from a passed value
		// like "package.Type"
		"splitDotLast": func(s string) RichString {
			_, second := RichString(s).DivideLastOr0('.')
			return second
		},
	}
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
	var w io.Writer = os.Stdout
	if len(outputFile) > 0 {
		f, err := os.Create(outputFile)
		if err != nil {
			Fail(err)
		}
		defer f.Close()
		w = f
	}

	Debug("Execute template\n")
	err = tmpl.Execute(w, context)
	if err != nil {
		Fail(err)
	}
}

func generate(templateFile, outputFile string, force bool, deps []string, types, others Pairs) {
	Debug("generate '%s' '%s' %v %+v %+v\n", templateFile, outputFile, force, deps, types)

	foundTemplate := findTemplateFileFromPath(templateFile)
	than := templateFile

	youngestDep := foundTemplate

	if outputFile == "" && len(types.TValues()) > 0 {
		keys := strings.Join(types.TValues(), "_")
		tf, _ := RichString(templateFile).DivideLastOr0('.')
		tf = RichString(tf).RemoveBeforeLast('/').ToLower()
		outputFile = (RichString(keys).ToLower() + "_" + tf + ".go").String()
		Debug("default output now '%s'\n", outputFile)
	} else if outputFile == "-" {
		outputFile = "" // implies stdout
	}

	otherDeps := NewFileMeta(false, deps...)
	youngestDep = youngestDep.Younger(YoungestFile(otherDeps...))

	outputInfo := SingleFileMeta(outputFile, "")

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

	context := CreateContext(foundTemplate, outputFile, types, others)
	Debug("context %+v\n", context)

	runTheTemplate(foundTemplate.Path, outputFile, context)
	Info("Generated %s.\n", outputFile)
}
