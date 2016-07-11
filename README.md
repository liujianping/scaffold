scaffold
===

scaffold, 基于数据库表定义的代码生成工具, 已支持代码生成模板包括：

* model  对象生成
* portal 管理平台生成

用户还可以根据自己自定义的模板进行代码生成，前提是了解scaffold生成代码的规则。

以下是 portal 管理平台生成的用户账号表的定义，所有scaffold所需要的信息除了表格本身定义外，大部分在comment部分进行定义。

````sql
DROP TABLE IF EXISTS `user_accounts`;

CREATE TABLE `user_accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"名称" column:"y" update:"y" query:"like" widget:"text" valid:"required(),min(6),max(16)"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'caption:"邮箱" column:"y" query:"like" widget:"email" valid:"required(),email()"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"性别" column:"y" update:"y" widget:"selection" relation:"user_accounts_sex"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"描述" update:"y" widget:"textarea"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"密码" update:"y" widget:"password" valid:"required()"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"头像" update:"y" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"状态" column:"y" update:"y" query:"eq" widget:"selection" relation:"user_accounts_status"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"创建时间" widget:"datetime"',
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'caption:"更新时间" column:"y" widget:"datetime"',
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'caption:"删除时间" gotype:"*time.Time" ignore:"y" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"会员账号" index:"y" import:"y" export:"y"';

````
从以上定义可以看出，scaffold 从数据表定义中获取信息的方式和golang语言中json、xml等对象的方式一样，通过类似于字段的tag标签的方式设置对象在该字段上提供的属性信息。

````go
type SomeJsonthing struct{
  FieldX int64 `json:"x"`
  FieldY string `json:"y"`
}
````
但是，scaffold中具体tag的意义只与相应代码生成模板有关。例如，portal模板与model模板中的tag的定义完全是独立的，所以如果是自定义模板，具体tag也完全由用户自己定义。

## 快速开始



#### Thanks 

[jaywcjlove](https://github.com/jaywcjlove) 提供的datetime控件js脚本