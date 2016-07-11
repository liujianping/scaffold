model 数据表定义
===

例子:

````sql
CREATE TABLE `user_accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT,
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT,
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'gotype:"*time.Time"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"会员账号"';

````

主要提供了两个tag功能:

- 	可选:	给表加一个 caption 的comment注释, 这样生成代码时会出现相遇的注释信息。

- 	可选:	给字段加一个 gotype 的comment注释, 这样生成该字段类型时, 就会使用自定的 gotype 做为生成的字段类型。
