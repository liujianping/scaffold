package models
[[set . "t_class" (.table.Name | singular | camel)]]
[[set . "t_module" (.table.Name | module)]]
import (
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var (
	Default[[.t_class]]      = [[.t_class]]{}
	Default[[.t_class]]Query = [[.t_class]]Query{}
	Default[[.t_class]]Sort  = [[.t_class]]Sort{1}
	Default[[.t_class]]Page  = [[.t_class]]Page{0, 25}
)

type [[.t_class]] struct {
	[[range .table.Columns]]
    [[.Field | camel | lint]]	[[convert "mysql" .Type (.Tag "gotype")]]	`db:"[[.Field]]"    json:"[[.Field]]"`[[end]]
}

func (obj [[.t_class]]) TableName() string {
	return "[[.table.Name]]"
}

func (obj [[.t_class]]) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&obj).Count(&total)
	return total
}

func (obj [[.t_class]]) Validate(v *revel.Validation) {[[range .table.Columns]]
    [[if ne (.Tag "valid") ""]]v.Check(obj.[[.Field | camel | lint]], [[.Tag "valid"]])[[end]][[end]]
}

func (obj [[.t_class]]) Search(db *gorm.DB,
	query [[.t_class]]Query,
	sort [[.t_class]]Sort,
	page [[.t_class]]Page) (total int64, results [][[.t_class]], err error) {

	search := db.Model(&obj).Select(Columns(obj))
	search = query.Query(search)
	search.Count(&total)
	search = page.Page(sort.Sort(search))
	search.Find(&results)
	err = search.Error
	return
}

type [[.t_class]]Query struct {[[range .table.Columns]]
	[[if or (eq (.Tag "query") "range") (eq (.Tag "find") "range")]][[.Field | camel | lint]]From    [[convert "mysql" .Type (.Tag "gotype")]]     `db:"[[.Field]]"    query:"gt"`
    [[.Field | camel | lint]]To 	[[convert "mysql" .Type (.Tag "gotype")]]     `db:"[[.Field]]"    query:"lt"`
    [[else if or (ne (.Tag "query") "") (ne (.Tag "find") "")]][[.Field | camel | lint]]		[[convert "mysql" .Type (.Tag "gotype")]]    `db:"[[.Field]]"    query:"[[.Tag "query"]]"`[[end]][[end]]
}


type [[.t_class]]Sort struct {
	Value int64
}

type [[.t_class]]Page struct {
	No   int64
	Size int64
}

func (obj [[.t_class]]Query) Query(db *gorm.DB) *gorm.DB {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		value := v.Field(i)
		if IsZero(value) == false {
			db = db.Where(BinaryExpression(tag.Get("db"), tag.Get("query"), value.Interface()).Statment())
		}
	}
	return db
}

func (obj [[.t_class]]Sort) Sort(db *gorm.DB) *gorm.DB {
	switch obj.Value {[[range $k, $v := .table.Columns]]
	[[if $v.IsPrimary]]case [[add (multiply 2 $k) 1]]:
		db = db.Order("[[$v.Field]] DESC")
	case [[add (multiply 2 $k) 2]]:
		db = db.Order("[[$v.Field]] ASC")
	[[else if eq ($v.Tag "sort") "y"]]case [[add (multiply 2 $k) 1]]:
		db = db.Order("[[$v.Field]] DESC")
	case [[add (multiply 2 $k) 2]]:
		db = db.Order("[[$v.Field]] ASC")
	[[end]][[end]]
	}
	return db
}

func (obj [[.t_class]]Page) Page(db *gorm.DB) *gorm.DB {
	db = db.Offset(int(obj.No * obj.Size))
	db = db.Limit(int(obj.Size))
	return db
}
