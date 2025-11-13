CREATE TABLE `file_info`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `file_id`      VARCHAR(64)     NOT NULL COMMENT '文件唯一ID (UUID)',
    `file_name`    VARCHAR(255)    NOT NULL COMMENT '原始文件名',
    `bucket`       VARCHAR(64)     NOT NULL COMMENT '存储桶',
    `object_name`  VARCHAR(255)    NOT NULL COMMENT '存储路径/对象名',
    `size`         BIGINT          NOT NULL COMMENT '文件大小（字节）',
    `content_type` VARCHAR(128)    NOT NULL COMMENT '文件类型（MIME）',
    `uploader`     BIGINT UNSIGNED     NOT NULL COMMENT '上传者用户ID',
    `upload_time`  BIGINT        NOT NULL COMMENT '上传时间',
    `hash`         VARCHAR(255)     NOT NULL COMMENT '文件哈希值',
    `description`  VARCHAR(255)     NOT NULL COMMENT '文件描述',
    `deleted_time` BIGINT        NOT NULL COMMENT '删除时间',
    `status`       TINYINT         NOT NULL DEFAULT 1 COMMENT '状态（1=正常，0=已删除）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_file_id` (`file_id`),
    KEY `idx_uploader` (`uploader`),
    KEY `idx_upload_time` (`upload_time`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='文件元数据表';
