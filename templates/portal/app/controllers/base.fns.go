package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"time"

	"github.com/liujianping/scaffold/symbol"
	"github.com/revel/revel"
)

type Page struct {
	No       int64
	Size     int64
	Head     bool
	Tail     bool
	Normal   bool
	Selected bool
}

var (
	TimeLocation       *time.Location
	TimeLocationBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			for _, f := range revel.TimeFormats {
				if r, err := time.ParseInLocation(f, val, TimeLocation); err == nil {
					return reflect.ValueOf(r)
				}
			}
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			var (
				t       = val.(time.Time)
				format  = revel.DateTimeFormat
				h, m, s = t.Clock()
			)
			if h == 0 && m == 0 && s == 0 {
				format = revel.DateFormat
			}
			output[name] = t.Format(format)
		},
	}

	Widget = func(widget string, mode string, name string, value interface{}, options ...interface{}) template.HTML {
		tname := "text"
		if widget != "" {
			tname = strings.ToLower(widget)
		}

		tmpl, err := revel.MainTemplateLoader.Template(fmt.Sprintf("widget/%s.html", tname))
		if err != nil {
			return template.HTML(fmt.Sprintf("widget/%s.html not exist.", widget))
		}

		wr := bytes.NewBuffer([]byte(""))
		if err := tmpl.Render(wr, map[string]interface{}{
			"mode":    mode,
			"name":    name,
			"value":   value,
			"options": options,
		}); err != nil {
			return template.HTML(err.Error())
		}

		return template.HTML(wr.String())
	}

	Many = func(belong string, many string, relation string, mode string, id int64, value interface{}, options ...interface{}) template.HTML {
		tmpl, err := revel.MainTemplateLoader.Template(fmt.Sprintf("view.%s/many.html", symbol.ModuleName(many)))
		if err != nil {
			return template.HTML(fmt.Sprintf("view.%s/many.html not exist.", symbol.ModuleName(many)))
		}

		wr := bytes.NewBuffer([]byte(""))
		if err := tmpl.Render(wr, map[string]interface{}{
			"belong":   belong,
			"many":     many,
			"relation": relation,
			"mode":     mode,
			"id":       id,
			"value":    value,
			"options":  options,
		}); err != nil {
			return template.HTML(err.Error())
		}

		return template.HTML(wr.String())
	}

	DomID = func(src string) string {
		return strings.Replace(src, ".", "_", -1)
	}

	DateFormat = func() string {
		return revel.Config.StringDefault("format.date", "2006/01/02")
	}

	DatetimeFormat = func() string {
		return revel.Config.StringDefault("format.datetime", "2006/01/02 15:04:05")
	}

	JsonEncode = func(value interface{}) string {
		if b, err := json.Marshal(value); err == nil {
			return string(b)
		}

		return ""
	}

	SecureEscape = func(src interface{}) string {
		return template.HTMLEscapeString(template.JSEscapeString(fmt.Sprintf("%v", src)))
	}

	JavascriptEscape = func(src interface{}) string {
		return template.JSEscapeString(fmt.Sprintf("%v", src))
	}

	Pagination = func(widget, name string, total, page_size, page_no int64) template.HTML {
		pages := []Page{}
		if page_size == 0 || total == 0 {
			return template.HTML("")
		}

		num := total / page_size
		if total%page_size > 0 {
			num = num + 1
		}

		var i int64 = 0
		for ; i < num; i++ {
			var page Page
			page.No = i
			page.Size = page_size
			page.Normal = true
			if i == 0 {
				page.Head = true
				page.Normal = false
			}
			if i == page_no {
				page.Selected = true
			}
			if i+1 == num {
				page.Tail = true
				page.Normal = false
			}
			pages = append(pages, page)
		}
		return Widget(widget, "r", DomID(name), map[string]interface{}{
			"total":     total,
			"page_size": page_size,
			"pages":     pages,
		})
	}
)

func init() {
	TimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	revel.TypeBinders[reflect.TypeOf(time.Time{})] = TimeLocationBinder
	revel.TemplateFuncs["widget"] = Widget
	revel.TemplateFuncs["many"] = Many
	revel.TemplateFuncs["domid"] = DomID
	revel.TemplateFuncs["date_format"] = DateFormat
	revel.TemplateFuncs["datetime_format"] = DatetimeFormat
	revel.TemplateFuncs["module"] = symbol.ModuleName
	revel.TemplateFuncs["class"] = symbol.ClassName
	revel.TemplateFuncs["camel"] = symbol.Camel
	revel.TemplateFuncs["lint"] = symbol.Lint
	revel.TemplateFuncs["json"] = JsonEncode
	revel.TemplateFuncs["pagination"] = Pagination
	revel.TemplateFuncs["secure_escape"] = SecureEscape
	revel.TemplateFuncs["javascript_escape"] = JavascriptEscape
}
