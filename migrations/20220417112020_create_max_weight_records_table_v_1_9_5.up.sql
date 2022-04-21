CREATE TABLE IF NOT EXISTS max_weight_records (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '最佳紀錄id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `action_id`      INT(11) UNSIGNED COMMENT '動作id',
    `weight`         FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量(公斤)',
    `update_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_max_weight_records` (`user_id`,`action_id`),
    CONSTRAINT fk_max_weight_records_action_id_to_actions_id FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE,
    CONSTRAINT fk_max_weight_records_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;