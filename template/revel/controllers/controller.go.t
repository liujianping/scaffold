package controllers

import (
	"fmt"
	"reflect"
	"time"

	"github.com/revel/revel"
	"[[.project]]/app/models"
)

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
)

var (
	model *models.Model
)

func InitModel() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	host := revel.Config.StringDefault("db.host", "127.0.0.1")
	port := revel.Config.IntDefault("db.port", 3306)
	user := revel.Config.StringDefault("db.username", "")
	pass := revel.Config.StringDefault("db.password", "")
	name := revel.Config.StringDefault("db.datebase", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&strict=true&sql_notes=false",
		user, pass, host, port, name)

	m, err := models.NewModel(driver, dsn)
	if err != nil {
		panic(err)
	}

	model = m
	model.TraceOn("[model]", revel.INFO)
}

func init() {
	TimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	revel.TypeBinders[reflect.TypeOf(time.Time{})] = TimeLocationBinder
	revel.OnAppStart(InitModel)
}
