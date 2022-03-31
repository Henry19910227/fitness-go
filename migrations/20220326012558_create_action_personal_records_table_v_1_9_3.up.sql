CREATE TABLE IF NOT EXISTS action_personal_records (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '課表報表id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `action_id`      INT(11) UNSIGNED COMMENT '動作id',
    `weight`         FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量(公斤)',
    `reps`           INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '次數',
    `distance`       FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '距離(公里)',
    `duration`       INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '時長(秒)',
    `incline`        FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '坡度',
    `update_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_action_personal_records` (`user_id`,`action_id`),
    CONSTRAINT fk_action_personal_records_action_id_to_actions_id FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;