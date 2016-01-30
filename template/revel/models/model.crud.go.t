package models
[[set . "ClassName" (.table.Name | singular | camel)]]
[[set . "ModuleName" (.table.Name | module)]]
import (
	"reflect"
	"time"

	bsql "github.com/liujianping/bsql"
	revel "github.com/revel/revel"
)

var (
	Default[[.ClassName]]       = [[.ClassName]]{}
	Default[[.ClassName]]Query  = [[.ClassName]]Query{}
	Default[[.ClassName]]SortBy = [[.ClassName]]SortBy{1}
	Default[[.ClassName]]Page   = [[.ClassName]]Page{0, DefaultPageSize}
)

//! ===========================================================================
type [[.ClassName]] struct {[[range .table.Columns]]
    [[.Field | camel | lint]]   [[.Type | convert "mysql"]] `db:"[[.Field]]"    json:"[[.Field]]"`[[end]]
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

func (obj [[.ClassName]]) Count(m *Model) int64 {
	t := bsql.TABLE(obj.TableName())
	SQL := bsql.NewQuerySQL(t)
	stmt := SQL.CountStatment()
	count, _ := m.SelectInt(stmt.SQLFormat(), stmt.SQLParams()...)
	return count
}

func (obj [[.ClassName]]) Search(m *Model,
	query [[.ClassName]]Query,
	sort [[.ClassName]]SortBy,
	page [[.ClassName]]Page) (total int64, results [][[.ClassName]], err error) {

	t := bsql.TABLE(obj.TableName())
	t.PrimaryKey(obj.PrimaryKey())
	t.Columns(obj.Columns()...)

	SQL := bsql.NewQuerySQL(t)

	query.SQL(SQL)
	sort.SQL(SQL)
	page.SQL(SQL)

	count := SQL.CountStatment()

	total, err = m.SelectInt(count.SQLFormat(), count.SQLParams()...)
	if err != nil {
		return
	}

	statment := SQL.Statment()
	_, err = m.Select(&results, statment.SQLFormat(), statment.SQLParams()...)
	return
}

//! model validation
func (obj [[.ClassName]]) Validate(v *revel.Validation) {
    [[range .table.Columns]]
    [[if ne (.Tag "valid") ""]]v.Check(obj.[[.Field | camel | lint]], [[.Tag "valid"]])[[end]][[end]]
}

//! ===========================================================================

//! model query
type [[.ClassName]]Query struct {
    [[range .table.Columns]]
    [[if ne (.Tag "query") ""]][[.Field | camel | lint]]    [[.Type | convert "mysql"]]     `db:"[[.Field]]"    query:"[[.Tag "query"]]"`[[else if ne (.Tag "find") ""]]
    [[.Field | camel | lint]]    [[.Type | convert "mysql"]]     `db:"[[.Field]]"    query:"[[.Tag "query"]]"`[[end]][[end]]
}

func (obj [[.ClassName]]Query) SQL(query *bsql.QuerySQL) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		value := v.Field(i)
		if ZeroValue(value) == false {
			query.Where(bsql.BinaryExpression(tag.Get("db"), tag.Get("query"), value.Interface()))
		}
	}
}

//! ===========================================================================
//! model query sortby
type [[.ClassName]]SortBy struct {
	Value int64
}

func (sortBy [[.ClassName]]SortBy) SQL(query *bsql.QuerySQL) {	
	t := query.Table(Default[[.ClassName]].TableName())
	switch sortBy.Value {
	case 1:
		query.OrderByDesc(t.Column(Default[[.ClassName]].PrimaryKey()))
	case 2:
		query.OrderByAsc(t.Column(Default[[.ClassName]].PrimaryKey()))				
	}
}

//! ===========================================================================
//! model page
type [[.ClassName]]Page struct {
	No   int64
	Size int64
}

func (page [[.ClassName]]Page) SQL(query *bsql.QuerySQL) {
	query.Page(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(Default[[.ClassName]].TableName(),
		Default[[.ClassName]].PrimaryKey(),
		Default[[.ClassName]])
}
