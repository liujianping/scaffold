package models

import (
	"database/sql"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

const (
	DefaultPageSize           = 20
	DefaultMaxIdleConnections = 10
	DefaultMaxOpenConnections = 128
)

func PageSize(size int) int {
	if size <= 0 {
		return DefaultPageSize
	}
	return size
}

func ZeroValue(v reflect.Value) bool {
	valid := true
	switch v.Kind() {
	case reflect.String:
		valid = len(v.String()) != 0
	case reflect.Ptr, reflect.Interface:
		valid = !v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = v.Len() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = v.Float() != 0
	case reflect.Bool:
		valid = v.Bool()
	case reflect.Invalid:
		valid = false // always invalid
	}
	return !valid
}

type Model struct {
	*gorp.DbMap
}

func NewModel(provider, dsn string) (*Model, error) {
	db, err := database(provider, dsn)
	if err != nil {
		return nil, err
	}

	for table, object := range tableObjects {
		db.AddTableWithName(object, table.Name).SetKeys(true, table.PrimaryKey)
	}

	return &Model{db}, nil
}

//! util fns - database
func database(provider, dsn string) (*gorp.DbMap, error) {
	conn, err := sql.Open(provider, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(DefaultMaxIdleConnections)
	conn.SetMaxOpenConns(DefaultMaxOpenConnections)

	return &gorp.DbMap{
		Db: conn,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}, nil
}

//! util fns - regist
type TableObject struct {
	Name       string
	PrimaryKey string
}

var tableObjects map[TableObject]interface{} = make(map[TableObject]interface{})

func registTableObject(tablename string, pk string, object interface{}) {
	tableObjects[TableObject{tablename, pk}] = object
}
