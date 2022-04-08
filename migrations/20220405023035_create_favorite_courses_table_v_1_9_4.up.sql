CREATE TABLE IF NOT EXISTS favorite_courses (
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `course_id`      INT(11) UNSIGNED COMMENT '課表id',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    PRIMARY KEY (`user_id`,`course_id`),
    CONSTRAINT `fk_favorite_courses_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_favorite_courses_course_id_to_courses_id` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;