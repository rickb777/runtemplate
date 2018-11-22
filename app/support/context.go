package support

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

const Prefix = "Prefix"

func abort(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

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

func setAnyInContext(k, v string, context map[string]interface{}) {
	if !strings.HasSuffix(k, Prefix) {
		Debug("setAnyInContext %s=%s\n", k, v)
		context[k] = v
		context["U"+k] = "Any"
		context["L"+k] = "any"
		context[k+"IsPtr"] = false
		context[k+"Star"] = ""
		context[k+"Amp"] = ""
		context[k+"Zero"] = "nil"
	}
}

func setTypeInContext(k, v string, context map[string]interface{}) {
	p := v
	ptr := false

	if len(v) > 0 && v[0] == '*' {
		v = v[1:]
		ptr = true
	}

	Debug("setTypeInContext %s=%s for %s, ptr=%v\n", k, v, p, ptr)

	rs := RichString(v).NoDots()
	context[k] = v
	context["U"+k] = rs.FirstUpper().String()
	context["L"+k] = rs.FirstLower().String()

	if !strings.HasSuffix(k, Prefix) {
		context["P"+k] = p
		context[k+"IsPtr"] = ptr
		if ptr {
			context[k+"Star"] = "*"
			context[k+"Amp"] = "&"
			context[k+"Zero"] = "nil"
		} else {
			context[k+"Star"] = ""
			context[k+"Amp"] = ""
			switch v {
			case "string":
				context[k+"Zero"] = `""`
			case "bool":
				context[k+"Zero"] = `false`
			case "int", "int8", "int16", "int32", "int64",
				"uint", "uint8", "uint16", "uint32", "uint64",
				"float32", "float64", "byte", "rune":
				context[k+"Zero"] = `0`
			default:
				context[k+"Zero"] = fmt.Sprintf("*(new(%s))", v)
			}
		}
	}
}

func setPairTypeInContext(pp Triple, context map[string]interface{}) {
	k := pp.Key
	v := pp.Val
	switch v {
	case "true":
		context[k] = true
	case "false":
		context[k] = false
	case "interface{}":
		setAnyInContext(k, v, context)
	default:
		setTypeInContext(k, v, context)
	}
	context["Has"+k] = true
}

func addPairInContext(pp Triple, context map[string]interface{}) {
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

func contextInfo(others Triples, context map[string]interface{}) {
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

func CreateContext(templateFile FileMeta, outputFile string, types, others Triples, appVersion string) map[string]interface{} {
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
		if strings.HasSuffix(p.Key, "Type") {
			l := len(p.Key)
			k := p.Key[:l-4]
			setTypeInContext(k+Prefix, "", context)
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
