DROP TABLE IF EXISTS `user_accounts`;

CREATE TABLE `user_accounts` (
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

DROP TABLE IF EXISTS `user_posts`;

CREATE TABLE `user_posts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"编号"',
  `user_account_id` INT UNSIGNED     NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"作者" update:"n" query:"eq"',
  `title`       VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"标题" update:"y" query:"like" valid:"required(),min(12)"',
  `content`     TEXT             NOT NULL  COMMENT 'index:"n" caption:"内容" update:"y"',
  `image_url`   VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'index:"y" caption:"图片" update:"y"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'index:"y" caption:"状态" update:"y" query:"eq"',
  `create_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'caption:"创建时间"',
  `update_at`   TIMESTAMP(6)     NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'index:"y" caption:"更新时间"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'index:"y" caption:"会员帖子"';