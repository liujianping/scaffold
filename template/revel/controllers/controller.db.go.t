package controllers

import (
	"database/sql"
	"fmt"

	"[[.project]]/app/models"
	"github.com/revel/revel"
	gorp "gopkg.in/gorp.v1"
)

var (
	db *models.Model
)

type DBController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *DBController) Index() revel.Result {
	return c.RenderTemplate("index.html")
}

func (c *DBController) Begin() revel.Result {
	txn, err := db.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *DBController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *DBController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	host := revel.Config.StringDefault("db.host", "127.0.0.1")
	port := revel.Config.IntDefault("db.port", 3306)
	user := revel.Config.StringDefault("db.username", "")
	pass := revel.Config.StringDefault("db.password", "")
	name := revel.Config.StringDefault("db.database", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&strict=true&sql_notes=false",
		user, pass, host, port, name)

	m, err := models.NewModel(driver, dsn)
	if err != nil {
		panic(err)
	}

	db = m
	db.TraceOn("[db]", revel.INFO)
}

func init() {
	revel.OnAppStart(InitDB)
}
