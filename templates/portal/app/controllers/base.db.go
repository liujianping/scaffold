package controllers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"

	"[[.project]]/app/models"
)

var db *gorm.DB

func InitDatabase() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	host := revel.Config.StringDefault("db.host", "127.0.0.1")
	port := revel.Config.IntDefault("db.port", 3306)
	user := revel.Config.StringDefault("db.username", "")
	pass := revel.Config.StringDefault("db.password", "")
	name := revel.Config.StringDefault("db.database", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&strict=true&sql_notes=false",
		user, pass, host, port, name)

	var err error
	if db, err = models.DB(driver, dsn); err != nil {
		panic(err)
	}

	db.LogMode(revel.Config.BoolDefault("db.debug", false))
}

func init() {
	revel.OnAppStart(InitDatabase)
}
