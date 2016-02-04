package models
[[set . "ClassName" (.table.Name | singular | camel)]]
[[set . "ModuleName" (.table.Name | module)]]
var (
    Default[[.ClassName]]       = [[.ClassName]]{}
)

type [[.ClassName]] struct {[[range .table.Columns]]
    [[.Field | camel | lint]]   [[.Type | convert "mysql"]] `db:"[[.Field]]"`[[end]]
}

func (obj [[.ClassName]]) TableName() string {
    return "[[.table.Name]]"
}

func (obj [[.ClassName]]) PrimaryKey() string {
    return "[[.table.PrimaryColumn.Field]]"
}

func init() {
    registTableObject(Default[[.ClassName]].TableName(),
        Default[[.ClassName]].PrimaryKey(),
        Default[[.ClassName]])
}
