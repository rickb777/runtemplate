// Runtemplate
// This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
// use-case is for source code generation.
//
// Please see the [README](https://github.com/rickb777/runtemplate/blob/master/README.md).

package app

import (
	"github.com/gobuffalo/packr"
	. "github.com/rickb777/runtemplate/v3/app/support"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

func findTemplateFileFromPath(templateFile string, builtins []packr.Box) FileMeta {
	Debug("findTemplateFileFromPath %q\n", templateFile)

	templatePath := os.Getenv("TEMPLATEPATH")
	if templatePath == "" {
		templatePath = "."
	}
	Debug("TEMPLATEPATH=%s\n", templatePath)

	x := strings.Split(templatePath, ":")
	Debug("searching template path %+v\n", x)
	for _, p := range x {
		fp := p + "/" + templateFile
		file := SingleFileMeta(path.Clean(fp), templateFile)
		if file.Exists() {
			return file
		}
	}

	var err error
	var content string
	for _, box := range builtins {
		content, err = box.FindString(templateFile)
		if err == nil {
			st := "builtin/" + templateFile
			return EmbeddedFileMeta(st, templateFile, content)
		}
	}

	Fail(templateFile, err)
	return FileMeta{}
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

func runTheTemplate(foundTemplate FileMeta, outputFile string, context map[string]interface{}) {
	var s string
	var err error

	if foundTemplate.Embedded != "" {
		Debug("ReadFile builtin %+v\n", foundTemplate)
		s = foundTemplate.Embedded
	} else {
		Debug("ReadFile %+v\n", foundTemplate)
		s, err = readFile(foundTemplate.Path)
		if err != nil {
			Fail(foundTemplate.Path, err)
		}
	}

	funcMap := makeFuncMap()
	Debug("Parse template\n")
	tmpl, err := template.New(foundTemplate.Path).Funcs(funcMap).Parse(s)
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
		Fail("execute template:", err)
	}
}

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Generate(templateFile, outputFile string, force bool, deps []string, types Tuples, others Pairs, builtins []packr.Box, appVersion string) {
	Debug("generate %s %q %v %+v %#v\n", templateFile, outputFile, force, deps, types)

	foundTemplate := findTemplateFileFromPath(templateFile, builtins)
	than := templateFile

	youngestDep := foundTemplate

	if outputFile == "" && len(types.TValues()) > 0 {
		keys := strings.Join(types.TValues(), "_")
		tf, _ := RichString(templateFile).DivideLastOr0('.')
		tf = RichString(tf).RemoveBeforeLast('/').ToLower()
		outputFile = (RichString(keys).ToLower() + "_" + tf + ".go").String()
		Debug("default output now %s\n", outputFile)
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
			Progress("%s is already newer than %s.\n", outputFile, than)
			return
		}
	}

	context := CreateContext(foundTemplate, outputFile, types, others, appVersion)
	Debug("context %+v\n", context)

	runTheTemplate(foundTemplate, outputFile, context)
	Progress("Generated %s.\n", outputFile)
}
