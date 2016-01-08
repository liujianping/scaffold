scaffold
===

scaffold, generate revel project by database schema

脚手架工具, 通过定义数据表,一键生成Revel管理平台项目。

### 特点

**通过数据表定义, 一键生成管理平台** 

[索引页]

![index](http://7xjh31.com1.z0.glb.clouddn.com/scaffold_index.png)

[列表页]

![list](http://7xjh31.com1.z0.glb.clouddn.com/user_account_index.png)

[新增页]

![add](http://7xjh31.com1.z0.glb.clouddn.com/user_account_add.png)

[更新页]

![update](http://7xjh31.com1.z0.glb.clouddn.com/user_account_update.png)

[详情页]

![detail](http://7xjh31.com1.z0.glb.clouddn.com/user_account_detail.png)

### 安装

本工具会执行 goimports 工具格式化生成的文件, 请提前安装好该工具.

````
$: go get github.com/liujianping/scaffold

````

### 快速开始

#### 创建 Revel 项目

````
$: revel new [project/path]

````
编辑 项目配置 conf/app.conf 文件, 增加相应数据库配置。

````

db.driver = mysql
db.host = 127.0.0.1
db.port = 3306
db.username = [username]
db.password = [password]
db.database = [database]

````

#### 定义数据结构

定义数据结构说明如下:

````
CREATE TABLE `accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"名称" update:"y" query:"like" valid:"required(),min(6),max(16)"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"邮箱" update:"y" query:"like" valid:"required(),email()"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"性别" update:"y"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"描述" update:"y"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"密码" update:"y" valid:"required()"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"头像" update:"y"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"状态" update:"y" query:"eq"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"创建时间"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"更新时间"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"会员账号"';

````

通过在数据表定义中的注释说明中, 增加脚手架生成规则。具体规则如下:

** 表规则 **

````
-   index
    index:"y" 表示该表需要在列表页(/)中建立索引, 否则不建立索引

-   caption
    caption:"xxxx" 表示该表的可读性名称

````
** 列规则 **

````
-   index
    index:"y" 表示该列需要在列表页(/accounts.index)中展示, 否则不展示

-   caption
    caption:"xxxx" 表示该列的可读性名称

-   query
    query:"like" 表示在列表页(/accounts.index)查询条件面板中增加相应查询字段。
                 具体查询运算符包括: like,llike,rlike,eq,gt,ge,lt,le

-   valid
    valid:"required(),min(6),max(16),email(),length(8)"
                表示在进行该列的增加、更新操作时, 需要进行的验证规则

-   update
    update:"y"  表示该列在更新页面需要进行更新操作，否则会使用hidden控件进行隐藏操作。

````

#### 生成代码

````
//! 初始化 
$: scaffold -f revel init [project/path]  

//! 生成模块代码 [mvc]
$: scaffold -f revel module  [project/path]  [table_name1] [table_name2] ...

//! 生成模型代码 [m]
$: scaffold -f revel model  [project/path]  [table_name1] [table_name2] ...

//! 生成视图代码 [v]
$: scaffold -f revel view  [project/path]  [table_name1] [table_name2] ...

//! 生成控制器代码 [c]
$: scaffold -f revel controller  [project/path]  [table_name1] [table_name2] ...

//! 建立索引路由
$: scaffold -f revel index [project/path] 

````

#### 运行 Revel 项目

````
$: revel run [project/path]

````

#### TODO

* 通用 HTML 控件生成
* 其它框架支持(beego等)
