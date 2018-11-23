package support

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

const Prefix = "Prefix"

func choosePackage(outputFile string) (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		Fail(err)
	}

	pkg := RichString(wd).RemoveBeforeLast('/')

	if strings.IndexByte(outputFile, '/') > 0 {
		dir, _ := RichString(outputFile).DivideLastOr0('/')
		if dir.IndexByte('/') > 0 {
			dir = RichString(dir).RemoveBeforeLast('/')
		}
		if dir != "." {
			pkg = dir
		}
	}

	return wd, pkg.String()
}

func setIdentInContext(pp Tuple, context map[string]interface{}) {
	Debug("setIdentInContext %+v\n", pp)

	k := pp.Key
	rs := pp.Ident().NoDots()
	context[k] = pp
	context["U"+k] = rs.FirstUpper().String()
	context["L"+k] = rs.FirstLower().String()
}

func setTypeInContext(pp Tuple, context map[string]interface{}) {
	Debug("setTypeInContext %+v\n", pp)

	k := pp.Key

	if !strings.HasSuffix(k, Prefix) {
		context["P"+k] = pp.s
		context[k+"IsPtr"] = pp.IsPtr()
		context[k+"Star"] = pp.Star()
		context[k+"Amp"] = pp.Amp()
		context[k+"Zero"] = pp.Zero()
	}
}

func setPairTypeInContext(pp Tuple, context map[string]interface{}) {
	k := pp.Key
	v := pp.s
	switch v {
	case "true":
		context[k] = true
	case "false":
		context[k] = false
	default:
		setIdentInContext(pp, context)
		setTypeInContext(pp, context)
	}
	context["Has"+k] = true
}

func addPairInContext(pp Pair, context map[string]interface{}) {
	k := pp.Key
	v := pp.Val
	switch v {
	case "true":
		context[k] = true
	case "false":
		context[k] = false
	default:
		e, ok := context[k]
		if ok {
			s, ok := e.(string)
			if ok {
				context[k] = []string{s, v}
			} else {
				ss := e.([]string)
				ss = append(ss, v)
				context[k] = ss
			}
		} else {
			context[k] = v
			context["Has"+k] = true
		}
	}
}

func copyOf(context map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range context {
		cp[k] = v
	}
	return cp
}

func contextInfo1(key string, context map[string]interface{}) {
	fmt.Printf("%-14s= %v\n", key, context[key])
	delete(context, key)
}

func contextInfo(others Pairs, context map[string]interface{}) {
	contextInfo1("AppVersion", context)
	contextInfo1("PWD", context)
	contextInfo1("Package", context)
	contextInfo1("Outfile", context)
	contextInfo1("TemplateFile", context)
	contextInfo1("TemplatePath", context)

	fmt.Printf("GOARCH = %s, GOOS = %s, GOPATH = %s, GOROOT = %s\n", context["GOARCH"], context["GOOS"], context["GOPATH"], context["GOROOT"])

	delete(context, "GOARCH")
	delete(context, "GOOS")
	delete(context, "GOPATH")
	delete(context, "GOROOT")

	for _, p := range others {
		if _, exists := context[p.Key]; exists {
			contextInfo1(p.Key, context)
		}
	}

	keys := make([]string, 0, len(context))
	for k, _ := range context {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-14s= %v\n", k, context[k])
	}
}

func CreateContext(templateFile FileMeta, outputFile string, types Tuples, others Pairs, appVersion string) map[string]interface{} {
	// Context will be passed to the template as a map.
	context := make(map[string]interface{})
	context["GOARCH"] = runtime.GOARCH
	context["GOOS"] = runtime.GOOS
	context["GOPATH"] = os.Getenv("GOPATH")
	context["GOROOT"] = os.Getenv("GOROOT")

	context["PWD"], context["Package"] = choosePackage(outputFile)

	// set up some special context values just in case they are wanted.
	context["OutFile"] = outputFile
	context["TemplateFile"] = templateFile.Name
	context["TemplatePath"] = templateFile.Path
	context["AppVersion"] = appVersion

	// define automatic prefix template values with default blank value.
	for _, p := range types {
		if strings.HasSuffix(p.Key, "Tuple") {
			l := len(p.Key)
			k := p.Key[:l-4]
			setIdentInContext(NewTuple(k+Prefix+"="), context)
		}
	}

	// copy the key/vals to template values
	for _, p := range types {
		setPairTypeInContext(p, context)
	}

	for _, p := range others {
		addPairInContext(p, context)
	}

	if ShowContextInfo {
		contextInfo(others, copyOf(context))
	}
	return context
}
