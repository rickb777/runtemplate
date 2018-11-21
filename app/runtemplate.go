// Runtemplate
// This application provides a simple way of exxecuting standard Go templates from the command line. The obvious
// use-case is for source code generation.
//
// Please see the [README](https://github.com/rickb777/runtemplate/blob/master/README.md).

package app

import (
	"fmt"
	"github.com/go-playground/statics/static"
	. "github.com/rickb777/runtemplate/app/support"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var builtins *static.Files

func mustLoadBuiltins() {
	var err error
	builtins, err = newStaticBuiltins(&static.Config{UseStaticFiles: true})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func findTemplateFileFromPath(templateFile string) FileMeta {
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

	st := "/builtin/" + templateFile
	f, err := builtins.GetHTTPFile(st)
	if err != nil {
		Fail(err)
	}

	fi, err := f.Stat()
	if err != nil {
		Fail(err)
	}

	mt := fi.ModTime()
	f.Close()
	return EmbeddedFileMeta(st, templateFile, mt)
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
	Debug("ReadFile %s\n", foundTemplate)
	var b []byte
	var err error

	if foundTemplate.Embedded {
		b, err = builtins.ReadFile("/builtin/" + foundTemplate.Name)
	} else {
		b, err = ioutil.ReadFile(foundTemplate.Path)
	}

	if err != nil {
		Fail(err)
	}

	funcMap := makeFuncMap()
	Debug("Parse template\n")
	tmpl, err := template.New(foundTemplate.Path).Funcs(funcMap).Parse(string(b))
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

func Generate(templateFile, outputFile string, force bool, deps []string, types, others Pairs, appVersion string) {
	Debug("generate %s %s %v %+v %+v\n", templateFile, outputFile, force, deps, types)

	mustLoadBuiltins()
	foundTemplate := findTemplateFileFromPath(templateFile)
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
			Info("%s is already newer than %s.\n", outputFile, than)
			return
		}
	}

	context := CreateContext(foundTemplate, outputFile, types, others, appVersion)
	Debug("context %+v\n", context)

	runTheTemplate(foundTemplate, outputFile, context)
	Info("Generated %s.\n", outputFile)
}
