CREATE TABLE IF NOT EXISTS reviews (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '評論id',
    `course_id` INT(11) UNSIGNED COMMENT '課表id',
    `user_id` INT(11) UNSIGNED COMMENT '用戶id',
    `score` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分數',
    `body` VARCHAR (400) NOT NULL DEFAULT '' COMMENT '頻論內容',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    UNIQUE KEY `review_unique` (`course_id`,`user_id`),
    CONSTRAINT fk_reviews_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_reviews_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;