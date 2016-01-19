package models

import (
	"reflect"

	bsql "github.com/liujianping/bsql"
	revel "github.com/revel/revel"
)

var (
	DefaultSelection       = Selection{}
	DefaultSelectionQuery  = SelectionQuery{}
	DefaultSelectionSortBy = SelectionSortBy{1}
	DefaultSelectionPage   = SelectionPage{0, DefaultPageSize}
)

//! ===========================================================================
type Selection struct {
	ID          int64  `db:"id"    json:"id"`
	Name        string `db:"name"    json:"name"`
	Code        string `db:"code"    json:"code"`
	OptionName  string `db:"option_name"    json:"option_name"`
	OptionCode  string `db:"option_code"    json:"option_code"`
	OptionValue int64  `db:"option_value"    json:"option_value"`
	Description string `db:"description"    json:"description"`
}

//! model table_name
func (obj Selection) TableName() string {
	return "selections"
}

func (obj Selection) PrimaryKey() string {
	return "id"
}

func (obj Selection) Columns() []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}
	return columns
}

func (obj Selection) Count(m *Model) int64 {
	t := bsql.TABLE(obj.TableName())
	SQL := bsql.NewQuerySQL(t)
	stmt := SQL.CountStatment()
	count, _ := m.SelectInt(stmt.SQLFormat(), stmt.SQLParams()...)
	return count
}

func (obj Selection) Search(m *Model,
	query SelectionQuery,
	sort SelectionSortBy,
	page SelectionPage) (total int64, results []Selection, err error) {

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
func (obj Selection) Validate(v *revel.Validation) {

	v.Check(obj.Name, required())
	v.Check(obj.Code, required())
	v.Check(obj.OptionName, required())
	v.Check(obj.OptionCode, required())
	v.Check(obj.OptionValue, required())

}

//! ===========================================================================

//! model query
type SelectionQuery struct {
	Name string `db:"name"    query:"like"`

	Code string `db:"code"    query:"eq"`

	OptionValue int64 `db:"option_value"    query:"eq"`
}

func (obj SelectionQuery) SQL(query *bsql.QuerySQL) {
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
type SelectionSortBy struct {
	Value int64
}

func (sortBy SelectionSortBy) SQL(query *bsql.QuerySQL) {
	t := query.Table(DefaultSelection.TableName())
	switch sortBy.Value {
	case 1:
		query.OrderByDesc(t.Column(DefaultSelection.PrimaryKey()))
	case 2:
		query.OrderByAsc(t.Column(DefaultSelection.PrimaryKey()))
	}
}

//! ===========================================================================
//! model page
type SelectionPage struct {
	No   int64
	Size int64
}

func (page SelectionPage) SQL(query *bsql.QuerySQL) {
	query.Limit(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(DefaultSelection.TableName(),
		DefaultSelection.PrimaryKey(),
		DefaultSelection)
}
