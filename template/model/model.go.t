package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

const (
	DefaultPageSize           = 20
	DefaultMaxIdleConnections = 10
	DefaultMaxOpenConnections = 128
)

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
