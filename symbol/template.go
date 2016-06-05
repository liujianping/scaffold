package symbol

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strconv"
	"text/template"
)

var (
	delimeter_begin = "[["
	delimeter_end   = "]]"
	extends         = map[string]interface{}{
		"set": func(renderArgs map[string]interface{}, key string, value interface{}) string {
			renderArgs[key] = value
			return ""
		},
	}
)

func Template(fpath string) (*template.Template, error) {
	_, name := path.Split(fpath)
	return template.New(name).Funcs(extends).Delims(delimeter_begin, delimeter_end).ParseFiles(fpath)
}

func RenderString(src string, data interface{}) (string, error) {
	t, err := template.New("str").Funcs(extends).Delims(delimeter_begin, delimeter_end).Parse(src)
	if err != nil {
		return "", err
	}

	b := bytes.NewBufferString("")
	if err := t.Execute(b, data); err != nil {
		return "", err
	}

	return b.String(), nil
}

func RenderTemplate(src, dest string, data interface{}, force bool) error {
	if IsFileExist(dest) && force == false {
		return nil
	}

	t, err := Template(src)
	if err != nil {
		return err
	}

	dir, _ := path.Split(dest)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	tmp := dest + ".tmp"
	fd, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer fd.Close()

	if err = t.Execute(fd, data); err != nil {
		return err
	}

	if err := os.Rename(tmp, dest); err != nil {
		return err
	}

	log.Println("render file => ", dest)
	//goimports
	gofmt := exec.Command("gofmt", "-w", dest)
	gofmt.Run()

	goimports := exec.Command("goimports", "-w", dest)
	goimports.Run()
	return nil
}

func Extend(name string, fn interface{}) {
	if reflect.TypeOf(fn).Kind() == reflect.Func {
		extends[name] = fn
	}
}

func init() {
	Extend("camel", Camel)
	Extend("plural", Plural)
	Extend("singular", Singular)
	Extend("separate", Separate)
	Extend("lint", Lint)
	Extend("quote", strconv.Quote)
	Extend("convert", Convert)
	Extend("module", ModuleName)
	Extend("class", ClassName)
	Extend("add", Add)
	Extend("sub", Sub)
	Extend("divide", Divide)
	Extend("multiply", Multiply)
}
