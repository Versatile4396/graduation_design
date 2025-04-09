-- 删除数据库（如果存在）
DROP DATABASE IF EXISTS appDB;

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
        nickname VARCHAR(20) NOT NULL COMMENT '昵称',
        overview VARCHAR(255) DEFAULT NULL COMMENT '简介',
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
        deleted TINYINT (1) NOT NULL DEFAULT 0 COMMENT '是否已删除',
        role TINYINT (1) NOT NULL DEFAULT 0 COMMENT '用户角色 0:普通用户 1:管理员 ',
        deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
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
        article_index INT (11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章序号' PRIMARY KEY,
        article_id VARCHAR(20) COMMENT '文章ID',
        title VARCHAR(255) NOT NULL COMMENT '文章标题',
        content TEXT NOT NULL COMMENT '文章内容',
        user_id VARCHAR(20) NOT NULL COMMENT '文章作者',
        category_id INT NOT NULL COMMENT '分类id',
        tag_id INT NOT NULL COMMENT '标签id',
        topic_id INT NOT NULL COMMENT '话题id',
        abstract VARCHAR(200) NOT NULL COMMENT '文章摘要',
        cover VARCHAR(200) NOT NULL COMMENT '文章封面',
        gpt_summarize VARCHAR(200) COMMENT 'gpt生成的文章摘要',
        view_count INT NOT NULL DEFAULT 0,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        article_status INT NOT NULL DEFAULT 0,
        -- 外键关联文章分类表，设置级联更新和删除
        FOREIGN KEY (category_id) REFERENCES article_categories (category_id) ON UPDATE CASCADE ON DELETE RESTRICT,
        -- 为作者 ID 和创建时间添加索引
        KEY user_id (user_id),
        UNIQUE KEY article_id (article_id),
        KEY created_at (created_at)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章标签表
CREATE TABLE
    IF NOT EXISTS tags (
        tag_id INT AUTO_INCREMENT PRIMARY KEY,
        tag_name VARCHAR(50) NOT NULL,
        tag_desc VARCHAR(200),
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
        FOREIGN KEY (tag_id) REFERENCES tags (tag_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和标签 ID 添加联合索引
        UNIQUE KEY article_tag (article_id, tag_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章话题表
CREATE TABLE
    IF NOT EXISTS topics (
        topic_id INT AUTO_INCREMENT PRIMARY KEY,
        topic_name VARCHAR(50) NOT NULL,
        topic_desc VARCHAR(200),
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
        FOREIGN KEY (topic_id) REFERENCES topics (topic_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和标签 ID 添加联合索引
        UNIQUE KEY article_topic (article_id, topic_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建文章评论表
CREATE TABLE
    IF NOT EXISTS article_comments (
        comment_index INT AUTO_INCREMENT PRIMARY KEY,
        comment_id VARCHAR(20) UNIQUE COMMENT '评论ID',
        article_id VARCHAR(20) NOT NULL,
        user_id VARCHAR(20) NOT NULL,
        content TEXT NOT NULL,
        create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        parent_comment_id VARCHAR(20) NULL,
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
    IF NOT EXISTS article_likes (
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

-- 创建文章收藏表
CREATE TABLE
    IF NOT EXISTS article_collections (
        -- 收藏记录 ID，自增主键
        collection_id INT AUTO_INCREMENT PRIMARY KEY,
        -- 关联的文章 ID，外键，引用 articles 表的 article_id
        article_id VARCHAR(20) NOT NULL,
        -- 收藏用户的 ID，可以根据实际情况修改为合适的数据类型
        user_id VARCHAR(20) NOT NULL,
        -- 收藏时间，默认值为当前时间
        collected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        -- 外键约束，确保 article_id 存在于 articles 表中，设置级联更新和删除
        FOREIGN KEY (article_id) REFERENCES articles (article_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 为文章 ID 和用户 ID 添加联合唯一索引，避免重复收藏
        UNIQUE KEY article_user (article_id, user_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建视频表
CREATE TABLE
    videos (
        video_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id VARCHAR(20) NOT NULL,
        title VARCHAR(255) NOT NULL,
        abstract VARCHAR(255) NOT NULL,
        upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        duration INT,
        views INT DEFAULT 0,
        likes INT DEFAULT 0,
        dislikes INT DEFAULT 0,
        tag_ids JSON,
        category VARCHAR(100),
        video_url VARCHAR(255) NOT NULL
    );

-- 信息表
DROP TABLE IF EXISTS `messages`;

CREATE TABLE
    IF NOT EXISTS `messages` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
        `created_at` datetime (3) DEFAULT NULL COMMENT '创建时间',
        `updated_at` datetime (3) DEFAULT NULL COMMENT '更新时间',
        `deleted_at` bigint unsigned DEFAULT NULL COMMENT '删除时间戳',
        `from_user_id` varchar(20) DEFAULT NULL COMMENT '发送人ID',
        `to_user_id` varchar(20) DEFAULT NULL COMMENT '发送对象ID',
        `content` varchar(2500) DEFAULT NULL COMMENT '消息内容',
        `url` varchar(350) DEFAULT NULL COMMENT '''文件或者图片地址''',
        `pic` text COMMENT '缩略图',
        `message_type` smallint DEFAULT NULL COMMENT '''消息类型：1单聊，2群聊''',
        `content_type` smallint DEFAULT NULL COMMENT '''消息内容类型：1文字，2语音，3视频''',
        PRIMARY KEY (`id`),
        KEY `idx_messages_deleted_at` (`deleted_at`),
        KEY `idx_messages_from_user_id` (`from_user_id`),
        KEY `idx_messages_to_user_id` (`to_user_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '消息表';

-- 创建互补帮学表
CREATE TABLE
    IF NOT EXISTS assistances (
        assistance_index INT (11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '互助序号' PRIMARY KEY,
        assistance_id VARCHAR(20) COMMENT '互助ID',
        title VARCHAR(255) NOT NULL COMMENT '互助标题',
        content TEXT NOT NULL COMMENT '互助内容',
        user_id VARCHAR(20) NOT NULL COMMENT '用户ID',
        category_id INT NOT NULL COMMENT '分类id',
        tag_id INT COMMENT '标签id',
        topic_id INT COMMENT '话题id',
        abstract VARCHAR(200) COMMENT '互助摘要',
        cover VARCHAR(200) NOT NULL COMMENT '互助封面',
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        assistance_status INT NOT NULL DEFAULT 0,
        -- 外键关联互助分类表，设置级联更新和删除
        FOREIGN KEY (category_id) REFERENCES article_categories (category_id) ON UPDATE CASCADE ON DELETE RESTRICT,
        UNIQUE KEY assistance_id (assistance_id),
        KEY created_at (created_at)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建互助模块评论表
CREATE TABLE
    IF NOT EXISTS assistance_comments (
        comment_index INT AUTO_INCREMENT PRIMARY KEY,
        comment_id VARCHAR(20) UNIQUE COMMENT '评论ID',
        assistance_id VARCHAR(20) NOT NULL,
        user_id VARCHAR(20) NOT NULL,
        content TEXT NOT NULL,
        create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        parent_comment_id VARCHAR(20) NULL,
        -- 外键关联文章表，设置级联更新和删除
        FOREIGN KEY (assistance_id) REFERENCES assistances (assistance_id) ON UPDATE CASCADE ON DELETE CASCADE,
        -- 外键关联自身，用于多级评论，设置级联更新和删除
        FOREIGN KEY (parent_comment_id) REFERENCES assistance_comments (comment_id) ON UPDATE CASCADE ON DELETE SET NULL,
        -- 为文章 ID 和用户 ID 添加索引
        KEY assistance_id (assistance_id),
        KEY user_id (user_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建user_friends表，用于存储用户之间的关注关系
DROP TABLE IF EXISTS `user_friends`;

CREATE TABLE
    IF NOT EXISTS `user_friends` (
        `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
        `created_at` datetime (3) DEFAULT NULL COMMENT '创建时间',
        `updated_at` datetime (3) DEFAULT NULL COMMENT '更新时间',
        `deleted_at` bigint unsigned DEFAULT NULL COMMENT '删除时间戳',
        `user_id` int DEFAULT NULL COMMENT '用户ID',
        `friend_id` int DEFAULT NULL COMMENT '好友ID',
        PRIMARY KEY (`id`),
        KEY `idx_user_friends_user_id` (`user_id`),
        KEY `idx_user_friends_friend_id` (`friend_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '好友信息表';

-- 插入 article_categories 数据
INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('计算机类', NULL);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('农学类', NULL);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('机械自动化类', null);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('电子信息类', null);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('外语学类', null);

INSERT INTO
    article_categories (category_name, parent_id)
VALUES
    ('管理学类', null);

INSERT into
    article_categories (category_name, parent_id)
values
    ('校园OKR', null);

INSERT INTO
    topics (topic_name, topic_desc)
values
    ('新人报道', '新人报道赛道');

-- 插入数学学习方法主题
INSERT INTO
    topics (topic_name, topic_desc)
VALUES
    ('数学学习方法', '探讨数学学科的高效学习方法与技巧');

-- 插入英语单词记忆主题
INSERT INTO
    topics (topic_name, topic_desc)
VALUES
    ('英语单词记忆', '分享英语单词的记忆方法和策略');

-- 插入编程学习经验主题
INSERT INTO
    topics (topic_name, topic_desc)
VALUES
    ('编程学习经验', '交流编程学习过程中的经验和遇到的问题');

-- 插入历史事件分析主题
INSERT INTO
    topics (topic_name, topic_desc)
VALUES
    ('历史事件分析', '对重要历史事件进行深入分析和解读');

-- 插入物理实验技巧主题
INSERT INTO
    topics (topic_name, topic_desc)
VALUES
    ('物理实验技巧', '分享物理实验中的操作技巧和注意事项');