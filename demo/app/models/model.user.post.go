package models

import (
	"reflect"
	"time"

	bsql "github.com/liujianping/bsql"
	revel "github.com/revel/revel"
)

var (
	DefaultUserPost       = UserPost{}
	DefaultUserPostQuery  = UserPostQuery{}
	DefaultUserPostSortBy = UserPostSortBy{1}
	DefaultUserPostPage   = UserPostPage{0, DefaultPageSize}
)

//! ===========================================================================
type UserPost struct {
	ID            int64     `db:"id"    json:"id"`
	UserAccountID int64     `db:"user_account_id"    json:"user_account_id"`
	Title         string    `db:"title"    json:"title"`
	Content       string    `db:"content"    json:"content"`
	ImageURL      string    `db:"image_url"    json:"image_url"`
	Status        int64     `db:"status"    json:"status"`
	CreateAt      time.Time `db:"create_at"    json:"create_at"`
	UpdateAt      time.Time `db:"update_at"    json:"update_at"`
}

//! model table_name
func (obj UserPost) TableName() string {
	return "user_posts"
}

func (obj UserPost) PrimaryKey() string {
	return "id"
}

func (obj UserPost) Columns() []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}
	return columns
}

func (obj UserPost) Count(m *Model) int64 {
	t := bsql.TABLE(obj.TableName())
	SQL := bsql.NewQuerySQL(t)
	stmt := SQL.CountStatment()
	count, _ := m.SelectInt(stmt.SQLFormat(), stmt.SQLParams()...)
	return count
}

func (obj UserPost) Search(m *Model,
	query UserPostQuery,
	sort UserPostSortBy,
	page UserPostPage) (total int64, results []UserPost, err error) {

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
func (obj UserPost) Validate(v *revel.Validation) {

	v.Check(obj.Title, required(), min(12))

}

//! ===========================================================================

//! model query
type UserPostQuery struct {
	UserAccountID int64 `db:"user_account_id"    query:"eq"`

	Title string `db:"title"    query:"like"`

	Status int64 `db:"status"    query:"eq"`
}

func (obj UserPostQuery) SQL(query *bsql.QuerySQL) {
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
type UserPostSortBy struct {
	Value int64
}

func (sortBy UserPostSortBy) SQL(query *bsql.QuerySQL) {
	t := query.Table(DefaultUserPost.TableName())
	switch sortBy.Value {
	case 1:
		query.OrderByDesc(t.Column(DefaultUserPost.PrimaryKey()))
	case 2:
		query.OrderByAsc(t.Column(DefaultUserPost.PrimaryKey()))
	}
}

//! ===========================================================================
//! model page
type UserPostPage struct {
	No   int64
	Size int64
}

func (page UserPostPage) SQL(query *bsql.QuerySQL) {
	query.Limit(page.No, page.Size)
}

//! ===========================================================================
//! model init
func init() {
	registTableObject(DefaultUserPost.TableName(),
		DefaultUserPost.PrimaryKey(),
		DefaultUserPost)
}
