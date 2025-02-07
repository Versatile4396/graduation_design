-- 创建数据库，指定字符集和排序规则
CREATE DATABASE IF NOT EXISTS appDB CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

USE appDB;

-- 删除用户表（如果存在）
DROP TABLE IF EXISTS users;

-- 创建用户表
CREATE TABLE
    IF NOT EXISTS users (
        user_index INT (11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
        user_id VARCHAR(20) NOT NULL COMMENT '用户ID',
        user_name VARCHAR(20) NOT NULL COMMENT '用户名',
        email VARCHAR(320) COMMENT '邮箱',
        phone_number VARCHAR(11) DEFAULT NULL COMMENT '手机号',
        password VARCHAR(255) NOT NULL COMMENT '密码',
        confirm_password VARCHAR(255) NOT NULL COMMENT '重复密码',
        gender TINYINT (2) NOT NULL DEFAULT 0 COMMENT '性别 0:未知 1:男 2:女',
        is_email_verified TINYINT (1) NOT NULL DEFAULT 0 COMMENT '邮箱是否已验证',
        is_phone_verified TINYINT (1) NOT NULL DEFAULT 0 COMMENT '手机号是否已验证',
        third_party_accounts JSON DEFAULT NULL COMMENT '第三方账号信息',
        avatar VARCHAR(255) DEFAULT NULL COMMENT '头像token',
        cover VARCHAR(50) DEFAULT NULL COMMENT '封面图片token',
        follower_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注我的人数',
        followee_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我关注的人数',
        following_article_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我关注的文章数',
        following_question_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我关注的问题数',
        following_topic_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我关注的话题数',
        article_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我发表的文章数量',
        question_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我发表的问题数量',
        answer_count INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '我发表的回答数量',
        notification_unread INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '未读通知数',
        inbox_unread INT (11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '未读私信数',
        headline VARCHAR(40) DEFAULT NULL COMMENT '一句话介绍',
        bio VARCHAR(160) DEFAULT NULL COMMENT '个人简介',
        company VARCHAR(255) DEFAULT NULL COMMENT '公司名称',
        location VARCHAR(255) DEFAULT NULL COMMENT '地址',
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        PRIMARY KEY (user_index),
        UNIQUE KEY user_id (user_id),
        KEY user_name (user_name),
        KEY email (email),
        KEY created_at (created_at)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

-- 创建文章分类表
CREATE TABLE
    IF NOT EXISTS article_categories (
        category_id INT AUTO_INCREMENT PRIMARY KEY,
        category_name VARCHAR(100) NOT NULL,
        parent_id INT NULL,
        -- 外键关联自身，用于多级分类，设置级联更新和删除
        FOREIGN KEY (parent_id) REFERENCES article_categories (category_id) ON UPDATE CASCADE ON DELETE SET NULL
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章表
CREATE TABLE
    IF NOT EXISTS articles (
        article_id VARCHAR(20) COMMENT '文章ID' PRIMARY KEY,
        title VARCHAR(255) NOT NULL COMMENT '文章标题',
        content TEXT NOT NULL COMMENT '文章内容',
        user_id VARCHAR(20) NOT NULL COMMENT '文章作者',
        category_id INT NOT NULL COMMENT '分类id',
        cover VARCHAR(200) NOT NULL COMMENT '文章封面',
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        view_count INT NOT NULL DEFAULT 0,
        is_published TINYINT (1) NOT NULL DEFAULT 0,
        -- 外键关联文章分类表，设置级联更新和删除
        FOREIGN KEY (category_id) REFERENCES article_categories (category_id) ON UPDATE CASCADE ON DELETE RESTRICT,
        -- 为作者 ID 和创建时间添加索引
        KEY user_id (user_id),
        UNIQUE KEY article_id (article_id),
        KEY created_at (created_at)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章标签表
CREATE TABLE
    IF NOT EXISTS article_tags (
        tag_id INT AUTO_INCREMENT PRIMARY KEY,
        tag_name VARCHAR(50) NOT NULL,
        UNIQUE KEY tag_name (tag_name)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章 - 标签关联表
CREATE TABLE
    IF NOT EXISTS article_tag_relations (
        relation_id INT AUTO_INCREMENT PRIMARY KEY,
        article_id VARCHAR(20) NOT NULL,
        tag_id INT NOT NULL,
        -- 外键关联文章表，设置级联更新和删除
        FOREIGN KEY (article_id) REFERENCES articles (article_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 外键关联文章标签表，设置级联更新和删除
        FOREIGN KEY (tag_id) REFERENCES article_tags (tag_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和标签 ID 添加联合索引
        UNIQUE KEY article_tag (article_id, tag_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章话题表
CREATE TABLE
    IF NOT EXISTS article_topics (
        topic_id INT AUTO_INCREMENT PRIMARY KEY,
        topic_name VARCHAR(50) NOT NULL,
        UNIQUE KEY topic_name (topic_name)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章 - 话题关联表
CREATE TABLE
    IF NOT EXISTS article_topic_relations (
        relation_id INT AUTO_INCREMENT PRIMARY KEY,
        article_id VARCHAR(20) NOT NULL,
        topic_id INT NOT NULL,
        -- 外键关联文章表，设置级联更新和删除
        FOREIGN KEY (article_id) REFERENCES articles (article_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 外键关联文章标签表，设置级联更新和删除
        FOREIGN KEY (topic_id) REFERENCES article_topics (topic_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和标签 ID 添加联合索引
        UNIQUE KEY article_topic (article_id, topic_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章评论表
CREATE TABLE
    IF NOT EXISTS article_comments (
        comment_id INT AUTO_INCREMENT PRIMARY KEY,
        article_id VARCHAR(20) NOT NULL,
        user_id VARCHAR(20) NOT NULL,
        content TEXT NOT NULL,
        create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        parent_comment_id INT NULL,
        -- 外键关联文章表，设置级联更新和删除
        FOREIGN KEY (article_id) REFERENCES articles (article_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 外键关联自身，用于多级评论，设置级联更新和删除
        FOREIGN KEY (parent_comment_id) REFERENCES article_comments (comment_id) ON UPDATE CASCADE ON DELETE SET NULL,
        -- 为文章 ID 和用户 ID 添加索引
        KEY article_id (article_id),
        KEY user_id (user_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章点赞表
CREATE TABLE
    IF NOT EXISTS article_like (
        -- 点赞记录 ID，自增主键
        like_id INT AUTO_INCREMENT PRIMARY KEY,
        -- 关联的文章 ID，外键，引用 articles 表的 article_id
        article_id VARCHAR(20) NOT NULL,
        -- 点赞用户的 ID，可以根据实际情况修改为合适的数据类型
        user_id VARCHAR(20) NOT NULL,
        -- 点赞时间，默认值为当前时间
        liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        -- 外键约束，确保 article_id 存在于 articles 表中，设置级联更新和删除
        FOREIGN KEY (article_id) REFERENCES articles (article_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和用户 ID 添加联合唯一索引，避免重复点赞
        UNIQUE KEY article_user (article_id, user_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 插入三条 article_categories 数据
INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('农学类', NULL);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('数学信息学类', NULL);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('水产医学', 1);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('兽医', 1);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('计算机科学与技术', 2);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('软件工程', 2);