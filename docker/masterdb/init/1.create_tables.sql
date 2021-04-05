CREATE DATABASE IF NOT EXISTS lesson;

CREATE TABLE lesson.`posts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'シーケンスID',
    `title` varchar(255) ,
    `body` TEXT COMMENT 'ブログ本文',
    `created` datetime DEFAULT NULL COMMENT '登録日',
    `modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '変更履歴',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ブログテーブル';

CREATE TABLE lesson.`comments` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'オートインクリメントID',
    `commenter` varchar(255) DEFAULT '' COMMENT 'コメント者',
    `body` varchar(255) DEFAULT '' COMMENT 'コメント内容',
    `post_id` int(10) NOT NULL COMMENT 'postテーブルのID',
    `created` datetime DEFAULT NULL COMMENT '投稿時間',
    `modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `index_time` (`created`,`id`),
    KEY `index_comment` (`commenter`,`body`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コメントテーブル';