-- 删除数据库中的所有表数据
SET FOREIGN_KEY_CHECKS = 0;
-- 获取所有表名并拼接成 DROP TABLE 语句，这里假设表名之间不存在空格等特殊情况干扰
SELECT CONCAT('DROP TABLE ', table_name, ';') INTO @drop_tables_sql
FROM information_schema.tables
WHERE table_schema = DATABASE() AND table_type = 'BASE TABLE';

-- 执行拼接好的语句来删除表
PREPARE stmt FROM @drop_tables_sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE
    `user` (
        `user_id` int (11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
        `username` varchar(20) NOT NULL COMMENT '用户名',
        `email` varchar(320) NOT NULL COMMENT '邮箱',
        `avatar` varchar(255) DEFAULT NULL COMMENT '头像token',
        `cover` varchar(50) DEFAULT NULL COMMENT '封面图片token',
        `password` varchar(255) NOT NULL COMMENT '密码',
        `create_ip` varchar(80) DEFAULT NULL COMMENT '注册IP',
        `create_location` varchar(100) DEFAULT NULL COMMENT '注册地址',
        `last_login_time` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
        `last_login_ip` varchar(80) DEFAULT NULL COMMENT '最后登陆IP',
        `last_login_location` varchar(100) DEFAULT NULL COMMENT '最后登录地址',
        `follower_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '关注我的人数',
        `followee_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我关注的人数',
        `following_article_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我关注的文章数',
        `following_question_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我关注的问题数',
        `following_topic_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我关注的话题数',
        `article_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我发表的文章数量',
        `question_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我发表的问题数量',
        `answer_count` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '我发表的回答数量',
        `notification_unread` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '未读通知数',
        `inbox_unread` int (11) unsigned NOT NULL DEFAULT '0' COMMENT '未读私信数',
        `headline` varchar(40) DEFAULT NULL COMMENT '一句话介绍',
        `bio` varchar(160) DEFAULT NULL COMMENT '个人简介',
        `blog` varchar(255) DEFAULT NULL COMMENT '个人主页',
        `company` varchar(255) DEFAULT NULL COMMENT '公司名称',
        `location` varchar(255) DEFAULT NULL COMMENT '地址',
        `create_time` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
        `update_time` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
        `disable_time` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '禁用时间',
        PRIMARY KEY (`user_id`),
        KEY `user_name` (`username`),
        KEY `email` (`email`),
        KEY `follower_count` (`follower_count`),
        KEY `create_time` (`create_time`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';