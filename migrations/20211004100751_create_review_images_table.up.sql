CREATE TABLE IF NOT EXISTS review_images (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '圖片id',
    `course_id` INT(11) UNSIGNED NOT NULL COMMENT '課表id',
    `user_id` INT(11) UNSIGNED NOT NULL COMMENT '用戶id',
    `image` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '圖片',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_review_images_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    CONSTRAINT fk_review_images_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;