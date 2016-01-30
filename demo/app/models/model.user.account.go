package models

import (
	"reflect"
	"time"

	bsql "github.com/liujianping/bsql"
	revel "github.com/revel/revel"
)

var (
	DefaultUserAccount       = UserAccount{}
	DefaultUserAccountQuery  = UserAccountQuery{}
	DefaultUserAccountSortBy = UserAccountSortBy{1}
	DefaultUserAccountPage   = UserAccountPage{0, DefaultPageSize}
)

//! ===========================================================================
type UserAccount struct {
	ID          int64     `db:"id"    json:"id"`
	Name        string    `db:"name"    json:"name"`
	Mailbox     string    `db:"mailbox"    json:"mailbox"`
	Sex         int64     `db:"sex"    json:"sex"`
	Description string    `db:"description"    json:"description"`
	Password    string    `db:"password"    json:"password"`
	HeadURL     string    `db:"head_url"    json:"head_url"`
	Status      int64     `db:"status"    json:"status"`
	CreateAt    time.Time `db:"create_at"    json:"create_at"`
	UpdateAt    time.Time `db:"update_at"    json:"update_at"`
}

//! model table_name
func (obj UserAccount) TableName() string {
	return "user_accounts"
}

func (obj UserAccount) PrimaryKey() string {
	return "id"
}

func (obj UserAccount) Columns() []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}
	return columns
}

func (obj UserAccount) Count(m *Model) int64 {
	t := bsql.TABLE(obj.TableName())
	SQL := bsql.NewQuerySQL(t)
	stmt := SQL.CountStatment()
	count, _ := m.SelectInt(stmt.SQLFormat(), stmt.SQLParams()...)
	return count
}

func (obj UserAccount) Search(m *Model,
	query UserAccountQuery,
	sort UserAccountSortBy,
	page UserAccountPage) (total int64, results []UserAccount, err error) {

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
func (obj UserAccount) Validate(v *revel.Validation) {

	v.Check(obj.Name, required(), min(6), max(16))
	v.Check(obj.Mailbox, required(), email())

	v.Check(obj.Password, required())

}

//! ===========================================================================

//! model query
type UserAccountQuery struct {
	Name    string `db:"name"    query:"like"`
	Mailbox string `db:"mailbox"    query:"like"`

	Status int64 `db:"status"    query:"eq"`
}

func (obj UserAccountQuery) SQL(query *bsql.QuerySQL) {
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
type UserAccountSortBy struct {
	Value int64
}

func (sortBy UserAccountSortBy) SQL(query *bsql.QuerySQL) {
	t := query.Table(DefaultUserAccount.TableName())
	switch sortBy.Value {
	case 1:
		query.OrderByDesc(t.Column(DefaultUserAccount.PrimaryKey()))
	case 2:
		query.OrderByAsc(t.Column(DefaultUserAccount.PrimaryKey()))
	}
}

//! ===========================================================================
//! model page
type UserAccountPage struct {
	No   int64
	Size int64
}

func (page UserAccountPage) SQL(query *bsql.QuerySQL) {
	query.Page(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(DefaultUserAccount.TableName(),
		DefaultUserAccount.PrimaryKey(),
		DefaultUserAccount)
}
