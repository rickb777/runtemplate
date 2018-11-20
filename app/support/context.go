package support

import (
	"fmt"
	"os"
	"runtime"
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

func setTypeInContext(k, v string, context map[string]interface{}) {
	p := v
	ptr := false

	if len(v) > 0 && v[0] == '*' {
		v = v[1:]
		ptr = true
	}

	Debug("setTypeInContext %s=%s for %s, ptr=%v\n", k, v, p, ptr)

	rs := RichString(v).RemoveBeforeLast('.')
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

func setPairTypeInContext(pp Pair, context map[string]interface{}) {
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

func CreateContext(templateFile FileMeta, outputFile string, types, others Pairs, appVersion string) map[string]interface{} {
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

	return context
}
