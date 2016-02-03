-- CREATE DATABASE IF NOT EXISTS `scaffold`;
-- GRANT ALL PRIVILEGES ON *.* TO `scaffold_user`@'%' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON *.* TO 'scaffold_user'@'localhost' IDENTIFIED BY 'scaffold_pass';
-- GRANT ALL PRIVILEGES ON `scaffold`.* TO `scaffold_user`@'%';
-- FLUSH PRIVILEGES;

DROP TABLE IF EXISTS `user_accounts`;

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

DROP TABLE IF EXISTS `user_posts`;

CREATE TABLE `user_posts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"no"',
  `user_account_id` INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"author" update:"w" query:"eq" widget:"finder" relation:"user_accounts"',
  `title`       VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"title" update:"w" finder:"like" query:"like" widget:"text" valid:"required(),min(12)"',
  `content`     TEXT             NOT NULL  COMMENT 'index:"n" caption:"content" update:"w" widget:"textarea"',
  `image_url`   VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'index:"n" caption:"image_url" update:"w" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"status" update:"w" finder:"eq" query:"eq" widget:"selection" relation:"user_posts_status"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"createtime" widget:"datetime"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"updatetime" update:"w" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"posts"';

-- system options
DROP TABLE IF EXISTS `options`;

CREATE TABLE `options` (
  `id`                   INT UNSIGNED        NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"no"',
  `name`                 VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"name" update:"w" query:"like" widget:"text" valid:"required()"',
  `code`                 VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"code" update:"w" query:"eq" widget:"text" valid:"required()"',
  `option_name`          VARCHAR(256) DEFAULT "" COMMENT 'index:"y" caption:"option_name" update:"w" widget:"text" valid:"required()"',
  `option_code`          VARCHAR(32) DEFAULT "" COMMENT 'index:"y" caption:"option_code" update:"w" widget:"text" valid:"required()"',
  `option_value`         INT NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"option_value" update:"w" query:"eq" widget:"number" valid:"required()"',
  `description`          VARCHAR(256) DEFAULT "" COMMENT 'caption:"description" update:"w" widget:"text"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"options"';

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "sex" , "user_accounts_sex", "male", "male", 1,  "male"),
( "sex" , "user_accounts_sex", "female", "female", 2,  "female");

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "account_status" , "user_accounts_status", "unactived", "unactived", 1,  "account unactived"),
( "account_status" , "user_accounts_status", "normal", "normal", 2,  "account normal"),
( "account_status" , "user_accounts_status", "forbidden", "forbidden", 3,  "account forbidden");

INSERT INTO `options`(
  `name`, `code`, `option_name`, `option_code`, `option_value`, `description`     
) VALUES 
( "post_status" , "user_posts_status", "draft", "draft", 1,  "draft post"),
( "post_status" , "user_posts_status", "public", "public", 2,  "public post"),
( "post_status" , "user_posts_status", "trash", "trash", 3,  "trash post");
