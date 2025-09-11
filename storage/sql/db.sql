CREATE TABLE `user`
(
    `id`                        BIGINT       NOT NULL AUTO_INCREMENT UNIQUE COMMENT '用户唯一ID',
    `address`                   VARCHAR(255) NOT NULL DEFAULT '' COMMENT '钱包地址',
    `user_name`                 VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar_url`                VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像链接',
    `points`                    BIGINT       NOT NULL DEFAULT 0 COMMENT '积分',
    `badges`                    JSON         NOT NULL DEFAULT '[]' COMMENT '用户徽章, jsn list: [1,2,3,4]',
    `sbt_token_id`              BIGINT       NOT NULL DEFAULT 0 COMMENT 'sbt nft token id',
    `task_status`               VARCHAR(300) NOT NULL COMMENT '任务状态, 一个长度300的字符串，0表示没有参与，1表示参与了。例如：00011111',
    `update_at`                 DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_at`                 DATETIME              DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_address` (`address`)
);

CREATE TABLE `user_points_log`
(
    `id`        BIGINT       NOT NULL AUTO_INCREMENT UNIQUE COMMENT '用户积分变化日志唯一ID',
    `user_id`   BIGINT       NOT NULL COMMENT '用户唯一ID',
    `total_points`   INT          NOT NULL COMMENT '积分变化',
    `base_points`    INT          NOT NULL COMMENT '积分变化',
    `extra_points`   INT          NOT NULL COMMENT '积分变化',
    `type`      VARCHAR(255) NOT NULL COMMENT   '积分类型',
    `update_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
);

CREATE TABLE `badge`
(
    `id`        BIGINT NOT NULL AUTO_INCREMENT UNIQUE COMMENT 'badge id',
    `name`      VARCHAR(255) NOT NULL COMMENT   '徽章名称',
    `image_url` VARCHAR(255) NOT NULL COMMENT   '徽章图片链接',
    `unlock_conditions` VARCHAR(255) NOT NULL COMMENT   '徽章解锁条件',
    `update_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
);

CREATE TABLE `user_daily_task_log`
(
    `id`        BIGINT       NOT NULL AUTO_INCREMENT UNIQUE COMMENT '用户积分变化日志唯一ID',
    `user_id`   BIGINT       NOT NULL COMMENT '用户唯一ID',
    `answer`    TEXT NOT NULL COMMENT   '用户每日回答',
    `current_day` INT        NOT NULL COMMENT '记录回答问题天数offset，从2025-09-10 开始是第一天',
    `badges`    JSON         NOT NULL DEFAULT '[]' COMMENT '本次交互用户获取的徽章， jsn list: [1,2,3,4]',
    `tx_hash`   VARCHAR(255) NOT NULL COMMENT   '交易hash',
    `points`    INT          NOT NULL COMMENT '积分变化',
    `extra_points`    INT          NOT NULL COMMENT '积分变化',
    `update_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_tx_hash` (`tx_hash`)

);

CREATE TABLE `questions`
(
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '问题标题',
    `content` text NOT NULL COMMENT '问题详细内容',
    `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '问题类型：1-单选题，2-多选题，3-填空题，4-问答题，5-判断题',
    `update_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
);