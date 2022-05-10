CREATE TABLE IF NOT EXISTS user_course_assets (
    `id`              INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '課表資產id',
    `order_id`        VARCHAR(20) COMMENT '訂單 id',
    `user_id`         INT(11) UNSIGNED COMMENT '用戶 id',
    `course_id`       INT(11) UNSIGNED COMMENT '課表 id',
    `available`       TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否可用(0:不可用/1:可用)',
    `create_at`       DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`       DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_user_course` (`user_id`,`course_id`),
    CONSTRAINT fk_purchase_courses_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_purchase_courses_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;