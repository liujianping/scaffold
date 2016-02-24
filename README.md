scaffold
===

[english](README.en.md)

Go Scaffold,  Generate code from database schema by template

脚手架工具, 通过数据表和自定义模板生成代码。

### 特点

**通过数据表定义, 一键生成管理平台** 

**支持自定义控件模板**

[索引页]

![home](http://7xjh31.com1.z0.glb.clouddn.com/home.png)

[控件]

![list](http://7xjh31.com1.z0.glb.clouddn.com/widget.png)

[查询控件]

![add](http://7xjh31.com1.z0.glb.clouddn.com/find.png)

[新增页]

![update](http://7xjh31.com1.z0.glb.clouddn.com/add.png)

[列表页]

![index](http://7xjh31.com1.z0.glb.clouddn.com/index.png)

以上图片效果均为实际生成效果, 未编写一行代码:)

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
请根据实际环境配置, 增加并修改项目配置[conf/app.conf]:

````
# format (datetime控件默认日期格式,如需修改请相应修改 widget/datetime.html 日期格式。需保持一致。)
format.date             = 2006/01/02
format.datetime         = 2006/01/02 15:04:05

# database (数据库配置)
db.driver               = mysql
db.host                 = 127.0.0.1
db.port                 = 3306
db.username             = 
db.password             = 
db.database             = 

# upload support qiniu (上传功能支持七牛云存储配置项)
qiniu.enable            = false
qiniu.access            = 
qiniu.secret            = 
qiniu.bucket            = 
qiniu.domain            = 

````

#### 定义数据结构

Demo项目中定义数据结构说明如下:

````
DROP TABLE IF EXISTS `user_accounts`;

CREATE TABLE `user_accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"名称" update:"w" finder:"like" query:"like" widget:"text" valid:"required(),min(6),max(16)"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"邮箱" update:"r" finder:"like" query:"like" widget:"email" valid:"required(),email()"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"性别" update:"w" widget:"selection" relation:"user_accounts_sex"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"描述" update:"w" widget:"textarea"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"密码" update:"r" widget:"password" valid:"required()"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"头像" update:"w" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"状态" update:"w" finder:"eq" query:"eq" widget:"selection" relation:"user_accounts_status"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"创建时间" update:"r" widget:"datetime"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"更新时间" update:"w" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"会员账号"';

DROP TABLE IF EXISTS `user_posts`;

CREATE TABLE `user_posts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `user_account_id` INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"作者" update:"w" query:"eq" widget:"finder" relation:"user_accounts"',
  `title`       VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"标题" update:"w" finder:"like" query:"like" widget:"text" valid:"required(),min(12)"',
  `content`     TEXT             NOT NULL  COMMENT 'index:"n" caption:"内容" update:"w" widget:"textarea"',
  `image_url`   VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'index:"n" caption:"图片" update:"w" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"状态" update:"w" finder:"eq" query:"eq" widget:"selection" relation:"user_posts_status"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"创建时间" widget:"datetime"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"更新时间" update:"w" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"会员帖子"';

-- system options
DROP TABLE IF EXISTS `options`;

CREATE TABLE `options` (
  `id`                   INT UNSIGNED        NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`                 VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"名称" update:"w" query:"like" widget:"text" valid:"required()"',
  `code`                 VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"代码" update:"w" query:"eq" widget:"text" valid:"required()"',
  `option_name`          VARCHAR(256) DEFAULT "" COMMENT 'index:"y" caption:"选项名称" update:"w" widget:"text" valid:"required()"',
  `option_code`          VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"选项代码" update:"w" widget:"text" valid:"required()"',
  `option_value`         INT NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"选项值" update:"w" query:"eq" widget:"number" valid:"required()"',
  `description`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"描述" update:"w" widget:"text"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"系统状态"';

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "性别" , "user_accounts_sex", "男", "male", 1,  "男"),
( "性别" , "user_accounts_sex", "女", "female", 2,  "女");

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "账号状态" , "user_accounts_status", "待激活", "unactived", 1,  "账号待激活"),
( "账号状态" , "user_accounts_status", "正常", "normal", 2,  "账号正常"),
( "账号状态" , "user_accounts_status", "禁用", "forbidden", 3,  "账号禁用");

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "帖子状态" , "user_posts_status", "待发布", "draft", 1,  "帖子待发布"),
( "帖子状态" , "user_posts_status", "已发布", "public", 2,  "帖子已发布"),
( "帖子状态" , "user_posts_status", "已作废", "trash", 3,  "帖子已作废");

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
-   caption
    caption:"xxxx" 表示该列的可读性名称

-   index
    index:"y" 表示该列需要在列表页(/module.name.index)中展示, 否则不展示

-   query
    query:"like" 表示在列表页(/module.name.index)查询条件面板中增加相应查询字段。
                 具体查询运算符包括: like,llike,rlike,eq,gt,ge,lt,le

-   finder
    finder:"like" 表示在查找控件列表页(/module.name.finder.index)查询条件面板中增加相应查询字段。
                 具体查询运算符包括: like,llike,rlike,eq,gt,ge,lt,le

-   valid
    valid:"required(),min(6),max(16),email(),length(8)"
                表示在进行该列的增加、更新操作时, 需要进行的验证规则

-   update
    update:"y"  表示该列在更新页面需要进行更新操作，否则会使用hidden控件进行隐藏操作。

-   widget
    widget:"text"  表示该列在新增或更新页面需要进行生成的HTML控件模板, 具体控件模板可以进行自定义。
                目前支持的控件包括: 

                - text        文本输入控件
                - number      数字输入控件
                - password    密码输入控件   
                - email       邮箱输入控件
                - datetime    时间输入控件 
                - file        文件上传输入控件
                - finder      管理表记录查找控件
                - textarea    编辑区域输入控件
                - finder      查询关联表控件
                - selection   查询关联selection控件 

    以上所有控件, 待项目生成后, 均可以在 views/widget/xxx.html 进行自定义设置。

-   relation
    relation:"name" 当widget:"finder"时,表示该字段关联的表名
    relation:"name" 当widget:"selection"时,表示该字段关联的options表中的code
    
````

#### 生成代码

````
//! 初始化
$: scaffold -f revel init [project/path]  

//! 生成所有模块代码 [mvc]
$: scaffold -f revel module  [project/path]  "*"

//! 建立索引路由
$: scaffold -f revel index [project/path] 

````

#### 运行 Revel 项目

````
$: revel run [project/path]

````

#### 控件模板

可以在 [project/path]/app/view/widget/ 目录中自定义自己的控件模板。

#### TODO

* 其它框架支持(beego等)

#### Thanks 

[jaywcjlove](https://github.com/jaywcjlove) 提供的datetime控件js脚本