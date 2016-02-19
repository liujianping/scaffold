scaffold
===

[中文](README.cn.md)

scaffold, generate code by database schema definition through template. 

now scaffold tool supports the following templates:

- revel web project template

- data model template

### Installation

````
$: go get github.com/liujianping/scaffold

````
### Quick Start

````
$: scaffold -f=template/model -d=database -u=use -p=password github.com/liujianping/example/models

````

#### 1. Generate Data Model Definition Code

example command:

````
$: scaffold model generate -d=database -u=user -p=password github.com/liujianping/example/models

````

#### 2. Generate Revel Web Application

##### 2.1 create revel project by revel command

````
$: revel new [project/path]

```` 

##### 2.2 preparing database schema definitions

setup the database schema at mysql database.

example [schema.en.sql](demo/conf/schema.en.sql):

Rules for database schema definition,

scaffold use database field commet as the scaffold generating code rules, includes:

** table rule **

````
-   index
    index:"y" means home page index the table

-   caption
    caption:"xxxx" means home page indexing name

````
** column rule **

````
-   caption
    caption:"xxxx" means column generate caption name

-   index
    index:"y" means column showing at the list page columns

-   query
    query:"like" means table list page(/table.name.index) search fields
                 symbols includes: like,llike,rlike,eq,gt,ge,lt,le

-   finder
    finder:"like" means finder popup page (/table.name.finder.index) search fields
                 symbols includes: like,llike,rlike,eq,gt,ge,lt,le

-   valid         means create/update the table object will do the validators
    valid:"required(),min(6),max(16),email(),length(8)"
                

-   update        means at update page, the table object which fields can be updated
    update:"y"    "y" means yes. "n" means no.

-   widget
    widget:"text"  means generate the field with the widget definition

                - text        text input widget
                - number      number input widget
                - password    password input widget
                - email       email input widget
                - datetime    datetime input widget
                - file        file upload widget
                - finder      object finder widget
                - textarea    textarea widget
                - selection   selection widget use the system options definitions

    you can custom a new widget with a widget/xxxx.html  definition.

-   relation
    relation:"name" widget:"finder" relate table name
    relation:"name" widget:"selection" relate system options code
    
````
example:

````
CREATE TABLE `user_accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"no"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"name" update:"w" finder:"like" query:"like" widget:"text" valid:"required(),min(6),max(16)"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"mailbox" update:"r" finder:"like" query:"like" widget:"email" valid:"required(),email()"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"sex" update:"w" widget:"selection" relation:"user_accounts_sex"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"description" update:"w" widget:"textarea"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"password" update:"r" widget:"password" valid:"required()"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"head_url" update:"w" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"status" update:"w" finder:"eq" query:"eq" widget:"selection" relation:"user_accounts_status"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"createtime" update:"r" widget:"datetime"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"updatetime" update:"w" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"accounts"';
````


##### 2.3 modify revel project conf with databse configs

for example, revel project *conf/app.conf*

````
# datetime format 
format.date             = 2006/01/02
format.datetime         = 2006/01/02 15:04:05

# database 
db.driver               = mysql
db.host                 = 127.0.0.1
db.port                 = 3306
db.username             = 
db.password             = 
db.database             = 

# upload support qiniu cloud storage
qiniu.enable            = false
qiniu.access            = 
qiniu.secret            = 
qiniu.bucket            = 
qiniu.domain            = 

````

##### 2.4 generate revel project code

use the following commands:

````
$: scaffold -f revel init [project/path]  

$: scaffold -f revel module  [project/path]  "*"

$: scaffold -f revel index [project/path] 

````

##### 2.5 running revel project 

````
$: revel run [project/path]

````

[home]

![home](http://7xjh31.com1.z0.glb.clouddn.com/home.png)

[widget]

![list](http://7xjh31.com1.z0.glb.clouddn.com/widget.png)

[widget finder]

![finder](http://7xjh31.com1.z0.glb.clouddn.com/find.png)

[add]

![update](http://7xjh31.com1.z0.glb.clouddn.com/add.png)

[list]

![index](http://7xjh31.com1.z0.glb.clouddn.com/index.png)


### TODO

-   more database driver support
-   more scaffold template projects

### Thanks 

[jaywcjlove](https://github.com/jaywcjlove) for providing scaffold revel project datetime.js
