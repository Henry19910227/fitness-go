CREATE TABLE IF NOT EXISTS feedbacks (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `version`               VARCHAR(50) NOT NULL DEFAULT '' COMMENT '軟體版本',
    `platform`              VARCHAR(50) NOT NULL DEFAULT '' COMMENT '平台(ios/android)',
    `os_version`            VARCHAR(50) NOT NULL DEFAULT '' COMMENT '軟體版本',
    `phone_model`           VARCHAR(50) NOT NULL DEFAULT '' COMMENT '手機機型',
    `body`                  VARCHAR(400) NOT NULL DEFAULT '' COMMENT '內文',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_feedbacks_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;