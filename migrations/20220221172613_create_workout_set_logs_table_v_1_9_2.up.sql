CREATE TABLE IF NOT EXISTS workout_set_logs (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訂閱紀錄id',
    `workout_log_id` INT(11) UNSIGNED COMMENT '訓練歷史id',
    `workout_set_id` INT(11) UNSIGNED COMMENT '訓練組id',
    `weight`         FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量(公斤)',
    `reps`           INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '次數',
    `distance`       FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '距離(公里)',
    `duration`       INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '時長(秒)',
    `incline`        FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '坡度',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_workout_set_logs_workout_log_id_to_workout_logs_id FOREIGN KEY (workout_log_id) REFERENCES workout_logs(id) ON DELETE CASCADE,
    CONSTRAINT fk_workout_set_logs_workout_set_id_to_workout_sets_id FOREIGN KEY (workout_set_id) REFERENCES workout_sets(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;