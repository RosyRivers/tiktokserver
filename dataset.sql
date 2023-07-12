CREATE DATABASE IF NOT EXISTS `tiktok` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `tiktok`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`              varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `token`             varchar(128)        NOT NULL DEFAULT '' COMMENT '昵称+密码，验证用户',
    `followcount`       int(10)             NOT NULL DEFAULT 1 COMMENT '关注数量',
    `followercount`     int(10)             NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    `isfollow`          boolean             NOT NULL DEFAULT FALSE COMMENT '是否被关注',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'zhanglei', 'zhangleidouyin', 1, '0', FALSE);