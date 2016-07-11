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

-- 系统账户
DROP TABLE IF EXISTS `system_accounts`;
CREATE TABLE `system_accounts` (
  `id`             INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`           VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"名称" column:"y" update:"y" query:"like" widget:"text" valid:"required()"',
  `mailbox`        VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"邮箱" update:"y" widget:"email" valid:"required(),email()"',
  `password`       VARCHAR(64)      NOT NULL  DEFAULT '' COMMENT 'caption:"密码" update:"y" cipher:"md5" widget:"password" valid:"required()"',
  `role`           TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"角色" column:"y" update:"y" query:"eq" widget:"selection" relation:"system_account_role"',
  `status`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"状态" column:"y" update:"y" finder:"eq" query:"eq" widget:"selection" relation:"system_account_status"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"创建时间" widget:"datetime"',
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'caption:"更新时间" column:"y" widget:"datetime"',
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'caption:"删除时间" gotype:"*time.Time" ignore:"y" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"系统账户" index:"y"';
INSERT INTO `system_accounts`(`name`,`mailbox`, `password`, `role`, `status`) VALUES('super', 'super@bomotong.com', MD5('bomotong'), 1, 1);

-- 系统令牌
DROP TABLE IF EXISTS `system_tokens`;
CREATE TABLE `system_tokens` (
  `id`                INT UNSIGNED     NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `system_account_id` INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'caption:"系统账户" column:"y" widget:"finder" relation:"system_accounts" field:"name" valid:"required()"',
  `secret`            VARCHAR(64)      NOT NULL DEFAULT '' COMMENT 'caption:"令牌秘钥" cipher:"uuid" widget:"text"',
  `remote`            VARCHAR(32)      NOT NULL DEFAULT '' COMMENT 'caption:"远程地址" column:"y" query:"like" widget:"text"',
  `strict`            TINYINT(1)       NOT NULL DEFAULT 0 COMMENT 'caption:"加强校验" widget:"selection" relation:"bool"',
  `times`             INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'caption:"过期次数" column:"y" update:"y" widget:"number"',
  `hits`              INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'caption:"访问次数" column:"y" widget:"number"',
  `expired_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"过期时间" column:"y" widget:"datetime"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"创建时间" widget:"datetime"',
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'caption:"更新时间" column:"y" widget:"datetime"',
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'caption:"删除时间" gotype:"*time.Time" ignore:"y" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"系统令牌" index:"y"';
CREATE INDEX `idx_tokens_secret` ON `system_tokens` (`secret`);

-- system system_system_options
DROP TABLE IF EXISTS `system_system_options`;

CREATE TABLE `system_system_options` (
  `id`                   INT UNSIGNED        NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `name`                 VARCHAR(32) DEFAULT "" COMMENT 'caption:"名称" column:"y" update:"y" query:"like" widget:"text" valid:"required()"',
  `code`                 VARCHAR(32) DEFAULT "" COMMENT 'caption:"代码" column:"y" update:"y" query:"eq" widget:"text" valid:"required()"',
  `relate_code`          VARCHAR(32) DEFAULT "" COMMENT 'caption:"关联代码" column:"y" update:"y" query:"eq" widget:"text"',
  `option_code`          VARCHAR(32) DEFAULT "" COMMENT 'caption:"选项代码" column:"y" update:"y" widget:"text" valid:"required()"',
  `option_name`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"选项名称" column:"y" update:"y" widget:"text" valid:"required()"',
  `option_value`         INT NOT NULL DEFAULT 0 COMMENT 'caption:"选项值" column:"y" update:"y" query:"eq" widget:"number" valid:"required()"',
  `description`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"描述" update:"y" widget:"text"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"系统选项" index:"y" index:"y"';



INSERT INTO `system_options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "管理账户角色" , "portal_account_role", "超级管理员", "super", 1,  "超级管理员"),
( "管理账户角色" , "portal_account_role", "管理员",  "administrator", 2,  "管理员"),
( "管理账户角色" , "portal_account_role", "运营操作员", "operator", 3,  "运营操作员");

INSERT INTO `system_options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "管理账户状态" , "portal_account_status", "已启用", "normal", 1,  "已启用"),
( "管理账户状态" , "portal_account_status", "已禁用", "fibidden", 4,  "已禁用");

INSERT INTO `system_options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "性别" , "user_accounts_sex", "男", "male", 1,  "男"),
( "性别" , "user_accounts_sex", "女", "female", 2,  "女");

INSERT INTO `system_options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "账号状态" , "user_accounts_status", "待激活", "unactived", 1,  "账号待激活"),
( "账号状态" , "user_accounts_status", "正常", "normal", 2,  "账号正常"),
( "账号状态" , "user_accounts_status", "禁用", "forbidden", 3,  "账号禁用");

INSERT INTO `system_options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "帖子状态" , "user_posts_status", "待发布", "draft", 1,  "帖子待发布"),
( "帖子状态" , "user_posts_status", "已发布", "public", 2,  "帖子已发布"),
( "帖子状态" , "user_posts_status", "已作废", "trash", 3,  "帖子已作废");
