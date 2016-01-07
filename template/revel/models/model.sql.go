package models

import (
	"fmt"
	"log"
	"strings"
)

type Statment struct {
	format string
	values []interface{}
}

func (statment *Statment) SQLFormat() string {
	return statment.format
}

func (statment *Statment) SQLParams() []interface{} {
	return statment.values
}

func Join(statments []*Statment, sep string) *Statment {
	if len(statments) == 0 {
		return nil
	}

	if len(statments) == 1 {
		return statments[0]
	}

	var fmts []string
	var vals []interface{}
	for _, statment := range statments {
		fmts = append(fmts, statment.format)
		vals = append(vals, statment.values...)
	}

	return &Statment{
		format: strings.Join(fmts, sep),
		values: vals,
	}
}

type QuerySQL struct {
	table      string
	primary    string
	columns    []string
	conditions []Expression
	sort       string
	order_by   []string
	group_by   []string
	offset     int
	limit      int
}

func NewQuerySQL(table string) *QuerySQL {
	return &QuerySQL{
		table:      table,
		columns:    []string{},
		conditions: []Expression{},
		order_by:   []string{},
		group_by:   []string{},
	}
}

func (sql *QuerySQL) PrimaryKey(pk string) *QuerySQL {
	sql.primary = pk
	return sql
}

func (sql *QuerySQL) Columns(columns ...string) *QuerySQL {
	if len(columns) > 0 {
		sql.columns = append(sql.columns, columns...)
	}
	return sql
}

func (sql *QuerySQL) Where(exps ...Expression) *QuerySQL {
	if len(exps) > 0 {
		sql.conditions = append(sql.conditions, exps...)
	}
	return sql
}

func (sql *QuerySQL) OrderByAsc(fields ...string) *QuerySQL {
	if sql.sort != "ASC" {
		sql.order_by = []string{}
		sql.sort = "ASC"
	}
	sql.order_by = append(sql.order_by, fields...)
	return sql
}

func (sql *QuerySQL) OrderByDesc(fields ...string) *QuerySQL {
	if sql.sort != "DESC" {
		sql.order_by = []string{}
		sql.sort = "DESC"
	}
	sql.order_by = append(sql.order_by, fields...)
	return sql
}

func (sql *QuerySQL) Limit(page_no, page_size int) *QuerySQL {
	sql.limit = page_size
	sql.offset = page_no * page_size
	return sql
}

func (sql *QuerySQL) GroupBy(fields ...string) *QuerySQL {
	if len(fields) > 0 {
		sql.group_by = append(sql.group_by, fields...)
	}
	return sql
}

func (sql *QuerySQL) CountStatment() *Statment {
	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "SELECT")
	if sql.primary != "" {
		fmts = append(fmts, fmt.Sprintf("COUNT(%s)", sql.primary))
	} else {
		fmts = append(fmts, fmt.Sprintf("COUNT(%s)", "*"))
	}
	fmts = append(fmts, "FROM")
	fmts = append(fmts, sql.table)

	if len(sql.conditions) > 0 {
		condition := AND(sql.conditions...).Statment()
		if condition != nil {
			fmts = append(fmts, "WHERE")
			fmts = append(fmts, condition.SQLFormat())
			vals = append(vals, condition.SQLParams()...)
		}
	}

	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}

func (sql *QuerySQL) QueryStatment() *Statment {
	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "SELECT")
	if len(sql.columns) > 0 {
		fmts = append(fmts, fmt.Sprintf("%s", strings.Join(sql.columns, ", ")))
	} else {
		fmts = append(fmts, fmt.Sprintf("%s", "*"))
	}
	fmts = append(fmts, "FROM")
	fmts = append(fmts, sql.table)

	if len(sql.conditions) > 0 {
		condition := AND(sql.conditions...).Statment()
		if condition != nil {
			fmts = append(fmts, "WHERE")
			fmts = append(fmts, condition.SQLFormat())
			vals = append(vals, condition.SQLParams()...)
		}
	}

	if len(sql.group_by) > 0 {
		fmts = append(fmts, "GROUP BY")
		fmts = append(fmts, strings.Join(sql.group_by, ", "))
	}

	if len(sql.order_by) > 0 {
		fmts = append(fmts, "ORDER BY")
		fmts = append(fmts, strings.Join(sql.order_by, ", "))
		fmts = append(fmts, sql.sort)
	}

	if sql.limit > 0 {
		fmts = append(fmts, "LIMIT ?,?")
		vals = append(vals, sql.offset, sql.limit)
	}

	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}

type Expression struct {
	insides []Expression
	field   string
	symbol  string
	value   interface{}
}

func BinaryExpression(field, symbol string, value interface{}) Expression {
	log.Println(field, symbol, value)
	switch strings.ToUpper(symbol) {
	case "LIKE":
		return LIKE(field, value)
	case "LLIKE":
		return LLIKE(field, value)
	case "RLIKE":
		return RLIKE(field, value)
	case "EQ":
		return EQ(field, value)
	case "NE":
		return NE(field, value)
	case "IN":
		return IN(field, value)
	case "GT":
		return GT(field, value)
	case "GE":
		return GE(field, value)
	case "LT":
		return LT(field, value)
	case "LE":
		return LE(field, value)
	}
	return NONE(field, value)
}

func (exp Expression) Statment() *Statment {
	switch strings.ToUpper(exp.symbol) {
	case "AND":
		var statments []*Statment
		for _, inside := range exp.insides {
			if inside.Statment() != nil {
				statments = append(statments, inside.Statment())
			}
		}
		return Join(statments, " AND ")
	case "OR":
		var statments []*Statment
		for _, inside := range exp.insides {
			if inside.Statment() != nil {
				statments = append(statments, inside.Statment())
			}
		}
		return Join(statments, " OR ")
	case "LIKE", "LLIKE", "RLIKE":
		return &Statment{
			format: fmt.Sprintf("%s LIKE ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "EQ":
		return &Statment{
			format: fmt.Sprintf("%s = ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "NE":
		return &Statment{
			format: fmt.Sprintf("%s != ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "GT":
		return &Statment{
			format: fmt.Sprintf("%s > ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "GE":
		return &Statment{
			format: fmt.Sprintf("%s >= ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "LT":
		return &Statment{
			format: fmt.Sprintf("%s < ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "LE":
		return &Statment{
			format: fmt.Sprintf("%s <= ?", exp.field),
			values: []interface{}{exp.value},
		}
	case "IN":
		return nil
	}
	return nil
}

func AND(exps ...Expression) Expression {
	return Expression{
		insides: exps,
		symbol:  "AND",
	}
}

func OR(exps ...Expression) Expression {
	return Expression{
		insides: exps,
		symbol:  "OR",
	}
}

func NONE(field string, value interface{}) Expression {
	return Expression{
		field:  "",
		symbol: "",
		value:  "",
	}
}

func RANGE(field string, start, end interface{}) Expression {
	return AND(GT(field, start), LT(field, end))
}

func LIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%%%s%%", value),
	}
}

func LLIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%%%s", value),
	}
}

func RLIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%s%%", value),
	}
}

func EQ(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "EQ",
		value:  value,
	}
}

func NE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "NE",
		value:  value,
	}
}

func IN(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "IN",
		value:  value,
	}
}

func GT(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "GT",
		value:  value,
	}
}

func GE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "GE",
		value:  value,
	}
}

func LT(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LT",
		value:  value,
	}
}

func LE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LE",
		value:  value,
	}
}
