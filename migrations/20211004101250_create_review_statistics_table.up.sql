CREATE TABLE IF NOT EXISTS review_statistics (
    `course_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '課表id',
    `score_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '評分累計',
    `amount` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '評分筆數',
    `five_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '五分總筆數',
    `four_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '四分總筆數',
    `three_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '三分總筆數',
    `two_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '二分總筆數',
    `one_total` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '一分總筆數',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_review_statistics_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;