package model
[[set . "t_class" (.table.Name | singular | camel)]]
import (
	"time"
)

//! [[.table.Tag "caption"]]
type [[.t_class]] struct {
	[[range .table.Columns]]
	[[.Field | camel | lint]]	[[convert "mysql" .Type (.Tag "gotype")]]	`db:"[[.Field]]"`[[end]]
}

func (obj [[.t_class]]) TableName() string {
	return "[[.table.Name]]"
}

