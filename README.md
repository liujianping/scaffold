scaffold
===

一款基于数据库定义的代码生成器。

### 它是如何工作的？

正如我们所知, go 中进行 json 字符串的编码/解码过程中, 可以通过对象定义时字段的tag定义, 对字段进行补充说明。如下例:

````go
type JsonSomething struct{
  AField  int64     `json:"x"`
  BField  string    `json:"y"`
}
````
同样的方法, scaffold 通过数据库定义中的字段(或表)的 COMMENT 定义来对相应字段(或表)进行补充说明, 在根据模板进行代码生成。如:

````sql
CREATE TABLE `users` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"名称"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'caption:"邮箱"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"性别"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"描述"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"密码"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"头像"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"状态"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"创建时间"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"会员"';

````
如定义表结构后, scaffold 就可以通过 模板函数读取到 comment 中的 caption 字段, 并根据模板生成代码了。

### 快速开始

##### 安装

````shell
$: go get github.com/liujianping/scaffold

````

##### 生成model代码

[表定义详解](/doc/model.md)

````shell
$: scaffold -i=.go -t=model generate -d="database" -u="root" -p="pass" github.com/yourname/model

````

##### 生成管理平台

[表定义例子](/doc/portal.sql)

[表定义详解](/doc/portal.md)

````shell
# 生成项目
$: scaffold -i=.go -i=.html -i=routes -t=portal generate -d="database" -u="root" -p="pass" github.com/yourname/portal

# 修改数据库配置 github.com/yourname/portal/conf/app.conf

# 运行项目
$: revel run github.com/yourname/portal
````

##### 自定义模板

[模版定义详解](/doc/template.md)

#### Thanks 

[jaywcjlove](https://github.com/jaywcjlove) 提供的datetime控件js脚本