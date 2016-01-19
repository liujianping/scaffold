package models

import (
	"reflect"

	bsql "github.com/liujianping/bsql"
	revel "github.com/revel/revel"
)

var (
	DefaultOption       = Option{}
	DefaultOptionQuery  = OptionQuery{}
	DefaultOptionSortBy = OptionSortBy{1}
	DefaultOptionPage   = OptionPage{0, DefaultPageSize}
)

//! ===========================================================================
type Option struct {
	ID          int64  `db:"id"    json:"id"`
	Name        string `db:"name"    json:"name"`
	Code        string `db:"code"    json:"code"`
	OptionName  string `db:"option_name"    json:"option_name"`
	OptionCode  string `db:"option_code"    json:"option_code"`
	OptionValue int64  `db:"option_value"    json:"option_value"`
	Description string `db:"description"    json:"description"`
}

//! model table_name
func (obj Option) TableName() string {
	return "options"
}

func (obj Option) PrimaryKey() string {
	return "id"
}

func (obj Option) Columns() []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}
	return columns
}

func (obj Option) Count(m *Model) int64 {
	t := bsql.TABLE(obj.TableName())
	SQL := bsql.NewQuerySQL(t)
	stmt := SQL.CountStatment()
	count, _ := m.SelectInt(stmt.SQLFormat(), stmt.SQLParams()...)
	return count
}

func (obj Option) Search(m *Model,
	query OptionQuery,
	sort OptionSortBy,
	page OptionPage) (total int64, results []Option, err error) {

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
func (obj Option) Validate(v *revel.Validation) {

	v.Check(obj.Name, required())
	v.Check(obj.Code, required())
	v.Check(obj.OptionName, required())
	v.Check(obj.OptionCode, required())
	v.Check(obj.OptionValue, required())

}

//! ===========================================================================

//! model query
type OptionQuery struct {
	Name string `db:"name"    query:"like"`

	Code string `db:"code"    query:"eq"`

	OptionValue int64 `db:"option_value"    query:"eq"`
}

func (obj OptionQuery) SQL(query *bsql.QuerySQL) {
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
type OptionSortBy struct {
	Value int64
}

func (sortBy OptionSortBy) SQL(query *bsql.QuerySQL) {
	t := query.Table(DefaultOption.TableName())
	switch sortBy.Value {
	case 1:
		query.OrderByDesc(t.Column(DefaultOption.PrimaryKey()))
	case 2:
		query.OrderByAsc(t.Column(DefaultOption.PrimaryKey()))
	}
}

//! ===========================================================================
//! model page
type OptionPage struct {
	No   int64
	Size int64
}

func (page OptionPage) SQL(query *bsql.QuerySQL) {
	query.Limit(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(DefaultOption.TableName(),
		DefaultOption.PrimaryKey(),
		DefaultOption)
}
