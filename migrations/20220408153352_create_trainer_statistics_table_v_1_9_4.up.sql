CREATE TABLE IF NOT EXISTS trainer_statistics (
    `user_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `student_count`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '學生總數(訓練過該教練上架的課表)',
    `course_count`   INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '課表總數',
    `review_score`   FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '課表總評分',
    `update_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT `fk_trainer_statistics_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;