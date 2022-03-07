CREATE TABLE IF NOT EXISTS user_plan_statistics (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '課表報表id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `plan_id`        INT(11) UNSIGNED COMMENT '計畫id',
    `finish_workout_count`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '完成訓練總數量(去除重複並累加)',
    `duration`       INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '總花費時間(秒)',
    `update_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_user_plan_statistics` (`user_id`,`plan_id`),
    CONSTRAINT fk_user_plan_statistics_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_plan_statistics_plan_id_to_plans_id FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;