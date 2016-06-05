package models

import (
	"fmt"
	"strings"
)

type Expression struct {
	field  string
	symbol string
	value  interface{}
}

func BinaryExpression(field, symbol string, value interface{}) Expression {
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

func (exp Expression) Statment() (format string, data interface{}) {
	format = fmt.Sprintf("%s %s ?", exp.field, exp.symbol)
	data = exp.value
	return
}

func NONE(field string, value interface{}) Expression {
	return Expression{
		field:  "",
		symbol: "",
		value:  "",
	}
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
		symbol: "=",
		value:  value,
	}
}

func NE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "!=",
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
		symbol: ">",
		value:  value,
	}
}

func GE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: ">=",
		value:  value,
	}
}

func LT(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "<",
		value:  value,
	}
}

func LE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "<=",
		value:  value,
	}
}
