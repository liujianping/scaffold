package controllers

import (
	"[[.project]]/app/models"
	"github.com/revel/revel"
)

func GetTableObject(table string, id int64) interface{} {
	return models.TableObject(db, table, id)
}

func init() {
	revel.TemplateFuncs["table_object"] = GetTableObject
}
