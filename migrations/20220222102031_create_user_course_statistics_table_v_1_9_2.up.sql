CREATE TABLE IF NOT EXISTS user_course_statistics (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '課表報表id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `course_id`      INT(11) UNSIGNED COMMENT '課表id',
    `finish_workout_count`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '完成訓練總數量(去除重複並累加)',
    `total_finish_workout_count`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訓練總量(可重複並累加)',
    `duration`       INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '總花費時間(秒)',
    `update_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_user_course_statistics` (`user_id`,`course_id`),
    CONSTRAINT fk_user_course_statistics_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_course_statistics_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;