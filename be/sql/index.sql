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
        `phone_number` varchar(11) DEFAULT NULL COMMENT '手机号',
        `password` varchar(255) NOT NULL COMMENT '密码',
        `avatar` varchar(255) DEFAULT NULL COMMENT '头像token',
        `cover` varchar(50) DEFAULT NULL COMMENT '封面图片token',
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
        `created_at` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
        `updated_at` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
        `deleted_at` int (10) unsigned NOT NULL DEFAULT '0' COMMENT '禁用时间',
        PRIMARY KEY (`user_id`),
        KEY `user_name` (`username`),
        KEY `email` (`email`),
        KEY `create_at` (`create_time`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';