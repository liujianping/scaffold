-- CREATE DATABASE IF NOT EXISTS `scaffold`;
-- GRANT ALL PRIVILEGES ON *.* TO `scaffold_user`@'%' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON *.* TO 'scaffold_user'@'localhost' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON `scaffold`.* TO `scaffold_user`@'%';
-- FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON *.* TO `maiya_user`@'%' IDENTIFIED BY 'maiya_pass';
GRANT ALL PRIVILEGES ON `maiya`.* TO `maiya_user`@'localhost';
GRANT ALL PRIVILEGES ON `maiya`.* TO `maiya_user`@'%';

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
