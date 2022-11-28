CREATE TABLE IF NOT EXISTS course_status_update_logs (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '紀錄 id',
    `course_id` INT(11) UNSIGNED COMMENT '課表 id',
    `course_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)',
    `comment` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '註解',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    CONSTRAINT fk_course_status_update_logs_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;