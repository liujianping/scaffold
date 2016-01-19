package controllers

import (
	"bytes"
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

	Widget = func(widget string, name string, value interface{}, options ...string) template.HTML {
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
			"name":    name,
			"value":   value,
			"options": options,
		}); err != nil {
			return template.HTML(err.Error())
		}

		return template.HTML(wr.String())
	}

	DomID = func(src string) string {
		return strings.Replace(src, ".", "-", -1)
	}

	DateFormat = func() string {
		return revel.Config.StringDefault("format.date", "yyyy/MM/dd")
	}

	DatetimeFormat = func() string {
		return revel.Config.StringDefault("format.datetime", "yyyy/MM/dd hh:mm:ss")
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
		return Widget(widget, DomID(name), map[string]interface{}{
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
	revel.TemplateFuncs["domid"] = DomID
	revel.TemplateFuncs["date_format"] = DateFormat
	revel.TemplateFuncs["datetime_format"] = DatetimeFormat
	revel.TemplateFuncs["module"] = symbol.ModuleName
	revel.TemplateFuncs["class"] = symbol.ClassName
	revel.TemplateFuncs["pagination"] = Pagination
}
