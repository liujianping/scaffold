package models
[[set . "ClassName" (.table.Name | singular | camel)]]
[[set . "ModuleName" (.table.Name | module)]]
import (
	"reflect"
	"time"

	revel "github.com/revel/revel"
)

var (
	Default[[.ClassName]]       = [[.ClassName]]{}
	Default[[.ClassName]]Query  = [[.ClassName]]Query{}
	Default[[.ClassName]]SortBy = [[.ClassName]]SortBy{1}
	Default[[.ClassName]]Page   = [[.ClassName]]Page{0, DefaultPageSize}
)

//! ===========================================================================
type [[.ClassName]] struct {
    [[range .table.Columns]][[.Field | camel | lint]]   [[.Type | convert "mysql"]] `db:"[[.Field]]"    json:"[[.Field]]"`
    [[end]]
}

//! model table_name
func (obj [[.ClassName]]) TableName() string {
	return "[[.table.Name]]"
}

func (obj [[.ClassName]]) PrimaryKey() string {
	return "[[.table.PrimaryColumn.Field]]"
}

func (obj [[.ClassName]]) Columns() []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}
	return columns
}

func (obj [[.ClassName]]) Execute(m *Model,
	query [[.ClassName]]Query,
	sort [[.ClassName]]SortBy,
	page [[.ClassName]]Page) (total int64, results [][[.ClassName]], err error) {

	SQL := NewQuerySQL(obj.TableName())
	SQL.PrimaryKey(obj.PrimaryKey())
	SQL.Columns(obj.Columns()...)

	query.SQL(SQL)
	sort.SQL(SQL)
	page.SQL(SQL)

	count := SQL.CountStatment()

	total, err = m.SelectInt(count.SQLFormat(), count.SQLParams()...)
	if err != nil {
		return
	}

	statment := SQL.QueryStatment()
	_, err = m.Select(&results, statment.SQLFormat(), statment.SQLParams()...)
	return
}

//! model validation
func (obj [[.ClassName]]) Validate(v *revel.Validation) {
    [[range .table.Columns]][[if ne (.Tag "valid") ""]]v.Check(obj.[[.Field | camel | lint]], [[.Tag "valid"]])[[end]]
    [[end]]
}

//! ===========================================================================

//! model query
type [[.ClassName]]Query struct {
    [[range .table.Columns]][[if ne (.Tag "query") ""]][[.Field | camel | lint]]    [[.Type | convert "mysql"]]     `db:"[[.Field]]"    query:"[[.Tag "query"]]"`[[end]]
    [[end]]
}

func (obj [[.ClassName]]Query) SQL(sql *QuerySQL) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		value := v.Field(i)
		if ZeroValue(value) == false {
			sql.Where(BinaryExpression(tag.Get("db"), tag.Get("query"), value.Interface()))
		}
	}
}

//! ===========================================================================
//! model query sortby
type [[.ClassName]]SortBy struct {
	Value int
}

func (sortBy [[.ClassName]]SortBy) SQL(sql *QuerySQL) {
	switch sortBy.Value {
	case 1:
		sql.OrderByDesc(Default[[.ClassName]].PrimaryKey())
	case 2:
		sql.OrderByAsc(Default[[.ClassName]].PrimaryKey())
	}
}

//! ===========================================================================
//! model page
type [[.ClassName]]Page struct {
	No   int
	Size int
}

func (page [[.ClassName]]Page) SQL(sql *QuerySQL) {
	if page.Size == 0 {
		page.Size = DefaultPageSize
	}
	sql.Limit(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(Default[[.ClassName]].TableName(),
		Default[[.ClassName]].PrimaryKey(),
		Default[[.ClassName]])
}
