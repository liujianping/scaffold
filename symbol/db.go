package symbol

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func DSNFormat(host string,
	port int,
	user string,
	pass string,
	database string) string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&strict=true&sql_notes=false",
		user, pass, host, port, database)
}

func sqlstr(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

type TypeConvert interface {
	Convert(string) string
}

type Table struct {
	name    string
	rows    int64
	comment sql.NullString
	columns []*Column
}

type Column struct {
	sql_field      sql.NullString
	sql_type       sql.NullString
	sql_collation  sql.NullString
	sql_null       sql.NullString
	sql_key        sql.NullString
	sql_default    sql.NullString
	sql_extra      sql.NullString
	sql_privileges sql.NullString
	sql_comment    sql.NullString
}

type Option struct {
	Name        string
	Code        string
	OptionName  string
	OptionCode  string
	OptionValue int64
}

func GetOptions(db *sql.DB) (map[string][]Option, error) {
	options := make(map[string][]Option)

	rows, err := db.Query(`SELECT name, code, option_code, option_value , option_name FROM options`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var opt Option
		if err := rows.Scan(&opt.Name, &opt.Code, &opt.OptionCode, &opt.OptionValue, &opt.OptionName); err != nil {
			log.Println("scan failed:", err)
			return nil, err
		}

		if sel, ok := options[opt.Code]; ok {
			sel = append(sel, opt)
			options[opt.Code] = sel
		} else {
			sel = []Option{}
			sel = append(sel, opt)
			options[opt.Code] = sel
		}
	}

	return options, nil
}

func GetAllTables(db *sql.DB) ([]*Table, error) {
	var tables []*Table

	// table comment
	rows_tb, err := db.Query(`select table_name, table_rows, table_comment from information_schema.tables where table_schema = DATABASE()`)
	if err != nil {
		return nil, err
	}
	defer rows_tb.Close()

	for rows_tb.Next() {
		var table Table
		if err := rows_tb.Scan(&table.name, &table.rows, &table.comment); err != nil {
			return nil, err
		}

		rows_col, err := db.Query(fmt.Sprintf("SHOW FULL COLUMNS FROM %s", table.name))
		if err != nil {
			return nil, err
		}
		for rows_col.Next() {
			var column Column
			if err := rows_col.Scan(&column.sql_field,
				&column.sql_type,
				&column.sql_collation,
				&column.sql_null,
				&column.sql_key,
				&column.sql_default,
				&column.sql_extra,
				&column.sql_privileges,
				&column.sql_comment); err != nil {
				return nil, err
			}

			table.columns = append(table.columns, &column)
		}
		rows_col.Close()

		tables = append(tables, &table)
	}

	return tables, nil
}

func GetTable(db *sql.DB, name string) (*Table, error) {
	var table Table
	table.name = name

	// table comment
	rows_tb, err := db.Query(`select table_comment 
                           from information_schema.tables 
                           where table_schema = DATABASE() and table_name = ?`,
		name)
	if err != nil {
		return nil, err
	}
	defer rows_tb.Close()

	for rows_tb.Next() {
		if err := rows_tb.Scan(&table.comment); err != nil {
			return nil, err
		}
	}

	// table columns
	rows_col, err := db.Query(fmt.Sprintf("SHOW FULL COLUMNS  FROM %s", name))
	if err != nil {
		return nil, err
	}
	defer rows_col.Close()

	for rows_col.Next() {
		var column Column
		if err := rows_col.Scan(&column.sql_field,
			&column.sql_type,
			&column.sql_collation,
			&column.sql_null,
			&column.sql_key,
			&column.sql_default,
			&column.sql_extra,
			&column.sql_privileges,
			&column.sql_comment); err != nil {
			return nil, err
		}

		table.columns = append(table.columns, &column)
	}

	return &table, nil
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) Comment() string {
	return sqlstr(t.comment)
}

func (t *Table) Tag(tag string) string {
	return StructTag(sqlstr(t.comment)).Get(tag)
}

func (t *Table) Columns() []*Column {
	return t.columns
}

func (t *Table) Column(name string) *Column {
	for _, col := range t.columns {
		if col.sql_field.Valid &&
			strings.ToLower(col.sql_field.String) == strings.ToLower(name) {
			return col
		}
	}
	return nil
}

func (t *Table) PrimaryColumn() *Column {
	for _, col := range t.columns {
		if col.IsPrimary() {
			return col
		}
	}
	return nil
}

func (t *Table) Fields() []string {
	var fields []string
	for _, col := range t.columns {
		fields = append(fields, col.Field())
	}
	return fields
}

func (t *Table) FieldComment(field string) string {
	col := t.Column(field)
	if col != nil {
		return sqlstr(col.sql_comment)
	}

	return ""
}

func (t *Table) FieldComments() map[string]string {
	var results map[string]string = make(map[string]string)
	for _, col := range t.columns {
		results[sqlstr(col.sql_field)] = sqlstr(col.sql_comment)
	}
	return results
}

func (c *Column) IsPrimary() bool {
	if sqlstr(c.sql_key) == "PRI" {
		return true
	}
	return false
}

func (c *Column) Field() string {
	return sqlstr(c.sql_field)
}

func (c *Column) Comment() string {
	return sqlstr(c.sql_comment)
}

func (c *Column) Tag(tag string) string {
	return StructTag(sqlstr(c.sql_comment)).Get(tag)
}

func (c *Column) Type() string {
	return sqlstr(c.sql_type)
}

func (c *Column) Default() string {
	return sqlstr(c.sql_default)
}
