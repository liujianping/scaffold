package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DefaultPageSize           = 20
	DefaultMaxIdleConnections = 10
	DefaultMaxOpenConnections = 128
)

func DB(driver, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(DefaultMaxOpenConnections)
	db.DB().SetMaxIdleConns(DefaultMaxIdleConnections)
	return db, nil
}

func Columns(obj interface{}) []string {
	var columns []string

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("db") != "" && field.Tag.Get("db") != "-" {
			columns = append(columns, field.Tag.Get("db"))
		}
	}
	return columns
}

func TableObject(db *gorm.DB, table string, id int64) interface{} {
	switch table {
	[[range .tables]]
	case "[[.Name]]":
		var obj [[.Name | class]]
		if err := db.Find(&obj, id).Error; err != nil {
			return nil
		}
		return obj
	[[end]]
	}
	return nil
}

func DefaultTableObject(table string) interface{} {
	switch table {
	[[range .tables]]
	case "[[.Name]]":
		return Default[[.Name | class]]
	[[end]]
	}
	return nil
}
