### 模板定义

scaffold 模板支持所有 golang 中 text/template 包提供的模版功能。同时还增加了以下特性:

-   模板分割符
    
    使用"[["与"]]".

-   模板函数:

    - 驼峰函数(camel) 
````go
    转化例子: users => Users
````
    - 复数函数(plural)
````go
    user => users
````
    - 单数函数(singular)
````go
    users => user
````
    - 函数(lint)

    - 函数(quote)
````go
    users => "users"
````
    - 函数(convert)

    主要用于数据库字段类型的转化
````go
    covert("mysql", "tinyint", "") => int64
````
    - 函数(module)
````go
    user_accounts => user.account
````
    - 函数(class)
````go
    user_accounts => UserAccount
````
    - 加函数(add)
    - 减函数(sub)
    - 乘函数(multiply)
    - 除函数(divide)

-   模板数据   

    - 表对象接口

    ````
    table.Name() => 返回表格名
    table.Columns() => 返回表格所有列表对象
    table.Column(name string) => 返回表格指定列表对象
    table.Comment() => 返回表格注释
    table.Tag(tag string) => 返回表格注释中指定tag
    ````

    - 列表对象接口
    ````
    column.Name() => 返回列表名
    column.Type() => 返回列表类型
    column.Comment() => 返回列表注释
    column.Tag(tag string) => 返回表格注释中指定tag
    ````

````go
  data := map[string]interface{}{
    "project": "path/to/project", //!用户调用命令输入的生成项目路径,
    "tables":  tables,            //!用户定义的所有数据库表结构定义数组,
    "table": table,               //!当前模板使用的表结构定义,
    "index": index,               //!当前模板使用的表结构位置,
  }
````

例子: 根据数据库表定义生成model定义代码，模板定义如下:

生成单个文件模板：

模板名: models.go

模板内容:

````go
package model

import (
  "time"
)

[[range .tables]]
//! [[.Tag "caption"]]
type [[.Name | singular | camel]] struct {
  [[range .Columns]]
  [[.Field | camel | lint]] [[convert "mysql" .Type (.Tag "gotype")]] `db:"[[.Field]]"    json:"[[.Field]]"`[[end]]
}

func (obj [[.Name | singular | camel]]) TableName() string {
  return "[[.Name]]"
}
[[end]]

````
    
生成多个文件模板：

模板名: model.[[.table.Name | module]].go

模板内容:

````go
package model
[[set . "t_class" (.table.Name | singular | camel)]]
import (
  "time"
)

//! [[.table.Tag "caption"]]
type [[.t_class]] struct {
  [[range .table.Columns]]
  [[.Field | camel | lint]] [[convert "mysql" .Type (.Tag "gotype")]] `db:"[[.Field]]"    json:"[[.Field]]"`[[end]]
}

func (obj [[.t_class]]) TableName() string {
  return "[[.table.Name]]"
}

````