-- CREATE DATABASE IF NOT EXISTS `scaffold`;
-- GRANT ALL PRIVILEGES ON *.* TO `scaffold_user`@'%' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON *.* TO 'scaffold_user'@'localhost' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON `scaffold`.* TO `scaffold_user`@'%';
-- FLUSH PRIVILEGES;

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

DROP TABLE IF EXISTS `user_posts`;

CREATE TABLE `user_posts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `user_account_id` INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'caption:"作者" column:"y" update:"y" query:"eq" widget:"finder" relation:"user_accounts" field:"id"',
  `title`       VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"标题" column:"y" update:"y" query:"like" widget:"text" valid:"required(),min(12)"',
  `content`     TEXT             NOT NULL  COMMENT 'caption:"内容" update:"y" widget:"textarea"',
  `image_url`   VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"图片" update:"y" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"状态" column:"y" update:"y" query:"eq" widget:"selection" relation:"user_posts_status"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'column:"y" caption:"创建时间" widget:"datetime"',
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'caption:"更新时间" column:"y" widget:"datetime"',
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'caption:"删除时间" gotype:"*time.Time" ignore:"y" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"会员帖子" index:"y" import:"y" export:"y"';

-- system options
DROP TABLE IF EXISTS `options`;

CREATE TABLE `options` (
  `id`                   INT UNSIGNED        NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`                 VARCHAR(32) DEFAULT "" COMMENT 'caption:"名称" column:"y" update:"y" query:"like" widget:"text" valid:"required()"',
  `code`                 VARCHAR(32) DEFAULT "" COMMENT 'caption:"代码" column:"y" update:"y" query:"eq" widget:"text" valid:"required()"',
  `option_name`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"选项名称" column:"y" update:"y" widget:"text" valid:"required()"',
  `option_code`          VARCHAR(32) DEFAULT "" COMMENT 'caption:"选项代码" column:"y" update:"y" widget:"text" valid:"required()"',
  `option_value`         INT NOT NULL DEFAULT 0 COMMENT 'caption:"选项值" column:"y" update:"y" query:"eq" widget:"number" valid:"required()"',
  `description`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"描述" update:"y" widget:"text"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"系统状态" index:"y" import:"y" export:"y"';

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
