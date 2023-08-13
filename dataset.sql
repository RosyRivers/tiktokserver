CREATE DATABASE IF NOT EXISTS `tiktok` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `tiktok`;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
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

INSERT INTO `users`
VALUES (1, 'zhanglei', 'zhangleidouyin', 1, '0', FALSE);
use tiktok;
ALTER TABLE users ADD INDEX idx_video_author (token);
CREATE TABLE IF NOT EXISTS `video`
(
    `id`              bigint(20) unique not null auto_increment comment '主键ID',
    `token`           varchar(128) not null comment '视频发布者ID',
    `playurl`         varchar(500) not null default 'https://www.w3schools.com/html/movie.mp4' comment '播放地址',
    `coverurl`        varchar(500) not null default 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg' comment '封面',
    `favouritecount`   int(10) not null default 0 comment '点赞总数',
    `commentcount`    int(10) not null default 0 comment '评论数',
    `title`           varchar(128) not null default '' comment '视频标题',
    `uploadtime`    datetime not null default '1970-01-01' comment '视频上传时间',
    PRIMARY KEY (`id`),
    CONSTRAINT video_author FOREIGN KEY (token) references users (token)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 comment = '视频表';