CREATE TABLE IF NOT EXISTS banners (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `course_id`             INT(11) UNSIGNED COMMENT '課表id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `url`                   VARCHAR(500) NOT NULL DEFAULT '' COMMENT '連結',
    `image`                 VARCHAR(50) NOT NULL DEFAULT '' COMMENT '圖片',
    `type`                  TINYINT(1) NOT NULL DEFAULT '0' COMMENT '類型(1:課表/2:教練/3:訂閱/4:連結)',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_banners_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE SET NULL,
    CONSTRAINT fk_banners_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;