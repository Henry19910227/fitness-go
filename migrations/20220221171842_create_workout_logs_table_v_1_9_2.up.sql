CREATE TABLE IF NOT EXISTS workout_logs (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訂閱紀錄id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `workout_id`     INT(11) UNSIGNED COMMENT '訓練id',
    `duration`       INT(11) UNSIGNED COMMENT '訓練時長',
    `intensity`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)',
    `place`          TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '地點(0:未指定/1:住家/2:健身房/3:戶外)',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_workout_logs_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_workout_logs_workout_id_to_workouts_id FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;