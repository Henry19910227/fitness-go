CREATE TABLE IF NOT EXISTS course_training_avg_statistics (
    `id`        INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `course_id` INT(11) UNSIGNED COMMENT '課表 id',
    `rate`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '平均訓練率',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_course_training_avg_statistics_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;